package core

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var UserCookies = []*http.Cookie{}
var endpoint = "http://localhost:8595/sendtx"

var str = "{\"data\":[{\"hash\":\"0xa7dd50bf3026c8d26f492faa105e99b7fceb6ac8041bce7d26f5a490fa783574\",\"amount\":\"1000000000000\",\"token\":\"ETH\",\"chain\":\"ethereum\"}],\"err\":0,\"message\":\"ok\"}"

func ClaimToken[T any](chain string, toAddress string) (T, error) {
	var result T

	// req := &TestnetFaucet{
	// 	Chain:     chain,
	// 	ToAddress: toAddress,
	// }

	// bytes, err := json.Marshal(req)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return result, err
	// }

	// resp, err := req.PostWithCookie(endpoint, bytes, UserCookies)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return result, err
	// }

	// defer resp.Body.Close()

	// bytes, err = io.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return result, err
	// }

	bytes := []byte(str)

	if err := json.Unmarshal(bytes, &result); err != nil {
		fmt.Println(err)
		return result, err
	}

	return result, nil
}
