package handlers

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func (h *Handler) GetProjects(w http.ResponseWriter, r *http.Request) {
	tableName, ok := r.Context().Value("TableThree").(string)
	if !ok {
		log.Error("TableName not found in context")
		InternalError500(w, "TableThree", ErrTableNameNotFound)
		return
	}

	fields := log.Fields{"TableThree": tableName, "full_name": "EliFuchsman"}

	projects, err := h.h.ReturnProjects(tableName)
	if err != nil {
		log.Errorf("%+v", err)
		InternalError500(w, "Projects", err)
		return
	}

	if projects == nil {
		log.WithFields(fields).Errorf("%+v", err)
		NotFound404(w, "Projects")
		return
	}
	OK200(w, projects)
}
