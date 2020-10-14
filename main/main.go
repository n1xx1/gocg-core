package main

import (
	"fmt"
	"io/ioutil"
	"ocgcore"
	"path/filepath"
)

func test() error {
	db := newCardDatabase()
	fmt.Println(len(db.cardList))

	path, err := filepath.Abs("ocgcore.dll")
	if err != nil {
		return err
	}
	core, err := ocgcore.NewOcgCore(path)
	if err != nil {
		return err
	}

	duelHandle := core.CreateDuel(ocgcore.CreateDuelOptions{
		Seed:  0,
		Flags: 0,
		CardReader: func(code uint32) ocgcore.RawCardData {
			fmt.Println("card reader:", code)
			return ocgcore.RawCardData{
				Code:       0,
				Alias:      0,
				SetCodes:   nil,
				CardType:   0,
				Level:      0,
				Attribute:  0,
				Race:       0,
				Attack:     0,
				Defense:    0,
				LScale:     0,
				RScale:     0,
				LinkMarker: 0,
			}
		},
		ScriptReader: func(path string) []byte {
			contents, err := ioutil.ReadFile(filepath.Join("script", path))
			if err != nil {
				fmt.Println(err)
				return nil
			}
			return contents
		},
	})

	constantFile, err := ioutil.ReadFile("script/constant.lua")
	core.LoadScript(duelHandle, "constant.lua", constantFile)
	utilityFile, err := ioutil.ReadFile("script/utility.lua")
	core.LoadScript(duelHandle, "utility.lua", utilityFile)

	for i, card := range myDeck {
		core.DuelNewCard(duelHandle, 0, card, 0, ocgcore.LocationDeck, i, 1)
	}
	for i, card := range myDeck {
		core.DuelNewCard(duelHandle, 1, card, 1, ocgcore.LocationDeck, i, 1)
	}

	core.StartDuel(duelHandle)

	xd := core.DuelProcess(duelHandle)
	fmt.Println(xd)

	return nil
}

var myDeck = []uint32{96005454, 55878038, 88774734, 67748760, 61901281, 26655293, 15381421, 99234526, 68464358, 5969957, 57143342, 80250185, 30227494, 30227494, 81035362, 45894482, 62957424, 99745551, 35272499, 35272499, 35272499, 81275020, 20758643, 61677004, 10802915, 10802915, 56410040, 56410040, 15981690, 15981690, 53932291, 43694650, 48686504, 48686504, 48686504, 19353570, 19353570, 19353570, 8972398, 1845204, 47325505, 47325505, 47325505, 54693926, 54693926, 54693926, 81439173, 99266988, 99266988, 99266988, 24224830, 24224830, 24224830, 31443476, 31443476, 31443476, 67723438, 36668118, 62265044, 61740673}

func main() {
	err := test()
	if err != nil {
		panic(err)
	}
}
