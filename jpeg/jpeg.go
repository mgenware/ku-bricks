package jpeg

import (
	"github.com/mgenware/ku-builder"
)

var repo = &ku.RepoInfo{
	Url:  "https://github.com/libjpeg-turbo/libjpeg-turbo",
	Name: "libjpeg-turbo",
	Tag:  "3.2.1",
}

func BuildJpeg(be *ku.BuildEnv) {
	p := ku.NewCMakeProject(repo, be, ku.LibTypeStatic)
	p.Init(&ku.ProjectInitOptions{
		Args: []string{
			"-DENABLE_STATIC=1",
			"-DENABLE_SHARED=0",
			"-DWITH_JAVA=0",
		},
	})
	p.Build()
	p.Install("libjpeg", nil)
}
