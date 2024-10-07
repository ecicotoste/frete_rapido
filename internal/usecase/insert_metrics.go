package usecase

import "api.frete.rapido/internal/entity"

type InsertMetricsInputDto struct {
	ID         string
	IdTransp   string
	Company    string
	FinalPrice float64
}

type InsertMetricsUsecase struct {
	MetricsRepository entity.MetricsRepository
}

func NewInsertMetricsUsecase(metricsRepository entity.MetricsRepository) *InsertMetricsUsecase {
	return &InsertMetricsUsecase{MetricsRepository: metricsRepository}
}

func (u *InsertMetricsUsecase) Execute(input InsertMetricsInputDto) error {
	metrics := entity.NewMetrics(input.ID, input.IdTransp, input.Company, input.FinalPrice)
	err := u.MetricsRepository.Insert(metrics)
	if err != nil {
		return err
	}

	return nil
}
