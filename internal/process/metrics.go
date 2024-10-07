package process

import (
	"database/sql"
	"log/slog"

	"api.frete.rapido/internal/repository"
	"api.frete.rapido/internal/usecase"
)

func InsertMetrics(id string, registered_number string, company_name string, final_price float64) {

	slog.Info("=======> InsertMetrics ")

	db, err := sql.Open("mysql", "root:root@tcp(db:3306)/dbmysql")
	if err != nil {
		slog.Error("=======> sql.Open: " + err.Error())
		return
	}
	defer db.Close()

	slog.Info("=======> sql.Open OK ")

	repository := repository.NewMetricsRepositoryMysql(db)
	insertMetricsUsecase := usecase.NewInsertMetricsUsecase(repository)
	var input usecase.InsertMetricsInputDto
	input.ID = id
	input.IdTransp = registered_number
	input.Company = company_name
	input.FinalPrice = final_price
	err = insertMetricsUsecase.Execute(input)
	if err != nil {
		slog.Error("=======> insertMetricsUsecase.Execute(input): " + err.Error())
		return
	}

	slog.Info("=======> insertMetricsUsecase.Execute(input) OK ")
}
