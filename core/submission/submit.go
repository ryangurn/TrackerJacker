package submission

import (
	"TrackerJacker/core/enc"
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

// Send example: submission.Send(data, result, payload[i].ID)
func  Send(data string, result bool, check int) {
	// get method from env
	method := os.Getenv("SCORING_METHOD")
	baseURL := os.Getenv("SERVER")
	// change the uri based on scoring method
	var baseURI string
	switch method {
	case "0":
		baseURI = baseURL + "/api/score/compact"
		break
	case "1":
		baseURI = baseURL + "/api/score/hybrid"
		break
	case "2":
		baseURI = baseURL + "/api/score/verbose"
		break
	}

	// encode response to json
	jsonData, err := json.Marshal(map[string]string{
		"data": data,
		"result": strconv.FormatBool(result),
	})
	if err != nil {
		fmt.Println(err)
	}

	// calculate the fingerprint of the result
	fingerprint := enc.Hash(string(jsonData))

	// compose the form elements
	// reference: https://golangbyexample.com/http-mutipart-form-body-golang/
	requestBody := &bytes.Buffer{}
	writer := multipart.NewWriter(requestBody)
	writer.WriteField("check_id", strconv.FormatInt(int64(check), 10))
	writer.WriteField("fingerprint", fingerprint)
	// only send the correct information based on the method (0 = result, otherwise = json)
	switch method {
	case "0":
		writer.WriteField("result", strconv.FormatBool(result))
		break
	default:
		writer.WriteField("json", string(jsonData))
		break
	}

	// create the new request
	req, err := http.NewRequest("POST", baseURI, bytes.NewReader(requestBody.Bytes()))
	if err != nil {
		fmt.Println(err)
	}

	// get authorization from env
	auth := os.Getenv("AUTH_TOKEN")
	
	// add bearer
	req.Header.Add("Authorization", "Bearer " + auth)
	// set content type
	req.Header.Add("Content-Type", writer.FormDataContentType())

	// send request using client
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close() // close the body

	// read out the response from the server
	switch res.StatusCode {
	case 200:
		fmt.Println("Sent check information successfully.")
		break
	default:
		fmt.Println("Error sending check information. Please check internet connection or authorization token.")
		break
	}
}
