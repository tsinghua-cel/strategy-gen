package syncwrong

import (
	"context"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/tsinghua-cel/strategy-gen/globalinfo"
	"github.com/tsinghua-cel/strategy-gen/types"
	"github.com/tsinghua-cel/strategy-gen/utils"
	"time"
)

type Instance struct {
}

func (o *Instance) Name() string {
	return "syncwrong"
}

func (o *Instance) Description() string {
	//	desc_cn := `
	//假设当前epoch = 0, 那么 在epoch=1 时，所有做恶验证者的投票不广播;
	//在 epoch=2 时，所有做恶验证者的投票不广播;
	//在 epoch=2的 最后一个做恶节点出块时，打包之前的所有做恶验证者的投票，并在 epoch=4的最后一个slot广播区块.
	//`
	desc_eng := `	Assume that the current epoch = 0, then in epoch = 1, the votes of all 
	malicious validators are not broadcast;
	In epoch = 2, the votes of all malicious validators are not broadcast;
	When the last malicious node in epoch = 2 produces a block, package the votes of
	all malicious validators before and broadcast the block at the last slot of epoch = 4.`
	return desc_eng
}

func (o *Instance) Run(ctx context.Context, params types.LibraryParams, feedbacker types.FeedBacker) {
	log.WithField("name", o.Name()).Info("start to run strategy")
	var latestEpoch int64 = -1
	ticker := time.NewTicker(time.Second * 3)
	slotTool := utils.SlotTool{SlotsPerEpoch: globalinfo.ChainBaseInfo().SlotsPerEpoch}
	for {
		select {
		case <-ctx.Done():
			log.WithField("name", o.Name()).Info("stop to run strategy")
			return
		case <-ticker.C:
			slot, err := utils.GetCurSlot(params.Attacker)
			if err != nil {
				log.WithField("error", err).Error("failed to get slot")
				continue
			}
			log.WithFields(log.Fields{
				"slot":      slot,
				"lastEpoch": latestEpoch,
			}).Info("get slot")
			epoch := slotTool.SlotToEpoch(int64(slot))
			// generate new strategy at the end of last epoch.
			if int64(slot) < slotTool.EpochEnd(epoch) {
				continue
			}
			if epoch == latestEpoch {
				continue
			}
			latestEpoch = epoch

			{
				nextEpoch := epoch + 1

				duties, err := utils.GetEpochDuties(params.Attacker, nextEpoch)
				if err != nil {
					log.WithFields(log.Fields{
						"error": err,
						"epoch": epoch + 1,
					}).Error("failed to get duties")
					latestEpoch = epoch - 1
					continue
				}
				strategy := types.Strategy{}
				strategy.Uid = uuid.NewString()
				strategy.Slots = GenSlotStrategy(params.FillterHackDuties(duties), nextEpoch)
				if err = utils.UpdateStrategy(params.Attacker, strategy); err != nil {
					log.WithField("error", err).Error("failed to update strategy")
				} else {
					log.WithFields(log.Fields{
						"epoch":    nextEpoch,
						"strategy": strategy,
					}).Info("update strategy successfully")
				}
			}
		}
	}
}
