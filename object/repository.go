package object

import (
	"os"
	"path"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
)

// InitRepo intialize Github repository required by WAFLab
func InitRepo() {
	// make directory for storing repository
	os.MkdirAll("repos", os.ModePerm)
	crsPath := path.Join("repos", "coreruleset")
	wafbenchPath := path.Join("repos", "WAFBench")

	// initialize coreruleset
	// check if crsPath has been created
	if _, err := os.Stat(crsPath); os.IsNotExist(err) {
		g, err := git.PlainClone(crsPath, false, &git.CloneOptions{
			URL:      "https://github.com/coreruleset/coreruleset.git",
			Progress: os.Stdout,
		})
		if err != nil {
			panic(err)
		}

		w, err := g.Worktree()
		if err != nil {
			panic(err)
		}

		err = g.Fetch(&git.FetchOptions{
			RefSpecs: []config.RefSpec{"refs/*:refs/*", "HEAD:refs/heads/HEAD"},
		})
		if err != nil {
			panic(err)
		}

		w.Checkout(&git.CheckoutOptions{
			Branch: plumbing.ReferenceName("refs/heads/v3.2/master"),
		})
	}

	// initialize coreruleset
	// check if wafbencPath has been created
	if _, err := os.Stat(wafbenchPath); os.IsNotExist(err) {
		_, err := git.PlainClone(wafbenchPath, false, &git.CloneOptions{
			URL:      "https://github.com/microsoft/WAFBench.git",
			Progress: os.Stdout,
		})
		if err != nil {
			panic(err)
		}
	}
}
