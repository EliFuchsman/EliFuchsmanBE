package handlers

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func (h *Handler) GetSummary(w http.ResponseWriter, r *http.Request) {
	filePath, ok := r.Context().Value("SummaryFile").(string)
	if !ok {
		log.Error("SummaryFile not found in context")
		InternalError500(w, "SummaryFile", ErrSummaryFileNotFound)
		return
	}

	fields := log.Fields{"filePath": filePath}
	summary, err := h.h.ReturnSummary(filePath)
	if err != nil {
		log.WithFields(fields).Errorf("%+v", err)
		InternalError500(w, "Summary", err)
		return
	}

	OK200(w, summary)
}
