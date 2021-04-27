# dt-go-public
## Go： 学习实践

---
> 执行前先设置好环境变量“GOPROXY”和“GO111MODULE”：
- export GO111MODULE=on
- export GOPROXY=https://goproxy.io
- 或者export GOPROXY=https://goproxy.cn
```
"""
安装依赖
"""
go get -u github.com/gin-gonic/gin   web框架
go get github.com/Unknwon/goconfig
go get -u gopkg.in/ini.v1            init 依赖
go get -u github.com/jinzhu/gorm     数据库驱动软件
go get -u gorm.io/driver/mysql       mysql 操作依赖
go get -u github.com/dgrijalva/jwt-go  jwt token

```
### 创建项目
#### 项目学习链接：https://gitee.com/wejectchan/ginblog/tree/master