package emailsender

import (
	resp "doodocs/internal/lib/api/response"
	"doodocs/internal/lib/logger/sl"
	"doodocs/internal/lib/validator"
	"net/http"

	"github.com/go-chi/render"
	"golang.org/x/exp/slog"
)

const nRouter = 3

func New(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.archiving.New"

		log := log.With(
			slog.String("op", op),
		)

		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			log.Error("files more than 10MB", sl.Err(err))

			render.JSON(w, r, resp.Error("your file is too big"))

			return
		}

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			log.Error("Unable to get file from form", sl.Err(err))

			render.JSON(w, r, resp.Error("choose a file"))

			return
		}
		defer file.Close()
		if err := validator.ValidatorOfzipType(fileHeader.Header, nRouter); err != nil {
			log.Error("non valid file", sl.Err(err))

			render.JSON(w, r, resp.Error("check type of file"))

			return
		}

		emails := r.Form["emails"]

		err = sendEmails(emails, file, fileHeader)
		if err != nil {
			log.Error("Error sending emails", sl.Err(err))

			render.JSON(w, r, resp.Error("something wrong with your mail, check if you have written correctly"))

			return
		}

		w.WriteHeader(http.StatusOK)
		log.Info("Emails sent successfully")
	}
}
