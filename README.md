### Ecommerce Basic Example using
1. Go Language
2. Mux Router : `go get github.com/gorilla/mux`
3. MongoDB : `go get github.com/mongodb/mongo-go-driver/mongo`
4. Env : `github.com/joho/godotenv`


## Ecommerce Modules
- User Sign-In, Signup
- Wallet: Credit, Debit
- Products & variations
- Orders: Create, Update and Cancel


## Run
### Production Mode
```
docker-compose -f docker-compose.yml up
```
### Development Mode
```
docker-compose -f docker-compose.dev.yml up
```