package handlers

import (
	"net/http"

	"github.com/phyrexxxxx/exhibition/utils"
)

func HandlerErr(w http.ResponseWriter, r *http.Request) {
	// 400 Bad Request
	utils.RespondWithError(w, http.StatusBadRequest, "Bad request")
}
