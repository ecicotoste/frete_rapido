package process

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log/slog"
	"strconv"

	"api.frete.rapido/internal/entity"
	"api.frete.rapido/internal/repository"
	"api.frete.rapido/internal/usecase"
	"api.frete.rapido/internal/utils"
	"api.frete.rapido/internal/validations"
	"api.frete.rapido/internal/ws"
)

func QuoteSimulate(jsonIN entity.JsonIN) (entity.JsonReturn, error) {

	var jsonReturn entity.JsonReturn

	strErros := validations.ValidationIn(jsonIN)
	if len(strErros) > 0 {
		jsonReturn.Erro.RC = 1
		jsonReturn.Erro.DescErro = strErros
		return jsonReturn, nil
	}

	var jsonFrReq entity.JsonFrReqQuoteSimulate
	jsonFrReq.Shipper.RegisteredNumber = "25438296000158"
	jsonFrReq.Shipper.Token = "1d52a9b6b78cf07b08586152459a5c90"
	jsonFrReq.Shipper.PlatformCode = "5AKVkHqCn"
	jsonFrReq.SimulationType = []int{0}
	jsonFrReq.Recipient.Type = 1
	jsonFrReq.Recipient.Country = "BRA"
	cep, _ := strconv.Atoi(jsonIN.Recipient.Address.Zipcode)
	jsonFrReq.Recipient.Zipcode = cep
	var disparcher entity.Dispatcher
	disparcher.RegisteredNumber = "25438296000158"
	disparcher.Zipcode = cep
	for _, volume := range jsonIN.Volumes {
		var frVolume entity.Volume
		frVolume.Amount = volume.Amount
		frVolume.Category = fmt.Sprint(volume.Category)
		frVolume.Height = volume.Height
		frVolume.Width = volume.Width
		frVolume.Length = volume.Length
		frVolume.UnitaryPrice = volume.Price
		frVolume.UnitaryWeight = volume.UnitaryWeight
		disparcher.Volumes = append(disparcher.Volumes, frVolume)
		jsonFrReq.Dispatchers = append(jsonFrReq.Dispatchers, disparcher)
	}

	payloadBytes, err := json.Marshal(jsonFrReq)
	if err != nil {
		slog.Error("json.Marshal(jsonFrReq): " + err.Error())
		jsonReturn.Erro.RC = 9999
		return jsonReturn, err
	}
	slog.Info("======> Request : " + string(payloadBytes))

	body, err := ws.WSRequest("POST", "https://sp.freterapido.com/api/v3/quote/simulate", payloadBytes)
	if err != nil {
		slog.Error("ws.WSRequest: " + err.Error())
		jsonReturn.Erro.RC = 9999
		return jsonReturn, err
	}
	slog.Info("======> Response: " + string(body))

	var jsonFrRespQuoteSimulate entity.JsonFrRespQuoteSimulate
	err = json.Unmarshal(body, &jsonFrRespQuoteSimulate)
	if err != nil {
		slog.Error("json.Unmarshal(body, &jsonFrRespQuoteSimulate): " + err.Error())
		jsonReturn.Erro.RC = 9999
		return jsonReturn, err
	}

	for _, disparcher := range jsonFrRespQuoteSimulate.Dispatchers {
		for _, offer := range disparcher.Offers {
			var carrier entity.Carrier
			carrier.Name = offer.Carrier.Name
			carrier.Service = offer.Service
			if len(offer.DeliveryTime.EstimatedDate) > 0 {
				deadline, err := utils.DifDias(offer.DeliveryTime.EstimatedDate)
				if err != nil {
					slog.Error("utils.DifDias(offer.DeliveryTime.EstimatedDate): " + err.Error())
					jsonReturn.Erro.RC = 9999
					return jsonReturn, err
				}
				carrier.Deadline = fmt.Sprint(deadline)
			} else {
				carrier.Deadline = "1"
			}
			carrier.Price = offer.CostPrice
			jsonReturn.Carrier.Carriers = append(jsonReturn.Carrier.Carriers, carrier)
			go InsertMetrics(disparcher.ID, offer.Carrier.RegisteredNumber, offer.Carrier.CompanyName, offer.FinalPrice)
		}
	}

	return jsonReturn, nil
}

func LastQuotesMetrics(limQuotes string) (entity.JsonReturn, error) {
	var jsonReturn entity.JsonReturn

	slog.Info("=======> Last Quotes Matrics ")

	db, err := sql.Open("mysql", "root:root@tcp(db:3306)/dbmysql")
	if err != nil {
		slog.Error("=======> sql.Open: " + err.Error())
		jsonReturn.Erro.RC = 9999
		return jsonReturn, err
	}
	defer db.Close()

	slog.Info("=======> sql.Open OK ")

	repository := repository.NewMetricsRepositoryMysql(db)
	queryMatricsUseCase := usecase.NewQueryMetricsUseCase(repository)
	if len(limQuotes) > 0 {
		limQuotes = " LIMIT " + limQuotes
	}
	queryMetricsOut, err := queryMatricsUseCase.Execute(limQuotes)
	if err != nil {
		slog.Error("=======> queryMatricsUseCase.Execute(input): " + err.Error())
		jsonReturn.Erro.RC = 9999
		return jsonReturn, err
	}

	slog.Info("=======> queryMatricsUseCase.Execute(limQuotes) OK ")
	totReg := len(queryMetricsOut)
	srt := fmt.Sprintf("Total de registros Metrics: %d ", totReg)
	slog.Info(srt)

	if totReg > 0 {
		jsonReturn.RettMetrics.MinAllPrice = queryMetricsOut[0].MinAllPrice
		jsonReturn.RettMetrics.MaxAllPrice = queryMetricsOut[0].MaxAllPrice
		for _, metrics := range queryMetricsOut {
			var MetricsCarrier entity.MetricsCarrier
			MetricsCarrier.TotResTransp = metrics.TotResTransp
			MetricsCarrier.CompanyName = metrics.Carrier
			MetricsCarrier.TotalFinalPrice = metrics.TotalFinalPrice
			MetricsCarrier.AverageFinalPrice = metrics.AverageFinalPrice
			jsonReturn.RettMetrics.MetricsCarriers = append(jsonReturn.RettMetrics.MetricsCarriers, MetricsCarrier)
		}

	}

	return jsonReturn, nil
}
