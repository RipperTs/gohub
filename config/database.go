package config

import (
	"gohub/pkg/config"
)

func init() {

	config.Add("database", func() map[string]interface{} {
		return map[string]interface{}{

			// 默认数据库
			"connection": config.Env("DB_CONNECTION", "mysql"),
			// 表前缀
			"table_prefix": config.Env("DB_TABLE_PREFIX", "gohub_"),

			"mysql": map[string]interface{}{

				// 数据库连接信息
				"host":     config.Env("DB_HOST", "127.0.0.1"),
				"port":     config.Env("DB_PORT", "3306"),
				"database": config.Env("DB_DATABASE", "gohub"),
				"username": config.Env("DB_USERNAME", ""),
				"password": config.Env("DB_PASSWORD", ""),
				"charset":  "utf8mb4",

				// 设置空闲连接池中连接的最大数量
				"max_idle_connections": config.Env("DB_MAX_IDLE_CONNECTIONS", 100),
				// 设置打开数据库连接的最大数量
				"max_open_connections": config.Env("DB_MAX_OPEN_CONNECTIONS", 25),
				//设置了连接可复用的最大时间
				"max_life_seconds": config.Env("DB_MAX_LIFE_SECONDS", 5*60),
			},

			"sqlite": map[string]interface{}{
				"database": config.Env("DB_SQL_FILE", "database/database.db"),
			},
		}
	})
}
