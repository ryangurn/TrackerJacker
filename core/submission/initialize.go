package submission

import (
	"TrackerJacker/core/parsing"
	"encoding/json"
	"errors"
	"github.com/gen2brain/dlgs"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Initialize(baseURL string) error {
	authentication, _, err := dlgs.Password("Initialization", "Enter your API key")
	if err != nil {
		return err
	}

	// get image based on auth token
	req, err := http.NewRequest("POST", baseURL + "/api/user/auth", nil)
	if err != nil {
		return err
	}

	// add bearer
	req.Header.Add("Authorization", "Bearer " + authentication)

	// send request using client
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close() // close the body

	// read output
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	switch res.StatusCode {
	case 200:
		// unmarshal response
		output := parsing.InitialAuthorization{}
		json.Unmarshal(body, &output)

		id := strconv.FormatInt(int64(output.GetID()), 10)
		method := strconv.FormatInt(int64(output.GetMethod()), 10)

		env, _ := godotenv.Unmarshal("AUTH_TOKEN="+authentication+"\nBUGSNAG_KEY=1073d19819b48d408710876c02dd126b\nSERVER="+baseURL+"\nIMAGE="+id+"\nSCORING_METHOD="+method)
		godotenv.Write(env, "./.env")

		err := godotenv.Load()
		if err != nil {
			return err
		}

		dlgs.Info("Init Completed", "TrackerJacker has successfully completed initialization")
		break
	case 500:
		dlgs.Warning("Init Failed", "Unable to start, authentication token not tied to image yet.")
		return errors.New("Unable to start")
		break
	case 401:
		dlgs.Warning("Init Failed", "Authentication token not valid")
		return errors.New("Auth token invalid")
		break
	}

	return nil
}
