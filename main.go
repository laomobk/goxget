package main

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"regexp"
)

func GoXGetRun(req string, gopath string) bool {
	cur, err := user.Current()
	if err != nil {
		fmt.Println("[E] 获取当前用户信息失败！")
		return false
	}

	_, err = os.Stat(path.Join(cur.HomeDir, CONFIG_XML_PATH))
	if os.IsNotExist(err) {
		fmt.Printf("[W] 未找到配置文件，将从默认地址(%s)下载\n", DEFAULT_CONFIG_XML_URL)

		if !downloadConfigXML() {
			fmt.Println("[E] 配置文件下载失败")
			return false
		}

		fmt.Println("[E] 配置文件下载完成")

	} else if err != nil {
		fmt.Println("[E] 获取配置文件信息失败")
		return false
	}

	repos := listAllPkgConfig()

	gtr := NewGetter(req, gopath, repos)

	return gtr.Get()
}

func GoXGetCheckUpdate() {
	fmt.Println("[W] 检查配置文件更新...")

	ucode, msg := checkConfigUpdate()
	if ucode == 1 {
		fmt.Printf("[U] %s\n", msg)
		if askYN("是否更新？(Y/n): ") {
			if !downloadConfigXML() {
				fmt.Println("[W] 更新失败")
			} else {
				fmt.Println("[G] 更新成功")
			}
		}
	} else if ucode == -1 {
		fmt.Printf("[W] %s\n", msg)
	} else if ucode == 0 {
		fmt.Printf("[G] %s\n", msg)
	}
}

func GoXGetListRepo(searchPattern string) {
	search := searchPattern != ""
	var rexp *regexp.Regexp

	if search {
		var err error
		rexp, err = regexp.Compile(searchPattern)
		search = err == nil
	}

	fmt.Println("[G] 列出可用的包和其替换库...\n")

	cfg := readConfigXml()
	if cfg == nil {
		return
	}
	root := cfg.SelectElement("xrepos")

	v := getVersionNumber(root)

	fmt.Printf("版本号：%d\n\n", v)

	repos := listAllPkgConfig()

	var searchCount int

	for _, repo := range repos {
		if search {
			if !rexp.MatchString(repo.pkgName) {
				continue
			}
		}

		searchCount++
		fmt.Printf("- %s\n", repo.pkgName)
		fmt.Printf("	替换库：%s\n", repo.replRepoUrl)
		fmt.Printf("	托管平台：%s\n", repo.website)
		fmt.Println("")
	}

	if !search {
		fmt.Printf("读取完毕，一共 %d 个\n", len(repos))
	} else {
		fmt.Printf("读取完毕，搜索到 %d 个条目\n", searchCount)
	}
}

func main() {
	carg := parseArg()
	if carg == nil {
		printUsage()
		os.Exit(1)
	}

	if carg.update {
		GoXGetCheckUpdate()
		return

	} else if carg.list {
		GoXGetListRepo("")
		return

	} else if carg.search {
		GoXGetListRepo(carg.pkgName)
		return

	} else if carg.help {
		printUsage()
		return
	} else if carg.version {
		fmt.Println(VERSION)
		return
	}

	pkg := carg.pkgName

	if os.Getenv("GOPATH") == "" {
		fmt.Println("[E] 请设置 GOPATH 环境变量")
		return
	}

	gopath := os.Getenv("GOPATH")

	code := 0

	if !GoXGetRun(pkg, gopath) {
		code = 1
	}

	os.Exit(code)
}
