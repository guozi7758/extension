package lib

import (
	// "fmt"
	"io/ioutil"
	"strings"
	"net/http"

)

func XFormPost( str_url,had, para string )( []byte, error ){
	client := &http.Client{}
	req,err := http.NewRequest("POST",str_url,strings.NewReader( para ))
	if err != nil {
		return make([]byte, 0), err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("apikey",had)
	resp, err_a := client.Do(req)
	if err_a != nil {
		return make([]byte, 0), err_a
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

