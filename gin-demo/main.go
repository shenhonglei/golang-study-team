package main

import (
	_ "gin-demo/docs"
	"gin-demo/initRouter"
)

// @title Gin swagger
// @version 1.0
// @description Gin swagger Demo

// @contact.name Yang KUn
// @contact.url http://localhost:8080
// @contact.email yang@mail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
