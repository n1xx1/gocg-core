package database

import (
	"database/sql"
	"fmt"
	"ocgcore"
	"ocgcore/lib"

	_ "github.com/mattn/go-sqlite3"
)

type CardMonsterPendulum struct {
	Description string `json:"description"`
	LScale      int    `json:"l_scale"`
	RScale      int    `json:"r_scale"`
}

type CardMonster struct {
	Attribute ocgcore.CardMonsterAttribute `json:"attribute"`
	Ability   ocgcore.CardMonsterAbility   `json:"ability"`
	Frame     ocgcore.CardMonsterFrame     `json:"frame"`
	Type      ocgcore.CardMonsterType      `json:"type"`
	Tuner     bool                         `json:"tuner"`
	Attack    int                          `json:"attack"`
	Defense   int                          `json:"defense"`

	Pendulum *CardMonsterPendulum     `json:"pendulum,omitempty"`
	Link     []ocgcore.CardLinkMarker `json:"link,omitempty"`
}

type CardSpell struct {
	Type ocgcore.CardSpellType `json:"type"`
}

type CardTrap struct {
	Type ocgcore.CardTrapType `json:"type"`
}

type Card struct {
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Type        ocgcore.CardType `json:"type"`
	Monster     *CardMonster     `json:"monster,omitempty"`
	Spell       *CardSpell       `json:"spell,omitempty"`
	Trap        *CardTrap        `json:"trap,omitempty"`
	Token       *CardMonster     `json:"token,omitempty"`
}

type CardEntry struct {
	Card        Card
	Raw         ocgcore.RawCardData
	Name        string
	Description string

	stringIndexes [16]int
	strings       []string
}

func (c *CardEntry) Str(i int) string {
	return c.strings[c.stringIndexes[i]]
}

type CardDatabase map[uint32]*CardEntry

func NewDatabase() CardDatabase {
	return CardDatabase{}
}

type sqliteDatabaseSelect struct {
	dataId        uint32
	dataOt        uint32
	dataAlias     uint32
	dataType      uint32
	dataLevel     uint32
	dataRace      uint32
	dataAttribute uint32
	dataCategory  uint32
	dataSetCode   uint64
	dataAtk       int32
	dataDef       int32
	name          string
	desc          string
	str           [16]string
}

func (s *sqliteDatabaseSelect) toCard(raw *ocgcore.RawCardData, card *Card) {
	card.Name = s.name
	card.Description = s.desc

	card.Type = ocgcore.ParseCardType(raw.Type)

	switch card.Type {
	case ocgcore.CardTypeMonster:
		card.Monster = new(CardMonster)

		var pendulum bool
		card.Monster.Frame, card.Monster.Type, card.Monster.Ability, card.Monster.Tuner, pendulum =
			ocgcore.ParseCardTypeMonster(raw.Type)

		if pendulum {
			card.Monster.Pendulum = &CardMonsterPendulum{
				Description: s.desc,
				LScale:      int(raw.LScale),
				RScale:      int(raw.RScale),
			}
		}
		if card.Monster.Frame == ocgcore.CardMonsterFrameLink {
			// TODO: parse link markers
			card.Monster.Link = ocgcore.ParseLinkMarkers(raw.LinkMarker)
		}
	case ocgcore.CardTypeSpell:
		card.Spell = &CardSpell{
			Type: ocgcore.ParseCardTypeSpell(raw.Type),
		}
	case ocgcore.CardTypeTrap:
		card.Trap = &CardTrap{
			Type: ocgcore.ParseCardTypeTrap(raw.Type),
		}
	case ocgcore.CardTypeToken:
		card.Token = new(CardMonster)
		card.Monster.Frame, card.Monster.Type, card.Monster.Ability, card.Monster.Tuner, _ =
			ocgcore.ParseCardTypeMonster(raw.Type)
	}
}

func (s *sqliteDatabaseSelect) toRaw(card *ocgcore.RawCardData) {
	card.Code = s.dataId
	card.Alias = s.dataAlias

	card.SetCodes = nil
	for i := 0; i < 4; i++ {
		setCode := uint16((s.dataSetCode >> (i * 16)) & 0xffff)
		if setCode != 0 {
			card.SetCodes = append(card.SetCodes, setCode)
		}
	}

	card.Type = lib.CardType(s.dataType)
	card.Attack = s.dataAtk
	card.Defense = s.dataDef

	if (card.Type & lib.CardTypeLink) != 0 {
		card.LinkMarker = lib.LinkMarker(s.dataDef)
		card.Defense = 0
	}

	if s.dataLevel < 0 {
		card.Level = -(s.dataLevel & 0xff)
	} else {
		card.Level = s.dataLevel & 0xff
	}
	card.LScale = (s.dataLevel >> 16) & 0xff
	card.RScale = (s.dataLevel >> 24) & 0xff
	card.Race = lib.Race(s.dataRace)
	card.Attribute = lib.Attribute(s.dataAttribute)
	return
}

func (c CardDatabase) Load(fileName string) error {
	db, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?mode=ro", fileName))
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

	var c1 sqliteDatabaseSelect

	for datas.Next() {
		err = datas.Scan(
			&c1.dataId, &c1.dataOt, &c1.dataAlias, &c1.dataSetCode, &c1.dataType, &c1.dataAtk, &c1.dataDef, &c1.dataLevel, &c1.dataRace, &c1.dataAttribute, &c1.dataCategory,
			&c1.name, &c1.desc, &c1.str[0], &c1.str[1], &c1.str[2], &c1.str[3], &c1.str[4], &c1.str[5], &c1.str[6], &c1.str[7], &c1.str[8], &c1.str[9], &c1.str[10], &c1.str[11], &c1.str[12], &c1.str[13], &c1.str[14], &c1.str[15],
		)
		if err != nil {
			return err
		}

		card := &CardEntry{}

		c1.toRaw(&card.Raw)
		c1.toCard(&card.Raw, &card.Card)

		card.Name = c1.name
		card.Description = c1.desc

		for i := 0; i < 16; i++ {
			if c1.str[i] != "" {
				card.stringIndexes[i] = len(card.strings)
				card.strings = append(card.strings, c1.str[i])
			} else {
				card.stringIndexes[i] = 0
			}
		}

		c[c1.dataId] = card
	}
	return nil

}
