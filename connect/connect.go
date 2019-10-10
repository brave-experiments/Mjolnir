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

func (sshClient *SshClient) Dial(args []string) (err error) {
	userAndHost := sshClient.user + "@" + sshClient.hostname
	err = sshClient.addIdentity(sshClient.keyPath)

	cmdArgs := []string{userAndHost, "-A", "-t"}
	cmdArgs = append(cmdArgs, args...)
	cmd := exec.Command("ssh", cmdArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()

	return err
}

func (sshClient *SshClient) addIdentity(identityFilePath string) (err error) {
	cmd := exec.Command("ssh-add", identityFilePath)
	err = cmd.Run()

	return err
}
