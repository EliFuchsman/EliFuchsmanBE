package handlers

import (
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var ErrTableNameNotFound = errors.New("tableName not found in context")

func (h *Handler) GetBasicInfo(w http.ResponseWriter, r *http.Request) {
	tableName, ok := r.Context().Value("tableName").(string)
	if !ok {
		log.Error("tableName not found in context")
		InternalError500(w, "tableName", ErrTableNameNotFound)
		return
	}

	fields := log.Fields{"tableName": tableName}

	basicInfo, err := h.h.ReturnBasicInfo(tableName)
	if err != nil {
		log.Errorf("%+v", err)
		InternalError500(w, "region", err)
		return
	}

	if basicInfo == nil {
		log.WithFields(fields).Errorf("%+v", err)
		NotFound404(w, "basic_info")
		return
	}
	OK200(w, basicInfo)
}
