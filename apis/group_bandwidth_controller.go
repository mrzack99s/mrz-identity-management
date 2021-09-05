package apis

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/mrz-identity-management/ent"
	"github.com/mrzack99s/mrz-identity-management/services"
)

type BandwidthGroupController struct {
	router gin.IRouter
}

func (ctl *BandwidthGroupController) Create(c *gin.Context) {

	obj := ent.GroupBandwidth{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	CheckLdapConnection()

	usr, err := services.CreateGroupBandwidth(obj)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return

	} else {
		c.JSON(200, usr)
	}
}

func (ctl *BandwidthGroupController) GetAll(c *gin.Context) {
	usr, err := services.ReadAllAllGroupBandwidth()
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

func (ctl *BandwidthGroupController) Get(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	usr, err := services.ReadGroupBandwidth(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return

	} else {
		c.JSON(200, usr)

	}

}

func (ctl *BandwidthGroupController) Update(c *gin.Context) {

	obj := ent.GroupBandwidth{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	CheckLdapConnection()

	usr, err := services.UpdateGroupBandwidth(obj)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return

	} else {
		c.JSON(200, usr)
	}
}

func (ctl *BandwidthGroupController) GetPagination(c *gin.Context) {

	page, _ := strconv.Atoi(c.Param("page"))
	perPage, _ := strconv.Atoi(c.Param("perPage"))

	usr, total_page, err := services.ReadGroupBandwidthWithPagination(page, perPage)
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

func (ctl *BandwidthGroupController) Delete(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	CheckLdapConnection()

	err := services.DeleteGroupBandwidth(id)
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

// BandwidthGroupController creates and register handles
func NewBandwidthGroupController(router gin.IRouter) *BandwidthGroupController {
	pc := &BandwidthGroupController{
		router: router,
	}

	pc.register()

	return pc

}

func (ctl *BandwidthGroupController) register() {
	router := ctl.router.Group("/bw-groups")
	router.POST("create", ctl.Create)
	router.GET("get", ctl.GetAll)
	router.GET("get/:id", ctl.Get)
	router.GET("get-pagination/:page/:perPage", ctl.GetPagination)
	router.POST("update", ctl.Update)
	router.DELETE("delete/:id", ctl.Delete)
}
