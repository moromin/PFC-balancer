# PFC-balancer
PFC-balancer is a microservice that allows you to create and retrieve information about foods and recipes based on them.

This service was created with reference to [mercari-microservices-example](https://github.com/mercari/mercari-microservices-example).

## Architecture
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

## Microservices
### Gateway
- Port: `4000`
- This is the only microservice which is facing the out side of the services and acts as just a proxy.
### Auth
- Port: `50051`
- This microservice is responsible for issuing the access token (JWT) for the user.
- This microservice also provides the ability to validate the Auth information contained in the request.
### User
- Port: `50052`
- This microservice is responsible for storing the user information to the database and providing it as APIs.
### Recipe
- Port: `50053`
- This microservice is responsible for storing the recipe information to the database and providing it as APIs.
### Food
- Port: `50054`
- This microservice is responsible for providing food information from database as APIs.
### DB
- Port: `50000`
- This microservice is responsible for storing the user, recipe and food information.

## Get started
``` bash
# Build each service Docker images and
# run containers based on these.
$ make
```

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

## Examples
### Register
```bash
curl -s -XPOST -d '{"email": "moromisterioso@example.com", "password": "secret"}' localhost:4000/auth/register | jq .
```
```json
{
  "user": {
    "id": "4",
    "email": "moromisterioso@example.com"
  }
}
```
### Login
```bash
curl -s -XPOST -d '{"email": "moromisterioso@example.com", "password": "secret"}' localhost:4000/auth/login | jq .
```
```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTAwMzgyMzksImlzcyI6ImdvLWdycGMtYXV0aC1zZXJ2aWNlIiwiSWQiOjQsIkVtYWlsIjoibW9yb21pc3Rlcmlvc29AZXhhbXBsZS5jb20ifQ.pCcaM2FqrAV6OI8kUYXrOOeu4ydaJR8vIrNCZgqaRoc"
}
```
or
```bash
TOKEN=$(curl -s -XPOST -d '{"email": "moromisterioso@example.com", "password": "secret"}' localhost:4000/auth/login | jq .token -r)
```
### Find food by ID
```bash
curl -s -XGET localhost:4000/foods/42 | jq .
```
```json
{
  "food": {
    "id": "42",
    "name": "こむぎ ［パン類］ ロールパン",
    "protein": 10.1,
    "fat": 9,
    "carbohydrate": 48.6,
    "category": "1"
  }
}
```
### Search foods
```
curl -s -XGET localhost:4000/foods/search/もやし | jq .
```
```json
{
  "foodList": [
    {
      "id": "809",
      "name": "（もやし類） アルファルファもやし 生",
      "protein": 1.6,
      "fat": 0.1,
      "carbohydrate": 2,
      "category": "6"
    },
    {
      "id": "810",
      "name": "（もやし類） だいずもやし 生",
      "protein": 3.7,
      "fat": 1.5,
      "carbohydrate": 2.3,
      "category": "6"
    },
    ...
  ]
}
```
### Create recipe
```bash
curl -s -XPOST -d '{"name": "Test Recipe", "food_amounts": [{"food_id": 333,"amount": 15.12}, {"food_id": 29, "amount": 35.3}], "procedures": ["fire", "shake, shake, shake", "EAAAAAT!!!"]}' -H "authorization: bearer $TOKEN" localhost:4000/recipes | jq .
```
```json
{
  "id": "1"
}
```

### Find Recipe by ID
```bash
curl -s XGET localhost:4000/recipes/1 | jq .
```
```json
{
  "recipe": {
    "id": "1",
    "recipeName": "Test Recipe",
    "foodPfcAmounts": [
      {
        "food": {
          "id": "333",
          "name": "だいず ［全粒・全粒製品］ 全粒 青大豆 国産 乾",
          "protein": 33.5,
          "fat": 19.3,
          "carbohydrate": 30.1,
          "category": "4"
        },
        "amount": 15.12
      },
      {
        "food": {
          "id": "29",
          "name": "こむぎ ［パン類］ 角形食パン 食パン",
          "protein": 8.9,
          "fat": 4.1,
          "carbohydrate": 46.4,
          "category": "1"
        },
        "amount": 35.3
      }
    ],
    "procedures": [
      "fire",
      "shake, shake, shake",
      "EAAAAAT!!!"
    ],
    "userName": "moromisterioso@example.com"
  }
}
```
## DB schema
```mermaid
erDiagram
    users ||--o{ recipes: ""
    recipes ||--o| procedures: ""
    recipes ||--o| recipe_food: ""
    foods ||--o{ recipe_food: ""

    users {
        integer id
        string email
        string password
    }

    recipes {
        integer id
        string name
        integer user_id
    }

    procedures {
        integer id
        string text
        integer recipe_id
    }

    recipe_food {
        integer id
        integer recipe_id
        integer food_id
        float amount
    }

    foods {
        integer id
        string name
        float protein
        float fat
        float carbohydrate
        integer category
    }
```

## Reference
- [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)
    - The `google` directory was copied from [googleapis](https://github.com/googleapis/googleapis) to generate stubs for the gPRC gateway with the `protoc` command.

- [Microservices in Go with gRPC, API Gateway, and Authentication — Part 1/2](https://levelup.gitconnected.com/microservices-with-go-grpc-api-gateway-and-authentication-part-1-2-393ad9fc9d30)

- [mercari-microservices-example](https://github.com/mercari/mercari-microservices-example)

- [日本食品標準成分表・資源に関する取組:文部科学省](https://www.mext.go.jp/a_menu/syokuhinseibun/)
