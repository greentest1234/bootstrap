package command

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing[ CMD ]: %s\n", strings.Join(cmd.Args, " "))
}

func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("\t--> Error: %s\n", err.Error()))
	}
}

func printOutput(o []byte) {
	if len(o) > 0 {
		fmt.Printf("\t--> Out: %s\n", string(o))
	}
}

func Exec(c string, args ...string) error {
	cmd := exec.Command(c, args...)
	printCommand(cmd)
	output, err := cmd.CombinedOutput()
	printError(err)
	printOutput(output)
	return err
}

func ExecGit(c string, args ...string) (string, error) {
	cmd := exec.Command(c, args...)
	// Combine stdout and stderr
	printCommand(cmd)
	output, err := cmd.CombinedOutput()
	printError(err)
	printOutput(output)
	var out string
	if len(output) > 0 {
		out = string(output)
	}
	return out, err
}

func ExecWaitSts(c string, args ...string) (err error) {

	cmd := exec.Command(c, args...)
	printCommand(cmd)
	var waitStatus syscall.WaitStatus
	err = cmd.Run()
	if err != nil {
		printError(err)
		// Did the command fail because of an unsuccessful exit code
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus = exitError.Sys().(syscall.WaitStatus)
			printOutput([]byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
		}
		return
	} else {
		// Command was successful
		waitStatus = cmd.ProcessState.Sys().(syscall.WaitStatus)
		printOutput([]byte(fmt.Sprintf("%d", waitStatus.ExitStatus())))
	}

	return
}
