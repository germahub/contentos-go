// +build tests

package constants

const (
	SingletonId      = 1
	COSChainName     = "contentos"
	COSTokenDecimals = 1000000
	COSInitSupply    = 6500000000 * COSTokenDecimals
	COSConsensusName = "saBFT"
	COSInitMiner     = "initminer"
	COSSysAccount    = "contentos"

	CoinSymbol = "COS"
	VestSymbol = "VEST"

	BlockInterval       = 1  // 1000 ms for one block produce
	BlockProdRepetition = 5 // each producer produces 5 blocks in a row

	NoticeOpPre        = "oppre"
	NoticeOpPost       = "oppost"
	NoticeTrxPre       = "trxpre"
	NoticeTrxPost      = "trxpost"
	NoticeTrxPending   = "trxpending"
	NoticeTrxApplied   = "trxapplyresult"
	NoticeBlockApplied = "blockapply"
	NoticeAddTrx       = "addTrx"
	NoticeCashout      = "rewardCashout"
	//NoticeLIB          = "lastIrreversibleBlock"
	NoticeState        = "blockstate"

	GenesisTime = 0

	MaxTransactionSize = 1024 * 256

	MaxBlockSize           = 1024 * 1024 * 2
	MaxUncommittedBlockNum = 2000
	MinBlockSize           = 115

	InitminerPubKey  = "COS5JVLLcTPhq4Unr194JzWPDNSYGoMcam8yxnsjgRVo3Nb7ioyFW"
	InitminerPrivKey = "4DjYx2KAGh1NP3dai7MZTLUBMMhMBPmwouKE8jhVSESywccpVZ"

	RpcPageSizeLimit = 100

	MaxWitnessCount = 21

	MinVestBalance = 0 * COSTokenDecimals

	PostInvalidId        = 0
	PostMaxDepth         = 8
	PostCashOutDelayBlock = 60 * 10
	//VpDecayTime = 60 * 60 * 24 * 1.5
	VpDecayTime = PostCashOutDelayBlock * 1.5

	MaxBpVoteCount       = 30

	PerVoterCanVoteWitness   = 1
	VoteCountPerVest         = 1

	MaxAccountNameLength     = 16
	MinAccountNameLength     = 6

	// resource limit
	MinStaminaFree       = 0
	DefaultStaminaFree   = 100000
	MaxStaminaFree       = 10000000

	MinTPSExpected       = 100
	DefaultTPSExpected   = 100
	MaxTPSExpected       = 2000

	MinAccountCreateFee      = 1
	DefaultAccountCreateFee  = 1
	MaxAccountCreateFee      = 10000 * COSTokenDecimals

	BlocksPerDay = 24 * 60 * 60 / BlockInterval

	MaxUndoHistory = 10000

	MinVoteInterval = 0 // per 260s 1/1000 vp will restore. Between the 260s any vote operations are valueless if its vp has been exhausted
	MinPostInterval = 0 // 5 * 60 TODO for unit test

	PERCENT = 10000

	VoteRegenerateTime = 10000

	VoteLimitDuringRegenerate = 30


	TrxMaxExpirationTime = 60

	// from total minted
	RewardRateCreator = 7500
	RewardRateBP     = 1500
	RewardRateDapp   = 1000

	// from Creator
	//RewardRateAuthor = 7000
	RewardRateAuthor = 7500
	RewardRateReply = 1500
	RewardRateVoter = 1000
	//RewardRateReport = 500

	ConvertWeeks = 13

	BaseRate               = 1e6
	PowerDownBlockInterval = 100

	ReportCashout = 1000

	// 10 billion
	TotalCurrency = 100 * 1e8

	BlockApplierVersion = 0x00000001

	// resource parameter
	LimitPrecision       = 1000 * 1000
	NetConsumePointNum   = 10
	NetConsumePointDen   = 1
	CpuConsumePointNum   = 1
	CpuConsumePointDen   = 100
	MaxGasPerCall        = 20000 * CpuConsumePointDen
	MaxStaminaPerBlock   = 100000
	WindowSize           = 60 * 60 * 24
	FreeStamina          = 100000
	OneDayStamina        = MaxStaminaPerBlock * WindowSize
	CommonOpStamina      = 100
	StakeFreezeTime      = WindowSize * 3
	TpsWindowSize        = 60
	FreeStaminaOverFlow  = "freeStaminaOverFlow"
	StakeStaminaOverFlow = "stakeStaminaOverFlow"
	EnableResourceControl = true

	MinReputation       = 0
	MaxReputation       = 10000
	DefaultReputation   = 100

	CopyrightUnkown             = 0
	CopyrightInfringement       = 1
	CopyrightConfirmation       = 2

	MaxTicketsPerTurn = 1e5
	PerTicketPrice = 1
	PerTicketPriceStr = "1.000000"
	PerTicketWeight = 1e7
	InitTopN = 500
	//InitEpochDuration = 60 * 60 * 24 * 30
	InitEpochDuration = 600
	MaxTopN = 10000
	MinEpochDuration = 600
	MinTicketPrice = 100

	InitPostWeightedVps = "0"
	InitReplyWeightedVps = "0"
)

var GlobalId int32 = 1