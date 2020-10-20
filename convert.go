package ocgcore

import (
	"math/bits"
	"ocgcore/lib"
)

//       other              our
// _____szone__mzone _____szone__mzone
// 00000000 00000000 00000000 00000000
func parsePlaceFlag(f uint32) []Place {
	ret := make([]Place, bits.OnesCount32(^f))

	x := 0
	parsePlaceFlagPlayer((^f)&0xffff, 0, &x, ret)
	parsePlaceFlagPlayer((^f)>>16, 1, &x, ret)
	return ret
}

func parsePlaceFlagPlayer(flag uint32, player int, x *int, ret []Place) {
	mask := uint32(1)
	for i := 0; i < 7; i++ {
		if tryMask(flag, &mask) {
			ret[*x] = Place{Player: player, Location: LocationMonsterZone, Sequence: i}
			*x++
		}
	}
	// TODO: figure out if it's important or not
	mask = mask << 1
	for i := 0; i < 5; i++ {
		if tryMask(flag, &mask) {
			ret[*x] = Place{Player: player, Location: LocationSpellZone, Sequence: i}
			*x++
		}
	}
	if tryMask(flag, &mask) {
		ret[*x] = Place{Player: player, Location: LocationFieldZone}
		*x++
	}
	for i := 0; i < 2; i++ {
		if tryMask(flag, &mask) {
			ret[*x] = Place{Player: player, Location: LocationPendulumZone, Sequence: i}
			*x++
		}
	}
}

func tryMask(flag uint32, mask *uint32) (r bool) {
	if flag&*mask == *mask {
		r = true
	}
	*mask = *mask << 1
	return
}

func parseCorePhaseDetailed(p lib.Phase) DetailedPhase {
	switch p {
	case lib.PhaseDraw:
		return DetailedPhaseDraw
	case lib.PhaseStandby:
		return DetailedPhaseStandby
	case lib.PhaseMain1:
		return DetailedPhaseMain1
	case lib.PhaseBattleStart:
		return DetailedPhaseBattleStart
	case lib.PhaseBattleStep:
		return DetailedPhaseBattleStep
	case lib.PhaseDamage:
		return DetailedPhaseDamage
	case lib.PhaseDamageCalculation:
		return DetailedPhaseDamageCalculation
	case lib.PhaseBattle:
		return DetailedPhaseBattle
	case lib.PhaseMain2:
		return DetailedPhaseMain2
	case lib.PhaseEnd:
		return DetailedPhaseEnd
	}
	return DetailedPhaseUnknown
}

func parseCorePhase(p lib.Phase) Phase {
	switch p {
	case lib.PhaseDraw:
		return PhaseDP
	case lib.PhaseStandby:
		return PhaseSP
	case lib.PhaseMain1:
		return PhaseM1
	case lib.PhaseBattleStart, lib.PhaseBattleStep, lib.PhaseDamage, lib.PhaseDamageCalculation, lib.PhaseBattle:
		return PhaseBP
	case lib.PhaseMain2:
		return PhaseM2
	case lib.PhaseEnd:
		return PhaseEP
	}
	return PhaseUnknown
}

func parseCorePosition(p lib.Position) Position {
	switch p {
	case lib.PositionFaceUpAttack:
		return PositionFaceUpAttack
	case lib.PositionFaceDownAttack:
		return PositionFaceDownAttack
	case lib.PositionFaceUpDefense:
		return PositionFaceUpDefense
	case lib.PositionFaceDownDefense:
		return PositionFaceDownDefense
	case lib.PositionFaceDown:
		return PositionFaceDownAttack
	case lib.PositionFaceUp:
		return PositionFaceUpAttack
	case lib.PositionDefense:
		return PositionFaceUpDefense
	case lib.PositionAttack:
		return PositionFaceUpAttack
	default:
		return PositionUnknown
	}
}

func convertPosition(p Position) lib.Position {
	switch p {
	case PositionFaceUpAttack:
		return lib.PositionFaceUpAttack
	case PositionFaceDownAttack:
		return lib.PositionFaceDownAttack
	case PositionFaceUpDefense:
		return lib.PositionFaceUpDefense
	case PositionFaceDownDefense:
		return lib.PositionFaceDownDefense
	}
	return 0
}

func parseCorePositions(l lib.Position) []Position {
	var positions []Position
	if l&lib.PositionFaceUpAttack != 0 {
		positions = append(positions, PositionFaceUpAttack)
	}
	if l&lib.PositionFaceDownAttack != 0 {
		positions = append(positions, PositionFaceDownAttack)
	}
	if l&lib.PositionFaceUpDefense != 0 {
		positions = append(positions, PositionFaceUpDefense)
	}
	if l&lib.PositionFaceDownDefense != 0 {
		positions = append(positions, PositionFaceDownDefense)
	}
	return positions
}

func convertLocation(l Location) lib.Location {
	switch l {
	case LocationDeck:
		return lib.LocationDeck
	case LocationHand:
		return lib.LocationHand
	case LocationGrave:
		return lib.LocationGrave
	case LocationBanished:
		return lib.LocationRemoved
	case LocationExtraDeck:
		return lib.LocationExtra
	case LocationOverlay:
		return lib.LocationOverlay
	case LocationMonsterZone:
		return lib.LocationMZone
	case LocationSpellZone:
		return lib.LocationSZone
	case LocationFieldZone:
		return lib.LocationFZone
	case LocationPendulumZone:
		return lib.LocationPZone
	}
	return 0
}

func parseCoreLocation(l lib.Location) Location {
	if l&lib.LocationPZone != 0 {
		return LocationPendulumZone
	}
	if l&lib.LocationFZone != 0 {
		return LocationFieldZone
	}
	switch l {
	case lib.LocationDeck:
		return LocationDeck
	case lib.LocationHand:
		return LocationHand
	case lib.LocationGrave:
		return LocationGrave
	case lib.LocationRemoved:
		return LocationBanished
	case lib.LocationExtra:
		return LocationExtraDeck
	case lib.LocationOverlay:
		return LocationOverlay
	case lib.LocationMZone:
		return LocationMonsterZone
	case lib.LocationSZone:
		return LocationSpellZone
	}
	return LocationUnknown
}

func ParseCardType(ot lib.CardType) CardType {
	switch {
	case ot&lib.CardTypeMonster != 0:
		return CardTypeMonster
	case ot&lib.CardTypeSpell != 0:
		return CardTypeSpell
	case ot&lib.CardTypeTrap != 0:
		return CardTypeTrap
	case ot&lib.CardTypeToken != 0:
		return CardTypeToken
	}
	return 0
}

func ParseCardTypeMonster(ot lib.CardType) (mf CardMonsterFrame, mt CardMonsterType, ma CardMonsterAbility, mtu bool, mp bool) {
	switch {
	case ot&lib.CardTypeNormal != 0:
		mf = CardMonsterFrameNormal
	case ot&lib.CardTypeEffect != 0:
		mf = CardMonsterFrameEffect
	case ot&lib.CardTypeFusion != 0:
		mf = CardMonsterFrameFusion
	case ot&lib.CardTypeRitual != 0:
		mf = CardMonsterFrameRitual
	case ot&lib.CardTypeSynchro != 0:
		mf = CardMonsterFrameSynchro
	case ot&lib.CardTypeXyz != 0:
		mf = CardMonsterFrameXyz
	case ot&lib.CardTypeLink != 0:
		mf = CardMonsterFrameLink
	}
	switch {
	case ot&lib.CardTypeSpirit != 0:
		ma = CardMonsterAbilitySpirit
	case ot&lib.CardTypeUnion != 0:
		ma = CardMonsterAbilityUnion
	case ot&lib.CardTypeGemini != 0:
		ma = CardMonsterAbilityGemini
	case ot&lib.CardTypeFlip != 0:
		ma = CardMonsterAbilityFlip
	case ot&lib.CardTypeToon != 0:
		ma = CardMonsterAbilityToon
	}
	if ot&lib.CardTypeTuner != 0 {
		mtu = true
	}
	if ot&lib.CardTypePendulum != 0 {
		mp = true
	}
	return
}

func ParseCardTypeSpell(ot lib.CardType) CardSpellType {
	switch {
	case ot&lib.CardTypeQuickPlay != 0:
		return CardSpellTypeQuickPlay
	case ot&lib.CardTypeContinuous != 0:
		return CardSpellTypeContinuous
	case ot&lib.CardTypeEquip != 0:
		return CardSpellTypeEquip
	case ot&lib.CardTypeField != 0:
		return CardSpellTypeField
	case ot&lib.CardTypeRitual != 0:
		return CardSpellTypeRitual
	}
	return CardSpellTypeNormal
}

func ParseCardTypeTrap(ot lib.CardType) CardTrapType {
	switch {
	case ot&lib.CardTypeContinuous != 0:
		return CardTrapTypeContinuous
	case ot&lib.CardTypeCounter != 0:
		return CardTrapTypeCounter
	}
	return CardTrapTypeNormal
}

func ParseLinkMarkers(m lib.LinkMarker) []CardLinkMarker {
	ms := make([]CardLinkMarker, bits.OnesCount32(uint32(m)))
	msi := 0
	if m&lib.LinkMarkerBottomLeft != 0 {
		ms[msi] = CardLinkMarkerBottomLeft
		msi++
	}
	if m&lib.LinkMarkerBottom != 0 {
		ms[msi] = CardLinkMarkerBottom
		msi++
	}
	if m&lib.LinkMarkerBottomRight != 0 {
		ms[msi] = CardLinkMarkerBottomRight
		msi++
	}
	if m&lib.LinkMarkerLeft != 0 {
		ms[msi] = CardLinkMarkerLeft
		msi++
	}
	if m&lib.LinkMarkerRight != 0 {
		ms[msi] = CardLinkMarkerRight
		msi++
	}
	if m&lib.LinkMarkerTopLeft != 0 {
		ms[msi] = CardLinkMarkerTopLeft
		msi++
	}
	if m&lib.LinkMarkerTop != 0 {
		ms[msi] = CardLinkMarkerTop
		msi++
	}
	if m&lib.LinkMarkerTopRight != 0 {
		ms[msi] = CardLinkMarkerTopRight
		msi++
	}
	return ms
}
