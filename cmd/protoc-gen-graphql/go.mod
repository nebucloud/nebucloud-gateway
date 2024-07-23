module github.com/nebucloud/nebucloud-gateway/cmd/protoc-gen-graphql

go 1.22.0

replace github.com/nebucloud/nebucloud-gateway => ../..

require (
	github.com/nebucloud/nebucloud-gateway v0.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.34.2
)

require github.com/iancoleman/strcase v0.3.0 // indirect
