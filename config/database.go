package config

import "crgo/infra/conf"

//数据库配置
func init() {
	//暂时没用
	conf.Add("database1", func() map[string]interface{} {
		return map[string]interface{}{
			"connection": conf.Env("DB_CONNECTION", "mysql"),
			"mysql": map[string]interface{}{
				// 数据库连接信息
				"host":     conf.Env("DB_HOST", "127.0.0.1"),
				"port":     conf.Env("DB_PORT", "3306"),
				"database": conf.Env("DB_DATABASE", "test"),
				"username": conf.Env("DB_USERNAME", "root"),
				"password": conf.Env("DB_PASSWORD", ""),
				"charset":  "utf8mb4",
				// 连接池配置
				"max_idle_connections": conf.Env("DB_MAX_IDLE_CONNECTIONS", 100),
				"max_open_connections": conf.Env("DB_MAX_OPEN_CONNECTIONS", 25),
				"max_life_seconds":     conf.Env("DB_MAX_LIFE_SECONDS", 5*60),
			},
		}
	})
}
