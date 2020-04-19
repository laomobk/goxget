# GoXGet
## Golang包镜像库下载器

#### golang.org 登不上？
#### github.com 克隆太慢？
### 试试 GoXGet 吧！

**GoXGet 的镜像库尽量采用码云的库，保证了在国内的克隆速度。**

## GoXGet 是如何运作的？

1. GoXGet 在运行的时候会进行配置文件搜索，若搜索不到则会同步远程库的配置文件

2. 接下来会搜索配置文件：
    * 若找到镜像库，则会使用 git 克隆到 `$GOPATH/src` 当中（确保计算机安装了git）。
    * 若找不到镜像库，则会使用 go get 来进行下载，go get 的参数为 `-v` `-u`。

3. 安装完成

## 维护配置文件

GoXGet 同步的配置文件是采用库 `https://gitee.com/LaomoBK/goxget-config.git` 中的配置文件。
若发现镜像版本过旧，欢迎进行更新😃，配置文件规范在该库的 README 文件中。


