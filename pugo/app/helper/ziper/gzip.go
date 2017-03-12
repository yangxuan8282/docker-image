package ziper

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
	"os"
	"path"
)

// Gzip compresses bytes
func Gzip(data []byte) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	writer, _ := gzip.NewWriterLevel(buf, gzip.BestCompression)
	if _, err := writer.Write(data); err != nil {
		return nil, err
	}
	writer.Close()
	return buf.Bytes(), nil
}

// GzipFile compresses file bytes to bytes
func GzipFile(file string) ([]byte, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return Gzip(data)
}

// GzipFileBase64 compresses file bytes and outputs base64 string
func GzipFileBase64(file string) (string, error) {
	data, err := GzipFile(file)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(data), nil
}

// UnGzip uncompresses bytes
func UnGzip(data []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	return ioutil.ReadAll(reader)
}

// UnGzipFile uncompresses data to file
func UnGzipFile(data []byte, file string) error {
	output, err := UnGzip(data)
	if err != nil {
		return err
	}
	os.MkdirAll(path.Dir(file), os.ModePerm)
	return ioutil.WriteFile(file, output, os.ModePerm)
}

// UnGzipFileBase64 uncompresses base64 string to file
func UnGzipFileBase64(data, file string) error {
	output, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return err
	}
	return UnGzipFile(output, file)
}
