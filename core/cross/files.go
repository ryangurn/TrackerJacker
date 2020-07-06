package cross

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
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

func FileParse(args []string, result interface{}) (retBool bool) {
	retBool = false

	if len(args) != 2 {
		return
	}

	if args[0] == "exist" {
		if FileExists(args[1]) == result {
			retBool = true
		}
	} else if args[0] == "hash" {
		res, err := FileHash(args[1])
		if err != nil {
			return
		}

		if res == result {
			retBool = true
		}
	} else {
		fmt.Printf("Unrecognized Command: %s\n", args[0])
		return
	}

	return
}