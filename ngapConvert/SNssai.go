package ngapConvert

import (
	"encoding/hex"

	"github.com/sadhasiva1984/ngap/logger"
	"github.com/sadhasiva1984/ngap/ngapType"
	"github.com/sadhasiva1984/openapi/models"
)

func SNssaiToModels(ngapSnssai ngapType.SNSSAI) (modelsSnssai models.Snssai) {
	modelsSnssai.Sst = int32(ngapSnssai.SST.Value[0])
	if ngapSnssai.SD != nil {
		modelsSnssai.Sd = hex.EncodeToString(ngapSnssai.SD.Value)
	}
	return
}

func SNssaiToNgap(modelsSnssai models.Snssai) ngapType.SNSSAI {
	var ngapSnssai ngapType.SNSSAI
	ngapSnssai.SST.Value = []byte{byte(modelsSnssai.Sst)}

	if modelsSnssai.Sd != "" {
		ngapSnssai.SD = new(ngapType.SD)
		if sdTmp, err := hex.DecodeString(modelsSnssai.Sd); err != nil {
			logger.NgapLog.Warnf("Decode snssai.sd failed: %+v", err)
		} else {
			ngapSnssai.SD.Value = sdTmp
		}
	}
	return ngapSnssai
}
