# Test implementation of Cloud Profiler gRPC server

Using Google's provided Cloud Profiler agent, the only methods that should be called are:

+ [`CreateProfile`]()
+ [`UpdateProfile`]()

The other method was implemented to test whether it was being called by the agent:

+ [`CreateOfflineProfile`]()

The code also tests use of Google's [`google.rpc.details`] to responses.

## References

+ [profiler.proto](https://github.com/googleapis/googleapis/blob/master/google/devtools/cloudprofiler/v2/profiler.proto)
+ [profiler.pb.go](https://github.com/googleapis/go-genproto/blob/main/googleapis/devtools/cloudprofiler/v2/profiler.pb.go)


## `protoc`

The `profiler.proto` Protobuf Go source ([`profile.pb.go`](https://github.com/googleapis/go-genproto/blob/main/googleapis/devtools/cloudprofiler/v2/profiler.pb.go)) exists but the gRPC service (`profiler_grpc.pb.go`) does not appear to exist.

Created both locally:

```bash
GOOGLEAPIS="/home/dazwilkin/Projects/googleapis"
PROJECT="/home/dazwilkin/Projects/grpc-google-cloud-profiler"

protoc \
--proto_path=${GOOGLEAPIS} \
--go_out=${PROJECT} \
--go_opt=paths=source_relative \
--go-grpc_out=${PROJECT} \
--go-grpc_opt=paths=source_relative \
${GOOGLEAPIS}/google/devtools/cloudprofiler/v2/profiler.proto
```

## Tailscale

Reusing an existing `hades-canyon.orca-allosaurus.ts.net.[crt|key]` pair

Because:

```bash
openssl x509 -enddate -noutout -in hades-canyon.orca-allosaurus.ts.net.crt
notAfter=May 18 17:03:29 2023 GMT
```

## `github-webhook`

Configured:

```golang
// Google Cloud Profiler
cfg := profiler.Config{
	Service:        subsystem,
	ServiceVersion: version,
	DebugLogging:   true,
	APIAddr:        "hades-canyon.orca-allosaurus.ts.net:50051",
}
```

## gRPC

```bash
ENDPOINT="hades-canyon.orca-allosaurus.ts.net:50051"

grpcurl ${ENDPOINT} list

google.devtools.cloudprofiler.v2.ProfilerService
grpc.health.v1.Health
grpc.reflection.v1alpha.ServerReflection

grpcurl ${ENDPOINT} grpc.health.v1.Health/Check

{
  "status": "SERVING"
}

grpcurl ${ENDPOINT} \
list google.devtools.cloudprofiler.v2.ProfilerService

google.devtools.cloudprofiler.v2.ProfilerService.CreateOfflineProfile
google.devtools.cloudprofiler.v2.ProfilerService.CreateProfile
google.devtools.cloudprofiler.v2.ProfilerService.UpdateProfile
```

For the last case, see [methods for adding google.rpc.details to gRPC status](https://github.com/grpc/grpc-go/issues/1233)

`profiler.proto` includes "The backoff duration is returned in `google.rpc.RetryInfo` extension on the response status" but `google.rpc.RetryInfo` was unfamiliar to me.

[`error_details.proto`](https://github.com/googleapis/googleapis/blob/master/google/rpc/error_details.proto) is package `google.rpc` and includes `RetryInfo` (among other e.g. `QuotaFailure`).

So, the code that effects this is:

```golang
retry := &errdetails.RetryInfo{
	RetryDelay: durationpb.New(5 * time.Minute),
}
x, _ := status.New(codes.Unavailable, "").WithDetails(retry)
...

return nil, x.Err()
```

And, when called:

```bash
grpcurl ${ENDPOINT} \ 
google.devtools.cloudprofiler.v2.ProfilerService/CreateProfile

ERROR:
  Code: Unavailable
  Message: msg
  Details:
  1)	{"@type":"type.googleapis.com/google.rpc.RetryInfo","retryDelay":"300s"}
```

Added (!) `QuotaFailure`:

```golang
details := []protoiface.MessageV1{
	&errdetails.RetryInfo{
		RetryDelay: durationpb.New(5 * time.Minute),
	},
	&errdetails.QuotaFailure{
		Violations: []*errdetails.QuotaFailure_Violation{
			{
				Subject:     "subject",
				Description: "description",
			},
		},
	},
}
```
And this returns:

```bash
grpcurl ${ENDPOINT} \
google.devtools.cloudprofiler.v2.ProfilerService/CreateProfile

ERROR:
  Code: Unavailable
  Message: msg
  Details:
  1)	{"@type":"type.googleapis.com/google.rpc.RetryInfo","retryDelay":"300s"}
  2)	{"@type":"type.googleapis.com/google.rpc.QuotaFailure","violations":[{"subject":"subject","description":"description"}]}
```

So, it's become lost in the implementation, but you can see that the `Details` are [`Any`](https://protobuf.dev/reference/protobuf/google.protobuf/#any) type and the format used for the `@type` is `type.googleapis.com` and then the fully-qualified (i.e. with package) message name.

## `cloudprofiler.googleapis.com:443`

```JSON
{
    "parent":"projects/ackal-230322",
    "deployment":{
        "projectId":"ackal-230322",
        "target":"github",
        "labels":{
            "language":"go",
            "version":"v0.0.1"
        }
    },
    "profileType":[
        "CPU",
        "HEAP",
        "THREADS",
        "HEAP_ALLOC"
    ]
}
```

## Invoke w/ `gRPCurl`

```bash
ENDPOINT="cloudprofiler.googleapis.com:443"
ROOT="/home/dazwilkin/Projects/googleapis"
PACKAGE="google/devtools/cloudprofiler/v2"

METHOD="google.devtools.cloudprofiler.v2.ProfilerService/CreateProfile"

TOKEN=$(gcloud auth print-access-token)

grpcurl \
--import-path=${ROOT} --proto=${ROOT}/${PACKAGE}/profiler.proto \
-H "Authorization: Bearer ${TOKEN}" \
-d '{
    "parent":"projects/ackal-230322",
    "deployment":{
        "projectId":"ackal-230322",
        "target":"github",
        "labels":{
            "language":"go",
            "version":"v0.0.1"
        }
    },
    "profileType":[
        "CPU",
        "HEAP",
        "THREADS",
        "HEAP_ALLOC"
    ]
}' \
${ENDPOINT} ${METHOD}
```
Yields:
```JSON
{
  "name": "projects/ackal-230322/profiles/5738e46935a57059",
  "profileType": "HEAP",
  "deployment": {
    "projectId": "ackal-230322",
    "target": "github",
    "labels": {
      "language": "go",
      "version": "v0.0.1"
    }
  },
  "duration": "10s"
}
```