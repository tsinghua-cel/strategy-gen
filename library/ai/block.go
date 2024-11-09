package aiattack

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tsinghua-cel/strategy-gen/ai"
	"github.com/tsinghua-cel/strategy-gen/types"
	"os"
	"strings"
	"sync"
)

var (
	prompt   = ""
	agent    *ai.Session
	initonce = sync.Once{}
)

func initAgent() {
	promptPath := os.Getenv("OPENAI_PROMPT_FILE")
	if data, err := os.ReadFile(promptPath); err != nil {
		log.WithError(err).Error("read prompt failed")
		return
	} else {
		prompt = string(data)
		agent = ai.NewAI().NewSession(context.TODO(), prompt)
		log.WithField("prompt", prompt).Info("read prompt success")
	}
}

func AddFeedBack(strategy types.Strategy, info types.FeedBackInfo) error {
	if agent == nil {
		return errors.New("agent is nil")
	}
	format := "I have a feedback for you, this feedback is about the strategy you provided me history. Detail strategy is `%s` and the feedback is `%s`."

	content, err := agent.Ask(fmt.Sprintf(format, strategy, info))
	if err != nil {
		log.WithError(err).Error("agent.Ask() failed")
		return err
	}
	log.WithField("response", content).Info("add feedback success")
	return nil
}

func GenSlotStrategy(allHacks []types.ProposerDuty) []types.SlotStrategy {
	if len(allHacks) == 0 {
		return nil
	}
	initonce.Do(initAgent)
	var strategies []types.SlotStrategy
	slots := ""
	for i, duty := range allHacks {
		if i == (len(allHacks) - 1) {
			slots = fmt.Sprintf("%s%s", slots, duty.Slot)
		} else {
			slots = fmt.Sprintf("%s%s,", slots, duty.Slot)
		}
	}
	content, err := agent.Ask(fmt.Sprintf("produce several penetration test plans for me and slot should between [%s].", slots))
	if err != nil {
		log.WithError(err).Error("agent.Ask() failed")
		return nil
	} else {
		content = strings.Replace(content, "```json", "", -1)
		content = strings.Replace(content, "```", "", -1)

		if err = json.Unmarshal([]byte(content), &strategies); err != nil {
			log.WithError(err).Error("json.Unmarshal() failed")
		} else {
			log.WithField("strategys", strategies).Info("generate new strategy success")
		}
	}

	return strategies

}
