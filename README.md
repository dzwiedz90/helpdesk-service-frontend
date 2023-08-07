# Helpdesk Service Frontend

#### Helpdesk is test project of helpdesk page using microservices

## Helpdesk Service Frontend

- Build with go 1.20.2
- Uses the [Protocol Buffers](https://protobuf.dev/)
- Uses [gRPC](https://grpc.io/)
- Uses the [chi router](https://github.com/go-ch/chi) for handling incoming public HTTP requests

-------------
-------------
### Helpdesk has below services:

- [Proto](https://github.com/dzwiedz90/helpdesk-proto)
- [service-agents]() - to manage agents
- [service-frontend](https://github.com/dzwiedz90/helpdesk-service-frontend) - to serve frontend UI of application and communicate with other services when necessary
- [service-notifications]() - to manage and send notification about events for tickets
- [service-tickets]() - to manage tickets
- [service-users](https://github.com/dzwiedz90/helpdesk-service-users) - to manage users

-------------
-------------
## Endpoints
### Users
---
### POST /users/user.create
Endpoint used to create a new user</br>
Request:
```json
{
    "username": "james.kirk",
    "password": "password",
    "email": "james.kirk@helpdesk.com",
    "firstName": "James",
    "lastName": "Kirk",
    "age": 32,
    "gender": "male",
    "address": {
        "street": "Teststreet 8",
        "city": "Somecity",
        "postalCode": "2137",
        "Country": "USA"
    }
}
```
Response:
```json
{
    "code": 200,
    "message": "User created",
    "createUserResponse": {
        "id": 1
    }
}
```