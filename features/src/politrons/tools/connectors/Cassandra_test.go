package connectors

/*
I have seen far by standing on the shoulders of giants.
Example implemented in top of [cluster "github.com/gocql/gocql"]
*/
import (
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"testing"
)

// We embrace TDD so we create a specific Type per attribute in our program
type Name struct{ value string }
type Sex struct{ value string }
type Age struct{ value string }

type User struct {
	id   gocql.UUID
	name Name
	age  Age
	sex  Sex
}

type ResponseUsers struct {
	users []User
	error error
}

type ResponseUser struct {
	user  User
	error error
}

func TestCassandra(t *testing.T) {
	session := createCassandraSession()
	defer session.Close()

	insertQuery(session)

	selectChannel := make(chan ResponseUser)
	selectQuery(session, User{}, "Politrons", selectChannel)
	responseUser := <-selectChannel
	if responseUser.error != nil {
		log.Fatal(responseUser.error, nil)
	}
	fmt.Println("Full User:", responseUser.user.id, responseUser.user.name, responseUser.user.age, responseUser.user.sex)

	selectAllChannel := make(chan ResponseUsers)
	selectAllQuery(session, User{}, selectAllChannel)
	responseUsers := <-selectAllChannel
	if responseUser.error != nil {
		log.Fatal(responseUsers.error, nil)
	}
	for _, user := range responseUsers.users {
		fmt.Println("User:", user.id, user.name, user.age, user.sex)
	}
}

/*
In order to create a INSERT query we use the [session][Query] which expect a first [string] argument which is the query
followed by a number of [interface{}] which are the arguments to be inserted.
We can specify in the query the Consistency level. In this case One it's ok since this is a pet project.
Then we use the extended method of Query [Exec()] which execute the query.
In case there's a problem we receive an Error, otherwise it can be consider the command was correct and it was inserted
*/
func insertQuery(session *gocql.Session) {
	err := session.Query(`INSERT INTO demodb.user (id, name, age,sex) VALUES (?, ?, ?,?)`, gocql.TimeUUID(), "Politrons", "38", "Male").
		Consistency(gocql.One).
		Exec()
	if err != nil {
		log.Fatal(err)
	}
}

/*
In order to make Queries of SELECT it's similar than INSERT, it use the [session][Query] which expect a first [string] argument which is the query
followed by a number of [interface{}] which are the filter arguments in this case.
We can specify in the query the Consistency level. In this case One it's ok since this is a pet project.
Then we use the extended method of Query [Scan()] passing to the method the argument we want to bind with the query result, in the specific order.
In case there's a problem we receive an Error, otherwise it can be consider the command was correct and it was inserted

In order to make this whole transaction async, we use go routines and channels. To get further details go to the concurrency section of this project.
*/
func selectQuery(session *gocql.Session, user User, filterValue string, channel chan ResponseUser) {
	go func() {
		err := session.Query(`SELECT * FROM demodb.user WHERE name = ? LIMIT 1 ALLOW FILTERING`, filterValue).
			Consistency(gocql.One).
			Scan(&user.id, &user.name.value, &user.age.value, &user.sex.value)
		if err != nil {
			channel <- ResponseUser{User{}, err}
		}
		channel <- ResponseUser{User{user.id, Name{user.name.value}, Age{user.age.value}, Sex{user.sex.value}}, nil}
	}()
}

/*
When we want to SELECT all queries, we use same session and Query described before, but this time, instead of use
[Scan] to execute the query and bond the columns we use [Iter] which executes the query and returns an iterator
capable of iterating over all results.

Now for every iteration we use operator [Scan()] which bond the column of the query with the variables that we specify

In order to make this whole transaction async, we use go routines and channels. To get further details go to the concurrency section of this project.
*/
func selectAllQuery(session *gocql.Session, user User, selectAllChannel chan ResponseUsers) {
	go func() {
		users := []User{}
		iter := session.Query(`SELECT id, name FROM demodb.user`).Iter()
		for iter.Scan(&user.id, &user.name.value) {
			user := User{user.id, Name{user.name.value}, Age{}, Sex{}}
			users = append(users, user)
		}
		if err := iter.Close(); err != nil {
			selectAllChannel <- ResponseUsers{nil, err}
		}
		selectAllChannel <- ResponseUsers{users, nil}

	}()
}

/*
Using the [gocql] we follow the API to create a session with the next steps:

* Create the cluster using the [gocql] [NewCluster],
* Specify the keyspace that we create in Cassandra.
* Set the Consistency level setting in the cluster instance [Consistency]
* Create the session using [CreateSession] in case something goes wrong it return error, other wise the session.
*/
func createCassandraSession() *gocql.Session {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "demodb"
	cluster.Consistency = gocql.One
	session, err := cluster.CreateSession()
	if err != nil {
		fmt.Sprintln("Error creating Cassandra session since ", err)
	}
	return session
}
