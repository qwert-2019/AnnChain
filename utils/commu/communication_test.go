package commu

import (
	"testing"
	"time"

	"github.com/dappledger/AnnChain/eth/common"
)

var hash string

func TestSendPayload(t *testing.T) {
	byt, err := SendPayload("http://127.0.0.1:8000", []string{}, []byte("01234567891111"))
	if err != nil {
		t.Log(err)
		return
	}
	hash = common.Bytes2Hex(byt)
	t.Log(hash)
	time.Sleep(time.Second)
}

func TestGetPayload(t *testing.T) {
	payload, err := GetPayload("http://127.0.0.1:8000", hash)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(string(payload))
}
