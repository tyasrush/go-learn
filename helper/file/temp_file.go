package file

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func CreateTempFile(b []byte) (data string, err error) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "walletpass-")
	if err != nil {
		return data, err
	}
	// remove after we file uploaded
	defer os.Remove(tmpFile.Name())

	if _, err = tmpFile.Write(b); err != nil {
		return data, err
	}

	// get location file from temp
	data = tmpFile.Name()

	// close tmpfile process
	if err = tmpFile.Close(); err != nil {
		return data, nil
	}
	return data, nil
}

// return 4 items
// file for closing if the file are done
// path file to remove if we done with those file
// body response from http, close if we done with http get
// err if any error happened
func DownloadFile(nameFile string, link string) (f *os.File, path string, body io.ReadCloser, err error) {
	// download logo file
	f, err = ioutil.TempFile(os.TempDir(), nameFile)
	if err != nil {
		return f, path, body, err
	}

	resp, err := http.Get(link)
	if err != nil {
		return f, path, body, err
	}

	return f, f.Name(), resp.Body, err
}
