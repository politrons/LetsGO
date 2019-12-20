# ![My image](../../../../../img/grpc.png)  ![My image](../../../../../img/goGRPC.jpg)    

### Create server and transport messages source

Create your own [proto file](login/login.proto) where you define the transport messages and contract functions of server.
Having this file you will have a contract between Client-Server to be sure that communication between them remain unaltered.
You can consider this contract file, like an API Rest contract, where if you want to alter the contract between publisher-consumer you, as a publisher must version the file,
and distribute to all your consumers, in order to generate the new source version for communications.

Once you have your contract defined, then execute the `protoc` command, which it will generate the gRPC.go file with all implementation, to be used from client and server.

```
protoc -I login/ login/login.proto --go_out=plugins=grpc:login

```

Here `login` is the directory where you have the contract, and where you want to create the sources.

### Account example

Here you can see the generated source by `Protoc` using the [proto file](login/login.proto) to link Client and Server.

* **[Server-Messages](login/login.pb.go)**

And here the **Account** example implementation of Client and Server that you need to provide using gRPC.

* **[Client-Server](login/ClientServerGRPC_test.go)**