# Helpdesk Service Frontend

#### Helpdesk is test project of helpdesk page using microservices

## Helpdesk Service Frontend

- Build with go 1.20.2
- Uses the [Protocol Buffers](https://protobuf.dev/)
- Uses [gRPC](https://grpc.io/)
- Uses the [chi router](https://github.com/go-ch/chi) for handling incoming public HTTP requests

-------------
### Helpdesk has below services:

- [Proto](https://github.com/dzwiedz90/helpdesk-proto)
- [service-agents]() - to manage agents
- [service-frontend](https://github.com/dzwiedz90/helpdesk-service-frontend) - to serve frontend UI of application and communicate with other services when necessary
- [service-notifications]() - to manage and send notification about events for tickets
- [service-tickets]() - to manage tickets
- [service-users](https://github.com/dzwiedz90/helpdesk-service-users) - to manage users

-------------
### Configuration before first run
- git pull origin master
- set up Postgres database
- create .env file and fill it with information as in the example below which will be loaded to the app's config:
```
HTTPAddress=0.0.0.0
HTTPPort=8080
Timeout=5
UsersGRPCPort=5002
UsersGRPCAddress=0.0.0.0
```
- run app using run.sh file or with command ```go build -o helpdesk-service-frontend main.go && ./helpdesk-service-frontend | tee logs/console.log```

-------------
### Endpoints
---
---
## USER

### POST /users/user.create
Endpoint used to create a new user</br>
Example request:
```json
{
    "username": "mr.spock",
    "password": "k1rkSuckZ",
    "email": "mr.spock@federation.com",
    "firstName": "Mr",
    "lastName": "Spock",
    "age": 89,
    "gender": "male",
    "address": {
        "street": "Valcanioan 69",
        "city": "Voclanium",
        "postal_code": "2137",
        "country": "Volcan"
    }
}
```
Example response:
```json
{
    "code": 200,
    "message": "User created",
    "createUserResponse": {
        "id": 1
    }
}
```

---
### GET /users/user.get
Endpoint used to get user by id</br>
Example request:
```json
{
     "id": 12
}
```
Example response:
```json
{
    "code": 200,
    "message": "",
    "getUserResponse": {
        "user": {
            "username": "mr.spock",
            "email": "mr.spock@federation.com",
            "firstName": "Mr",
            "lastName": "Spock",
            "age": 89,
            "gender": "male",
            "address": {
                "street": "Valcanioan 69",
                "city": "Voclanium",
                "country": "Volcan"
            }
        }
    }
}
```

---
### GET /users/user.get_all
Endpoint used to get all users</br>
Example response:
```json
{
    "code": 200,
    "message": "",
    "getAllUsersResponse": {
        "users": [
            {
                "user_id": 1,
                "username": "james.kirk",
                "email": "james.kirk@federation.com",
                "firstName": "James",
                "lastName": "Kirk",
                "age": 32,
                "gender": "male",
                "address": {
                    "street": "Spocksuckz 2137",
                    "city": "Somecity",
                    "country": "Earth"
                }
            },
            {
                "user_id": 2,
                "username": "mr.spock",
                "email": "mr.spock@federation.com",
                "firstName": "Mr",
                "lastName": "Spock",
                "age": 99,
                "gender": "male",
                "address": {
                    "street": "Kirksuckz 69",
                    "city": "Somecity",
                    "country": "Volcan"
                }
            }
        ]
    }
}
```