package web

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const tmpDir = "./templates"
const tmpZip = "./tmp.zip"
const tmpMd5 = "d82e09dae684a494ccb9b969ab013e45"
const tmpUrl = "https://github.com/ryangurn/fakebank/raw/master/docs/trackerjacker.zip"

func Setup() (retBool bool) {
	retBool = false
	fmt.Println("--------------")
	fmt.Println("--------------")
	fmt.Println("--------------")
	err := download(tmpZip, tmpUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println("downloaded web templates")
	fmt.Println("--------------")
	fmt.Println("checking hash for template")
	hash, err := hash("tmp.zip")
	if err != nil {
		fmt.Println("Unable to generate hash")
	}
	fmt.Println("("+tmpMd5+") == ["+hash+"]")
	fmt.Println("--------------")
	zipErr := unzip(tmpZip, "./")
	fmt.Println("unzipping web templates")
	fmt.Println("--------------")
	fmt.Println("--------------")
	fmt.Println("--------------")

	if zipErr == nil && tmpMd5 == hash {
		retBool = true
	}

	os.Remove(tmpZip)
	os.RemoveAll("./__MACOSX")

	return
}

func Clean() bool {
	os.RemoveAll(tmpDir)

	return true
}

// ref: https://golangcode.com/download-a-file-from-a-url/
func download(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// ref: https://stackoverflow.com/questions/20357223/easy-way-to-unzip-file-with-golang
func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}