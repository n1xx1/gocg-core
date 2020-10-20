package lib

import (
	"bytes"
	"encoding/binary"
	"io"
	"ocgcore/utils"
)

func ParseQuery(data []byte) ParsedQueryResult {
	b := bytes.NewBuffer(data)

	res := ParsedQueryResult{}
	for {
		var length uint16
		err := binary.Read(b, binary.LittleEndian, &length)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		query := utils.ReadInt32(b)
		queryData := make([]byte, length-4)
		querySize, err := b.Read(queryData)

		if err != nil || querySize != int(length-4) {
			panic(err)
		}
		res[Query(query)] = queryData
	}
	return res
}

func ParseQueryLocation(data []byte) []ParsedQueryResult {
	b := bytes.NewBuffer(data)

	var res []ParsedQueryResult
	size := utils.ReadUint32(b)
	if size == 0 {
		return nil
	}

	for b.Len() > 0 {
		cardRes := ParsedQueryResult{}
		for {
			var length uint16
			err := binary.Read(b, binary.LittleEndian, &length)
			if err == io.EOF {
				break
			} else if err != nil {
				panic(err)
			}
			if length == 0 {
				break
			}

			query := utils.ReadUint32(b)
			queryData := make([]byte, length-4)
			querySize, err := b.Read(queryData)

			if err != nil || querySize != int(length-4) {
				panic(err)
			}
			if query == uint32(QueryEnd) {
				break
			}
			cardRes[Query(query)] = queryData
		}
		if len(cardRes) > 0 {
			res = append(res, cardRes)
		} else {
			res = append(res, nil)
		}
	}
	return res
}

func ParseQueryField(data []byte) ParsedQueryField {
	b := bytes.NewBuffer(data)

	var field ParsedQueryField
	field.duelOptions = utils.ReadInt32(b)
	parsePlayer(b, &field.player1)
	parsePlayer(b, &field.player2)
	chainSize := utils.ReadInt32(b)
	for i := 0; i < int(chainSize); i++ {
		field.chain = append(field.chain, ParsedQueryFieldChain{
			code:                 utils.ReadInt32(b),
			controller:           utils.ReadUint8(b),
			location:             utils.ReadUint8(b),
			sequence:             utils.ReadUint32(b),
			position:             utils.ReadUint32(b),
			triggeringController: utils.ReadUint8(b),
			triggeringLocation:   utils.ReadUint8(b),
			triggeringSequence:   utils.ReadUint32(b),
			description:          utils.ReadUint64(b),
		})
	}
	return field
}

func parsePlayer(b *bytes.Buffer, player *ParsedQueryFieldPlayer) {
	player.lp = utils.ReadInt32(b)
	for i := 0; i < 7; i++ {
		player.monsters[i].present = utils.ReadUint8(b) != 0
		if player.monsters[i].present {
			player.monsters[i].position = utils.ReadInt8(b)
			player.monsters[i].materials = utils.ReadInt32(b)
		}
	}
	for i := 0; i < 8; i++ {
		player.spells[i].present = utils.ReadUint8(b) != 0
		if player.spells[i].present {
			player.spells[i].position = utils.ReadInt8(b)
			player.spells[i].materials = utils.ReadInt32(b)
		}
	}
	player.mainCount = utils.ReadUint32(b)
	player.handCount = utils.ReadUint32(b)
	player.graveCount = utils.ReadUint32(b)
	player.banishCount = utils.ReadUint32(b)
	player.extraCount = utils.ReadUint32(b)
	player.extraPCount = utils.ReadUint32(b)
}

type ParsedQueryResult map[Query][]byte

type ParsedQueryField struct {
	duelOptions int32
	player1     ParsedQueryFieldPlayer
	player2     ParsedQueryFieldPlayer
	chain       []ParsedQueryFieldChain
}

type ParsedQueryFieldChain struct {
	code                 int32
	controller           uint8
	location             uint8
	sequence             uint32
	position             uint32
	triggeringController uint8
	triggeringLocation   uint8
	triggeringSequence   uint32
	description          uint64
}

type ParsedQueryFieldPlayer struct {
	lp          int32
	monsters    [7]ParsedQueryFieldCard
	spells      [8]ParsedQueryFieldCard
	mainCount   uint32
	handCount   uint32
	graveCount  uint32
	banishCount uint32
	extraCount  uint32
	extraPCount uint32
}

type ParsedQueryFieldCard struct {
	present   bool
	position  int8
	materials int32
}
