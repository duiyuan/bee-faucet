package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var UserCookies = []*http.Cookie{}
var endpoint = "http://localhost:8091/sendtx"

func ClaimToken(chain string, toAddress string) (map[string]interface{}, error) {
	var result map[string]interface{}

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
		fmt.Println(err)
		return result, err
	}

	return result, nil
}
