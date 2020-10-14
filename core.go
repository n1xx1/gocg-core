package ocgcore

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"golang.org/x/sys/windows"
	"io"
	"math/rand"
	"unsafe"
)

type DuelHandle uintptr

type OcgCore struct {
	dll *windows.DLL

	// OCGAPI void OCG_GetVersion(int* major, int* minor);
	procGetVersion *windows.Proc
	// OCGAPI int OCG_CreateDuel(OCG_Duel* duel, OCG_DuelOptions options);
	procCreateDuel *windows.Proc
	// OCGAPI void OCG_DestroyDuel(OCG_Duel duel);
	procDestroyDuel *windows.Proc
	// OCGAPI void OCG_DuelNewCard(OCG_Duel duel, OCG_NewCardInfo info);
	procDuelNewCard *windows.Proc
	// OCGAPI void OCG_StartDuel(OCG_Duel duel);
	procStartDuel *windows.Proc
	// OCGAPI int OCG_DuelProcess(OCG_Duel duel);
	procDuelProcess *windows.Proc
	// OCGAPI void* OCG_DuelGetMessage(OCG_Duel duel, uint32_t* length);
	procDuelGetMessage *windows.Proc
	// OCGAPI void OCG_DuelSetResponse(OCG_Duel duel, const void* buffer, uint32_t length);
	procDuelSetResponse *windows.Proc
	// OCGAPI int OCG_LoadScript(OCG_Duel duel, const char* buffer, uint32_t length, const char* name);
	procLoadScript *windows.Proc
	// OCGAPI uint32_t OCG_DuelQueryCount(OCG_Duel duel, uint8_t team, uint32_t loc);
	procDuelQueryCount *windows.Proc
	// OCGAPI void* OCG_DuelQuery(OCG_Duel duel, uint32_t* length, OCG_QueryInfo info);
	procDuelQuery *windows.Proc
	// OCGAPI void* OCG_DuelQueryLocation(OCG_Duel duel, uint32_t* length, OCG_QueryInfo info);
	procDuelQueryLocation *windows.Proc
	// OCGAPI void* OCG_DuelQueryField(OCG_Duel duel, uint32_t* length);
	procDuelQueryField *windows.Proc
}

func NewOcgCore(path string) (*OcgCore, error) {
	dll, err := windows.LoadDLL(path)
	if err != nil {
		return nil, err
	}

	// OCGAPI void OCG_GetVersion(int* major, int* minor);
	procGetVersion, err := dll.FindProc("OCG_GetVersion")
	if err != nil {
		return nil, err
	}
	// OCGAPI int OCG_CreateDuel(OCG_Duel* duel, OCG_DuelOptions options);
	procCreateDuel, err := dll.FindProc("OCG_CreateDuel")
	if err != nil {
		return nil, err
	}
	// OCGAPI void OCG_DestroyDuel(OCG_Duel duel);
	procDestroyDuel, err := dll.FindProc("OCG_DestroyDuel")
	if err != nil {
		return nil, err
	}
	// OCGAPI void OCG_DuelNewCard(OCG_Duel duel, OCG_NewCardInfo info);
	procDuelNewCard, err := dll.FindProc("OCG_DuelNewCard")
	if err != nil {
		return nil, err
	}
	// OCGAPI void OCG_StartDuel(OCG_Duel duel);
	procStartDuel, err := dll.FindProc("OCG_StartDuel")
	if err != nil {
		return nil, err
	}
	// OCGAPI int OCG_DuelProcess(OCG_Duel duel);
	procDuelProcess, err := dll.FindProc("OCG_DuelProcess")
	if err != nil {
		return nil, err
	}
	// OCGAPI void* OCG_DuelGetMessage(OCG_Duel duel, uint32_t* length);
	procDuelGetMessage, err := dll.FindProc("OCG_DuelGetMessage")
	if err != nil {
		return nil, err
	}
	// OCGAPI void OCG_DuelSetResponse(OCG_Duel duel, const void* buffer, uint32_t length);
	procDuelSetResponse, err := dll.FindProc("OCG_DuelSetResponse")
	if err != nil {
		return nil, err
	}
	// OCGAPI int OCG_LoadScript(OCG_Duel duel, const char* buffer, uint32_t length, const char* name);
	procLoadScript, err := dll.FindProc("OCG_LoadScript")
	if err != nil {
		return nil, err
	}
	// OCGAPI uint32_t OCG_DuelQueryCount(OCG_Duel duel, uint8_t team, uint32_t loc);
	procDuelQueryCount, err := dll.FindProc("OCG_DuelQueryCount")
	if err != nil {
		return nil, err
	}
	// OCGAPI void* OCG_DuelQuery(OCG_Duel duel, uint32_t* length, OCG_QueryInfo info);
	procDuelQuery, err := dll.FindProc("OCG_DuelQuery")
	if err != nil {
		return nil, err
	}
	// OCGAPI void* OCG_DuelQueryLocation(OCG_Duel duel, uint32_t* length, OCG_QueryInfo info);
	procDuelQueryLocation, err := dll.FindProc("OCG_DuelQueryLocation")
	if err != nil {
		return nil, err
	}
	// OCGAPI void* OCG_DuelQueryField(OCG_Duel duel, uint32_t* length);
	procDuelQueryField, err := dll.FindProc("OCG_DuelQueryField")
	if err != nil {
		return nil, err
	}

	core := &OcgCore{
		dll:                   dll,
		procGetVersion:        procGetVersion,
		procCreateDuel:        procCreateDuel,
		procDestroyDuel:       procDestroyDuel,
		procDuelNewCard:       procDuelNewCard,
		procStartDuel:         procStartDuel,
		procDuelProcess:       procDuelProcess,
		procDuelGetMessage:    procDuelGetMessage,
		procDuelSetResponse:   procDuelSetResponse,
		procLoadScript:        procLoadScript,
		procDuelQueryCount:    procDuelQueryCount,
		procDuelQuery:         procDuelQuery,
		procDuelQueryLocation: procDuelQueryLocation,
		procDuelQueryField:    procDuelQueryField,
	}
	return core, nil
}

func bytesFromPtr(cptr uintptr, size uintptr) []byte {
	cbytes := (*[1<<30 - 1]byte)(unsafe.Pointer(cptr))
	cbytesCopy := make([]byte, size)
	copy(cbytesCopy, cbytes[:size])
	return cbytesCopy
}

func stringFromCStringPtr(cptr uintptr) string {
	cbytes := (*[1<<30 - 1]byte)(unsafe.Pointer(cptr))
	return stringFromCString(cbytes[:])
}

func stringFromCString(cbytes []byte) string {
	size := bytes.IndexByte(cbytes, 0)
	cbytesCopy := make([]byte, size)
	copy(cbytesCopy, cbytes[:size])
	return string(cbytesCopy)
}

type CreateDuelOptions struct {
	Seed         uint32
	Flags        CoreDuelMode
	CardReader   func(code uint32) RawCardData
	ScriptReader func(path string) []byte
}

func (c *OcgCore) CreateDuel(options CreateDuelOptions) DuelHandle {
	loadedCards := map[uint32]RawCardData{}

	duelHandle := uintptr(0)
	duelOptions := typeDuelOptions{
		seed:  0,
		flags: uint32(options.Flags),
		team1: typePlayer{
			startingLP:        8000,
			startingDrawCount: 5,
			drawCountPerTurn:  1,
		},
		team2: typePlayer{
			startingLP:        8000,
			startingDrawCount: 5,
			drawCountPerTurn:  1,
		},
		cardReader: windows.NewCallback(func(payload uintptr, code uint32, data uintptr) uintptr {
			cardData := (*typeCardData)(unsafe.Pointer(data))
			_ = cardData

			loadedCard, ok := loadedCards[code]
			if !ok {
				loadedCard = options.CardReader(code)
				setCodes := make([]uint16, 0, len(loadedCard.SetCodes)+1)
				loadedCard.SetCodes = append(setCodes, loadedCard.SetCodes...)
				loadedCard.SetCodes = append(loadedCard.SetCodes, 0)
				loadedCards[code] = loadedCard
			}
			cardData.code = loadedCard.Code
			cardData.alias = loadedCard.Alias
			cardData.setcodes = uintptr(unsafe.Pointer(&loadedCard.SetCodes[0]))
			cardData.cardtype = loadedCard.CardType
			cardData.level = loadedCard.Level
			cardData.attribute = loadedCard.Attribute
			cardData.race = loadedCard.Race
			cardData.attack = loadedCard.Attack
			cardData.defense = loadedCard.Defense
			cardData.lscale = loadedCard.LScale
			cardData.rscale = loadedCard.RScale
			cardData.link_marker = loadedCard.LinkMarker

			return 0
		}),
		scriptReader: windows.NewCallback(func(payload uintptr, duel DuelHandle, namePtr uintptr) uintptr {
			name := stringFromCStringPtr(namePtr)
			contents := options.ScriptReader(name)
			if len(contents) == 0 {
				return 0
			}
			c.LoadScript(duel, name, contents)
			return 1
		}),
		logHandler: windows.NewCallback(func(payload uintptr, messagePtr uintptr, messageType uintptr) uintptr {
			message := stringFromCStringPtr(messagePtr)
			fmt.Println("log handler: ", message)
			return 0
		}),
		cardReaderDone: windows.NewCallback(func(payload uintptr, data uintptr) uintptr {
			// fmt.Println("card reader done")
			return 0
		}),
	}

	_, _, err := c.procCreateDuel.Call(uintptr(unsafe.Pointer(&duelHandle)), uintptr(unsafe.Pointer(&duelOptions)))
	if err != nil && err != windows.ERROR_SUCCESS {
		panic(err)
	}
	return DuelHandle(duelHandle)
}

func (c *OcgCore) LoadScript(duel DuelHandle, name string, content []byte) {
	contentBytes := make([]byte, len(content))
	copy(contentBytes, content)
	contentPtr := uintptr(unsafe.Pointer(&contentBytes[0]))

	nameBytes := make([]byte, len(name))
	copy(nameBytes, name)
	namePtr := uintptr(unsafe.Pointer(&nameBytes[0]))

	_, _, err := c.procLoadScript.Call(uintptr(duel), contentPtr, uintptr(len(contentBytes)), namePtr, uintptr(len(nameBytes)))
	if err != nil && err != windows.ERROR_SUCCESS {
		panic(err)
	}
}

func (c *OcgCore) duelNewCard(duel DuelHandle, cardInfo typeNewCardInfo) {
	_, _, err := c.procDuelNewCard.Call(uintptr(duel), uintptr(unsafe.Pointer(&cardInfo)))
	if err != nil && err != windows.ERROR_SUCCESS {
		panic(err)
	}
}

func (c *OcgCore) DuelProcess(duel DuelHandle) ProcessorFlag {
	ret, _, err := c.procDuelProcess.Call(uintptr(duel))
	if err != nil && err != windows.ERROR_SUCCESS {
		panic(err)
	}
	return ProcessorFlag(ret)
}

func (c *OcgCore) StartDuel(duel DuelHandle) {
	_, _, err := c.procStartDuel.Call(uintptr(duel))
	if err != nil && err != windows.ERROR_SUCCESS {
		panic(err)
	}
}

func (c *OcgCore) DuelGetMessage(duel DuelHandle) [][]byte {
	var length uint32
	dataPtr, _, err := c.procDuelGetMessage.Call(uintptr(duel), uintptr(unsafe.Pointer(&length)))
	if err != nil && err != windows.ERROR_SUCCESS {
		panic(err)
	}

	if length == 0 {
		return nil
	}

	data := bytesFromPtr(dataPtr, uintptr(length))
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
func (c *OcgCore) Debug(duel DuelHandle) {
	fieldCounts := c.parseQueryField(c.duelQueryField(duel))

	var field Field

	for i := 0; i < 5; i++ {
		if fieldCounts.player1.monsters[i].present {
			m := c.getFieldCard(duel, 0, coreLocationMZone, uint32(i))
			field.Player1.Monsters[i] = &m
		}
		if fieldCounts.player1.spells[i].present {
			s := c.getFieldCard(duel, 0, coreLocationMZone, uint32(i))
			field.Player1.Spells[i] = &s
		}
		if fieldCounts.player2.monsters[i].present {
			m := c.getFieldCard(duel, 0, coreLocationMZone, uint32(i))
			field.Player2.Monsters[i] = &m
		}
		if fieldCounts.player2.spells[i].present {
			s := c.getFieldCard(duel, 0, coreLocationMZone, uint32(i))
			field.Player2.Spells[i] = &s
		}
	}
	for i := 0; i < 2; i++ {
		if fieldCounts.player1.monsters[5+i].present {
			m := c.getFieldCard(duel, 0, coreLocationMZone, uint32(i))
			field.Player1.ExtraMonsters[i] = &m
		}
		if fieldCounts.player2.monsters[5+i].present {
			m := c.getFieldCard(duel, 0, coreLocationMZone, uint32(i))
			field.Player2.ExtraMonsters[i] = &m
		}
	}
	for i := 0; i < int(fieldCounts.player2.handCount); i++ {
		field.Player1.Hand = append(field.Player1.Hand, c.getFieldDeckCard(duel, 0, coreLocationHand, uint32(i)))
	}
}

func (c *OcgCore) getFieldCard(duel DuelHandle, con uint8, loc coreLocation, seq uint32) (card FieldCard) {
	flagsField := coreQueryCode |
		coreQueryLevel | coreQueryPosition |
		coreQueryAttack | coreQueryDefense | coreQueryEquipCard |
		coreQueryCounters | coreQueryLScale | coreQueryRScale

	data := c.parseQuery(c.duelQuery(duel, flagsField, con, loc, seq))

	card.Code = binary.LittleEndian.Uint32(data[coreQueryCode])
	card.Position = Position(binary.LittleEndian.Uint32(data[coreQueryPosition]))
	card.Level = int(binary.LittleEndian.Uint32(data[coreQueryLevel]))
	card.Defense = int(binary.LittleEndian.Uint32(data[coreQueryDefense]))
	card.Attack = int(binary.LittleEndian.Uint32(data[coreQueryAttack]))
	card.LScale = int(binary.LittleEndian.Uint32(data[coreQueryLScale]))
	card.RScale = int(binary.LittleEndian.Uint32(data[coreQueryRScale]))
	return
}

func (c *OcgCore) getFieldDeckCard(duel DuelHandle, con uint8, loc coreLocation, seq uint32) (card FieldDeckCard) {
	data := c.parseQuery(c.duelQuery(duel, coreQueryCode, con, loc, seq))

	card.Code = binary.LittleEndian.Uint32(data[coreQueryCode])
	return
}

type Field struct {
	Player1 FieldPlayer
	Player2 FieldPlayer
}

type FieldPlayer struct {
	Deck          []FieldDeckCard
	Hand          []FieldDeckCard
	Grave         []FieldCard
	Monsters      [5]*FieldCard
	Spells        [5]*FieldCard
	ExtraMonsters [2]*FieldCard
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
	Code uint32 `json:"code"`
}

func (c *OcgCore) FieldData(duel DuelHandle) {

}

func (c *OcgCore) parseQuery(data []byte) map[coreQuery][]byte {
	b := bytes.NewBuffer(data)

	res := map[coreQuery][]byte{}
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

func (c *OcgCore) duelQueryOverlay(duel DuelHandle, flags coreQuery, con uint8, loc coreLocation, seq uint32, overlaySeq uint32) []byte {
	return c.duelQueryInfo(duel, typeQueryInfo{
		flags:      flags,
		con:        con,
		loc:        loc,
		seq:        seq,
		overlaySeq: overlaySeq,
	})
}

func (c *OcgCore) duelQuery(duel DuelHandle, flags coreQuery, con uint8, loc coreLocation, seq uint32) []byte {
	return c.duelQueryInfo(duel, typeQueryInfo{
		flags: flags,
		con:   con,
		loc:   loc,
		seq:   seq,
	})
}

func (c *OcgCore) duelQueryInfo(duel DuelHandle, info typeQueryInfo) []byte {
	var length uint32
	dataPtr, _, err := c.procDuelQuery.Call(uintptr(duel), uintptr(unsafe.Pointer(&length)), uintptr(unsafe.Pointer(&info)))
	if err != nil && err != windows.ERROR_SUCCESS {
		panic(err)
	}
	return bytesFromPtr(dataPtr, uintptr(length))
}

func (c *OcgCore) duelQueryField(duel DuelHandle) []byte {
	var length uint32
	dataPtr, _, err := c.procDuelQueryField.Call(uintptr(duel), uintptr(unsafe.Pointer(&length)))
	if err != nil && err != windows.ERROR_SUCCESS {
		panic(err)
	}
	return bytesFromPtr(dataPtr, uintptr(length))
}

func (c *OcgCore) parseQueryField(data []byte) coreQueryField {
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

func (c *OcgCore) SetupDeck(duel DuelHandle, player int, mainDeck []uint32, extraDeck []uint32, shuffle bool) {
	if shuffle {
		rand.Shuffle(len(mainDeck), func(i, j int) {
			mainDeck[i], mainDeck[j] = mainDeck[j], mainDeck[i]
		})
	}

	var cardInfo typeNewCardInfo
	cardInfo.duelist = 0
	cardInfo.team = uint8(player)
	cardInfo.con = uint8(player)

	cardInfo.pos = uint32(corePositionFaceDownDefense)

	cardInfo.loc = uint32(coreLocationDeck)
	for _, card := range mainDeck {
		cardInfo.code = card
		c.duelNewCard(duel, cardInfo)
	}
	cardInfo.loc = uint32(coreLocationExtra)
	for _, card := range extraDeck {
		cardInfo.code = card
		c.duelNewCard(duel, cardInfo)
	}
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

type typeQueryInfo struct {
	flags      coreQuery
	con        uint8
	loc        coreLocation
	seq        uint32
	overlaySeq uint32
}

type typeNewCardInfo struct {
	team    uint8
	duelist uint8
	code    uint32
	con     uint8
	loc     uint32
	seq     uint32
	pos     uint32
}

type typePlayer struct {
	startingLP        uint32
	startingDrawCount uint32
	drawCountPerTurn  uint32
}

type typeDuelOptions struct {
	seed           uint32
	flags          uint32
	team1          typePlayer
	team2          typePlayer
	cardReader     uintptr // void fn(void* payload, uint32_t code, OCG_CardData* data);
	payload1       uintptr
	scriptReader   uintptr // int fn(void* payload, OCG_Duel duel, const char* name);
	payload2       uintptr
	logHandler     uintptr // void fn(void* payload, const char* string, int type);
	payload3       uintptr
	cardReaderDone uintptr // void fn(void* payload, OCG_CardData* data);
	payload4       uintptr
}

type typeCardData struct {
	code        uint32
	alias       uint32
	setcodes    uintptr // uint16_t*
	cardtype    uint32
	level       uint32
	attribute   uint32
	race        uint32
	attack      int32
	defense     int32
	lscale      uint32
	rscale      uint32
	link_marker uint32
}

type RawCardData struct {
	Code       uint32
	Alias      uint32
	SetCodes   []uint16
	CardType   uint32
	Level      uint32
	Attribute  uint32
	Race       uint32
	Attack     int32
	Defense    int32
	LScale     uint32
	RScale     uint32
	LinkMarker uint32
}
