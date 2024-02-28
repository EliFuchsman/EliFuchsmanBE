package handlers

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func (h *Handler) GetEducationHistory(w http.ResponseWriter, r *http.Request) {
	tableName, ok := r.Context().Value("TableOne").(string)
	if !ok {
		log.Error("TableName not found in context")
		InternalError500(w, "TableOne", ErrTableNameNotFound)
		return
	}

	fields := log.Fields{"TableOne": tableName, "full_name": "EliFuchsman"}

	educationHistory, err := h.h.ReturnEducationHistory(tableName)
	if err != nil {
		log.Errorf("%+v", err)
		InternalError500(w, "Education History", err)
		return
	}

	if educationHistory == nil {
		log.WithFields(fields).Errorf("%+v", err)
		NotFound404(w, "Education History")
		return
	}
	OK200(w, educationHistory)
}
