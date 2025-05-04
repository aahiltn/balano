// cmd/server/gen.go
package main

//go:generate go tool oapi-codegen -config=../../config/oapi-codegen.yaml -o=../../api/server.gen.go ../../openapi.yaml
