package archiveinfo

import (
	"archive/zip"
	"io"
	"mime"
	"mime/multipart"
	"os"
	"path/filepath"
)

func GetArchiveInfo(file multipart.File, fileHeader *multipart.FileHeader) (*FileHeader, error) {
	out, err := os.Create(fileHeader.Filename)
	if err != nil {
		return nil, err
	}
	defer out.Close()
	defer os.RemoveAll(fileHeader.Filename)

	_, err = io.Copy(out, file)
	if err != nil {
		return nil, err
	}

	reader, err := zip.OpenReader(fileHeader.Filename)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	filesInfo := []File{}
	counterOfFiles := 0
	sumOfSizes := 0

	for _, r := range reader.File {
		if r.UncompressedSize64 != 0 {

			counterOfFiles++
			sumOfSizes += int(r.UncompressedSize64)
			mimetype := mime.TypeByExtension(filepath.Ext(r.Name))

			fileInfo := File{
				FilePath: r.Name,
				Size:     float32(r.UncompressedSize64),
				Mimetype: mimetype,
			}

			filesInfo = append(filesInfo, fileInfo)

		}
	}
	archiveInfo := FileHeader{
		FileName:    fileHeader.Filename,
		ArchiveSize: float32(fileHeader.Size),
		TotalSize:   float32(sumOfSizes),
		TotalFiles:  float32(counterOfFiles),
		Files:       filesInfo,
	}
	return &archiveInfo, nil
}
