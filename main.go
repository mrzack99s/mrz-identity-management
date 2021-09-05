package main

import (
	"context"
	"fmt"
	"log"
	"time"

	env "github.com/Netflix/go-env"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	cors "github.com/itsjamie/gin-cors"
	"github.com/joho/godotenv"
	"github.com/mrzack99s/mrz-identity-management/apis"
	"github.com/mrzack99s/mrz-identity-management/config"
	"github.com/mrzack99s/mrz-identity-management/services"
)

func main() {

	godotenv.Load()

	_, err := env.UnmarshalFromEnviron(&config.SystemConfigVar)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	if config.SystemConfigVar.Prod {
		gin.SetMode(gin.ReleaseMode)
	}

	//Unsecure API
	unsecureAPI := router.Group("/unsecure/api")
	apis.NewUnsecureController(unsecureAPI)

	//Secure API
	secureAPI := router.Group("/api", services.APIAuthentication)
	apis.DefaultSystem(secureAPI)

	//A2A API
	initAccountA2A := gin.Accounts{}
	initAccountA2A["a2a"] = config.SystemConfigVar.Security.A2ASecret
	a2aAPI := router.Group("/a2a", gin.BasicAuth(initAccountA2A))
	apis.NewA2AController(a2aAPI)

	db_client, err := config.OpenDB()
	if err != nil {
		log.Fatalf("fail to open mysql: %v", err)
	}
	defer db_client.Close()

	if err := db_client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	ldap_client, err := config.OpenLDAP()
	if err != nil {
		log.Fatalf("fail to open ldap: %v", err)
	}
	defer ldap_client.Close()

	services.DB_CLIENT = db_client
	services.LDAP_CLIENT = ldap_client
	services.BASE_DN = config.SystemConfigVar.LDAP.BaseDN
	services.DOMAIN_NAME = config.SystemConfigVar.LDAP.DomainName
	services.EMAIL_DOMAIN_NAME = config.SystemConfigVar.LDAP.EmailDomainName

	router.Run(fmt.Sprintf(":%d", config.SystemConfigVar.Port))

}
