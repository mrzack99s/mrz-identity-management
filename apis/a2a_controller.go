package apis

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/mrz-identity-management/ent"
	"github.com/mrzack99s/mrz-identity-management/services"
)

type A2AController struct {
	router gin.IRouter
}

func (ctl *A2AController) Create(c *gin.Context) {

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

func (ctl *A2AController) Authentication(c *gin.Context) {

	obj := services.AuthSchema{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	res := services.GetAuthentication(obj.Username, obj.Password, false)
	if res.Status != "success" {
		c.JSON(400, gin.H{
			"error": res.Message,
		})
		return

	} else {
		c.JSON(200, res)
	}
}

func (ctl *A2AController) Get(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	usr, err := services.ReadGroup(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return

	} else {
		c.JSON(200, usr)

	}

}

// GroupController creates and register handles
func NewA2AController(router gin.IRouter) *A2AController {
	pc := &A2AController{
		router: router,
	}

	pc.register()

	return pc

}

func (ctl *A2AController) register() {
	session := ctl.router.Group("/session")
	session.POST("create", ctl.Create)
	session.GET("get/:id", ctl.Get)
	auth := ctl.router.Group("/authentication")
	auth.POST("", ctl.Authentication)
}
