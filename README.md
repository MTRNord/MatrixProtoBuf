# Matrix Protobuf/gRPC Proxy

This Repo contains a Protobuf and gRPC Proxy named MatrixProtoBuf for the [Matrix Protocol](https://matrix.org) together with example clients for both of them.
They are designed to be run localy on a client or on Servers as they take a HS address on every request.

### Join the discussion

Join the discussion about this and help making decisions in [#MatrixProtoBuf:matrix.org](https://matrix.to/#/#MatrixProtoBuf:matrix.org)

## How to build this?

1. Install GB from [getgb.io](getgb.io)
2. Run `gb vendor restore`
3. Run `gb build`

## How and what to run?

To run only the Proxy which provides the access to the gRPC Server and HTTP Protobuf Server just run the MatrixProtoBuf* file in `bin`.
The Proxy currently listens on port `50051` fo gRPC and on `3000` for Protobuf over HTTP

## Where is the Protobuf definition?

the .proto file is in `src/github.com/Nordgedanken/MatrixProtoBuf/matrixProtos` named matrix.proto
