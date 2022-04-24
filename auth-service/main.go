package main

import (
	"github.com/swiggy-2022-bootcamp/cdp-team3/auth-service/app"
)

// @title           Swagger Auth Service API
// @version         1.0
// @description     Users ( Admin, Customer, etc ) can login and get a token and use it to access other APIs
// @termsOfService  http://swagger.io/terms/

// @contact.name   Rishabh Mishra
// @contact.email  swiggyb2026@datascience.manipal.edu

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3012

// @securityDefinitions.apiKey ApiKeyAuth
// @type apiKey
// @in cookie
// @name Authorization
func main() {
	app.Start()
}
