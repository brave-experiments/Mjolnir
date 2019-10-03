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

type privateKey struct {
	location string
	signer   ssh.Signer
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
		// Validate here
		return nil, nil
	}

	return ssh.Dial(sshClient.network, sshClient.hostname, sshClient.Config)
}
