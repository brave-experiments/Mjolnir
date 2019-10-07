package connect

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	"os"
)

type SshClient struct {
	Config   *ssh.ClientConfig
	hostname string
	network  string
}

type DialError struct {
	Message string
}

type privateKey struct {
	location string
	signer   ssh.Signer
}

func (err DialError) Error() (message string) {
	return err.Message
}

func (sshClient *SshClient) New(user string, hostname string, keyPath string) (err error) {
	authorizationKey := privateKey{location: keyPath}
	err, method := authorizationKey.method()

	if nil != err {
		return err
	}

	sshClient.Config = &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			method,
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	sshClient.hostname = hostname
	sshClient.network = "tcp"

	return nil
}

func (sshClient *SshClient) Dial() (client *ssh.Client, err error) {
	if nil == sshClient.Config {
		return nil, DialError{Message: "Config not present"}
	}

	client, err = sshClient.attachPipes()

	return client, err
}

func (key *privateKey) method() (err error, method ssh.AuthMethod) {
	priv, err := ioutil.ReadFile(key.location)

	if nil != err {
		return err, nil
	}

	signer, err := ssh.ParsePrivateKey(priv)

	if nil != err {
		return err, nil
	}

	key.signer = signer

	return err, ssh.PublicKeys(signer)
}

func (sshClient *SshClient) attachPipes() (client *ssh.Client, err error) {
	addr := sshClient.hostname + ":22"
	fmt.Println("Address: ", addr)
	client, err = ssh.Dial(sshClient.network, addr, sshClient.Config)

	if nil != err {
		return client, err
	}

	session, err := client.NewSession()

	if err != nil {
		return client, err
	}

	err = attachStdIn(*session)

	if err != nil {
		return client, err
	}

	err = attachStdOut(*session)

	if err != nil {
		return client, err
	}

	err = attachStdErr(*session)

	return client, err
}

func attachStdIn(session ssh.Session) (err error) {
	stdin, err := session.StdinPipe()

	if err != nil {
		return err
	}

	go func() {
		_, err = io.Copy(stdin, os.Stdin)

		if nil != err {
			panic(err)
		}
	}()

	return nil
}

func attachStdOut(session ssh.Session) (err error) {
	stdOut, err := session.StdoutPipe()

	if err != nil {
		return err
	}

	go func() {
		_, err = io.Copy(os.Stdin, stdOut)

		if nil != err {
			panic(err)
		}
	}()

	return nil
}

func attachStdErr(session ssh.Session) (err error) {
	stdErr, err := session.StderrPipe()

	if err != nil {
		return err
	}

	go func() {
		_, err = io.Copy(os.Stderr, stdErr)

		if nil != err {
			panic(err)
		}
	}()

	return nil
}
