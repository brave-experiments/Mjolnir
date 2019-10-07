package connect

import (
	"golang.org/x/crypto/ssh"
	"os"
	"os/exec"
)

type SshClient struct {
	Config   *ssh.ClientConfig
	hostname string
	user     string
	keyPath  string
}

type DialError struct {
	Message string
}

func (err DialError) Error() (message string) {
	return err.Message
}

func (sshClient *SshClient) New(user string, hostname string, keyPath string) {
	sshClient.user = user
	sshClient.hostname = hostname
	sshClient.keyPath = keyPath
}

func (sshClient *SshClient) Dial() (err error) {
	userAndHost := sshClient.user + "@" + sshClient.hostname

	cmdArgs := []string{userAndHost, "-i", sshClient.keyPath}
	cmd := exec.Command("ssh", cmdArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()

	return err
}
