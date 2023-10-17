package validator

import (
	"errors"
	"net/textproto"
)

const (
	ContentType = "Content-Type"
	zipType     = "application/zip"
	first       = 1
	second      = 2
	third       = 3
)

func ValidatorOfzipType(mimeType textproto.MIMEHeader, nRouter int) error {
	m := map[int][]string{
		first:  {"application/zip"},
		second: {"application/vnd.openxmlformats-officedocument.wordprocessingml.document", "application/xml", "image/jpeg", "image/png"},
		third:  {"application/vnd.openxmlformats-officedocument.wordprocessingml.document", "application/pdf"},
	}

	switch nRouter {
	case first:
		if mimeType[ContentType][0] != m[first][0] {
			return errors.New("file is not a application/zip")
		}
	case second:
		if !contains(m[second], mimeType[ContentType][0]) {
			return errors.New("the file does not fit the type")
		}
	case third:
		if !contains(m[third], mimeType[ContentType][0]) {
			return errors.New("the file does not fit the type")
		}
	}

	return nil
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
