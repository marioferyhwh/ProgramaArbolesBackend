package commons

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/marioferyhwh/IMFBackend_forest/models"
)

//DisplayMessage estructura de respuesta
func DisplayMessage(w http.ResponseWriter, m *models.Message) {
	j, err := json.Marshal(m)
	if err != nil {
		log.Fatal("no se comvirtio mensaje->", err)
	}
	w.WriteHeader(m.Code)
	w.Write(j)
}
