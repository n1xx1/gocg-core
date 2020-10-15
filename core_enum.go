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

	coreLocationOnField coreLocation = 0x0c
	coreLocationAll     coreLocation = 0x3ff
)

type corePosition uint32

const (
	corePositionFaceUpAttack corePosition = 1 << iota
	corePositionFaceDownAttack
	corePositionFaceUpDefense
	corePositionFaceDownDefense
)

const (
	corePositionFaceUp   corePosition = 0b0101
	corePositionFaceDown corePosition = 0b1010
	corePositionAttack   corePosition = 0b0011
	corePositionDefense  corePosition = 0b1100
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

type coreDuelMode uint32

const (
	coreDuelTestMode coreDuelMode = 1 << iota
	coreDuelAttackFirstTurn
	coreDuelUseTrapsInNewChain
	coreDuel6StepBattleStep
	coreDuelPseudoShuffle
	coreDuelTriggerWhenPrivateKnowledge
	coreDuelSimpleAI
	coreDuelRelay
	coreDuelObsoleteIgnition
	coreDuel1stTurnDraw
	coreDuel1FaceUpField
	coreDuelPZone
	coreDuelSeparatePZone
	coreDuelEMZone
	coreDuelFSXMMZone
	coreDuelTrapMonstersNotUseZone
	coreDuelReturnToExtraDeckTriggers
	coreDuelTriggerOnlyInLocation
	coreDuelSPSummonOnceOldNegate
	coreDuelCannotSummonOathOld
	coreDuelNoStandbyPhase
	coreDuelNoMainPhase2
	coreDuel3ColumnsField
	coreDuelDrawUntil5
	coreDuelNoHandLimit
	coreDuelUnlimitedSummons
	coreDuelInvertedQuickPriority
	coreDuelEquipNotSentIfMissingTarget
	coreDuel0AtkDestroyed
	coreDuelStoreAttackReplays
	coreDuelSingleChainInDamageSubStep
	coreDuelReposAfterControlSwitch
)

const (
	coreDuelModeSpeed coreDuelMode = coreDuel3ColumnsField | coreDuelNoMainPhase2 | coreDuelTriggerOnlyInLocation
	coreDuelModeRush  coreDuelMode = coreDuel3ColumnsField | coreDuelNoMainPhase2 | coreDuelNoStandbyPhase | coreDuel1stTurnDraw | coreDuelInvertedQuickPriority | coreDuelDrawUntil5 | coreDuelNoHandLimit | coreDuelUnlimitedSummons | coreDuelTriggerOnlyInLocation
	coreDuelModeMR1   coreDuelMode = coreDuelObsoleteIgnition | coreDuel1stTurnDraw | coreDuel1FaceUpField | coreDuelSPSummonOnceOldNegate | coreDuelReturnToExtraDeckTriggers | coreDuelCannotSummonOathOld
	coreDuelModeGoat  coreDuelMode = coreDuelModeMR1 | coreDuelUseTrapsInNewChain | coreDuel6StepBattleStep | coreDuelTriggerWhenPrivateKnowledge | coreDuelEquipNotSentIfMissingTarget | coreDuel0AtkDestroyed | coreDuelStoreAttackReplays | coreDuelSingleChainInDamageSubStep | coreDuelReposAfterControlSwitch
	coreDuelModeMR2   coreDuelMode = coreDuel1stTurnDraw | coreDuel1FaceUpField | coreDuelSPSummonOnceOldNegate | coreDuelReturnToExtraDeckTriggers | coreDuelCannotSummonOathOld
	coreDuelModeMR3   coreDuelMode = coreDuelPZone | coreDuelSeparatePZone | coreDuelSPSummonOnceOldNegate | coreDuelReturnToExtraDeckTriggers | coreDuelCannotSummonOathOld
	coreDuelModeMR4   coreDuelMode = coreDuelPZone | coreDuelEMZone | coreDuelSPSummonOnceOldNegate | coreDuelReturnToExtraDeckTriggers | coreDuelCannotSummonOathOld
	coreDuelModeMR5   coreDuelMode = coreDuelPZone | coreDuelEMZone | coreDuelFSXMMZone | coreDuelTrapMonstersNotUseZone | coreDuelTriggerOnlyInLocation
)

type processorFlag int

const (
	processorFlagEnd      processorFlag = 0
	processorFlagWaiting  processorFlag = 0x1
	processorFlagContinue processorFlag = 0x2
)

type processor int

const (
	processorAdjust             processor = 1
	processorHint               processor = 2
	processorTurn               processor = 3
	processorRefreshLoc         processor = 5
	processorStartUp            processor = 6
	processorSelectIdleCMD      processor = 10
	processorSelectEffectYN     processor = 11
	processorSelectBattleCMD    processor = 12
	processorSelectYesNo        processor = 13
	processorSelectOption       processor = 14
	processorSelectCard         processor = 15
	processorSelectChain        processor = 16
	processorSelectUnselectCard processor = 17
	processorSelectPlace        processor = 18
	processorSelectPosition     processor = 19
	processorSelectTributeP     processor = 20
	processorSortChain          processor = 21
	processorSelectCounter      processor = 22
	processorSelectSum          processor = 23
	processorSelectDisfield     processor = 24
	processorSortCard           processor = 25
	processorSelectRelease      processor = 26
	processorSelectTribute      processor = 27
	processorPointEvent         processor = 30
	processorQuickEffect        processor = 31
	processorIdleCommand        processor = 32
	processorPhaseEvent         processor = 33
	processorBattleCommand      processor = 34
	processorDamageStep         processor = 35
	processorForcedBattle       processor = 36
	processorAddChain           processor = 40
	processorSolveChain         processor = 42
	processorSolveContinuous    processor = 43
	processorExecuteCost        processor = 44
	processorExecuteOperation   processor = 45
	processorExecuteTarget      processor = 46
	processorDestroy            processor = 50
	processorRelease            processor = 51
	processorSendTo             processor = 52
	processorMoveToField        processor = 53
	processorChangePos          processor = 54
	processorOperationReplace   processor = 55
	processorDestroyReplace     processor = 56
	processorReleaseReplace     processor = 57
	processorSendToReplace      processor = 58
	processorSummonRule         processor = 60
	processorSPSummonRule       processor = 61
	processorSPSummon           processor = 62
	processorFlipSummon         processor = 63
	processorMSet               processor = 64
	processorSSet               processor = 65
	processorSPSummonStep       processor = 66
	processorSSetG              processor = 67
	processorDraw               processor = 70
	processorDamage             processor = 71
	processorRecover            processor = 72
	processorEquip              processor = 73
	processorGetControl         processor = 74
	processorSwapControl        processor = 75
	processorControlAdjust      processor = 76
	processorSelfDestroy        processor = 77
	processorTrapMonsterAdjust  processor = 78
	processorPayLPCost          processor = 80
	processorRemoveCounter      processor = 81
	processorAttackDisable      processor = 82
	processorActivateEffect     processor = 83
	processorAnnounceRace       processor = 110
	processorAnnounceAttrib     processor = 111
	processorAnnounceLevel      processor = 112
	processorAnnounceCard       processor = 113
	processorAnnounceType       processor = 114
	processorAnnounceNumber     processor = 115
	processorAnnounceCoin       processor = 116
	processorTossDice           processor = 117
	processorTossCoin           processor = 118
	processorRockPaperScissors  processor = 119
	processorSelectFusion       processor = 131
	processorDiscardHand        processor = 150
	processorDiscardDeck        processor = 151
	processorSortDeck           processor = 152
	processorRemoveOverlay      processor = 160
)

type cardType uint32

const (
	cardTypeMonster     cardType = 0x1
	cardTypeSpell       cardType = 0x2
	cardTypeTrap        cardType = 0x4
	cardTypeNormal      cardType = 0x10
	cardTypeEffect      cardType = 0x20
	cardTypeFusion      cardType = 0x40
	cardTypeRitual      cardType = 0x80
	cardTypeTrapMonster cardType = 0x100
	cardTypeSpirit      cardType = 0x200
	cardTypeUnion       cardType = 0x400
	cardTypeGemini      cardType = 0x800
	cardTypeTuner       cardType = 0x1000
	cardTypeSynchro     cardType = 0x2000
	cardTypeToken       cardType = 0x4000
	cardTypeQuickPlay   cardType = 0x10000
	cardTypeContinuous  cardType = 0x20000
	cardTypeEquip       cardType = 0x40000
	cardTypeField       cardType = 0x80000
	cardTypeCounter     cardType = 0x100000
	cardTypeFlip        cardType = 0x200000
	cardTypeToon        cardType = 0x400000
	cardTypeXyz         cardType = 0x800000
	cardTypePendulum    cardType = 0x1000000
	cardTypeSPSummon    cardType = 0x2000000
	cardTypeLink        cardType = 0x4000000
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

	coreQueryEnd coreQuery = 0x80000000
)
