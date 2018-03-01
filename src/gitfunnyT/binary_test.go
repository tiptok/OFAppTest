package gitfunnyT

import (
	"log"
	"testing"

	"bytes"

	"github.com/funny/binary"
)

func TestBinary(t *testing.T) {
	/*通过binary读取*/
	data := []byte{0x01, 0x00, 0x00, 0x00}
	vLEInt32 := binary.GetUint32LE(data)
	vBEInt32 := binary.GetUint32BE(data)
	log.Println("Get Int32LE:", vLEInt32)
	log.Println("Get Int32BE:", vBEInt32)
	// 2018/03/01 14:57:48 Get Int32LE: 1
	// 2018/03/01 14:57:48 Get Int32BE: 16777216

	/*通过read读取*/
	bR := binary.Reader{}
	bR.R = bytes.NewBuffer([]byte{0x01, 0x00, 0x00, 0x00, 0x02, 0x00})
	// 2018/03/01 14:57:48 Read Int32LE: 1
	// 2018/03/01 14:57:48 Read Int16LE: 2

	bRInt32 := bR.ReadInt32LE()
	bRInt16 := bR.ReadInt16LE()
	log.Println("Read Int32LE:", bRInt32)
	log.Println("Read Int16LE:", bRInt16)

	//binary.BufioOptimizer
	//binary.Reader
	//binary.Writer

	//binary.Buffer
	t.Log("end")
}
