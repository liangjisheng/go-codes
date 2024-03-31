package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

func ssh1() {
	runSSH()
}

func SSHConnect(user, password, host, keyPath string, port int) (*ssh.Session, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))
	auth = append(auth, publicKeyAuthFunc(keyPath))

	hostKeyCallbk := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}

	clientConfig = &ssh.ClientConfig{
		User:            user,
		Auth:            auth,
		Timeout:         5 * time.Second,
		HostKeyCallback: hostKeyCallbk,
	}

	// connect to ssh
	addr = fmt.Sprintf("%s:%d", host, port)
	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}
	return session, nil
}

func runSSH() {
	var stdOut, stdErr bytes.Buffer

	session, err := SSHConnect("", "", "", "", 22)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	session.Stdout = &stdOut
	session.Stderr = &stdErr

	//session.Run("if [ -d liujx/project ]; then echo 0; else echo 1; fi")
	session.Run("ls ~")
	//ret, err := strconv.Atoi(strings.Replace(stdOut.String(), "\n", "", -1))
	//if err != nil {
	//	panic(err)
	//}

	fmt.Println(stdOut.String())
	//fmt.Printf("%s, %s\n", stdOut.String(), stdErr.String())
}

func publicKeyAuthFunc(keyPath string) ssh.AuthMethod {
	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		log.Fatal("ssh key file read failed", err)
	}
	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatal("ssh key signer failed", err)
	}
	return ssh.PublicKeys(signer)
}
