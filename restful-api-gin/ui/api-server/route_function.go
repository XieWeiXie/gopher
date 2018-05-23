package apiserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GResource(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "resource",
	})

}

func GName(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "name",
	})
}

func GResources(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "resources",
	})
}

func GNames(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "names",
	})
}
