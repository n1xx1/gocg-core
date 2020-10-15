package ocgcore

import (
	"bytes"
	"encoding/binary"
	"io"
)

func parseQuery(data []byte) coreQueryResult {
	b := bytes.NewBuffer(data)

	res := coreQueryResult{}
	for {
		var length uint16
		err := binary.Read(b, binary.LittleEndian, &length)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		query := readInt32(b)
		queryData := make([]byte, length-4)
		querySize, err := b.Read(queryData)

		if err != nil || querySize != int(length-4) {
			panic(err)
		}
		res[coreQuery(query)] = queryData
	}
	return res
}

func parseQueryLocation(data []byte) []coreQueryResult {
	b := bytes.NewBuffer(data)

	var res []coreQueryResult
	size := readUint32(b)
	if size == 0 {
		return nil
	}

	for b.Len() > 0 {
		cardRes := coreQueryResult{}
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

			query := readUint32(b)
			queryData := make([]byte, length-4)
			querySize, err := b.Read(queryData)

			if err != nil || querySize != int(length-4) {
				panic(err)
			}
			if query == uint32(coreQueryEnd) {
				break
			}
			cardRes[coreQuery(query)] = queryData
		}
		if len(cardRes) > 0 {
			res = append(res, cardRes)
		} else {
			res = append(res, nil)
		}
	}
	return res
}

func parseQueryField(data []byte) coreQueryField {
	b := bytes.NewBuffer(data)

	var field coreQueryField
	field.duelOptions = readInt32(b)
	parsePlayer(b, &field.player1)
	parsePlayer(b, &field.player2)
	chainSize := readInt32(b)
	for i := 0; i < int(chainSize); i++ {
		field.chain = append(field.chain, coreQueryFieldChain{
			code:                 readInt32(b),
			controller:           readUint8(b),
			location:             readUint8(b),
			sequence:             readUint32(b),
			position:             readUint32(b),
			triggeringController: readUint8(b),
			triggeringLocation:   readUint8(b),
			triggeringSequence:   readUint32(b),
			description:          readUint64(b),
		})
	}
	return field
}

func parsePlayer(b *bytes.Buffer, player *coreQueryFieldPlayer) {
	player.lp = readInt32(b)
	for i := 0; i < 7; i++ {
		player.monsters[i].present = readUint8(b) != 0
		if player.monsters[i].present {
			player.monsters[i].position = readInt8(b)
			player.monsters[i].materials = readInt32(b)
		}
	}
	for i := 0; i < 8; i++ {
		player.spells[i].present = readUint8(b) != 0
		if player.spells[i].present {
			player.spells[i].position = readInt8(b)
			player.spells[i].materials = readInt32(b)
		}
	}
	player.mainCount = readUint32(b)
	player.handCount = readUint32(b)
	player.graveCount = readUint32(b)
	player.banishCount = readUint32(b)
	player.extraCount = readUint32(b)
	player.extraPCount = readUint32(b)
}

type coreQueryResult map[coreQuery][]byte

type coreQueryField struct {
	duelOptions int32
	player1     coreQueryFieldPlayer
	player2     coreQueryFieldPlayer
	chain       []coreQueryFieldChain
}

type coreQueryFieldChain struct {
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

type coreQueryFieldPlayer struct {
	lp          int32
	monsters    [7]coreQueryFieldCard
	spells      [8]coreQueryFieldCard
	mainCount   uint32
	handCount   uint32
	graveCount  uint32
	banishCount uint32
	extraCount  uint32
	extraPCount uint32
}

type coreQueryFieldCard struct {
	present   bool
	position  int8
	materials int32
}
