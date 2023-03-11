package services

import (
	"genfile2/appconfig"
	"genfile2/helps"
)

func init() {
	appconfig.LoadInitConfig()
}
func Generate() {
	filepath := appconfig.Config.FilePath
	var end helps.Endpoints
	end.GenFile(filepath)

}
