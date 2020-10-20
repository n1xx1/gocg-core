package ocgcore

import (
	"bytes"
	"ocgcore/lib"
	"ocgcore/utils"
)

//go:generate go run ocgcore/cmd/enumer -type=ResponseType,BattleAction,IdleAction -json -transform=snake -output response_enumer.go -trimprefix BattleAction,IdleAction
//go:generate go run ocgcore/cmd/interfacer -method=responseType -returns=ResponseType -field=response_type -interface Response -output response_interfacer.go

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

func (r ResponseSelectBattleCMD) responseType() ResponseType {
	return ResponseTypeSelectBattleCMD
}

func (r ResponseSelectBattleCMD) responseWrite() []byte {
	var b bytes.Buffer
	utils.WriteUint32(&b, (uint32(r.Action)&0xff)|((uint32(r.Index)&0xff)<<16))
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

func (r ResponseSelectIdleCMD) responseType() ResponseType {
	return ResponseTypeSelectIdleCMD
}

func (r ResponseSelectIdleCMD) responseWrite() []byte {
	var b bytes.Buffer
	utils.WriteUint32(&b, (uint32(r.Action)&0xff)|((uint32(r.Index)&0xff)<<16))
	return b.Bytes()
}

type ResponseSelectEffectYN struct {
	Yes bool `json:"yes"`
}

func (r ResponseSelectEffectYN) responseType() ResponseType {
	return ResponseTypeSelectEffectYN
}

func (r ResponseSelectEffectYN) responseWrite() []byte {
	var b bytes.Buffer
	if r.Yes {
		utils.WriteInt32(&b, 1)
	} else {
		utils.WriteInt32(&b, 0)
	}
	return b.Bytes()
}

type ResponseSelectYesNo struct {
	Yes bool `json:"yes"`
}

func (r ResponseSelectYesNo) responseType() ResponseType {
	return ResponseTypeSelectYesNo
}

func (r ResponseSelectYesNo) responseWrite() []byte {
	var b bytes.Buffer
	if r.Yes {
		utils.WriteInt32(&b, 1)
	} else {
		utils.WriteInt32(&b, 0)
	}
	return b.Bytes()
}

type ResponseSelectOption struct {
	Option int `json:"option"`
}

func (r ResponseSelectOption) responseType() ResponseType {
	return ResponseTypeSelectOption
}

func (r ResponseSelectOption) responseWrite() []byte {
	var b bytes.Buffer
	utils.WriteInt32(&b, int32(r.Option))
	return b.Bytes()
}

type ResponseSelectCard struct {
	Cancel bool  `json:"cancel"`
	Select []int `json:"select,omitempty"`
}

func (r ResponseSelectCard) responseType() ResponseType {
	return ResponseTypeSelectCard
}

func (r ResponseSelectCard) responseWrite() []byte {
	var b bytes.Buffer
	if r.Cancel {
		utils.WriteInt32(&b, -1)
	} else {
		utils.WriteInt32(&b, 2) // only support type 2
		utils.WriteInt32(&b, int32(len(r.Select)))
		for _, c := range r.Select {
			utils.WriteInt8(&b, int8(c))
		}
	}
	return b.Bytes()
}

type ResponseSelectChain struct {
	Chain int `json:"chain"`
}

func (r ResponseSelectChain) responseType() ResponseType {
	return ResponseTypeSelectChain
}

func (r ResponseSelectChain) responseWrite() []byte {
	var b bytes.Buffer
	utils.WriteInt32(&b, int32(r.Chain))
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

func (r ResponseSelectPlace) responseType() ResponseType {
	return ResponseTypeSelectPlace
}

func (r ResponseSelectPlace) responseWrite() []byte {
	var b bytes.Buffer
	for _, p := range r.Places {
		seq := uint8(p.Sequence)
		loc := convertLocation(p.Location)
		if loc == lib.LocationPZone {
			loc = lib.LocationSZone
			seq -= 6
		}
		utils.WriteUint8(&b, uint8(p.Player))
		utils.WriteUint8(&b, uint8(loc))
		utils.WriteUint8(&b, seq)
	}
	return b.Bytes()
}

type ResponseSelectPosition struct {
	Position Position `json:"position"`
}

func (r ResponseSelectPosition) responseType() ResponseType {
	return ResponseTypeSelectPosition
}

func (r ResponseSelectPosition) responseWrite() []byte {
	var b bytes.Buffer
	utils.WriteInt32(&b, int32(convertPosition(r.Position)))
	return b.Bytes()
}

type ResponseSelectUnselectCard struct {
	Cancel    bool `json:"cancel"`
	Selection int  `json:"selection"`
}

func (r ResponseSelectUnselectCard) responseType() ResponseType {
	return ResponseTypeSelectUnselectCard
}

func (r ResponseSelectUnselectCard) responseWrite() []byte {
	var b bytes.Buffer
	if r.Cancel {
		utils.WriteInt32(&b, -1)
	} else {
		utils.WriteInt32(&b, 1)
		utils.WriteInt32(&b, int32(r.Selection))
	}
	return b.Bytes()
}
