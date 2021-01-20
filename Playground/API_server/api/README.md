## REST Client
https://marketplace.visualstudio.com/items?itemName=humao.rest-client


.httpまたは.restファイルを作成し、次のような形式でファイルを作成すると、APIへのリクエストが作成できる

```
### TEST
GET http://localhost:8080/hello-world

### List shopping items
GET http://localhost:8080/shopping-items

### Create new shopping item
POST  http://localhost:8080/shopping-items
Content-Type: application/json

{
    "name": "Pizza"
}

### Remove shopping item
DELETE http://localhost:8080/shopping-items/uuid
```