package handlers

import (
	"errors"

	elifuchsman "github.com/efuchsman/EliFuchsmanBE/internal/eli_fuchsman"
)

type Handler struct {
	h elifuchsman.Client
}

var ErrTableNameNotFound = errors.New("TableName not found in context")
var ErrInfoFileNotFound = errors.New("InfoFile not found in context")
var ErrSummaryFileNotFound = errors.New("SummaryFile not found in context")

func NewHandler(h elifuchsman.Client) *Handler {
	return &Handler{
		h: h,
	}
}
