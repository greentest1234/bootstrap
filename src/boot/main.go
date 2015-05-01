package main
import (
	"boot/bootstrap"
	"boot/config"
	"boot/models"
	"fmt"
)

func main() {

	//load config
	config.LoadConfig()

	//Initialize inputs
	user := &models.User{}
	user.ApiToken = "yZmnGXSURLIjHJWnepcrKhbuMYdPKgxO"
	projId := "cf95869a-eea8-11e4-83d1-0242ac11038b"
	projName := "Vijeay"

	serv := models.NewService()
	serv.Repository = "https://github.com/greentest1234/gittest.git"
	serv.Name = "gruntjs-hello-world"
	serv.ServiceID = "846a4bf4-eea9-11e4-83d5-0242ac11038b"

	//serv1 := models.NewService()
	//serv1.Repository = "https://github.com/CiscoCloud/shipped.git"
	//serv1.Name = "gruntjs-hello-world"
	//serv1.ServiceID = "ebb3d9ab-eea8-11e4-83d3-0242ac11038b"

	b := bootstrap.NewBootstrap(user, projId, projName, []models.Service{*serv})
	if er := b.Run(); er == nil {
		fmt.Println("Bootstap Completed !!")
	}
}
