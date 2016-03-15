package main

import (
	// gRPC
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "helloworld"

	// Command Runner
	"github.com/codeskyblue/go-sh"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	out, err := sh.Command("echo", "hello\tworld").Command("cut", "-f2").Output()
	if err != nil {
		log.Fatal(err, out)
	}
	return &pb.HelloReply{Message: "Hello " + in.Name + "" + string(out)}, nil
}
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}

// transmission-daemon
// transmission-remote --add "magnet/torrent_file"
// amulecmd
// amuled "ED2K_LINK"
