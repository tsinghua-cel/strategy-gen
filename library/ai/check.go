package aiattack

import (
	"github.com/tsinghua-cel/strategy-gen/utils"
	"strconv"
)

func CheckDuties(maxValidatorIndex int, duties []utils.ProposerDuty) ([]interface{}, bool) {
	result := make([]interface{}, 0)
	for _, duty := range duties {
		// filter out duties with ValidatorIndex <= maxValidatorIndex
		dutyValIdx, _ := strconv.Atoi(duty.ValidatorIndex)
		if dutyValIdx <= maxValidatorIndex {
			result = append(result, duty)
		}
	}
	if len(result) > 0 {
		return result, true
	}
	return nil, false
}
