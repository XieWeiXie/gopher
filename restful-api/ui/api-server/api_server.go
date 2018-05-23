package apiserver

import (
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
)

type APIServer struct {
	container *restful.Container
}

type APIServers []APIServer

func (a *APIServer) init() {
	wsContainer := restful.NewContainer()
	wsContainer.EnableContentEncoding(true)
	wsContainer.Router(restful.CurlyRouter{})
	a.container = wsContainer
}

func (a *APIServer) register() {
	ws := new(restful.WebService)
	ws.Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON)

	ws.Path("/users")
	ws.Route(ws.GET("/api/v1/tenants/hello").To(GetUser))
	ws.Route(ws.PUT("/api/v1/tenants/{user}").To(UpdateUser))
	ws.Route(ws.POST("/api/v1/tenants/{user}").To(PostUser))
	ws.Route(ws.DELETE("/api/v1/tenants/{user}").To(DeleteUser))

	ws.Route(ws.GET("/api/v1/tenants/all").To(GetName))
	ws.Route(ws.PUT("/api/v1/tenants/{name}").To(PutName))
	ws.Route(ws.DELETE("/api/v1/tenants/{name}").To(DeleteName))
	ws.Route(ws.POST("/api/v1/tenants/{name}").To(PostName))

	a.container.Add(ws)

}

func (a *APIServer) Start() {
	server := &http.Server{Addr: ":8080", Handler: a.container}
	log.Fatal(server.ListenAndServe())
}

func New() *APIServer {
	apiServer := &APIServer{container: nil}
	apiServer.init()
	apiServer.register()
	return apiServer
}
