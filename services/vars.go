package services

import (
	"github.com/go-ldap/ldap/v3"
	"github.com/mrzack99s/mrz-identity-management/ent"
)

var (
	DB_CLIENT         *ent.Client
	LDAP_CLIENT       *ldap.Conn
	BASE_DN           string
	DOMAIN_NAME       string
	EMAIL_DOMAIN_NAME string
	NEVER_EXPIRE      int = 9223372036854775807
)
