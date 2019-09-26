package commu

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/dappledger/AnnChain/eth/common"
	"github.com/dappledger/AnnChain/gemmill/go-crypto"
)

var hash string

type ST struct {
	H string
	D []byte
}

var sts []*ST

func TestSendPayload(t *testing.T) {
	var byt25519 [crypto.PrivKeyLenEd25519]byte
	copy(byt25519[:], common.Hex2Bytes("4A3597AE16309060A7634704989ECE522E058F5400C5274B0995F66F1C88B6B119FFA7BAA1D48D0B9357AF3C93620D2DC9CA4E82C2957A4AB4A3D057B6743CE0"))
	for i := 0; i < 10000; i++ {
		data := []byte("01234567sdfgsdfgrrr911" + fmt.Sprintf("%v", i))
		byt, err := SendPayloadWithSign("http://127.0.0.1:8000", []string{"971D0EB6F0FECA0B7365E621FD9EC5E6D281604DBDD82A3A85931F62B19AE7F9"}, data, crypto.PrivKeyEd25519(byt25519))
		if err != nil {
			t.Log(err)
			return
		}
		sts = append(sts, &ST{H: common.Bytes2Hex(byt), D: data})
	}
	time.Sleep(time.Second)
}

func TestGetPayload(t *testing.T) {
	for _, st := range sts {
		payload, err := GetPayload("http://127.0.0.1:8001", st.H)
		if err != nil {
			t.Log(err)
			continue
		}
		if 0 != bytes.Compare(payload, st.D) {
			t.Log(st.H, string(payload), string(st.D))
		}
	}
}

//func TestEquel(t *testing.T) {
//	for i := 0; i < 10000; i++ {
//		data := []byte("01234567sdfgsdfgrrr911" + fmt.Sprintf("%v", i))
//		payload, err := GetPayload("http://127.0.0.1:8001", common.Bytes2Hex(Hash(data)))
//		if err != nil {
//			t.Log(err)
//			continue
//		}
//		if 0 != bytes.Compare(payload, data) {
//			t.Log(string(payload), string(data))
//		}
//	}
//}
