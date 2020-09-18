package html

import (
	"context"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestHTMLToPDF(t *testing.T) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "bp-*.html")
	if err != nil {
		t.FailNow()
		t.Errorf("error init temp file")
	}
	// remove after we file uploaded
	defer os.Remove(tmpFile.Name())

	if _, err = tmpFile.Write([]byte(BoardingPassHTML)); err != nil {
		t.FailNow()
		t.Errorf("error write file")
	}

	// get location file from temp
	tempFileLocation := tmpFile.Name()

	// close tmpfile process
	if err = tmpFile.Close(); err != nil {
		t.FailNow()
		t.Errorf("error close file")
	}

	t.Log("file location - ", tempFileLocation)

	file, err := CreateHTMLFile(context.Background(), tempFileLocation)
	if err != nil {
		t.FailNow()
		t.Errorf("error create pdf file")
	}

	// err = ioutil.WriteFile("./output.pdf", file, 2048)

	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		t.FailNow()
		t.Errorf("error out file")
	}
}
