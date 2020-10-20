package lib

type Location uint32

const (
	LocationDeck Location = 1 << iota
	LocationHand
	LocationMZone
	LocationSZone
	LocationGrave
	LocationRemoved
	LocationExtra
	LocationOverlay
	LocationFZone
	LocationPZone

	LocationOnField Location = 0x0c
	LocationAll     Location = 0x3ff
)

type Position uint32

const (
	PositionFaceUpAttack Position = 1 << iota
	PositionFaceDownAttack
	PositionFaceUpDefense
	PositionFaceDownDefense
)

const (
	PositionFaceUp   Position = 0b0101
	PositionFaceDown Position = 0b1010
	PositionAttack   Position = 0b0011
	PositionDefense  Position = 0b1100
)

type Attribute uint32

const (
	AttributeEarth Attribute = 1 << iota
	AttributeWater
	AttributeFire
	AttributeWind
	AttributeLight
	AttributeDark
	AttributeDivine
)

type Race uint32

const (
	RaceWarrior Race = 1 << iota
	RaceSpellCaster
	RaceFairy
	RaceFiend
	RaceZombie
	RaceMachine
	RaceAqua
	RacePyro
	RaceRock
	RaceWingedBeast
	RacePlant
	RaceInsect
	RaceThunder
	RaceDragon
	RaceBeast
	RaceBeastWarrior
	RaceDinosaur
	RaceFish
	RaceSeaSerpent
	RaceReptile
	RacePsychic
	RaceDivine
	RaceCreatorGod
	RaceWyrm
	RaceCyberse
)

type DuelMode uint32

const (
	DuelTestMode DuelMode = 1 << iota
	DuelAttackFirstTurn
	DuelUseTrapsInNewChain
	Duel6StepBattleStep
	DuelPseudoShuffle
	DuelTriggerWhenPrivateKnowledge
	DuelSimpleAI
	DuelRelay
	DuelObsoleteIgnition
	Duel1stTurnDraw
	Duel1FaceUpField
	DuelPZone
	DuelSeparatePZone
	DuelEMZone
	DuelFSXMMZone
	DuelTrapMonstersNotUseZone
	DuelReturnToExtraDeckTriggers
	DuelTriggerOnlyInLocation
	DuelSPSummonOnceOldNegate
	DuelCannotSummonOathOld
	DuelNoStandbyPhase
	DuelNoMainPhase2
	Duel3ColumnsField
	DuelDrawUntil5
	DuelNoHandLimit
	DuelUnlimitedSummons
	DuelInvertedQuickPriority
	DuelEquipNotSentIfMissingTarget
	Duel0AtkDestroyed
	DuelStoreAttackReplays
	DuelSingleChainInDamageSubStep
	DuelReposAfterControlSwitch
)

const (
	DuelModeSpeed DuelMode = Duel3ColumnsField | DuelNoMainPhase2 | DuelTriggerOnlyInLocation
	DuelModeRush  DuelMode = Duel3ColumnsField | DuelNoMainPhase2 | DuelNoStandbyPhase | Duel1stTurnDraw | DuelInvertedQuickPriority | DuelDrawUntil5 | DuelNoHandLimit | DuelUnlimitedSummons | DuelTriggerOnlyInLocation
	DuelModeMR1   DuelMode = DuelObsoleteIgnition | Duel1stTurnDraw | Duel1FaceUpField | DuelSPSummonOnceOldNegate | DuelReturnToExtraDeckTriggers | DuelCannotSummonOathOld
	DuelModeGoat  DuelMode = DuelModeMR1 | DuelUseTrapsInNewChain | Duel6StepBattleStep | DuelTriggerWhenPrivateKnowledge | DuelEquipNotSentIfMissingTarget | Duel0AtkDestroyed | DuelStoreAttackReplays | DuelSingleChainInDamageSubStep | DuelReposAfterControlSwitch
	DuelModeMR2   DuelMode = Duel1stTurnDraw | Duel1FaceUpField | DuelSPSummonOnceOldNegate | DuelReturnToExtraDeckTriggers | DuelCannotSummonOathOld
	DuelModeMR3   DuelMode = DuelPZone | DuelSeparatePZone | DuelSPSummonOnceOldNegate | DuelReturnToExtraDeckTriggers | DuelCannotSummonOathOld
	DuelModeMR4   DuelMode = DuelPZone | DuelEMZone | DuelSPSummonOnceOldNegate | DuelReturnToExtraDeckTriggers | DuelCannotSummonOathOld
	DuelModeMR5   DuelMode = DuelPZone | DuelEMZone | DuelFSXMMZone | DuelTrapMonstersNotUseZone | DuelTriggerOnlyInLocation
)

type ProcessorFlag int

const (
	ProcessorFlagEnd ProcessorFlag = iota
	ProcessorFlagWaiting
	ProcessorFlagContinue
)

type CardType uint32

const (
	CardTypeMonster CardType = 1 << iota
	CardTypeSpell
	CardTypeTrap
	CardTypeUnused1
	CardTypeNormal
	CardTypeEffect
	CardTypeFusion
	CardTypeRitual
	CardTypeTrapMonster
	CardTypeSpirit
	CardTypeUnion
	CardTypeGemini
	CardTypeTuner
	CardTypeSynchro
	CardTypeToken
	CardTypeUnused2
	CardTypeQuickPlay
	CardTypeContinuous
	CardTypeEquip
	CardTypeField
	CardTypeCounter
	CardTypeFlip
	CardTypeToon
	CardTypeXyz
	CardTypePendulum
	CardTypeSPSummon
	CardTypeLink
)

type Message uint8

const (
	MessageRetry              Message = 1
	MessageHint               Message = 2
	MessageWaiting            Message = 3
	MessageStart              Message = 4
	MessageWin                Message = 5
	MessageUpdateData         Message = 6
	MessageUpdateCard         Message = 7
	MessageRequestDeck        Message = 8
	MessageSelectBattleCMD    Message = 10
	MessageSelectIdleCMD      Message = 11
	MessageSelectEffectYN     Message = 12
	MessageSelectYesNo        Message = 13
	MessageSelectOption       Message = 14
	MessageSelectCard         Message = 15
	MessageSelectChain        Message = 16
	MessageSelectPlace        Message = 18
	MessageSelectPosition     Message = 19
	MessageSelectTribute      Message = 20
	MessageSortChain          Message = 21
	MessageSelectCounter      Message = 22
	MessageSelectSum          Message = 23
	MessageSelectDisfield     Message = 24
	MessageSortCard           Message = 25
	MessageSelectUnselectCard Message = 26
	MessageConfirmDeckTop     Message = 30
	MessageConfirmCards       Message = 31
	MessageShuffleDeck        Message = 32
	MessageShuffleHand        Message = 33
	MessageRefreshDeck        Message = 34
	MessageSwapGraveDeck      Message = 35
	MessageShuffleSetCard     Message = 36
	MessageReverseDeck        Message = 37
	MessageDeckTop            Message = 38
	MessageShuffleExtra       Message = 39
	MessageNewTurn            Message = 40
	MessageNewPhase           Message = 41
	MessageConfirmExtraTop    Message = 42
	MessageMove               Message = 50
	MessagePosChange          Message = 53
	MessageSet                Message = 54
	MessageSwap               Message = 55
	MessageFieldDisabled      Message = 56
	MessageSummoning          Message = 60
	MessageSummoned           Message = 61
	MessageSPSummoning        Message = 62
	MessageSPSummoned         Message = 63
	MessageFlipSummoning      Message = 64
	MessageFlipSummoned       Message = 65
	MessageChaining           Message = 70
	MessageChained            Message = 71
	MessageChainSolving       Message = 72
	MessageChainSolved        Message = 73
	MessageChainEnd           Message = 74
	MessageChainNegated       Message = 75
	MessageChainDisabled      Message = 76
	MessageCardSelected       Message = 80
	MessageRandomSelected     Message = 81
	MessageBecomeTarget       Message = 83
	MessageDraw               Message = 90
	MessageDamage             Message = 91
	MessageRecover            Message = 92
	MessageEquip              Message = 93
	MessageLPUpdate           Message = 94
	MessageUnequip            Message = 95
	MessageCardTarget         Message = 96
	MessageCancelTarget       Message = 97
	MessagePayLPCost          Message = 100
	MessageAddCounter         Message = 101
	MessageRemoveCounter      Message = 102
	MessageAttack             Message = 110
	MessageBattle             Message = 111
	MessageAttackDisabled     Message = 112
	MessageDamageStepStart    Message = 113
	MessageDamageStepEnd      Message = 114
	MessageMissedEffect       Message = 120
	MessageBeChainTarget      Message = 121
	MessageCreateRelation     Message = 122
	MessageReleaseRelation    Message = 123
	MessageTossCoin           Message = 130
	MessageTossDice           Message = 131
	MessageRockPaperScissors  Message = 132
	MessageHandRes            Message = 133
	MessageAnnounceRace       Message = 140
	MessageAnnounceAttribute  Message = 141
	MessageAnnounceCard       Message = 142
	MessageAnnounceNumber     Message = 143
	MessageCardHint           Message = 160
	MessageTagSwap            Message = 161
	MessageReloadField        Message = 162
	MessageAIName             Message = 163
	MessageShowHint           Message = 164
	MessagePlayerHint         Message = 165
	MessageMatchKill          Message = 170
	MessageCustomMessage      Message = 180
	MessageRemoveCards        Message = 190
)

type Phase uint16

const (
	PhaseDraw Phase = 1 << iota
	PhaseStandby
	PhaseMain1
	PhaseBattleStart
	PhaseBattleStep
	PhaseDamage
	PhaseDamageCalculation
	PhaseBattle
	PhaseMain2
	PhaseEnd
)

type Query uint32

const (
	QueryCode Query = 1 << iota
	QueryPosition
	QueryAlias
	QueryType
	QueryLevel
	QueryRank
	QueryAttribute
	QueryRace
	QueryAttack
	QueryDefense
	QueryBaseAttack
	QueryBaseDefense
	QueryReason
	QueryReasonCard
	QueryEquipCard
	QueryTargetCard
	QueryOverlayCard
	QueryCounters
	QueryOwner
	QueryStatus
	QueryIsPublic
	QueryLScale
	QueryRScale
	QueryLink
	QueryIsHidden
	QueryCover

	QueryEnd Query = 0x80000000
)

type LinkMarker uint32

const (
	LinkMarkerBottomLeft LinkMarker = 1 << iota
	LinkMarkerBottom
	LinkMarkerBottomRight
	LinkMarkerLeft
	LinkMarkerUnused
	LinkMarkerRight
	LinkMarkerTopLeft
	LinkMarkerTop
	LinkMarkerTopRight
)
