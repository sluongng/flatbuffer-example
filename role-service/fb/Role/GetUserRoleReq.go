// automatically generated by the FlatBuffers compiler, do not modify

package Role

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type GetUserRoleReq struct {
	_tab flatbuffers.Table
}

func GetRootAsGetUserRoleReq(buf []byte, offset flatbuffers.UOffsetT) *GetUserRoleReq {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &GetUserRoleReq{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *GetUserRoleReq) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *GetUserRoleReq) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *GetUserRoleReq) UserId() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *GetUserRoleReq) MutateUserId(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func GetUserRoleReqStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func GetUserRoleReqAddUserId(builder *flatbuffers.Builder, userId int32) {
	builder.PrependInt32Slot(0, userId, 0)
}
func GetUserRoleReqEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
