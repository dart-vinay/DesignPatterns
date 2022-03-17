package Creational

// This is used if we want to copy the object again and again. This prevents us from the overhead of fetching the data for the object again and again.
// Example places to use: Cloning products that belong to same Batch, clone network connection config objects, etc.
// Remember: Important to create a copy of the actual object than returning the copy of reference.

type Clone interface {
	Clone() CommonObject
}

type CommonObject interface {
}

type ClonnablePhone struct {
	Name  string
	Color string
	Brand string
}

type ClonnableTCPConnection struct {
	MaxTries       int
	MaxConnections int
	IdleTimeout    int
}

// Implements Clone Interface
func (obj *ClonnablePhone) Clone() CommonObject {
	newPhone := *obj
	return &newPhone
}

// Implements Clone Interface
func (obj *ClonnableTCPConnection) Clone() CommonObject {
	newConnection := *obj
	return &newConnection
}
