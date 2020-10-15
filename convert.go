package ocgcore

import "math/bits"

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

func parseCorePhaseDetailed(p corePhase) DetailedPhase {
	switch p {
	case corePhaseDraw:
		return DetailedPhaseDraw
	case corePhaseStandby:
		return DetailedPhaseStandby
	case corePhaseMain1:
		return DetailedPhaseMain1
	case corePhaseBattleStart:
		return DetailedPhaseBattleStart
	case corePhaseBattleStep:
		return DetailedPhaseBattleStep
	case corePhaseDamage:
		return DetailedPhaseDamage
	case corePhaseDamageCalculation:
		return DetailedPhaseDamageCalculation
	case corePhaseBattle:
		return DetailedPhaseBattle
	case corePhaseMain2:
		return DetailedPhaseMain2
	case corePhaseEnd:
		return DetailedPhaseEnd
	}
	return DetailedPhaseUnknown
}

func parseCorePhase(p corePhase) Phase {
	switch p {
	case corePhaseDraw:
		return PhaseDP
	case corePhaseStandby:
		return PhaseSP
	case corePhaseMain1:
		return PhaseM1
	case corePhaseBattleStart, corePhaseBattleStep, corePhaseDamage, corePhaseDamageCalculation, corePhaseBattle:
		return PhaseBP
	case corePhaseMain2:
		return PhaseM2
	case corePhaseEnd:
		return PhaseEP
	}
	return PhaseUnknown
}

func parseCorePosition(p corePosition) Position {
	switch p {
	case corePositionFaceUpAttack:
		return PositionFaceUpAttack
	case corePositionFaceDownAttack:
		return PositionFaceDownAttack
	case corePositionFaceUpDefense:
		return PositionFaceUpDefense
	case corePositionFaceDownDefense:
		return PositionFaceDownDefense
	case corePositionFaceDown:
		return PositionFaceDownAttack
	case corePositionFaceUp:
		return PositionFaceUpAttack
	case corePositionDefense:
		return PositionFaceUpDefense
	case corePositionAttack:
		return PositionFaceUpAttack
	default:
		return PositionUnknown
	}
}

func convertPosition(p Position) corePosition {
	switch p {
	case PositionFaceUpAttack:
		return corePositionFaceUpAttack
	case PositionFaceDownAttack:
		return corePositionFaceDownAttack
	case PositionFaceUpDefense:
		return corePositionFaceUpDefense
	case PositionFaceDownDefense:
		return corePositionFaceDownDefense
	}
	return 0
}

func parseCorePositions(l corePosition) []Position {
	var positions []Position
	if l&corePositionFaceUpAttack != 0 {
		positions = append(positions, PositionFaceUpAttack)
	}
	if l&corePositionFaceDownAttack != 0 {
		positions = append(positions, PositionFaceDownAttack)
	}
	if l&corePositionFaceUpDefense != 0 {
		positions = append(positions, PositionFaceUpDefense)
	}
	if l&corePositionFaceDownDefense != 0 {
		positions = append(positions, PositionFaceDownDefense)
	}
	return positions
}

func convertLocation(l Location) coreLocation {
	switch l {
	case LocationDeck:
		return coreLocationDeck
	case LocationHand:
		return coreLocationHand
	case LocationGrave:
		return coreLocationGrave
	case LocationBanished:
		return coreLocationRemoved
	case LocationExtraDeck:
		return coreLocationExtra
	case LocationOverlay:
		return coreLocationOverlay
	case LocationMonsterZone:
		return coreLocationMZone
	case LocationSpellZone:
		return coreLocationSZone
	case LocationFieldZone:
		return coreLocationFZone
	case LocationPendulumZone:
		return coreLocationPZone
	}
	return 0
}

func parseCoreLocation(l coreLocation) Location {
	if l&coreLocationPZone != 0 {
		return LocationPendulumZone
	}
	if l&coreLocationFZone != 0 {
		return LocationFieldZone
	}
	switch l {
	case coreLocationDeck:
		return LocationDeck
	case coreLocationHand:
		return LocationHand
	case coreLocationGrave:
		return LocationGrave
	case coreLocationRemoved:
		return LocationBanished
	case coreLocationExtra:
		return LocationExtraDeck
	case coreLocationOverlay:
		return LocationOverlay
	case coreLocationMZone:
		return LocationMonsterZone
	case coreLocationSZone:
		return LocationSpellZone
	}
	return LocationUnknown
}
