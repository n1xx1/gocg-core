package ocgcore

//go:generate go run ocgcore/cmd/enumer -type=Location,FieldPlace,DetailedPhase,Phase,BattlePosition,FacePosition,Position -json -transform=snake -output enum_enumer.go -trimprefix Location,FieldPlace,DetailedPhase,Phase,BattlePosition,FacePosition,Position

type Location uint

const (
	LocationUnknown Location = iota
	LocationDeck
	LocationHand
	LocationGrave
	LocationBanished
	LocationExtraDeck
	LocationOverlay
	LocationMonsterZone
	LocationSpellZone
	LocationFieldZone
	LocationPendulumZone
)

func (l Location) OnField() bool {
	return l >= LocationMonsterZone && l <= LocationPendulumZone
}

func (l Location) IsEMZ(seq int) bool {
	return l == LocationMonsterZone && (seq == 5 || seq == 6)
}

type FieldPlace int

const (
	Monster1 FieldPlace = iota
	Monster2
	Monster3
	Monster4
	Monster5
	MonsterExtra1
	MonsterExtra2
	Spell0
	Spell1
	Spell2
	Spell3
	Spell4
	SpellField
	SpellPendulum1
	SpellPendulum2
)

type DetailedPhase uint

const (
	DetailedPhaseUnknown DetailedPhase = iota
	DetailedPhaseDraw
	DetailedPhaseStandby
	DetailedPhaseMain1
	DetailedPhaseBattleStart
	DetailedPhaseBattleStep
	DetailedPhaseDamage
	DetailedPhaseDamageCalculation
	DetailedPhaseBattle
	DetailedPhaseMain2
	DetailedPhaseEnd
)

type Phase uint

const (
	PhaseUnknown Phase = iota
	PhaseDP
	PhaseSP
	PhaseM1
	PhaseBP
	PhaseM2
	PhaseEP
)

type BattlePosition int

const (
	BattlePositionUnknown BattlePosition = iota
	BattlePositionAttack
	BattlePositionDefense
)

type FacePosition int

const (
	FacePositionUnknown FacePosition = iota
	FacePositionUp
	FacePositionDown
)

type Position uint

const (
	PositionUnknown Position = iota
	PositionFaceUpAttack
	PositionFaceDownAttack
	PositionFaceUpDefense
	PositionFaceDownDefense
)

func (p Position) Battle() BattlePosition {
	switch p {
	case PositionFaceUpAttack:
		return BattlePositionAttack
	case PositionFaceDownAttack:
		return BattlePositionAttack
	case PositionFaceUpDefense:
		return BattlePositionDefense
	case PositionFaceDownDefense:
		return BattlePositionDefense
	default:
		return BattlePositionUnknown
	}
}
func (p Position) Face() FacePosition {
	switch p {
	case PositionFaceUpAttack:
		return FacePositionUp
	case PositionFaceDownAttack:
		return FacePositionDown
	case PositionFaceUpDefense:
		return FacePositionUp
	case PositionFaceDownDefense:
		return FacePositionDown
	default:
		return FacePositionUnknown
	}
}

func (p Position) FaceUp() bool {
	return p == PositionFaceUpAttack || p == PositionFaceUpDefense
}
func (p Position) FaceDown() bool {
	return p == PositionFaceDownAttack || p == PositionFaceDownDefense
}
func (p Position) Defense() bool {
	return p == PositionFaceUpDefense || p == PositionFaceDownDefense
}
func (p Position) Attack() bool {
	return p == PositionFaceUpAttack || p == PositionFaceDownAttack
}
