package main

import (
	"alif/internal/api"
	"alif/internal/jobs"
)

// @title API документация
// @version 1.0.0
// @description Документация к сервису
// @termsOfService http://swagger.io/terms/
// @contact.name Isfandiyor
// @contact.email isfand.zabirov@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost
// @accept json
// @produce json
// @schemes http
func main() {
	go jobs.StartJobs()
	api.RunApp()
}
