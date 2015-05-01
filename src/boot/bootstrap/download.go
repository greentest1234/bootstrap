package bootstrap

import (
	"boot/log"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func DownloadFile(url string, pathToSave string, fileName string) error {

	//Validations
	if len(url) < 1 {
		return fmt.Errorf("Download.Url cannot be empty")
	}
	if len(fileName) < 1 {
		return fmt.Errorf("Download.FileName cannot be empty")
	}

	fileSavePath := path.Join(pathToSave, fileName)
	log.Info("Starting Download for ", url, " -To- ", fileSavePath)

	if _, err := os.Stat(pathToSave); os.IsNotExist(err) {
		if err := os.MkdirAll(pathToSave, 0777); err != nil {
			return err
		}
	}

	if _, err := os.Stat(fileSavePath); err == nil {
		//no need to download file file already exists
		log.Info("File Already Exists. File: ", fileName)
		return nil
	}

	response, err := http.Get(url)
	if err != nil {
		log.Error("Error while downloading", url, "-", err)
		return err
	}

	output, err := os.Create(fileSavePath)
	if err != nil {
		log.Error("Error while creating", fileName, "-", err)
		return err
	}
	defer output.Close()

	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		log.Error("Error while downloading", url, "-", err)
		return err
	}

	log.Info(n, " bytes downloaded.")
	return nil
}
