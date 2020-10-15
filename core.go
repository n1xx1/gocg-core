package ocgcore

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"golang.org/x/sys/windows"
	"io"
	"unsafe"
)

type duelHandle uintptr

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

type CardReader func(code uint32) RawCardData
type ScriptReader func(path string) []byte

type CreateDuelOptions struct {
	Seed uint32
	Mode DuelMode

	CardReader   CardReader
	ScriptReader ScriptReader
}

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

func (c *OcgCore) CreateDuel(options CreateDuelOptions) *OcgDuel {
	loadedCards := map[uint32]RawCardData{}

	var flags coreDuelMode

	switch options.Mode {
	case DuelModeSpeed:
		flags |= coreDuelModeSpeed
	case DuelModeRush:
		flags |= coreDuelModeRush
	case DuelModeMR1:
		flags |= coreDuelModeMR1
	case DuelModeGoat:
		flags |= coreDuelModeGoat
	case DuelModeMR2:
		flags |= coreDuelModeMR2
	case DuelModeMR3:
		flags |= coreDuelModeMR3
	case DuelModeMR4:
		flags |= coreDuelModeMR4
	case DuelModeMR5:
		flags |= coreDuelModeMR5
	}

	var handle duelHandle
	duelOptions := typeDuelOptions{
		seed:  0,
		flags: flags,
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
			cardData.setCodes = uintptr(unsafe.Pointer(&loadedCard.SetCodes[0]))
			cardData.cardType = loadedCard.CardType
			cardData.level = loadedCard.Level
			cardData.attribute = loadedCard.Attribute
			cardData.race = loadedCard.Race
			cardData.attack = loadedCard.Attack
			cardData.defense = loadedCard.Defense
			cardData.lscale = loadedCard.LScale
			cardData.rscale = loadedCard.RScale
			cardData.linkMarker = loadedCard.LinkMarker

			return 0
		}),
		scriptReader: windows.NewCallback(func(payload uintptr, duel duelHandle, namePtr uintptr) uintptr {
			name := stringFromCStringPtr(namePtr)
			contents := options.ScriptReader(name)
			if len(contents) == 0 {
				return 0
			}
			c.loadScript(duel, name, contents)
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

	_, _, err := c.procCreateDuel.Call(uintptr(unsafe.Pointer(&handle)), uintptr(unsafe.Pointer(&duelOptions)))
	if err != nil && err != windows.ERROR_SUCCESS {
		panic(err)
	}

	c.loadScript(handle, "constant.lua", options.ScriptReader("constant.lua"))
	c.loadScript(handle, "utility.lua", options.ScriptReader("utility.lua"))

	return &OcgDuel{
		c: c,
		h: handle,
	}
}

func (c *OcgCore) loadScript(duel duelHandle, name string, content []byte) {
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

func (c *OcgCore) destroyDuel(duel duelHandle) {
	_, _, err := c.procDestroyDuel.Call(uintptr(duel))
	if err != nil && err != windows.ERROR_SUCCESS {
		panic(err)
	}
}

func (c *OcgCore) duelNewCard(duel duelHandle, cardInfo typeNewCardInfo) {
	_, _, err := c.procDuelNewCard.Call(uintptr(duel), uintptr(unsafe.Pointer(&cardInfo)))
	if err != nil && err != windows.ERROR_SUCCESS {
		panic(err)
	}
}

func (c *OcgCore) duelProcess(duel duelHandle) processorFlag {
	ret, _, err := c.procDuelProcess.Call(uintptr(duel))
	if err != nil && err != windows.ERROR_SUCCESS {
		panic(err)
	}
	return processorFlag(ret)
}

func (c *OcgCore) startDuel(duel duelHandle) {
	_, _, err := c.procStartDuel.Call(uintptr(duel))
	if err != nil && err != windows.ERROR_SUCCESS {
		panic(err)
	}
}

func (c *OcgCore) duelGetMessage(duel duelHandle) [][]byte {
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

func (c *OcgCore) duelSetResponse(duel duelHandle, data []byte) {
	if len(data) == 0 {
		return
	}

	_, _, err := c.procDuelSetResponse.Call(uintptr(duel), uintptr(unsafe.Pointer(&data[0])), uintptr(len(data)))
	if err != nil && err != windows.ERROR_SUCCESS {
		panic(err)
	}
}

func (c *OcgCore) FieldStatus(duel duelHandle) (field Field) {
	c.loadFieldPlayer(duel, &field.Player1, 0)
	c.loadFieldPlayer(duel, &field.Player2, 1)
	return
}

func (c *OcgCore) loadFieldPlayer(duel duelHandle, player *FieldPlayer, con uint8) {
	flagsField := coreQueryCode |
		coreQueryLevel | coreQueryPosition |
		coreQueryAttack | coreQueryDefense | coreQueryEquipCard |
		coreQueryCounters | coreQueryLScale | coreQueryRScale
	flagsDeck := coreQueryCode | coreQueryPosition

	player.Deck = parseFieldDeckCards(c.duelQueryLocation(duel, typeQueryInfo{flags: flagsDeck, con: con, loc: coreLocationDeck}))
	player.ExtraDeck = parseFieldDeckCards(c.duelQueryLocation(duel, typeQueryInfo{flags: flagsDeck, con: con, loc: coreLocationDeck}))
	player.Grave = parseFieldDeckCards(c.duelQueryLocation(duel, typeQueryInfo{flags: flagsDeck, con: con, loc: coreLocationDeck}))
	player.Banished = parseFieldDeckCards(c.duelQueryLocation(duel, typeQueryInfo{flags: flagsDeck, con: con, loc: coreLocationDeck}))
	player.Hand = parseFieldDeckCards(c.duelQueryLocation(duel, typeQueryInfo{flags: flagsDeck, con: con, loc: coreLocationHand}))

	monsters := c.duelQueryLocation(duel, typeQueryInfo{flags: flagsField, con: con, loc: coreLocationMZone})
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

	spells := c.duelQueryLocation(duel, typeQueryInfo{flags: flagsField, con: con, loc: coreLocationSZone})
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

func parseFieldDeckCards(cards []coreQueryResult) []FieldDeckCard {
	res := make([]FieldDeckCard, len(cards))
	for i, data := range cards {
		res[i] = parseFieldDeckCard(data)
	}
	return res
}

func parseFieldDeckCard(data coreQueryResult) (card FieldDeckCard) {
	card.Code = binary.LittleEndian.Uint32(data[coreQueryCode])
	card.Position = parseCorePosition(corePosition(binary.LittleEndian.Uint32(data[coreQueryPosition]))).Face()
	return
}

func parseFieldCard(data coreQueryResult) (card FieldCard) {
	card.Code = binary.LittleEndian.Uint32(data[coreQueryCode])
	card.Position = parseCorePosition(corePosition(binary.LittleEndian.Uint32(data[coreQueryPosition])))
	card.Level = int(binary.LittleEndian.Uint32(data[coreQueryLevel]))
	card.Defense = int(binary.LittleEndian.Uint32(data[coreQueryDefense]))
	card.Attack = int(binary.LittleEndian.Uint32(data[coreQueryAttack]))
	card.LScale = int(binary.LittleEndian.Uint32(data[coreQueryLScale]))
	card.RScale = int(binary.LittleEndian.Uint32(data[coreQueryRScale]))
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

func (c *OcgCore) duelQueryOverlay(duel duelHandle, flags coreQuery, con uint8, loc coreLocation, seq uint32, overlaySeq uint32) coreQueryResult {
	return c.duelQueryInfo(duel, typeQueryInfo{
		flags:      flags,
		con:        con,
		loc:        loc,
		seq:        seq,
		overlaySeq: overlaySeq,
	})
}

func (c *OcgCore) duelQuery(duel duelHandle, flags coreQuery, con uint8, loc coreLocation, seq uint32) coreQueryResult {
	return c.duelQueryInfo(duel, typeQueryInfo{
		flags: flags,
		con:   con,
		loc:   loc,
		seq:   seq,
	})
}

func (c *OcgCore) duelQueryInfo(duel duelHandle, info typeQueryInfo) coreQueryResult {
	var length uint32
	dataPtr, _, err := c.procDuelQuery.Call(uintptr(duel), uintptr(unsafe.Pointer(&length)), uintptr(unsafe.Pointer(&info)))
	if err != nil && err != windows.ERROR_SUCCESS {
		panic(err)
	}
	return parseQuery(bytesFromPtr(dataPtr, uintptr(length)))
}

func (c *OcgCore) duelQueryLocation(duel duelHandle, info typeQueryInfo) []coreQueryResult {
	var length uint32
	dataPtr, _, err := c.procDuelQueryLocation.Call(uintptr(duel), uintptr(unsafe.Pointer(&length)), uintptr(unsafe.Pointer(&info)))
	if err != nil && err != windows.ERROR_SUCCESS {
		panic(err)
	}
	return parseQueryLocation(bytesFromPtr(dataPtr, uintptr(length)))
}

func (c *OcgCore) duelQueryField(duel duelHandle) coreQueryField {
	var length uint32
	dataPtr, _, err := c.procDuelQueryField.Call(uintptr(duel), uintptr(unsafe.Pointer(&length)))
	if err != nil && err != windows.ERROR_SUCCESS {
		panic(err)
	}
	return parseQueryField(bytesFromPtr(dataPtr, uintptr(length)))
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
	flags          coreDuelMode
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
	code       uint32
	alias      uint32
	setCodes   uintptr // uint16_t*
	cardType   uint32
	level      uint32
	attribute  uint32
	race       uint32
	attack     int32
	defense    int32
	lscale     uint32
	rscale     uint32
	linkMarker uint32
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
