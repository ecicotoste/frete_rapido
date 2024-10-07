package main

import (
	"encoding/json"
	"log"
	"net/http"

	"api.frete.rapido/internal/handlefunc"
)

func main() {

	//---> GET ECHO
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("ECHO FRETE.RAPIDO V0001")
	})
	//---> GET ECHO

	//---> POST FreteRapido: Quote Simulate
	http.HandleFunc("/quote", handlefunc.QuoteSimulate)
	//--->

	//---> GET FreteRapido: metrics
	http.HandleFunc("/metrics", handlefunc.LastQuotesMetrics)
	//--->

	log.Fatal(http.ListenAndServe(":6379", nil))
}

//docker-compose up -d --build --force-recreate
