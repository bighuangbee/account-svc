package protocol

import "github.com/google/wire"

var ProviderSet = wire.NewSet(wire.Struct(new(PbServer), "*"), NewHTTPServer, NewGRPCServer)
