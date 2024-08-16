package manifold

import (
	"encoding/json"
)

// ProfitCached holds cached profit data for different time periods.
type ProfitCached struct {
	Daily   float64 `json:"daily"`   // Daily profit
	Weekly  float64 `json:"weekly"`  // Weekly profit
	Monthly float64 `json:"monthly"` // Monthly profit
	AllTime float64 `json:"allTime"` // All-time profit
}

// User represents a user in the system with various attributes.
type User struct {
	ID                   string       `json:"id"`                             // Unique identifier for the user
	CreatedTime          int64        `json:"createdTime"`                    // Timestamp when the user was created
	Name                 string       `json:"name"`                           // Full name of the user
	Username             string       `json:"username"`                       // Username of the user
	URL                  string       `json:"url"`                            // URL to the user's profile
	AvatarUrl            *string      `json:"avatarUrl,omitempty"`            // URL to the user's avatar image (optional)
	Bio                  *string      `json:"bio,omitempty"`                  // User's biography (optional)
	BannerUrl            *string      `json:"bannerUrl,omitempty"`            // URL to the user's banner image (optional)
	Website              *string      `json:"website,omitempty"`              // User's personal website (optional)
	TwitterHandle        *string      `json:"twitterHandle,omitempty"`        // User's Twitter handle (optional)
	DiscordHandle        *string      `json:"discordHandle,omitempty"`        // User's Discord handle (optional)
	IsBot                *bool        `json:"isBot,omitempty"`                // Indicates if the user is a bot (optional)
	IsAdmin              *bool        `json:"isAdmin,omitempty"`              // Indicates if the user is an admin (optional)
	IsTrustworthy        *bool        `json:"isTrustworthy,omitempty"`        // Indicates if the user is trustworthy (optional)
	IsBannedFromPosting  *bool        `json:"isBannedFromPosting,omitempty"`  // Indicates if the user is banned from posting (optional)
	UserDeleted          *bool        `json:"userDeleted,omitempty"`          // Indicates if the user has been deleted (optional)
	Balance              float64      `json:"balance"`                        // Current balance of the user
	TotalDeposits        float64      `json:"totalDeposits"`                  // Total deposits made by the user
	LastBetTime          *int64       `json:"lastBetTime,omitempty"`          // Timestamp of the user's last bet (optional)
	CurrentBettingStreak *int         `json:"currentBettingStreak,omitempty"` // User's current betting streak (optional)
	ProfitCached         ProfitCached `json:"profitCached"`                   // Cached profit data for the user
}

// DisplayUser represents a simplified view of a user, often used for display purposes.
type DisplayUser struct {
	ID        string  `json:"id"`                  // Unique identifier for the user
	Name      string  `json:"name"`                // Full name of the user
	Username  string  `json:"username"`            // Username of the user
	AvatarURL *string `json:"avatarUrl,omitempty"` // URL to the user's avatar image (optional)
}

// Group represents a group with various attributes.
type Group struct {
	ID                string          `json:"id"`                          // Unique identifier for the group
	Slug              string          `json:"slug"`                        // Slug for the group (usually URL-friendly)
	Name              string          `json:"name"`                        // Name of the group
	About             json.RawMessage `json:"about,omitempty"`             // Detailed information about the group (optional)
	CreatorID         string          `json:"creatorId"`                   // ID of the group creator
	CreatedTime       int64           `json:"createdTime"`                 // Timestamp when the group was created
	AnyoneCanJoin     *bool           `json:"anyoneCanJoin,omitempty"`     // Indicates if anyone can join the group (optional)
	TotalMembers      int             `json:"totalMembers"`                // Total number of members in the group
	PostIDs           []string        `json:"postIds"`                     // List of post IDs associated with the group
	CachedLeaderboard *Leaderboard    `json:"cachedLeaderboard,omitempty"` // Cached leaderboard data for the group (optional)
	BannerURL         *string         `json:"bannerUrl,omitempty"`         // URL to the group's banner image (optional)
	PrivacyStatus     string          `json:"privacyStatus"`               // Privacy status of the group (e.g., "public", "private")
	ImportanceScore   float64         `json:"importanceScore"`             // Importance score of the group
}

// Leaderboard represents the cached leaderboard for a group.
type Leaderboard struct {
	TopTraders  []Trader `json:"topTraders"`  // List of top traders in the group
	TopCreators []Trader `json:"topCreators"` // List of top content creators in the group
}

// Trader represents a user and their score on the leaderboard.
type Trader struct {
	UserID string  `json:"userId"` // ID of the user
	Score  float64 `json:"score"`  // Score of the user on the leaderboard
}

// FullMarket represents a comprehensive view of a market, extending LiteMarket.
type FullMarket struct {
	LiteMarket

	Answers               *[]ApiAnswer `json:"answers,omitempty"`               // List of possible answers (optional)
	ShouldAnswersSumToOne *bool        `json:"shouldAnswersSumToOne,omitempty"` // Indicates if answers should sum to one (optional)
	AddAnswersMode        *string      `json:"addAnswersMode,omitempty"`        // Mode for adding answers ("ANYONE", "ONLY_CREATOR", "DISABLED") (optional)
	Options               *[]struct {
		Text  string `json:"text"`  // Text of the option
		Votes int    `json:"votes"` // Number of votes for the option
	} `json:"options,omitempty"` // List of options and their votes (optional)
	TotalBounty     *float64        `json:"totalBounty,omitempty"`   // Total bounty for the market (optional)
	BountyLeft      *float64        `json:"bountyLeft,omitempty"`    // Bounty left for the market (optional)
	Description     json.RawMessage `json:"description"`             // Detailed description of the market
	TextDescription string          `json:"textDescription"`         // Text-based description of the market
	CoverImageUrl   *string         `json:"coverImageUrl,omitempty"` // URL to the market's cover image (optional)
	GroupSlugs      *[]string       `json:"groupSlugs,omitempty"`    // List of group slugs associated with the market (optional)
}

// LiteMarket represents a basic view of a market with essential fields.
type LiteMarket struct {
	ID                    string             `json:"id"`                              // Unique identifier for the market
	CreatorID             string             `json:"creatorId"`                       // ID of the market creator
	CreatorUsername       string             `json:"creatorUsername"`                 // Username of the market creator
	CreatorName           string             `json:"creatorName"`                     // Full name of the market creator
	CreatedTime           int64              `json:"createdTime"`                     // Timestamp when the market was created
	CreatorAvatarURL      *string            `json:"creatorAvatarUrl,omitempty"`      // URL to the creator's avatar image (optional)
	CloseTime             *int64             `json:"closeTime,omitempty"`             // Timestamp when the market closes (optional)
	Question              string             `json:"question"`                        // Question posed by the market
	Slug                  string             `json:"slug"`                            // Slug for the market (usually URL-friendly)
	URL                   string             `json:"url"`                             // URL to the market
	OutcomeType           string             `json:"outcomeType"`                     // Type of outcome for the market (e.g., "BINARY", "MULTIPLE")
	Mechanism             string             `json:"mechanism"`                       // Mechanism used in the market (e.g., "CPMM")
	Pool                  map[string]float64 `json:"pool,omitempty"`                  // Pool of funds in the market (optional)
	Probability           *float64           `json:"probability,omitempty"`           // Current probability for the market (optional)
	P                     *float64           `json:"p,omitempty"`                     // Additional probability field (optional)
	TotalLiquidity        *float64           `json:"totalLiquidity,omitempty"`        // Total liquidity in the market (optional)
	Value                 *float64           `json:"value,omitempty"`                 // Value of the market (optional)
	Min                   *float64           `json:"min,omitempty"`                   // Minimum value for the market (optional)
	Max                   *float64           `json:"max,omitempty"`                   // Maximum value for the market (optional)
	Volume                float64            `json:"volume"`                          // Total volume of the market
	Volume24Hours         float64            `json:"volume24Hours"`                   // Volume in the last 24 hours
	IsResolved            bool               `json:"isResolved"`                      // Indicates if the market is resolved
	Resolution            *string            `json:"resolution,omitempty"`            // Resolution of the market (optional)
	ResolutionTime        *int64             `json:"resolutionTime,omitempty"`        // Timestamp when the market was resolved (optional)
	ResolutionProbability *float64           `json:"resolutionProbability,omitempty"` // Probability at the time of resolution (optional)
	UniqueBettorCount     int                `json:"uniqueBettorCount"`               // Number of unique bettors in the market
	LastUpdatedTime       *int64             `json:"lastUpdatedTime,omitempty"`       // Timestamp when the market was last updated (optional)
	LastBetTime           *int64             `json:"lastBetTime,omitempty"`           // Timestamp of the last bet (optional)
	MarketTier            *string            `json:"marketTier,omitempty"`            // Tier of the market (optional)
}

// Answer represents a possible answer in a market.
type Answer struct {
	ID                    string   `json:"id"`                              // Unique identifier for the answer
	Index                 int      `json:"index"`                           // Index of the answer in the list
	ContractID            string   `json:"contractId"`                      // ID of the associated contract
	UserID                string   `json:"userId"`                          // ID of the user who created the answer
	Text                  string   `json:"text"`                            // Text of the answer
	CreatedTime           int64    `json:"createdTime"`                     // Timestamp when the answer was created
	Color                 *string  `json:"color,omitempty"`                 // Color associated with the answer (optional)
	PoolYes               float64  `json:"poolYes"`                         // Pool of "yes" votes
	PoolNo                float64  `json:"poolNo"`                          // Pool of "no" votes
	Prob                  float64  `json:"prob"`                            // Current probability of the answer
	TotalLiquidity        float64  `json:"totalLiquidity"`                  // Total liquidity for the answer
	SubsidyPool           float64  `json:"subsidyPool"`                     // Subsidy pool for the answer
	IsOther               *bool    `json:"isOther,omitempty"`               // Indicates if this is an "other" answer (optional)
	Resolution            *string  `json:"resolution,omitempty"`            // Resolution of the answer (optional)
	ResolutionTime        *int64   `json:"resolutionTime,omitempty"`        // Timestamp when the answer was resolved (optional)
	ResolutionProbability *float64 `json:"resolutionProbability,omitempty"` // Probability at the time of resolution (optional)
	ResolverID            *string  `json:"resolverId,omitempty"`            // ID of the user who resolved the answer (optional)
	ProbChanges           struct {
		Day   float64 `json:"day"`   // Probability change over the last day
		Week  float64 `json:"week"`  // Probability change over the last week
		Month float64 `json:"month"` // Probability change over the last month
	} `json:"probChanges"` // Changes in probability over different time periods
	LoverUserID *string `json:"loverUserId,omitempty"` // ID of a user associated with the answer (optional)
}

// ApiAnswer represents the API's version of an answer, with adjusted fields.
type ApiAnswer struct {
	Answer

	Probability float64            `json:"probability"` // Current probability of the answer
	Pool        map[string]float64 `json:"pool"`        // Pool of funds associated with the answer
}

// ContractMetric represents a metric related to a contract.
type ContractMetric struct {
	ID               int                      `json:"id"`                         // Unique identifier for the metric
	ContractID       string                   `json:"contractId"`                 // ID of the associated contract
	From             map[string]PeriodMetrics `json:"from,omitempty"`             // Profit and investment metrics from specific periods (optional)
	HasNoShares      bool                     `json:"hasNoShares"`                // Indicates if the user holds no shares
	HasShares        bool                     `json:"hasShares"`                  // Indicates if the user holds shares
	HasYesShares     bool                     `json:"hasYesShares"`               // Indicates if the user holds "yes" shares
	Invested         float64                  `json:"invested"`                   // Amount invested by the user
	Loan             float64                  `json:"loan"`                       // Loan amount associated with the contract
	MaxSharesOutcome *string                  `json:"maxSharesOutcome,omitempty"` // Outcome with the maximum shares (optional)
	Payout           float64                  `json:"payout"`                     // Payout amount for the contract
	Profit           float64                  `json:"profit"`                     // Profit made from the contract
	ProfitPercent    float64                  `json:"profitPercent"`              // Profit percentage from the contract
	TotalShares      map[string]float64       `json:"totalShares"`                // Total shares held by the user
	UserID           string                   `json:"userId"`                     // ID of the user associated with the metric
	UserUsername     string                   `json:"userUsername"`               // Username of the user
	UserName         string                   `json:"userName"`                   // Full name of the user
	UserAvatarURL    string                   `json:"userAvatarUrl"`              // URL to the user's avatar image
	LastBetTime      int64                    `json:"lastBetTime"`                // Timestamp of the last bet made by the user
	AnswerID         *string                  `json:"answerId,omitempty"`         // ID of the associated answer (optional)
	ProfitAdjustment *float64                 `json:"profitAdjustment,omitempty"` // Profit adjustment for the contract (optional)
}

// PeriodMetrics represents the profit and investment metrics for a specific period.
type PeriodMetrics struct {
	Profit        float64 `json:"profit"`        // Profit during the period
	ProfitPercent float64 `json:"profitPercent"` // Profit percentage during the period
	Invested      float64 `json:"invested"`      // Amount invested during the period
	PrevValue     float64 `json:"prevValue"`     // Previous value before the period
	Value         float64 `json:"value"`         // Current value after the period
}

// Fees represents the fees associated with a bet or a fill.
type Fees struct {
	CreatorFee   float64 `json:"creatorFee"`   // Fee taken by the creator
	PlatformFee  float64 `json:"platformFee"`  // Fee taken by the platform
	LiquidityFee float64 `json:"liquidityFee"` // Fee taken for liquidity
}

// Fill represents a record of a transaction that partially or fully fills a bet order.
type Fill struct {
	MatchedBetID string  `json:"matchedBetId"`     // ID of the matched bet (null if matched by pool)
	Amount       float64 `json:"amount"`           // Amount of the fill
	Shares       float64 `json:"shares"`           // Shares bought/sold in the fill
	Timestamp    int64   `json:"timestamp"`        // Timestamp when the fill occurred
	Fees         Fees    `json:"fees"`             // Fees associated with the fill
	IsSale       *bool   `json:"isSale,omitempty"` // Indicates if this was a sale (optional)
}

// LimitProps represents properties specific to limit orders.
type LimitProps struct {
	OrderAmount float64 `json:"orderAmount"`         // Amount of the order
	LimitProb   float64 `json:"limitProb"`           // Probability limit for the order
	IsFilled    bool    `json:"isFilled"`            // Indicates if the order is filled
	IsCancelled bool    `json:"isCancelled"`         // Indicates if the order is cancelled
	Fills       []Fill  `json:"fills"`               // List of fills associated with the order
	ExpiresAt   *int64  `json:"expiresAt,omitempty"` // Expiration time of the order (optional)
}

// Bet represents a bet placed in a contract.
type Bet struct {
	ID               string      `json:"id"`                         // Unique identifier for the bet
	UserID           string      `json:"userId"`                     // ID of the user who placed the bet
	ContractID       string      `json:"contractId"`                 // ID of the associated contract
	AnswerID         *string     `json:"answerId,omitempty"`         // ID of the associated answer for multi-binary contracts (optional)
	CreatedTime      int64       `json:"createdTime"`                // Timestamp when the bet was placed
	UpdatedTime      *int64      `json:"updatedTime,omitempty"`      // Timestamp when the bet was last updated (optional)
	Amount           float64     `json:"amount"`                     // Amount of the bet
	LoanAmount       *float64    `json:"loanAmount,omitempty"`       // Loan amount associated with the bet (optional)
	Outcome          string      `json:"outcome"`                    // Outcome chosen for the bet
	Shares           float64     `json:"shares"`                     // Number of shares bought/sold
	ProbBefore       float64     `json:"probBefore"`                 // Probability before the bet
	ProbAfter        float64     `json:"probAfter"`                  // Probability after the bet
	Fees             Fees        `json:"fees"`                       // Fees associated with the bet
	IsApi            *bool       `json:"isApi,omitempty"`            // Indicates if the bet was placed via API (optional)
	IsRedemption     bool        `json:"isRedemption"`               // Indicates if the bet is a redemption
	ReplyToCommentID *string     `json:"replyToCommentId,omitempty"` // ID of the comment the bet replies to (optional)
	BetGroupID       *string     `json:"betGroupId,omitempty"`       // ID of the group associated with the bet (optional)
	LimitProps       *LimitProps `json:"limitProps,omitempty"`       // Limit order properties (optional)
}

// AnyTxnType represents the generic type of transaction.
type AnyTxnType struct {
	Category string `json:"category"` // Category of the transaction
}

// Txn represents a transaction within the system. This can involve different entities,
// such as users or contracts, and may include various types of financial operations.
type Txn struct {
	ID          string                 `json:"id"`                    // Unique identifier for the transaction
	CreatedTime int64                  `json:"createdTime"`           // Timestamp when the transaction was created
	FromID      string                 `json:"fromId"`                // ID of the entity sending the transaction
	FromType    string                 `json:"fromType"`              // Type of the sender (e.g., "user", "contract")
	ToID        string                 `json:"toId"`                  // ID of the entity receiving the transaction
	ToType      string                 `json:"toType"`                // Type of the recipient (e.g., "user", "contract")
	Amount      float64                `json:"amount"`                // Amount of the transaction
	Token       string                 `json:"token"`                 // Token or currency type used in the transaction
	Category    string                 `json:"category"`              // Category of the transaction, derived from AnyTxnType
	Description *string                `json:"description,omitempty"` // Optional description of the transaction
	Data        map[string]interface{} `json:"data,omitempty"`        // Extra data related to the transaction, if any
	AnyTxnType                         // Embedding AnyTxnType to include its fields
}

// Resolution represents the outcome of a resolution process, typically associated
// with a market or contract. It details the answer selected and the percentage
// allocated to it.
type Resolution struct {
	Answer int `json:"answer"` // Index of the selected answer in the resolution
	Pct    int `json:"pct"`    // Percentage allocated to the answer
}

// Comment represents a user comment associated with a bet, contract, or other entities.
// Comments may include replies, visibility settings, and status indicators (e.g., pinned, hidden).
type Comment struct {
	ID               string          `json:"id"`                         // Unique identifier for the comment
	ReplyToCommentID *string         `json:"replyToCommentId,omitempty"` // Optional ID of the comment being replied to
	UserID           string          `json:"userId"`                     // ID of the user who made the comment
	Text             *string         `json:"text,omitempty"`             // Deprecated: Use Content instead
	Content          json.RawMessage `json:"content"`                    // Content of the comment, potentially rich text or JSON
	CreatedTime      int64           `json:"createdTime"`                // Timestamp when the comment was created
	UserName         string          `json:"userName"`                   // Full name of the user who made the comment
	UserUsername     string          `json:"userUsername"`               // Username of the user who made the comment
	UserAvatarURL    *string         `json:"userAvatarUrl,omitempty"`    // Optional URL to the user's avatar
	Likes            *int            `json:"likes,omitempty"`            // Deprecated: Number of likes (still used, but no longer maintained)
	Hidden           *bool           `json:"hidden,omitempty"`           // Optional flag indicating if the comment is hidden
	HiddenTime       *int64          `json:"hiddenTime,omitempty"`       // Optional timestamp when the comment was hidden
	HiderID          *string         `json:"hiderId,omitempty"`          // Optional ID of the user who hid the comment
	Pinned           *bool           `json:"pinned,omitempty"`           // Optional flag indicating if the comment is pinned
	PinnedTime       *int64          `json:"pinnedTime,omitempty"`       // Optional timestamp when the comment was pinned
	PinnerID         *string         `json:"pinnerId,omitempty"`         // Optional ID of the user who pinned the comment
	Visibility       string          `json:"visibility"`                 // Visibility status of the comment (e.g., "public", "private")
	EditedTime       *int64          `json:"editedTime,omitempty"`       // Optional timestamp when the comment was last edited
	IsApi            *bool           `json:"isApi,omitempty"`            // Optional flag indicating if the comment was posted via API
}
