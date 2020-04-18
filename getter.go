package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

type Getter struct {
	req string
	gopath string
	repos []*Repo
}

func NewGetter(req string, gopath string, repos []*Repo) *Getter {
	g := new(Getter)

	g.req = req
	g.gopath = gopath
	g.repos = repos

	return g
}

func (self *Getter) Get() (ok bool) {
	req := self.req

	var repo *Repo

	for _, r := range self.repos {
		if r.pkgName == req {
			repo = r
		}
	}

	if repo == nil {
		fmt.Println("[W] 无替换仓库，使用默认 go get")

		gexe, err := exec.LookPath("go")
		if err != nil {
			fmt.Println("[E] 找不到 go 可执行文件")
			return false
		}

		gcmd := exec.Command(gexe, "get", "-v", "-u", req)
		gcmd.Stdout = os.Stdout
		gcmd.Stdin = os.Stdin
		gcmd.Stderr = os.Stderr
		err  = gcmd.Run()

		if err != nil {
			fmt.Println("[E] Go get 命令执行出错")
			return false
		}

		return true
	}

	gitexe, err := exec.LookPath("git")

	if err != nil {
		fmt.Println("[E] 找不到 git ，使用前请安装")
		return false
	}

	clonePath := path.Join(self.gopath, self.req)

	cmd := exec.Command(gitexe, repo.replRepoUrl, clonePath)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	fmt.Printf("即将使用来自 %s 的库 %s 来替代 %s", repo.website, repo.replRepoUrl, repo.pkgName)

	err = cmd.Run()
	if err != nil {
		fmt.Println("[E] git 指令执行出错")
		return false
	}

	return true
}