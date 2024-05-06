// Package desc contains "rich descriptors" for protocol buffers. The built-in
// descriptor types are simple protobuf messages, each one representing a
// different kind of element in the AST of a .proto source file.
//
// Because of this inherent "tree" quality, these build-in descriptors cannot
// refer to their enclosing file descriptor. Nor can a field descriptor refer to
// a message or enum descriptor that represents the field's type (for enum and
// nested message fields). All such links must instead be stringly typed. This
// limitation makes them much harder to use for doing interesting things with
// reflection.
//
// Without this package, resolving references to types is particularly complex.
// For example, resolving a field's type, the message type an extension extends,
// or the request and response types of an RPC method all require searching
// through symbols defined not only in the file in which these elements are
// declared but also in its transitive closure of dependencies.
//
// "Rich descriptors" avoid the need to deal with the complexities described
// above. A rich descriptor has all type references resolved and provides
// methods to access other rich descriptors for all referenced elements. Each
// rich descriptor has a usefully broad API, but does not try to mimic the full
// interface of the underlying descriptor proto. Instead, every rich descriptor
// provides access to that underlying proto, for extracting descriptor
// properties that are not immediately accessible through rich descriptor's
// methods.
//
// Also see the grpcreflect, dynamic, and grpcdynamic packages in this same
// repo to see just how useful rich descriptors really are.
//
// # Loading Descriptors
//
// Rich descriptors can be accessed in similar ways as their "poor" cousins
// (descriptor protos). Instead of using proto.FileDescriptor, use
// desc.LoadFileDescriptor. Message descriptors and extension field descriptors
// can also be easily accessed using desc.LoadMessageDescriptor and
// desc.LoadFieldDescriptorForExtension, respectively.
//
// If you are using the protoc-gen-gosrcinfo plugin (also in this repo), then
// the descriptors returned from these Load* functions will include source code
// information, and thus include comments for elements.
//
// # Creating Descriptors
//
// It is also possible create rich descriptors for proto messages that a given
// Go program doesn't even know about. For example, they could be loaded from a
// FileDescriptorSet file (which can be generated by protoc) or loaded from a
// server. This enables interesting things like dynamic clients: where a Go
// program can be an RPC client of a service it wasn't compiled to know about.
//
// You cannot create a message descriptor without also creating its enclosing
// file, because the enclosing file is what contains other relevant information
// like other symbols and dependencies/imports, which is how type references
// are resolved (such as when a field in a message has a type that is another
// message or enum).
//
// So the functions in this package for creating descriptors are all for
// creating *file* descriptors. See the various Create* functions for more
// information.
//
// Also see the desc/builder sub-package, for another API that makes it easier
// to synthesize descriptors programmatically.
//
// Deprecated: This module was created for use with the older "v1" Protobuf API
// in github.com/golang/protobuf. However, much of this module is no longer
// necessary as the newer "v2" API in google.golang.org/protobuf provides similar
// capabilities. Instead of using this github.com/davron112/protoreflect/desc package,
// see [google.golang.org/protobuf/reflect/protoreflect].
//
// [google.golang.org/protobuf/reflect/protoreflect]: https://pkg.go.dev/google.golang.org/protobuf/reflect/protoreflect
package desc
