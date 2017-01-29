package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func prepareGatlingDirectory(fromURL string, gatlingRoot string) error {

	// create staging directory
	timestamp := time.Now().UnixNano()
	stageDirectoryName := fmt.Sprintf("/tmp/staging%d", timestamp)
	Log("creating staging directory: %s", stageDirectoryName)
	if err := os.Mkdir(stageDirectoryName, 0777); err != nil {
		return err
	}

	// download artifact into staging directory
	localArtifactFile := fmt.Sprintf("%s/artifact.tar.bz2", stageDirectoryName)
	if err := downloadArtifact(fromURL, localArtifactFile); err != nil {
		return err
	}

	Log("Extracting artifact..")
	if err := execute([]string{}, "tar", "-xjf", localArtifactFile, "-C", stageDirectoryName); err != nil {
		return err
	}

	// TODO Validate extracted artifact

	// copy files from extracted artifact into gatling directory
	Log("Copying simulations..")
	// TODO copy * from sim/* to simulations/* without using bash file interpolation.
	if err := execute([]string{}, "cp", "-rv", fmt.Sprintf("%s/sim/com", stageDirectoryName), fmt.Sprintf("%s/user-files/simulations/", gatlingRoot)); err != nil {
		return err
	}
	Log("Copying libraries..")
	if err := execute([]string{}, "cp", "-rv", fmt.Sprintf("%s/lib/", stageDirectoryName), fmt.Sprintf("%s/", gatlingRoot)); err != nil {
		return err
	}
	Log("Copying misc other files..")
	if err := execute([]string{}, "cp", "-rv", fmt.Sprintf("%s/misc/", stageDirectoryName), fmt.Sprintf("%s/", gatlingRoot)); err != nil {
		return err
	}

	return nil
}

func downloadArtifact(fromURL string, toLocalFile string) error {
	Log("Downloading artifact from url=%s to local file=%s", fromURL, toLocalFile)

	out, err := os.Create(toLocalFile)
	if err != nil {
		return err
	}

	resp, err := http.Get(fromURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
