package apiserver

import (
	"io"

	"github.com/emicklei/go-restful"
)

func GetUser(request *restful.Request, response *restful.Response) {
	response.WriteEntity("Hello world: Users API")
}

func UpdateUser(request *restful.Request, response *restful.Response) {

	var name string

	name = request.PathParameter("user")
	response.WriteEntity("hello " + name)

}

func PostUser(request *restful.Request, response *restful.Response) {

	var name string

	name = request.PathParameter("user")
	io.WriteString(response.ResponseWriter, "this would be a normal response")
	response.WriteEntity("hello " + name)
}

func DeleteUser(request *restful.Request, response *restful.Response) {

	var name string

	name = request.PathParameter("name")
	response.WriteEntity("delete " + name)
}

func GetName(request *restful.Request, response *restful.Response) {

	response.WriteEntity("Hello world: Names API")
}

func PutName(request *restful.Request, response *restful.Response) {}

func DeleteName(request *restful.Request, response *restful.Response) {}

func PostName(request *restful.Request, response *restful.Response) {}
