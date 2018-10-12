package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kprbc/kprbc/router"
)

func main() {
	server := gin.Default()
	server = router.Route(server)
	server.Run()
}
