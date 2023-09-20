package handlers

import (
	"net/http"

	"github.com/phyrexxxxx/exhibition/utils"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}
