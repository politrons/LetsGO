#  ![My image](../../img/REST.png)  ![My image](../../img/kafkalogo.jpg) ![My image](../../img/grpc.png)   

### Communications

An example of how three services, can communicate between each other using `Rest`, `gRPC` and `Kafka events`.

The flow of the program
````
Client ---> [REST request] ---> [KAKA PUBLISHER] ---> [KAFKA CONSUMER] ---> [gRPC CLIENT]

       ---> [gRPC SERVER]  ---> [KAKA PUBLISHER] ---> [KAFKA CONSUMER] ---> [REST response] ---> Client
````

The logs of the transaction show:

```
2019/07/07 16:58:48 REST request: hello world from rest 
2019/07/07 16:58:53 KAFKA Publisher: hello world from rest and Kafka publisher 
2019/07/07 16:58:54 KAFKA Consumer: hello world from rest and Kafka publisher and Kafka consumer 
2019/07/07 16:58:54 gRPC Client:hello world from rest and Kafka publisher and Kafka consumer and gRPC client 
2019/07/07 16:58:54 gRPC Server:hello world from rest and Kafka publisher and Kafka consumer and gRPC client and gRPC server 
2019/07/07 16:58:54 KAFKA Publisher: hello world from rest and Kafka publisher and Kafka consumer and gRPC client and gRPC server and Kafka publisher 
2019/07/07 16:58:55 KAFKA Consumer: hello world from rest and Kafka publisher and Kafka consumer and gRPC client and gRPC server and Kafka publisher and Kafka consumer 
2019/07/07 16:58:55 REST response: hello world from rest and Kafka publisher and Kafka consumer and gRPC client and gRPC server and Kafka publisher and Kafka consumer 
2019/07/07 16:58:55 ############################################################################################
2019/07/07 16:58:55 End of transaction with Message:
2019/07/07 16:58:55 HELLO WORLD FROM REST AND KAFKA PUBLISHER AND KAFKA CONSUMER AND GRPC CLIENT AND GRPC SERVER AND KAFKA PUBLISHER AND KAFKA CONSUMER
2019/07/07 16:58:55 ############################################################################################

```

The response to the client show:

```
"HELLO WORLD FROM REST AND KAFKA PUBLISHER AND KAFKA CONSUMER AND GRPC CLIENT AND GRPC SERVER AND KAFKA PUBLISHER AND KAFKA CONSUMER"
```