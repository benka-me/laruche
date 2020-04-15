#Generator

The generator create entry point files for what we need:
- Files structure
- Protobuf files
- Main application for **service** or **gateway** or **client**
- Empty clients file provider untill we add new service dependencies

It will re-generate the client file each time we add a service dependency.

## Example
### Service created as:
- Name: aaaa
- Author: benka-me
- Dev langage: golang

```
.
├── bee.yaml                         ** bee config file
├── go-pkg                           
│   ├── aaaa                         ** Package 
│   │   ├── aaaa.pb.go               ** Protobuf definitions generated for the server and consumers
│   │   └── rpc-aaaa.pb.go           ** gRPC 
│   └── http
│       └── rpc
│           ├── clients.go           ** Dependencies connections providers
│           ├── hello-world.go       ** Dummy example
│           └── server-grpc-2.0.go   ** Entry point service application
├── js-pkg
│   └── src
│       └── protobuf
│           ├── aaaa_pb.js           ** Javascript protobuf generated for consumers 
│           └── rpc-aaaa_pb.js
├── main.go                          ** Go main package
└── protobuf
    ├── aaaa.proto                   ** Protobuf definitions
    └── rpc-aaaa.proto
```

### Client created as:
```
Here is a service created as:
- Dev langage: golang
.
├── bee.yaml
├── go-pkg
│   └── http
│       └── rpc
│           └── clients.go           ** Dependencies connections providers
└── main.go                          ** Entrypoint for the app
```

