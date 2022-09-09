// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package model

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type FileHints struct {
	_tab flatbuffers.Table
}

func GetRootAsFileHints(buf []byte, offset flatbuffers.UOffsetT) *FileHints {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &FileHints{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsFileHints(buf []byte, offset flatbuffers.UOffsetT) *FileHints {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &FileHints{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *FileHints) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *FileHints) Table() flatbuffers.Table {
	return rcv._tab
}

// / IsBinary can be set to hint that the contents are binary data.
func (rcv *FileHints) IsBinary() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

// / IsBinary can be set to hint that the contents are binary data.
func (rcv *FileHints) MutateIsBinary(n bool) bool {
	return rcv._tab.MutateBoolSlot(4, n)
}

// / FileName hints at the name of the file that should be written. It's
// / truncated to 255 bytes if longer. It may be empty to suggest that the
// / file should not be written to disk. It may be equal to "_CONSOLE" to
// / suggest the data should not be written to disk.
func (rcv *FileHints) FileName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

// / FileName hints at the name of the file that should be written. It's
// / truncated to 255 bytes if longer. It may be empty to suggest that the
// / file should not be written to disk. It may be equal to "_CONSOLE" to
// / suggest the data should not be written to disk.
// / ModTime format allowed: RFC3339, contains the modification time of the file, or the zero time if not applicable.
func (rcv *FileHints) ModTime() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

// / ModTime format allowed: RFC3339, contains the modification time of the file, or the zero time if not applicable.
func FileHintsStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func FileHintsAddIsBinary(builder *flatbuffers.Builder, isBinary bool) {
	builder.PrependBoolSlot(0, isBinary, false)
}
func FileHintsAddFileName(builder *flatbuffers.Builder, fileName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(fileName), 0)
}
func FileHintsAddModTime(builder *flatbuffers.Builder, modTime flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(modTime), 0)
}
func FileHintsEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
