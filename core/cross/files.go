package cross

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)


func FileExists(path string) (retBool bool) {
	retBool = false
	if _, err := os.Stat(path); err == nil {
		retBool = true
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