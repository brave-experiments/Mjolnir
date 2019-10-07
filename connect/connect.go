package connect

import (
	"fmt"
	"golang.org/x/crypto/ssh"
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

	defer func() {
		err = session.Close()

		if nil != err {
			fmt.Println(err)
		}
	}()

	if err != nil {
		return client, err
	}

	session.Stdin = os.Stdin
	session.Stderr = os.Stderr
	session.Stdout = os.Stdout

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	err = session.RequestPty("xterm", 80, 40, modes)

	if err != nil {
		return client, err
	}

	err = session.Shell()

	if err != nil {
		return client, err
	}

	err = session.Wait()

	return client, err
}
