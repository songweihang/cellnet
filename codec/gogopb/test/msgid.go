// Generated by github.com/songweihang/cellnet/protoc-gen-msg
// DO NOT EDIT!
// Source: pb.proto

package test

import (
	"github.com/songweihang/cellnet"
	"reflect"
	_ "github.com/songweihang/cellnet/codec/gogopb"
	"github.com/songweihang/cellnet/codec"
)

func init() {

	// pb.proto
	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("gogopb"),
		Type:  reflect.TypeOf((*ContentACK)(nil)).Elem(),
		ID:    60952,
	})
}