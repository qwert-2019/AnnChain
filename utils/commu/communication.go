package commu

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/dappledger/AnnChain/eth/rlp"
	"github.com/dappledger/AnnChain/gemmill/go-crypto"
	"golang.org/x/crypto/sha3"
)

var DefaultHost string

func Hash(v interface{}) []byte {
	h := make([]byte, 32)
	hw := sha3.NewLegacyKeccak256()
	rlp.Encode(hw, v)
	hw.Sum(h[:0])
	return h
}

func SendPayloadWithSign(host string, pubkeys []string, data []byte, signer crypto.PrivKey) ([]byte, error) {
	if len(host) > 0 {
		DefaultHost = host
	}
	url := DefaultHost + "/v1/transaction/withsignature"
	payloadHash := Hash(data)
	params := struct {
		PubKeys []string `json:"public_keys"`
		Value   []byte   `json:"value"`
		Sign    []byte   `json:"sign"`
	}{}
	params.PubKeys = pubkeys
	params.Value = data
	params.Sign = signer.Sign(payloadHash).Bytes()
	bytParams, err := json.Marshal(params)
	if err != nil {
		return payloadHash, err
	}
	buff := bytes.NewBuffer(bytParams)
	req, err := http.NewRequest("PUT", url, buff)
	if err != nil {
		return payloadHash, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "close")

	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		return payloadHash, err
	}
	byt, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return payloadHash, err
	}
	result := &struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	if err = json.Unmarshal(byt, result); err != nil {
		return payloadHash, err
	}
	if !result.Success {
		return payloadHash, errors.New(result.Message)
	}
	return payloadHash, nil
}

func SendPayload(host string, pubkeys []string, data []byte) ([]byte, error) {
	if len(host) > 0 {
		DefaultHost = host
	}
	url := DefaultHost + "/v1/transaction"
	payloadHash := Hash(data)
	params := struct {
		PubKeys []string `json:"public_keys"`
		Value   []byte   `json:"value"`
	}{}
	params.PubKeys = pubkeys
	params.Value = data
	bytParams, err := json.Marshal(params)
	if err != nil {
		return payloadHash, err
	}
	buff := bytes.NewBuffer(bytParams)
	req, err := http.NewRequest("PUT", url, buff)
	if err != nil {
		return payloadHash, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Connection", "close")

	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		return payloadHash, err
	}
	byt, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return payloadHash, err
	}
	result := &struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}{}

	if err = json.Unmarshal(byt, result); err != nil {
		return payloadHash, err
	}
	if !result.Success {
		return payloadHash, errors.New(result.Message)
	}
	return payloadHash, nil
}

func GetPayload(host, key string) (payload []byte, err error) {
	if len(host) > 0 {
		DefaultHost = host
	}
	rsp, err := http.Get(DefaultHost + "/v1/transaction/" + key)
	if err != nil {
		return nil, err
	}
	byt, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	result := &struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Data    []byte `json:"data"`
	}{}
	if err = json.Unmarshal(byt, result); err != nil {
		return nil, err
	}
	if !result.Success {
		return nil, errors.New(result.Message)
	}
	return result.Data, nil
}
