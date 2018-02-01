package domain

import (
	"fmt"
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type DbConfig struct {
	// Optional.
	Username, Password string

	// Host of the db instance.
	//
	// If set, UnixSocket should be unset.
	Host string

	// Port of the db instance.
	//
	// If set, UnixSocket should be unset.
	Port int

	// UnixSocket is the filepath to a unix socket.
	//
	// If set, Host and Port should be unset.
	UnixSocket string
}

func NewMysqlConn(options *DbConfig) (*gorm.DB, error) {

	conn, err := gorm.Open("mysql", options.getConnectionString("todo-db"))
	if err != nil {
		return nil, fmt.Errorf("mysql: could not get a connection: %v", err)
	}
	
	return conn, nil
}

func (config DbConfig) getConnectionString(databaseName string) string {

	var connString string
	// [username[:password]@]
	if config.Username != "" {
		connString = config.Username
		if config.Password != "" {
			connString = connString + ":" + config.Password
		}
		connString = connString + "@"
	}

	if config.UnixSocket != "" {
		return fmt.Sprintf("%sunix(%s)/%s?parseTime=true", connString, config.UnixSocket, databaseName)
	}
	return fmt.Sprintf("%stcp([%s]:%d)/%s?parseTime=true", connString, config.Host, config.Port, databaseName)
}