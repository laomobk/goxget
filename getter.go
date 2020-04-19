package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

type Getter struct {
	req    string
	gopath string
	repos  []*Repo
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

	clonePath := path.Join(self.gopath, "src", self.req)

	_, err := os.Stat(clonePath)
	if !os.IsNotExist(err) {
		fmt.Println("[T] 同名包已存在")
		return true
	}

	fmt.Println("[N] 正在寻找可替换的库")

	for _, r := range self.repos {
		if r.pkgName == req {
			repo = r
		}
	}

	if repo == nil {
		gexe, err := exec.LookPath("go")
		if err != nil {
			fmt.Println("[E] 找不到 go 可执行文件")
			return false
		}
		fmt.Println("[W] 无替换仓库，使用默认 go get")

		if !askYN("\n是否继续？(Y/n)") {
			return false
		}

		gcmd := exec.Command(gexe, "get", "-v", "-u", req)
		gcmd.Stdout = os.Stdout
		gcmd.Stdin = os.Stdin
		gcmd.Stderr = os.Stderr
		err = gcmd.Run()

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

	cmd := exec.Command(gitexe, "clone", "--recurse-submodules", repo.replRepoUrl, clonePath)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	fmt.Printf("替代包信息：\n\t库：%s\n\t托管平台：%s\n",
		repo.replRepoUrl, repo.website)
	fmt.Printf("原包信息：\n\t包名：%s\n", repo.pkgName)

	if !askYN("\n是否继续？(Y/n): ") {
		return false
	}

	err = cmd.Run()
	if err != nil {
		fmt.Println("[E] git 指令执行出错")
		return false
	}

	return true
}
