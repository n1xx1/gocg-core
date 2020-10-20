package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"ocgcore"
	"ocgcore/lib"
)

type cardDatabaseEntry struct {
	data    ocgcore.RawCardData
	name    string
	desc    string
	strings map[int]string
}

type cardDatabase map[uint32]*cardDatabaseEntry

const cardTypeLink uint32 = 0x4000000

func loadDatabase(fileName string, cards cardDatabase) error {
	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		return err
	}

	datas, err := db.Query(`
SELECT
	d.id, d.ot, d.alias, d.setcode, d.type, d.atk, d.def, d.level, d.race, d.attribute, d.category,
	t.name, t.desc, t.str1, t.str2, t.str3, t.str4, t.str5, t.str6, t.str7, t.str8, t.str9, t.str10, t.str11, t.str12, t.str13, t.str14, t.str15, t.str16
FROM datas d
LEFT JOIN texts t ON d.id = t.id`)

	if err != nil {
		return err
	}

	var dataId, dataOt, dataAlias, dataType, dataLevel, dataRace, dataAttribute, dataCategory uint32
	var dataSetCode uint64
	var dataAtk, dataDef int32
	var name, desc string
	var str [16]string

	for datas.Next() {
		err = datas.Scan(
			&dataId, &dataOt, &dataAlias, &dataSetCode, &dataType, &dataAtk, &dataDef, &dataLevel, &dataRace, &dataAttribute, &dataCategory,
			&fileName, &desc, &str[0], &str[1], &str[2], &str[3], &str[4], &str[5], &str[6], &str[7], &str[8], &str[9], &str[10], &str[11], &str[12], &str[13], &str[14], &str[15],
		)
		if err != nil {
			return err
		}

		card := &cardDatabaseEntry{}

		card.data.Code = dataId
		card.data.Alias = dataAlias

		card.data.SetCodes = nil
		for i := 0; i < 4; i++ {
			setCode := uint16((dataSetCode >> (i * 16)) & 0xffff)
			if setCode != 0 {
				card.data.SetCodes = append(card.data.SetCodes, setCode)
			}
		}

		card.data.Type = lib.CardType(dataType)
		card.data.Attack = dataAtk
		card.data.Defense = dataDef

		if (dataType & cardTypeLink) != 0 {
			card.data.LinkMarker = lib.LinkMarker(dataDef)
			card.data.Defense = 0
		}

		if dataLevel < 0 {
			card.data.Level = -(dataLevel & 0xff)
		} else {
			card.data.Level = dataLevel & 0xff
		}
		card.data.LScale = (dataLevel >> 16) & 0xff
		card.data.RScale = (dataLevel >> 24) & 0xff
		card.data.Race = lib.Race(dataRace)
		card.data.Attribute = lib.Attribute(dataAttribute)

		card.name = name
		card.desc = desc
		card.strings = map[int]string{}

		for i := 0; i < 16; i++ {
			if str[i] != "" {
				card.strings[i] = str[i]
			}
		}
		cards[card.data.Code] = card
	}
	return nil
}

func newCardDatabase() cardDatabase {
	cards := cardDatabase{}
	err := loadDatabase("cards.cdb", cards)
	if err != nil {
		panic(err)
	}
	err = loadDatabase("release.cdb", cards)
	if err != nil {
		panic(err)
	}
	return cards
}
