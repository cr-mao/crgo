/*

配置读取的顺序为

+ fromENV
+ fromConfigFile

尽量将配置映射为 var，可以在其他地方直接调用如：

conf.Net.GRPC_ADDR

或

conf.IsDev()

*/
package conf


