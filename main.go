package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	// create an argument parser to accept the repo_url argument
	repoURL := flag.String("repo-url", "", "URL of the GitHub repository to clone")
	flag.Parse()

	// ensure that the repoURL argument is not empty
	if *repoURL == "" {
		fmt.Fprintf(os.Stderr, "Error: repo URL argument cannot be empty. The format has to be https://github.com/<owner>/<repo_name>.git\n")
		os.Exit(1)
	}

	// get the repository name from the URL
	ownerRepo := filepath.Base(*repoURL)
	if len(ownerRepo) == 0 || filepath.Ext(ownerRepo) != ".git" {
		fmt.Fprintf(os.Stderr, "Invalid repository URL: %s\n", *repoURL)
		os.Exit(1)
	}

	var owner, repo string
	dir := filepath.Dir(*repoURL)
	if dir != "." {
		ownerRepo = filepath.Base(dir) + "/" + ownerRepo
		owner, repo = filepath.Split(ownerRepo)
		owner = owner[:len(owner)-1]
	} else {
		repo = ownerRepo[:len(ownerRepo)-len(filepath.Ext(ownerRepo))]
	}

	// specify the directory of the cloned GitHub repository (without the .git suffix)
	repoDirectory := strings.TrimSuffix(repo, ".git")

	// check if the repository directory already exists
	if _, err := os.Stat(repoDirectory); os.IsNotExist(err) {
		// clone the repository to the current working directory
		cmd := exec.Command("git", "clone", *repoURL)
		if err := cmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error cloning repository: %v\n", err)
			os.Exit(1)
		}
	} else {
		fmt.Printf("The repository directory %s already exists, using it to get URLs.\n", repoDirectory)
	}

	// iterate through all files in the repository and its subdirectories (including symbolic links)
	filepath.Walk(repoDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error walking path %s: %v\n", path, err)
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".md" {
			return nil
		}
		// construct the full file URL relative to the repository root directory
		fileURL, err := filepath.Rel(repoDirectory, path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error getting relative file path for %s: %v\n", path, err)
			return err
		}
		// construct the full raw.githubusercontent.com URL and print it
		fullURL := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/main/%s", owner, repoDirectory, fileURL)
		fmt.Println(fullURL)
		return nil
	})
}
