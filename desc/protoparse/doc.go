// Package protoparse provides functionality for parsing *.proto source files
// into descriptors that can be used with other protoreflect packages, like
// dynamic messages and dynamic GRPC clients.
//
// This package links in other packages that include compiled descriptors for
// the various "google/protobuf/*.proto" files that are included with protoc.
// That way, like when invoking protoc, programs need not supply copies of these
// "builtin" files. Though if copies of the files are provided, they will be
// used instead of the builtin descriptors.
//
// Deprecated: This protoparse package is now just a thin veneer around a newer
// replacement parser/compiler: [github.com/bufbuild/protocompile]. Users are
// highly encouraged to directly use protocompile instead of this package.
//
// [github.com/bufbuild/protocompile]: https://pkg.go.dev/github.com/bufbuild/protocompile
package protoparse
