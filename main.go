package main

import (
	"fmt"
	"os"
	"os/user"
	"path"
)

func GoXGetRun(req string, gopath string) bool {
	cur, err := user.Current()
	if err != nil {
		fmt.Println("[E] 获取当前用户信息失败！")
		return false
	}

	_, err := os.Stat(path.Join(cur.HomeDir, CONFIG_XML_PATH))
	if os.IsNotExist(err) {
		_downloadConfigXML()
	} else if err != nil {
		fmt.Println("[E] 获取配置文件信息失败")
		return false
	}

	gtr := NewGetter(req, gopath)
	gtr.Get()
}
