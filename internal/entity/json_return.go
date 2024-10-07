package entity

type JsonReturn struct {
	Erro        JsonRetErr      `json:"erro"`
	Carrier     JsonRetCarriers `json:"carriers"`
	RettMetrics ResMetrics      `json:"ret_metrics"`
}

type JsonRetErr struct {
	RC       int      `json:"rc"`
	DescErro []string `json:"descerro"`
}

type JsonRetCarriers struct {
	Carriers []Carrier `json:"carrier"`
}

type Carrier struct {
	Name     string  `json:"name"`
	Service  string  `json:"service"`
	Deadline string  `json:"deadline"`
	Price    float64 `json:"price"`
}

type ResMetrics struct {
	MinAllPrice     float64          `json:"minim_all_price"`
	MaxAllPrice     float64          `json:"maxim_all_price"`
	MetricsCarriers []MetricsCarrier `json:"metrics_carrier"`
}

type MetricsCarrier struct {
	TotResTransp      int     `json:"total_result_carrier"`
	CompanyName       string  `json:"company_name"`
	TotalFinalPrice   float64 `json:"total_final_price"`
	AverageFinalPrice float64 `json:"average_final_price"`
}
