package backend

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/cheggaaa/pb/v3"
	"github.com/spf13/viper"
)

// Correctly formats a path for the program
func UsePath(path string, pathType string) string {

	// If the type is a directory, add a trailing slash
	if path[len(path)-1:] != "/" && pathType == "dir" {
		path = path + "/"
	}

	// If type is a file, remove the trailing slash
	if path[len(path)-1:] == "/" && pathType == "file" {
		path = path[:len(path)-1]
	}

	// If short notation for the home directory is used, expand it.
	if path[:2] == "~/" {
		homeDir, _ := os.UserHomeDir()
		path = filepath.Join(homeDir, path[2:])
	}

	return path
}

func ClearTemp() error {
	err := os.RemoveAll(viper.GetString("app.temp_storage"))

	if err != nil {
		return err
	}

	return nil
}

func DownloadFile(path string, url string) (os.FileInfo, error) {

	// Create the path if it does not exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0755)

		if _, err := os.Stat(path); os.IsNotExist(err) {
			return nil, err
		}
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

	// Set up a progress bar
	tmpl := `{{ cycle . "⠃" "⠆" "⠤" "⠰" "⠘" "⠉" }} Installing {{string . "src"}} [{{percent .}} | {{speed . "%s/s"}} | {{ rtime .}}]`
	bar := pb.ProgressBarTemplate(tmpl).Start64(resp.ContentLength).Set("src", strings.Split(url, "/")[len(strings.Split(url, "/"))-1])
	reader := bar.NewProxyReader(resp.Body)

	defer resp.Body.Close()

	// Write the data to the file
	_, err = io.Copy(out, reader)

	if err != nil {
		return nil, err
	}

	// Close the file
	out.Close()

	// Check if the file is valid
	if _, err := os.Stat(path); err != nil {
		return nil, err
	}

	bar.Finish()

	// Get the downloaded file and return it
	return os.Stat(path)
}

// Extracts a tarball to the given path
func ExtractTar(path string, r io.Reader) error {

	// Create the path if it does not exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0755)

		if _, err := os.Stat(path); os.IsNotExist(err) {
			return err
		}
	}

	gzr, err := gzip.NewReader(r)

	if err != nil {
		return err
	}

	defer gzr.Close()
	tr := tar.NewReader(gzr)

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

func HumanReadableSize(bytes int64) (int64, string) {
	// COnvert the bytes to a human readable file size

	// If the file size is less than 1 KB, return the size

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
