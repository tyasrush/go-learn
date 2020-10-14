package file_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/tyasrush/go-learn/helper/file"
)

func TestZipGenerate(t *testing.T) {
	// files := []string{"example.csv", "data.csv"}
	output := "done.zip"

	f, path, body, err := file.DownloadFile("image-icon-*.png", "https://storage.googleapis.com/cabeen-dev/wallet/material/icon.png")
	if err != nil {
		t.FailNow()
		t.Logf("error - %v", err)
	}

	defer f.Close()
	defer os.Remove(path)
	defer body.Close()

	if err := file.GenerateZipFile([]string{"./sample_files/icon.png", path}, output); err != nil {
		panic(err)
	}
	fmt.Println("Zipped File:", output)
}
