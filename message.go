package ocgcore

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"sync"
)

//go:generate go run ocgcore/cmd/enumer -type=MessageType -json -transform=snake -output message_enumer.go -trimprefix MessageType

type cardLocation struct {
	controller int
	location   coreLocation
	sequence   int
	position   corePosition
}

func readMessage(contents []byte) Message {
	b := bytes.NewBuffer(contents)

	id := readUint8(b)

	switch coreMessage(id) {
	case coreMessageRetry:
		return ReadMessageRetry(b)
	case coreMessageHint:
		return ReadMessageHint(b)
	case coreMessageWaiting:
		return ReadMessageWaiting(b)
	case coreMessageStart:
		return ReadMessageStart(b)
	case coreMessageWin:
		return ReadMessageWin(b)
	case coreMessageUpdateData:
		return ReadMessageUpdateData(b)
	case coreMessageUpdateCard:
		return ReadMessageUpdateCard(b)
	case coreMessageRequestDeck:
		return ReadMessageRequestDeck(b)
	case coreMessageSelectBattleCMD:
		return ReadMessageSelectBattleCMD(b)
	case coreMessageSelectIdleCMD:
		return ReadMessageSelectIdleCMD(b)
	case coreMessageSelectEffectYN:
		return ReadMessageSelectEffectYN(b)
	case coreMessageSelectYesNo:
		return ReadMessageSelectYesNo(b)
	case coreMessageSelectOption:
		return ReadMessageSelectOption(b)
	case coreMessageSelectCard:
		return ReadMessageSelectCard(b)
	case coreMessageSelectChain:
		return ReadMessageSelectChain(b)
	case coreMessageSelectPlace:
		return ReadMessageSelectPlace(b)
	case coreMessageSelectPosition:
		return ReadMessageSelectPosition(b)
	case coreMessageSelectTribute:
		return ReadMessageSelectTribute(b)
	case coreMessageSortChain:
		return ReadMessageSortChain(b)
	case coreMessageSelectCounter:
		return ReadMessageSelectCounter(b)
	case coreMessageSelectSum:
		return ReadMessageSelectSum(b)
	case coreMessageSelectDisfield:
		return ReadMessageSelectDisfield(b)
	case coreMessageSortCard:
		return ReadMessageSortCard(b)
	case coreMessageSelectUnselectCard:
		return ReadMessageSelectUnselectCard(b)
	case coreMessageConfirmDeckTop:
		return ReadMessageConfirmDeckTop(b)
	case coreMessageConfirmCards:
		return ReadMessageConfirmCards(b)
	case coreMessageShuffleDeck:
		return ReadMessageShuffleDeck(b)
	case coreMessageShuffleHand:
		return ReadMessageShuffleHand(b)
	case coreMessageRefreshDeck:
		return ReadMessageRefreshDeck(b)
	case coreMessageSwapGraveDeck:
		return ReadMessageSwapGraveDeck(b)
	case coreMessageShuffleSetCard:
		return ReadMessageShuffleSetCard(b)
	case coreMessageReverseDeck:
		return ReadMessageReverseDeck(b)
	case coreMessageDeckTop:
		return ReadMessageDeckTop(b)
	case coreMessageShuffleExtra:
		return ReadMessageShuffleExtra(b)
	case coreMessageNewTurn:
		return ReadMessageNewTurn(b)
	case coreMessageNewPhase:
		return ReadMessageNewPhase(b)
	case coreMessageConfirmExtraTop:
		return ReadMessageConfirmExtraTop(b)
	case coreMessageMove:
		return ReadMessageMove(b)
	case coreMessagePosChange:
		return ReadMessagePosChange(b)
	case coreMessageSet:
		return ReadMessageSet(b)
	case coreMessageSwap:
		return ReadMessageSwap(b)
	case coreMessageFieldDisabled:
		return ReadMessageFieldDisabled(b)
	case coreMessageSummoning:
		return ReadMessageSummoning(b)
	case coreMessageSummoned:
		return ReadMessageSummoned(b)
	case coreMessageSPSummoning:
		return ReadMessageSPSummoning(b)
	case coreMessageSPSummoned:
		return ReadMessageSPSummoned(b)
	case coreMessageFlipSummoning:
		return ReadMessageFlipSummoning(b)
	case coreMessageFlipSummoned:
		return ReadMessageFlipSummoned(b)
	case coreMessageChaining:
		return ReadMessageChaining(b)
	case coreMessageChained:
		return ReadMessageChained(b)
	case coreMessageChainSolving:
		return ReadMessageChainSolving(b)
	case coreMessageChainSolved:
		return ReadMessageChainSolved(b)
	case coreMessageChainEnd:
		return ReadMessageChainEnd(b)
	case coreMessageChainNegated:
		return ReadMessageChainNegated(b)
	case coreMessageChainDisabled:
		return ReadMessageChainDisabled(b)
	case coreMessageCardSelected:
		return ReadMessageCardSelected(b)
	case coreMessageRandomSelected:
		return ReadMessageRandomSelected(b)
	case coreMessageBecomeTarget:
		return ReadMessageBecomeTarget(b)
	case coreMessageDraw:
		return ReadMessageDraw(b)
	case coreMessageDamage:
		return ReadMessageDamage(b)
	case coreMessageRecover:
		return ReadMessageRecover(b)
	case coreMessageEquip:
		return ReadMessageEquip(b)
	case coreMessageLPUpdate:
		return ReadMessageLPUpdate(b)
	case coreMessageUnequip:
		return ReadMessageUnequip(b)
	case coreMessageCardTarget:
		return ReadMessageCardTarget(b)
	case coreMessageCancelTarget:
		return ReadMessageCancelTarget(b)
	case coreMessagePayLPCost:
		return ReadMessagePayLPCost(b)
	case coreMessageAddCounter:
		return ReadMessageAddCounter(b)
	case coreMessageRemoveCounter:
		return ReadMessageRemoveCounter(b)
	case coreMessageAttack:
		return ReadMessageAttack(b)
	case coreMessageBattle:
		return ReadMessageBattle(b)
	case coreMessageAttackDisabled:
		return ReadMessageAttackDisabled(b)
	case coreMessageDamageStepStart:
		return ReadMessageDamageStepStart(b)
	case coreMessageDamageStepEnd:
		return ReadMessageDamageStepEnd(b)
	case coreMessageMissedEffect:
		return ReadMessageMissedEffect(b)
	case coreMessageBeChainTarget:
		return ReadMessageBeChainTarget(b)
	case coreMessageCreateRelation:
		return ReadMessageCreateRelation(b)
	case coreMessageReleaseRelation:
		return ReadMessageReleaseRelation(b)
	case coreMessageTossCoin:
		return ReadMessageTossCoin(b)
	case coreMessageTossDice:
		return ReadMessageTossDice(b)
	case coreMessageRockPaperScissors:
		return ReadMessageRockPaperScissors(b)
	case coreMessageHandRes:
		return ReadMessageHandRes(b)
	case coreMessageAnnounceRace:
		return ReadMessageAnnounceRace(b)
	case coreMessageAnnounceAttribute:
		return ReadMessageAnnounceAttribute(b)
	case coreMessageAnnounceCard:
		return ReadMessageAnnounceCard(b)
	case coreMessageAnnounceNumber:
		return ReadMessageAnnounceNumber(b)
	case coreMessageCardHint:
		return ReadMessageCardHint(b)
	case coreMessageTagSwap:
		return ReadMessageTagSwap(b)
	case coreMessageReloadField:
		return ReadMessageReloadField(b)
	case coreMessageAIName:
		return ReadMessageAIName(b)
	case coreMessageShowHint:
		return ReadMessageShowHint(b)
	case coreMessagePlayerHint:
		return ReadMessagePlayerHint(b)
	case coreMessageMatchKill:
		return ReadMessageMatchKill(b)
	case coreMessageCustomMessage:
		return ReadMessageCustomMessage(b)
	case coreMessageRemoveCards:
		return ReadMessageRemoveCards(b)
	default:
		fmt.Println("unhandled message:", id, "size:", len(contents)-1)
		return nil
	}
}

func readCardLocation(b *bytes.Buffer) cardLocation {
	return cardLocation{
		controller: int(readUint8(b)),
		location:   coreLocation(readUint8(b)),
		sequence:   int(readUint32(b)),
		position:   corePosition(readUint32(b)),
	}
}

type MessageType uint8

const (
	MessageTypeWaitingResponse MessageType = iota
	MessageTypeRetry
	MessageTypeHint
	MessageTypeWaiting
	MessageTypeStart
	MessageTypeWin
	MessageTypeUpdateData
	MessageTypeUpdateCard
	MessageTypeRequestDeck
	MessageTypeSelectBattleCMD
	MessageTypeSelectIdleCMD
	MessageTypeSelectEffectYN
	MessageTypeSelectYesNo
	MessageTypeSelectOption
	MessageTypeSelectCard
	MessageTypeSelectChain
	MessageTypeSelectPlace
	MessageTypeSelectPosition
	MessageTypeSelectTribute
	MessageTypeSortChain
	MessageTypeSelectCounter
	MessageTypeSelectSum
	MessageTypeSelectDisfield
	MessageTypeSortCard
	MessageTypeSelectUnselectCard
	MessageTypeConfirmDeckTop
	MessageTypeConfirmCards
	MessageTypeShuffleDeck
	MessageTypeShuffleHand
	MessageTypeRefreshDeck
	MessageTypeSwapGraveDeck
	MessageTypeShuffleSetCard
	MessageTypeReverseDeck
	MessageTypeDeckTop
	MessageTypeShuffleExtra
	MessageTypeNewTurn
	MessageTypeNewPhase
	MessageTypeConfirmExtraTop
	MessageTypeMove
	MessageTypePosChange
	MessageTypeSet
	MessageTypeSwap
	MessageTypeFieldDisabled
	MessageTypeSummoning
	MessageTypeSummoned
	MessageTypeSPSummoning
	MessageTypeSPSummoned
	MessageTypeFlipSummoning
	MessageTypeFlipSummoned
	MessageTypeChaining
	MessageTypeChained
	MessageTypeChainSolving
	MessageTypeChainSolved
	MessageTypeChainEnd
	MessageTypeChainNegated
	MessageTypeChainDisabled
	MessageTypeCardSelected
	MessageTypeRandomSelected
	MessageTypeBecomeTarget
	MessageTypeDraw
	MessageTypeDamage
	MessageTypeRecover
	MessageTypeEquip
	MessageTypeLPUpdate
	MessageTypeUnequip
	MessageTypeCardTarget
	MessageTypeCancelTarget
	MessageTypePayLPCost
	MessageTypeAddCounter
	MessageTypeRemoveCounter
	MessageTypeAttack
	MessageTypeBattle
	MessageTypeAttackDisabled
	MessageTypeDamageStepStart
	MessageTypeDamageStepEnd
	MessageTypeMissedEffect
	MessageTypeBeChainTarget
	MessageTypeCreateRelation
	MessageTypeReleaseRelation
	MessageTypeTossCoin
	MessageTypeTossDice
	MessageTypeRockPaperScissors
	MessageTypeHandRes
	MessageTypeAnnounceRace
	MessageTypeAnnounceAttribute
	MessageTypeAnnounceCard
	MessageTypeAnnounceNumber
	MessageTypeCardHint
	MessageTypeTagSwap
	MessageTypeReloadField
	MessageTypeAIName
	MessageTypeShowHint
	MessageTypePlayerHint
	MessageTypeMatchKill
	MessageTypeCustomMessage
	MessageTypeRemoveCards
)

type Message interface {
	messageType() MessageType
}

type ChainInfo struct {
	Code        int      `json:"code"`
	Controller  int      `json:"controller"`
	Location    Location `json:"location"`
	Sequence    int      `json:"sequence"`
	Description uint64   `json:"description"`
	ClientMode  uint8    `json:"client_mode"`
}

type AttackInfo struct {
	Code       int      `json:"code"`
	Controller int      `json:"controller"`
	Location   Location `json:"location"`
	Sequence   int      `json:"sequence"`
	Direct     bool     `json:"direct"`
}

type CardInfo struct {
	Code       int      `json:"code"`
	Controller int      `json:"controller"`
	Location   Location `json:"location"`
	Sequence   int      `json:"sequence"`
}

type FieldCardInfo struct {
	Code int `json:"code"`
	CardLocation
}

type CounterCardInfo struct {
	Code       int      `json:"code"`
	Controller int      `json:"controller"`
	Location   Location `json:"location"`
	Sequence   int      `json:"sequence"`
	Count      int      `json:"count"`
}

type CardChainInfo struct {
	Code int `json:"code"`
	CardLocation
	Description uint64 `json:"description"`
	ClientMode  uint8  `json:"client_mode"`
}

type TributeCardInfo struct {
	Code         int      `json:"code"`
	Controller   int      `json:"controller"`
	Location     Location `json:"location"`
	Sequence     int      `json:"sequence"`
	ReleaseParam int      `json:"release"`
}

type DrawnCardInfo struct {
	Code     int          `json:"code"`
	Position FacePosition `json:"position"`
}

type CardLocation struct {
	Controller int      `json:"controller"`
	Location   Location `json:"location"`
	Sequence   int      `json:"sequence"`
	Position   Position `json:"position"`
}

type MessageWaitingResponse struct{}

func (MessageWaitingResponse) messageType() MessageType {
	return MessageTypeWaitingResponse
}

type MessageRetry struct{}

func ReadMessageRetry(*bytes.Buffer) (msg MessageRetry) {
	return
}

func (MessageRetry) messageType() MessageType {
	return MessageTypeRetry
}

type MessageHint struct {
	Hint   int    `json:"hint"`
	Player int    `json:"player"`
	Desc   uint64 `json:"desc"`
}

func ReadMessageHint(b *bytes.Buffer) (msg MessageHint) {
	msg.Hint = int(readUint8(b))
	msg.Player = int(readUint8(b))
	msg.Desc = readUint64(b)
	return
}

func (MessageHint) messageType() MessageType {
	return MessageTypeHint
}

type MessageWaiting struct{}

func ReadMessageWaiting(*bytes.Buffer) (msg MessageWaiting) {
	return
}

func (MessageWaiting) messageType() MessageType {
	return MessageTypeWaiting
}

type MessageStart struct{}

func ReadMessageStart(*bytes.Buffer) (msg MessageStart) {
	return
}

func (MessageStart) messageType() MessageType {
	return MessageTypeStart
}

type MessageWin struct {
	Player int `json:"player"`
	Reason int `json:"reason"`
}

func ReadMessageWin(b *bytes.Buffer) (msg MessageWin) {
	msg.Player = int(readUint8(b))
	msg.Reason = int(readUint8(b))
	return
}

func (MessageWin) messageType() MessageType {
	return MessageTypeWin
}

type MessageUpdateData struct{}

func ReadMessageUpdateData(*bytes.Buffer) (msg MessageUpdateData) {
	return
}

func (MessageUpdateData) messageType() MessageType {
	return MessageTypeUpdateData
}

type MessageUpdateCard struct{}

func ReadMessageUpdateCard(*bytes.Buffer) (msg MessageUpdateCard) {
	return
}

func (MessageUpdateCard) messageType() MessageType {
	return MessageTypeUpdateCard
}

type MessageRequestDeck struct{}

func ReadMessageRequestDeck(*bytes.Buffer) (msg MessageRequestDeck) {
	return
}

func (MessageRequestDeck) messageType() MessageType {
	return MessageTypeRequestDeck
}

type MessageSelectBattleCMD struct {
	Player  int          `json:"player"`
	Chains  []ChainInfo  `json:"chains"`
	Attacks []AttackInfo `json:"attacks"`
	ToM2    bool         `json:"to_m2"`
	ToEP    bool         `json:"to_ep"`
}

func ReadMessageSelectBattleCMD(b *bytes.Buffer) (msg MessageSelectBattleCMD) {
	msg.Player = int(readUint8(b))

	selectChainsSize := readUint32(b)
	msg.Chains = make([]ChainInfo, selectChainsSize)
	for i := range msg.Chains {
		msg.Chains[i] = ChainInfo{
			Code:        int(readUint32(b)),
			Controller:  int(readUint8(b)),
			Location:    parseCoreLocation(coreLocation(readUint8(b))),
			Sequence:    int(readUint32(b)),
			Description: readUint64(b),
			ClientMode:  readUint8(b),
		}
	}

	attackableSize := readUint32(b)
	msg.Attacks = make([]AttackInfo, attackableSize)
	for i := range msg.Attacks {
		msg.Attacks[i] = AttackInfo{
			Code:       int(readUint32(b)),
			Controller: int(readUint8(b)),
			Location:   parseCoreLocation(coreLocation(readUint8(b))),
			Sequence:   int(readUint32(b)),
			Direct:     readUint8(b) != 0,
		}
	}
	msg.ToM2 = readUint8(b) != 0
	msg.ToEP = readUint8(b) != 0
	return
}

func (MessageSelectBattleCMD) messageType() MessageType {
	return MessageTypeSelectBattleCMD
}

type MessageSelectIdleCMD struct {
	Player      int         `json:"player"`
	Summons     []CardInfo  `json:"summons"`
	SpSummons   []CardInfo  `json:"sp_summons"`
	PosChanges  []CardInfo  `json:"pos_changes"`
	MonsterSets []CardInfo  `json:"monster_sets"`
	SpellSets   []CardInfo  `json:"spell_sets"`
	Activate    []ChainInfo `json:"activate"`
	ToBP        bool        `json:"to_bp"`
	ToEP        bool        `json:"to_ep"`
	Shuffle     bool        `json:"shuffle"`
}

func ReadMessageSelectIdleCMD(b *bytes.Buffer) (msg MessageSelectIdleCMD) {
	msg.Player = int(readUint8(b))

	summonableSize := readUint32(b)
	msg.Summons = make([]CardInfo, summonableSize)
	for i := range msg.Summons {
		msg.Summons[i] = CardInfo{
			Code:       int(readUint32(b)),
			Controller: int(readUint8(b)),
			Location:   parseCoreLocation(coreLocation(readUint8(b))),
			Sequence:   int(readUint32(b)),
		}
	}

	spSummonableSize := readUint32(b)
	msg.SpSummons = make([]CardInfo, spSummonableSize)
	for i := range msg.SpSummons {
		msg.SpSummons[i] = CardInfo{
			Code:       int(readUint32(b)),
			Controller: int(readUint8(b)),
			Location:   parseCoreLocation(coreLocation(readUint8(b))),
			Sequence:   int(readUint32(b)),
		}
	}

	posChangeSize := readUint32(b)
	msg.PosChanges = make([]CardInfo, posChangeSize)
	for i := range msg.PosChanges {
		msg.PosChanges[i] = CardInfo{
			Code:       int(readUint32(b)),
			Controller: int(readUint8(b)),
			Location:   parseCoreLocation(coreLocation(readUint8(b))),
			Sequence:   int(readUint32(b)),
		}
	}

	monsterSetSize := readUint32(b)
	msg.MonsterSets = make([]CardInfo, monsterSetSize)
	for i := range msg.MonsterSets {
		msg.MonsterSets[i] = CardInfo{
			Code:       int(readUint32(b)),
			Controller: int(readUint8(b)),
			Location:   parseCoreLocation(coreLocation(readUint8(b))),
			Sequence:   int(readUint32(b)),
		}
	}

	spellSetSize := readUint32(b)
	msg.SpellSets = make([]CardInfo, spellSetSize)
	for i := range msg.SpellSets {
		msg.SpellSets[i] = CardInfo{
			Code:       int(readUint32(b)),
			Controller: int(readUint8(b)),
			Location:   parseCoreLocation(coreLocation(readUint8(b))),
			Sequence:   int(readUint32(b)),
		}
	}

	activateSize := readUint32(b)
	msg.Activate = make([]ChainInfo, activateSize)
	for i := range msg.Activate {
		msg.Activate[i] = ChainInfo{
			Code:        int(readUint32(b)),
			Controller:  int(readUint8(b)),
			Location:    parseCoreLocation(coreLocation(readUint8(b))),
			Sequence:    int(readUint32(b)),
			Description: readUint64(b),
			ClientMode:  readUint8(b),
		}
	}

	msg.ToBP = readUint8(b) != 0
	msg.ToEP = readUint8(b) != 0
	msg.Shuffle = readUint8(b) != 0
	return
}

func (MessageSelectIdleCMD) messageType() MessageType {
	return MessageTypeSelectIdleCMD
}

type MessageSelectEffectYN struct {
	Player      int      `json:"player"`
	Code        uint32   `json:"code"`
	Controller  int      `json:"controller"`
	Location    Location `json:"location"`
	Sequence    int      `json:"sequence"`
	Position    Position `json:"position"`
	Description uint64   `json:"description"`
}

func ReadMessageSelectEffectYN(b *bytes.Buffer) (msg MessageSelectEffectYN) {
	msg.Player = int(readUint8(b))
	msg.Code = readUint32(b)
	loc := readCardLocation(b)
	msg.Controller = loc.controller
	msg.Location = parseCoreLocation(loc.location)
	msg.Sequence = loc.sequence
	msg.Position = parseCorePosition(loc.position)
	msg.Description = readUint64(b)
	return
}

func (MessageSelectEffectYN) messageType() MessageType {
	return MessageTypeSelectEffectYN
}

type MessageSelectYesNo struct {
	Player      int    `json:"player"`
	Description uint64 `json:"description"`
}

func ReadMessageSelectYesNo(b *bytes.Buffer) (msg MessageSelectYesNo) {
	msg.Player = int(readUint8(b))
	msg.Description = readUint64(b)
	return
}

func (MessageSelectYesNo) messageType() MessageType {
	return MessageTypeSelectYesNo
}

type MessageSelectOption struct {
	Player  int      `json:"player"`
	Options []uint64 `json:"options"`
}

func ReadMessageSelectOption(b *bytes.Buffer) (msg MessageSelectOption) {
	msg.Player = int(readUint8(b))

	optionsSize := readUint8(b)
	msg.Options = make([]uint64, optionsSize)
	for i := range msg.Options {
		msg.Options[i] = readUint64(b)
	}
	return
}

func (MessageSelectOption) messageType() MessageType {
	return MessageTypeSelectOption
}

type MessageSelectCard struct {
	Player      int             `json:"player"`
	Cancellable bool            `json:"cancellable"`
	Min         int             `json:"min"`
	Max         int             `json:"max"`
	Cards       []FieldCardInfo `json:"cards"`
}

func parseCardLocation(loc cardLocation) CardLocation {
	return CardLocation{
		Controller: loc.controller,
		Location:   parseCoreLocation(loc.location),
		Sequence:   loc.sequence,
		Position:   parseCorePosition(loc.position),
	}
}

func ReadMessageSelectCard(b *bytes.Buffer) (msg MessageSelectCard) {
	msg.Player = int(readUint8(b))
	msg.Cancellable = readUint8(b) != 0
	msg.Min = int(readUint32(b))
	msg.Max = int(readUint32(b))

	cardsSize := readUint32(b)
	msg.Cards = make([]FieldCardInfo, cardsSize)
	for i := range msg.Cards {
		msg.Cards[i] = FieldCardInfo{
			Code:         int(readUint32(b)),
			CardLocation: parseCardLocation(readCardLocation(b)),
		}
	}
	return
}

func (MessageSelectCard) messageType() MessageType {
	return MessageTypeSelectCard
}

type MessageSelectChain struct {
	Player           int             `json:"player"`
	SpeCount         int             `json:"spe_count"`
	Forced           bool            `json:"forced"`
	HintTimingPlayer uint32          `json:"hint_timing_player"`
	HintTimingOther  uint32          `json:"hint_timing_other"`
	Chains           []CardChainInfo `json:"chains"`
}

func ReadMessageSelectChain(b *bytes.Buffer) (msg MessageSelectChain) {
	msg.Player = int(readUint8(b))
	msg.SpeCount = int(readUint8(b))
	msg.Forced = readUint8(b) != 0
	msg.HintTimingPlayer = readUint32(b)
	msg.HintTimingOther = readUint32(b)

	chainsSize := readUint32(b)
	msg.Chains = make([]CardChainInfo, chainsSize)
	for i := range msg.Chains {
		msg.Chains[i] = CardChainInfo{
			Code:         int(readUint32(b)),
			CardLocation: parseCardLocation(readCardLocation(b)),
			Description:  readUint64(b),
			ClientMode:   readUint8(b),
		}
	}
	return
}

func (MessageSelectChain) messageType() MessageType {
	return MessageTypeSelectChain
}

type MessageSelectPlace struct {
	Player int     `json:"player"`
	Count  int     `json:"count"`
	Places []Place `json:"places"`
}

func ReadMessageSelectPlace(b *bytes.Buffer) (msg MessageSelectPlace) {
	msg.Player = int(readUint8(b))
	msg.Count = int(readUint8(b))
	msg.Places = parsePlaceFlag(readUint32(b))
	return
}

func (MessageSelectPlace) messageType() MessageType {
	return MessageTypeSelectPlace
}

type MessageSelectPosition struct {
	Player    int        `json:"player"`
	Code      uint32     `json:"code"`
	Positions []Position `json:"positions"`
}

func ReadMessageSelectPosition(b *bytes.Buffer) (msg MessageSelectPosition) {
	msg.Player = int(readUint8(b))
	msg.Code = readUint32(b)
	msg.Positions = parseCorePositions(corePosition(readUint8(b)))
	return
}

func (MessageSelectPosition) messageType() MessageType {
	return MessageTypeSelectPosition
}

type MessageSelectTribute struct {
	Player      int               `json:"player"`
	Cancellable bool              `json:"cancellable"`
	Min         int               `json:"min"`
	Max         int               `json:"max"`
	Cards       []TributeCardInfo `json:"cards"`
}

func ReadMessageSelectTribute(b *bytes.Buffer) (msg MessageSelectTribute) {
	msg.Player = int(readUint8(b))
	msg.Cancellable = readUint8(b) != 0
	msg.Min = int(readUint32(b))
	msg.Max = int(readUint32(b))

	tributeSize := readUint32(b)
	msg.Cards = make([]TributeCardInfo, tributeSize)
	for i := range msg.Cards {
		msg.Cards[i] = TributeCardInfo{
			Code:         int(readUint32(b)),
			Controller:   int(readUint8(b)),
			Location:     parseCoreLocation(coreLocation(readUint8(b))),
			Sequence:     int(readUint32(b)),
			ReleaseParam: int(readUint8(b)),
		}
	}
	return
}

func (MessageSelectTribute) messageType() MessageType {
	return MessageTypeSelectTribute
}

type MessageSortChain struct {
	Player int        `json:"player"`
	Cards  []CardInfo `json:"cards"`
}

func ReadMessageSortChain(b *bytes.Buffer) (msg MessageSortChain) {
	msg.Player = int(readUint8(b))
	cardsSize := readUint32(b)
	msg.Cards = make([]CardInfo, cardsSize)
	for i := range msg.Cards {
		msg.Cards[i] = CardInfo{
			Code:       int(readUint32(b)),
			Controller: int(readUint8(b)),
			Location:   parseCoreLocation(coreLocation(readUint8(b))),
			Sequence:   int(readUint32(b)),
		}
	}
	return
}

func (MessageSortChain) messageType() MessageType {
	return MessageTypeSortChain
}

type MessageSelectCounter struct {
	Player      int               `json:"player"`
	CounterType int               `json:"counter_type"`
	Count       int               `json:"count"`
	Cards       []CounterCardInfo `json:"cards"`
}

func ReadMessageSelectCounter(b *bytes.Buffer) (msg MessageSelectCounter) {
	msg.Player = int(readUint8(b))
	msg.CounterType = int(readUint16(b))
	msg.Count = int(readUint16(b))
	cardsSize := readUint32(b)
	msg.Cards = make([]CounterCardInfo, cardsSize)
	for i := range msg.Cards {
		msg.Cards[i] = CounterCardInfo{
			Code:       int(readUint32(b)),
			Controller: int(readUint8(b)),
			Location:   parseCoreLocation(coreLocation(readUint8(b))),
			Sequence:   int(readUint8(b)),
			Count:      int(readUint16(b)),
		}
	}
	return
}

func (MessageSelectCounter) messageType() MessageType {
	return MessageTypeSelectCounter
}

type MessageSelectSum struct {
	Player      int               `json:"player"`
	HasMax      bool              `json:"has_max"`
	Acc         int               `json:"acc"`
	Min         int               `json:"min"`
	Max         int               `json:"max"`
	MustSelects []CounterCardInfo `json:"must_selects"`
	Selects     []CounterCardInfo `json:"selects"`
}

func ReadMessageSelectSum(b *bytes.Buffer) (msg MessageSelectSum) {
	msg.Player = int(readUint8(b))
	msg.HasMax = readUint8(b) != 0
	msg.Acc = int(readUint32(b))
	msg.Min = int(readUint32(b))
	msg.Max = int(readUint32(b))
	mustSelectsSize := readUint32(b)
	msg.MustSelects = make([]CounterCardInfo, mustSelectsSize)
	for i := range msg.MustSelects {
		msg.MustSelects[i] = CounterCardInfo{
			Code:       int(readUint32(b)),
			Controller: int(readUint8(b)),
			Location:   parseCoreLocation(coreLocation(readUint8(b))),
			Sequence:   int(readUint32(b)),
			Count:      int(readUint32(b)),
		}
	}
	selectsSize := readUint32(b)
	msg.Selects = make([]CounterCardInfo, selectsSize)
	for i := range msg.Selects {
		msg.Selects[i] = CounterCardInfo{
			Code:       int(readUint32(b)),
			Controller: int(readUint8(b)),
			Location:   parseCoreLocation(coreLocation(readUint8(b))),
			Sequence:   int(readUint32(b)),
			Count:      int(readUint32(b)),
		}
	}
	return
}

func (MessageSelectSum) messageType() MessageType {
	return MessageTypeSelectSum
}

type MessageSelectDisfield struct {
	Player int    `json:"player"`
	Count  int    `json:"count"`
	Flag   uint32 `json:"flag"`
}

func ReadMessageSelectDisfield(b *bytes.Buffer) (msg MessageSelectDisfield) {
	msg.Player = int(readUint8(b))
	msg.Count = int(readUint8(b))
	msg.Flag = readUint32(b)
	return
}

func (MessageSelectDisfield) messageType() MessageType {
	return MessageTypeSelectDisfield
}

type MessageSortCard struct {
	Player int        `json:"player"`
	Cards  []CardInfo `json:"cards"`
}

func ReadMessageSortCard(b *bytes.Buffer) (msg MessageSortCard) {
	msg.Player = int(readUint8(b))
	cardsSize := readUint32(b)
	msg.Cards = make([]CardInfo, cardsSize)
	for i := range msg.Cards {
		msg.Cards[i] = CardInfo{
			Code:       int(readUint32(b)),
			Controller: int(readUint8(b)),
			Location:   parseCoreLocation(coreLocation(readUint8(b))),
			Sequence:   int(readUint32(b)),
		}
	}
	return
}

func (MessageSortCard) messageType() MessageType {
	return MessageTypeSortCard
}

type MessageSelectUnselectCard struct {
	Player      int             `json:"player"`
	Finishable  bool            `json:"finishable"`
	Cancellable bool            `json:"cancellable"`
	Min         int             `json:"min"`
	Max         int             `json:"max"`
	Selects     []FieldCardInfo `json:"selects"`
	Unselects   []FieldCardInfo `json:"unselects"`
}

func ReadMessageSelectUnselectCard(b *bytes.Buffer) (msg MessageSelectUnselectCard) {
	msg.Player = int(readUint8(b))
	msg.Finishable = readUint8(b) != 0
	msg.Cancellable = readUint8(b) != 0
	msg.Min = int(readUint32(b))
	msg.Max = int(readUint32(b))

	selectsSize := readUint32(b)
	msg.Selects = make([]FieldCardInfo, selectsSize)
	for i := range msg.Selects {
		msg.Selects[i] = FieldCardInfo{
			Code:         int(readUint32(b)),
			CardLocation: parseCardLocation(readCardLocation(b)),
		}
	}
	unselectsSize := readUint32(b)
	msg.Unselects = make([]FieldCardInfo, unselectsSize)
	for i := range msg.Unselects {
		msg.Unselects[i] = FieldCardInfo{
			Code:         int(readUint32(b)),
			CardLocation: parseCardLocation(readCardLocation(b)),
		}
	}
	return
}

func (MessageSelectUnselectCard) messageType() MessageType {
	return MessageTypeSelectUnselectCard
}

type MessageConfirmDeckTop struct {
	Player int        `json:"player"`
	Cards  []CardInfo `json:"cards"`
}

func ReadMessageConfirmDeckTop(b *bytes.Buffer) (msg MessageConfirmDeckTop) {
	msg.Player = int(readUint8(b))
	cardsSize := readUint32(b)
	msg.Cards = make([]CardInfo, cardsSize)
	for i := range msg.Cards {
		msg.Cards[i] = CardInfo{
			Code:       int(readUint32(b)),
			Controller: int(readUint8(b)),
			Location:   parseCoreLocation(coreLocation(readUint8(b))),
			Sequence:   int(readUint32(b)),
		}
	}
	return
}

func (MessageConfirmDeckTop) messageType() MessageType {
	return MessageTypeConfirmDeckTop
}

type MessageConfirmCards struct {
	Player int        `json:"player"`
	Cards  []CardInfo `json:"cards"`
}

func ReadMessageConfirmCards(b *bytes.Buffer) (msg MessageConfirmCards) {
	msg.Player = int(readUint8(b))
	cardsSize := readUint32(b)
	msg.Cards = make([]CardInfo, cardsSize)
	for i := range msg.Cards {
		msg.Cards[i] = CardInfo{
			Code:       int(readUint32(b)),
			Controller: int(readUint8(b)),
			Location:   parseCoreLocation(coreLocation(readUint8(b))),
			Sequence:   int(readUint32(b)),
		}
	}
	return
}

func (MessageConfirmCards) messageType() MessageType {
	return MessageTypeConfirmCards
}

type MessageShuffleDeck struct {
	Player int `json:"player"`
}

func ReadMessageShuffleDeck(b *bytes.Buffer) (msg MessageShuffleDeck) {
	msg.Player = int(readUint8(b))
	return
}

func (MessageShuffleDeck) messageType() MessageType {
	return MessageTypeShuffleDeck
}

type MessageShuffleHand struct {
}

func ReadMessageShuffleHand(b *bytes.Buffer) (msg MessageShuffleHand) {
	// TODO: implement
	panic("not implemented")
}

func (MessageShuffleHand) messageType() MessageType {
	return MessageTypeShuffleHand
}

type MessageRefreshDeck struct {
}

func ReadMessageRefreshDeck(b *bytes.Buffer) (msg MessageRefreshDeck) {
	// TODO: implement
	panic("not implemented")
}

func (MessageRefreshDeck) messageType() MessageType {
	return MessageTypeRefreshDeck
}

type MessageSwapGraveDeck struct {
}

func ReadMessageSwapGraveDeck(b *bytes.Buffer) (msg MessageSwapGraveDeck) {
	// TODO: implement
	panic("not implemented")
}

func (MessageSwapGraveDeck) messageType() MessageType {
	return MessageTypeSwapGraveDeck
}

type MessageShuffleSetCard struct {
}

func ReadMessageShuffleSetCard(b *bytes.Buffer) (msg MessageShuffleSetCard) {
	// TODO: implement
	panic("not implemented")
}

func (MessageShuffleSetCard) messageType() MessageType {
	return MessageTypeShuffleSetCard
}

type MessageReverseDeck struct {
}

func ReadMessageReverseDeck(b *bytes.Buffer) (msg MessageReverseDeck) {
	// TODO: implement
	panic("not implemented")
}

func (MessageReverseDeck) messageType() MessageType {
	return MessageTypeReverseDeck
}

type MessageDeckTop struct {
}

func ReadMessageDeckTop(b *bytes.Buffer) (msg MessageDeckTop) {
	// TODO: implement
	panic("not implemented")
}

func (MessageDeckTop) messageType() MessageType {
	return MessageTypeDeckTop
}

type MessageShuffleExtra struct {
}

func ReadMessageShuffleExtra(b *bytes.Buffer) (msg MessageShuffleExtra) {
	// TODO: implement
	panic("not implemented")
}

func (MessageShuffleExtra) messageType() MessageType {
	return MessageTypeShuffleExtra
}

type MessageNewTurn struct {
	Player int `json:"player"`
}

func ReadMessageNewTurn(b *bytes.Buffer) (msg MessageNewTurn) {
	msg.Player = int(readUint8(b))
	return
}

func (MessageNewTurn) messageType() MessageType {
	return MessageTypeNewTurn
}

type MessageNewPhase struct {
	Phase         Phase         `json:"phase"`
	DetailedPhase DetailedPhase `json:"detailed_phase"`
}

func ReadMessageNewPhase(b *bytes.Buffer) (msg MessageNewPhase) {
	phase := corePhase(readUint16(b))
	msg.DetailedPhase = parseCorePhaseDetailed(phase)
	msg.Phase = parseCorePhase(phase)
	return
}

func (MessageNewPhase) messageType() MessageType {
	return MessageTypeNewPhase
}

type MessageConfirmExtraTop struct {
}

func ReadMessageConfirmExtraTop(b *bytes.Buffer) (msg MessageConfirmExtraTop) {
	// TODO: implement
	panic("not implemented")
}

func (MessageConfirmExtraTop) messageType() MessageType {
	return MessageTypeConfirmExtraTop
}

type MessageMove struct {
	Card     FieldCardInfo `json:"card"`
	Previous CardLocation  `json:"previous"`
	Reason   uint32        `json:"reason"`
}

func ReadMessageMove(b *bytes.Buffer) (msg MessageMove) {
	msg.Card.Code = int(readUint32(b))
	msg.Card.CardLocation = parseCardLocation(readCardLocation(b))
	msg.Previous = parseCardLocation(readCardLocation(b))
	msg.Reason = readUint32(b)
	return
}

func (MessageMove) messageType() MessageType {
	return MessageTypeMove
}

type MessagePosChange struct {
}

func ReadMessagePosChange(b *bytes.Buffer) (msg MessagePosChange) {
	// TODO: implement
	panic("not implemented")
}

func (MessagePosChange) messageType() MessageType {
	return MessageTypePosChange
}

type MessageSet struct {
}

func ReadMessageSet(b *bytes.Buffer) (msg MessageSet) {
	// TODO: implement
	panic("not implemented")
}

func (MessageSet) messageType() MessageType {
	return MessageTypeSet
}

type MessageSwap struct {
}

func ReadMessageSwap(b *bytes.Buffer) (msg MessageSwap) {
	// TODO: implement
	panic("not implemented")
}

func (MessageSwap) messageType() MessageType {
	return MessageTypeSwap
}

type MessageFieldDisabled struct {
}

func ReadMessageFieldDisabled(b *bytes.Buffer) (msg MessageFieldDisabled) {
	// TODO: implement
	panic("not implemented")
}

func (MessageFieldDisabled) messageType() MessageType {
	return MessageTypeFieldDisabled
}

type MessageSummoning struct {
	Card FieldCardInfo `json:"card"`
}

func ReadMessageSummoning(b *bytes.Buffer) (msg MessageSummoning) {
	msg.Card.Code = int(readUint32(b))
	msg.Card.CardLocation = parseCardLocation(readCardLocation(b))
	return
}

func (MessageSummoning) messageType() MessageType {
	return MessageTypeSummoning
}

type MessageSummoned struct{}

func ReadMessageSummoned(*bytes.Buffer) (msg MessageSummoned) {
	return
}

func (MessageSummoned) messageType() MessageType {
	return MessageTypeSummoned
}

type MessageSPSummoning struct {
	Card FieldCardInfo `json:"card"`
}

func ReadMessageSPSummoning(b *bytes.Buffer) (msg MessageSPSummoning) {
	msg.Card.Code = int(readUint32(b))
	msg.Card.CardLocation = parseCardLocation(readCardLocation(b))
	return
}

func (MessageSPSummoning) messageType() MessageType {
	return MessageTypeSPSummoning
}

type MessageSPSummoned struct{}

func ReadMessageSPSummoned(*bytes.Buffer) (msg MessageSPSummoned) {
	return
}

func (MessageSPSummoned) messageType() MessageType {
	return MessageTypeSPSummoned
}

type MessageFlipSummoning struct {
}

func ReadMessageFlipSummoning(b *bytes.Buffer) (msg MessageFlipSummoning) {
	// TODO: implement
	panic("not implemented")
}

func (MessageFlipSummoning) messageType() MessageType {
	return MessageTypeFlipSummoning
}

type MessageFlipSummoned struct {
}

func ReadMessageFlipSummoned(b *bytes.Buffer) (msg MessageFlipSummoned) {
	// TODO: implement
	panic("not implemented")
}

func (MessageFlipSummoned) messageType() MessageType {
	return MessageTypeFlipSummoned
}

type MessageChaining struct {
	Card              FieldCardInfo `json:"card"`
	TriggerController int           `json:"trigger_controller"`
	TriggerLocation   Location      `json:"trigger_location"`
	TriggerSequence   int           `json:"trigger_sequence"`
	Description       uint64        `json:"description"`
	Count             int           `json:"count"`
}

func ReadMessageChaining(b *bytes.Buffer) (msg MessageChaining) {
	msg.Card.Code = int(readUint32(b))
	msg.Card.CardLocation = parseCardLocation(readCardLocation(b))
	msg.TriggerController = int(readUint8(b))
	msg.TriggerLocation = parseCoreLocation(coreLocation(readUint8(b)))
	msg.TriggerSequence = int(readUint8(b))
	msg.Description = readUint64(b)
	msg.Count = int(readUint32(b))
	return
}

func (MessageChaining) messageType() MessageType {
	return MessageTypeChaining
}

type MessageChained struct {
	Count int `json:"count"`
}

func ReadMessageChained(b *bytes.Buffer) (msg MessageChained) {
	msg.Count = int(readUint8(b))
	return
}

func (MessageChained) messageType() MessageType {
	return MessageTypeChained
}

type MessageChainSolving struct {
	Count int `json:"count"`
}

func ReadMessageChainSolving(b *bytes.Buffer) (msg MessageChainSolving) {
	msg.Count = int(readUint8(b))
	return
}

func (MessageChainSolving) messageType() MessageType {
	return MessageTypeChainSolving
}

type MessageChainSolved struct {
	Count int `json:"count"`
}

func ReadMessageChainSolved(b *bytes.Buffer) (msg MessageChainSolved) {
	msg.Count = int(readUint8(b))
	return
}

func (MessageChainSolved) messageType() MessageType {
	return MessageTypeChainSolved
}

type MessageChainEnd struct{}

func ReadMessageChainEnd(*bytes.Buffer) (msg MessageChainEnd) {
	return
}

func (MessageChainEnd) messageType() MessageType {
	return MessageTypeChainEnd
}

type MessageChainNegated struct {
}

func ReadMessageChainNegated(b *bytes.Buffer) (msg MessageChainNegated) {
	// TODO: implement
	panic("not implemented")
}

func (MessageChainNegated) messageType() MessageType {
	return MessageTypeChainNegated
}

type MessageChainDisabled struct {
}

func ReadMessageChainDisabled(b *bytes.Buffer) (msg MessageChainDisabled) {
	// TODO: implement
	panic("not implemented")
}

func (MessageChainDisabled) messageType() MessageType {
	return MessageTypeChainDisabled
}

type MessageCardSelected struct {
}

func ReadMessageCardSelected(b *bytes.Buffer) (msg MessageCardSelected) {
	// TODO: implement
	panic("not implemented")
}

func (MessageCardSelected) messageType() MessageType {
	return MessageTypeCardSelected
}

type MessageRandomSelected struct {
}

func ReadMessageRandomSelected(b *bytes.Buffer) (msg MessageRandomSelected) {
	// TODO: implement
	panic("not implemented")
}

func (MessageRandomSelected) messageType() MessageType {
	return MessageTypeRandomSelected
}

type MessageBecomeTarget struct {
	Targets []CardLocation `json:"targets"`
}

func ReadMessageBecomeTarget(b *bytes.Buffer) (msg MessageBecomeTarget) {
	targetsLen := readUint32(b)
	msg.Targets = make([]CardLocation, targetsLen)
	for i := range msg.Targets {
		msg.Targets[i] = parseCardLocation(readCardLocation(b))
	}
	return
}

func (MessageBecomeTarget) messageType() MessageType {
	return MessageTypeBecomeTarget
}

type MessageDraw struct {
	Player int             `json:"player"`
	Cards  []DrawnCardInfo `json:"cards"`
}

func ReadMessageDraw(b *bytes.Buffer) (msg MessageDraw) {
	msg.Player = int(readUint8(b))
	cardsSize := readUint32(b)
	msg.Cards = make([]DrawnCardInfo, cardsSize)
	for i := range msg.Cards {
		msg.Cards[i] = DrawnCardInfo{
			Code:     int(readUint32(b)),
			Position: parseCorePosition(corePosition(readUint32(b))).Face(),
		}
	}
	return
}

func (MessageDraw) messageType() MessageType {
	return MessageTypeDraw
}

type MessageDamage struct {
}

func ReadMessageDamage(b *bytes.Buffer) (msg MessageDamage) {
	// TODO: implement
	panic("not implemented")
}

func (MessageDamage) messageType() MessageType {
	return MessageTypeDamage
}

type MessageRecover struct {
}

func ReadMessageRecover(b *bytes.Buffer) (msg MessageRecover) {
	// TODO: implement
	panic("not implemented")
}

func (MessageRecover) messageType() MessageType {
	return MessageTypeRecover
}

type MessageEquip struct {
}

func ReadMessageEquip(b *bytes.Buffer) (msg MessageEquip) {
	// TODO: implement
	panic("not implemented")
}

func (MessageEquip) messageType() MessageType {
	return MessageTypeEquip
}

type MessageLPUpdate struct {
}

func ReadMessageLPUpdate(b *bytes.Buffer) (msg MessageLPUpdate) {
	// TODO: implement
	panic("not implemented")
}

func (MessageLPUpdate) messageType() MessageType {
	return MessageTypeLPUpdate
}

type MessageUnequip struct {
}

func ReadMessageUnequip(b *bytes.Buffer) (msg MessageUnequip) {
	// TODO: implement
	panic("not implemented")
}

func (MessageUnequip) messageType() MessageType {
	return MessageTypeUnequip
}

type MessageCardTarget struct {
}

func ReadMessageCardTarget(b *bytes.Buffer) (msg MessageCardTarget) {
	// TODO: implement
	panic("not implemented")
}

func (MessageCardTarget) messageType() MessageType {
	return MessageTypeCardTarget
}

type MessageCancelTarget struct {
}

func ReadMessageCancelTarget(b *bytes.Buffer) (msg MessageCancelTarget) {
	// TODO: implement
	panic("not implemented")
}

func (MessageCancelTarget) messageType() MessageType {
	return MessageTypeCancelTarget
}

type MessagePayLPCost struct {
}

func ReadMessagePayLPCost(b *bytes.Buffer) (msg MessagePayLPCost) {
	// TODO: implement
	panic("not implemented")
}

func (MessagePayLPCost) messageType() MessageType {
	return MessageTypePayLPCost
}

type MessageAddCounter struct {
}

func ReadMessageAddCounter(b *bytes.Buffer) (msg MessageAddCounter) {
	// TODO: implement
	panic("not implemented")
}

func (MessageAddCounter) messageType() MessageType {
	return MessageTypeAddCounter
}

type MessageRemoveCounter struct {
}

func ReadMessageRemoveCounter(b *bytes.Buffer) (msg MessageRemoveCounter) {
	// TODO: implement
	panic("not implemented")
}

func (MessageRemoveCounter) messageType() MessageType {
	return MessageTypeRemoveCounter
}

type MessageAttack struct {
}

func ReadMessageAttack(b *bytes.Buffer) (msg MessageAttack) {
	// TODO: implement
	panic("not implemented")
}

func (MessageAttack) messageType() MessageType {
	return MessageTypeAttack
}

type MessageBattle struct {
}

func ReadMessageBattle(b *bytes.Buffer) (msg MessageBattle) {
	// TODO: implement
	panic("not implemented")
}

func (MessageBattle) messageType() MessageType {
	return MessageTypeBattle
}

type MessageAttackDisabled struct {
}

func ReadMessageAttackDisabled(b *bytes.Buffer) (msg MessageAttackDisabled) {
	// TODO: implement
	panic("not implemented")
}

func (MessageAttackDisabled) messageType() MessageType {
	return MessageTypeAttackDisabled
}

type MessageDamageStepStart struct {
}

func ReadMessageDamageStepStart(b *bytes.Buffer) (msg MessageDamageStepStart) {
	// TODO: implement
	panic("not implemented")
}

func (MessageDamageStepStart) messageType() MessageType {
	return MessageTypeDamageStepStart
}

type MessageDamageStepEnd struct {
}

func ReadMessageDamageStepEnd(b *bytes.Buffer) (msg MessageDamageStepEnd) {
	// TODO: implement
	panic("not implemented")
}

func (MessageDamageStepEnd) messageType() MessageType {
	return MessageTypeDamageStepEnd
}

type MessageMissedEffect struct {
}

func ReadMessageMissedEffect(b *bytes.Buffer) (msg MessageMissedEffect) {
	// TODO: implement
	panic("not implemented")
}

func (MessageMissedEffect) messageType() MessageType {
	return MessageTypeMissedEffect
}

type MessageBeChainTarget struct {
}

func ReadMessageBeChainTarget(b *bytes.Buffer) (msg MessageBeChainTarget) {
	// TODO: implement
	panic("not implemented")
}

func (MessageBeChainTarget) messageType() MessageType {
	return MessageTypeBeChainTarget
}

type MessageCreateRelation struct {
}

func ReadMessageCreateRelation(b *bytes.Buffer) (msg MessageCreateRelation) {
	// TODO: implement
	panic("not implemented")
}

func (MessageCreateRelation) messageType() MessageType {
	return MessageTypeCreateRelation
}

type MessageReleaseRelation struct {
}

func ReadMessageReleaseRelation(b *bytes.Buffer) (msg MessageReleaseRelation) {
	// TODO: implement
	panic("not implemented")
}

func (MessageReleaseRelation) messageType() MessageType {
	return MessageTypeReleaseRelation
}

type MessageTossCoin struct {
}

func ReadMessageTossCoin(b *bytes.Buffer) (msg MessageTossCoin) {
	// TODO: implement
	panic("not implemented")
}

func (MessageTossCoin) messageType() MessageType {
	return MessageTypeTossCoin
}

type MessageTossDice struct {
}

func ReadMessageTossDice(b *bytes.Buffer) (msg MessageTossDice) {
	// TODO: implement
	panic("not implemented")
}

func (MessageTossDice) messageType() MessageType {
	return MessageTypeTossDice
}

type MessageRockPaperScissors struct {
}

func ReadMessageRockPaperScissors(b *bytes.Buffer) (msg MessageRockPaperScissors) {
	// TODO: implement
	panic("not implemented")
}

func (MessageRockPaperScissors) messageType() MessageType {
	return MessageTypeRockPaperScissors
}

type MessageHandRes struct {
}

func ReadMessageHandRes(b *bytes.Buffer) (msg MessageHandRes) {
	// TODO: implement
	panic("not implemented")
}

func (MessageHandRes) messageType() MessageType {
	return MessageTypeHandRes
}

type MessageAnnounceRace struct {
}

func ReadMessageAnnounceRace(b *bytes.Buffer) (msg MessageAnnounceRace) {
	// TODO: implement
	panic("not implemented")
}

func (MessageAnnounceRace) messageType() MessageType {
	return MessageTypeAnnounceRace
}

type MessageAnnounceAttribute struct {
}

func ReadMessageAnnounceAttribute(b *bytes.Buffer) (msg MessageAnnounceAttribute) {
	// TODO: implement
	panic("not implemented")
}

func (MessageAnnounceAttribute) messageType() MessageType {
	return MessageTypeAnnounceAttribute
}

type MessageAnnounceCard struct {
}

func ReadMessageAnnounceCard(b *bytes.Buffer) (msg MessageAnnounceCard) {
	// TODO: implement
	panic("not implemented")
}

func (MessageAnnounceCard) messageType() MessageType {
	return MessageTypeAnnounceCard
}

type MessageAnnounceNumber struct {
}

func ReadMessageAnnounceNumber(b *bytes.Buffer) (msg MessageAnnounceNumber) {
	// TODO: implement
	panic("not implemented")
}

func (MessageAnnounceNumber) messageType() MessageType {
	return MessageTypeAnnounceNumber
}

type MessageCardHint struct {
}

func ReadMessageCardHint(b *bytes.Buffer) (msg MessageCardHint) {
	// TODO: implement
	panic("not implemented")
}

func (MessageCardHint) messageType() MessageType {
	return MessageTypeCardHint
}

type MessageTagSwap struct {
}

func ReadMessageTagSwap(b *bytes.Buffer) (msg MessageTagSwap) {
	// TODO: implement
	panic("not implemented")
}

func (MessageTagSwap) messageType() MessageType {
	return MessageTypeTagSwap
}

type MessageReloadField struct {
}

func ReadMessageReloadField(b *bytes.Buffer) (msg MessageReloadField) {
	// TODO: implement
	panic("not implemented")
}

func (MessageReloadField) messageType() MessageType {
	return MessageTypeReloadField
}

type MessageAIName struct {
}

func ReadMessageAIName(b *bytes.Buffer) (msg MessageAIName) {
	// TODO: implement
	panic("not implemented")
}

func (MessageAIName) messageType() MessageType {
	return MessageTypeAIName
}

type MessageShowHint struct {
}

func ReadMessageShowHint(b *bytes.Buffer) (msg MessageShowHint) {
	// TODO: implement
	panic("not implemented")
}

func (MessageShowHint) messageType() MessageType {
	return MessageTypeShowHint
}

type MessagePlayerHint struct {
}

func ReadMessagePlayerHint(b *bytes.Buffer) (msg MessagePlayerHint) {
	// TODO: implement
	panic("not implemented")
}

func (MessagePlayerHint) messageType() MessageType {
	return MessageTypePlayerHint
}

type MessageMatchKill struct {
}

func ReadMessageMatchKill(b *bytes.Buffer) (msg MessageMatchKill) {
	// TODO: implement
	panic("not implemented")
}

func (MessageMatchKill) messageType() MessageType {
	return MessageTypeMatchKill
}

type MessageCustomMessage struct {
}

func ReadMessageCustomMessage(b *bytes.Buffer) (msg MessageCustomMessage) {
	// TODO: implement
	panic("not implemented")
}

func (MessageCustomMessage) messageType() MessageType {
	return MessageTypeCustomMessage
}

type MessageRemoveCards struct {
}

func ReadMessageRemoveCards(b *bytes.Buffer) (msg MessageRemoveCards) {
	// TODO: implement
	panic("not implemented")
}

func (MessageRemoveCards) messageType() MessageType {
	return MessageTypeRemoveCards
}

type jsonMessageName struct {
	MessageType MessageType `json:"message_type"`
}

func MessageToJSON(m Message) ([]byte, error) {
	structTypesCacheLock.Lock()
	defer structTypesCacheLock.Unlock()

	v := createStructTypesCache(m.messageType(), reflect.TypeOf(m))
	v.Field(0).Set(reflect.ValueOf(m))
	return json.Marshal(v.Interface())
}

func JSONToMessage(b []byte) (Message, error) {
	structTypesCacheLock.Lock()
	defer structTypesCacheLock.Unlock()

	var m jsonMessageName
	err := json.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}
	typ, ok := messageUnmarshalMap[m.MessageType]
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

var messageUnmarshalMap = map[MessageType]reflect.Type{
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
