package login

//To understand how the structure of the [Client] - [Server] works, take a look into [login.proto] file.

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"testing"
	"time"
)

/*
Runner of the server and client. Since the server get block listening for communications,
we run in another thread using go routines.
In this test we run the server, we create the user in server and then we login.
*/
func TestCreateUserAndLogin(t *testing.T) {
	go createServer()
	createUser()
	log.Println("*********************************************")
	loginClient()
}

/*
In this test we do the same than previous test but async every operation making composition of channels
*/
func TestAsync(t *testing.T) {
	channelCreate := make(chan bool)
	channelLogin := make(chan bool)
	go createServer()
	go func(channelCreate chan bool) {
		createUser()
		channelCreate <- true
	}(channelCreate)
	go func(channelCreate chan bool, channelLogin chan bool) {
		a := <-channelCreate
		fmt.Printf("Create user finished %t \n", a)
		loginClient()
		channelLogin <- true
	}(channelCreate, channelLogin)
	fmt.Printf("Login finished %t \n", <-channelLogin)
}

/*
In this test we run the server, we try to login, and we receive an [UserNotFound] Error response.
*/
func TestErrorUserNotFound(t *testing.T) {
	go createServer()
	loginClient()
}

var timeout = false

/*
In this test we prove how the context with timeout works between client-server
*/
func TestErrorTimeout(t *testing.T) {
	timeout = true
	go createServer()
	createUser()
	loginClient()
}

//######################################//
//		  gRPC CLIENT
//#####################################//

/*
Run gRPC request against a server following the next steps:

* Create connection
* Create the client, associated with the type of service we define in the proto file. In this case using factory with format
	[New[service name]Client] providing a client with format [service name[Client]] which means AccountClient.
* Create the [LoginMessage] defined in proto file as transport message.
* Make the gRPC call using the client previously created pointing to [LoginUser] which is the function defied in the service
  in the proto file.
*/
func loginClient() {
	conn, err := createConnection()
	if conn != nil {
		defer conn.Close()
	}
	accountClient := NewAccountClient(conn)
	loginMessage := &LoginMessage{Username: "politrons", Password: "12345"}
	response, err := accountClient.LoginUser(context.Background(), loginMessage)
	if err != nil {
		log.Fatalf("Error in Login process: %s", err)
	}
	log.Printf("Response, User with info:")
	log.Printf("Name: %s", response.Name)
	log.Printf("Age: %s", response.Age)
	log.Printf("Sex: %s", response.Sex)
}

/*
Same steps than loginClient, but instead of invoke [LoginUser] with [LoginMessage] we invoke [CreateUser] with [CreateUserMessage].

For this communication we create a context with timeout of 500ms where the server must return the response.
*/
func createUser() {
	conn, err := createConnection()
	if conn != nil {
		defer conn.Close()
	}
	accountClient := NewAccountClient(conn)
	createUserMessage := &CreateUserMessage{Username: "politrons", UserMessage: &UserMessage{Name: "Paul", Age: "38", Sex: "Male"}}
	ctx, _ := context.WithTimeout(context.Background(), 500*time.Millisecond)
	response, err := accountClient.CreateUser(ctx, createUserMessage)
	if err != nil {
		log.Fatalf("Error in creation process: %s", err)
	}
	log.Printf("Response, User created with info:")
	log.Printf("Name: %s", response.Name)
	log.Printf("Age: %s", response.Age)
	log.Printf("Sex: %s", response.Sex)
}

/*
We create a new connection using [grpc.Dial] passing the port, and Security strategy, in this case
[WithInsecure]
*/
func createConnection() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(":1981", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	return conn, err
}

//######################################//
//		  gRPC SERVER
//#####################################//

/*
Start a gRPC server following next steps:

* Create a listener on TCP port 1981
* Create a server instance
* Create a gRPC server instance
* Attach the Service to the server using the class generate with the proto file. Format[Register[service]Server]
  Now all the communication from client that reach the ip/port of server it will be redirect to all extend methods,
  of Server[Account]
* Start the server using the [net.Listen] specifying the network protocol and port
*/
func createServer() {
	listener := createServerListener()
	server := Server{}
	grpcServer := grpc.NewServer()
	RegisterAccountServer(grpcServer, &server)
	startServer(grpcServer, listener)
}

/*
We use this function to create a Listener in a network protocol and port.
*/
func createServerListener() net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 1981))
	if err != nil {
		log.Fatalf("Error to listen: %v", err)
	}
	return lis
}

/*
Using the instance of the Server and the listener we run the server and we control the error in
the initialization.
*/
func startServer(grpcServer *grpc.Server, listener net.Listener) {
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error running server: %server", err)
	}
}

// Server represents the gRPC server
type Server struct {
	users map[string]User
}

/*
Implementation of the service [Account]-[CreateUser] defined in [login.proto] in case the user does not exist in database we return the error.
In this communication we receive a context from the client which specify the timeout before the server must return the call.

In this function to prove the timeout of the context, if the flag is true in a test we will sleep the process for a second.
*/
func (server *Server) CreateUser(ctx context.Context, message *CreateUserMessage) (*UserMessage, error) {
	log.Printf("Request to create user with username %s", message.Username)
	if timeout {
		time.Sleep(1 * time.Second)
	}
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
		return nil, ctx.Err()
	default:
		if server.users == nil {
			server.users = make(map[string]User)
		}
		user := User{message.UserMessage.Name, message.UserMessage.Age, message.UserMessage.Sex}
		server.users[message.Username] = user
		return message.UserMessage, nil
	}

}

/*
Implementation of the service [Account]-[LoginUser] defined in [login.proto] in case the user does not exist in database we return the error.
*/
func (server *Server) LoginUser(ctx context.Context, message *LoginMessage) (*UserMessage, error) {
	user, found := server.users[message.Username]
	if found {
		log.Printf("Request to login user with username %s", message.Username)
		return &UserMessage{
			Name: user.name,
			Age:  user.age,
			Sex:  user.sex,
		}, nil
	} else {
		return nil, UserNotFound{fmt.Sprintf("User %s not found", message.Username)}
	}
}

type User struct {
	name string
	age  string
	sex  string
}

type UserNotFound struct {
	Cause string
}

func (e UserNotFound) Error() string {
	return e.Cause
}
