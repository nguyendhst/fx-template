package database

type Database interface {
	Connect() error
	Close() error
	GetDBConnection() any
	String() string
}
