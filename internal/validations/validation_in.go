package validations

import (
	"fmt"
	"strconv"

	"api.frete.rapido/internal/entity"
)

func ValidationIn(jsonIN entity.JsonIN) []string {

	var arrErros []string

	if len(jsonIN.Recipient.Address.Zipcode) > 8 {
		arrErros = append(arrErros, "Recipient.Address.Zipcode INVALID: More Than 8 Characters")
	}

	_, err := strconv.Atoi(jsonIN.Recipient.Address.Zipcode)
	if err != nil {
		arrErros = append(arrErros, "Recipient.Address.Zipcode INVALID: Not Numeric")
	}

	if len(jsonIN.Volumes) == 0 {
		arrErros = append(arrErros, "Volumes INVALID: Must Be More Than Zero")
	} else {
		for idx, volume := range jsonIN.Volumes {
			if volume.Amount <= 0 {
				strErr := fmt.Sprintf("Volumes[%d].Amount INVALID", idx)
				arrErros = append(arrErros, strErr)
			}
			if volume.Price <= 0 {
				strErr := fmt.Sprintf("Volumes[%d].Price INVALID", idx)
				arrErros = append(arrErros, strErr)
			}
			if volume.UnitaryWeight <= 0 {
				strErr := fmt.Sprintf("Volumes[%d].UnitaryWeight INVALID", idx)
				arrErros = append(arrErros, strErr)
			}
			if volume.Height <= 0 {
				strErr := fmt.Sprintf("Volumes[%d].Height INVALID", idx)
				arrErros = append(arrErros, strErr)
			}
			if volume.Width <= 0 {
				strErr := fmt.Sprintf("Volumes[%d].Width INVALID", idx)
				arrErros = append(arrErros, strErr)
			}
			if volume.Length <= 0 {
				strErr := fmt.Sprintf("Volumes[%d].Length INVALID", idx)
				arrErros = append(arrErros, strErr)
			}
			if volume.Category <= 0 {
				strErr := fmt.Sprintf("Volumes[%d].Category INVALID", idx)
				arrErros = append(arrErros, strErr)
			}
		}
	}
	return arrErros
}
