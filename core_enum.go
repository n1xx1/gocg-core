package ocgcore

type coreLocation uint32

const (
	coreLocationDeck coreLocation = 1 << iota
	coreLocationHand
	coreLocationMZone
	coreLocationSZone
	coreLocationGrave
	coreLocationRemoved
	coreLocationExtra
	coreLocationOverlay
	coreLocationFZone
	coreLocationPZone

	CoreLocationOnField coreLocation = 0x0c
	CoreLocationAll     coreLocation = 0x3ff
)

type corePosition uint32

const (
	corePositionFaceUpAttack corePosition = 1 << iota
	corePositionFaceDownAttack
	corePositionFaceUpDefense
	corePositionFaceDownDefense
)

const (
	corePositionFaceUp   corePosition = 0x5
	corePositionFaceDown corePosition = 0xa
	corePositionAttack   corePosition = 0x3
	corePositionDefense  corePosition = 0xc
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

type CoreDuelMode uint32

const (
	CoreDuelTestMode                    CoreDuelMode = 0x01
	CoreDuelAttackFirstTurn             CoreDuelMode = 0x02
	CoreDuelUseTrapsInNewChain          CoreDuelMode = 0x04
	CoreDuel6StepBattleStep             CoreDuelMode = 0x08
	CoreDuelPseudoShuffle               CoreDuelMode = 0x10
	CoreDuelTriggerWhenPrivateKnowledge CoreDuelMode = 0x20
	CoreDuelSimpleAI                    CoreDuelMode = 0x40
	CoreDuelRelay                       CoreDuelMode = 0x80
	CoreDuelObsoleteIgnition            CoreDuelMode = 0x100
	CoreDuel1stTurnDraw                 CoreDuelMode = 0x200
	CoreDuel1FaceUpField                CoreDuelMode = 0x400
	CoreDuelPZone                       CoreDuelMode = 0x800
	CoreDuelSeparatePZone               CoreDuelMode = 0x1000
	CoreDuelEMZone                      CoreDuelMode = 0x2000
	CoreDuelFSXMMZone                   CoreDuelMode = 0x4000
	CoreDuelTrapMonstersNotUseZone      CoreDuelMode = 0x8000
	CoreDuelReturnToExtraDeckTriggers   CoreDuelMode = 0x10000
	CoreDuelTriggerOnlyInLocation       CoreDuelMode = 0x20000
	CoreDuelSPSummonOnceOldNegate       CoreDuelMode = 0x40000
	CoreDuelCannotSummonOathOld         CoreDuelMode = 0x80000
	CoreDuelNoStandbyPhase              CoreDuelMode = 0x100000
	CoreDuelNoMainPhase2                CoreDuelMode = 0x200000
	CoreDuel3ColumnsField               CoreDuelMode = 0x400000
	CoreDuelDrawUntil5                  CoreDuelMode = 0x800000
	CoreDuelNoHandLimit                 CoreDuelMode = 0x1000000
	CoreDuelUnlimitedSummons            CoreDuelMode = 0x2000000
	CoreDuelInvertedQuickPriority       CoreDuelMode = 0x4000000
	CoreDuelEquipNotSentIfMissingTarget CoreDuelMode = 0x8000000
	CoreDuel0AtkDestroyed               CoreDuelMode = 0x10000000
	CoreDuelStoreAttackReplays          CoreDuelMode = 0x20000000
	CoreDuelSingleChainInDamageSubStep  CoreDuelMode = 0x40000000
	CoreDuelReposAfterControlSwitch     CoreDuelMode = 0x80000000
)
const (
	CoreDuelModeSpeed CoreDuelMode = CoreDuel3ColumnsField | CoreDuelNoMainPhase2 | CoreDuelTriggerOnlyInLocation
	CoreDuelModeRush  CoreDuelMode = CoreDuel3ColumnsField | CoreDuelNoMainPhase2 | CoreDuelNoStandbyPhase | CoreDuel1stTurnDraw | CoreDuelInvertedQuickPriority | CoreDuelDrawUntil5 | CoreDuelNoHandLimit | CoreDuelUnlimitedSummons | CoreDuelTriggerOnlyInLocation
	CoreDuelModeMR1   CoreDuelMode = CoreDuelObsoleteIgnition | CoreDuel1stTurnDraw | CoreDuel1FaceUpField | CoreDuelSPSummonOnceOldNegate | CoreDuelReturnToExtraDeckTriggers | CoreDuelCannotSummonOathOld
	CoreDuelModeGoat  CoreDuelMode = CoreDuelModeMR1 | CoreDuelUseTrapsInNewChain | CoreDuel6StepBattleStep | CoreDuelTriggerWhenPrivateKnowledge | CoreDuelEquipNotSentIfMissingTarget | CoreDuel0AtkDestroyed | CoreDuelStoreAttackReplays | CoreDuelSingleChainInDamageSubStep | CoreDuelReposAfterControlSwitch
	CoreDuelModeMR2   CoreDuelMode = CoreDuel1stTurnDraw | CoreDuel1FaceUpField | CoreDuelSPSummonOnceOldNegate | CoreDuelReturnToExtraDeckTriggers | CoreDuelCannotSummonOathOld
	CoreDuelModeMR3   CoreDuelMode = CoreDuelPZone | CoreDuelSeparatePZone | CoreDuelSPSummonOnceOldNegate | CoreDuelReturnToExtraDeckTriggers | CoreDuelCannotSummonOathOld
	CoreDuelModeMR4   CoreDuelMode = CoreDuelPZone | CoreDuelEMZone | CoreDuelSPSummonOnceOldNegate | CoreDuelReturnToExtraDeckTriggers | CoreDuelCannotSummonOathOld
	CoreDuelModeMR5   CoreDuelMode = CoreDuelPZone | CoreDuelEMZone | CoreDuelFSXMMZone | CoreDuelTrapMonstersNotUseZone | CoreDuelTriggerOnlyInLocation
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

type coreMessage uint8

const (
	coreMessageRetry              coreMessage = 1
	coreMessageHint               coreMessage = 2
	coreMessageWaiting            coreMessage = 3
	coreMessageStart              coreMessage = 4
	coreMessageWin                coreMessage = 5
	coreMessageUpdateData         coreMessage = 6
	coreMessageUpdateCard         coreMessage = 7
	coreMessageRequestDeck        coreMessage = 8
	coreMessageSelectBattleCMD    coreMessage = 10
	coreMessageSelectIdleCMD      coreMessage = 11
	coreMessageSelectEffectYN     coreMessage = 12
	coreMessageSelectYesNo        coreMessage = 13
	coreMessageSelectOption       coreMessage = 14
	coreMessageSelectCard         coreMessage = 15
	coreMessageSelectChain        coreMessage = 16
	coreMessageSelectPlace        coreMessage = 18
	coreMessageSelectPosition     coreMessage = 19
	coreMessageSelectTribute      coreMessage = 20
	coreMessageSortChain          coreMessage = 21
	coreMessageSelectCounter      coreMessage = 22
	coreMessageSelectSum          coreMessage = 23
	coreMessageSelectDisfield     coreMessage = 24
	coreMessageSortCard           coreMessage = 25
	coreMessageSelectUnselectCard coreMessage = 26
	coreMessageConfirmDeckTop     coreMessage = 30
	coreMessageConfirmCards       coreMessage = 31
	coreMessageShuffleDeck        coreMessage = 32
	coreMessageShuffleHand        coreMessage = 33
	coreMessageRefreshDeck        coreMessage = 34
	coreMessageSwapGraveDeck      coreMessage = 35
	coreMessageShuffleSetCard     coreMessage = 36
	coreMessageReverseDeck        coreMessage = 37
	coreMessageDeckTop            coreMessage = 38
	coreMessageShuffleExtra       coreMessage = 39
	coreMessageNewTurn            coreMessage = 40
	coreMessageNewPhase           coreMessage = 41
	coreMessageConfirmExtraTop    coreMessage = 42
	coreMessageMove               coreMessage = 50
	coreMessagePosChange          coreMessage = 53
	coreMessageSet                coreMessage = 54
	coreMessageSwap               coreMessage = 55
	coreMessageFieldDisabled      coreMessage = 56
	coreMessageSummoning          coreMessage = 60
	coreMessageSummoned           coreMessage = 61
	coreMessageSPSummoning        coreMessage = 62
	coreMessageSPSummoned         coreMessage = 63
	coreMessageFlipSummoning      coreMessage = 64
	coreMessageFlipSummoned       coreMessage = 65
	coreMessageChaining           coreMessage = 70
	coreMessageChained            coreMessage = 71
	coreMessageChainSolving       coreMessage = 72
	coreMessageChainSolved        coreMessage = 73
	coreMessageChainEnd           coreMessage = 74
	coreMessageChainNegated       coreMessage = 75
	coreMessageChainDisabled      coreMessage = 76
	coreMessageCardSelected       coreMessage = 80
	coreMessageRandomSelected     coreMessage = 81
	coreMessageBecomeTarget       coreMessage = 83
	coreMessageDraw               coreMessage = 90
	coreMessageDamage             coreMessage = 91
	coreMessageRecover            coreMessage = 92
	coreMessageEquip              coreMessage = 93
	coreMessageLPUpdate           coreMessage = 94
	coreMessageUnequip            coreMessage = 95
	coreMessageCardTarget         coreMessage = 96
	coreMessageCancelTarget       coreMessage = 97
	coreMessagePayLPCost          coreMessage = 100
	coreMessageAddCounter         coreMessage = 101
	coreMessageRemoveCounter      coreMessage = 102
	coreMessageAttack             coreMessage = 110
	coreMessageBattle             coreMessage = 111
	coreMessageAttackDisabled     coreMessage = 112
	coreMessageDamageStepStart    coreMessage = 113
	coreMessageDamageStepEnd      coreMessage = 114
	coreMessageMissedEffect       coreMessage = 120
	coreMessageBeChainTarget      coreMessage = 121
	coreMessageCreateRelation     coreMessage = 122
	coreMessageReleaseRelation    coreMessage = 123
	coreMessageTossCoin           coreMessage = 130
	coreMessageTossDice           coreMessage = 131
	coreMessageRockPaperScissors  coreMessage = 132
	coreMessageHandRes            coreMessage = 133
	coreMessageAnnounceRace       coreMessage = 140
	coreMessageAnnounceAttribute  coreMessage = 141
	coreMessageAnnounceCard       coreMessage = 142
	coreMessageAnnounceNumber     coreMessage = 143
	coreMessageCardHint           coreMessage = 160
	coreMessageTagSwap            coreMessage = 161
	coreMessageReloadField        coreMessage = 162
	coreMessageAIName             coreMessage = 163
	coreMessageShowHint           coreMessage = 164
	coreMessagePlayerHint         coreMessage = 165
	coreMessageMatchKill          coreMessage = 170
	coreMessageCustomMessage      coreMessage = 180
	coreMessageRemoveCards        coreMessage = 190
)

type corePhase uint16

const (
	corePhaseDraw corePhase = 1 << iota
	corePhaseStandby
	corePhaseMain1
	corePhaseBattleStart
	corePhaseBattleStep
	corePhaseDamage
	corePhaseDamageCalculation
	corePhaseBattle
	corePhaseMain2
	corePhaseEnd
)

type coreQuery uint32

const (
	coreQueryCode coreQuery = 1 << iota
	coreQueryPosition
	coreQueryAlias
	coreQueryType
	coreQueryLevel
	coreQueryRank
	coreQueryAttribute
	coreQueryRace
	coreQueryAttack
	coreQueryDefense
	coreQueryBaseAttack
	coreQueryBaseDefense
	coreQueryReason
	coreQueryReasonCard
	coreQueryEquipCard
	coreQueryTargetCard
	coreQueryOverlayCard
	coreQueryCounters
	coreQueryOwner
	coreQueryStatus
	coreQueryIsPublic
	coreQueryLScale
	coreQueryRScale
	coreQueryLink
	coreQueryIsHidden
	coreQueryCover
)
