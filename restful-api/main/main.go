package main

import apiserver "gopher/restful-api/ui/api-server"

func main() {

	apiserver.NewAPIServer().Start()
}
