package apis

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/mrz-identity-management/ent"
	"github.com/mrzack99s/mrz-identity-management/services"
)

type OnlineSessionController struct {
	router gin.IRouter
}

func (ctl *OnlineSessionController) Create(c *gin.Context) {

	obj := ent.OnlineSession{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	usr, err := services.CreateOnlineSession(obj)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return

	} else {
		c.JSON(200, usr)
	}
}

func (ctl *OnlineSessionController) CountAll(c *gin.Context) {
	usr, err := services.CountAllOnlineSession()
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return

	} else {
		c.JSON(200, usr)
	}
}

func (ctl *OnlineSessionController) GetAll(c *gin.Context) {
	usr, err := services.ReadAllOnlineSession()
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return

	} else {
		c.JSON(200, gin.H{
			"data": usr,
		})
	}
}

func (ctl *OnlineSessionController) Delete(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	CheckLdapConnection()

	err := services.DeleteOnlineSession(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return

	} else {
		c.JSON(200, gin.H{
			"success": true,
		})

	}

}

// OnlineSessionController creates and register handles
func NewOnlineSessionController(router gin.IRouter) *OnlineSessionController {
	pc := &OnlineSessionController{
		router: router,
	}

	pc.register()

	return pc

}

func (ctl *OnlineSessionController) register() {
	router := ctl.router.Group("/online-session")
	router.POST("create", ctl.Create)
	router.GET("get", ctl.GetAll)
	router.GET("count", ctl.CountAll)
	router.DELETE("delete/:id", ctl.Delete)
}
