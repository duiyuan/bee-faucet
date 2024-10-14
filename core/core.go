package core

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

type TestnetFaucet struct {
	Chain     string `json:"chain"`
	ToAddress string `json:"recipient"`
}

func (t *TestnetFaucet) PostWithCookie(endpoint string, data []byte, cookies []*http.Cookie) (*http.Response, error) {
	buf := bytes.NewBuffer(data)

	req, err := http.NewRequest("POST", endpoint, buf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	client := &http.Client{
		// Timeout: time.Second * 30,
	}

	return client.Do(req)
}
