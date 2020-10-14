package ocgcore

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

func (l Location) GoString() string {
	if l, ok := locationNames[l]; ok {
		return l
	}
	return "unknown"
}

func (l Location) OnField() bool {
	return l >= LocationMonsterZone && l <= LocationPendulumZone
}

func (l Location) MarshalJSON() ([]byte, error) {
	return []byte(`"` + l.GoString() + `"`), nil
}

var locationNames = map[Location]string{
	LocationDeck:         "deck",
	LocationHand:         "hand",
	LocationGrave:        "grave",
	LocationBanished:     "banished",
	LocationExtraDeck:    "extra_deck",
	LocationOverlay:      "overlay",
	LocationMonsterZone:  "monster_zone",
	LocationSpellZone:    "spell_zone",
	LocationFieldZone:    "field_zone",
	LocationPendulumZone: "pendulum_zone",
}

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

func (d DetailedPhase) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.GoString() + `"`), nil
}

func (d DetailedPhase) GoString() string {
	if l, ok := detailedPhaseNames[d]; ok {
		return l
	}
	return "unknown"
}

var detailedPhaseNames = map[DetailedPhase]string{
	DetailedPhaseDraw:              "draw",
	DetailedPhaseStandby:           "standby",
	DetailedPhaseMain1:             "main1",
	DetailedPhaseBattleStart:       "battle_start",
	DetailedPhaseBattleStep:        "battle_step",
	DetailedPhaseDamage:            "damage",
	DetailedPhaseDamageCalculation: "damage_calculation",
	DetailedPhaseBattle:            "battle",
	DetailedPhaseMain2:             "main2",
	DetailedPhaseEnd:               "end",
}

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

func (p Phase) MarshalJSON() ([]byte, error) {
	return []byte(`"` + p.GoString() + `"`), nil
}

func (p Phase) GoString() string {
	if l, ok := phaseNames[p]; ok {
		return l
	}
	return "unknown"
}

var phaseNames = map[Phase]string{
	PhaseDP: "DP",
	PhaseSP: "SP",
	PhaseM1: "M1",
	PhaseBP: "BP",
	PhaseM2: "M2",
	PhaseEP: "EP",
}

type Position uint

const (
	PositionUnknown Position = iota
	PositionFaceUpAttack
	PositionFaceDownAttack
	PositionFaceUpDefense
	PositionFaceDownDefense
)

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

func (p Position) MarshalJSON() ([]byte, error) {
	return []byte(`"` + p.GoString() + `"`), nil
}

func (p Position) GoString() string {
	if l, ok := positionNames[p]; ok {
		return l
	}
	return "unknown"
}

var positionNames = map[Position]string{
	PositionFaceUpAttack:    "faceup_attack",
	PositionFaceDownAttack:  "facedown_attack",
	PositionFaceUpDefense:   "faceup_defense",
	PositionFaceDownDefense: "facedown_defense",
}
