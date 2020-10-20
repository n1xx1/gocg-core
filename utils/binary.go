package utils

import (
	"bytes"
	"encoding/binary"
)

func ReadUint8(b *bytes.Buffer) (v uint8) {
	_ = binary.Read(b, binary.LittleEndian, &v)
	return
}
func ReadUint16(b *bytes.Buffer) (v uint16) {
	_ = binary.Read(b, binary.LittleEndian, &v)
	return
}
func ReadUint32(b *bytes.Buffer) (v uint32) {
	_ = binary.Read(b, binary.LittleEndian, &v)
	return
}
func ReadUint64(b *bytes.Buffer) (v uint64) {
	_ = binary.Read(b, binary.LittleEndian, &v)
	return
}
func ReadInt8(b *bytes.Buffer) (v int8) {
	_ = binary.Read(b, binary.LittleEndian, &v)
	return
}
func ReadInt16(b *bytes.Buffer) (v int16) {
	_ = binary.Read(b, binary.LittleEndian, &v)
	return
}
func ReadInt32(b *bytes.Buffer) (v int32) {
	_ = binary.Read(b, binary.LittleEndian, &v)
	return
}
func ReadInt64(b *bytes.Buffer) (v int64) {
	_ = binary.Read(b, binary.LittleEndian, &v)
	return
}

func WriteUint8(b *bytes.Buffer, v uint8) {
	_ = binary.Write(b, binary.LittleEndian, v)
}
func WriteUint16(b *bytes.Buffer, v uint16) {
	_ = binary.Write(b, binary.LittleEndian, v)
}
func WriteUint32(b *bytes.Buffer, v uint32) {
	_ = binary.Write(b, binary.LittleEndian, v)
}
func WriteUint64(b *bytes.Buffer, v uint64) {
	_ = binary.Write(b, binary.LittleEndian, v)
}
func WriteInt8(b *bytes.Buffer, v int8) {
	_ = binary.Write(b, binary.LittleEndian, v)
}
func WriteInt16(b *bytes.Buffer, v int16) {
	_ = binary.Write(b, binary.LittleEndian, v)
}
func WriteInt32(b *bytes.Buffer, v int32) {
	_ = binary.Write(b, binary.LittleEndian, v)
}
func WriteInt64(b *bytes.Buffer, v int64) {
	_ = binary.Write(b, binary.LittleEndian, v)
}
