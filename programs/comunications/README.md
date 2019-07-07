#  ![My image](../../img/REST.png)  ![My image](../../img/kafkalogo.jpg) ![My image](../../img/grpc.png)   

### Communications

An example how we can communicate programs using `Rest`, `gRPC` and `Kafka events`.

The flow of the program
````
Client ---> [REST request] ---> [KAKA PUBLISHER] ---> [KAFKA CONSUMER] ---> [gRPC CLIENT]

       ---> [gRPC SERVER]  ---> [KAKA PUBLISHER] ---> [KAFKA CONSUMER] ---> [REST response] ---> Client
````

The logs of the transaction show:

```
2019/07/07 16:45:43 KAFKA Publisher: hello world from rest and Kafka publisher 
2019/07/07 16:45:44 KAFKA Consumer: hello world from rest and Kafka publisher and Kafka consumer 
2019/07/07 16:45:44 gRPC Client:hello world from rest and Kafka publisher and Kafka consumer and gRPC client 
2019/07/07 16:45:44 gRPC Server:hello world from rest and Kafka publisher and Kafka consumer and gRPC client and gRPC server 
2019/07/07 16:45:44 KAFKA Publisher: hello world from rest and Kafka publisher and Kafka consumer and gRPC client and gRPC server and Kafka publisher 
2019/07/07 16:45:45 KAFKA Consumer: hello world from rest and Kafka publisher and Kafka consumer and gRPC client and gRPC server and Kafka publisher and Kafka consumer 
2019/07/07 16:45:45 #####################################
2019/07/07 16:45:45 End of transaction with Message:
2019/07/07 16:45:45 HELLO WORLD FROM REST AND KAFKA PUBLISHER AND KAFKA CONSUMER AND GRPC CLIENT AND GRPC SERVER AND KAFKA PUBLISHER AND KAFKA CONSUMER
2019/07/07 16:45:45 #####################################

```

The output of the flow must response

```
"HELLO WORLD FROM REST AND KAFKA PUBLISHER AND KAFKA CONSUMER AND GRPC CLIENT AND GRPC SERVER AND KAFKA PUBLISHER AND KAFKA CONSUMER"
```