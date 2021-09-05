package apis

import (
	"encoding/json"
	"io/ioutil"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/mrz-identity-management/ent"
	"github.com/mrzack99s/mrz-identity-management/services"
)

type UserController struct {
	router gin.IRouter
}

func (ctl *UserController) Create(c *gin.Context) {

	obj := ent.Users{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	CheckLdapConnection()

	usr, err := services.CreateUser(obj)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return

	} else {
		c.JSON(200, usr)
	}
}

func (ctl *UserController) ResetPassword(c *gin.Context) {

	obj := ent.Users{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	CheckLdapConnection()

	err := services.ResetPassword(obj.UPid)
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

func (ctl *UserController) MultiCreate(c *gin.Context) {

	obj := []services.MultiCreateSchema{}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	err = json.Unmarshal(body, &obj)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	CheckLdapConnection()

	msgErr := services.MultiCreateUser(obj)
	if len(msgErr) > 0 {
		c.JSON(400, gin.H{
			"error": msgErr,
		})
		return

	} else {
		c.JSON(200, gin.H{
			"success": true,
		})
	}
}

func (ctl *UserController) CountAll(c *gin.Context) {
	usr, err := services.CountAllUser()
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return

	} else {
		c.JSON(200, usr)
	}
}

func (ctl *UserController) GetAll(c *gin.Context) {
	usr, err := services.ReadAllUser()
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

func (ctl *UserController) Get(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	usr, err := services.ReadUser(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return

	} else {
		c.JSON(200, usr)

	}

}

func (ctl *UserController) GetPagination(c *gin.Context) {

	page, _ := strconv.Atoi(c.Param("page"))
	perPage, _ := strconv.Atoi(c.Param("perPage"))

	usr, total_page, err := services.ReadUserWithPagination(page, perPage)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return

	} else {
		c.JSON(200, gin.H{
			"record_count": total_page,
			"record_list":  usr,
		})

	}

}

func (ctl *UserController) GetSearch(c *gin.Context) {

	type SearchSchema struct {
		Search  string `json:"search"`
		Page    int    `json:"page"`
		PerPage int    `json:"p_page"`
	}

	obj := SearchSchema{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	usr, total_page, err := services.ReadUserWithSearch(obj.Search, obj.Page, obj.PerPage)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return

	} else {
		c.JSON(200, gin.H{
			"record_count": total_page,
			"record_list":  usr,
		})

	}

}

func (ctl *UserController) Update(c *gin.Context) {

	obj := ent.Users{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	CheckLdapConnection()

	usr, err := services.UpdateUser(obj)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return

	} else {
		c.JSON(200, usr)
	}
}

func (ctl *UserController) Delete(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	CheckLdapConnection()

	err := services.DeleteUser(id)
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

// UserController creates and register handles
func NewUserController(router gin.IRouter) *UserController {
	pc := &UserController{
		router: router,
	}

	pc.register()

	return pc

}

func (ctl *UserController) register() {
	router := ctl.router.Group("/users")
	router.POST("create", ctl.Create)
	router.POST("multi-create", ctl.MultiCreate)
	router.GET("get", ctl.GetAll)
	router.GET("count", ctl.CountAll)
	router.POST("search", ctl.GetSearch)
	router.GET("get/:id", ctl.Get)
	router.GET("get-pagination/:page/:perPage", ctl.GetPagination)
	router.POST("update", ctl.Update)
	router.POST("reset-password", ctl.ResetPassword)
	router.DELETE("delete/:id", ctl.Delete)

}
