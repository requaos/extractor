package pdfinspect

import (
	"bytes"
	"strings"

	"github.com/pkg/errors"
	"github.com/unidoc/unidoc/pdf/extractor"
	"github.com/unidoc/unidoc/pdf/model"
)

func Inspect(b []byte) (string, error) {
	br := bytes.NewReader(b)

	pr, err := model.NewPdfReader(br)
	if err != nil {
		return "", errors.Wrap(err, "error calling NewPdfReader")
	}

	var s string
	sb := strings.Builder{}
	for i := range pr.PageList {
		ex, err := extractor.New(pr.PageList[i])
		if err != nil {
			return "", errors.Wrap(err, "error during pagelist iteration")
		}
		s, err = ex.ExtractText()
		if err != nil {
			return "", errors.Wrap(err, "error extracting text")
		}
		_, err = sb.WriteString(s + " ")
		if err != nil {
			return "", errors.Wrap(err, "error writing to string builder")
		}
	}
	return sb.String(), nil
}
