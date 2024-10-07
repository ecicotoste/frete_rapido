package usecase

import "api.frete.rapido/internal/entity"

type QueryMetricsOutputDto struct {
	TotResTransp      int
	Carrier           string
	TotalFinalPrice   float64
	AverageFinalPrice float64
	MinAllPrice       float64
	MaxAllPrice       float64
}

type QueryMetricsUseCase struct {
	MetricsRepository entity.MetricsRepository
}

func NewQueryMetricsUseCase(MetricsRepository entity.MetricsRepository) *QueryMetricsUseCase {
	return &QueryMetricsUseCase{MetricsRepository: MetricsRepository}
}

func (u *QueryMetricsUseCase) Execute(limQuote string) ([]*QueryMetricsOutputDto, error) {
	ListMetrics, err := u.MetricsRepository.Query(limQuote)
	if err != nil {
		println(err.Error())
		return nil, err
	}

	var ListQueryMetricsOutputDto []*QueryMetricsOutputDto
	for _, Metrics := range ListMetrics {
		ListQueryMetricsOutputDto = append(ListQueryMetricsOutputDto, &QueryMetricsOutputDto{
			TotResTransp:      Metrics.TotResTransp,
			Carrier:           Metrics.Carrier,
			TotalFinalPrice:   Metrics.TotalFinalPrice,
			AverageFinalPrice: Metrics.AverageFinalPrice,
			MinAllPrice:       Metrics.MinAllPrice,
			MaxAllPrice:       Metrics.MaxAllPrice,
		})
	}

	return ListQueryMetricsOutputDto, nil
}
