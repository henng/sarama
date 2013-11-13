package sarama

import (
	"bytes"
	"testing"
)

func TestFetchRequestSerialization(t *testing.T) {

	exp := new(FetchRequestExpectation).
		AddMessage("my_topic", 0, nil, ByteEncoder([]byte{0x00, 0xEE}), 3)

	expected := []byte{
		0x00, 0x00, 0x00, 0x01, // number of topics
		0x00, 0x08, 'm', 'y', '_', 't', 'o', 'p', 'i', 'c', // topic name
		0x00, 0x00, 0x00, 0x01, // number of blocks for this topic
		0x00, 0x00, 0x00, 0x00, // partition id
		0x00, 0x00, // error
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // high water mark offset
		// messageSet
		0x00, 0x00, 0x00, 0x1C, // messageset size
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, // offset
		// message
		0x00, 0x00, 0x00, 0x10, // length of message (?)
		0x23, 0x96, 0x4a, 0xf7, // CRC32
		0x00,                   // format
		0x00,                   // attribute (compression)
		0xFF, 0xFF, 0xFF, 0xFF, // key (nil)
		0x00, 0x00, 0x00, 0x02, 0x00, 0xEE, // value
	}

	actual := exp.ResponseBytes()
	if bytes.Compare(actual, expected) != 0 {
		t.Error("\nExpected\n", expected, "\nbut got\n", actual)
	}
}
