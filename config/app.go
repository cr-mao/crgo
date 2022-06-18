package config

import "crgo/infra/conf"

//应用配置
func init() {
	conf.Add("app", func() map[string]interface{} {
		return map[string]interface{}{
			// 应用名称
			"name": conf.Env("app_name", "cr"),

			// 当前环境，用以区分多环境，一般为 local, stage, production, test
			"env": conf.Env("app_env", "local"),

			// 是否进入调试模式
			"debug": conf.Env("debug", false),

			// 应用服务端口
			"http_port":    conf.Env("http_port", "3000"),
			"http_addr":    conf.Env("http_addr", "127.0.0.1"),
			"http_connect": conf.Env("http_connect", "127.0.0.1:3000"),
			"grpc_port":    conf.Env("grpc_port", "3001"),
			"grpc_addr":    conf.Env("grpc_addr", "127.0.0.1"),
			"grpc_connect": conf.Env("grpc_connect", "127.0.0.1:3001"),


			// 加密会话、JWT 加密
			"jwt_key": conf.Env("jwt_key", "33446a9dcf9ea060a0a6532b166da32f304af0de"),

			// 设置时区，JWT 里会使用，日志记录里也会使用到
			"timezone": conf.Env("timezone", "Asia/Shanghai"),
		}
	})
}
