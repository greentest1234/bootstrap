package bootstrap

import (
	"boot/log"
	"boot/models"
	"fmt"
	"os"
	"os/user"
	"path"
	//"path/filepath"
	"boot/command"
	"boot/config"
	"runtime"
	"strings"
)

type Bootstrap struct {
	ProjectID   string
	User        *models.User
	ProjectName string
	Services    []models.Service
}

func NewBootstrap(user *models.User, project_id string, projectName string, services []models.Service) *Bootstrap {
	return &Bootstrap{project_id, user, projectName, services}
}

const (
	OS_MAC = "drawin"
)

var curdir, _ = os.Getwd()
var VargantCwd = path.Join(curdir, ".shipped")
var DownloadDirPath string
var ProjectUrl string

func exec(f func() error) {
	log.Info("-----------------------------------")
	if e := f(); e != nil {
		log.Error(e)
		panic(e)
	}
}

func (this *Bootstrap) Run() error {

	if len(this.ProjectID) < 1 {
		return fmt.Errorf("Found Null ProjectId")
	}
	ProjectUrl = fmt.Sprintf("%s/cli/%s", config.Url.HostPort, this.ProjectID)

	//clone Repositories
	exec(this.cloneRepositoties)

	//Start Pre Verification
	//exec(this.preverify)

	//Download Scripts
	//exec(this.downloadScripts)

	//Install Virtualbox
	//exec(this.installVirtualBox)

	return nil
}

//# Function to clone repos
func (this *Bootstrap) cloneRepositoties() error {
	log.Info("Cloning Repositories Start")
	if len(this.Services) < 1 {
		return fmt.Errorf("Models.Services obj found null.")
	}

	//Check project Name directory exists
	//if [[ {{ escape .ProjectName }} != ${PWD##*/} ]]; then
	//  if [[ ! -d {{ escape .ProjectName }} ]]; then
	//    mkdir {{ escape .ProjectName }}
	//  fi
	//  cd {{ escape .ProjectName }}
	//fi

	if strings.Index(curdir, this.ProjectName) < 1 {
		if err := os.MkdirAll(this.ProjectName, 0777); err != nil {
			return err
		}
	}
	//TODO: check makefile is file not dir. if exists->delette this file
	//[[ -f .shipped/Makefile ]] && rm  .shipped/Makefile

	//{{range .Services}}
	//  clone_git_repo {{.SshUrl}} {{escape .Name}}
	//{{end}}
	for _, serv := range this.Services {
		var name string
		if name = serv.Name; len(name) < 1 {
			return fmt.Errorf("Models.Service.name is found empty.")
		}
		var sshUrl string
		if sshUrl = serv.SshUrl(); len(sshUrl) < 1 {
			return fmt.Errorf("Models.Service.SshUrl is found empty.")
		}
		if res, e := cloneGitRepo(sshUrl, name); e != nil {
			log.Error("ServiceID '" + serv.ServiceID + "' Got error while git clone. Err: " + e.Error())

		} else {
			log.Info("ServiceID '"+serv.ServiceID+"' Got error while git clone. Err: ", res)
		}

	}
	log.Info("Cloning Repositories End")
	return nil

}

func cloneGitRepo(src string, destn string) (string, error) {

	//git -C $2 status -s >/dev/null 2>&1 ; status=$?
	//    if [[ $status -ne 0 ]]; then
	//        git clone $1 $2
	//    fi

	out, er := command.ExecGit("git", "-C", "status -s >/dev/null 2>", src)
	if er != nil {
		return out, er
	} else {
		log.Info(out)
		if out, er := command.ExecGit("sh", "-c", "git clone", src, destn); er != nil {
			return out, er
		}
	}
	return "", nil
}

//Pre-verification function
// Also setup local directory for downloads
func (this *Bootstrap) preverify() error {
	log.Info("Pre-Verification Start")
	//check OS
	if runtime.GOOS != OS_MAC {
		return fmt.Errorf("Mac OS X not detected")
	}

	//get home dir
	usr, err := user.Current()
	if err != nil {
		return err
	}

	DownloadDirPath = path.Join(usr.HomeDir, "Downloads", "shipped")
	if _, err := os.Stat(DownloadDirPath); os.IsNotExist(err) {
		//path does not exist, lets create this
		if err := os.MkdirAll(DownloadDirPath, 0777); err != nil {
			return err
		}

	}
	log.Info("Pre-Verification End")
	return nil
}

//Download Script from git
func (this *Bootstrap) downloadScripts() error {

	log.Info("Download configs")

	log.Info("Project URL - ", ProjectUrl)

	//DownloadFile function declaration is DownloadFile(url string, pathToSave string, fileName string) error

	vargrantUrl := fmt.Sprintf("%s/Vagrantfile", ProjectUrl)
	//download shipped/vargantfile.
	if err := DownloadFile(vargrantUrl, ".shipped", "Vagrantfile"); err != nil {
		return err
	}

	makefileUrl := fmt.Sprintf("%s/makefile", ProjectUrl)
	//download Make file
	if err := DownloadFile(makefileUrl, ".", "Makefile"); err != nil {
		return err
	}

	//Git Clone download code will go here.
	//...
	//...I need to check if i can use use git commands using GO or I need to run shell for the same.
	//..

	log.Info("Verifed all config file for project")
	return nil
}

//# Function to install VirtualBox if not installed
func (this *Bootstrap) installVirtualBox() error {

	log.Info("Downloading virtual box.")

	//TODO: Check if VBox is already installed
	vbfilename := "virtualbox.dmg"
	//download shipped/vargantfile.
	if err := DownloadFile(config.Url.VbUrl, DownloadDirPath, vbfilename); err != nil {
		return err
	}

	log.Info("Mount VirtualBox disk image...")
	//hdiutil mount -quiet "$DOWNLOAD_DIR/virtualbox.dmg"
	if er := command.Exec("sh", "-c", "hdiutil mount -quite", fmt.Sprintf("%s/%s", DownloadDirPath, vbfilename)); er != nil {
		return er
	}

	log.Info("Install VirtualBox...")
	// sudo installer -pkg /Volumes/VirtualBox/VirtualBox.pkg -target /
	if er := command.Exec("sh", "-c", "sudo installer -pkg -quite", "Volumes/VirtualBox/VirtualBox.pkg", "-target", "/"); er != nil {
		return er
	}

	log.Info("Unmount VirtualBox disk image...")
	//umount /Volumes/VirtualBox/
	if er := command.Exec("sh", "-c", "sumount", "/Volumes/VirtualBox/"); er != nil {
		return er
	}

	log.Info("Installation done for virtual box ")
	return nil
}
