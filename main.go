package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/go-git/go-git/v5"
)

func main() {
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(tmpDir)

	configPath := filepath.Join(tmpDir, "config")
	if r, err := git.PlainClone(configPath, false, &git.CloneOptions{
		URL: "https://github.com/qlik-oss/qliksense-k8s",
	}); err != nil {
		panic(err)
	} else if workTree, err := r.Worktree(); err != nil {
		panic(err)
	} else if hash, err := r.ResolveRevision("v0.0.8"); err != nil {
		panic(err)
	} else if err := workTree.Checkout(&git.CheckoutOptions{
		Hash: *hash,
	}); err != nil {
		panic(err)
	}

	cmd := exec.Command("git", "status")
	cmd.Dir = configPath
	if out, err := cmd.Output(); err != nil {
		panic(err)
	} else {
		fmt.Println(string(out))
	}
}
