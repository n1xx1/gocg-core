package main

import (
	"fmt"
	"io/ioutil"
	"ocgcore"
	"path/filepath"
)

func test() error {
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

	return nil
}

func main() {
	err := test()
	if err != nil {
		panic(err)
	}
}
