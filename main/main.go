package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"ocgcore"
	"path/filepath"
	"regexp"
	"time"
)

func scriptReader() func(path string) []byte {
	cardScriptRegex := regexp.MustCompile(`c\d+\.lua`)

	return func(path string) []byte {
		if cardScriptRegex.MatchString(path) {
			path = filepath.Join("official", path)
		}

		contents, err := ioutil.ReadFile(filepath.Join("script", path))
		if err != nil {
			fmt.Println(err)
			return nil
		}
		return contents
	}
}

func cardReader() func(code uint32) ocgcore.RawCardData {
	db := newCardDatabase()

	return func(code uint32) ocgcore.RawCardData {
		card := db[code]
		if card == nil {
			fmt.Println("card not found:", code)
			return ocgcore.RawCardData{}
		}
		// fmt.Println("card reader:", card.name)
		return card.data
	}
}

func printJson(v interface{}) {
	s, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(s))
}

func test() error {
	path, err := filepath.Abs("ocgcore.dll")
	if err != nil {
		return err
	}
	core, err := ocgcore.NewOcgCore(path)
	if err != nil {
		return err
	}

	rand.Seed(time.Now().Unix())

	duel := core.CreateDuel(ocgcore.CreateDuelOptions{
		Seed:         rand.Uint32(),
		Mode:         ocgcore.DuelModeMR5,
		CardReader:   cardReader(),
		ScriptReader: scriptReader(),
	})

	duel.SetupDeck(0, mainDeck, extraDeck, false)
	duel.SetupDeck(1, mainDeck, extraDeck, false)

	messages := duel.Start()

	step := 0
	lastIsEmptyChain := false
	graffLocation := 0
	for {
		m1, ok := <-messages
		if !ok {
			break
		}
		switch m := m1.(type) {
		case ocgcore.MessageSelectChain:
			if len(m.Chains) == 0 {
				lastIsEmptyChain = true
				continue
			}
		case ocgcore.MessageSelectCard:
			for i, c := range m.Cards {
				if c.Code == 20758643 {
					graffLocation = i
				}
			}
		case ocgcore.MessageWaitingResponse:
			if lastIsEmptyChain {
				duel.SendResponse(ocgcore.ResponseSelectChain{Chain: -1})
				break
			}

			switch step {
			case 0, 1, 4:
				duel.SendResponse(ocgcore.ResponseSelectChain{Chain: -1})
			case 2:
				duel.SendResponse(ocgcore.ResponseSelectIdleCMD{Action: ocgcore.IdleActionSummon, Index: 0})
			case 3:
				duel.SendResponse(ocgcore.ResponseSelectPlace{Places: []ocgcore.Place{{Player: 0, Location: ocgcore.LocationMonsterZone, Sequence: 2}}})
			case 5:
				duel.SendResponse(ocgcore.ResponseSelectIdleCMD{Action: ocgcore.IdleActionActivate, Index: 2})
			case 6:
				duel.SendResponse(ocgcore.ResponseSelectPlace{Places: []ocgcore.Place{{Player: 0, Location: ocgcore.LocationSpellZone, Sequence: 2}}})
			case 7:
				duel.SendResponse(ocgcore.ResponseSelectCard{Select: []int{0}})
			case 8:
				duel.SendResponse(ocgcore.ResponseSelectPlace{Places: []ocgcore.Place{{Player: 0, Location: ocgcore.LocationMonsterZone, Sequence: 1}}})
			case 9:
				duel.SendResponse(ocgcore.ResponseSelectPosition{Position: ocgcore.PositionFaceUpDefense})
			case 10:
				duel.SendResponse(ocgcore.ResponseSelectIdleCMD{Action: ocgcore.IdleActionSpSummon, Index: 0})
			case 11:
				duel.SendResponse(ocgcore.ResponseSelectUnselectCard{Selection: 0})
			case 12:
				duel.SendResponse(ocgcore.ResponseSelectUnselectCard{Selection: 0})
			case 13:
				duel.SendResponse(ocgcore.ResponseSelectPlace{Places: []ocgcore.Place{{Player: 0, Location: ocgcore.LocationMonsterZone, Sequence: 5}}})
			case 14:
				duel.SendResponse(ocgcore.ResponseSelectIdleCMD{Action: ocgcore.IdleActionActivate, Index: 2})
			case 15:
				duel.SendResponse(ocgcore.ResponseSelectCard{Select: []int{graffLocation}})
			case 16:
				duel.SendResponse(ocgcore.ResponseSelectCard{Select: []int{0}})
			case 17:
				duel.SendResponse(ocgcore.ResponseSelectEffectYN{Yes: true})
			case 18:
				duel.SendResponse(ocgcore.ResponseSelectCard{Select: []int{1}})
			case 19:
				duel.SendResponse(ocgcore.ResponseSelectPlace{Places: []ocgcore.Place{{Player: 0, Location: ocgcore.LocationMonsterZone, Sequence: 4}}})
			}
			step++
		}
		out, err := ocgcore.MessageToJSON(m1)
		if err != nil {
			panic(err)
		}

		msg, err := ocgcore.JSONToMessage(out)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%#v\n", msg)
		lastIsEmptyChain = false
	}
	return nil
}

var mainDeck = []uint32{55878038, 88774734, 67748760, 61901281, 26655293, 15381421, 99234526, 68464358, 5969957, 57143342, 80250185, 30227494, 30227494, 81035362, 45894482, 62957424, 99745551, 35272499, 35272499, 35272499, 81275020, 20758643, 61677004, 10802915, 10802915, 56410040, 56410040, 15981690, 15981690, 53932291, 43694650, 48686504, 48686504, 48686504, 19353570, 19353570, 19353570, 8972398, 1845204, 47325505, 47325505, 47325505, 54693926, 54693926, 54693926, 81439173, 99266988, 99266988, 99266988, 24224830, 24224830, 24224830, 31443476, 31443476, 31443476, 67723438, 36668118, 62265044, 96005454, 61740673}
var extraDeck = []uint32{17881964, 27548199, 63767246, 4280258, 85289965, 58699500, 98095162, 23935886, 11969228, 86148577, 13143275, 65330383, 2857636, 38342335, 73539069}

func main() {
	err := test()
	if err != nil {
		panic(err)
	}
}
