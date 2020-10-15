package ocgcore

import (
	"bytes"
	"encoding/json"
	"reflect"
	"sync"
)

//go:generate go run ocgcore/cmd/enumer -type=BattleAction,IdleAction -json -transform=snake -output response_enumer.go -trimprefix BattleAction,IdleAction

type Response interface {
	responseType() ResponseType
	responseWrite() []byte
}

type ResponseType int

const (
	ResponseTypeSelectBattleCMD ResponseType = iota
	ResponseTypeSelectIdleCMD
	ResponseTypeSelectEffectYN
	ResponseTypeSelectYesNo
	ResponseTypeSelectOption
	ResponseTypeSelectCard
	ResponseTypeSelectChain
	ResponseTypeSelectPlace
	ResponseTypeSelectPosition
	ResponseTypeSelectUnselectCard
)

type BattleAction int

const (
	BattleActionChain BattleAction = iota
	BattleActionAttack
	BattleActionToM2
	BattleActionToEP
)

type ResponseSelectBattleCMD struct {
	Action BattleAction `json:"action"`
	Index  int          `json:"index"`
}

func (r ResponseSelectBattleCMD) responseWrite() []byte {
	var b bytes.Buffer
	writeUint32(&b, (uint32(r.Action)&0xff)|((uint32(r.Index)&0xff)<<16))
	return b.Bytes()
}

type IdleAction int

const (
	IdleActionSummon IdleAction = iota
	IdleActionSpSummon
	IdleActionPosChange
	IdleActionMonsterSet
	IdleActionSpellSet
	IdleActionActivate
	IdleActionToBP
	IdleActionToEP
	IdleActionShuffle
)

type ResponseSelectIdleCMD struct {
	Action IdleAction `json:"action"`
	Index  int        `json:"index"`
}

func (r ResponseSelectIdleCMD) responseWrite() []byte {
	var b bytes.Buffer
	writeUint32(&b, (uint32(r.Action)&0xff)|((uint32(r.Index)&0xff)<<16))
	return b.Bytes()
}

type ResponseSelectEffectYN struct {
	Yes bool `json:"yes"`
}

func (r ResponseSelectEffectYN) responseWrite() []byte {
	var b bytes.Buffer
	if r.Yes {
		writeInt32(&b, 1)
	} else {
		writeInt32(&b, 0)
	}
	return b.Bytes()
}

type ResponseSelectYesNo struct {
	Yes bool `json:"yes"`
}

func (r ResponseSelectYesNo) responseWrite() []byte {
	var b bytes.Buffer
	if r.Yes {
		writeInt32(&b, 1)
	} else {
		writeInt32(&b, 0)
	}
	return b.Bytes()
}

type ResponseSelectOption struct {
	Option int `json:"option"`
}

func (r ResponseSelectOption) responseWrite() []byte {
	var b bytes.Buffer
	writeInt32(&b, int32(r.Option))
	return b.Bytes()
}

type ResponseSelectCard struct {
	Cancel bool  `json:"cancel"`
	Select []int `json:"select,omitempty"`
}

func (r ResponseSelectCard) responseWrite() []byte {
	var b bytes.Buffer
	if r.Cancel {
		writeInt32(&b, -1)
	} else {
		writeInt32(&b, 2) // only support type 2
		writeInt32(&b, int32(len(r.Select)))
		for _, c := range r.Select {
			writeInt8(&b, int8(c))
		}
	}
	return b.Bytes()
}

type ResponseSelectChain struct {
	Chain int `json:"chain"`
}

func (r ResponseSelectChain) responseWrite() []byte {
	var b bytes.Buffer
	writeInt32(&b, int32(r.Chain))
	return b.Bytes()
}

type ResponseSelectPlace struct {
	Places []Place `json:"places"`
}
type Place struct {
	Player   int      `json:"player"`
	Location Location `json:"location"`
	Sequence int      `json:"sequence"`
}

func (r ResponseSelectPlace) responseWrite() []byte {
	var b bytes.Buffer
	for _, p := range r.Places {
		seq := uint8(p.Sequence)
		loc := convertLocation(p.Location)
		if loc == coreLocationPZone {
			loc = coreLocationSZone
			seq -= 6
		}
		writeUint8(&b, uint8(p.Player))
		writeUint8(&b, uint8(loc))
		writeUint8(&b, seq)
	}
	return b.Bytes()
}

type ResponseSelectPosition struct {
	Position Position `json:"position"`
}

func (r ResponseSelectPosition) responseWrite() []byte {
	var b bytes.Buffer
	writeInt32(&b, int32(convertPosition(r.Position)))
	return b.Bytes()
}

type ResponseSelectUnselectCard struct {
	Cancel    bool `json:"cancel"`
	Selection int  `json:"selection"`
}

func (r ResponseSelectUnselectCard) responseWrite() []byte {
	var b bytes.Buffer
	if r.Cancel {
		writeInt32(&b, -1)
	} else {
		writeInt32(&b, 1)
		writeInt32(&b, int32(r.Selection))
	}
	return b.Bytes()
}

type jsonResponseName struct {
	MessageType MessageType `json:"message_type"`
}

func ResponseToJSON(m Message) ([]byte, error) {
	structTypesCacheLock.Lock()
	defer structTypesCacheLock.Unlock()

	v := createStructTypesCache(m.messageType(), reflect.TypeOf(m))
	v.Field(0).Set(reflect.ValueOf(m))
	return json.Marshal(v.Interface())
}

func JSONToResponse(b []byte) (Message, error) {
	structTypesCacheLock.Lock()
	defer structTypesCacheLock.Unlock()

	var m jsonMessageName
	err := json.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}
	typ, ok := unmarshalMap[m.MessageType]
	if !ok {
		return nil, errors.New("type not found")
	}
	v := createStructTypesCache(m.MessageType, typ)
	typV := reflect.New(v.Type())

	err = json.Unmarshal(b, typV.Interface())
	if err != nil {
		return nil, err
	}
	return typV.Elem().Field(0).Interface().(Message), nil
}

var unmarshalMap = map[MessageType]reflect.Type{
	MessageTypeWaitingResponse:    reflect.TypeOf(MessageWaitingResponse{}),
	MessageTypeRetry:              reflect.TypeOf(MessageRetry{}),
	MessageTypeHint:               reflect.TypeOf(MessageHint{}),
	MessageTypeWaiting:            reflect.TypeOf(MessageWaiting{}),
	MessageTypeStart:              reflect.TypeOf(MessageStart{}),
	MessageTypeWin:                reflect.TypeOf(MessageWin{}),
	MessageTypeUpdateData:         reflect.TypeOf(MessageUpdateData{}),
	MessageTypeUpdateCard:         reflect.TypeOf(MessageUpdateCard{}),
	MessageTypeRequestDeck:        reflect.TypeOf(MessageRequestDeck{}),
	MessageTypeSelectBattleCMD:    reflect.TypeOf(MessageSelectBattleCMD{}),
	MessageTypeSelectIdleCMD:      reflect.TypeOf(MessageSelectIdleCMD{}),
	MessageTypeSelectEffectYN:     reflect.TypeOf(MessageSelectEffectYN{}),
	MessageTypeSelectYesNo:        reflect.TypeOf(MessageSelectYesNo{}),
	MessageTypeSelectOption:       reflect.TypeOf(MessageSelectOption{}),
	MessageTypeSelectCard:         reflect.TypeOf(MessageSelectCard{}),
	MessageTypeSelectChain:        reflect.TypeOf(MessageSelectChain{}),
	MessageTypeSelectPlace:        reflect.TypeOf(MessageSelectPlace{}),
	MessageTypeSelectPosition:     reflect.TypeOf(MessageSelectPosition{}),
	MessageTypeSelectTribute:      reflect.TypeOf(MessageSelectTribute{}),
	MessageTypeSortChain:          reflect.TypeOf(MessageSortChain{}),
	MessageTypeSelectCounter:      reflect.TypeOf(MessageSelectCounter{}),
	MessageTypeSelectSum:          reflect.TypeOf(MessageSelectSum{}),
	MessageTypeSelectDisfield:     reflect.TypeOf(MessageSelectDisfield{}),
	MessageTypeSortCard:           reflect.TypeOf(MessageSortCard{}),
	MessageTypeSelectUnselectCard: reflect.TypeOf(MessageSelectUnselectCard{}),
	MessageTypeConfirmDeckTop:     reflect.TypeOf(MessageConfirmDeckTop{}),
	MessageTypeConfirmCards:       reflect.TypeOf(MessageConfirmCards{}),
	MessageTypeShuffleDeck:        reflect.TypeOf(MessageShuffleDeck{}),
	MessageTypeShuffleHand:        reflect.TypeOf(MessageShuffleHand{}),
	MessageTypeRefreshDeck:        reflect.TypeOf(MessageRefreshDeck{}),
	MessageTypeSwapGraveDeck:      reflect.TypeOf(MessageSwapGraveDeck{}),
	MessageTypeShuffleSetCard:     reflect.TypeOf(MessageShuffleSetCard{}),
	MessageTypeReverseDeck:        reflect.TypeOf(MessageReverseDeck{}),
	MessageTypeDeckTop:            reflect.TypeOf(MessageDeckTop{}),
	MessageTypeShuffleExtra:       reflect.TypeOf(MessageShuffleExtra{}),
	MessageTypeNewTurn:            reflect.TypeOf(MessageNewTurn{}),
	MessageTypeNewPhase:           reflect.TypeOf(MessageNewPhase{}),
	MessageTypeConfirmExtraTop:    reflect.TypeOf(MessageConfirmExtraTop{}),
	MessageTypeMove:               reflect.TypeOf(MessageMove{}),
	MessageTypePosChange:          reflect.TypeOf(MessagePosChange{}),
	MessageTypeSet:                reflect.TypeOf(MessageSet{}),
	MessageTypeSwap:               reflect.TypeOf(MessageSwap{}),
	MessageTypeFieldDisabled:      reflect.TypeOf(MessageFieldDisabled{}),
	MessageTypeSummoning:          reflect.TypeOf(MessageSummoning{}),
	MessageTypeSummoned:           reflect.TypeOf(MessageSummoned{}),
	MessageTypeSPSummoning:        reflect.TypeOf(MessageSPSummoning{}),
	MessageTypeSPSummoned:         reflect.TypeOf(MessageSPSummoned{}),
	MessageTypeFlipSummoning:      reflect.TypeOf(MessageFlipSummoning{}),
	MessageTypeFlipSummoned:       reflect.TypeOf(MessageFlipSummoned{}),
	MessageTypeChaining:           reflect.TypeOf(MessageChaining{}),
	MessageTypeChained:            reflect.TypeOf(MessageChained{}),
	MessageTypeChainSolving:       reflect.TypeOf(MessageChainSolving{}),
	MessageTypeChainSolved:        reflect.TypeOf(MessageChainSolved{}),
	MessageTypeChainEnd:           reflect.TypeOf(MessageChainEnd{}),
	MessageTypeChainNegated:       reflect.TypeOf(MessageChainNegated{}),
	MessageTypeChainDisabled:      reflect.TypeOf(MessageChainDisabled{}),
	MessageTypeCardSelected:       reflect.TypeOf(MessageCardSelected{}),
	MessageTypeRandomSelected:     reflect.TypeOf(MessageRandomSelected{}),
	MessageTypeBecomeTarget:       reflect.TypeOf(MessageBecomeTarget{}),
	MessageTypeDraw:               reflect.TypeOf(MessageDraw{}),
	MessageTypeDamage:             reflect.TypeOf(MessageDamage{}),
	MessageTypeRecover:            reflect.TypeOf(MessageRecover{}),
	MessageTypeEquip:              reflect.TypeOf(MessageEquip{}),
	MessageTypeLPUpdate:           reflect.TypeOf(MessageLPUpdate{}),
	MessageTypeUnequip:            reflect.TypeOf(MessageUnequip{}),
	MessageTypeCardTarget:         reflect.TypeOf(MessageCardTarget{}),
	MessageTypeCancelTarget:       reflect.TypeOf(MessageCancelTarget{}),
	MessageTypePayLPCost:          reflect.TypeOf(MessagePayLPCost{}),
	MessageTypeAddCounter:         reflect.TypeOf(MessageAddCounter{}),
	MessageTypeRemoveCounter:      reflect.TypeOf(MessageRemoveCounter{}),
	MessageTypeAttack:             reflect.TypeOf(MessageAttack{}),
	MessageTypeBattle:             reflect.TypeOf(MessageBattle{}),
	MessageTypeAttackDisabled:     reflect.TypeOf(MessageAttackDisabled{}),
	MessageTypeDamageStepStart:    reflect.TypeOf(MessageDamageStepStart{}),
	MessageTypeDamageStepEnd:      reflect.TypeOf(MessageDamageStepEnd{}),
	MessageTypeMissedEffect:       reflect.TypeOf(MessageMissedEffect{}),
	MessageTypeBeChainTarget:      reflect.TypeOf(MessageBeChainTarget{}),
	MessageTypeCreateRelation:     reflect.TypeOf(MessageCreateRelation{}),
	MessageTypeReleaseRelation:    reflect.TypeOf(MessageReleaseRelation{}),
	MessageTypeTossCoin:           reflect.TypeOf(MessageTossCoin{}),
	MessageTypeTossDice:           reflect.TypeOf(MessageTossDice{}),
	MessageTypeRockPaperScissors:  reflect.TypeOf(MessageRockPaperScissors{}),
	MessageTypeHandRes:            reflect.TypeOf(MessageHandRes{}),
	MessageTypeAnnounceRace:       reflect.TypeOf(MessageAnnounceRace{}),
	MessageTypeAnnounceAttribute:  reflect.TypeOf(MessageAnnounceAttribute{}),
	MessageTypeAnnounceCard:       reflect.TypeOf(MessageAnnounceCard{}),
	MessageTypeAnnounceNumber:     reflect.TypeOf(MessageAnnounceNumber{}),
	MessageTypeCardHint:           reflect.TypeOf(MessageCardHint{}),
	MessageTypeTagSwap:            reflect.TypeOf(MessageTagSwap{}),
	MessageTypeReloadField:        reflect.TypeOf(MessageReloadField{}),
	MessageTypeAIName:             reflect.TypeOf(MessageAIName{}),
	MessageTypeShowHint:           reflect.TypeOf(MessageShowHint{}),
	MessageTypePlayerHint:         reflect.TypeOf(MessagePlayerHint{}),
	MessageTypeMatchKill:          reflect.TypeOf(MessageMatchKill{}),
	MessageTypeCustomMessage:      reflect.TypeOf(MessageCustomMessage{}),
	MessageTypeRemoveCards:        reflect.TypeOf(MessageRemoveCards{}),
}

var structTypesCacheLock sync.Mutex
var structTypesCache = map[MessageType]reflect.Value{}

func createStructTypesCache(messageType MessageType, typ reflect.Type) reflect.Value {
	if x, ok := structTypesCache[messageType]; ok {
		return x
	}
	t := reflect.StructOf([]reflect.StructField{
		{
			Name:      typ.Name(),
			Type:      typ,
			Anonymous: true,
		},
		{
			Name: "MessageType",
			Type: reflect.TypeOf(""),
			Tag:  `json:"message_type"`,
		},
	})
	v := reflect.New(t).Elem()
	v.Field(1).Set(reflect.ValueOf(messageType.String()))
	structTypesCache[messageType] = v
	return v
}
