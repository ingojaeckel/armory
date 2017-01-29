package main

import (
	"fmt"
)

func executeSSH(sshUsername string, sshHost string, pathPemFile string, remoteCmd string) error {
	remoteDest := fmt.Sprintf("%s@%s", sshUsername, sshHost)
	return execute([]string{}, "/usr/bin/ssh", "-o", "StrictHostKeyChecking=no", "-i", pathPemFile, remoteDest, remoteCmd)
}

func executeScp(sshUsername string, sshHost string, pathPemFile string, fromPath string, toPath string) error {
	remoteDest := fmt.Sprintf("%s@%s:%s", sshUsername, sshHost, toPath)
	return execute([]string{}, "/usr/bin/scp", "-p", "-o", "StrictHostKeyChecking=no", "-i", pathPemFile, fromPath, remoteDest)
}

func scpWithRetries(sshUsername string, sshHost string, pathPemFile string, fromPath string, toPath string) error {
	return retry(true, 1000, 2000, 3, func() (bool, error) {
		if err := executeScp(sshUsername, sshHost, pathPemFile, fromPath, toPath); err != nil {
			Log("Failed to copy latest version of executable to worker %s: %s", sshHost, err.Error())
			return false, err
		}
		return true, nil
	})
}
