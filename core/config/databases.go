package config

import "PACKAGENAME/utilities"

// Redis Configures
var RedisDSN string = utilities.GetEnv(
	"REDIS_DSN",
	"127.0.0.1:6379",
)

// Mysql Configures
var MysqlDSN string = utilities.GetEnv(
	"MYSQL_DSN",
	"root:123456@(127.0.0.1:3306)/sso?parseTime=true",
)

// Session Configures
// var SessionName string = utilities.GetEnv(
// 	"SESSION_NAME",
// 	"sso-session",
// )
// var SessionSecret string = utilities.GetEnv(
// 	"SESSION_SECRET",
// 	"sso-session-secret",
// )
