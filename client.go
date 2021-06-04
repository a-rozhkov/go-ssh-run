package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <host>", os.Args[0])
	}
    command_to_run := "sudo systemctl status"
    hostname := os.Args[1]+":22"
	client, session, err := connectToHost(hostname)
	if err != nil {
		panic(err)
	}
	out, err := session.CombinedOutput(command_to_run)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
	client.Close()
}

func connectToHost(host string) (*ssh.Client, *ssh.Session, error) {
	var pass string = "ENTER-PASSWORD-HERE"
    var user string = "ENTER-USERNAME-HERE"

    sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{ssh.Password(pass)},
	}
	sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	client, err := ssh.Dial("tcp", host, sshConfig)
	if err != nil {
		return nil, nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		client.Close()
		return nil, nil, err
	}

	return client, session, nil
}
