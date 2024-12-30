package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	cloudprofiler "github.com/brabantcourt/grpc-google-cloud-profiler/google/devtools/cloudprofiler/v2"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/durationpb"
)

var _ cloudprofiler.ProfilerServiceServer = (*Server)(nil)

type Server struct {
	cloudprofiler.UnimplementedProfilerServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) CreateProfile(ctx context.Context, rqst *cloudprofiler.CreateProfileRequest) (*cloudprofiler.Profile, error) {
	log.Print("[CreateProfile] entered")
	log.Printf("%+v", rqst)
	b, err := protojson.Marshal(rqst)
	if err == nil {
		log.Printf("%s", b)
	}

	// Artificial delay
	// The document says:
	// > The server ensures that the new profiles are created at a constant rate per deployment,
	// > so the creation request may hang for some time until the next profile session is available.
	time.Sleep(30 * time.Second)

	// x := status.New(codes.Unavailable, "msg")

	// details := []protoiface.MessageV1{
	// 	&errdetails.RetryInfo{
	// 		RetryDelay: durationpb.New(5 * time.Minute),
	// 	},
	// 	&errdetails.QuotaFailure{
	// 		Violations: []*errdetails.QuotaFailure_Violation{
	// 			{
	// 				Subject:     "subject",
	// 				Description: "description",
	// 			},
	// 		},
	// 	},
	// }

	// y, err := x.WithDetails(details...)
	// if err != nil {
	// 	log.Print("unable to add retry details to status")
	// 	// Return previous status (x) because y failed
	// 	return nil, x.Err()
	// }

	profile, err := randomHex(16)
	if err != nil {
		profile = "0000000000000000"
	}
	name := fmt.Sprintf("%s/profiles/%s", rqst.Parent, profile)

	profileType := func() cloudprofiler.ProfileType {
		// 1<=r<=7
		r := rand.Intn(7) + 1
		return cloudprofiler.ProfileType(r)
	}()

	projectID := func(parent string) string {
		result := strings.Split(parent, "/")
		return result[1]
	}(rqst.Parent)

	log.Print("[CreateProfile] exited")
	// return nil, y.Err()
	return &cloudprofiler.Profile{
		Name:        name,
		ProfileType: profileType,
		Deployment: &cloudprofiler.Deployment{
			ProjectId: projectID,
			Target:    rqst.Deployment.Target,
			Labels:    rqst.Deployment.Labels,
		},
		Duration: durationpb.New(10 * time.Second),
	}, nil
}
func (s *Server) CreateOfflinProfile(ctx context.Context, rqst *cloudprofiler.CreateOfflineProfileRequest) (*cloudprofiler.Profile, error) {
	log.Print("[CreateOfflineProfile] entered")
	log.Printf("%+v", rqst)
	b, err := protojson.Marshal(rqst)
	if err == nil {
		log.Printf("%s", b)
	}
	log.Print("[CreateOfflineProfile] exited")
	return &cloudprofiler.Profile{}, nil
}
func (s *Server) UpdateProfile(ctx context.Context, rqst *cloudprofiler.UpdateProfileRequest) (*cloudprofiler.Profile, error) {
	log.Print("[UpdateProfile] entered")
	log.Printf("%+v", rqst)
	b, err := protojson.Marshal(rqst)
	if err == nil {
		log.Printf("%s", b)
	}
	log.Print("[UpdateProfile] exited")
	return &cloudprofiler.Profile{}, nil
}
