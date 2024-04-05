package lib

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

// TestDownloadFromURL_FileNameMatch : Check expected filename exist when downloaded
func TestDownloadFromURL_FileNameMatch(t *testing.T) {

	hashiURL := "https://releases.hashicorp.com/terraform/"
	installVersion := "terraform_"
	tempDir := t.TempDir()
	installPath := fmt.Sprintf(tempDir + string(os.PathSeparator) + ".terraform.versions_test")
	macOS := "_darwin_amd64.zip"

	home, err := homedir.Dir()
	if err != nil {
		logger.Fatalf("Could not detect home directory")
	}

	logger.Infof("Current home directory: %q", home)
	var installLocation = ""
	if runtime.GOOS != "windows" {
		installLocation = filepath.Join(home, installPath)
	} else {
		installLocation = installPath
	}
	logger.Infof("Install Location: %v", installLocation)

	// create /.terraform.versions_test/ directory to store code
	if _, err := os.Stat(installLocation); os.IsNotExist(err) {
		t.Logf("Creating directory for terraform: %v", installLocation)
		err = os.MkdirAll(installLocation, 0755)
		if err != nil {
			t.Logf("Unable to create directory for terraform: %v", installLocation)
			t.Error("Test fail")
		}
	}

	/* test download old terraform version */
	lowestVersion := "0.11.0"

	url := hashiURL + lowestVersion + "/" + installVersion + lowestVersion + macOS
	expectedFile := filepath.Join(installLocation, installVersion+lowestVersion+macOS)
	installedFile, errDownload := DownloadFromURL(installLocation, url)

	if errDownload != nil {
		t.Logf("Expected file name %v to be downloaded", expectedFile)
		t.Error("Download not possible (unexpected)")
	}

	if installedFile == expectedFile {
		t.Logf("Expected file %v", expectedFile)
		t.Logf("Downloaded file %v", installedFile)
		t.Log("Download file matches expected file")
	} else {
		t.Logf("Expected file %v", expectedFile)
		t.Logf("Downloaded file %v", installedFile)
		t.Error("Download file mismatches expected file (unexpected)")
	}

	//check file name is what is expected
	_, err = os.Stat(expectedFile)
	if err != nil {
		t.Logf("Expected file does not exist %v", expectedFile)
	}

	t.Cleanup(func() {
		defer os.Remove(tempDir)
		logger.Infof("Cleanup temporary directory %q", tempDir)
	})
}

// // TestDownloadFromURL_Valid : Test if https://releases.hashicorp.com/terraform/ is still valid
func TestDownloadFromURL_Valid(t *testing.T) {

	hashiURL := "https://releases.hashicorp.com/terraform/"

	url, err := url.ParseRequestURI(hashiURL)
	if err != nil {
		t.Errorf("Invalid URL %v [unexpected]", err)
	} else {
		t.Logf("Valid URL from %v [expected]", url)
	}
}
