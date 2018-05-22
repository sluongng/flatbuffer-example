package main

import (
	"fmt"

	flatbuffers "github.com/google/flatbuffers/go"
	user "github.com/sluongng/flatbuffer-example/user-service/fb/User"
)

func makeGetUserRequest(b *flatbuffers.Builder, id int32) []byte {

	b.Reset()

	user.GetUserReqStart(b)
	user.GetUserReqAddId(b, id)

	getUserRequest := user.GetUserReqEnd(b)
	b.Finish(getUserRequest)

	return b.Bytes[b.Head():]
}

func readGetUserRequest(buf []byte) (id int32) {
	getUserRequest := user.GetRootAsGetUserReq(buf, 0)

	id = getUserRequest.Id()

	return id
}

func main() {
	b := flatbuffers.NewBuilder(0)

	buf := makeGetUserRequest(b, 1)

	id := readGetUserRequest(buf)

	fmt.Printf("The id is %d", id)
}
