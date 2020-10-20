package ocgcore

import (
	"bytes"
	"fmt"
	"ocgcore/lib"
	"ocgcore/utils"
)

//go:generate go run ocgcore/cmd/enumer -type=MessageType -json -transform=snake -output message_enumer.go -trimprefix MessageType
//go:generate go run ocgcore/cmd/interfacer -method=messageType -returns=MessageType -field=message_type -interface Message -output message_interfacer.go

type cardLocation struct {
	controller int
	location   lib.Location
	sequence   int
	position   lib.Position
}

func readMessage(contents []byte) Message {
	b := bytes.NewBuffer(contents)

	id := utils.ReadUint8(b)

	switch lib.Message(id) {
	case lib.MessageRetry:
		return ReadMessageRetry(b)
	case lib.MessageHint:
		return ReadMessageHint(b)
	case lib.MessageWaiting:
		return ReadMessageWaiting(b)
	case lib.MessageStart:
		return ReadMessageStart(b)
	case lib.MessageWin:
		return ReadMessageWin(b)
	case lib.MessageUpdateData:
		return ReadMessageUpdateData(b)
	case lib.MessageUpdateCard:
		return ReadMessageUpdateCard(b)
	case lib.MessageRequestDeck:
		return ReadMessageRequestDeck(b)
	case lib.MessageSelectBattleCMD:
		return ReadMessageSelectBattleCMD(b)
	case lib.MessageSelectIdleCMD:
		return ReadMessageSelectIdleCMD(b)
	case lib.MessageSelectEffectYN:
		return ReadMessageSelectEffectYN(b)
	case lib.MessageSelectYesNo:
		return ReadMessageSelectYesNo(b)
	case lib.MessageSelectOption:
		return ReadMessageSelectOption(b)
	case lib.MessageSelectCard:
		return ReadMessageSelectCard(b)
	case lib.MessageSelectChain:
		return ReadMessageSelectChain(b)
	case lib.MessageSelectPlace:
		return ReadMessageSelectPlace(b)
	case lib.MessageSelectPosition:
		return ReadMessageSelectPosition(b)
	case lib.MessageSelectTribute:
		return ReadMessageSelectTribute(b)
	case lib.MessageSortChain:
		return ReadMessageSortChain(b)
	case lib.MessageSelectCounter:
		return ReadMessageSelectCounter(b)
	case lib.MessageSelectSum:
		return ReadMessageSelectSum(b)
	case lib.MessageSelectDisfield:
		return ReadMessageSelectDisfield(b)
	case lib.MessageSortCard:
		return ReadMessageSortCard(b)
	case lib.MessageSelectUnselectCard:
		return ReadMessageSelectUnselectCard(b)
	case lib.MessageConfirmDeckTop:
		return ReadMessageConfirmDeckTop(b)
	case lib.MessageConfirmCards:
		return ReadMessageConfirmCards(b)
	case lib.MessageShuffleDeck:
		return ReadMessageShuffleDeck(b)
	case lib.MessageShuffleHand:
		return ReadMessageShuffleHand(b)
	case lib.MessageRefreshDeck:
		return ReadMessageRefreshDeck(b)
	case lib.MessageSwapGraveDeck:
		return ReadMessageSwapGraveDeck(b)
	case lib.MessageShuffleSetCard:
		return ReadMessageShuffleSetCard(b)
	case lib.MessageReverseDeck:
		return ReadMessageReverseDeck(b)
	case lib.MessageDeckTop:
		return ReadMessageDeckTop(b)
	case lib.MessageShuffleExtra:
		return ReadMessageShuffleExtra(b)
	case lib.MessageNewTurn:
		return ReadMessageNewTurn(b)
	case lib.MessageNewPhase:
		return ReadMessageNewPhase(b)
	case lib.MessageConfirmExtraTop:
		return ReadMessageConfirmExtraTop(b)
	case lib.MessageMove:
		return ReadMessageMove(b)
	case lib.MessagePosChange:
		return ReadMessagePosChange(b)
	case lib.MessageSet:
		return ReadMessageSet(b)
	case lib.MessageSwap:
		return ReadMessageSwap(b)
	case lib.MessageFieldDisabled:
		return ReadMessageFieldDisabled(b)
	case lib.MessageSummoning:
		return ReadMessageSummoning(b)
	case lib.MessageSummoned:
		return ReadMessageSummoned(b)
	case lib.MessageSPSummoning:
		return ReadMessageSPSummoning(b)
	case lib.MessageSPSummoned:
		return ReadMessageSPSummoned(b)
	case lib.MessageFlipSummoning:
		return ReadMessageFlipSummoning(b)
	case lib.MessageFlipSummoned:
		return ReadMessageFlipSummoned(b)
	case lib.MessageChaining:
		return ReadMessageChaining(b)
	case lib.MessageChained:
		return ReadMessageChained(b)
	case lib.MessageChainSolving:
		return ReadMessageChainSolving(b)
	case lib.MessageChainSolved:
		return ReadMessageChainSolved(b)
	case lib.MessageChainEnd:
		return ReadMessageChainEnd(b)
	case lib.MessageChainNegated:
		return ReadMessageChainNegated(b)
	case lib.MessageChainDisabled:
		return ReadMessageChainDisabled(b)
	case lib.MessageCardSelected:
		return ReadMessageCardSelected(b)
	case lib.MessageRandomSelected:
		return ReadMessageRandomSelected(b)
	case lib.MessageBecomeTarget:
		return ReadMessageBecomeTarget(b)
	case lib.MessageDraw:
		return ReadMessageDraw(b)
	case lib.MessageDamage:
		return ReadMessageDamage(b)
	case lib.MessageRecover:
		return ReadMessageRecover(b)
	case lib.MessageEquip:
		return ReadMessageEquip(b)
	case lib.MessageLPUpdate:
		return ReadMessageLPUpdate(b)
	case lib.MessageUnequip:
		return ReadMessageUnequip(b)
	case lib.MessageCardTarget:
		return ReadMessageCardTarget(b)
	case lib.MessageCancelTarget:
		return ReadMessageCancelTarget(b)
	case lib.MessagePayLPCost:
		return ReadMessagePayLPCost(b)
	case lib.MessageAddCounter:
		return ReadMessageAddCounter(b)
	case lib.MessageRemoveCounter:
		return ReadMessageRemoveCounter(b)
	case lib.MessageAttack:
		return ReadMessageAttack(b)
	case lib.MessageBattle:
		return ReadMessageBattle(b)
	case lib.MessageAttackDisabled:
		return ReadMessageAttackDisabled(b)
	case lib.MessageDamageStepStart:
		return ReadMessageDamageStepStart(b)
	case lib.MessageDamageStepEnd:
		return ReadMessageDamageStepEnd(b)
	case lib.MessageMissedEffect:
		return ReadMessageMissedEffect(b)
	case lib.MessageBeChainTarget:
		return ReadMessageBeChainTarget(b)
	case lib.MessageCreateRelation:
		return ReadMessageCreateRelation(b)
	case lib.MessageReleaseRelation:
		return ReadMessageReleaseRelation(b)
	case lib.MessageTossCoin:
		return ReadMessageTossCoin(b)
	case lib.MessageTossDice:
		return ReadMessageTossDice(b)
	case lib.MessageRockPaperScissors:
		return ReadMessageRockPaperScissors(b)
	case lib.MessageHandRes:
		return ReadMessageHandRes(b)
	case lib.MessageAnnounceRace:
		return ReadMessageAnnounceRace(b)
	case lib.MessageAnnounceAttribute:
		return ReadMessageAnnounceAttribute(b)
	case lib.MessageAnnounceCard:
		return ReadMessageAnnounceCard(b)
	case lib.MessageAnnounceNumber:
		return ReadMessageAnnounceNumber(b)
	case lib.MessageCardHint:
		return ReadMessageCardHint(b)
	case lib.MessageTagSwap:
		return ReadMessageTagSwap(b)
	case lib.MessageReloadField:
		return ReadMessageReloadField(b)
	case lib.MessageAIName:
		return ReadMessageAIName(b)
	case lib.MessageShowHint:
		return ReadMessageShowHint(b)
	case lib.MessagePlayerHint:
		return ReadMessagePlayerHint(b)
	case lib.MessageMatchKill:
		return ReadMessageMatchKill(b)
	case lib.MessageCustomMessage:
		return ReadMessageCustomMessage(b)
	case lib.MessageRemoveCards:
		return ReadMessageRemoveCards(b)
	default:
		fmt.Println("unhandled message:", id, "size:", len(contents)-1)
		return nil
	}
}

func readCardLocation(b *bytes.Buffer) cardLocation {
	return cardLocation{
		controller: int(utils.ReadUint8(b)),
		location:   lib.Location(utils.ReadUint8(b)),
		sequence:   int(utils.ReadUint32(b)),
		position:   lib.Position(utils.ReadUint32(b)),
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
	msg.Hint = int(utils.ReadUint8(b))
	msg.Player = int(utils.ReadUint8(b))
	msg.Desc = utils.ReadUint64(b)
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
	msg.Player = int(utils.ReadUint8(b))
	msg.Reason = int(utils.ReadUint8(b))
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
	msg.Player = int(utils.ReadUint8(b))

	selectChainsSize := utils.ReadUint32(b)
	msg.Chains = make([]ChainInfo, selectChainsSize)
	for i := range msg.Chains {
		msg.Chains[i] = ChainInfo{
			Code:        int(utils.ReadUint32(b)),
			Controller:  int(utils.ReadUint8(b)),
			Location:    parseCoreLocation(lib.Location(utils.ReadUint8(b))),
			Sequence:    int(utils.ReadUint32(b)),
			Description: utils.ReadUint64(b),
			ClientMode:  utils.ReadUint8(b),
		}
	}

	attackableSize := utils.ReadUint32(b)
	msg.Attacks = make([]AttackInfo, attackableSize)
	for i := range msg.Attacks {
		msg.Attacks[i] = AttackInfo{
			Code:       int(utils.ReadUint32(b)),
			Controller: int(utils.ReadUint8(b)),
			Location:   parseCoreLocation(lib.Location(utils.ReadUint8(b))),
			Sequence:   int(utils.ReadUint32(b)),
			Direct:     utils.ReadUint8(b) != 0,
		}
	}
	msg.ToM2 = utils.ReadUint8(b) != 0
	msg.ToEP = utils.ReadUint8(b) != 0
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
	msg.Player = int(utils.ReadUint8(b))

	summonableSize := utils.ReadUint32(b)
	msg.Summons = make([]CardInfo, summonableSize)
	for i := range msg.Summons {
		msg.Summons[i] = CardInfo{
			Code:       int(utils.ReadUint32(b)),
			Controller: int(utils.ReadUint8(b)),
			Location:   parseCoreLocation(lib.Location(utils.ReadUint8(b))),
			Sequence:   int(utils.ReadUint32(b)),
		}
	}

	spSummonableSize := utils.ReadUint32(b)
	msg.SpSummons = make([]CardInfo, spSummonableSize)
	for i := range msg.SpSummons {
		msg.SpSummons[i] = CardInfo{
			Code:       int(utils.ReadUint32(b)),
			Controller: int(utils.ReadUint8(b)),
			Location:   parseCoreLocation(lib.Location(utils.ReadUint8(b))),
			Sequence:   int(utils.ReadUint32(b)),
		}
	}

	posChangeSize := utils.ReadUint32(b)
	msg.PosChanges = make([]CardInfo, posChangeSize)
	for i := range msg.PosChanges {
		msg.PosChanges[i] = CardInfo{
			Code:       int(utils.ReadUint32(b)),
			Controller: int(utils.ReadUint8(b)),
			Location:   parseCoreLocation(lib.Location(utils.ReadUint8(b))),
			Sequence:   int(utils.ReadUint32(b)),
		}
	}

	monsterSetSize := utils.ReadUint32(b)
	msg.MonsterSets = make([]CardInfo, monsterSetSize)
	for i := range msg.MonsterSets {
		msg.MonsterSets[i] = CardInfo{
			Code:       int(utils.ReadUint32(b)),
			Controller: int(utils.ReadUint8(b)),
			Location:   parseCoreLocation(lib.Location(utils.ReadUint8(b))),
			Sequence:   int(utils.ReadUint32(b)),
		}
	}

	spellSetSize := utils.ReadUint32(b)
	msg.SpellSets = make([]CardInfo, spellSetSize)
	for i := range msg.SpellSets {
		msg.SpellSets[i] = CardInfo{
			Code:       int(utils.ReadUint32(b)),
			Controller: int(utils.ReadUint8(b)),
			Location:   parseCoreLocation(lib.Location(utils.ReadUint8(b))),
			Sequence:   int(utils.ReadUint32(b)),
		}
	}

	activateSize := utils.ReadUint32(b)
	msg.Activate = make([]ChainInfo, activateSize)
	for i := range msg.Activate {
		msg.Activate[i] = ChainInfo{
			Code:        int(utils.ReadUint32(b)),
			Controller:  int(utils.ReadUint8(b)),
			Location:    parseCoreLocation(lib.Location(utils.ReadUint8(b))),
			Sequence:    int(utils.ReadUint32(b)),
			Description: utils.ReadUint64(b),
			ClientMode:  utils.ReadUint8(b),
		}
	}

	msg.ToBP = utils.ReadUint8(b) != 0
	msg.ToEP = utils.ReadUint8(b) != 0
	msg.Shuffle = utils.ReadUint8(b) != 0
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
	msg.Player = int(utils.ReadUint8(b))
	msg.Code = utils.ReadUint32(b)
	loc := readCardLocation(b)
	msg.Controller = loc.controller
	msg.Location = parseCoreLocation(loc.location)
	msg.Sequence = loc.sequence
	msg.Position = parseCorePosition(loc.position)
	msg.Description = utils.ReadUint64(b)
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
	msg.Player = int(utils.ReadUint8(b))
	msg.Description = utils.ReadUint64(b)
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
	msg.Player = int(utils.ReadUint8(b))

	optionsSize := utils.ReadUint8(b)
	msg.Options = make([]uint64, optionsSize)
	for i := range msg.Options {
		msg.Options[i] = utils.ReadUint64(b)
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
	msg.Player = int(utils.ReadUint8(b))
	msg.Cancellable = utils.ReadUint8(b) != 0
	msg.Min = int(utils.ReadUint32(b))
	msg.Max = int(utils.ReadUint32(b))

	cardsSize := utils.ReadUint32(b)
	msg.Cards = make([]FieldCardInfo, cardsSize)
	for i := range msg.Cards {
		msg.Cards[i] = FieldCardInfo{
			Code:         int(utils.ReadUint32(b)),
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
	msg.Player = int(utils.ReadUint8(b))
	msg.SpeCount = int(utils.ReadUint8(b))
	msg.Forced = utils.ReadUint8(b) != 0
	msg.HintTimingPlayer = utils.ReadUint32(b)
	msg.HintTimingOther = utils.ReadUint32(b)

	chainsSize := utils.ReadUint32(b)
	msg.Chains = make([]CardChainInfo, chainsSize)
	for i := range msg.Chains {
		msg.Chains[i] = CardChainInfo{
			Code:         int(utils.ReadUint32(b)),
			CardLocation: parseCardLocation(readCardLocation(b)),
			Description:  utils.ReadUint64(b),
			ClientMode:   utils.ReadUint8(b),
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
	msg.Player = int(utils.ReadUint8(b))
	msg.Count = int(utils.ReadUint8(b))
	msg.Places = parsePlaceFlag(utils.ReadUint32(b))
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
	msg.Player = int(utils.ReadUint8(b))
	msg.Code = utils.ReadUint32(b)
	msg.Positions = parseCorePositions(lib.Position(utils.ReadUint8(b)))
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
	msg.Player = int(utils.ReadUint8(b))
	msg.Cancellable = utils.ReadUint8(b) != 0
	msg.Min = int(utils.ReadUint32(b))
	msg.Max = int(utils.ReadUint32(b))

	tributeSize := utils.ReadUint32(b)
	msg.Cards = make([]TributeCardInfo, tributeSize)
	for i := range msg.Cards {
		msg.Cards[i] = TributeCardInfo{
			Code:         int(utils.ReadUint32(b)),
			Controller:   int(utils.ReadUint8(b)),
			Location:     parseCoreLocation(lib.Location(utils.ReadUint8(b))),
			Sequence:     int(utils.ReadUint32(b)),
			ReleaseParam: int(utils.ReadUint8(b)),
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
	msg.Player = int(utils.ReadUint8(b))
	cardsSize := utils.ReadUint32(b)
	msg.Cards = make([]CardInfo, cardsSize)
	for i := range msg.Cards {
		msg.Cards[i] = CardInfo{
			Code:       int(utils.ReadUint32(b)),
			Controller: int(utils.ReadUint8(b)),
			Location:   parseCoreLocation(lib.Location(utils.ReadUint8(b))),
			Sequence:   int(utils.ReadUint32(b)),
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
	msg.Player = int(utils.ReadUint8(b))
	msg.CounterType = int(utils.ReadUint16(b))
	msg.Count = int(utils.ReadUint16(b))
	cardsSize := utils.ReadUint32(b)
	msg.Cards = make([]CounterCardInfo, cardsSize)
	for i := range msg.Cards {
		msg.Cards[i] = CounterCardInfo{
			Code:       int(utils.ReadUint32(b)),
			Controller: int(utils.ReadUint8(b)),
			Location:   parseCoreLocation(lib.Location(utils.ReadUint8(b))),
			Sequence:   int(utils.ReadUint8(b)),
			Count:      int(utils.ReadUint16(b)),
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
	msg.Player = int(utils.ReadUint8(b))
	msg.HasMax = utils.ReadUint8(b) != 0
	msg.Acc = int(utils.ReadUint32(b))
	msg.Min = int(utils.ReadUint32(b))
	msg.Max = int(utils.ReadUint32(b))
	mustSelectsSize := utils.ReadUint32(b)
	msg.MustSelects = make([]CounterCardInfo, mustSelectsSize)
	for i := range msg.MustSelects {
		msg.MustSelects[i] = CounterCardInfo{
			Code:       int(utils.ReadUint32(b)),
			Controller: int(utils.ReadUint8(b)),
			Location:   parseCoreLocation(lib.Location(utils.ReadUint8(b))),
			Sequence:   int(utils.ReadUint32(b)),
			Count:      int(utils.ReadUint32(b)),
		}
	}
	selectsSize := utils.ReadUint32(b)
	msg.Selects = make([]CounterCardInfo, selectsSize)
	for i := range msg.Selects {
		msg.Selects[i] = CounterCardInfo{
			Code:       int(utils.ReadUint32(b)),
			Controller: int(utils.ReadUint8(b)),
			Location:   parseCoreLocation(lib.Location(utils.ReadUint8(b))),
			Sequence:   int(utils.ReadUint32(b)),
			Count:      int(utils.ReadUint32(b)),
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
	msg.Player = int(utils.ReadUint8(b))
	msg.Count = int(utils.ReadUint8(b))
	msg.Flag = utils.ReadUint32(b)
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
	msg.Player = int(utils.ReadUint8(b))
	cardsSize := utils.ReadUint32(b)
	msg.Cards = make([]CardInfo, cardsSize)
	for i := range msg.Cards {
		msg.Cards[i] = CardInfo{
			Code:       int(utils.ReadUint32(b)),
			Controller: int(utils.ReadUint8(b)),
			Location:   parseCoreLocation(lib.Location(utils.ReadUint8(b))),
			Sequence:   int(utils.ReadUint32(b)),
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
	msg.Player = int(utils.ReadUint8(b))
	msg.Finishable = utils.ReadUint8(b) != 0
	msg.Cancellable = utils.ReadUint8(b) != 0
	msg.Min = int(utils.ReadUint32(b))
	msg.Max = int(utils.ReadUint32(b))

	selectsSize := utils.ReadUint32(b)
	msg.Selects = make([]FieldCardInfo, selectsSize)
	for i := range msg.Selects {
		msg.Selects[i] = FieldCardInfo{
			Code:         int(utils.ReadUint32(b)),
			CardLocation: parseCardLocation(readCardLocation(b)),
		}
	}
	unselectsSize := utils.ReadUint32(b)
	msg.Unselects = make([]FieldCardInfo, unselectsSize)
	for i := range msg.Unselects {
		msg.Unselects[i] = FieldCardInfo{
			Code:         int(utils.ReadUint32(b)),
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
	msg.Player = int(utils.ReadUint8(b))
	cardsSize := utils.ReadUint32(b)
	msg.Cards = make([]CardInfo, cardsSize)
	for i := range msg.Cards {
		msg.Cards[i] = CardInfo{
			Code:       int(utils.ReadUint32(b)),
			Controller: int(utils.ReadUint8(b)),
			Location:   parseCoreLocation(lib.Location(utils.ReadUint8(b))),
			Sequence:   int(utils.ReadUint32(b)),
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
	msg.Player = int(utils.ReadUint8(b))
	cardsSize := utils.ReadUint32(b)
	msg.Cards = make([]CardInfo, cardsSize)
	for i := range msg.Cards {
		msg.Cards[i] = CardInfo{
			Code:       int(utils.ReadUint32(b)),
			Controller: int(utils.ReadUint8(b)),
			Location:   parseCoreLocation(lib.Location(utils.ReadUint8(b))),
			Sequence:   int(utils.ReadUint32(b)),
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
	msg.Player = int(utils.ReadUint8(b))
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
	msg.Player = int(utils.ReadUint8(b))
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
	phase := lib.Phase(utils.ReadUint16(b))
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
	msg.Card.Code = int(utils.ReadUint32(b))
	msg.Card.CardLocation = parseCardLocation(readCardLocation(b))
	msg.Previous = parseCardLocation(readCardLocation(b))
	msg.Reason = utils.ReadUint32(b)
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
	msg.Card.Code = int(utils.ReadUint32(b))
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
	msg.Card.Code = int(utils.ReadUint32(b))
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
	msg.Card.Code = int(utils.ReadUint32(b))
	msg.Card.CardLocation = parseCardLocation(readCardLocation(b))
	msg.TriggerController = int(utils.ReadUint8(b))
	msg.TriggerLocation = parseCoreLocation(lib.Location(utils.ReadUint8(b)))
	msg.TriggerSequence = int(utils.ReadUint8(b))
	msg.Description = utils.ReadUint64(b)
	msg.Count = int(utils.ReadUint32(b))
	return
}

func (MessageChaining) messageType() MessageType {
	return MessageTypeChaining
}

type MessageChained struct {
	Count int `json:"count"`
}

func ReadMessageChained(b *bytes.Buffer) (msg MessageChained) {
	msg.Count = int(utils.ReadUint8(b))
	return
}

func (MessageChained) messageType() MessageType {
	return MessageTypeChained
}

type MessageChainSolving struct {
	Count int `json:"count"`
}

func ReadMessageChainSolving(b *bytes.Buffer) (msg MessageChainSolving) {
	msg.Count = int(utils.ReadUint8(b))
	return
}

func (MessageChainSolving) messageType() MessageType {
	return MessageTypeChainSolving
}

type MessageChainSolved struct {
	Count int `json:"count"`
}

func ReadMessageChainSolved(b *bytes.Buffer) (msg MessageChainSolved) {
	msg.Count = int(utils.ReadUint8(b))
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
	targetsLen := utils.ReadUint32(b)
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
	msg.Player = int(utils.ReadUint8(b))
	cardsSize := utils.ReadUint32(b)
	msg.Cards = make([]DrawnCardInfo, cardsSize)
	for i := range msg.Cards {
		msg.Cards[i] = DrawnCardInfo{
			Code:     int(utils.ReadUint32(b)),
			Position: parseCorePosition(lib.Position(utils.ReadUint32(b))).Face(),
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
