package ngapConvert

import (
	"github.com/sadhasiva1984/ngap/ngapType"
	"github.com/sadhasiva1984/openapi/models"
)

func AllowedNssaiToNgap(allowedNssaiModels []models.AllowedSnssai) (allowedNssaiNgap ngapType.AllowedNSSAI) {
	for _, allowedSnssai := range allowedNssaiModels {
		item := ngapType.AllowedNSSAIItem{
			SNSSAI: SNssaiToNgap(*allowedSnssai.AllowedSnssai),
		}
		allowedNssaiNgap.List = append(allowedNssaiNgap.List, item)
	}
	return
}

func AllowedNssaiToModels(allowedNssaiNgap ngapType.AllowedNSSAI) (allowedNssaiModels []models.AllowedSnssai) {
	for _, item := range allowedNssaiNgap.List {
		snssai := SNssaiToModels(item.SNSSAI)
		allowedSnssai := models.AllowedSnssai{
			AllowedSnssai: &snssai,
		}
		allowedNssaiModels = append(allowedNssaiModels, allowedSnssai)
	}
	return
}
