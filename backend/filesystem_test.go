package backend

import (
	"os"
	"testing"

	"github.com/spf13/viper"
)

// TODO: Add test for UsePath

func TestClearTemp(t *testing.T) {
	viper.Set("app.temp_storage", "./test/temp")

	err := ClearTemp()
	if err != nil {
		t.Error(err)
	}

	// Do cleanup
	t.Cleanup(func() {
		os.RemoveAll("./test_tmp/")
	})
}

func TestDownloadFile(t *testing.T) {
	fileData, err := DownloadFile("./test_tmp/download_test", "https://github.com/BitsOfAByte/proto/blob/main/README.md")

	if err != nil {
		t.Error(err)
	}

	if fileData.Name() != "download_test" {
		t.Error("File name is not correct")
	}

	// Do cleanup
	t.Cleanup(func() {
		os.RemoveAll("./test_tmp/")
	})
}

func TestExtractTar(t *testing.T) {
	// get .test_data/test.tar.gz from root of repo
	fileData, err := os.Open("../.test_data/test.tar.gz")
	if err != nil {
		t.Error(err)
	}

	// Extract the tarball to the temp directory.
	err = ExtractTar("./test_tmp/", fileData)
	if err != nil {
		t.Error(err)
	}

	// Check if the extracted files are correct.
	_, err = os.Stat("./test_tmp/file.txt")
	if err != nil {
		t.Error("File not extracted correctly")
	}

	// Do cleanup
	t.Cleanup(func() {
		os.RemoveAll("./test_tmp/")
	})
}

func TestMatchChecksum(t *testing.T) {

	match, err := MatchChecksum("../.test_data/file.txt", "../.test_data/file.sha512sum")

	if err != nil {
		t.Error(err)
	}

	if !match {
		t.Error("Checksum does not match")
	}

	// Do cleanup
	t.Cleanup(func() {
		os.RemoveAll("./test_tmp/")
	})
}

func TestDirSize(t *testing.T) {

	size, err := DirSize("../.test_data/")

	if err != nil {
		t.Error(err)
	}

	if size == 0 {
		t.Error("Size is zero")
	}

	// Do cleanup
	t.Cleanup(func() {
		os.RemoveAll("./test_tmp/")
	})
}

func TestHumanReadableSize(t *testing.T) {

	sizeKb, unitKb := HumanReadableSize(1024)

	if sizeKb != 1 && unitKb != "KB" {
		t.Error("Size is not correct (Kb test)")
	}

	sizeMb, unitMb := HumanReadableSize(1024 * 1024)

	if sizeMb != 1 && unitMb != "MB" {
		t.Error("Size is not correct (Mb test)")
	}

	sizeNil, unitNil := HumanReadableSize(0)

	if sizeNil != 0 && unitNil != "" {
		t.Error("Size is not correct (Nil test)")
	}

	// Do cleanup
	t.Cleanup(func() {
		os.RemoveAll("./test_tmp/")
	})
}
