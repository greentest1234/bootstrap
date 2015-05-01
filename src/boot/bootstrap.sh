#!/bin/bash

if [[ -z "$HOST_PORT" ]]; then
  HOST_PORT="shipped-api.shipped-cisco.com:80"
fi
project_url=http://${HOST_PORT}/cli/{{ .ProjectID }}

#########################################################################
# Function download all project related configs needed for local sandbox
# and hosted CICD build
#########################################################################
function download_file {
    if [[ ! -f "$1" ]] ; then
        curl -s -o $1 $2 ; status=$?
        if [[ $status -ne 0 ]]; then
            echo "Failed: download $2" error; exit 1
        fi

        grep -qvc "404 page not found" $1 ;status2=$?
        if [[ $status2 -ne 0 ]]; then
            echo "Failed: download $2" error; rm $1; exit 1
        fi
    fi
}

#########################################################################
# function to clone
#########################################################################
function  clone_git_repo {
    git -C $2 status -s >/dev/null 2>&1 ; status=$?
    if [[ $status -ne 0 ]]; then
        git clone $1 $2
    fi
}

#########################################################################
# Checks pre-requisite is satsfied to run script
#########################################################################
function pre_verify {
    export VAGRANT_CWD="$PWD/.shipped"

    if [[ "$(uname)" != "Darwin" ]]; then
        cecho "Error: Mac OS X not detected" error
        exit 1
    fi

    DOWNLOAD_DIR="$HOME/Downloads/shipped"
    if [[ ! -d $DOWNLOAD_DIR ]]; then
        cecho "Create shipped download directory..." h3
        mkdir -p "$DOWNLOAD_DIR"
    fi
}

function download_scripts {
    cecho "Download configs" h1

    mkdir -p .shipped
    download_file .shipped/Vagrantfile ${project_url}/Vagrantfile
		download_file Makefile ${project_url}/makefile
    {{range .Services}}
		git -C {{escape .Name}} reset .
		needsCommit=false
		if [[ ! -f {{escape .Name}}/.drone.yml ]]
		then
			download_file {{escape .Name}}/.drone.yml ${project_url}/{{.ServiceID}}/drone
			git -C {{escape .Name}} add .drone.yml
			needsCommit=true
		fi
		if [[ ! -f {{escape .Name}}/Dockerfile ]]
		then
			download_file {{escape .Name}}/Dockerfile ${project_url}/{{.ServiceID}}/Dockerfile
			git -C {{escape .Name}} add Dockerfile
			needsCommit=true
		fi
		if $needsCommit
		then
			git -C {{escape .Name}} commit -m "**Shipped -- Add Drone Configuration"
		fi
    {{end}}

    cecho "Verifed all config file for project"
}

#########################################################################
# Function to install VirtualBox if not installed
#########################################################################
function install_virtualbox {
    cecho "Install VirtualBox" h1
    VIRTUALBOX_URL="http://download.virtualbox.org/virtualbox/4.3.22/VirtualBox-4.3.22-98236-OSX.dmg"

    if [[ "$(command -v VirtualBox >/dev/null 2>&1)" -ne 0 ]]; then
        cecho "Download VirtualBox..."
        curl -L -o "$DOWNLOAD_DIR/virtualbox.dmg" "$VIRTUALBOX_URL"

        cecho "Mount VirtualBox disk image..."
        hdiutil mount -quiet "$DOWNLOAD_DIR/virtualbox.dmg"

        cecho "Install VirtualBox..."
        sudo installer -pkg /Volumes/VirtualBox/VirtualBox.pkg -target /

        cecho "Unmount VirtualBox disk image..."
        umount /Volumes/VirtualBox/
    fi
    cecho "Verified VirtualBox Installed"
}

#########################################################################
# Function to install vagrant if not installed
#########################################################################
function install_vagrant {
    cecho "Install Vagrant" h1
    VAGRANT_URL="https://dl.bintray.com/mitchellh/vagrant/vagrant_1.7.2.dmg"

    if [[ "$(command -v vagrant >/dev/null 2>&1)" -ne 0 ]]; then
        cecho "Download Vagrant..."
        curl -L -o "$DOWNLOAD_DIR/vagrant.dmg" "$VAGRANT_URL"

        cecho "Mount Vagrant disk image..."
        hdiutil mount -quiet "$DOWNLOAD_DIR/vagrant.dmg"

        cecho "Install Vagrant..."
        sudo installer -pkg /Volumes/Vagrant/Vagrant.pkg -target /

        cecho "Unmount Vagrant disk image..."
        umount /Volumes/Vagrant/
    fi
    cecho "Verified Vagrant setup"
}

##########################################################################
# Start vagrant vm if not running, first time it also download vagrantbox
##########################################################################
function sandbox_vm_up {
    cecho "Bootstrap sandbox VM" h1
    vagrant status | grep -qc running >/dev/null 2>&1; status=$?
    if [[ $status -ne 0 ]]; then
        cecho "Bringing up virtual machine" warn
        export current_wd=$PWD
        cd $VAGRANT_CWD
        time vagrant up --no-parallel --provider=docker
        cd $current_wd
    fi
    cecho "Verifed shipped developer sandbox VM is running"
}

# Util function set state for step on shipped
function set_shipped_state {
    STATUS_URL=${HOST_PORT}/cli/{{ .ProjectID }}/status
    curl -X POST ${STATUS_URL}/$1?api_token={{ .User.ApiToken }} >/dev/null 2>&1
}

# Util function for user friendly cecho
function cecho {
    local exp=$1;
    local header=$2;
    black=0; red=1; green=2; yellow=3; blue=4; magenta=5; cyan=6; white=7
    H1START="============================================================"; H1END=$H1START
    H2START="------------------------------------------------------------"; H2END=$H2START

    case $(echo $header | tr '[:upper:]' '[:lower:]') in
        h1) tput bold; tput setaf $green ;echo;echo "$H1START"; echo "                "; echo $exp; echo "$H1END" ;;

        h2) tput setaf $blue ;echo "$H2START"; echo $exp; echo "$H2END";;

        h3) tput setaf $magenta ; echo $exp ;;

        warn) tput setaf $yellow; echo $exp;;

        error) tput setaf $red; tput bold; echo $exp;;

        text|*) echo $exp ;;
    esac
    tput sgr0;
}

cecho "Cloning Repositories" h1
if [[ {{ escape .ProjectName }} != ${PWD##*/} ]]; then
  if [[ ! -d {{ escape .ProjectName }} ]]; then
    mkdir {{ escape .ProjectName }}
  fi
  cd {{ escape .ProjectName }}
fi

[[ -f .shipped/Makefile ]] && rm  .shipped/Makefile

{{range .Services}}
  clone_git_repo {{.SshUrl}} {{escape .Name}}
{{end}}

# pre_verify and download_script must be run first
pre_verify
set_shipped_state project_init

download_scripts
set_shipped_state download_cli

install_virtualbox
set_shipped_state install_vbox

install_vagrant
set_shipped_state install_vagrant

sandbox_vm_up
set_shipped_state sandbox_vm_up
set_shipped_state down_buildpack

cecho "Application running on http://localhost:${CONTAINER_PORT}" h2
