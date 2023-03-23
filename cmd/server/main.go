package main

import (
	"crypto/tls"
	"flag"
	"log"
	"net"

	pb "github.com/brabantcourt/grpc-google-cloud-profiler/google/devtools/cloudprofiler/v2"

	"tailscale.com/client/tailscale"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

var (
	endpoint = flag.String("endpoint", "0.0.0.0:50051", "gRPC service endpoint")
)

func main() {
	flag.Parse()

	c := &tls.Config{
		GetCertificate: tailscale.GetCertificate,
	}

	opts := grpc.Creds(credentials.NewTLS(c))
	s := grpc.NewServer(opts)
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
