
* 即将重构，TODO: 前后端分离，go只负责接口

## 项目构建方法：

3. 修改配置文件 `conf.json` , 配置文件名字也可任意

4. 编译 `[... Bubbles]$ go build -o bb.server .` 在项目目录下构建

5. 启动 `nohup ./bb.server -c conf.json &`  配置文件-c指定

* 注意go版本 `go version go1.13.1` , 至少要在 `1.11`以上才支持 `go mod`

// TODO: 使用Docker
bb.server -c conf.json
