package values

import "os"

var SONAR_KEY = map[string]string{
	"cloud":         os.Getenv("SONAR_CLOUD_KEY"),
	"onprem":        os.Getenv("SONAR_ONPREM_KEY"),
	"onprem-legacy": os.Getenv("SONAR_ONPREM_LEGACY_KEY"),
}
