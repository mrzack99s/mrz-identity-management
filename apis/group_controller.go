package apis

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/mrz-identity-management/ent"
	"github.com/mrzack99s/mrz-identity-management/services"
)

type GroupController struct {
	router gin.IRouter
}

func (ctl *GroupController) Create(c *gin.Context) {

	obj := ent.Groups{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	CheckLdapConnection()

	usr, err := services.CreateGroup(obj)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return

	} else {
		c.JSON(200, usr)
	}
}

func (ctl *GroupController) GetAll(c *gin.Context) {
	usr, err := services.ReadAllGroup()
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

func (ctl *GroupController) CountAll(c *gin.Context) {
	usr, err := services.CountAllGroup()
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return

	} else {
		c.JSON(200, usr)
	}
}

func (ctl *GroupController) Get(c *gin.Context) {

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

func (ctl *GroupController) Update(c *gin.Context) {

	obj := ent.Groups{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	CheckLdapConnection()

	usr, err := services.UpdateGroup(obj)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return

	} else {
		c.JSON(200, usr)
	}
}

func (ctl *GroupController) GetSearch(c *gin.Context) {

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

	usr, total_page, err := services.ReadGroupWithSearch(obj.Search, obj.Page, obj.PerPage)
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

func (ctl *GroupController) Delete(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	CheckLdapConnection()

	err := services.DeleteGroup(id)
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

func (ctl *GroupController) GetPagination(c *gin.Context) {

	page, _ := strconv.Atoi(c.Param("page"))
	perPage, _ := strconv.Atoi(c.Param("perPage"))

	usr, total_page, err := services.ReadGroupWithPagination(page, perPage)
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

// GroupController creates and register handles
func NewGroupController(router gin.IRouter) *GroupController {
	pc := &GroupController{
		router: router,
	}

	pc.register()

	return pc

}

func (ctl *GroupController) register() {
	router := ctl.router.Group("/groups")
	router.POST("create", ctl.Create)
	router.GET("get", ctl.GetAll)
	router.GET("get/:id", ctl.Get)
	router.POST("search", ctl.GetSearch)
	router.GET("get-pagination/:page/:perPage", ctl.GetPagination)
	router.GET("count", ctl.CountAll)
	router.POST("update", ctl.Update)
	router.DELETE("delete/:id", ctl.Delete)
}
