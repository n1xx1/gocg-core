package ocgcore

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"reflect"
	"sync"
)

type cardLocation struct {
	controller int
	location   coreLocation
	sequence   int
	position   corePosition
}

func ReadMessage(contents []byte) Message {
	b := bytes.NewBuffer(contents)

	var id uint8
	_ = binary.Read(b, binary.LittleEndian, &id)

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
func readUint8(b *bytes.Buffer) uint8 {
	var value uint8
	_ = binary.Read(b, binary.LittleEndian, &value)
	return value
}
func readUint16(b *bytes.Buffer) uint16 {
	var value uint16
	_ = binary.Read(b, binary.LittleEndian, &value)
	return value
}
func readUint32(b *bytes.Buffer) uint32 {
	var value uint32
	_ = binary.Read(b, binary.LittleEndian, &value)
	return value
}
func readUint64(b *bytes.Buffer) uint64 {
	var value uint64
	_ = binary.Read(b, binary.LittleEndian, &value)
	return value
}
func readInt8(b *bytes.Buffer) int8 {
	var value int8
	_ = binary.Read(b, binary.LittleEndian, &value)
	return value
}
func readInt16(b *bytes.Buffer) int16 {
	var value int16
	_ = binary.Read(b, binary.LittleEndian, &value)
	return value
}
func readInt32(b *bytes.Buffer) int32 {
	var value int32
	_ = binary.Read(b, binary.LittleEndian, &value)
	return value
}
func readInt64(b *bytes.Buffer) int64 {
	var value int64
	_ = binary.Read(b, binary.LittleEndian, &value)
	return value
}

type MessageType uint8

const (
	MessageTypeRetry MessageType = iota
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

var messageTypeNames = map[MessageType]string{
	MessageTypeRetry:              "retry",
	MessageTypeHint:               "hint",
	MessageTypeWaiting:            "waiting",
	MessageTypeStart:              "start",
	MessageTypeWin:                "win",
	MessageTypeUpdateData:         "update_data",
	MessageTypeUpdateCard:         "update_card",
	MessageTypeRequestDeck:        "request_deck",
	MessageTypeSelectBattleCMD:    "select_battlecmd",
	MessageTypeSelectIdleCMD:      "select_idlecmd",
	MessageTypeSelectEffectYN:     "select_effectyn",
	MessageTypeSelectYesNo:        "select_yesno",
	MessageTypeSelectOption:       "select_option",
	MessageTypeSelectCard:         "select_card",
	MessageTypeSelectChain:        "select_chain",
	MessageTypeSelectPlace:        "select_place",
	MessageTypeSelectPosition:     "select_position",
	MessageTypeSelectTribute:      "select_tribute",
	MessageTypeSortChain:          "sort_chain",
	MessageTypeSelectCounter:      "select_counter",
	MessageTypeSelectSum:          "select_sum",
	MessageTypeSelectDisfield:     "select_disfield",
	MessageTypeSortCard:           "sort_card",
	MessageTypeSelectUnselectCard: "select_unselect_card",
	MessageTypeConfirmDeckTop:     "confirm_decktop",
	MessageTypeConfirmCards:       "confirm_cards",
	MessageTypeShuffleDeck:        "shuffle_deck",
	MessageTypeShuffleHand:        "shuffle_hand",
	MessageTypeRefreshDeck:        "refresh_deck",
	MessageTypeSwapGraveDeck:      "swap_grave_deck",
	MessageTypeShuffleSetCard:     "shuffle_set_card",
	MessageTypeReverseDeck:        "reverse_deck",
	MessageTypeDeckTop:            "deck_top",
	MessageTypeShuffleExtra:       "shuffle_extra",
	MessageTypeNewTurn:            "new_turn",
	MessageTypeNewPhase:           "new_phase",
	MessageTypeConfirmExtraTop:    "confirm_extratop",
	MessageTypeMove:               "move",
	MessageTypePosChange:          "pos_change",
	MessageTypeSet:                "set",
	MessageTypeSwap:               "swap",
	MessageTypeFieldDisabled:      "field_disabled",
	MessageTypeSummoning:          "summoning",
	MessageTypeSummoned:           "summoned",
	MessageTypeSPSummoning:        "spsummoning",
	MessageTypeSPSummoned:         "spsummoned",
	MessageTypeFlipSummoning:      "flipsummoning",
	MessageTypeFlipSummoned:       "flipsummoned",
	MessageTypeChaining:           "chaining",
	MessageTypeChained:            "chained",
	MessageTypeChainSolving:       "chain_solving",
	MessageTypeChainSolved:        "chain_solved",
	MessageTypeChainEnd:           "chain_end",
	MessageTypeChainNegated:       "chain_negated",
	MessageTypeChainDisabled:      "chain_disabled",
	MessageTypeCardSelected:       "card_selected",
	MessageTypeRandomSelected:     "random_selected",
	MessageTypeBecomeTarget:       "become_target",
	MessageTypeDraw:               "draw",
	MessageTypeDamage:             "damage",
	MessageTypeRecover:            "recover",
	MessageTypeEquip:              "equip",
	MessageTypeLPUpdate:           "lpupdate",
	MessageTypeUnequip:            "unequip",
	MessageTypeCardTarget:         "card_target",
	MessageTypeCancelTarget:       "cancel_target",
	MessageTypePayLPCost:          "pay_lpcost",
	MessageTypeAddCounter:         "add_counter",
	MessageTypeRemoveCounter:      "remove_counter",
	MessageTypeAttack:             "attack",
	MessageTypeBattle:             "battle",
	MessageTypeAttackDisabled:     "attack_disabled",
	MessageTypeDamageStepStart:    "damage_step_start",
	MessageTypeDamageStepEnd:      "damage_step_end",
	MessageTypeMissedEffect:       "missed_effect",
	MessageTypeBeChainTarget:      "be_chain_target",
	MessageTypeCreateRelation:     "create_relation",
	MessageTypeReleaseRelation:    "release_relation",
	MessageTypeTossCoin:           "toss_coin",
	MessageTypeTossDice:           "toss_dice",
	MessageTypeRockPaperScissors:  "rock_paper_scissors",
	MessageTypeHandRes:            "hand_res",
	MessageTypeAnnounceRace:       "announce_race",
	MessageTypeAnnounceAttribute:  "announce_attrib",
	MessageTypeAnnounceCard:       "announce_card",
	MessageTypeAnnounceNumber:     "announce_number",
	MessageTypeCardHint:           "card_hint",
	MessageTypeTagSwap:            "tag_swap",
	MessageTypeReloadField:        "reload_field",
	MessageTypeAIName:             "ai_name",
	MessageTypeShowHint:           "show_hint",
	MessageTypePlayerHint:         "player_hint",
	MessageTypeMatchKill:          "match_kill",
	MessageTypeCustomMessage:      "custom_msg",
	MessageTypeRemoveCards:        "remove_cards",
}

type Message interface {
	Type() MessageType
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
	Code       int          `json:"code"`
	Controller int          `json:"controller"`
	Location   Location     `json:"location"`
	Sequence   int          `json:"sequence"`
	Position   corePosition `json:"position"`
}

type CounterCardInfo struct {
	Code       int      `json:"code"`
	Controller int      `json:"controller"`
	Location   Location `json:"location"`
	Sequence   int      `json:"sequence"`
	Count      int      `json:"count"`
}

type CardChainInfo struct {
	Code        int          `json:"code"`
	Controller  int          `json:"controller"`
	Location    Location     `json:"location"`
	Sequence    int          `json:"sequence"`
	Position    corePosition `json:"position"`
	Description uint64       `json:"description"`
	ClientMode  uint8        `json:"client_mode"`
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
	Position corePosition `json:"position"`
}

type MessageRetry struct {
}

func ReadMessageRetry(*bytes.Buffer) (msg MessageRetry) {
	return
}

func (MessageRetry) Type() MessageType {
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

func (MessageHint) Type() MessageType {
	return MessageTypeHint
}

type MessageWaiting struct {
}

func ReadMessageWaiting(*bytes.Buffer) (msg MessageWaiting) {
	return
}

func (MessageWaiting) Type() MessageType {
	return MessageTypeWaiting
}

type MessageStart struct {
}

func ReadMessageStart(*bytes.Buffer) (msg MessageStart) {
	return
}

func (MessageStart) Type() MessageType {
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

func (MessageWin) Type() MessageType {
	return MessageTypeWin
}

type MessageUpdateData struct {
}

func ReadMessageUpdateData(*bytes.Buffer) (msg MessageUpdateData) {
	return
}

func (MessageUpdateData) Type() MessageType {
	return MessageTypeUpdateData
}

type MessageUpdateCard struct {
}

func ReadMessageUpdateCard(*bytes.Buffer) (msg MessageUpdateCard) {
	return
}

func (MessageUpdateCard) Type() MessageType {
	return MessageTypeUpdateCard
}

type MessageRequestDeck struct {
}

func ReadMessageRequestDeck(*bytes.Buffer) (msg MessageRequestDeck) {
	return
}

func (MessageRequestDeck) Type() MessageType {
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
	for i := 0; i < int(selectChainsSize); i++ {
		msg.Chains = append(msg.Chains, ChainInfo{
			Code:        int(readUint32(b)),
			Controller:  int(readUint8(b)),
			Location:    parseCoreLocation(coreLocation(readUint8(b))),
			Sequence:    int(readUint32(b)),
			Description: readUint64(b),
			ClientMode:  readUint8(b),
		})
	}

	attackableSize := readUint32(b)
	for i := 0; i < int(attackableSize); i++ {
		msg.Attacks = append(msg.Attacks, AttackInfo{
			Code:       int(readUint32(b)),
			Controller: int(readUint8(b)),
			Location:   parseCoreLocation(coreLocation(readUint8(b))),
			Sequence:   int(readUint32(b)),
			Direct:     readUint8(b) != 0,
		})
	}
	msg.ToM2 = readUint8(b) != 0
	msg.ToEP = readUint8(b) != 0
	return
}

func (MessageSelectBattleCMD) Type() MessageType {
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
	for i := 0; i < int(summonableSize); i++ {
		msg.Summons = append(msg.Summons, CardInfo{
			Code:       int(readUint32(b)),
			Controller: int(readUint8(b)),
			Location:   parseCoreLocation(coreLocation(readUint8(b))),
			Sequence:   int(readUint32(b)),
		})
	}

	spSummonableSize := readUint32(b)
	for i := 0; i < int(spSummonableSize); i++ {
		msg.SpSummons = append(msg.SpSummons, CardInfo{
			Code:       int(readUint32(b)),
			Controller: int(readUint8(b)),
			Location:   parseCoreLocation(coreLocation(readUint8(b))),
			Sequence:   int(readUint32(b)),
		})
	}

	posChangeSize := readUint32(b)
	for i := 0; i < int(posChangeSize); i++ {
		msg.PosChanges = append(msg.PosChanges, CardInfo{
			Code:       int(readUint32(b)),
			Controller: int(readUint8(b)),
			Location:   parseCoreLocation(coreLocation(readUint8(b))),
			Sequence:   int(readUint32(b)),
		})
	}

	monsterSetSize := readUint32(b)
	for i := 0; i < int(monsterSetSize); i++ {
		msg.MonsterSets = append(msg.MonsterSets, CardInfo{
			Code:       int(readUint32(b)),
			Controller: int(readUint8(b)),
			Location:   parseCoreLocation(coreLocation(readUint8(b))),
			Sequence:   int(readUint32(b)),
		})
	}

	spellSetSize := readUint32(b)
	for i := 0; i < int(spellSetSize); i++ {
		msg.SpellSets = append(msg.SpellSets, CardInfo{
			Code:       int(readUint32(b)),
			Controller: int(readUint8(b)),
			Location:   parseCoreLocation(coreLocation(readUint8(b))),
			Sequence:   int(readUint32(b)),
		})
	}

	activateSize := readUint32(b)
	for i := 0; i < int(activateSize); i++ {
		msg.Activate = append(msg.Activate, ChainInfo{
			Code:        int(readUint32(b)),
			Controller:  int(readUint8(b)),
			Location:    parseCoreLocation(coreLocation(readUint8(b))),
			Sequence:    int(readUint32(b)),
			Description: readUint64(b),
			ClientMode:  readUint8(b),
		})
	}

	msg.ToBP = readUint8(b) != 0
	msg.ToEP = readUint8(b) != 0
	msg.Shuffle = readUint8(b) != 0
	return
}

func (MessageSelectIdleCMD) Type() MessageType {
	return MessageTypeSelectIdleCMD
}

type MessageSelectEffectYN struct {
	Player      int          `json:"player"`
	Code        uint32       `json:"code"`
	Controller  int          `json:"controller"`
	Location    Location     `json:"location"`
	Sequence    int          `json:"sequence"`
	Position    corePosition `json:"position"`
	Description uint64       `json:"description"`
}

func ReadMessageSelectEffectYN(b *bytes.Buffer) (msg MessageSelectEffectYN) {
	msg.Player = int(readUint8(b))
	msg.Code = readUint32(b)
	loc := readCardLocation(b)
	msg.Controller = loc.controller
	msg.Location = parseCoreLocation(loc.location)
	msg.Sequence = loc.sequence
	msg.Position = loc.position
	msg.Description = readUint64(b)
	return
}

func (MessageSelectEffectYN) Type() MessageType {
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

func (MessageSelectYesNo) Type() MessageType {
	return MessageTypeSelectYesNo
}

type MessageSelectOption struct {
	Player  int      `json:"player"`
	Options []uint64 `json:"options"`
}

func ReadMessageSelectOption(b *bytes.Buffer) (msg MessageSelectOption) {
	msg.Player = int(readUint8(b))

	optionsSize := readUint8(b)
	for i := 0; i < int(optionsSize); i++ {
		msg.Options = append(msg.Options, readUint64(b))
	}
	return
}

func (MessageSelectOption) Type() MessageType {
	return MessageTypeSelectOption
}

type MessageSelectCard struct {
	Player      int             `json:"player"`
	Cancellable bool            `json:"cancellable"`
	Min         int             `json:"min"`
	Max         int             `json:"max"`
	Cards       []FieldCardInfo `json:"cards"`
}

func ReadMessageSelectCard(b *bytes.Buffer) (msg MessageSelectCard) {
	msg.Player = int(readUint8(b))
	msg.Cancellable = readUint8(b) != 0
	msg.Min = int(readUint32(b))
	msg.Max = int(readUint32(b))

	cardsSize := readUint32(b)
	for i := 0; i < int(cardsSize); i++ {
		code := readUint32(b)
		loc := readCardLocation(b)
		msg.Cards = append(msg.Cards, FieldCardInfo{
			Code:       int(code),
			Controller: loc.controller,
			Location:   parseCoreLocation(loc.location),
			Sequence:   loc.sequence,
			Position:   loc.position,
		})
	}
	return
}

func (MessageSelectCard) Type() MessageType {
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
	for i := 0; i < int(chainsSize); i++ {
		code := readUint32(b)
		loc := readCardLocation(b)
		description := readUint64(b)
		clientMode := readUint8(b)
		msg.Chains = append(msg.Chains, CardChainInfo{
			Code:        int(code),
			Controller:  loc.controller,
			Location:    parseCoreLocation(loc.location),
			Sequence:    loc.sequence,
			Position:    loc.position,
			Description: description,
			ClientMode:  clientMode,
		})
	}
	return
}

func (MessageSelectChain) Type() MessageType {
	return MessageTypeSelectChain
}

type MessageSelectPlace struct {
	Player int    `json:"player"`
	Count  int    `json:"count"`
	Flag   uint32 `json:"flag"`
}

func ReadMessageSelectPlace(b *bytes.Buffer) (msg MessageSelectPlace) {
	msg.Player = int(readUint8(b))
	msg.Count = int(readUint8(b))
	msg.Flag = readUint32(b)
	return
}

func (MessageSelectPlace) Type() MessageType {
	return MessageTypeSelectPlace
}

type MessageSelectPosition struct {
	Player    int    `json:"player"`
	Code      uint32 `json:"code"`
	Positions uint8  `json:"positions"`
}

func ReadMessageSelectPosition(b *bytes.Buffer) (msg MessageSelectPosition) {
	msg.Player = int(readUint8(b))
	msg.Code = readUint32(b)
	msg.Positions = readUint8(b)
	return
}

func (MessageSelectPosition) Type() MessageType {
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
	for i := 0; i < int(tributeSize); i++ {
		msg.Cards = append(msg.Cards, TributeCardInfo{
			Code:         int(readUint32(b)),
			Controller:   int(readUint8(b)),
			Location:     parseCoreLocation(coreLocation(readUint8(b))),
			Sequence:     int(readUint32(b)),
			ReleaseParam: int(readUint8(b)),
		})
	}
	return
}

func (MessageSelectTribute) Type() MessageType {
	return MessageTypeSelectTribute
}

type MessageSortChain struct {
	Player int        `json:"player"`
	Cards  []CardInfo `json:"cards"`
}

func ReadMessageSortChain(b *bytes.Buffer) (msg MessageSortChain) {
	msg.Player = int(readUint8(b))
	cardsSize := readUint32(b)
	for i := 0; i < int(cardsSize); i++ {
		msg.Cards = append(msg.Cards, CardInfo{
			Code:       int(readUint32(b)),
			Controller: int(readUint8(b)),
			Location:   parseCoreLocation(coreLocation(readUint8(b))),
			Sequence:   int(readUint32(b)),
		})
	}
	return
}

func (MessageSortChain) Type() MessageType {
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
	for i := 0; i < int(cardsSize); i++ {
		msg.Cards = append(msg.Cards, CounterCardInfo{
			Code:       int(readUint32(b)),
			Controller: int(readUint8(b)),
			Location:   parseCoreLocation(coreLocation(readUint8(b))),
			Sequence:   int(readUint8(b)),
			Count:      int(readUint16(b)),
		})
	}
	return
}

func (MessageSelectCounter) Type() MessageType {
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
	for i := 0; i < int(mustSelectsSize); i++ {
		msg.MustSelects = append(msg.MustSelects, CounterCardInfo{
			Code:       int(readUint32(b)),
			Controller: int(readUint8(b)),
			Location:   parseCoreLocation(coreLocation(readUint8(b))),
			Sequence:   int(readUint32(b)),
			Count:      int(readUint32(b)),
		})
	}
	selectsSize := readUint32(b)
	for i := 0; i < int(selectsSize); i++ {
		msg.Selects = append(msg.Selects, CounterCardInfo{
			Code:       int(readUint32(b)),
			Controller: int(readUint8(b)),
			Location:   parseCoreLocation(coreLocation(readUint8(b))),
			Sequence:   int(readUint32(b)),
			Count:      int(readUint32(b)),
		})
	}
	return
}

func (MessageSelectSum) Type() MessageType {
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

func (MessageSelectDisfield) Type() MessageType {
	return MessageTypeSelectDisfield
}

type MessageSortCard struct {
	Player int        `json:"player"`
	Cards  []CardInfo `json:"cards"`
}

func ReadMessageSortCard(b *bytes.Buffer) (msg MessageSortCard) {
	msg.Player = int(readUint8(b))
	cardsSize := readUint32(b)
	for i := 0; i < int(cardsSize); i++ {
		msg.Cards = append(msg.Cards, CardInfo{
			Code:       int(readUint32(b)),
			Controller: int(readUint8(b)),
			Location:   parseCoreLocation(coreLocation(readUint8(b))),
			Sequence:   int(readUint32(b)),
		})
	}
	return
}

func (MessageSortCard) Type() MessageType {
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
	for i := 0; i < int(selectsSize); i++ {
		code := readUint32(b)
		loc := readCardLocation(b)
		msg.Selects = append(msg.Selects, FieldCardInfo{
			Code:       int(code),
			Controller: loc.controller,
			Location:   parseCoreLocation(loc.location),
			Sequence:   loc.sequence,
			Position:   loc.position,
		})
	}
	unselectsSize := readUint32(b)
	for i := 0; i < int(unselectsSize); i++ {
		code := readUint32(b)
		loc := readCardLocation(b)
		msg.Selects = append(msg.Selects, FieldCardInfo{
			Code:       int(code),
			Controller: loc.controller,
			Location:   parseCoreLocation(loc.location),
			Sequence:   loc.sequence,
			Position:   loc.position,
		})
	}
	return
}

func (MessageSelectUnselectCard) Type() MessageType {
	return MessageTypeSelectUnselectCard
}

type MessageConfirmDeckTop struct {
	Player int        `json:"player"`
	Cards  []CardInfo `json:"cards"`
}

func ReadMessageConfirmDeckTop(b *bytes.Buffer) (msg MessageConfirmDeckTop) {
	msg.Player = int(readUint8(b))
	cardsSize := readUint32(b)
	for i := 0; i < int(cardsSize); i++ {
		msg.Cards = append(msg.Cards, CardInfo{
			Code:       int(readUint32(b)),
			Controller: int(readUint8(b)),
			Location:   parseCoreLocation(coreLocation(readUint8(b))),
			Sequence:   int(readUint32(b)),
		})
	}
	return
}

func (MessageConfirmDeckTop) Type() MessageType {
	return MessageTypeConfirmDeckTop
}

type MessageConfirmCards struct {
	Player int        `json:"player"`
	Cards  []CardInfo `json:"cards"`
}

func ReadMessageConfirmCards(b *bytes.Buffer) (msg MessageConfirmCards) {
	msg.Player = int(readUint8(b))
	cardsSize := readUint32(b)
	for i := 0; i < int(cardsSize); i++ {
		msg.Cards = append(msg.Cards, CardInfo{
			Code:       int(readUint32(b)),
			Controller: int(readUint8(b)),
			Location:   parseCoreLocation(coreLocation(readUint8(b))),
			Sequence:   int(readUint32(b)),
		})
	}
	return
}

func (MessageConfirmCards) Type() MessageType {
	return MessageTypeConfirmCards
}

type MessageShuffleDeck struct {
	Player int `json:"player"`
}

func ReadMessageShuffleDeck(b *bytes.Buffer) (msg MessageShuffleDeck) {
	msg.Player = int(readUint8(b))
	return
}

func (MessageShuffleDeck) Type() MessageType {
	return MessageTypeShuffleDeck
}

type MessageShuffleHand struct {
}

func ReadMessageShuffleHand(b *bytes.Buffer) (msg MessageShuffleHand) {
	// TODO: implement
	panic("not implemented")
}

func (MessageShuffleHand) Type() MessageType {
	return MessageTypeShuffleHand
}

type MessageRefreshDeck struct {
}

func ReadMessageRefreshDeck(b *bytes.Buffer) (msg MessageRefreshDeck) {
	// TODO: implement
	panic("not implemented")
}

func (MessageRefreshDeck) Type() MessageType {
	return MessageTypeRefreshDeck
}

type MessageSwapGraveDeck struct {
}

func ReadMessageSwapGraveDeck(b *bytes.Buffer) (msg MessageSwapGraveDeck) {
	// TODO: implement
	panic("not implemented")
}

func (MessageSwapGraveDeck) Type() MessageType {
	return MessageTypeSwapGraveDeck
}

type MessageShuffleSetCard struct {
}

func ReadMessageShuffleSetCard(b *bytes.Buffer) (msg MessageShuffleSetCard) {
	// TODO: implement
	panic("not implemented")
}

func (MessageShuffleSetCard) Type() MessageType {
	return MessageTypeShuffleSetCard
}

type MessageReverseDeck struct {
}

func ReadMessageReverseDeck(b *bytes.Buffer) (msg MessageReverseDeck) {
	// TODO: implement
	panic("not implemented")
}

func (MessageReverseDeck) Type() MessageType {
	return MessageTypeReverseDeck
}

type MessageDeckTop struct {
}

func ReadMessageDeckTop(b *bytes.Buffer) (msg MessageDeckTop) {
	// TODO: implement
	panic("not implemented")
}

func (MessageDeckTop) Type() MessageType {
	return MessageTypeDeckTop
}

type MessageShuffleExtra struct {
}

func ReadMessageShuffleExtra(b *bytes.Buffer) (msg MessageShuffleExtra) {
	// TODO: implement
	panic("not implemented")
}

func (MessageShuffleExtra) Type() MessageType {
	return MessageTypeShuffleExtra
}

type MessageNewTurn struct {
	Player int `json:"player"`
}

func ReadMessageNewTurn(b *bytes.Buffer) (msg MessageNewTurn) {
	msg.Player = int(readUint8(b))
	return
}

func (MessageNewTurn) Type() MessageType {
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

func (MessageNewPhase) Type() MessageType {
	return MessageTypeNewPhase
}

type MessageConfirmExtraTop struct {
}

func ReadMessageConfirmExtraTop(b *bytes.Buffer) (msg MessageConfirmExtraTop) {
	// TODO: implement
	panic("not implemented")
}

func (MessageConfirmExtraTop) Type() MessageType {
	return MessageTypeConfirmExtraTop
}

type MessageMove struct {
}

func ReadMessageMove(b *bytes.Buffer) (msg MessageMove) {
	// TODO: implement
	panic("not implemented")
}

func (MessageMove) Type() MessageType {
	return MessageTypeMove
}

type MessagePosChange struct {
}

func ReadMessagePosChange(b *bytes.Buffer) (msg MessagePosChange) {
	// TODO: implement
	panic("not implemented")
}

func (MessagePosChange) Type() MessageType {
	return MessageTypePosChange
}

type MessageSet struct {
}

func ReadMessageSet(b *bytes.Buffer) (msg MessageSet) {
	// TODO: implement
	panic("not implemented")
}

func (MessageSet) Type() MessageType {
	return MessageTypeSet
}

type MessageSwap struct {
}

func ReadMessageSwap(b *bytes.Buffer) (msg MessageSwap) {
	// TODO: implement
	panic("not implemented")
}

func (MessageSwap) Type() MessageType {
	return MessageTypeSwap
}

type MessageFieldDisabled struct {
}

func ReadMessageFieldDisabled(b *bytes.Buffer) (msg MessageFieldDisabled) {
	// TODO: implement
	panic("not implemented")
}

func (MessageFieldDisabled) Type() MessageType {
	return MessageTypeFieldDisabled
}

type MessageSummoning struct {
}

func ReadMessageSummoning(b *bytes.Buffer) (msg MessageSummoning) {
	// TODO: implement
	panic("not implemented")
}

func (MessageSummoning) Type() MessageType {
	return MessageTypeSummoning
}

type MessageSummoned struct {
}

func ReadMessageSummoned(b *bytes.Buffer) (msg MessageSummoned) {
	// TODO: implement
	panic("not implemented")
}

func (MessageSummoned) Type() MessageType {
	return MessageTypeSummoned
}

type MessageSPSummoning struct {
}

func ReadMessageSPSummoning(b *bytes.Buffer) (msg MessageSPSummoning) {
	// TODO: implement
	panic("not implemented")
}

func (MessageSPSummoning) Type() MessageType {
	return MessageTypeSPSummoning
}

type MessageSPSummoned struct {
}

func ReadMessageSPSummoned(b *bytes.Buffer) (msg MessageSPSummoned) {
	// TODO: implement
	panic("not implemented")
}

func (MessageSPSummoned) Type() MessageType {
	return MessageTypeSPSummoned
}

type MessageFlipSummoning struct {
}

func ReadMessageFlipSummoning(b *bytes.Buffer) (msg MessageFlipSummoning) {
	// TODO: implement
	panic("not implemented")
}

func (MessageFlipSummoning) Type() MessageType {
	return MessageTypeFlipSummoning
}

type MessageFlipSummoned struct {
}

func ReadMessageFlipSummoned(b *bytes.Buffer) (msg MessageFlipSummoned) {
	// TODO: implement
	panic("not implemented")
}

func (MessageFlipSummoned) Type() MessageType {
	return MessageTypeFlipSummoned
}

type MessageChaining struct {
}

func ReadMessageChaining(b *bytes.Buffer) (msg MessageChaining) {
	// TODO: implement
	panic("not implemented")
}

func (MessageChaining) Type() MessageType {
	return MessageTypeChaining
}

type MessageChained struct {
}

func ReadMessageChained(b *bytes.Buffer) (msg MessageChained) {
	// TODO: implement
	panic("not implemented")
}

func (MessageChained) Type() MessageType {
	return MessageTypeChained
}

type MessageChainSolving struct {
}

func ReadMessageChainSolving(b *bytes.Buffer) (msg MessageChainSolving) {
	// TODO: implement
	panic("not implemented")
}

func (MessageChainSolving) Type() MessageType {
	return MessageTypeChainSolving
}

type MessageChainSolved struct {
}

func ReadMessageChainSolved(b *bytes.Buffer) (msg MessageChainSolved) {
	// TODO: implement
	panic("not implemented")
}

func (MessageChainSolved) Type() MessageType {
	return MessageTypeChainSolved
}

type MessageChainEnd struct {
}

func ReadMessageChainEnd(b *bytes.Buffer) (msg MessageChainEnd) {
	// TODO: implement
	panic("not implemented")
}

func (MessageChainEnd) Type() MessageType {
	return MessageTypeChainEnd
}

type MessageChainNegated struct {
}

func ReadMessageChainNegated(b *bytes.Buffer) (msg MessageChainNegated) {
	// TODO: implement
	panic("not implemented")
}

func (MessageChainNegated) Type() MessageType {
	return MessageTypeChainNegated
}

type MessageChainDisabled struct {
}

func ReadMessageChainDisabled(b *bytes.Buffer) (msg MessageChainDisabled) {
	// TODO: implement
	panic("not implemented")
}

func (MessageChainDisabled) Type() MessageType {
	return MessageTypeChainDisabled
}

type MessageCardSelected struct {
}

func ReadMessageCardSelected(b *bytes.Buffer) (msg MessageCardSelected) {
	// TODO: implement
	panic("not implemented")
}

func (MessageCardSelected) Type() MessageType {
	return MessageTypeCardSelected
}

type MessageRandomSelected struct {
}

func ReadMessageRandomSelected(b *bytes.Buffer) (msg MessageRandomSelected) {
	// TODO: implement
	panic("not implemented")
}

func (MessageRandomSelected) Type() MessageType {
	return MessageTypeRandomSelected
}

type MessageBecomeTarget struct {
}

func ReadMessageBecomeTarget(b *bytes.Buffer) (msg MessageBecomeTarget) {
	// TODO: implement
	panic("not implemented")
}

func (MessageBecomeTarget) Type() MessageType {
	return MessageTypeBecomeTarget
}

type MessageDraw struct {
	Player int             `json:"player"`
	Cards  []DrawnCardInfo `json:"cards"`
}

func ReadMessageDraw(b *bytes.Buffer) (msg MessageDraw) {
	msg.Player = int(readUint8(b))
	cardsSize := readUint32(b)
	for i := 0; i < int(cardsSize); i++ {
		msg.Cards = append(msg.Cards, DrawnCardInfo{
			Code:     int(readUint32(b)),
			Position: corePosition(readUint32(b)),
		})
	}
	return
}

func (MessageDraw) Type() MessageType {
	return MessageTypeDraw
}

type MessageDamage struct {
}

func ReadMessageDamage(b *bytes.Buffer) (msg MessageDamage) {
	// TODO: implement
	panic("not implemented")
}

func (MessageDamage) Type() MessageType {
	return MessageTypeDamage
}

type MessageRecover struct {
}

func ReadMessageRecover(b *bytes.Buffer) (msg MessageRecover) {
	// TODO: implement
	panic("not implemented")
}

func (MessageRecover) Type() MessageType {
	return MessageTypeRecover
}

type MessageEquip struct {
}

func ReadMessageEquip(b *bytes.Buffer) (msg MessageEquip) {
	// TODO: implement
	panic("not implemented")
}

func (MessageEquip) Type() MessageType {
	return MessageTypeEquip
}

type MessageLPUpdate struct {
}

func ReadMessageLPUpdate(b *bytes.Buffer) (msg MessageLPUpdate) {
	// TODO: implement
	panic("not implemented")
}

func (MessageLPUpdate) Type() MessageType {
	return MessageTypeLPUpdate
}

type MessageUnequip struct {
}

func ReadMessageUnequip(b *bytes.Buffer) (msg MessageUnequip) {
	// TODO: implement
	panic("not implemented")
}

func (MessageUnequip) Type() MessageType {
	return MessageTypeUnequip
}

type MessageCardTarget struct {
}

func ReadMessageCardTarget(b *bytes.Buffer) (msg MessageCardTarget) {
	// TODO: implement
	panic("not implemented")
}

func (MessageCardTarget) Type() MessageType {
	return MessageTypeCardTarget
}

type MessageCancelTarget struct {
}

func ReadMessageCancelTarget(b *bytes.Buffer) (msg MessageCancelTarget) {
	// TODO: implement
	panic("not implemented")
}

func (MessageCancelTarget) Type() MessageType {
	return MessageTypeCancelTarget
}

type MessagePayLPCost struct {
}

func ReadMessagePayLPCost(b *bytes.Buffer) (msg MessagePayLPCost) {
	// TODO: implement
	panic("not implemented")
}

func (MessagePayLPCost) Type() MessageType {
	return MessageTypePayLPCost
}

type MessageAddCounter struct {
}

func ReadMessageAddCounter(b *bytes.Buffer) (msg MessageAddCounter) {
	// TODO: implement
	panic("not implemented")
}

func (MessageAddCounter) Type() MessageType {
	return MessageTypeAddCounter
}

type MessageRemoveCounter struct {
}

func ReadMessageRemoveCounter(b *bytes.Buffer) (msg MessageRemoveCounter) {
	// TODO: implement
	panic("not implemented")
}

func (MessageRemoveCounter) Type() MessageType {
	return MessageTypeRemoveCounter
}

type MessageAttack struct {
}

func ReadMessageAttack(b *bytes.Buffer) (msg MessageAttack) {
	// TODO: implement
	panic("not implemented")
}

func (MessageAttack) Type() MessageType {
	return MessageTypeAttack
}

type MessageBattle struct {
}

func ReadMessageBattle(b *bytes.Buffer) (msg MessageBattle) {
	// TODO: implement
	panic("not implemented")
}

func (MessageBattle) Type() MessageType {
	return MessageTypeBattle
}

type MessageAttackDisabled struct {
}

func ReadMessageAttackDisabled(b *bytes.Buffer) (msg MessageAttackDisabled) {
	// TODO: implement
	panic("not implemented")
}

func (MessageAttackDisabled) Type() MessageType {
	return MessageTypeAttackDisabled
}

type MessageDamageStepStart struct {
}

func ReadMessageDamageStepStart(b *bytes.Buffer) (msg MessageDamageStepStart) {
	// TODO: implement
	panic("not implemented")
}

func (MessageDamageStepStart) Type() MessageType {
	return MessageTypeDamageStepStart
}

type MessageDamageStepEnd struct {
}

func ReadMessageDamageStepEnd(b *bytes.Buffer) (msg MessageDamageStepEnd) {
	// TODO: implement
	panic("not implemented")
}

func (MessageDamageStepEnd) Type() MessageType {
	return MessageTypeDamageStepEnd
}

type MessageMissedEffect struct {
}

func ReadMessageMissedEffect(b *bytes.Buffer) (msg MessageMissedEffect) {
	// TODO: implement
	panic("not implemented")
}

func (MessageMissedEffect) Type() MessageType {
	return MessageTypeMissedEffect
}

type MessageBeChainTarget struct {
}

func ReadMessageBeChainTarget(b *bytes.Buffer) (msg MessageBeChainTarget) {
	// TODO: implement
	panic("not implemented")
}

func (MessageBeChainTarget) Type() MessageType {
	return MessageTypeBeChainTarget
}

type MessageCreateRelation struct {
}

func ReadMessageCreateRelation(b *bytes.Buffer) (msg MessageCreateRelation) {
	// TODO: implement
	panic("not implemented")
}

func (MessageCreateRelation) Type() MessageType {
	return MessageTypeCreateRelation
}

type MessageReleaseRelation struct {
}

func ReadMessageReleaseRelation(b *bytes.Buffer) (msg MessageReleaseRelation) {
	// TODO: implement
	panic("not implemented")
}

func (MessageReleaseRelation) Type() MessageType {
	return MessageTypeReleaseRelation
}

type MessageTossCoin struct {
}

func ReadMessageTossCoin(b *bytes.Buffer) (msg MessageTossCoin) {
	// TODO: implement
	panic("not implemented")
}

func (MessageTossCoin) Type() MessageType {
	return MessageTypeTossCoin
}

type MessageTossDice struct {
}

func ReadMessageTossDice(b *bytes.Buffer) (msg MessageTossDice) {
	// TODO: implement
	panic("not implemented")
}

func (MessageTossDice) Type() MessageType {
	return MessageTypeTossDice
}

type MessageRockPaperScissors struct {
}

func ReadMessageRockPaperScissors(b *bytes.Buffer) (msg MessageRockPaperScissors) {
	// TODO: implement
	panic("not implemented")
}

func (MessageRockPaperScissors) Type() MessageType {
	return MessageTypeRockPaperScissors
}

type MessageHandRes struct {
}

func ReadMessageHandRes(b *bytes.Buffer) (msg MessageHandRes) {
	// TODO: implement
	panic("not implemented")
}

func (MessageHandRes) Type() MessageType {
	return MessageTypeHandRes
}

type MessageAnnounceRace struct {
}

func ReadMessageAnnounceRace(b *bytes.Buffer) (msg MessageAnnounceRace) {
	// TODO: implement
	panic("not implemented")
}

func (MessageAnnounceRace) Type() MessageType {
	return MessageTypeAnnounceRace
}

type MessageAnnounceAttribute struct {
}

func ReadMessageAnnounceAttribute(b *bytes.Buffer) (msg MessageAnnounceAttribute) {
	// TODO: implement
	panic("not implemented")
}

func (MessageAnnounceAttribute) Type() MessageType {
	return MessageTypeAnnounceAttribute
}

type MessageAnnounceCard struct {
}

func ReadMessageAnnounceCard(b *bytes.Buffer) (msg MessageAnnounceCard) {
	// TODO: implement
	panic("not implemented")
}

func (MessageAnnounceCard) Type() MessageType {
	return MessageTypeAnnounceCard
}

type MessageAnnounceNumber struct {
}

func ReadMessageAnnounceNumber(b *bytes.Buffer) (msg MessageAnnounceNumber) {
	// TODO: implement
	panic("not implemented")
}

func (MessageAnnounceNumber) Type() MessageType {
	return MessageTypeAnnounceNumber
}

type MessageCardHint struct {
}

func ReadMessageCardHint(b *bytes.Buffer) (msg MessageCardHint) {
	// TODO: implement
	panic("not implemented")
}

func (MessageCardHint) Type() MessageType {
	return MessageTypeCardHint
}

type MessageTagSwap struct {
}

func ReadMessageTagSwap(b *bytes.Buffer) (msg MessageTagSwap) {
	// TODO: implement
	panic("not implemented")
}

func (MessageTagSwap) Type() MessageType {
	return MessageTypeTagSwap
}

type MessageReloadField struct {
}

func ReadMessageReloadField(b *bytes.Buffer) (msg MessageReloadField) {
	// TODO: implement
	panic("not implemented")
}

func (MessageReloadField) Type() MessageType {
	return MessageTypeReloadField
}

type MessageAIName struct {
}

func ReadMessageAIName(b *bytes.Buffer) (msg MessageAIName) {
	// TODO: implement
	panic("not implemented")
}

func (MessageAIName) Type() MessageType {
	return MessageTypeAIName
}

type MessageShowHint struct {
}

func ReadMessageShowHint(b *bytes.Buffer) (msg MessageShowHint) {
	// TODO: implement
	panic("not implemented")
}

func (MessageShowHint) Type() MessageType {
	return MessageTypeShowHint
}

type MessagePlayerHint struct {
}

func ReadMessagePlayerHint(b *bytes.Buffer) (msg MessagePlayerHint) {
	// TODO: implement
	panic("not implemented")
}

func (MessagePlayerHint) Type() MessageType {
	return MessageTypePlayerHint
}

type MessageMatchKill struct {
}

func ReadMessageMatchKill(b *bytes.Buffer) (msg MessageMatchKill) {
	// TODO: implement
	panic("not implemented")
}

func (MessageMatchKill) Type() MessageType {
	return MessageTypeMatchKill
}

type MessageCustomMessage struct {
}

func ReadMessageCustomMessage(b *bytes.Buffer) (msg MessageCustomMessage) {
	// TODO: implement
	panic("not implemented")
}

func (MessageCustomMessage) Type() MessageType {
	return MessageTypeCustomMessage
}

type MessageRemoveCards struct {
}

func ReadMessageRemoveCards(b *bytes.Buffer) (msg MessageRemoveCards) {
	// TODO: implement
	panic("not implemented")
}

func (MessageRemoveCards) Type() MessageType {
	return MessageTypeRemoveCards
}

type jsonMessage struct {
	Message
}

func JSONMessage(m Message) interface{} {
	return jsonMessage{Message: m}
}

func (m jsonMessage) MarshalJSON() ([]byte, error) {
	structTypesCacheLock.Lock()
	defer structTypesCacheLock.Unlock()

	v := createStructTypesCache(m.Message)
	v.Field(0).Set(reflect.ValueOf(m.Message))
	return json.Marshal(v.Interface())
}

var structTypesCacheLock sync.Mutex
var structTypesCache = map[MessageType]reflect.Value{}

func createStructTypesCache(m Message) reflect.Value {
	messageType := m.Type()
	if x, ok := structTypesCache[messageType]; ok {
		return x
	}
	t := reflect.StructOf([]reflect.StructField{
		{
			Name:      reflect.TypeOf(m).Name(),
			Type:      reflect.TypeOf(m),
			Anonymous: true,
		},
		{
			Name: "Type",
			Type: reflect.TypeOf(""),
			Tag:  `json:"message_type"`,
		},
	})
	v := reflect.New(t).Elem()
	v.Field(1).Set(reflect.ValueOf(messageTypeNames[m.Type()]))
	structTypesCache[messageType] = v
	return v
}
