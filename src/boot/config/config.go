package config

import (
	"boot/log"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
)

var Url struct {
	VbUrl    string
	HostPort string
}

const (
	CONFIGFILENAME = "config/config.json"
)

func LoadConfig() {
	curdir, _ := os.Getwd()

	configPath := path.Join(curdir, CONFIGFILENAME)
	configContent, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Error(err)
		panic("Unable to read the config file " + CONFIGFILENAME)
	}

	decoder := json.NewDecoder(bytes.NewBuffer(configContent))
	var conf map[string]interface{}
	if err := decoder.Decode(&conf); err != nil {
		panic("The content of " + CONFIGFILENAME + " are not valid JSON contents.")
	}
	if _url, k := conf["url"]; k {
		if url, k := _url.(map[string]interface{}); k {
			if vburl, k := url["virtualboxurl"]; k {
				Url.VbUrl = vburl.(string)
			}
			if hp, k := url["hostport"]; k {
				Url.HostPort = hp.(string)
			}
		}
	}

}
