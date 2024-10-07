package handlefunc

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"api.frete.rapido/internal/entity"
	"api.frete.rapido/internal/process"
)

func QuoteSimulate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodPost {
		slog.Info("=========> FreteRapido: Quote Simulate")
		var jsonIN entity.JsonIN
		err := json.NewDecoder(r.Body).Decode(&jsonIN)
		if err != nil {
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}
		jsonReturn, err := process.QuoteSimulate(jsonIN)
		if err != nil {
			if jsonReturn.Erro.RC == 9999 {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		if jsonReturn.Erro.RC == 0 {
			json.NewEncoder(w).Encode(jsonReturn.Carrier)
		} else {
			json.NewEncoder(w).Encode(jsonReturn.Erro)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
