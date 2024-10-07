package confuse

import (
	"github.com/tsinghua-cel/strategy-gen/types"
	"github.com/tsinghua-cel/strategy-gen/utils"
	"strconv"
)

func CheckDuties(param types.LibraryParams, duties []utils.ProposerDuty) ([]interface{}, bool) {
	result := make([]interface{}, 0)
	for _, duty := range duties {
		// filter out duties with ValidatorIndex <= maxValidatorIndex
		dutyValIdx, _ := strconv.Atoi(duty.ValidatorIndex)
		if types.IsHackValidator(dutyValIdx, param) {
			result = append(result, duty)
		}
	}
	if len(result) > 0 {
		return result, true
	}
	return nil, false
}
