package main

import (
	_ "os"

	"github.com/mileapp_screening/cmd"
	_ "github.com/mileapp_screening/docs"
	_ "github.com/mileapp_screening/utils/env"
	_ "github.com/mileapp_screening/utils/log"
)

// @title           MileApp Screening Test
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host     localhost:5000
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	cmd.Execute()
}
