package cfg

import "github.com/aliforever/golang-backend-training/chapter12/section12.4/arg"

var PayApiUrl = ""

func init() {

	switch arg.DeploymentEnvironment {
	case "DEV":
		PayApiUrl = "http://localhost:8080/v1/pay"
	case "QA":
		PayApiUrl = "https://api.qa.example.com/v1/pay"
	case "STAGING":
		PayApiUrl = "https://api.staging.example.com/v1/pay"
	case "PROD":
		PayApiUrl = "https://api.example.com/v1/pay"
	default:
		panic("Unknown Deployment Environment: " + arg.DeploymentEnvironment)
	}
}
