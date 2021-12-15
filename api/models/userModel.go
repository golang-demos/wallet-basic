package models

import (
	"context"
	"log"
	"math/rand"

	"crypto/md5"
	"encoding/hex"

	"github.com/go-playground/validator/v10"
	"github.com/golang-demos/ecommerce-basic/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name" validate:"required,min=2,max=32" bson="name"`
	Mobile   string             `json:"mobile" validate:"required,min=2,max=32" bson="mobile"`
	Role     string             `json:"role" validate:"required,min=2,max=16" bson="role"`
	Password string             `json:"password" validate:"required,min=2,max=16" bson="password"`
	Token    string             `json:"token" validate:"required,min=24,max=24" bson="token"`
}

type UserSignupData struct {
	Name     string `json:"name" validate:"required,min=2,max=32"`
	Mobile   string `json:"mobile" validate:"required,min=2,max=32"`
	Password string `json:"password" validate:"required,min=2,max=16"`
}

type UserLoginData struct {
	Mobile   string `json:"mobile" validate:"required,min=2,max=32"`
	Password string `json:"password" validate:"required,min=2,max=16"`
}

func (u *User) Init(postData UserSignupData) *User {
	if u.Role == "" {
		u.Role = "user"
	}
	if u.Token == "" {
		u.Token = generateRandomString(24)
	}
	u.Name = postData.Name
	u.Mobile = postData.Mobile
	u.Password = postData.Password
	return u
}

func (u *User) ToShort() User {
	var user User
	user.ID = u.ID
	user.Name = u.Name
	user.Mobile = u.Mobile
	user.Role = u.Role
	return user
}

func (u *User) Print() {
	msg := "\nID     : " + u.ID.String()
	msg += "\nName   : " + u.Name
	msg += "\nMobile : " + u.Mobile
	msg += "\nRole   : " + u.Role
	log.Print(msg)
}

func (u *User) Validate() []*ErrorResp {
	return validateModelsForErrors(validator.New().Struct(u))
}

func (u *UserLoginData) Validate() []*ErrorResp {
	return validateModelsForErrors(validator.New().Struct(u))
}

func generateRandomString(expectedLength int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, expectedLength)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func checkIfExists(user *User) bool {
	ctx := context.Background()
	var foundUser User
	err := database.UserCollection.FindOne(ctx, bson.D{{"mobile", user.Mobile}}).Decode(&foundUser)
	return err == nil
}

func CreateUser(user *User) bool {
	alreadyExists := checkIfExists(user)
	if alreadyExists {
		return false
	}
	hash := md5.Sum([]byte(user.Password))
	user.Password = hex.EncodeToString(hash[:])
	result, _ := database.UserCollection.InsertOne(context.Background(), user)
	if objID, ok := result.InsertedID.(primitive.ObjectID); ok {
		user.ID = objID
	}
	isCreated := false
	if result != nil {
		isCreated = true
		CreateWallet(user)
	}
	return isCreated
}

func Login(mobile, password string) (bool, string) {

	hash := md5.Sum([]byte(password))

	var foundUser User
	err := database.UserCollection.FindOne(context.Background(), bson.D{{"mobile", mobile}, {"passord", hash}}).Decode(&foundUser)
	if err != nil {
		log.Print(err)
	}
	return false, ""
}
