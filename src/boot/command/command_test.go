package command

import (
	"log"
	"testing"
)

func Test_Command(t *testing.T) {

	if o, e := ExecGit("sh", "-c", "git -C gittest status -s"); e != nil {
		t.Fail()
	} else {
		log.Fatal("==OUT== ", o)
	}

}
