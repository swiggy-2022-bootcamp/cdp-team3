package main

import (
	"github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/app"
	_ "github.com/swiggy-2022-bootcamp/cdp-team3/transaction-service/docs"
)

// @title           Swagger Transaction Service API
// @version         1.0
// @description     Admin can add and view transaction amount for a customer based on custmer ID
// @termsOfService  http://swagger.io/terms/

// @contact.name   Jaithun Mahira
// @contact.email  swiggyb1035@datascience.manipal.edu

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3005

// @securityDefinitions.apiKey ApiKeyAuth
// @type apiKey
// @in header
// @name Authorization
func main() {
	app.Start()
}
