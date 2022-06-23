package cnet

type Server struct {
	IP        string
	Port      int
	IPVersion string //ip版本:   tcp4,tcp6,
	Name      string //服务器名

}

func (s *Server) Start() {
	panic("implement me")
}

func (s *Server) Stop() {
	panic("implement me")
}

func (s *Server) Serve() {
	panic("implement me")
}

func NewServer() *Server {
	return &Server{}
}
