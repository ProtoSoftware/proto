/*
Copyright © 2022 BitsOfAByte

*/
package backend

import (
	"archive/tar"
	"compress/gzip"
	"crypto"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/cheggaaa/pb/v3"
	"github.com/spf13/viper"
)

// Correctly formats a path for the program.
func UsePath(path string, trailSlash bool) string {

	Debug("UsePath: Attempting to format path: " + path)

	// If trail slash is true, add a trailing slash to the path
	if path[len(path)-1:] != "/" && trailSlash {
		path = path + "/"
	}

	// If trail slash is false, remove a trailing slash from the path
	if path[len(path)-1:] == "/" && !trailSlash {
		path = path[:len(path)-1]
	}

	// If short notation for the home directory is used, expand it.
	if path[:2] == "~/" {
		homeDir, _ := os.UserHomeDir()
		path = filepath.Join(homeDir, path[2:])
	}

	Debug("UsePath: Finished formatting path, result was: " + path)

	return path
}

func ClearTemp() error {
	err := os.RemoveAll(UsePath(viper.GetString("app.temp_storage"), false))
	if err != nil {
		return err
	}

	Debug("ClearTemp: Cleaned up temp directory")

	return nil
}

// Downloads the file from the given URL, following redirects if needed. The final file will be put at the given path
// and a progress bar will be output to the standard output while downloading.
func DownloadFile(path, url string) (os.FileInfo, error) {

	// If path doesnt exist create it
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
		if err != nil {
			return nil, err
		}

		Debug("DownloadFile: Created directory: " + filepath.Dir(path))
	}

	// Create the file
	out, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	defer out.Close()

	// Fetch the file from the URL
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	Debug("DownloadFile: Downloading file from: " + url)

	tmpl := `{{ cycle . "⠃" "⠆" "⠤" "⠰" "⠘" "⠉" }} Installing {{string . "src"}} [{{percent .}} | {{speed . "%s/s"}} | {{ rtime .}}]`
	bar := pb.ProgressBarTemplate(tmpl).Start64(resp.ContentLength).Set("src", strings.Split(url, "/")[len(strings.Split(url, "/"))-1])
	reader := bar.NewProxyReader(resp.Body)

	defer resp.Body.Close()

	// Write the data to the file
	_, err = io.Copy(out, reader)
	if err != nil {
		return nil, err
	}

	// Check if the file is valid
	if _, err := os.Stat(path); err != nil {
		return nil, err
	}

	Debug("DownloadFile: Downloaded file to: " + path)

	bar.Finish()

	// Get the downloaded file and return it
	return os.Stat(path)
}

// Extracts a tarball to the given path
func ExtractTar(path string, r io.Reader) error {

	// If path doesnt exist create it
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
		if err != nil {
			return err
		}

		Debug("ExtractTar: Created directory: " + filepath.Dir(path))
	}

	gzr, err := gzip.NewReader(r)
	if err != nil {
		return err
	}

	defer gzr.Close()
	tr := tar.NewReader(gzr)

	Debug("ExtractTar: Extracting tarball to: " + path)

	for {
		header, err := tr.Next()

		switch {
		// No more files to extract
		case err == io.EOF:
			return nil
		// Something went wrong, return the error.
		case err != nil:
			return err
		// Skip all nil headers.
		case header == nil:
			continue
		}

		Debug("ExtractTar: Inflating: " + header.Name)

		// Send all files to the path given
		target := filepath.Join(path, header.Name)

		switch header.Typeflag {

		// Create directory if it doesn't exist
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					return err
				}
			}

		// Create files with their contents
		case tar.TypeReg:
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

			// Copy the contents of the file
			if _, err := io.Copy(f, tr); err != nil {
				return err
			}

			// Close the file after copying, do not defer
			f.Close()
		}
	}
}

// Check a given files sum against the given sum (Sha512)
func MatchChecksum(filePath, sumPath string) (bool, error) {
	// Get the sum of the file with crypto inbuilt
	h := crypto.SHA512.New()
	f, err := os.Open(filePath)
	if err != nil {
		return false, err
	}

	defer f.Close()

	if _, err := io.Copy(h, f); err != nil {
		return false, err
	}

	// Get the sum of the file in the sum file
	sum, err := ioutil.ReadFile(sumPath)
	if err != nil {
		return false, err
	}

	// Check all lines for the files sum
	for _, line := range strings.Split(string(sum), "\n") {
		Debug("MatchChecksum: Attempting to match checksum: " + strings.TrimSuffix(line, "\n") + " to: " + string(h.Sum(nil)))
		if strings.HasPrefix(line, fmt.Sprintf("%x", h.Sum(nil))) {
			return true, nil
		}
	}

	return false, nil
}

// Gets the size of the given directory in bytes.
func DirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

// Converts the given bytes to a human readable amount of bytes and a unit.
func HumanReadableSize(bytes int64) (int64, string) {
	switch {
	case bytes < 1024:
		return bytes, "B"
	case bytes < 1024*1024:
		return bytes / 1024, "KB"
	case bytes < 1024*1024*1024:
		return bytes / (1024 * 1024), "MB"
	default:
		return bytes / (1024 * 1024 * 1024), "GB"
	}
}
