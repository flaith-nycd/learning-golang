package main

import (
	"fmt"
	"log"
	"os/exec"
	"reflect"
)

func main() {
    // Use os/exec command
    // Mandatory
    // name: 'cmd'
    // args: '/C' <- Need to use this argument for 'cmd'
    //       'dir', ...
	cmd := exec.Command("cmd", "/C", "dir", "/c", "/q")

    // CombinedOutput runs the command and returns its combined standard
    // output and standard error.
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

    // Print the result
	fmt.Print(string(stdout))
}
