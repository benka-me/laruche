package git

import (
	"fmt"
	gogit "gopkg.in/src-d/go-git.v4"
	"os"
	"strings"
)

var gopath = os.Getenv("GOPATH")

func Clone(repo string) (*gogit.Repository, error) {
	path, url := "", ""
	if strings.Contains(repo, "https") {
		r := strings.Replace(repo, "https://", "", -1)
		path = fmt.Sprintf("%s/src/%s", gopath, r)
		url = repo
	} else {
		path = fmt.Sprintf("%s/src/%s", gopath, repo)
		url = fmt.Sprintf("https://%s", repo)
	}

	return gogit.PlainClone(path, false, &gogit.CloneOptions{
		URL:               url,
		RecurseSubmodules: gogit.DefaultSubmoduleRecursionDepth,
	})
}
