package archiveinfo

import (
	resp "doodocs/internal/lib/api/response"
	"doodocs/internal/lib/logger/sl"
	"doodocs/internal/lib/validator"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/render"
	"golang.org/x/exp/slog"
)

type FileHeader struct {
	FileName    string  `json:"filename"`
	ArchiveSize float32 `json:"archive_size"`
	TotalSize   float32 `json:"total_size"`
	TotalFiles  float32 `json:"total_files"`
	Files       []File  `json:"files"`
}

type File struct {
	FilePath string  `json:"file_path"`
	Size     float32 `json:"size"`
	Mimetype string  `json:"mimetype"`
}

type Response struct {
	resp.Response
	FileHeader
}

const nRouter = 1

func New(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.archiveInfo.New"

		log := log.With(
			slog.String("op", op),
		)

		file, fileHeader, err := r.FormFile("file")
		if errors.Is(err, io.EOF) {
			log.Error("request body is empty")

			render.JSON(w, r, resp.Error("empty request"))

			return
		}
		if err != nil {
			log.Error("failed to get archive from client", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to get archive, choose a archive"))

			return

		}
		defer file.Close()

		log.Info("archive geted", slog.Any("request filename and mimetype", fmt.Sprintf("Name:%s; Type:%s", fileHeader.Filename, fileHeader.Header["Content-Type"][0])))

		if err := validator.ValidatorOfzipType(fileHeader.Header, nRouter); err != nil {
			log.Error("invalid request is't zip file", sl.Err(err))

			render.JSON(w, r, resp.Error("use zip archive"))

			return
		}

		archiveInfo, err := GetArchiveInfo(file, fileHeader)
		if err != nil {
			log.Error("failed to retrieve information from archive", sl.Err(err))

			render.JSON(w, r, resp.Error("we cannot process your archive"))

			return
		}

		log.Info("archive info geted")

		responseOK(w, r, archiveInfo)
	}
}

func responseOK(w http.ResponseWriter, r *http.Request, FileHeader *FileHeader) {
	render.JSON(w, r, Response{
		Response:   resp.OK(),
		FileHeader: *FileHeader,
	})
}
