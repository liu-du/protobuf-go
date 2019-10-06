proto: src/simple/simple.pb.go src/enum_example/enum_example.pb.go src/complex/complex.pb.go

src/simple/simple.pb.go:
	protoc -I src/simple/ --go_out=src/simple/ src/simple/simple.proto

src/enum_example/enum_example.pb.go:
	protoc -I src/enum_example/ --go_out=src/enum_example/ src/enum_example/enum_example.proto

src/complex/complex.pb.go:
	protoc -I src/complex/ --go_out=src/complex/ src/complex/complex.proto
