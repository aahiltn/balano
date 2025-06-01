// cmd/server/gen.go
package main

//go:generate go tool oapi-codegen -config=../../config/oapi-codegen.yaml -o=../../internal/service/server.gen.go ../../openapi.yaml
