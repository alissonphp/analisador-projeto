package values

import "os"

var SONAR_HOST = map[string]string{
	"cloud":         os.Getenv("SONAR_CLOUD_HOST"),
	"onprem":        os.Getenv("SONAR_ONPREM_HOST"),
	"onprem-legacy": os.Getenv("SONAR_ONPREM_LEGACY_HOST"),
}
