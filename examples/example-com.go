package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/govenue/cachecontrol"
)

func main() {
	req, _ := http.NewRequest("GET", "http://www.example.com/", nil)

	res, _ := http.DefaultClient.Do(req)
	_, _ = ioutil.ReadAll(res.Body)

	reasons, expires, _ := cachecontrol.CachableResponse(req, res, cachecontrol.Options{})

	fmt.Println("Reasons to not cache: ", reasons)
	fmt.Println("Expiration: ", expires.String())
}
