package connect

import (
	"golang.org/x/crypto/ssh"
	"io/ioutil"
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

	return nil
}

func (sshClient *SshClient) Dial() (client *ssh.Client, err error) {
	if nil == sshClient.Config {
		return nil, DialError{Message: "Config not present"}
	}

	return ssh.Dial(sshClient.network, sshClient.hostname, sshClient.Config)
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
