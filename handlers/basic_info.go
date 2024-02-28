package handlers

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func (h *Handler) GetBasicInfo(w http.ResponseWriter, r *http.Request) {
	filePath, ok := r.Context().Value("InfoFile").(string)
	if !ok {
		log.Error("InfoFile not found in context")
		InternalError500(w, "InfoFile", ErrTableNameNotFound)
		return
	}

	fields := log.Fields{"filePath": filePath}
	basicInfo, err := h.h.ReturnBasicInfo(filePath)
	if err != nil {
		log.WithFields(fields).Errorf("%+v", err)
		InternalError500(w, "Basic Info", err)
		return
	}

	OK200(w, basicInfo)
}
