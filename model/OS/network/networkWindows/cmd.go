package networkWindows

import (
	"bytes"
	"os/exec"
)

// `runcmd` returns the output of cmd command, and any errors.
func runcmd(tool string, command ...string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command(tool, command...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	return stdout.String(), stderr.String(), err
}

// `command` executes the cmd command.
// func command(tool string, command ...string) {
// 	out, errout, err := runcmd(tool, command...)

// 	if err != nil {
// 		log.Printf("error: %v\n", err)
// 		fmt.Print(errout)
// 	}

// 	fmt.Print(out)
// }
