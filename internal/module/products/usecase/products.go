package usecase

import (
	"errors"
	"service/pkg/crashy"
)

func (p *products) Check() (string, error) {
	err := errors.New("tes error wrap")
	return "", crashy.Wrap(err, crashy.ErrCodeForbidden, "message palsu")
}
