package apiserver

import "github.com/gin-gonic/gin"

type APIServer struct {
	engine *gin.Engine
}

func (a *APIServer) registry() {
	v1 := a.engine.Group("/v1")
	{
		v1.GET("/api/admin/resource", GResource)
		v1.GET("/api/admin/name", GName)
	}

	v2 := a.engine.Group("/v2")
	{
		v2.GET("/api/admin/resources", GResources)
		v2.GET("/api/admin/names", GNames)
	}
}

func (a *APIServer) init() {}

func (a *APIServer) Start() {
	a.registry()
	a.engine.Run()
}

func New() *APIServer {
	return &APIServer{
		engine: gin.Default(),
	}
}
