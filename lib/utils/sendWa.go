package utils

import (
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func SendWa(phone, text string) (interface{}, error) {
	url := "https://fonnte.com/api/send_message.php"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("phone", phone)
	_ = writer.WriteField("type", "text")
	_ = writer.WriteField("text", text)
	_ = writer.WriteField("delay", "1")
	_ = writer.WriteField("delay_req", "1")
	_ = writer.WriteField("schedule", "0")
	err := writer.Close()
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Y1yt8pNrPLCB92o3wVNn")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return string(body), nil
}
