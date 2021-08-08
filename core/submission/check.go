package submission

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetPayload() ([]byte, error) {
	// get image & auth
	auth := os.Getenv("AUTH_TOKEN")
	image := os.Getenv("IMAGE")
	baseURL := os.Getenv("SERVER")

	// create the new request
	req, err := http.NewRequest("GET", baseURL + "/api/check/image/" + image, nil)
	if err != nil {
		fmt.Println(err)
	}

	// add bearer
	req.Header.Add("Authorization", "Bearer " + auth)

	// send request using client
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close() // close the body

	// read output
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	return body, nil
}