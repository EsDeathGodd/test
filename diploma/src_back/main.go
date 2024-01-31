package studying

import (
    "log"
    "net"

    // import protobuf
    pbf "studying/test"

    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection" 
)

const (
    port = ":50051"
)



