package Creational

const (
	MongoFactory = 1
	SQLFactory   = 2
)

type DataAccessFactory interface {
	// ideally would initiate a connection pool
	CreateWriteStore()
	CreateReadStore()

	// ideally would return a connection from the connection pool
	GetWriteStore() WriteStore
	GetReadStore() ReadStore

	// Further extend the factory to implement following methods
	// GetType() string
	// GetAlgorithm() string // this would the indexing algorithm used by out datastore. Can return the actual algorithm function as well which the user can override as per the requirements
	// GetMemoryLimit() string // in bytes
}

type WriteStore interface {
	GetWritePath() string
}

type ReadStore interface {
	GetReadPath() string
}

type SQLDataAccessFactory struct {
	Algorithm   string
	MemoryLimit string
	Type        string // NOSQL

	WriteStore *SQLWriteStore
	ReadStore  *SQLReadStore
}

type SQLReadStore struct {
	DirectoryPath string
}

type SQLWriteStore struct {
	DirectoryPath string
}

type MongoDataAccessFactory struct {
	Algorithm   string
	MemoryLimit string
	Type        string

	WriteStore *MongoWriteStore
	ReadStore  *MongoReadStore
}

type MongoReadStore struct {
	DirectoryPath string
}

type MongoWriteStore struct {
	DirectoryPath string
}

func (dataFactory *SQLDataAccessFactory) CreateWriteStore() {
	writeStore := new(SQLWriteStore)
	writeStore.DirectoryPath = "SQL write path"
	dataFactory.WriteStore = writeStore
}

func (dataFactory *SQLDataAccessFactory) CreateReadStore() {
	readStore := new(SQLReadStore)
	readStore.DirectoryPath = "SQL read path"
	dataFactory.ReadStore = readStore
}

func (dataFactory *SQLDataAccessFactory) GetWriteStore() WriteStore {

	return dataFactory.WriteStore
}

func (dataFactory *SQLDataAccessFactory) GetReadStore() ReadStore {
	return dataFactory.ReadStore
}

func (readStore *SQLReadStore) GetReadPath() string {
	return readStore.DirectoryPath
}

func (writeStore *SQLWriteStore) GetWritePath() string {
	return writeStore.DirectoryPath
}

func (dataFactory *MongoDataAccessFactory) CreateReadStore() {
	readStore := new(MongoReadStore)
	readStore.DirectoryPath = "Mongo Read Path"
	dataFactory.ReadStore = readStore
}

func (dataFactory *MongoDataAccessFactory) CreateWriteStore() {
	writeStore := new(MongoWriteStore)
	writeStore.DirectoryPath = "Mongo Write Path"
	dataFactory.WriteStore = writeStore
}

func (dataFactory *MongoDataAccessFactory) GetReadStore() ReadStore {
	return dataFactory.ReadStore
}

func (dataFactory *MongoDataAccessFactory) GetWriteStore() WriteStore {
	return dataFactory.WriteStore
}

func (readStore *MongoReadStore) GetReadPath() string {
	return readStore.DirectoryPath
}

func (writeStore *MongoWriteStore) GetWritePath() string {
	return writeStore.DirectoryPath
}

func ReturnDataFactory(val int) DataAccessFactory {
	switch val {
	case MongoFactory:
		return new(MongoDataAccessFactory)
	case SQLFactory:
		return new(SQLDataAccessFactory)
	default:
		return nil
	}
}
