package distributesort

import (
	"bytes"
	"encoding/binary"
)

func IntToBytes(n int) []byte {
	data := int64(n)
	byteBuf := bytes.NewBuffer([]byte{})
	binary.Write(byteBuf, binary.BigEndian, data)
	return byteBuf.Bytes()
}

func BytesToInt(bts []byte) int {
	byteBuf := bytes.NewBuffer(bts)
	var data int64
	binary.Read(byteBuf, binary.BigEndian, &data)
	return int(data)
}
