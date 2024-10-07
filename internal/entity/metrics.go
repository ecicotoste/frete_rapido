package entity

type MetricsRepository interface {
	Insert(*Metrics) error
	Query(limQuote string) ([]*RMetrics, error)
}

type Metrics struct {
	ID         string
	IdTransp   string
	CompName   string
	FinalPrice float64
}

func NewMetrics(id string, idTransp string, compName string, finalPrice float64) *Metrics {
	return &Metrics{
		ID:         id,
		IdTransp:   idTransp,
		CompName:   compName,
		FinalPrice: finalPrice,
	}
}

type RMetrics struct {
	TotResTransp      int
	Carrier           string
	TotalFinalPrice   float64
	AverageFinalPrice float64
	MinAllPrice       float64
	MaxAllPrice       float64
}
