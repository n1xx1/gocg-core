package ocgcore

type Location uint32

const (
	LocationDeck    Location = 0x01
	LocationHand    Location = 0x02
	LocationMZone   Location = 0x04
	LocationSZone   Location = 0x08
	LocationGrave   Location = 0x10
	LocationRemoved Location = 0x20
	LocationExtra   Location = 0x40
	LocationOverlay Location = 0x80
	LocationOnField Location = 0x0c
	LocationFZone   Location = 0x100
	LocationPZone   Location = 0x200
	LocationAll     Location = 0x3ff
)

type Position uint32

const (
	PositionFaceUpAttack    Position = 0x1
	PositionFaceDownAttack  Position = 0x2
	PositionFaceUpDefense   Position = 0x4
	PositionFaceDownDefense Position = 0x8
	PositionFaceUp          Position = 0x5
	PositionFaceDown        Position = 0xa
	PositionAttack          Position = 0x3
	PositionDefense         Position = 0xc
)

type Attribute uint32

const (
	AttributeEarth  Attribute = 0x01
	AttributeWater  Attribute = 0x02
	AttributeFire   Attribute = 0x04
	AttributeWind   Attribute = 0x08
	AttributeLight  Attribute = 0x10
	AttributeDark   Attribute = 0x20
	AttributeDivine Attribute = 0x40
)

type Race uint32

const (
	RaceWarrior      Race = 0x1
	RaceSpellCaster  Race = 0x2
	RaceFairy        Race = 0x4
	RaceFiend        Race = 0x8
	RaceZombie       Race = 0x10
	RaceMachine      Race = 0x20
	RaceAqua         Race = 0x40
	RacePyro         Race = 0x80
	RaceRock         Race = 0x100
	RaceWingedBeast  Race = 0x200
	RacePlant        Race = 0x400
	RaceInsect       Race = 0x800
	RaceThunder      Race = 0x1000
	RaceDragon       Race = 0x2000
	RaceBeast        Race = 0x4000
	RaceBeastWarrior Race = 0x8000
	RaceDinosaur     Race = 0x10000
	RaceFish         Race = 0x20000
	RaceSeaSerpent   Race = 0x40000
	RaceReptile      Race = 0x80000
	RacePsychic      Race = 0x100000
	RaceDivine       Race = 0x200000
	RaceCreatorGod   Race = 0x400000
	RaceWyrm         Race = 0x800000
	RaceCyberse      Race = 0x1000000
)

type DuelMode uint32

const (
	DuelTestMode                    DuelMode = 0x01
	DuelAttackFirstTurn             DuelMode = 0x02
	DuelUseTrapsInNewChain          DuelMode = 0x04
	Duel6StepBattleStep             DuelMode = 0x08
	DuelPseudoShuffle               DuelMode = 0x10
	DuelTriggerWhenPrivateKnowledge DuelMode = 0x20
	DuelSimpleAI                    DuelMode = 0x40
	DuelRelay                       DuelMode = 0x80
	DuelObsoleteIgnition            DuelMode = 0x100
	Duel1stTurnDraw                 DuelMode = 0x200
	Duel1FaceUpField                DuelMode = 0x400
	DuelPZone                       DuelMode = 0x800
	DuelSeparatePZone               DuelMode = 0x1000
	DuelEMZone                      DuelMode = 0x2000
	DuelFSXMMZone                   DuelMode = 0x4000
	DuelTrapMonstersNotUseZone      DuelMode = 0x8000
	DuelReturnToExtraDeckTriggers   DuelMode = 0x10000
	DuelTriggerOnlyInLocation       DuelMode = 0x20000
	DuelSPSummonOnceOldNegate       DuelMode = 0x40000
	DuelCannotSummonOathOld         DuelMode = 0x80000
	DuelNoStandbyPhase              DuelMode = 0x100000
	DuelNoMainPhase2                DuelMode = 0x200000
	Duel3ColumnsField               DuelMode = 0x400000
	DuelDrawUntil5                  DuelMode = 0x800000
	DuelNoHandLimit                 DuelMode = 0x1000000
	DuelUnlimitedSummons            DuelMode = 0x2000000
	DuelInvertedQuickPriority       DuelMode = 0x4000000
	DuelEquipNotSentIfMissingTarget DuelMode = 0x8000000
	Duel0AtkDestroyed               DuelMode = 0x10000000
	DuelStoreAttackReplays          DuelMode = 0x20000000
	DuelSingleChainInDamageSubStep  DuelMode = 0x40000000
	DuelReposAfterControlSwitch     DuelMode = 0x80000000
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
	ProcessorFlagEnd      ProcessorFlag = 0
	ProcessorFlagWaiting  ProcessorFlag = 0x1
	ProcessorFlagContinue ProcessorFlag = 0x2
)

type Processor int

const (
	ProcessorAdjust             Processor = 1
	ProcessorHint               Processor = 2
	ProcessorTurn               Processor = 3
	ProcessorRefreshLoc         Processor = 5
	ProcessorStartUp            Processor = 6
	ProcessorSelectIdleCMD      Processor = 10
	ProcessorSelectEffectYN     Processor = 11
	ProcessorSelectBattleCMD    Processor = 12
	ProcessorSelectYesNo        Processor = 13
	ProcessorSelectOption       Processor = 14
	ProcessorSelectCard         Processor = 15
	ProcessorSelectChain        Processor = 16
	ProcessorSelectUnselectCard Processor = 17
	ProcessorSelectPlace        Processor = 18
	ProcessorSelectPosition     Processor = 19
	ProcessorSelectTributeP     Processor = 20
	ProcessorSortChain          Processor = 21
	ProcessorSelectCounter      Processor = 22
	ProcessorSelectSum          Processor = 23
	ProcessorSelectDisfield     Processor = 24
	ProcessorSortCard           Processor = 25
	ProcessorSelectRelease      Processor = 26
	ProcessorSelectTribute      Processor = 27
	ProcessorPointEvent         Processor = 30
	ProcessorQuickEffect        Processor = 31
	ProcessorIdleCommand        Processor = 32
	ProcessorPhaseEvent         Processor = 33
	ProcessorBattleCommand      Processor = 34
	ProcessorDamageStep         Processor = 35
	ProcessorForcedBattle       Processor = 36
	ProcessorAddChain           Processor = 40
	ProcessorSolveChain         Processor = 42
	ProcessorSolveContinuous    Processor = 43
	ProcessorExecuteCost        Processor = 44
	ProcessorExecuteOperation   Processor = 45
	ProcessorExecuteTarget      Processor = 46
	ProcessorDestroy            Processor = 50
	ProcessorRelease            Processor = 51
	ProcessorSendTo             Processor = 52
	ProcessorMoveToField        Processor = 53
	ProcessorChangePos          Processor = 54
	ProcessorOperationReplace   Processor = 55
	ProcessorDestroyReplace     Processor = 56
	ProcessorReleaseReplace     Processor = 57
	ProcessorSendToReplace      Processor = 58
	ProcessorSummonRule         Processor = 60
	ProcessorSPSummonRule       Processor = 61
	ProcessorSPSummon           Processor = 62
	ProcessorFlipSummon         Processor = 63
	ProcessorMSet               Processor = 64
	ProcessorSSet               Processor = 65
	ProcessorSPSummonStep       Processor = 66
	ProcessorSSetG              Processor = 67
	ProcessorDraw               Processor = 70
	ProcessorDamage             Processor = 71
	ProcessorRecover            Processor = 72
	ProcessorEquip              Processor = 73
	ProcessorGetControl         Processor = 74
	ProcessorSwapControl        Processor = 75
	ProcessorControlAdjust      Processor = 76
	ProcessorSelfDestroy        Processor = 77
	ProcessorTrapMonsterAdjust  Processor = 78
	ProcessorPayLPCost          Processor = 80
	ProcessorRemoveCounter      Processor = 81
	ProcessorAttackDisable      Processor = 82
	ProcessorActivateEffect     Processor = 83
	ProcessorAnnounceRace       Processor = 110
	ProcessorAnnounceAttrib     Processor = 111
	ProcessorAnnounceLevel      Processor = 112
	ProcessorAnnounceCard       Processor = 113
	ProcessorAnnounceType       Processor = 114
	ProcessorAnnounceNumber     Processor = 115
	ProcessorAnnounceCoin       Processor = 116
	ProcessorTossDice           Processor = 117
	ProcessorTossCoin           Processor = 118
	ProcessorRockPaperScissors  Processor = 119
	ProcessorSelectFusion       Processor = 131
	ProcessorDiscardHand        Processor = 150
	ProcessorDiscardDeck        Processor = 151
	ProcessorSortDeck           Processor = 152
	ProcessorRemoveOverlay      Processor = 160
)

type Type uint32

const (
	TypeMonster     Type = 0x1
	TypeSpell       Type = 0x2
	TypeTrap        Type = 0x4
	TypeNormal      Type = 0x10
	TypeEffect      Type = 0x20
	TypeFusion      Type = 0x40
	TypeRitual      Type = 0x80
	TypeTrapMonster Type = 0x100
	TypeSpirit      Type = 0x200
	TypeUnion       Type = 0x400
	TypeGemini      Type = 0x800
	TypeTuner       Type = 0x1000
	TypeSynchro     Type = 0x2000
	TypeToken       Type = 0x4000
	TypeQuickPlay   Type = 0x10000
	TypeContinuous  Type = 0x20000
	TypeEquip       Type = 0x40000
	TypeField       Type = 0x80000
	TypeCounter     Type = 0x100000
	TypeFlip        Type = 0x200000
	TypeToon        Type = 0x400000
	TypeXyz         Type = 0x800000
	TypePendulum    Type = 0x1000000
	TypeSPSummon    Type = 0x2000000
	TypeLink        Type = 0x4000000
)
