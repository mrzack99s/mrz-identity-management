package config

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/go-ldap/ldap/v3"
	"github.com/mrzack99s/mrz-identity-management/ent"
)

func OpenDB() (*ent.Client, error) {

	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
			SystemConfigVar.DB.Username,
			SystemConfigVar.DB.Password,
			SystemConfigVar.DB.Hostname,
			SystemConfigVar.DB.DBName,
		),
	)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB("mysql", db)
	return ent.NewClient(ent.Driver(drv)), nil
}

func OpenLDAP() (*ldap.Conn, error) {

	ldapURL := fmt.Sprintf("ldaps://%s", SystemConfigVar.LDAP.Hostname)
	l, err := ldap.DialURL(ldapURL, ldap.DialWithTLSConfig(&tls.Config{InsecureSkipVerify: true}))
	if err != nil {
		return nil, err
	}
	err = l.Bind(fmt.Sprintf("CN=%s,%s",
		SystemConfigVar.LDAP.Username,
		SystemConfigVar.LDAP.BaseDN,
	), SystemConfigVar.LDAP.Password)
	if err != nil {
		return nil, err
	}

	return l, nil
}
