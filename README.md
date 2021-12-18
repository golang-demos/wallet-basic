### Ecommerce Basic Example using

1. Go Language
2. Mux Router : `go get github.com/gorilla/mux`
3. MongoDB : `go get github.com/mongodb/mongo-go-driver/mongo`
4. Env : `github.com/joho/godotenv`


## Modules
- User Sign-In, Signup
- Wallet: Credit, Debit


## Run
### Production Mode
```
docker-compose -f docker-compose.yml up
```
### Development Mode
```
docker-compose -f docker-compose.dev.yml up
```

### Test Run


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


Example:
```
curl --location --request POST "http://localhost:3000/api/v1/signup" \
--header 'Content-Type: application/json' \
--data-raw '{
	"name": "Vinay Jeurkar",
	"mobile": "9766123123",
	"password": "12345"
}'
```

Response:
```json
{
    "success": true,
    "user": {
        "name": "Vinay Jeurkar",
        "mobile": "9766941950",
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

Example:
```
curl --location --request POST "http://localhost:3000/api/v1/session/login" \
--header 'Content-Type: application/json' \
--data-raw '{
	"mobile": "9766123123",
	"password": "12345"
}'
```

Response:
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

Example: (Please substitute `<SESS-TOKEN>` with right value)
```
curl --location --request GET "http://localhost:3000/api/v1/session/get" \
--header 'SESS-TOKEN: <SESS-TOKEN>'
```

Response:
```json
{
    "LoggedIn": true,
    "User": {
        "name": "Vinay Jeurkar",
        "mobile": "9766941956",
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

Example:
```
curl --location --request POST "http://localhost:3000/api/v1/wallet/make" \
--header 'SESS-TOKEN: <SESS-TOKEN>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "trans_type": "credit",
    "amount": 500
}'
```

Response:
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

Example:
```
curl --location --request POST "http://localhost:3000/api/v1/wallet/make" \
--header 'SESS-TOKEN: <SESS-TOKEN>' \
--header 'Content-Type: application/json' \
--data-raw '{
    "trans_type": "debit",
    "amount": 57
}'
```

Response:
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

Example:
```json
curl --location --request GET "http://localhost:3000/api/v1/wallet" \
--header 'SESS-TOKEN: <SESS-TOKEN>'
```

Response:
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

Example:
```json
curl --location --request GET "http://localhost:3000/api/v1/wallet/statement" \
--header 'SESS-TOKEN: <SESS-TOKEN>'
```

Response:
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
