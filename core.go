package ocgcore

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"ocgcore/lib"
)

type CardReader func(code uint32) RawCardData
type ScriptReader func(path string) []byte

type CreateDuelOptions struct {
	Seed uint32
	Mode DuelMode

	CardReader   CardReader
	ScriptReader ScriptReader
}

type RawCardData lib.CardData

type DuelMode int

const (
	DuelModeSpeed DuelMode = iota
	DuelModeRush
	DuelModeMR1
	DuelModeGoat
	DuelModeMR2
	DuelModeMR3
	DuelModeMR4
	DuelModeMR5
)

func CreateDuel(options CreateDuelOptions) *OcgDuel {
	var flags lib.DuelMode

	switch options.Mode {
	case DuelModeSpeed:
		flags |= lib.DuelModeSpeed
	case DuelModeRush:
		flags |= lib.DuelModeRush
	case DuelModeMR1:
		flags |= lib.DuelModeMR1
	case DuelModeGoat:
		flags |= lib.DuelModeGoat
	case DuelModeMR2:
		flags |= lib.DuelModeMR2
	case DuelModeMR3:
		flags |= lib.DuelModeMR3
	case DuelModeMR4:
		flags |= lib.DuelModeMR4
	case DuelModeMR5:
		flags |= lib.DuelModeMR5
	}

	duelOptions := lib.DuelOptions{
		Seed:  0,
		Flags: flags,
		Team1: lib.Player{
			StartingLP:        8000,
			StartingDrawCount: 5,
			DrawCountPerTurn:  1,
		},
		Team2: lib.Player{
			StartingLP:        8000,
			StartingDrawCount: 5,
			DrawCountPerTurn:  1,
		},
		CardReader: func(code uint32) (cardData lib.CardData) {
			return lib.CardData(options.CardReader(code))
		},
		ScriptReader: func(duel lib.Duel, name string) bool {
			contents := options.ScriptReader(name)
			if len(contents) == 0 {
				return false
			}
			lib.LoadScript(duel, contents, name)
			return true
		},
		LogHandler: func(message string, typ int) {
			fmt.Println("log handler: ", message, typ)
		},
		CardReaderDone: func(data lib.CardData) {},
	}

	duel := lib.CreateDuel(duelOptions)

	lib.LoadScript(duel, options.ScriptReader("constant.lua"), "constant.lua")
	lib.LoadScript(duel, options.ScriptReader("utility.lua"), "utility.lua")

	return newDuel(duel)
}

func duelGetMessage(duel lib.Duel) [][]byte {
	data := lib.DuelGetMessage(duel)
	dataBuffer := bytes.NewBuffer(data)

	var messages [][]byte
	for {
		var length uint32
		err := binary.Read(dataBuffer, binary.LittleEndian, &length)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		b := make([]byte, length)
		read, err := dataBuffer.Read(b)
		if err != nil || read != int(length) {
			panic(err)
		}
		messages = append(messages, b)
	}
	return messages
}

func FieldStatus(duel lib.Duel) (field Field) {
	loadFieldPlayer(duel, &field.Player1, 0)
	loadFieldPlayer(duel, &field.Player2, 1)
	return
}

func loadFieldPlayer(duel lib.Duel, player *FieldPlayer, con uint8) {
	flagsField := lib.QueryCode |
		lib.QueryLevel | lib.QueryPosition |
		lib.QueryAttack | lib.QueryDefense | lib.QueryEquipCard |
		lib.QueryCounters | lib.QueryLScale | lib.QueryRScale
	flagsDeck := lib.QueryCode | lib.QueryPosition

	player.Deck = parseFieldDeckCards(duelQueryLocation(duel, lib.QueryInfo{Flags: flagsDeck, Controller: con, Location: lib.LocationDeck}))
	player.ExtraDeck = parseFieldDeckCards(duelQueryLocation(duel, lib.QueryInfo{Flags: flagsDeck, Controller: con, Location: lib.LocationDeck}))
	player.Grave = parseFieldDeckCards(duelQueryLocation(duel, lib.QueryInfo{Flags: flagsDeck, Controller: con, Location: lib.LocationDeck}))
	player.Banished = parseFieldDeckCards(duelQueryLocation(duel, lib.QueryInfo{Flags: flagsDeck, Controller: con, Location: lib.LocationDeck}))
	player.Hand = parseFieldDeckCards(duelQueryLocation(duel, lib.QueryInfo{Flags: flagsDeck, Controller: con, Location: lib.LocationHand}))

	monsters := duelQueryLocation(duel, lib.QueryInfo{Flags: flagsField, Controller: con, Location: lib.LocationMZone})
	for i := 0; i < 5; i++ {
		if monsters[i] != nil {
			m := parseFieldCard(monsters[i])
			player.Monsters[i] = &m
		}
	}
	for i := 0; i < 2; i++ {
		if monsters[5+i] != nil {
			s := parseFieldCard(monsters[5+i])
			player.PendulumZones[i] = &s
		}
	}

	spells := duelQueryLocation(duel, lib.QueryInfo{Flags: flagsField, Controller: con, Location: lib.LocationSZone})
	for i := 0; i < 5; i++ {
		if spells[i] != nil {
			s := parseFieldCard(spells[i])
			player.Spells[i] = &s
		}
	}
	if spells[5] != nil {
		s := parseFieldCard(spells[5])
		player.FieldSpell = &s
	}
	for i := 0; i < 2; i++ {
		if spells[6+i] != nil {
			s := parseFieldCard(spells[6+i])
			player.PendulumZones[i] = &s
		}
	}
}

func parseFieldDeckCards(cards []lib.ParsedQueryResult) []FieldDeckCard {
	res := make([]FieldDeckCard, len(cards))
	for i, data := range cards {
		res[i] = parseFieldDeckCard(data)
	}
	return res
}

func parseFieldDeckCard(data lib.ParsedQueryResult) (card FieldDeckCard) {
	card.Code = binary.LittleEndian.Uint32(data[lib.QueryCode])
	card.Position = parseCorePosition(lib.Position(binary.LittleEndian.Uint32(data[lib.QueryPosition]))).Face()
	return
}

func parseFieldCard(data lib.ParsedQueryResult) (card FieldCard) {
	card.Code = binary.LittleEndian.Uint32(data[lib.QueryCode])
	card.Position = parseCorePosition(lib.Position(binary.LittleEndian.Uint32(data[lib.QueryPosition])))
	card.Level = int(binary.LittleEndian.Uint32(data[lib.QueryLevel]))
	card.Defense = int(binary.LittleEndian.Uint32(data[lib.QueryDefense]))
	card.Attack = int(binary.LittleEndian.Uint32(data[lib.QueryAttack]))
	card.LScale = int(binary.LittleEndian.Uint32(data[lib.QueryLScale]))
	card.RScale = int(binary.LittleEndian.Uint32(data[lib.QueryRScale]))
	return
}

type Field struct {
	Player1 FieldPlayer `json:"player1"`
	Player2 FieldPlayer `json:"player2"`
}

type FieldPlayer struct {
	Deck          []FieldDeckCard `json:"deck"`
	ExtraDeck     []FieldDeckCard `json:"extra_deck"`
	Hand          []FieldDeckCard `json:"hand"`
	Grave         []FieldDeckCard `json:"grave"`
	Banished      []FieldDeckCard `json:"banished"`
	Monsters      [5]*FieldCard   `json:"monsters"`
	Spells        [5]*FieldCard   `json:"spells"`
	FieldSpell    *FieldCard      `json:"field_spell"`
	PendulumZones [2]*FieldCard   `json:"pendulum_zones"`
	ExtraMonsters [2]*FieldCard   `json:"extra_monsters"`
}

type FieldCard struct {
	Code     uint32   `json:"code"`
	Position Position `json:"position"`
	Level    int      `json:"level,omitempty"`
	Attack   int      `json:"attack,omitempty"`
	Defense  int      `json:"defense,omitempty"`
	LScale   int      `json:"l_scale,omitempty"`
	RScale   int      `json:"r_scale,omitempty"`
}

type FieldDeckCard struct {
	Code     uint32       `json:"code"`
	Position FacePosition `json:"position"`
}

func duelQueryOverlay(duel lib.Duel, flags lib.Query, con uint8, loc lib.Location, seq uint32, overlaySeq uint32) lib.ParsedQueryResult {
	return duelQueryInfo(duel, lib.QueryInfo{
		Flags:           flags,
		Controller:      con,
		Location:        loc,
		Sequence:        seq,
		OverlaySequence: overlaySeq,
	})
}

func duelQuery(duel lib.Duel, flags lib.Query, controller uint8, location lib.Location, sequence uint32) lib.ParsedQueryResult {
	return duelQueryInfo(duel, lib.QueryInfo{
		Flags:      flags,
		Controller: controller,
		Location:   location,
		Sequence:   sequence,
	})
}

func duelQueryInfo(duel lib.Duel, info lib.QueryInfo) lib.ParsedQueryResult {
	return lib.ParseQuery(lib.DuelQuery(duel, info))
}

func duelQueryLocation(duel lib.Duel, info lib.QueryInfo) []lib.ParsedQueryResult {
	return lib.ParseQueryLocation(lib.DuelQueryLocation(duel, info))
}

func duelQueryField(duel lib.Duel) lib.ParsedQueryField {
	return lib.ParseQueryField(lib.DuelQueryField(duel))
}
