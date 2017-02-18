package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

var env = []string{fmt.Sprintf("JAVA_OPTS=\"-Xms%s -Xmx%s\"", conf.Base.GatlingInitialMemory, conf.Base.GatingMaxMemory)}
var gatlingBinary = fmt.Sprintf("%s/bin/gatling.sh", conf.Base.GatlingRoot)
var resultsFolder = fmt.Sprintf("%s/results/testname", conf.Base.GatlingRoot)

func executeGatling(simClass string) error {
	return execute(env, gatlingBinary, "-s", simClass, "-on", "testname", "-rd", "test", "-nr", "-rf", resultsFolder)
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
