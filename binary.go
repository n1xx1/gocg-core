package ocgcore

import (
	"bytes"
	"encoding/binary"
)

func readUint8(b *bytes.Buffer) (v uint8) {
	_ = binary.Read(b, binary.LittleEndian, &v)
	return
}
func readUint16(b *bytes.Buffer) (v uint16) {
	_ = binary.Read(b, binary.LittleEndian, &v)
	return
}
func readUint32(b *bytes.Buffer) (v uint32) {
	_ = binary.Read(b, binary.LittleEndian, &v)
	return
}
func readUint64(b *bytes.Buffer) (v uint64) {
	_ = binary.Read(b, binary.LittleEndian, &v)
	return
}
func readInt8(b *bytes.Buffer) (v int8) {
	_ = binary.Read(b, binary.LittleEndian, &v)
	return
}
func readInt16(b *bytes.Buffer) (v int16) {
	_ = binary.Read(b, binary.LittleEndian, &v)
	return
}
func readInt32(b *bytes.Buffer) (v int32) {
	_ = binary.Read(b, binary.LittleEndian, &v)
	return
}
func readInt64(b *bytes.Buffer) (v int64) {
	_ = binary.Read(b, binary.LittleEndian, &v)
	return
}

func writeUint8(b *bytes.Buffer, v uint8) {
	_ = binary.Write(b, binary.LittleEndian, v)
}
func writeUint16(b *bytes.Buffer, v uint16) {
	_ = binary.Write(b, binary.LittleEndian, v)
}
func writeUint32(b *bytes.Buffer, v uint32) {
	_ = binary.Write(b, binary.LittleEndian, v)
}
func writeUint64(b *bytes.Buffer, v uint64) {
	_ = binary.Write(b, binary.LittleEndian, v)
}
func writeInt8(b *bytes.Buffer, v int8) {
	_ = binary.Write(b, binary.LittleEndian, v)
}
func writeInt16(b *bytes.Buffer, v int16) {
	_ = binary.Write(b, binary.LittleEndian, v)
}
func writeInt32(b *bytes.Buffer, v int32) {
	_ = binary.Write(b, binary.LittleEndian, v)
}
func writeInt64(b *bytes.Buffer, v int64) {
	_ = binary.Write(b, binary.LittleEndian, v)
}
