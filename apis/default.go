package apis

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/mrz-identity-management/config"
	"github.com/mrzack99s/mrz-identity-management/services"
)

func DefaultSystem(router gin.IRouter) {
	NewGroupController(router)
	NewBandwidthGroupController(router)
	NewUserController(router)
	NewOnlineSessionController(router)
}

func CheckLdapConnection() {
	if services.LDAP_CLIENT.IsClosing() {
		l, err := config.OpenLDAP()
		if err != nil {
			log.Fatal(err)
		}
		services.LDAP_CLIENT = l
	}
}
