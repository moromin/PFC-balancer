# PFC-balancer
PFC-balancer is a microservice that allows you to create and retrieve information about foods and recipes based on them.

This service was created with reference to [mercari-microservices-example](https://github.com/mercari/mercari-microservices-example).

```mermaid
graph LR;
    A(gateway) --> B(auth)
    B --> C(user)
    A --> D(recipe)
    D --> B
    D --> C
    D --> E(food)
    A --> E
```

## Get started
``` bash
# start up postgres
docker compose up -d

# run each services
make
```

## Microservices
### Gateway
- Port: `localhost:4000`
### Auth
- Port: `localhost:50051`
### Users
- Port: `localhost:50052`
### Recipe
- Port: `localhost:50054`
### Food
- Port: `localhost:50055`
### DB
- Port: `localhost:5000`

## Endpoints
Specify the following path followed by `localhost:4000`
### Auth
| Service | Method | Endpoint       | Auth |
|---------|--------|----------------|------|
| Resister  | `POST` | `/auth/register/` | × |
| Login  | `POST` | `/auth/login/` | × |
### Food
| Service | Method | Endpoint       | Auth |
|---------|--------|----------------|------|
| List all foods  | `GET` | `/foods` | × |
| Find food by ID | `GET` | `/foods/{id}` | × |
| Search foods  | `GET` | `/foods/search/{name}` | × |
### Recipe
| Service | Method | Endpoint       | Auth |
|---------|--------|----------------|------|
| Create new recipe  | `POST` | `/recipes` | ✔︎ |
| List all recipes  | `GET` | `/recipes` | × |
| Find recipe by ID  | `GET` | `/recipes/{id}` | × |

## Reference
- [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)
    - The `google` directory was copied from [googleapis](https://github.com/googleapis/googleapis) to generate stubs for the gPRC gateway with the `protoc` command.

- [Microservices in Go with gRPC, API Gateway, and Authentication — Part 1/2](https://levelup.gitconnected.com/microservices-with-go-grpc-api-gateway-and-authentication-part-1-2-393ad9fc9d30)

- [mercari-microservices-example](https://github.com/mercari/mercari-microservices-example)

- [日本食品標準成分表・資源に関する取組:文部科学省](https://www.mext.go.jp/a_menu/syokuhinseibun/)
