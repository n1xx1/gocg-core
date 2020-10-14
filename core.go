package ocgcore

import (
	"bytes"
	"fmt"
	"golang.org/x/sys/windows"
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

func fromCStringPtr(cptr uintptr) string {
	cbytes := (*[1<<30 - 1]byte)(unsafe.Pointer(cptr))
	return fromCString(cbytes[:])
}

func fromCString(cbytes []byte) string {
	size := bytes.IndexByte(cbytes, 0)
	cbytesCopy := make([]byte, size)
	copy(cbytesCopy, cbytes[:size])
	return string(cbytesCopy)
}

type CreateDuelOptions struct {
	Seed         uint32
	Flags        uint32
	CardReader   func(code uint32) RawCardData
	ScriptReader func(path string) []byte
}

func (c *OcgCore) CreateDuel(options CreateDuelOptions) DuelHandle {
	loadedCards := map[uint32]RawCardData{}

	duelHandle := uintptr(0)
	duelOptions := typeDuelOptions{
		seed:  0,
		flags: 0,
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
			name := fromCStringPtr(namePtr)
			contents := options.ScriptReader(name)
			if len(contents) == 0 {
				return 0
			}
			c.LoadScript(duel, name, contents)
			return 1
		}),
		logHandler: windows.NewCallback(func(payload uintptr, messagePtr uintptr, messageType uintptr) uintptr {
			message := fromCStringPtr(messagePtr)
			fmt.Println("log handler: ", message)
			return 0
		}),
		cardReaderDone: windows.NewCallback(func(payload uintptr, data uintptr) uintptr {
			fmt.Println("card reader done")
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

func (c *OcgCore) DuelNewCard(duel DuelHandle, player int, code uint32, controller int, location Location, sequence int, position Position) {
	cardInfo := typeNewCardInfo{
		team:    uint8(player),
		duelist: uint8(player),
		code:    code,
		con:     uint8(controller),
		loc:     uint32(location),
		seq:     uint32(sequence),
		pos:     uint32(position),
	}
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
	fmt.Println(ret)
	return ProcessorFlag(ret)
}

func (c *OcgCore) StartDuel(duel DuelHandle) {
	_, _, err := c.procStartDuel.Call(uintptr(duel))
	if err != nil && err != windows.ERROR_SUCCESS {
		panic(err)
	}
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
