# gin 博客项目学习实践：后台部分
Go： 1.16.3

---

#### 项目架构
```
├─.idea
├─api              -----》》》 视图层view
│  └─v1
├─config           -----》》》 配置文件
├─log
├─middleware       -----》》》 中间件
├─model            -----》》》 模型操作：mysql
├─routes           -----》》》 路由配置
└─utils            -----》》》 
    ├─errmsg       -----》》》 状态码
    └─validator    -----》》》 翻译功能
```

#### 安装依赖
> 执行前先设置好环境变量“GOPROXY”和“GO111MODULE”：
- export GO111MODULE=on
- export GOPROXY=https://goproxy.io
- 或者export GOPROXY=https://goproxy.cn
```
go get -u github.com/gin-gonic/gin   web框架
go get github.com/Unknwon/goconfig
go get -u gopkg.in/ini.v1            init 依赖
go get -u github.com/jinzhu/gorm     数据库驱动软件
go get -u gorm.io/driver/mysql       mysql 操作依赖
go get -u github.com/dgrijalva/jwt-go  jwt token
go get -u github.com/sirupsen/logrus  日志处理
go get -u github.com/lestrrat-go/file-rotatelogs 日志分割
go get -u github.com/rifflock/lfshook
go get -u github.com/qiniu/go-sdk/v7   7牛云上传
go get -u github.com/gin-contrib/cors  跨域

```
项目学习链接：https://gitee.com/wejectchan/ginblog/tree/master