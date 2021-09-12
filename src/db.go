package main

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/go-sql-driver/mysql"
	// _ "github.com/go-sql-driver/mysql"
)

var (
	mysqlAutoCommit bool  = true
	errBadDbConf    error = errors.New("bad db conf")
)

func dbIsConnected(db *sql.DB) (bool, error) {
	err := db.Ping()
	if err != nil {
		db.Close()
		return false, err
	}
	return true, nil
}

func connectDB(sconf dbConf) (*sql.DB, error) {
	driverName := "mysql"
	// dsn = fmt.Sprintf("user:password@tcp(localhost:5555)/dbname?tls=skip-verify&autocommit=true")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?loc=Local", sconf.user, sconf.passwd, sconf.host, sconf.port, sconf.database)
	if mysqlAutoCommit {
		dsn += "&autocommit=true"
	}
	if sconf.serverCAPath != "" {
		dsn += "&tls=custom"
		rootCertPool := x509.NewCertPool()
		pem, err := ioutil.ReadFile(sconf.serverCAPath)
		if err != nil {
			return nil, err
		}
		if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
			return nil, errBadDbConf
		}
		if sconf.clientKeyPath != "" && sconf.clientCertPath != "" {
			// Both Key and Cert are set. Go with customer cert.
			clientCert := make([]tls.Certificate, 0, 1)
			certs, err := tls.LoadX509KeyPair(sconf.clientCertPath, sconf.clientKeyPath)
			if err != nil {
				return nil, err
			}
			clientCert = append(clientCert, certs)
			mysql.RegisterTLSConfig("custom", &tls.Config{
				// ServerName: "example.com",
				RootCAs:      rootCertPool,
				Certificates: clientCert,
			})
		} else if sconf.clientKeyPath == "" && sconf.clientCertPath == "" {
			// Neither Key or Cert is set. Proceed without customer cert.
			mysql.RegisterTLSConfig("custom", &tls.Config{
				// ServerName: "example.com",
				RootCAs: rootCertPool,
			})
		} else {
			// one of Key or Cert is set but not both, which is ILLEGAL.
			return nil, errBadDbConf
		}
	}

	db, err := sql.Open(driverName, dsn)
	if err != nil {
		return nil, err
	}

	if connected, err := dbIsConnected(db); !connected {
		return nil, err
	}

	return db, nil
}
