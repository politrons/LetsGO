
![My image](../../img/basket.png)    
 # Customer order service
  
Examples of a real Go Http service with Idempotent API, that use design patterns as Dependency Injection, CQRS, DDD,
Fail fast with Circuit breaker and NIO using go routines and channel.

### DDD Layers

* **[Application](src/application)**
* **[Domain](src/domain)**
* **[infrastructure](src/infrastructure)**

### Run

Just go to the file [BasketRunner_test.go](src/BasketRunner_test.go) and run the **Test** method.

Then just get a curl or postman and consume the API

* **Create order**

    **Request**
    
        POST:http://localhost:8080/order/create/
        
    **Response**
    
        {
            "OrderId": {
                "Id": "a5c2b3a2-74e9-48f8-9d84-13c9c6547564"
            },
            "Products": {},
            "TotalPrice": 0
        }

* **Find order**, using as uri param the orderId previously created.

    **Request**
    
        GET:http://localhost:8080/order/a5c2b3a2-74e9-48f8-9d84-13c9c6547564

    **Response**
    
        {
            "OrderId": {
                "Id": "a5c2b3a2-74e9-48f8-9d84-13c9c6547564"
            },
            "Products": {},
            "TotalPrice": 0
        }


* **Add product**, using as uri param the orderId previously created.

    **Request**
    
        POST:http://localhost:8080/order/addProduct/a5c2b3a2-74e9-48f8-9d84-13c9c6547564
   
    **Body**
    
        {
            "ProductId": "78844a80-a54a-443d-a062-26abe1462387",
            "ProductDescription":"7-up",
            "Price":1.95
        }
    
    **Response**
    
      {
          "OrderId": {
              "Id": "a5c2b3a2-74e9-48f8-9d84-13c9c6547564"
          },
          "Products": {
              "78844a80-a54a-443d-a062-26abe1462387": {
                  "Price": 1.95,
                  "Description": "7-up"
              }
          },
          "TotalPrice": 1.95
      }
  
* **Remove product**, using as uri param the orderId previously created.

    **Request**
    
        POST:http://localhost:8080/order/removeProduct/a5c2b3a2-74e9-48f8-9d84-13c9c6547564        
        
    **Body**
        
          {"ProductId":"78844a80-a54a-443d-a062-26abe1462387"}
               
    **Response**
    
          {
              "OrderId": {
                  "Id": "e4f16f9e-9df2-456c-8c4c-193761e675a9"
              },
              "Products": {},
              "TotalPrice": 0
          }
