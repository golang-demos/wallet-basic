# Wallet Basic System

## Table Of Contents
- [Dependencies](#dependencies)
- [Modules](#modules)
- [Run](#run)
  * [Production Mode](#production-mode)
  * [Development Mode](#development-mode)
- [Test Run](#test-run)
    + [User Signup](#user-signup)
    + [User Login](#user-login)
    + [Session Details](#session-details)
    + [Deposit to Wallet](#deposit-to-wallet)
    + [Withdraw from Wallet](#withdraw-from-wallet)
    + [Check Wallet Balance](#check-wallet-balance)
    + [Get Wallet Statement](#get-wallet-statement)



## Dependencies
1. Go Language & Modules
2. Fiber : `go get github.com/gofiber/fiber`
3. MongoDB : `go get github.com/mongodb/mongo-go-driver/mongo`
4. Env : `github.com/joho/godotenv`


## Modules
- User Sign-In, Signup
- Wallet: Credit, Debit


## Run
### Development Mode
```
docker-compose -f docker-compose.dev.yml up
cd ./api
go run main.go
```
### Production Mode
```
docker-compose -f docker-compose.yml up --build
```


## Test Run


#### User Signup
<hr>

Details:
| Property       | Value            |
| -------------- | -----------      |
| API            | `/api/v1/signup` |
| Method         | `POST`           |
| Authentication | `Not-Required`   |
| Description    | This API is for registering a new user with name, mobile and password.  |


Input:

	name     : User's name
	mobile   : User's mobile. Unique.
	password : Password for login


Sample Request:
```
curl --location --request POST "http://localhost:3000/api/v1/signup" \
--header 'Content-Type: application/json' \
--data-raw '{
	"name": "Vinay Jeurkar",
	"mobile": "9766123123",
	"password": "12345"
}'
```

Sample Response:
```json
{
    "success": true,
    "user": {
        "name": "Vinay Jeurkar",
        "mobile": "9766123123",
        "role": "user"
    }
}
```

#### User Login
<hr>

Details:
| Property       | Value            |
| -------------- | -----------      |
| API            | `/api/v1/session/login` |
| Method         | `POST`           |
| Authentication | `Not-Required`   |
| Description    | This API is for logging in using the credentials provided at the time of signup. |

Input: 

	mobile   : User's mobile number
	password : User's password

Sample Request:
```
curl --location --request POST "http://localhost:3000/api/v1/session/login" \
--header 'Content-Type: application/json' \
--data-raw '{
	"mobile": "9766123123",
	"password": "12345"
}'
```

Sample Response:
```json
{
    "isLoggedIn": true,
    "token": "XVlBzgbaiCMRAjWwhTHctcuA"
}
```


`NOTE:` After successful login, this API provides a `Token` in response which should be put as a value of `SESS-TOKEN` HTTP Header in all the rest of the APIs to indicate the session of a logged in user.

#### Session Details
<hr>

Details:
| Property       | Value            |
| -------------- | -----------      |
| API            | `/api/v1/signup` |
| Method         | `GET`           |
| Authentication | `Optional`   |
| Description    | This API is used for fetching currently logged in user's details. This information is derieved from the value of `SESS-TOKEN` HTTP Header. |

Sample Request: (Please substitute `<SESS-TOKEN>` with right value)
```
curl --location --request GET "http://localhost:3000/api/v1/session/get" \
--header 'SESS-TOKEN: <SESS-TOKEN>'
```

Sample Response:
```json
{
    "LoggedIn": true,
    "User": {
        "name": "Vinay Jeurkar",
        "mobile": "9766123123",
        "role": "user"
    }
}
```

#### Deposit to Wallet
<hr>

Details:
| Property       | Value            |
| -------------- | -----------      |
| API            | `/api/v1/wallet/make` |
| Method         | `POST`           |
| Authentication | `Required`   |
| Description    | This API is used for adding money to user's wallet. It also gives updated balance in response. |

Input: 

	trans_type : `credit` to indicate deposit transaction.
	amount     : Amount in float

Sample Request:
```
curl --location --request POST "http://localhost:3000/api/v1/wallet/make" \
--header 'SESS-TOKEN: <SESS-TOKEN>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "trans_type": "credit",
    "amount": 500
}'
```

Sample Response:
```json
{
    "balance": 500,
    "success": true
}
```


#### Withdraw from Wallet
<hr>

Details:
| Property       | Value            |
| -------------- | -----------      |
| API            | `/api/v1/wallet/make` |
| Method         | `POST`           |
| Authentication | `Required`   |
| Description    | This API is used for withdrawing money from user's wallet. It also gives updated balance in response. |

Input: 

	trans_type : `debit` to indicate withdraw transaction.
	amount     : Amount in float

Sample Request:
```
curl --location --request POST "http://localhost:3000/api/v1/wallet/make" \
--header 'SESS-TOKEN: <SESS-TOKEN>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "trans_type": "debit",
    "amount": 57
}'
```

Sample Response:
```json
{
    "balance": 500,
    "success": true
}
```


#### Check Wallet Balance
<hr>

Details:
| Property       | Value            |
| -------------- | -----------      |
| API            | `/api/v1/wallet` |
| Method         | `GET`           |
| Authentication | `Required`   |
| Description    | This API is used for withdrawing money from user's wallet. It also gives updated balance in response. |

Sample Request:
```
curl --location --request GET "http://localhost:3000/api/v1/wallet" \
--header 'SESS-TOKEN: <SESS-TOKEN>'
```

Sample Response:
```json
{
    "balance": 500,
    "success": true
}
```

#### Get Wallet Statement
<hr>

Details:
| Property       | Value            |
| -------------- | -----------      |
| API            | `/api/v1/wallet/statement` |
| Method         | `GET`           |
| Authentication | `Required`   |
| Description    | This API is used for fetching complete statement of wallet transactions. |

Sample Request:
```
curl --location --request GET "http://localhost:3000/api/v1/wallet/statement" \
--header 'SESS-TOKEN: <SESS-TOKEN>'
```

Sample Response:
```json
[
    {
        "amount": 500,
        "created_at": "2021-12-18 13:38:09 +0530 IST",
        "type": "credit"
    },
    {
        "amount": 22,
        "created_at": "2021-12-18 10:53:01 +0530 IST",
        "type": "debit"
    },
    {
        "amount": 57,
        "created_at": "2021-12-18 10:52:36 +0530 IST",
        "type": "debit"
    },
    {
        "amount": 250,
        "created_at": "2021-12-18 10:52:26 +0530 IST",
        "type": "credit"
    }
]
```


## Author
### Vinay Jeurkar

<p>
  <a href="https://www.linkedin.com/in/vinay-jeurkar/" rel="nofollow noreferrer" style="text-decoration:none;"><img src="https://img.shields.io/badge/LinkedIn-0077B5?style=flat&logo=linkedin&logoColor=white" /></a> 
	&nbsp; 
  <a href="https://github.com/vinay03" rel="nofollow noreferrer" style="text-decoration:none;"><img src="https://img.shields.io/badge/GitHub-100000?style=flat&logo=github&logoColor=white" /></a> 
	&nbsp; 
  <a href="https://twitter.com/Vinay_Jeurkar" rel="nofollow noreferrer" style="text-decoration:none;"><img src="https://img.shields.io/badge/Twitter-1DA1F2?style=flat&logo=twitter&logoColor=white" /></a>
</p>
