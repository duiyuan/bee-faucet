package thirdweb

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"time"
)

const Endpoint = "https://thirdweb.com/api/testnet-faucet/claim"

type TestnetFaucet struct {
	ChainID   int32  `json:"chainId"`
	ToAddress string `json:"toAddress"`
}

func (t *TestnetFaucet) PostWithCookie(data []byte, cookies []*http.Cookie) (*http.Response, error) {
	buf := bytes.NewBuffer(data)

	req, err := http.NewRequest("POST", Endpoint, buf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}

	return client.Do(req)
}
