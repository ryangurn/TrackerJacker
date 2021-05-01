package cross

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"os"
)


func FileExists(path string) (retBool bool, retData string) {
	retBool = false
	retData = ""
	if data, err := os.Stat(path); err == nil {
		retBool = true
		// marshal data
		if out, err := json.Marshal(data); err == nil {
			return retBool, string(out)
		}
	}
	return
}

func FileHash(path string) (string, error) {
	var h string

	file, err := os.Open(path)
	if err != nil {
		return h, err
	}

	defer file.Close()

	hash := md5.New()

	if _, err := io.Copy(hash, file); err != nil {
		return h, err
	}

	hashInBytes := hash.Sum(nil)[:16]
	h = hex.EncodeToString(hashInBytes)

	return h, nil
}