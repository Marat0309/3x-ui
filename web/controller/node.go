package controller

import (
	"net/http"

	"x-ui/web/service"

	"github.com/gin-gonic/gin"
)

// NodeController handles HTTP requests for managed nodes.
type NodeController struct {
	BaseController

	nodeService service.NodeService
}

func NewNodeController(g *gin.RouterGroup) *NodeController {
	a := &NodeController{}
	a.initRouter(g)
	return a
}

func (a *NodeController) initRouter(g *gin.RouterGroup) {
	g = g.Group("/nodes")
	g.Use(a.checkLogin)
	g.GET("", a.list)
	g.POST("", a.create)
}

func (a *NodeController) list(c *gin.Context) {
	nodes, err := a.nodeService.List()
	jsonObj(c, nodes, err)
}

func (a *NodeController) create(c *gin.Context) {
	var req struct {
		Name   string `json:"name"`
		ApiURL string `json:"apiUrl"`
		ApiKey string `json:"apiKey"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "msg": err.Error()})
		return
	}
	err := a.nodeService.Create(req.Name, req.ApiURL, req.ApiKey)
	jsonMsg(c, "", err)
}
