package bcapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// https://www.blockchain.com/api/blockchain_api

const (
	getRawTxURL       = "https://blockchain.info/rawtx/%s?format=hex"
	getRawBlockURL    = "https://blockchain.info/rawblock/%s?format=hex"
	getBlockHeightURL = "https://blockchain.info/block-height/%d?format=json"
)

type Client interface {
	GetRawTx(hash string) (string, error)
	GetRawBlockByHash(hash string) (string, error)
	GetRawBlockByHeight(h uint32) (string, error)
}

type SimpleClient struct {
	debug bool
}

func NewSimpleClient() Client {
	return SimpleClient{}
}

func (c SimpleClient) GetRawTx(hash string) (string, error) {
	url := fmt.Sprintf(getRawTxURL, hash)
	if c.debug {
		fmt.Println("GET", url)
	}
	resp, err := get(url)
	return string(resp), err
}

func (c SimpleClient) GetRawBlockByHash(hash string) (string, error) {
	url := fmt.Sprintf(getRawBlockURL, hash)
	if c.debug {
		fmt.Println("GET", url)
	}
	resp, err := get(url)
	return string(resp), err
}

func (c SimpleClient) GetRawBlockByHeight(h uint32) (string, error) {
	hash, err := c.getBlockHashByHeight(h)
	if err != nil {
		return "", err
	}

	return c.GetRawBlockByHash(hash)
}

func (c SimpleClient) getBlockHashByHeight(h uint32) (string, error) {

	type Block struct {
		Hash string `json:"hash"`
	}
	type Result struct {
		Blocks []Block `json:"blocks"`
	}

	url := fmt.Sprintf(getBlockHeightURL, h)
	if c.debug {
		fmt.Println("GET", url)
	}
	resp, err := get(url)
	if err != nil {
		return "", err
	}

	result := Result{}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		if c.debug {
			fmt.Println("resp:", string(resp))
		}
		return "", err
	}

	if n := len(result.Blocks); n != 1 {
		return "", fmt.Errorf("blocks: %d", n)
	}
	return result.Blocks[0].Hash, nil
}

func get(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
