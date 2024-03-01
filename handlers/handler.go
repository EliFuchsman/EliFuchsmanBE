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
var ErrAWSRegionNotFound = errors.New("AWS region not found in context")
var ErrBucketNotFound = errors.New("Bucket not found in context")
var ErrBucketKeyNotFound = errors.New("Bucket key not found in context")
var ErrFailedAWSSession = errors.New("Error starting an aws session")
var ErrNoResponseFromAWS = errors.New("Error retrieving from a response AWS")

func NewHandler(h elifuchsman.Client) *Handler {
	return &Handler{
		h: h,
	}
}
