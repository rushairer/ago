# apigg

## a

### 模块

| 模块        | 说明         |
| ----------- | ------------ |
| sso/core    | 核心算法模块 |
| sso/command | 命令行模块   |
| sso/cron    | 定时任务模块 |
| sso/api     | API 模块     |

### 如何开发

```bash
# 使用air进行热更新
go install github.com/cosmtrek/air@latest

# 热更新运行cron模块
air -d -c .air/cron

# 热更新运行web模块
air -d -c .air/web

# 运行命令行
go run command/main.go client:create NAME DOMAIN

# migrate db
migrate create -ext sql -dir core/databases/migrations -seq create_users_table
migrate -database mysql://$MYSQL_DSN -path core/databases/migrations up

# godoc生成文档
go install -v golang.org/x/tools/cmd/godoc@latest
godoc

```

### 如何构建

```bash
docker build -t apigg_api --build-arg BUILD_PATH=/web .
docker build -t apigg_cron --build-arg BUILD_PATH=/cron .
```

### 其他

[sqlx 文档](http://jmoiron.github.io/sqlx/)
