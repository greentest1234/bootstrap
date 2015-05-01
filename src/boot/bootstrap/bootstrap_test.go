package bootstrap

import (
	//"boot/config"
	//"boot/log"
	"boot/models"
	"testing"
)

func testRusult(t *testing.T, f func() error) {
	if e := f(); e != nil {
		t.Fail()
	}
}

func Test_Download(t *testing.T) {
	//test Preverify
	//testRusult(t, testPreverify)

	//test DownloadScripts
	testRusult(t, testDownload)

	//test Download
	//testRusult(t, testDownload)

}
func testDownload() (err error) {

	url := "http://download.geonames.org/export/dump/GB.zip"
	return DownloadFile(url, ".shipped", "Vagrantfile")
}

func testPreverify() (err error) {

	//Initialize inputs
	user := &models.User{}
	projId := ""
	projName := ""
	serv := &models.Service{ServiceID: ""}

	b := NewBootstrap(user, projId, projName, []models.Service{*serv})
	return b.preverify()
}

func testDownloadScripts() (err error) {

	//Initialize inputs
	user := &models.User{}
	projId := ""
	projName := ""
	serv := &models.Service{ServiceID: ""}

	b := NewBootstrap(user, projId, projName, []models.Service{*serv})
	return b.downloadScripts()
}
