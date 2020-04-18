package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/user"
	"path"

	"github.com/beevik/etree"
)

func _require(attr string) {
	fmt.Printf("[W] repo元素缺少属性 '%s'\n", attr)
}

func _incomplete() {
	fmt.Printf("[E] 配置文件不完整\n")
}

func downloadConfigXML() bool {
	resp, err := http.Get(DEFAULT_CONFIG_XML_URL)
	if err != nil {
		fmt.Println("[E] 配置文件下载失败")
		return false
	}

	cur, err := user.Current()
	if err != nil {
		fmt.Println("[E] 获取当前用户Home目录失败")
		return false
	}

	f, err := os.Create(path.Join(cur.HomeDir, CONFIG_XML_PATH))
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		panic(err)
	}

	return true
}

/*
	A pkg element should be like this:

	<repo>
		<pkg-name>golang.org/x/sys</pkg-name>

		<repl-repo>
			https://gitee.com/LaomoBK/golang-x-term.git
		</repl-repo>

		<website>gitee</website>
	</repo>

*/
func _parseSingle(ele *etree.Element) *Repo {
	repo := new(Repo)

	if name := ele.SelectElement("pkg-name"); name != nil {
		repo.pkgName = name.Text()
	} else {
		_require("pkg-name")
	}

	if replUrl := ele.SelectElement("repl-repo"); replUrl != nil {
		repo.replRepoUrl = replUrl.Text()
	} else {
		_require("repl-repo")
	}

	if website := ele.SelectElement("website"); website != nil {
		repo.website = website.Text()
	} else {
		_require("website")
	}

	return repo
}

func readConfigXml() *etree.Document {
	doc := etree.NewDocument()

	cur, err := user.Current()
	if err != nil {
		fmt.Println("[E] 读取配置文件失败！")
		return nil
	}

	err = doc.ReadFromFile(path.Join(cur.HomeDir, CONFIG_XML_PATH))

	if err != nil {
		fmt.Println("[E] 读取配置文件失败！")
		return nil
	}

	return doc
}

/*
	a goxget config xml should be like this:

	<xrepos>
		<repo>...</repo>
		<repo>...</repo>
		...
	</xrepos>
*/
func listAllPkgConfig() []*Repo {
	repos := make([]*Repo, 0)
	doc := readConfigXml()

	if doc == nil {
		return repos
	}

	root := doc.SelectElement("xrepos")
	if root == nil {
		_incomplete()
		return repos
	}

	for _, repoEle := range root.SelectElements("repo") {
		repos = append(repos, _parseSingle(repoEle))
	}

	return repos
}
