package main

import (
	"github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/app"
	_ "github.com/swiggy-2022-bootcamp/cdp-team3/admin-service/docs"
)

// @title           BuyItNow Admin Service
// @version         1.0
// @description
// @termsOfService  http://swagger.io/terms/

// @contact.name   Uttej Immadi
// @contact.email  swiggyb3014@datascience.manipal.edu

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3009

// @securityDefinitions.apiKey ApiKeyAuth
// @type apiKey
// @in header
// @name Authorization
func main() {
	app.Start()
}
