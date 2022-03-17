package Creational

// Used for instantiating a single instance for an object.
// Typically used for initializing database connections which are used again and again.

var dbConn *Connection

type Connection struct {
	ConnectionString string
	CollectionName   string
}

func GetDBConnection(val string) *Connection {

	if dbConn == nil {
		dbConn = new(Connection)
		dbConn.ConnectionString = val
		dbConn.CollectionName = "MyApp"
	}
	return dbConn
}
