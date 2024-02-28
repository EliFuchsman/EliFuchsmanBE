package handlers

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func (h *Handler) GetExperienceHistory(w http.ResponseWriter, r *http.Request) {
	tableName, ok := r.Context().Value("TableTwo").(string)
	if !ok {
		log.Error("TableName not found in context")
		InternalError500(w, "TableTwo", ErrTableNameNotFound)
		return
	}

	fields := log.Fields{"TableTwo": tableName, "full_name": "EliFuchsman"}

	expHistory, err := h.h.ReturnExperienceHistory(tableName)
	if err != nil {
		log.Errorf("%+v", err)
		InternalError500(w, "Experience History", err)
		return
	}

	if expHistory == nil {
		log.WithFields(fields).Errorf("%+v", err)
		NotFound404(w, "Experience History")
		return
	}
	OK200(w, expHistory)
}
