package aiattack

import (
	log "github.com/sirupsen/logrus"
	"github.com/tsinghua-cel/strategy-gen/types"
	"github.com/tsinghua-cel/strategy-gen/utils"
	"time"
)

type Instance struct {
}

func (o *Instance) Name() string {
	return "ai"
}

func (o *Instance) Description() string {
	desc_eng := `ai generate strategy.`
	return desc_eng
}

func (o *Instance) Run(params types.LibraryParams) {
	log.WithField("name", o.Name()).Info("start to run strategy")
	var latestEpoch int64
	ticker := time.NewTicker(time.Second * 3)
	slotTool := utils.SlotTool{SlotsPerEpoch: 32}
	for {
		select {
		case <-ticker.C:
			slot, err := utils.GetCurSlot(params.Attacker)
			if err != nil {
				log.WithField("error", err).Error("failed to get slot")
				continue
			}
			epoch := slotTool.SlotToEpoch(int64(slot))
			if epoch == latestEpoch {
				continue
			}
			if int64(slot) < slotTool.EpochEnd(epoch) {
				// only update strategy at the end of current epoch.
				continue
			}
			latestEpoch = epoch
			// get next epoch duties
			duties, err := utils.GetEpochDuties(params.Attacker, epoch+1)
			if err != nil {
				log.WithFields(log.Fields{
					"error": err,
					"epoch": epoch + 1,
				}).Error("failed to get duties")
				latestEpoch = epoch - 1
				continue
			}
			if hackDuties, happen := CheckDuties(params, duties); happen {
				strategy := types.Strategy{}
				strategy.Slots = GenSlotStrategy(hackDuties)
				if strategy.Slots == nil {
					log.Error("failed to generate slot strategy")
					continue
				}
				if err = utils.UpdateStrategy(params.Attacker, strategy); err != nil {
					log.WithField("strategy", strategy).WithError(err).Error("failed to update strategy")
				} else {
					log.WithFields(log.Fields{
						"epoch":    epoch + 1,
						"strategy": strategy,
					}).Info("update strategy successfully")
				}
			}
		}
	}
}
