package main

import (
	"os"

	"github.com/fsouza/go-dockerclient"
)

// NewClient returns a Client instance ready for communication with the given
// server endpoint. It will use the latest remote API version available in the
// server.

func main() {

	endpoint := "http://9.111.213.128:2375"
	client, _ := docker.NewClient(endpoint)

	exec, _ := client.CreateExec(docker.CreateExecOptions{
		Container:    "68512efde46c",
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Tty:          true,
		Cmd:          []string{"/bin/bash"},
	})

	success := make(chan struct{})
	execID := exec.ID

	opts := docker.StartExecOptions{
		OutputStream: os.Stdout,
		ErrorStream:  os.Stderr,
		InputStream:  os.Stdin,
		RawTerminal:  true,
		Success:      success,
	}
	go func() {
		if err := client.StartExec(execID, opts); err != nil {

		}
	}()
	<-success
}
