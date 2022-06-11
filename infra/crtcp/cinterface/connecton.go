package cinterface

import "net"

type ConnectionInterface interface {
	//启动
	Start()
	//关闭
	Stop()
	// 获得tcp 连接
	GetTCPConnection() *net.TCPConn
	// 获得连接id
	GetConnID() uint32

	RemoteAddr() net.Addr
}


