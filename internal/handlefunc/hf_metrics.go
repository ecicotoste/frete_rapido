package handlefunc

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"api.frete.rapido/internal/process"
)

func LastQuotesMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		slog.Info("=========> FreteRapido: Last Quotes Metrics")
		query := r.URL.Query()
		last_quotes := query.Get("last_quotes")
		slog.Info("=========> query.Get(last_quotes): " + last_quotes)
		if len(last_quotes) > 0 {
			_, err := strconv.Atoi(last_quotes)
			if err != nil {
				http.Error(w, "INVALID last_quotes", http.StatusBadRequest)
				return
			}
		}

		jsonReturn, err := process.LastQuotesMetrics(last_quotes)
		if err != nil {
			if jsonReturn.Erro.RC == 9999 {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
		if jsonReturn.Erro.RC == 0 {
			json.NewEncoder(w).Encode(jsonReturn.RettMetrics)
		} else {
			json.NewEncoder(w).Encode(jsonReturn.Erro)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
