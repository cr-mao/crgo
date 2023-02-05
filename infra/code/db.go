package code

const (
	// ErrConnectDB - 500: Init db error.
	ErrConnectDB int = iota + 100601

	// ErrConnectGRPC - 500: Connect to grpc error.
	ErrConnectGRPC
)
