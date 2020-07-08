module ai_py2go_grpc

go 1.14

require (
	github.com/golang/mock v1.1.1
	github.com/golang/protobuf v1.4.2
	google.golang.org/grpc v1.30.0
	google.golang.org/grpc/examples v0.0.0-20200627230533-68098483a7af
	google.golang.org/protobuf v1.25.0
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.30.0
