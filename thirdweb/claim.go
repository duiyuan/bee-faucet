package thirdweb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"githun.com/duiyuan/faucet/constants"
)

var UserToken = "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiIweDAxNjc1N2REZjJBYjZhOTk4YTQ3MjlBODBhMDkxMzA4ZDkwNTlFMTciLCJzdWIiOiIweDdGOTIwMzFGNjNlMDFBNkFEQTQ3NUY2RThBNjM3Q2U3NTJmOGQ3RDIiLCJhdWQiOiJ0aGlyZHdlYi5jb20iLCJleHAiOjE3Mjg5NTk5NDIsIm5iZiI6MTcyODcwMDEzNywiaWF0IjoxNzI4NzAwNzQyLCJqdGkiOiI4ODRiZTI1OTFlM2NmMjU1YjU3YjcxN2YxMjM4ZTI2MDQwZDdkZDNjMTMyNmI0MmFlODIyYzA3NTljM2FkMjJhIiwiY3R4Ijp7fX0.MHgzYzQyMTM5MzEyNTcwZGUxMDJjYjdhNjFhMTM3ZGI3NDRkZDA3ZmI2MTdkZThkMzViMDAxZDA0ZGViNGVkOGIzMzIxYTYxMTVhYzE2NGYzZDdhMTM4OWIxMTg2MDU0YzQ1YTk3MTk0Y2ZmM2ExZWY1ZmIwMzNlOWYyNzg3NTI5OTFi"

var UserCookies = []*http.Cookie{
	{Name: "tw_token_0x7F92031F63e01A6ADA475F6E8A637Ce752f8d7D2", Value: UserToken},
}

func ClaimToken(chainId int32, name string, toAddress string) error {
	req := &TestnetFaucet{
		ChainID:   chainId,
		ToAddress: toAddress,
	}

	bytes, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return err
	}

	resp, err := req.PostWithCookie(bytes, UserCookies)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer resp.Body.Close()

	bytes, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(bytes, &result); err != nil {
		fmt.Println(err)
		return err
	}

	if result["error"] != nil {
		format := fmt.Sprintf("  %%-%ds [error]: %%s", constants.MaxChainNameLen)
		e := fmt.Errorf(format, name, result["error"])
		fmt.Println(e)
		return e
	}
	format := fmt.Sprintf("  %%-%ds done, faucet %%s ETH!\n", constants.MaxChainNameLen)
	fmt.Printf(format, name, result["amount"])
	return nil
}
