package Structural

// This pattern provides a substitute for other object. It controls the access to original object and lets us execute some requirement before or after the
// request goes through original object

// Necessity: If the original object is massive and it is only required from time to time, for example a database, then a caching proxy can be used before the request
// moves to database.

// Solution: Proxy has the same interface implementation as that of the original object. The client object calls this proxy object and then the proxy initiates the
// original object when required.

// Example: 1) Credit card can be seen as a proxy to cash. Both would implement the payment interface.
// 2) Cache can be seen as a proxy to database.

type SearchService interface {
	Search(string) Record
}

type Record struct {
	Id string
}

type RecordDatabase struct {
	Records []Record
}

type RecordList []Record

func (records *RecordList) AddRecordToStack(record Record) {
	*records = append(*records, record)
}

func (recordDB *RecordDatabase) Search(recordId string) Record {
	for _, record := range recordDB.Records {
		if record.Id == recordId {

			return record
		}
	}
	return Record{}
}

// This is the proxy to RecordDatabase
type RecordCache struct {
	Database      *RecordDatabase
	CachedRecords RecordList // assumption is that this caching records has a better search capability than the actual database record.
}

func (cacheDb *RecordCache) Search(recordId string) Record {
	found := false
	response := Record{}
	for _, record := range cacheDb.CachedRecords {
		if record.Id == recordId {
			response = record
			found = true
		}
	}
	if !found {
		response = cacheDb.Database.Search(recordId)
	}
	if response.Id != "" {
		cacheDb.CachedRecords.AddRecordToStack(response)
	}
	return response
}

func ProxyService() SearchService {
	database := []Record{
		{"1"},
		{"2"},
		{"3"},
	}
	recordDatabase := new(RecordDatabase)
	recordDatabase.Records = database

	proxyService := new(RecordCache)
	proxyService.Database = recordDatabase
	proxyService.CachedRecords = []Record{}

	return proxyService

}
