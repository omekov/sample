package main

import "github.com/omekov/sample/internal/apiserver"

// @title Sample API
// @version 2.0
// @description This is a sample service for managment
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:9090
// @BasePath /
// @securitydefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	apiserver.Run()
}
