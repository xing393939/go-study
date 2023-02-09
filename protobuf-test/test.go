package main

import (
	"encoding/hex"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func bytesToDex(bytes []byte) string {
	result := ""
	for _, b := range bytes {
		byteData := []byte{b}
		hexStringData := hex.EncodeToString(byteData)
		result += hexStringData + " "
	}
	return result
}

// 当前目录下执行：protoc --go_out=../.. test.proto
func main() {
	fn := int64(1337)
	ss := &Student{
		UserName:       "Martin",
		FavoriteNumber: &fn,
		Interests:      []string{"daydreaming", "hacking"},
	}
	buffer, _ := proto.Marshal(ss)
	bufferDex := bytesToDex(buffer)
	fmt.Println("序列化之后的信息为：", bufferDex)

	data := &Student{}
	_ = proto.Unmarshal(buffer, data)
	fmt.Println("反序列化之后的信息为：", data)
}
