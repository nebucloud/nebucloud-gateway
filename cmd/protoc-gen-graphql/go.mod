module github.com/nebucloud/nebucloud-gateway/cmd/protoc-gen-graphql

go 1.22

require (
	github.com/nebucloud/nebucloud-gateway v0.1.0
	google.golang.org/protobuf v1.34.2
)

require (
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/iancoleman/strcase v0.3.0 // indirect
)

replace github.com/nebucloud/nebucloud-gateway => ../..
