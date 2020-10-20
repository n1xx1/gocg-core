package lib

type Duel uintptr

type DuelOptions struct {
	Seed           uint32
	Flags          DuelMode
	Team1          Player
	Team2          Player
	CardReader     func(code uint32) CardData
	ScriptReader   func(duel Duel, name string) bool
	LogHandler     func(s string, typ int)
	CardReaderDone func(data CardData)
}

type Player struct {
	StartingLP        uint32
	StartingDrawCount uint32
	DrawCountPerTurn  uint32
}

type CardData struct {
	Code       uint32
	Alias      uint32
	SetCodes   []uint16
	Type       CardType
	Level      uint32
	Attribute  Attribute
	Race       Race
	Attack     int32
	Defense    int32
	LScale     uint32
	RScale     uint32
	LinkMarker LinkMarker
}

type NewCardInfo struct {
	Team       uint8
	Duelist    uint8
	Code       uint32
	Controller uint8
	Location   Location
	Sequence   uint32
	Position   Position
}
type QueryInfo struct {
	Flags           Query
	Controller      uint8
	Location        Location
	Sequence        uint32
	OverlaySequence uint32
}
