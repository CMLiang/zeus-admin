# Zeus 宙斯权限&账号管理系统

[![golang](https://img.shields.io/badge/golang-1.12.1-green.svg?style=plastic)](https://www.golang.org/)
[![casbin](https://img.shields.io/badge/casbin-1.8.1-brightgreen.svg?style=plastic)](https://github.com/casbin/casbin)

## 项目介绍
> `Zeus 宙斯` 权限后台，为企业提供统一后台权限管理私有化Saas云服务。    
> - 项目使用 `golang gin + vue-element-admin` 框架开发，用`jwt + casbin`做权限管理,提供OAuth2.0 的Restful Api 接口。
> - 为企业后台系统提供统一登陆鉴权、菜单管理、权限管理、组织架构管理、员工管理、配置中心、日志管理等。
> - 支持企业微信、钉钉登陆和同步企业组织架构。
> - 统一管理员工入离职，强化权限审批流程化。
> - 打通开源软件、付费Saas软件，企业内部开发系统等，包括不限于jenkis、jira、gitlab、confluence、禅道、企业邮箱、OA、CRM、财务软件、企业Sass云服务等内外部系统，解决企业多个软件和平台账号不同步的痛点。     
> - `打造统一开放平台生态标准，为企业引进外部系统不再困难。`

## Features （目前实现功能）
- 登录/登出
- 权限管理
    - 用户管理(人员管理)
    - 角色管理(功能权限管理)
    - 部门管理
    - 项目管理
    - 菜单管理
    - 数据权限管理
- 个人帐户
    - 第三方登陆（钉钉）
    - 安全设置（[Google 2FA 二次验证](http://www.ruanyifeng.com/blog/2017/11/2fa-tutorial.html)）
    - 支持LDAP

### 快速开始
> 该操作在linux 下生效，需要golang 1.11+ & node v9 + 编译环境,设置git clone 权限

> 前后统一访问入口部署(前后统一)
````
git clone http://gitlab.lingeyun.com/lyb/zeus-admin.git
export GOPROXY=https://goproxy.cn
export GO111MODULE=on
#后端编译
go build -o zeus
#前端编译
cd pkg/webui
npm install
npm run build:work
cd ~/zeus-admin

export MYSQL_USERNAME=root
export MYSQL_PASSWORD=123456
export MYSQL_HOST=127.0.0.1
export MYSQL_DB=zeus
export MYSQL_PORT=3306
export REDIS_HOST=127.0.0.1
export REDIS_PORT=6379
export REDIS_PASSWORD=""

./zeus server -c ./config/in-local.yaml

````

> 前后不同入口部署(前后分离)

````
git clone http://gitlab.lingeyun.com/lyb/zeus-admin.git
export GOPROXY=https://goproxy.cn
export GO111MODULE=on
#后端编译
go build -o zeus
#前端编译
cd pkg/webui
npm install
#正常情况下，会生成dist目录，可自己部署web服务器(如nginx)，提供前端服务
npm run build:prod
cd ~/zeus-admin

export MYSQL_USERNAME=root
export MYSQL_PASSWORD=123456
export MYSQL_HOST=127.0.0.1
export MYSQL_DB=zeus
export MYSQL_PORT=3306
export REDIS_HOST=127.0.0.1
export REDIS_PORT=6379
export REDIS_PASSWORD=""

#修改in-local.yamln内部的project.merge为false,然后再启动
./zeus server -c ./config/in-local.yaml --cors=true
````


# 数据移值

```bash
# 执行 sql 语句
mysql> source ./scripts/init.sql;
```

## Git 工作流

[Git 协作工作流](docs/zh/CONTRIBUTING.md)

# openssl jwt 密钥生成
[openssl jwt 密钥](docs/zh/GenrsaKey.md)

# 接入权限系统 client demo
* [python-client](https://github.com/bullteam/zeusclient-python) 已提供
* [php-client](https://github.com/bullteam/zeusclient-php) 已提供
* [java-client](https://github.com/bullteam/zeusclient-java) 已提供
* [go-client](https://github.com/bullteam/zeusclient-go) 暂缺

## 开发者

* [wutongci](http://github.com/wutongci)
* [funlake](https://github.com/funlake)
* [Hyman](https://github.com/zhengcog)
* [severHo](https://github.com/qq330967496)
* [hodor-cn](https://github.com/hodor-cn)

更多请进入我们的官网了解我们  [公牛开源战队](http://www.bullteam.cn)