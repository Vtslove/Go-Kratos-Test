//go:build !windows
// +build !windows

//go:generate sh -c "protoc --proto_path=. --proto_path=../../../../../third_party --go_out=paths=source_relative:. ./*.proto"

package conf
