package conf

var (
	initializers []func() // 默认配置
	after        []func() // 读取配置后，设置
)

func AfterInit(f ...func()) {
	after = append(after, f...)
}
func OnInitialize(f ...func()) {
	initializers = append(initializers, f...)
}
