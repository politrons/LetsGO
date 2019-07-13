package connectors

import (
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"testing"
)

func TestCassandra(t *testing.T) {
	session := createCassandraSession()
	runQuery(session)

	var id gocql.UUID
	var name string
	var age string
	var sex string

	/* Search for a specific set of records whose 'timeline' column matches
	 * the value 'me'. The secondary index that we created earlier will be
	 * used for optimizing the search */
	err := session.Query(`SELECT id, name, age, sex FROM demodb.user WHERE name = ? LIMIT 1 ALLOW FILTERING`, "Politrons").
		Consistency(gocql.One).
		Scan(&id, &name, &age, &sex)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User:", id, name, age, sex)

}

/*
In order t
*/
func runQuery(session *gocql.Session) {
	err := session.Query(`INSERT INTO demodb.user (id, name, age,sex) VALUES (?, ?, ?,?)`, gocql.TimeUUID(), "Politrons", "38", "Male").Exec()
	if err != nil {
		log.Fatal(err)
	}
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
