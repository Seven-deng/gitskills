package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	//pb "google.golang.org/grpc/examples/helloworld/helloworld"
	pa "ai_py2go_grpc/src/go_dev/getscore/mahjong"
	pb "ai_py2go_grpc/src/go_dev/getscore/getscore"

)

const (
	port = ":50052"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedHuScoreServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) CheckHuType(ctx context.Context, in *pb.CardInfo) (*pb.CheckTypeScore, error) {
	var checker =  &pa.BaseRuleChecker{}
	var typeMapAllint32= in.GetTypeMapAll()
	typeMapAllint := make(map[int64]int)
	for i, i2 := range typeMapAllint32 {
		typeMapAllint[i]=int(i2)
	}

	message:= checker.CheckHuType(typeMapAllint,
		int(in.GetTilesNum()),
		in.GetDiscards(),
		in.GetHandCards(),
		in.GetSurplusCards(),
		in.GetOtherHands(),
		in.GetChowCards(),
		in.GetPongCards(),
		in.GetMingKong(),
		in.GetAnKong(),
		in.GetIsBanker(),
		in.GetIsZimo(),
		in.GetRobKong(),
		in.GetLastOptIsKong(),
		in.GetLastCardCpk(),
		in.GetLastCard(),
		in.GetTableFeng(),
		in.GetMenFeng())

	CheckTypeScoreint32:= make(map[int64]int32)
	for i, i2 := range message {
		CheckTypeScoreint32[i]=int32(i2)
	}

	return &pb.CheckTypeScore{Message: CheckTypeScoreint32},nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHuScoreServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
