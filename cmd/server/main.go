package main

import (
	"flag"
	"log"
	"net"

	pb "github.com/brabantcourt/grpc-google-cloud-profiler/google/devtools/cloudprofiler/v2"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

var (
	endpoint = flag.String("endpoint", "0.0.0.0:50051", "gRPC service endpoint")
)

func main() {
	flag.Parse()

	opts := []grpc.ServerOption{
		grpc.Creds(insecure.NewCredentials()),
	}
	s := grpc.NewServer(opts...)
	reflection.Register(s)

	server := NewServer()
	pb.RegisterProfilerServiceServer(s, server)

	healthcheck := health.NewServer()
	healthpb.RegisterHealthServer(s, healthcheck)

	listen, err := net.Listen("tcp", *endpoint)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(s.Serve(listen))
}
