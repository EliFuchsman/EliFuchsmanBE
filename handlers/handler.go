package handlers

import elifuchsman "github.com/efuchsman/EliFuchsmanBE/internal/eli_fuchsman"

type Handler struct {
	h elifuchsman.Client
}

func NewHandler(h elifuchsman.Client) *Handler {
	return &Handler{
		h: h,
	}
}
