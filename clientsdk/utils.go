package clientsdk

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func B64Encode(in string) string {
	out := base64.StdEncoding.EncodeToString([]byte(in))
	return out
}

func GET(url string, authToken string, headers map[string]string) ([]byte, error) {
	return RPC("GET", url, authToken, headers, nil)
}

func POST(url string, authToken string, headers map[string]string, content interface{}) ([]byte, error) {
	return RPC("POST", url, authToken, headers, content)
}

func RPC(method string, url string, authToken string, headers map[string]string, content interface{}) ([]byte, error) {
	var postBody io.Reader = nil
	var bodyBytes []byte
	if content != nil {
		bodyBytes, _ = json.Marshal(content)
		postBody = bytes.NewBuffer(bodyBytes)
	}

	req, err := http.NewRequest(method, url, postBody)
	if err != nil {
		log.Println("Error crating http request. ", err)
		return nil, err
	}
	//log.Printf("%s %s\n", method, url)

	req.Header.Set("Cache-Control", "no-cache")
	if content != nil {
		req.Header.Set("Content-Type", "application/json")
		//log.Printf("%s: %s\n", "Content-Type", "application/json")
	}

	if authToken != "" {
		req.Header.Set("Authorization", authToken)
		//log.Printf("%s: %s\n", "Authorization", authToken)
	}

	for k, v := range headers {
		req.Header.Set(k, v)
		//log.Printf("%s: %s\n", k, v)
	}

	if bodyBytes != nil {
		//log.Printf("%s\n", string(bodyBytes))
	}

	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error reading response. ", err)
		return nil, err
	}
	defer resp.Body.Close()
	// log.Printf("==== Response ====\n")
	// for k, v := range resp.Header {
	// 	log.Printf("%s: %v\n", k, v)
	// }

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//log.Fatal("Error reading body. ", err)
		return nil, err
	}

	//log.Printf("%s\n", string(body))
	return body, nil
}
