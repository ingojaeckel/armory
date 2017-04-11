package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

func executeGatling(simClass string) error {
	return execute(conf.EnvVars, conf.GatlingBinary, "-s", simClass, "-on", "testname", "-rd", "test", "-nr", "-rf", conf.ResultsFolder)
}

func execute(env []string, command string, args ...string) error {
	Log("execute(env=%v, command=%s, args=%v)", env, command, args)

	cmd := exec.Command(command, args...)
	outReader, err := cmd.StdoutPipe()
	errReader, _ := cmd.StderrPipe()
	if err != nil {
		return err
	}

	stdoutScanner := bufio.NewScanner(outReader)
	go func() {
		for stdoutScanner.Scan() {
			Log(fmt.Sprintf("[cmd-stdout] %s", stdoutScanner.Text()))
		}
	}()
	errScanner := bufio.NewScanner(errReader)
	go func() {
		for errScanner.Scan() {
			Log(fmt.Sprintf("[cmd-stderr] %s", errScanner.Text()))
		}
	}()
	return cmd.Run()
}
