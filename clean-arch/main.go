package main

import (
	"backend/clean-arch/configs"
	"backend/clean-arch/deliveries/http/routes"

	_ "backend/clean-arch/docs"

	"github.com/labstack/echo/v4"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8081
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	db := configs.InitPostgres()
	cleanArch := routes.NewCleanArch(&db)

	e := echo.New()
	routes.SetupRouter(e, cleanArch)
	e.Logger.Fatal(e.Start(":8081"))
}
