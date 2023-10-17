package archiving

import (
	"archive/zip"
	"bytes"
	"io"
	"mime/multipart"
)

func Archiving(files []*multipart.FileHeader) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer src.Close()

		zipEntry, err := zipWriter.Create(file.Filename)
		if err != nil {
			return nil, err
		}

		_, err = io.Copy(zipEntry, src)
		if err != nil {
			return nil, err
		}
	}

	err := zipWriter.Close()
	if err != nil {
		return nil, err
	}

	return buf, nil
}
