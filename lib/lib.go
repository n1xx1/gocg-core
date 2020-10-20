package lib

import (
	"fmt"
	"sync"
	"unsafe"
)

/*
#cgo LDFLAGS: ${SRCDIR}/libocgcore.a -lstdc++
#include "ocgapi.h"
#include <stdlib.h>
extern void goCallbackOCGDataReader(void* payload, uint32_t code, OCG_CardData* data);
extern void goCallbackOCGDataReaderDone(void* payload, OCG_CardData* data);
extern int goCallbackOCGScriptReader(void* payload, OCG_Duel duel, char* name);
extern void goCallbackOCGLogHandler(void* payload, char* string, int type);
*/
import "C"

var duelLock sync.Mutex
var lastDuelId = uintptr(0)
var mapDuels = map[C.OCG_Duel]uintptr{}
var mapCallbackDataReader = map[uintptr]func(code C.uint32_t, data *C.OCG_CardData){}
var mapCallbackDataReaderDone = map[uintptr]func(data *C.OCG_CardData){}
var mapCallbackScriptReader = map[uintptr]func(duel C.OCG_Duel, name *C.char) C.int{}
var mapCallbackLogHandler = map[uintptr]func(str *C.char, typ C.int){}

//export goCallbackOCGDataReader
func goCallbackOCGDataReader(p unsafe.Pointer, code C.uint32_t, data *C.OCG_CardData) {
	fn := mapCallbackDataReader[uintptr(p)]
	fn(code, data)
}

//export goCallbackOCGDataReaderDone
func goCallbackOCGDataReaderDone(p unsafe.Pointer, data *C.OCG_CardData) {
	fn := mapCallbackDataReaderDone[uintptr(p)]
	fn(data)
}

//export goCallbackOCGScriptReader
func goCallbackOCGScriptReader(p unsafe.Pointer, duel C.OCG_Duel, name *C.char) C.int {
	fn := mapCallbackScriptReader[uintptr(p)]
	return fn(duel, name)
}

//export goCallbackOCGLogHandler
func goCallbackOCGLogHandler(p unsafe.Pointer, str *C.char, typ C.int) {
	fn := mapCallbackLogHandler[uintptr(p)]
	fn(str, typ)
}

func CreateDuel(options DuelOptions) Duel {
	duelLock.Lock()
	defer duelLock.Unlock()

	var cduel C.OCG_Duel
	lastDuelId++

	if options.CardReader == nil {
		panic("card reader nil")
	}
	if options.ScriptReader == nil {
		panic("script reader nil")
	}

	mapCallbackDataReader[lastDuelId] = func(code C.uint32_t, data *C.OCG_CardData) {
		d := options.CardReader(uint32(code))
		setcodes := make([]C.uint16_t, len(d.SetCodes)+1)
		for i, sc := range d.SetCodes {
			setcodes[i] = C.uint16_t(sc)
		}
		setcodes[len(d.SetCodes)] = 0

		*data = C.OCG_CardData{
			code:     C.uint32_t(d.Code),
			alias:    C.uint32_t(d.Alias),
			setcodes: (*C.uint16_t)(&setcodes[0]),
			//type_:       C.uint32_t(d.Type),
			level:       C.uint32_t(d.Level),
			attribute:   C.uint32_t(d.Attribute),
			race:        C.uint32_t(d.Race),
			attack:      C.int32_t(d.Attack),
			defense:     C.int32_t(d.Defense),
			lscale:      C.uint32_t(d.LScale),
			rscale:      C.uint32_t(d.RScale),
			link_marker: C.uint32_t(d.LinkMarker),
		}
	}
	mapCallbackScriptReader[lastDuelId] = func(duel C.OCG_Duel, name *C.char) C.int {
		ok := options.ScriptReader(Duel(duel), C.GoString(name))
		if ok {
			return 1
		}
		return 0
	}
	mapCallbackLogHandler[lastDuelId] = func(str *C.char, typ C.int) {
		options.LogHandler(C.GoString(str), int(typ))
	}
	mapCallbackDataReaderDone[lastDuelId] = func(data *C.OCG_CardData) {
		// TODO: implement
	}

	status := C.OCG_CreateDuel(&cduel, C.OCG_DuelOptions{
		seed:  C.uint32_t(options.Seed),
		flags: C.uint32_t(options.Flags),
		team1: C.OCG_Player{
			startingLP:        C.uint32_t(options.Team1.StartingLP),
			startingDrawCount: C.uint32_t(options.Team1.StartingDrawCount),
			drawCountPerTurn:  C.uint32_t(options.Team1.DrawCountPerTurn),
		},
		team2: C.OCG_Player{
			startingLP:        C.uint32_t(options.Team2.StartingLP),
			startingDrawCount: C.uint32_t(options.Team2.StartingDrawCount),
			drawCountPerTurn:  C.uint32_t(options.Team2.DrawCountPerTurn),
		},
		cardReader:     C.OCG_DataReader(C.goCallbackOCGDataReader),
		payload1:       unsafe.Pointer(lastDuelId),
		scriptReader:   C.OCG_ScriptReader(C.goCallbackOCGScriptReader),
		payload2:       unsafe.Pointer(lastDuelId),
		logHandler:     C.OCG_LogHandler(C.goCallbackOCGLogHandler),
		payload3:       unsafe.Pointer(lastDuelId),
		cardReaderDone: C.OCG_DataReaderDone(C.goCallbackOCGDataReaderDone),
		payload4:       unsafe.Pointer(lastDuelId),
	})
	if status != 0 {
		cleanupCallbacks(lastDuelId)
		panic(fmt.Sprintf("invalid creation status: %v", status))
	}

	mapDuels[cduel] = lastDuelId
	return Duel(cduel)
}

func cleanupCallbacks(id uintptr) {
	delete(mapCallbackDataReader, id)
	delete(mapCallbackScriptReader, id)
	delete(mapCallbackLogHandler, id)
	delete(mapCallbackDataReaderDone, id)
}

func DestroyDuel(duel Duel) {
	duelLock.Lock()
	defer duelLock.Unlock()
	cduel := C.OCG_Duel(duel)
	C.OCG_DestroyDuel(cduel)
}

func DuelNewCard(duel Duel, info NewCardInfo) {
	cduel := C.OCG_Duel(duel)
	C.OCG_DuelNewCard(cduel, C.OCG_NewCardInfo{
		team:    C.uint8_t(info.Team),
		duelist: C.uint8_t(info.Duelist),
		code:    C.uint32_t(info.Code),
		con:     C.uint8_t(info.Controller),
		loc:     C.uint32_t(info.Location),
		seq:     C.uint32_t(info.Sequence),
		pos:     C.uint32_t(info.Position),
	})
}

func StartDuel(duel Duel) {
	cduel := C.OCG_Duel(duel)
	C.OCG_StartDuel(cduel)
}

func DuelProcess(duel Duel) ProcessorFlag {
	return ProcessorFlag(C.OCG_DuelProcess(C.OCG_Duel(duel)))
}
func DuelGetMessage(duel Duel) []byte {
	var length C.uint32_t
	msg := C.OCG_DuelGetMessage(C.OCG_Duel(duel), &length)
	return C.GoBytes(unsafe.Pointer(msg), C.int32_t(length))
}

func DuelSetResponse(duel Duel, buffer []byte) {
	C.OCG_DuelSetResponse(C.OCG_Duel(duel), unsafe.Pointer(&buffer[0]), C.uint32_t(len(buffer)))
}

func LoadScript(duel Duel, buffer []byte, name string) int {
	cname := C.CString(name)
	ret := C.OCG_LoadScript(C.OCG_Duel(duel), (*C.char)(unsafe.Pointer(&buffer[0])), C.uint32_t(len(buffer)), cname)
	C.free(unsafe.Pointer(cname))
	return int(ret)
}

func DuelQueryCount(duel Duel, team uint8, loc Location) uint32 {
	return uint32(C.OCG_DuelQueryCount(C.OCG_Duel(duel), C.uchar(team), C.uint32_t(uint32(loc))))
}

func DuelQuery(duel Duel, info QueryInfo) []byte {
	var length C.uint32_t
	query := C.OCG_DuelQuery(C.OCG_Duel(duel), &length, C.OCG_QueryInfo{
		flags:       C.uint32_t(info.Flags),
		con:         C.uint8_t(info.Controller),
		loc:         C.uint32_t(info.Location),
		seq:         C.uint32_t(info.Sequence),
		overlay_seq: C.uint32_t(info.OverlaySequence),
	})
	return C.GoBytes(unsafe.Pointer(query), C.int32_t(length))
}

func DuelQueryLocation(duel Duel, info QueryInfo) []byte {
	var length C.uint32_t
	query := C.OCG_DuelQueryLocation(C.OCG_Duel(duel), &length, C.OCG_QueryInfo{
		flags:       C.uint32_t(info.Flags),
		con:         C.uint8_t(info.Controller),
		loc:         C.uint32_t(info.Location),
		seq:         C.uint32_t(info.Sequence),
		overlay_seq: C.uint32_t(info.OverlaySequence),
	})
	return C.GoBytes(unsafe.Pointer(query), C.int32_t(length))
}

func DuelQueryField(duel Duel) []byte {
	var length C.uint32_t
	query := C.OCG_DuelQueryField(C.OCG_Duel(duel), &length)
	return C.GoBytes(unsafe.Pointer(query), C.int32_t(length))
}
