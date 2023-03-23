package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func TeamsServer(messageTitle string, messageText string) bool {
	body := []byte(`{
	  "@context": "https://schema.org/extensions",
	  "@type": "MessageCard",
	  "themeColor": "0072C6",
	  "title": "` + messageTitle + `",
	  "text": "` + messageText + `",
	}`)

	req, _ := http.NewRequest("POST", ApiUrl, bytes.NewBuffer(body))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		defer res.Body.Close()
		bodyBytes, err := io.ReadAll(res.Body)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Response: %v\n", string(bodyBytes))
		return true
	}
}
