# PFC-balancer

```mermaid
graph LR;
    A(gateway) --> B(auth)
    A(gateway) --> C(recipe)
    B --> D(user)
    C --> D
    C --> E(food)
```
## Microservices
### Gateway
### Auth
### Users
- Port: `localhost:50054`
### Recipe
### Food
### DB
- Port: `localhost:5000`

## Reference
### gRPC Gateway
The `google` directory was copied from [googleapis](https://github.com/googleapis/googleapis) to generate stubs for the gPRC gateway with the `protoc` command.
