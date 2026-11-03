package dav1d

import (
	"github.com/mgenware/ku-builder"
)

var repo = &ku.RepoInfo{
	Url:  "https://github.com/videolan/dav1d",
	Name: "dav1d",
	Tag:  "1.5.4",
}

func BuildDav1d(be *ku.BuildEnv) {
	args := []string{
		"-Denable_tools=false",
		"-Denable_tests=false",
	}
	p := ku.NewMesonProject(repo, be, ku.LibTypeStatic)
	p.Init(&ku.ProjectInitOptions{
		Args: args,
	})
	p.Build()
	p.Install("libdav1d", nil)
}
