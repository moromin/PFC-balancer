# PFC-balancer

```mermaid
graph LR;
    A(gateway) --> B(auth)
    A --> C(recipe)
    A --> E(food)
    B --> D(user)
    C --> D
    C --> E
```
## Microservices
### Gateway
- Port: `localhost:4000`
### Auth
- Port: `localhost:50051`
### Users
- Port: `localhost:50054`
### Recipe
### Food
- Port: `localhost:50052`
### DB
- Port: `localhost:5000`

## Reference
### gRPC Gateway
The `google` directory was copied from [googleapis](https://github.com/googleapis/googleapis) to generate stubs for the gPRC gateway with the `protoc` command.
