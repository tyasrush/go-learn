package file

import (
	"io/ioutil"
	"os"
)

func CreateTempFile() (string, error) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "walletpass-")
	if err != nil {
		return data, err
	}
	// remove after we file uploaded
	defer os.Remove(tmpFile.Name())

	if _, err = tmpFile.Write(z); err != nil {
		return data, err
	}

	// get location file from temp
	tempFileLocation := tmpFile.Name()

	// close tmpfile process
	if err = tmpFile.Close(); err != nil {
		return data, nil
	}
	return "", nil
}
