// git-get clones repos according to their dir structure in GitHub.
// Can be accessed as a subcommand of git. Ex: git get https://github.com/satyarohith/utils
package main

import (
	"fmt"
	"github.com/tcnksm/go-gitconfig"
	"log"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// This piece of code is obtained from https://github.com/atotto/clipboard/blob/master/clipboard_darwin.go
func copyToClipboard(text string) error {
	copyCmd := exec.Command("pbcopy")

	if runtime.GOOS == "linux" {
		copyCmd = exec.Command("xclip", "-in", "-selection", "clipboard")
	}

	in, err := copyCmd.StdinPipe()
	if err != nil {
		return err
	}

	if err := copyCmd.Start(); err != nil {
		return err
	}
	if _, err := in.Write([]byte(text)); err != nil {
		return err
	}
	if err := in.Close(); err != nil {
		return err
	}

	return copyCmd.Wait()
}

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	baseDir := homeDir + "/c/"

	args := os.Args[1:]
	cloneLink := args[0]
	var orgName, repoName string
	var absOrgPath, repoPath string
	var hostname string = "github.com"

	if strings.HasPrefix(cloneLink, "git@") {
		urlDetails := strings.Split(cloneLink, ":")
		hostname = strings.Split(urlDetails[0], "@")[1]
		names := strings.Split(urlDetails[1], "/")
		orgName = names[0]
		repoName = names[1]
	} else if strings.HasPrefix(cloneLink, "https://") {
		url, err := url.Parse(cloneLink)
		if err != nil {
			log.Fatal(err)
		}

		hostname = url.Hostname()
		names := strings.Split(url.Path, "/")
		orgName = names[1]
		repoName = names[2]
	} else if len(strings.Split(cloneLink, "/")) == 2 {
		names := strings.Split(cloneLink, "/")
		orgName = names[0]
		repoName = names[1]
		cloneLink = "git@" + hostname + ":" + orgName + "/" + repoName
	} else if len(strings.Split(cloneLink, "/")) == 1 {
		username, err := gitconfig.GithubUser()
		if err != nil {
			log.Fatal(err)
		}

		orgName = username
		repoName = cloneLink
		cloneLink = "git@" + hostname + ":" + orgName + "/" + repoName
	} else {
		log.Fatal("Could not parse the URL.")
	}

	// Organisation's absolute path
	absOrgPath = baseDir + hostname + "/" + orgName
	// Repo's absolute path
	repoPath = absOrgPath + "/" + repoName
	// Create all directories
	err = os.MkdirAll(absOrgPath, 0700)
	if err != nil {
		log.Fatal(err)
	}

	gitArgs := append([]string{"clone", cloneLink}, args[1:]...)
	// Executing git clone
	cmd := exec.Command("git", gitArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// Go to the organisation's directory and then clone
	cmd.Dir = absOrgPath
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// Copy the absolute path to clipboard on macos.
	err = copyToClipboard(repoPath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Cloned to", repoPath, "(in clipboard)")
}
