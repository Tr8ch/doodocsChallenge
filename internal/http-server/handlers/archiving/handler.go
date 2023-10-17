package archiving

import (
	"doodocs/internal/lib/logger/sl"
	"doodocs/internal/lib/validator"
	"errors"
	"net/http"

	resp "doodocs/internal/lib/api/response"

	"github.com/go-chi/render"
	"golang.org/x/exp/slog"
)

const nRouter = 2

func New(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.archiving.New"

		log := log.With(
			slog.String("op", op),
		)
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			log.Error("files more than 10MB", sl.Err(errors.New("")))

			render.JSON(w, r, resp.Error("your file is too big"))

			return
		}

		files := r.MultipartForm.File["files[]"]

		if len(files) == 0 {
			log.Error("empty request", sl.Err(errors.New("")))

			render.JSON(w, r, resp.Error("choose files"))

			return
		}

		for _, file := range files {
			if err := validator.ValidatorOfzipType(file.Header, nRouter); err != nil {
				log.Error("non valid file", sl.Err(err))

				render.JSON(w, r, resp.Error("check type of files"))

				return
			}
		}

		buf, err := Archiving(files)
		if err != nil {
			log.Error("failed to archive files", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to archive files"))

			return
		}

		w.Header().Set("Content-Type", "application/zip")
		w.Header().Set("Content-Disposition", "attachment; filename=doodocs.zip")
		w.WriteHeader(http.StatusOK)
		w.Write(buf.Bytes())
	}
}
