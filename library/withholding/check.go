package withholding

import (
	"github.com/tsinghua-cel/strategy-gen/types"
	"github.com/tsinghua-cel/strategy-gen/utils"
	"strconv"
)

func CheckDuties(param types.LibraryParams, duties []utils.ProposerDuty) ([]interface{}, bool) {
	result := make([]interface{}, 0)

	tmpsub := make([]utils.ProposerDuty, 0)
	for _, duty := range duties {
		valIdx, _ := strconv.Atoi(duty.ValidatorIndex)

		if types.IsHackValidator(valIdx, param) {
			tmpsub = append(tmpsub, duty)
		} else {
			if len(tmpsub) > 5 {
				result = append(result, tmpsub)
			}
			tmpsub = make([]utils.ProposerDuty, 0)
		}
	}
	if len(tmpsub) > 5 {
		result = append(result, tmpsub)
	}

	if len(result) > 0 {
		return result, true
	}

	return nil, false
}
