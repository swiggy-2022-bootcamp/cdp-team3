package main

import (
	"github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/app"
	_ "github.com/swiggy-2022-bootcamp/cdp-team3/rewards-service/docs"
)

// @title           BuyItNow Rewards Service
// @version         1.0
// @description
// @termsOfService  http://swagger.io/terms/

// @contact.name   Uttej Immadi
// @contact.email  swiggyb3014@datascience.manipal.edu

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3008

// @securityDefinitions.apiKey ApiKeyAuth
// @type apiKey
// @in header
// @name Authorization
func main() {
	app.Start()
}
