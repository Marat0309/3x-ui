package controller

import (
	"net/http"
	"strconv"

	"x-ui/database"
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
	g.GET(":id", a.get)
	g.PUT(":id", a.update)
	g.DELETE(":id", a.delete)
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

func (a *NodeController) get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		pureJsonMsg(c, http.StatusBadRequest, false, "invalid id")
		return
	}
	node, err := a.nodeService.Get(id)
	if database.IsNotFound(err) {
		pureJsonMsg(c, http.StatusNotFound, false, "node not found")
		return
	}
	if err != nil {
		pureJsonMsg(c, http.StatusInternalServerError, false, err.Error())
		return
	}
	jsonObj(c, node, nil)
}

func (a *NodeController) update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		pureJsonMsg(c, http.StatusBadRequest, false, "invalid id")
		return
	}
	var req struct {
		Name   string `json:"name"`
		ApiURL string `json:"apiUrl"`
		ApiKey string `json:"apiKey"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		pureJsonMsg(c, http.StatusBadRequest, false, err.Error())
		return
	}
	err = a.nodeService.Update(id, req.Name, req.ApiURL, req.ApiKey)
	if database.IsNotFound(err) {
		pureJsonMsg(c, http.StatusNotFound, false, "node not found")
		return
	}
	if err != nil {
		pureJsonMsg(c, http.StatusInternalServerError, false, err.Error())
		return
	}
	pureJsonMsg(c, http.StatusOK, true, "")
}

func (a *NodeController) delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		pureJsonMsg(c, http.StatusBadRequest, false, "invalid id")
		return
	}
	err = a.nodeService.Delete(id)
	if database.IsNotFound(err) {
		pureJsonMsg(c, http.StatusNotFound, false, "node not found")
		return
	}
	if err != nil {
		pureJsonMsg(c, http.StatusInternalServerError, false, err.Error())
		return
	}
	c.Status(http.StatusNoContent)
}
