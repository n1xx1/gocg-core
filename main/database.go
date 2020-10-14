package main

import (
	"database/sql"
	_ "github.com/alicebob/sqlittle/driver"
	"ocgcore"
)

type cardDatabaseEntry struct {
	data    ocgcore.RawCardData
	name    string
	desc    string
	strings map[int]string
}

type cardDatabase struct {
	cards    map[uint32]*cardDatabaseEntry
	cardList []*cardDatabaseEntry
}

func newCardDatabase() cardDatabase {
	db, err := sql.Open("sqlittle", "./cards.cdb")
	if err != nil {
		panic(err)
	}

	datas, err := db.Query(`
SELECT
	t.id, t.ot, t.alias, t.setcode, t.type, t.atk, t.def, t.level, t.race, t.attribute, t.category,
	d.name, d.desc, d.str1, d.str2, d.str3, d.str4, d.str5, d.str6, d.str7, d.str8, d.str9, d.str10, d.str11, d.str12, d.str13, d.str14, d.str15, d.str16
FROM datas d
LEFT JOIN texts t ON d.id = t.id`)

	if err != nil {
		panic(err)
	}

	var cardList []*cardDatabaseEntry
	cards := map[uint32]*cardDatabaseEntry{}

	var dataId, dataOt, dataAlias, dataSetCode, dataType, dataLevel, dataRace, dataAttribute, dataCategory, dataAtk, dataDef uint32
	var name, desc string
	var str [16]string

	for datas.Next() {
		err = datas.Scan(
			&dataId, &dataOt, &dataAlias, &dataSetCode, &dataType, &dataAtk, &dataDef, &dataLevel, &dataRace, &dataAttribute, &dataCategory,
			&name, &desc, &str[0], &str[1], &str[2], &str[3], &str[4], &str[5], &str[6], &str[7], &str[8], &str[9], &str[10], &str[11], &str[12], &str[13], &str[14], &str[15],
		)
		if err != nil {
			panic(err)
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

		card.data.CardType = dataType
		card.data.Attack = int32(dataAtk)
		card.data.Defense = int32(dataDef)

		if (dataType & uint32(ocgcore.TypeLink)) != 0 {
			card.data.LinkMarker = dataDef
			card.data.Defense = 0
		}

		if dataLevel < 0 {
			card.data.Level = -(dataLevel & 0xff)
		} else {
			card.data.Level = dataLevel & 0xff
		}
		card.data.LScale = (dataLevel >> 16) & 0xff
		card.data.RScale = (dataLevel >> 24) & 0xff
		card.data.Race = dataRace
		card.data.Attribute = dataAttribute

		card.name = name
		card.desc = desc
		card.strings = map[int]string{}

		for i := 0; i < 16; i++ {
			if str[i] != "" {
				card.strings[i] = str[i]
			}
		}

		cardList = append(cardList, card)
		cards[card.data.Code] = card
	}

	return cardDatabase{
		cards:    cards,
		cardList: cardList,
	}
}
