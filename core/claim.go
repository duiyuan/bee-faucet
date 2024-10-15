package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var UserCookies = []*http.Cookie{}
var endpoint = "http://localhost:8595/sendtx"

func ClaimToken[T any](chain string, toAddress string) (T, error) {
	var result T

	req := &TestnetFaucet{
		Chain:     chain,
		ToAddress: toAddress,
	}

	bytes, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return result, err
	}

	resp, err := req.PostWithCookie(endpoint, bytes, UserCookies)
	if err != nil {
		fmt.Println(err)
		return result, err
	}

	defer resp.Body.Close()

	bytes, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return result, err
	}

	if err := json.Unmarshal(bytes, &result); err != nil {
		return result, err
	}

	return result, nil
}
