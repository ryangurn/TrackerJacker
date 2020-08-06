package web

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func hash(path string) (string, error) {
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