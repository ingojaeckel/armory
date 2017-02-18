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
	stageDirectoryName := fmt.Sprintf("/tmp/staging%d", time.Now().UnixNano())
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
	return cp([]pair{
		// TODO copy * from sim/* to simulations/* without using bash file interpolation.
		pair{fmt.Sprintf("%s/sim/com", stageDirectoryName), fmt.Sprintf("%s/user-files/simulations/", gatlingRoot)},
		pair{fmt.Sprintf("%s/lib/", stageDirectoryName), fmt.Sprintf("%s/", gatlingRoot)},
		pair{fmt.Sprintf("%s/misc/", stageDirectoryName), fmt.Sprintf("%s/", gatlingRoot)},
	})
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
	return err
}
