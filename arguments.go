package main

import (
	"os"
)

type ConsoleArg struct {
	pkgName string
	help    bool
	update  bool
	list    bool
	search  bool
	version bool
}

const USAGE = `
欢迎使用 GoXget ！

GoXget 是一个用于下载 golang.org/x/ 系列库的工具
GoXget 使用国内码云(gitee)作为主要镜像库的托管平台

使用说明：
	goxget <命令> [命令参数]

可用命令：
	update              更新配置文件
	list                列出替换库
	get [包名]          尽可能从镜像库获取包
	list                列出可用的镜像库
	search [搜索样式]   搜索镜像库配置文件（支持正则表达式）
	help                获取帮助
	version             查看 GoXGet 版本
`

func printUsage() {
	println(USAGE)
}

func parseArg() *ConsoleArg {
	carg := new(ConsoleArg)

	var parseGet bool
	var haveArg bool

	for _, arg := range os.Args[1:] {
		switch arg {

		case "update":
			if haveArg {
				return nil
			}
			carg.update = true
			haveArg = true

		case "list":
			if haveArg {
				return nil
			}
			carg.list = true
			haveArg = true

		case "get":
			if haveArg {
				return nil
			}
			parseGet = true
			haveArg = true

		case "search":
			if haveArg {
				return nil
			}
			haveArg = true
			parseGet = true
			carg.search = true

		case "help":
			if haveArg {
				return nil
			}
			haveArg = true
			carg.help = true

		case "version":
			if haveArg {
				return nil
			}
			haveArg = true
			carg.version = true

		default:
			if parseGet {
				carg.pkgName = arg
				parseGet = false
			}
		}

	}

	if parseGet || !haveArg {
		return nil
	}

	return carg
}
