package main

import (
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

	duelHandle := core.CreateDuel(ocgcore.CreateDuelOptions{
		Seed:         rand.Uint32(),
		Flags:        ocgcore.CoreDuelModeMR5,
		CardReader:   cardReader(),
		ScriptReader: scriptReader(),
	})

	constantFile, err := ioutil.ReadFile("script/constant.lua")
	core.LoadScript(duelHandle, "constant.lua", constantFile)
	utilityFile, err := ioutil.ReadFile("script/utility.lua")
	core.LoadScript(duelHandle, "utility.lua", utilityFile)

	core.SetupDeck(duelHandle, 0, myDeck, []uint32{}, true)
	core.SetupDeck(duelHandle, 1, myDeck, []uint32{}, true)

	messages := core.DuelGetMessage(duelHandle)
	for _, message := range messages {
		msg := ocgcore.ReadMessage(message)
		fmt.Printf("%#v", msg)
	}

	core.StartDuel(duelHandle)

	for {
		status := core.DuelProcess(duelHandle)
		messages = core.DuelGetMessage(duelHandle)
		for _, message := range messages {
			msg := ocgcore.ReadMessage(message)
			if msg != nil {
				fmt.Printf("%#v\n", msg)
			}
		}

		if status != ocgcore.ProcessorFlagContinue {
			fmt.Println("status:", status)
			break
		}
	}

	core.Debug(duelHandle)

	return nil
}

var myDeck = []uint32{96005454, 55878038, 88774734, 67748760, 61901281, 26655293, 15381421, 99234526, 68464358, 5969957, 57143342, 80250185, 30227494, 30227494, 81035362, 45894482, 62957424, 99745551, 35272499, 35272499, 35272499, 81275020, 20758643, 61677004, 10802915, 10802915, 56410040, 56410040, 15981690, 15981690, 53932291, 43694650, 48686504, 48686504, 48686504, 19353570, 19353570, 19353570, 8972398, 1845204, 47325505, 47325505, 47325505, 54693926, 54693926, 54693926, 81439173, 99266988, 99266988, 99266988, 24224830, 24224830, 24224830, 31443476, 31443476, 31443476, 67723438, 36668118, 62265044, 61740673}

func main() {
	err := test()
	if err != nil {
		panic(err)
	}
}
