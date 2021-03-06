package network

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"log"
	"testing"
	"time"

	"github.com/iost-official/Go-IOS-Protocol/common"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRequest_isValidNode(t *testing.T) {
	Convey("register", t, func() {
		bn, _ := NewBaseNetwork(&NetConfig{RegisterAddr: "127.0.0.1:30304", ListenAddr: "127.0.0.1", NodeTablePath: "iost_db_"})
		isValid := isValidNode(&Request{From: []byte("127.0.0.1")}, bn)
		So(isValid, ShouldBeTrue)
		isValid = isValidNode(&Request{From: []byte("192.168.1.34")}, bn)
		So(isValid, ShouldBeTrue)
		isValid = isValidNode(&Request{From: []byte("13.232.79.7")}, bn)
		So(isValid, ShouldBeTrue)

		NetMode = PublicMode
		isValid = isValidNode(&Request{From: []byte("127.0.0.1")}, bn)
		So(isValid, ShouldBeFalse)
		isValid = isValidNode(&Request{From: []byte("192.168.1.34")}, bn)
		So(isValid, ShouldBeFalse)
		isValid = isValidNode(&Request{From: []byte("13.232.79.7")}, bn)
		So(isValid, ShouldBeTrue)

	})
}

func TestRequest_Unpack(t *testing.T) {
	tim := time.Now().UnixNano()
	req := newRequest(Message, "0.0.0.0", common.Int64ToBytes(tim))

	Convey("test unpack packet splicing", t, func() {
		testData, err := req.Pack()
		So(err, ShouldEqual, nil)
		buf := new(bytes.Buffer)
		buf.Write(testData)
		buf.Write(testData)
		buf.Write(testData)

		readerCh := make(chan Request, 3)
		// scanner
		reader(buf, readerCh)
		i := 0
		for {
			select {
			case req := <-readerCh:
				if len(req.Body) > 0 {
					So(common.BytesToInt64(req.Body), ShouldEqual, tim)
					i++
				}
				if i == 3 {
					return
				}
			case <-time.After(1 * time.Second):
				So("timeout", ShouldEqual, "")
				break

			}
		}
	})
}

func reader(buf *bytes.Buffer, readerCh chan Request) {
	scanner := bufio.NewScanner(buf)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if !atEOF && data[0] == 'i' {
			if len(data) > 8 {
				length := int32(0)
				binary.Read(bytes.NewReader(data[4:8]), binary.BigEndian, &length)
				if int(length)+8 <= len(data) {
					return int(length) + 8, data[:int(length)+8], nil
				}
			}
		}
		return
	})
	for scanner.Scan() {
		scannedPack := new(Request)
		scannedPack.Unpack(bytes.NewReader(scanner.Bytes()))
		readerCh <- *scannedPack
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("invalid data pack")
	}
}
