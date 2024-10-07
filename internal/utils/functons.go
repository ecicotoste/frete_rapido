package utils

import (
	"fmt"
	"time"
)

func DifDias(estimated_date string) (int, error) {

	deadline := 0
	parsedDate, err := time.Parse("2006-01-02", estimated_date)
	if err != nil {
		fmt.Println("Erro ao analisar a data:", err)
		return deadline, err
	}

	currentDate := time.Now().Truncate(24 * time.Hour)

	diff := parsedDate.Sub(currentDate)
	if diff.Hours() < 24 {
		deadline = 1
	} else {
		deadline = int(diff.Hours() / 24)
	}

	return deadline, nil
}
