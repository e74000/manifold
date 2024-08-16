package manifold

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

var (
	allowedMarketMarketsSort = []string{
		"created-time", "updated-time", "last-bet-time", "last-comment-time",
	}

	allowedMarketSearchSort = []string{
		"newest", "score", "daily-score", "freshness-score", "24-hour-vol", "most-popular", "liquidity", "subsidy",
		"last-updated", "close-date", "resolve-date", "random", "bounty-amount", "prob-descending", "prob-ascending",
	}

	allowedMarketSearchFilter = []string{
		"all", "open", "closed", "resolved", "closing-this-month", "closing-next-month",
	}

	allowedMarketSearchContractType = []string{
		"ALL", "BINARY", "MULTIPLE_CHOICE", "FREE-RESPONSE", "PSEUDO-NUMERIC", "BOUNTIED_QUESTION", "STONK", "POLL",
		"NUMBER",
	}
)

// MarketService provides methods for managing markets, including retrieving market data,
// creating new markets, adding liquidity or bounties, resolving markets, and more.
type MarketService struct {
	client *Client
}

// Markets retrieves a list of markets based on various filtering criteria.
//
// Parameters:
//   - limit: Limits the number of results returned. Must be between 0 and 1000. Optional.
//   - sort: Sorts the results based on one of the allowed sorting options (e.g., "created-time", "updated-time"). Optional.
//   - order: Specifies the order of the results, either "asc" or "desc". Optional.
//   - before: Retrieves markets created before this cursor. Optional.
//   - userID: Filters markets created by a specific user ID. Optional.
//   - groupID: Filters markets associated with a specific group ID. Optional.
//
// Returns:
//   - []LiteMarket: A slice of markets matching the specified criteria.
//   - error: An error object if the request fails or if input validation fails.
func (s *MarketService) Markets(limit *int, sort *string, order *string, before *string, userID *string, groupID *string) ([]LiteMarket, error) {
	params := make(map[string]string, 6)

	if limit != nil {
		if err := checkInRange(*limit, 0, 1000); err != nil {
			return nil, fmt.Errorf("Market: Markets(limit): %w", err)
		}

		params["limit"] = fmt.Sprintf("%d", *limit)
	}

	if sort != nil {
		if err := checkOneOf(*sort, allowedMarketMarketsSort...); err != nil {
			return nil, fmt.Errorf("Market: Markets(sort): %w", err)
		}

		params["sort"] = *sort
	}

	if order != nil {
		if err := checkOneOf(*order, "asc", "desc"); err != nil {
			return nil, fmt.Errorf("Market: Markets(order): %w", err)
		}

		params["order"] = *order
	}

	if before != nil {
		params["before"] = *before
	}

	if userID != nil {
		params["userID"] = *userID
	}

	if groupID != nil {
		params["groupID"] = *groupID
	}

	result, err := s.client.GET(
		"/markets", params,
	)
	if err != nil {
		return nil, fmt.Errorf("Market: Markets: %w: %w", ErrorGETFailed, err)
	}

	markets := make([]LiteMarket, 0)
	err = json.Unmarshal(result, &markets)
	if err != nil {
		return nil, fmt.Errorf("Market: Markets: %w: %w", ErrorFailedToParseResponse, err)
	}

	return markets, nil
}

// Market retrieves the details of a specific market using its ID.
//
// Parameters:
//   - id: The ID of the market to retrieve. Required.
//
// Returns:
//   - *FullMarket: A pointer to the retrieved market object.
//   - error: An error object if the request fails or if the response cannot be parsed.
func (s *MarketService) Market(id string) (*FullMarket, error) {
	result, err := s.client.GET(
		fmt.Sprintf("/market/%s", url.PathEscape(id)), nil,
	)
	if err != nil {
		return nil, fmt.Errorf("Market: Market: %w: %w", ErrorGETFailed, err)
	}

	market := new(FullMarket)
	err = json.Unmarshal(result, market)
	if err != nil {
		return nil, fmt.Errorf("Market: Market: %w: %w", ErrorFailedToParseResponse, err)
	}

	return market, nil
}

// Positions retrieves the positions for a specific market using its ID.
//
// Parameters:
//   - id: The ID of the market to retrieve positions for. Required.
//
// Returns:
//   - []ContractMetric: A slice of contract metrics representing the positions.
//   - error: An error object if the request fails or if the response cannot be parsed.
func (s *MarketService) Positions(id string) ([]ContractMetric, error) {
	result, err := s.client.GET(
		fmt.Sprintf("/market/%s/positions", url.PathEscape(id)), nil,
	)
	if err != nil {
		return nil, fmt.Errorf("Market: Positions: %w: %w", ErrorGETFailed, err)
	}

	positions := make([]ContractMetric, 0)
	err = json.Unmarshal(result, &positions)
	if err != nil {
		return nil, fmt.Errorf("Market: Positions: %w, %w", ErrorFailedToParseResponse, err)
	}

	return positions, nil
}

// Slug retrieves the details of a market using its slug.
//
// Parameters:
//   - slug: The slug of the market to retrieve. Required.
//
// Returns:
//   - *FullMarket: A pointer to the retrieved market object.
//   - error: An error object if the request fails or if the response cannot be parsed.
func (s *MarketService) Slug(slug string) (*FullMarket, error) {
	result, err := s.client.GET(
		fmt.Sprintf("/slug/%s", url.PathEscape(slug)), nil,
	)
	if err != nil {
		return nil, fmt.Errorf("Market: Slug: %w: %w", ErrorGETFailed, err)
	}

	market := new(FullMarket)
	err = json.Unmarshal(result, market)
	if err != nil {
		return nil, fmt.Errorf("Market: Slug: %w: %w", ErrorFailedToParseResponse, err)
	}

	return market, nil
}

// Search searches for markets based on various criteria.
//
// Parameters:
//   - term: The search term. Required.
//   - sort: Sorts the results based on one of the allowed sorting options (e.g., "newest", "score"). Optional.
//   - filter: Filters results based on their state (e.g., "open", "closed"). Optional.
//   - contractType: Filters results based on the type of contract (e.g., "BINARY", "POLL"). Optional.
//   - topicSlug: Filters results based on a topic slug. Optional.
//   - creatorID: Filters results based on the creator's user ID. Optional.
//   - limit: Limits the number of results returned. Must be between 0 and 1000. Optional.
//   - offset: Skips the specified number of results before returning. Must be 0 or greater. Optional.
//
// Returns:
//   - []LiteMarket: A slice of markets matching the specified criteria.
//   - error: An error object if the request fails or if input validation fails.
func (s *MarketService) Search(term string, sort *string, filter *string, contractType *string, topicSlug *string, creatorID *string, limit *int, offset *int) ([]LiteMarket, error) {
	params := make(map[string]string, 8)
	params[term] = term

	if sort != nil {
		if err := checkOneOf(*sort, allowedMarketSearchSort...); err != nil {
			return nil, fmt.Errorf("Market: Search(sort): %w", err)
		}

		params["sort"] = *sort
	}

	if filter != nil {
		if err := checkOneOf(*filter, allowedMarketSearchFilter...); err != nil {
			return nil, fmt.Errorf("Market: Search(filter): %w", err)
		}

		params["filter"] = *filter
	}

	if contractType != nil {
		if err := checkOneOf(*contractType, allowedMarketSearchContractType...); err != nil {
			return nil, fmt.Errorf("Market: Search(contractType): %w", err)
		}

		params["contractType"] = *contractType
	}

	if topicSlug != nil {
		params["topicSlug"] = *topicSlug
	}

	if creatorID != nil {
		params["creatorId"] = *creatorID
	}

	if limit != nil {
		if err := checkInRange(*limit, 0, 1000); err != nil {
			return nil, fmt.Errorf("Market: Search(limit): %w", err)
		}

		params["limit"] = fmt.Sprintf("%d", *limit)
	}

	if offset != nil {
		if *offset < 0 {
			return nil, fmt.Errorf("Market: Search(offset): invalid value: %v, must be greater than 0", *offset)
		}

		params["offset"] = fmt.Sprintf("%d", *offset)
	}

	result, err := s.client.GET(
		"/search-markets", params,
	)
	if err != nil {
		return nil, fmt.Errorf("Market: Search: %w: %w", ErrorGETFailed, err)
	}

	markets := make([]LiteMarket, 0)
	err = json.Unmarshal(result, &markets)
	if err != nil {
		return nil, fmt.Errorf("Market: Search: %w: %w", ErrorFailedToParseResponse, err)
	}

	return markets, nil
}

// Helper method to create a market.
func (s *MarketService) createMarket(params map[string]interface{}) (*LiteMarket, error) {
	result, err := s.client.POST("/market", params)
	if err != nil {
		return nil, fmt.Errorf("Market: createMarket: %w", err)
	}

	market := new(LiteMarket)
	err = json.Unmarshal(result, market)
	if err != nil {
		return nil, fmt.Errorf("Market: createMarket: %w", err)
	}

	return market, nil
}

// CreateBinary creates a binary market.
//
// Parameters:
//   - question: The question the market is based on. Required.
//   - initialProb: The initial probability (between 1 and 99) of the market outcome. Required.
//   - description: A description of the market. Optional.
//   - closeTime: The time when the market will close. Must be in the future. Optional.
//   - visibility: The visibility of the market ("public" or "unlisted"). Optional.
//   - extraLiquidity: The extra liquidity to add to the market. Optional.
//
// Returns:
//   - *LiteMarket: A pointer to the created market object.
//   - error: An error object if the request fails or if input validation fails.
func (s *MarketService) CreateBinary(question string, initialProb int, description *string, closeTime *time.Time, visibility *string, extraLiquidity *int) (*LiteMarket, error) {
	// Validate inputs
	if err := checkInRange(initialProb, 1, 99); err != nil {
		return nil, fmt.Errorf("Market: CreateBinary: %w", err)
	}

	params := map[string]interface{}{
		"outcomeType": "BINARY",
		"question":    question,
		"initialProb": initialProb,
	}

	if description != nil {
		params["description"] = *description
	}
	if closeTime != nil {
		if time.Now().After(*closeTime) {
			return nil, fmt.Errorf("Market: CreateBinary: closeTime cannot be in the past")
		}
		params["closeTime"] = closeTime.UnixMilli()
	}
	if visibility != nil {
		if err := checkOneOf(*visibility, "public", "unlisted"); err != nil {
			return nil, fmt.Errorf("Market: CreateBinary: %w", err)
		}
		params["visibility"] = *visibility
	}
	if extraLiquidity != nil {
		params["extraLiquidity"] = *extraLiquidity
	}

	return s.createMarket(params)
}

// CreatePseudoNumeric creates a pseudo-numeric market.
//
// Parameters:
//   - question: The question the market is based on. Required.
//   - min: The minimum value for the market. Required.
//   - max: The maximum value for the market. Required.
//   - initialValue: The initial value for the market, between min and max. Required.
//   - isLogScale: Whether the market uses a logarithmic scale. Required.
//   - description: A description of the market. Optional.
//   - closeTime: The time when the market will close. Must be in the future. Optional.
//   - visibility: The visibility of the market ("public" or "unlisted"). Optional.
//   - extraLiquidity: The extra liquidity to add to the market. Optional.
//
// Returns:
//   - *LiteMarket: A pointer to the created market object.
//   - error: An error object if the request fails or if input validation fails.
func (s *MarketService) CreatePseudoNumeric(question string, min, max, initialValue int, isLogScale bool, description *string, closeTime *time.Time, visibility *string, extraLiquidity *int) (*LiteMarket, error) {
	// Validate inputs
	if err := checkInRange(initialValue, min+1, max-1); err != nil {
		return nil, fmt.Errorf("Market: CreatePseudoNumeric: %w", err)
	}

	params := map[string]interface{}{
		"outcomeType":  "PSEUDO_NUMERIC",
		"question":     question,
		"min":          min,
		"max":          max,
		"initialValue": initialValue,
		"isLogScale":   isLogScale,
	}

	if description != nil {
		params["description"] = *description
	}
	if closeTime != nil {
		if time.Now().After(*closeTime) {
			return nil, fmt.Errorf("Market: CreatePseudoNumeric: closeTime cannot be in the past")
		}
		params["closeTime"] = closeTime.UnixMilli()
	}
	if visibility != nil {
		if err := checkOneOf(*visibility, "public", "unlisted"); err != nil {
			return nil, fmt.Errorf("Market: CreatePseudoNumeric: %w", err)
		}
		params["visibility"] = *visibility
	}
	if extraLiquidity != nil {
		params["extraLiquidity"] = *extraLiquidity
	}

	return s.createMarket(params)
}

// CreatePoll creates a poll market.
//
// Parameters:
//   - question: The question the poll is based on. Required.
//   - answers: The possible answers for the poll. Must include at least two answers. Required.
//   - description: A description of the poll. Optional.
//   - closeTime: The time when the poll will close. Must be in the future. Optional.
//   - visibility: The visibility of the poll ("public" or "unlisted"). Optional.
//
// Returns:
//   - *LiteMarket: A pointer to the created poll object.
//   - error: An error object if the request fails or if input validation fails.
func (s *MarketService) CreatePoll(question string, answers []string, description *string, closeTime *time.Time, visibility *string) (*LiteMarket, error) {
	// Validate inputs
	if len(answers) < 2 {
		return nil, fmt.Errorf("Market: CreatePoll: at least two answers are required")
	}

	params := map[string]interface{}{
		"outcomeType": "POLL",
		"question":    question,
		"answers":     answers,
	}

	if description != nil {
		params["description"] = *description
	}
	if closeTime != nil {
		if time.Now().After(*closeTime) {
			return nil, fmt.Errorf("Market: CreatePoll: closeTime cannot be in the past")
		}
		params["closeTime"] = closeTime.UnixMilli()
	}
	if visibility != nil {
		if err := checkOneOf(*visibility, "public", "unlisted"); err != nil {
			return nil, fmt.Errorf("Market: CreatePoll: %w", err)
		}
		params["visibility"] = *visibility
	}

	return s.createMarket(params)
}

// CreateBountiedQuestion creates a bountied question market.
//
// Parameters:
//   - question: The question the market is based on. Required.
//   - totalBounty: The total bounty amount for the question. Must be greater than zero. Required.
//   - description: A description of the market. Optional.
//   - closeTime: The time when the market will close. Must be in the future. Optional.
//   - visibility: The visibility of the market ("public" or "unlisted"). Optional.
//
// Returns:
//   - *LiteMarket: A pointer to the created market object.
//   - error: An error object if the request fails or if input validation fails.
func (s *MarketService) CreateBountiedQuestion(question string, totalBounty int, description *string, closeTime *time.Time, visibility *string) (*LiteMarket, error) {
	// Validate inputs
	if totalBounty <= 0 {
		return nil, fmt.Errorf("Market: CreateBountiedQuestion: totalBounty must be greater than zero")
	}

	params := map[string]interface{}{
		"outcomeType": "BOUNTIED_QUESTION",
		"question":    question,
		"totalBounty": totalBounty,
	}

	if description != nil {
		params["description"] = *description
	}
	if closeTime != nil {
		if time.Now().After(*closeTime) {
			return nil, fmt.Errorf("Market: CreateBountiedQuestion: closeTime cannot be in the past")
		}
		params["closeTime"] = closeTime.UnixMilli()
	}
	if visibility != nil {
		if err := checkOneOf(*visibility, "public", "unlisted"); err != nil {
			return nil, fmt.Errorf("Market: CreateBountiedQuestion: %w", err)
		}
		params["visibility"] = *visibility
	}

	return s.createMarket(params)
}

// Answer submits an answer to a market.
//
// Parameters:
//   - id: The ID of the market to answer. Required.
//   - text: The answer text to submit. Required.
//
// Returns:
//   - error: An error object if the request fails or if the response cannot be parsed.
func (s *MarketService) Answer(id string, text string) error {
	body := map[string]string{
		"text": text,
	}

	_, err := s.client.POST(
		fmt.Sprintf("/market/%s/answer", url.PathEscape(id)), body,
	)
	if err != nil {
		return fmt.Errorf("Market: Answer: %w: %w", ErrorPOSTFailed, err)
	}

	return nil
}

// AddLiquidity adds liquidity to a market.
//
// Parameters:
//   - id: The ID of the market to add liquidity to. Required.
//   - amount: The amount of liquidity to add. Must be greater than zero. Required.
//
// Returns:
//   - *Txn: A pointer to the transaction object representing the added liquidity.
//   - error: An error object if the request fails or if input validation fails.
func (s *MarketService) AddLiquidity(id string, amount float64) (*Txn, error) {
	if amount < 0 {
		return nil, fmt.Errorf("Market: AddLiquidity(amount): invalid value: %f must be >0", amount)
	}

	body := map[string]string{
		"amount": fmt.Sprintf("%f", amount),
	}

	response, err := s.client.POST(
		fmt.Sprintf("/market/%s/add-liquidity", url.PathEscape(id)), body,
	)
	if err != nil {
		return nil, fmt.Errorf("Market: AddLiquidity: %w: %w", ErrorPOSTFailed, err)
	}

	txn := new(Txn)
	err = json.Unmarshal(response, txn)
	if err != nil {
		return nil, fmt.Errorf("Market: AddLiquidity: %w: %w", ErrorFailedToParseResponse, err)
	}

	return txn, nil
}

// AddBounty adds a bounty to a market.
//
// Parameters:
//   - id: The ID of the market to add the bounty to. Required.
//   - amount: The amount of the bounty. Must be greater than zero. Required.
//
// Returns:
//   - *Txn: A pointer to the transaction object representing the added bounty.
//   - error: An error object if the request fails or if input validation fails.
func (s *MarketService) AddBounty(id string, amount float64) (*Txn, error) {
	if amount < 0 {
		return nil, fmt.Errorf("Market: AddBounty(amount): invalid value: %f must be >0", amount)
	}

	body := map[string]string{
		"amount": fmt.Sprintf("%f", amount),
	}

	response, err := s.client.POST(
		fmt.Sprintf("/market/%s/add-bounty", url.PathEscape(id)), body,
	)
	if err != nil {
		return nil, fmt.Errorf("Market: AddBounty: %w: %w", ErrorPOSTFailed, err)
	}

	txn := new(Txn)
	err = json.Unmarshal(response, txn)
	if err != nil {
		return nil, fmt.Errorf("Market: AddBounty: %w: %w", ErrorFailedToParseResponse, err)
	}

	return txn, nil
}

// AwardBounty awards a bounty to a specific comment on a market.
//
// Parameters:
//   - id: The ID of the market to award the bounty for. Required.
//   - amount: The amount of the bounty to award. Must be greater than zero. Required.
//   - commentID: The ID of the comment to award the bounty to. Required.
//
// Returns:
//   - *Txn: A pointer to the transaction object representing the awarded bounty.
//   - error: An error object if the request fails or if input validation fails.
func (s *MarketService) AwardBounty(id string, amount float64, commentID string) (*Txn, error) {
	if amount < 0 {
		return nil, fmt.Errorf("Market: AwardBounty(amount): invalid value: %f must be >0", amount)
	}

	body := map[string]string{
		"amount":    fmt.Sprintf("%f", amount),
		"commentId": commentID,
	}

	response, err := s.client.POST(
		fmt.Sprintf("/market/%s/award-bounty", url.PathEscape(id)), body,
	)
	if err != nil {
		return nil, fmt.Errorf("Market: AwardBounty: %w: %w", ErrorPOSTFailed, err)
	}

	txn := new(Txn)
	err = json.Unmarshal(response, txn)
	if err != nil {
		return nil, fmt.Errorf("Market: AwardBounty: %w: %w", ErrorFailedToParseResponse, err)
	}

	return txn, nil
}

// Close closes a market, setting a specific close time if provided.
//
// Parameters:
//   - id: The ID of the market to close. Required.
//   - closeTime: The time when the market should close. Must be in the future if provided. Optional.
//
// Returns:
//   - error: An error object if the request fails or if input validation fails.
func (s *MarketService) Close(id string, closeTime *time.Time) error {
	body := map[string]string{}

	if closeTime != nil {
		if time.Now().After(*closeTime) {
			return fmt.Errorf("Market: Close(closeTime): cannot close a market in the past")
		}

		body["closeTime"] = fmt.Sprintf("%d", closeTime.UnixMilli())
	}

	_, err := s.client.POST(
		fmt.Sprintf("/market/%s/close", url.PathEscape(id)), body,
	)
	if err != nil {
		return fmt.Errorf("Market: Close: %w: %w", ErrorPOSTFailed, err)
	}

	return nil
}

// Group adds or removes a market from a group.
//
// Parameters:
//   - id: The ID of the market to group. Required.
//   - groupID: The ID of the group to add or remove the market from. Required.
//   - remove: If true, removes the market from the group; otherwise, adds the market to the group. Optional.
//
// Returns:
//   - error: An error object if the request fails or if input validation fails.
func (s *MarketService) Group(id string, groupID string, remove *bool) error {
	body := map[string]string{
		"groupId": groupID,
	}

	if remove != nil {
		if *remove {
			body["remove"] = "true"
		} else {
			body["remove"] = "false"
		}
	}

	_, err := s.client.POST(
		fmt.Sprintf("/market/%s/group", url.PathEscape(id)), body,
	)
	if err != nil {
		return fmt.Errorf("Market: Group: %w: %w", ErrorPOSTFailed, err)
	}

	return nil
}

// Helper function to resolve a market.
func (s *MarketService) resolveMarket(id string, params map[string]interface{}) (*LiteMarket, error) {
	result, err := s.client.POST(fmt.Sprintf("/market/%s/resolve", url.PathEscape(id)), params)
	if err != nil {
		return nil, fmt.Errorf("Market: resolveMarket: %w", err)
	}

	market := new(LiteMarket)
	err = json.Unmarshal(result, market)
	if err != nil {
		return nil, fmt.Errorf("Market: resolveMarket: %w", err)
	}

	return market, nil
}

// ResolveBinary resolves a binary market.
//
// Parameters:
//   - id: The ID of the market to resolve. Required.
//   - outcome: The outcome of the market ("YES", "NO", "MKT", "CANCEL"). Required.
//   - probabilityInt: The probability integer (0-100) if the outcome is "MKT". Optional.
//
// Returns:
//   - *LiteMarket: A pointer to the resolved market object.
//   - error: An error object if the request fails or if input validation fails.
func (s *MarketService) ResolveBinary(id string, outcome string, probabilityInt *int) (*LiteMarket, error) {
	// Validate outcome
	if err := checkOneOf(outcome, "YES", "NO", "MKT", "CANCEL"); err != nil {
		return nil, fmt.Errorf("Market: ResolveBinary: %w", err)
	}

	// Validate probabilityInt if outcome is "MKT"
	if outcome == "MKT" && probabilityInt != nil {
		if err := checkInRange(*probabilityInt, 0, 100); err != nil {
			return nil, fmt.Errorf("Market: ResolveBinary: %w", err)
		}
	}

	// Prepare parameters
	params := map[string]interface{}{
		"outcome": outcome,
	}
	if probabilityInt != nil {
		params["probabilityInt"] = *probabilityInt
	}

	return s.resolveMarket(id, params)
}

// ResolveFreeResponse resolves a free response or multiple choice market.
//
// Parameters:
//   - id: The ID of the market to resolve. Required.
//   - outcome: The outcome of the market ("MKT", "CANCEL"). Required.
//   - resolutions: A slice of resolutions with percentages for each outcome if outcome is "MKT". Optional.
//
// Returns:
//   - *LiteMarket: A pointer to the resolved market object.
//   - error: An error object if the request fails or if input validation fails.
func (s *MarketService) ResolveFreeResponse(id string, outcome string, resolutions []Resolution) (*LiteMarket, error) {
	// Validate outcome
	if err := checkOneOf(outcome, "MKT", "CANCEL"); err != nil {
		if outcome == "MKT" && resolutions == nil {
			return nil, fmt.Errorf("Market: ResolveFreeResponse: outcome cannot be a specific answer without resolutions")
		}
	}

	// Prepare parameters
	params := map[string]interface{}{
		"outcome": outcome,
	}

	if outcome == "MKT" && len(resolutions) > 0 {
		totalPct := 0
		for _, resolution := range resolutions {
			totalPct += resolution.Pct
		}
		if totalPct != 100 {
			return nil, fmt.Errorf("Market: ResolveFreeResponse: total percentages of resolutions must add up to 100")
		}
		params["resolutions"] = resolutions
	}

	return s.resolveMarket(id, params)
}

// ResolveNumeric resolves a numeric market.
//
// Parameters:
//   - id: The ID of the market to resolve. Required.
//   - outcome: The outcome of the market ("CANCEL"). Required.
//   - value: The final value of the market. Optional.
//   - probabilityInt: The probability integer (0-100) if applicable. Optional.
//
// Returns:
//   - *LiteMarket: A pointer to the resolved market object.
//   - error: An error object if the request fails or if input validation fails.
func (s *MarketService) ResolveNumeric(id string, outcome string, value *float64, probabilityInt *int) (*LiteMarket, error) {
	// Validate outcome
	if err := checkOneOf(outcome, "CANCEL"); err != nil {
		return nil, fmt.Errorf("Market: ResolveNumeric: %w", err)
	}

	// Validate value and probabilityInt if provided
	if value != nil && probabilityInt != nil {
		if err := checkInRange(*probabilityInt, 0, 100); err != nil {
			return nil, fmt.Errorf("Market: ResolveNumeric: %w", err)
		}
	}

	// Prepare parameters
	params := map[string]interface{}{
		"outcome": outcome,
	}
	if value != nil {
		params["value"] = *value
	}
	if probabilityInt != nil {
		params["probabilityInt"] = *probabilityInt
	}

	return s.resolveMarket(id, params)
}

// Sell sells shares in a market.
//
// Parameters:
//   - id: The ID of the market to sell shares in. Required.
//   - outcome: The outcome to sell shares in ("YES" or "NO"). Optional.
//   - shares: The number of shares to sell. Must be greater than zero. Optional.
//   - answerID: The ID of the specific answer to sell shares in for multiple choice markets. Optional.
//
// Returns:
//   - *Bet: A pointer to the bet object representing the sale.
//   - error: An error object if the request fails or if input validation fails.
func (s *MarketService) Sell(id string, outcome *string, shares *float64, answerID *string) (*Bet, error) {
	body := map[string]string{}

	if outcome != nil {
		if err := checkOneOf(*outcome, "YES", "NO"); err != nil {
			return nil, fmt.Errorf("Market: Sell(outcome): %w", err)
		}

		body["outcome"] = *outcome
	}

	if shares != nil {
		if *shares <= 0 {
			return nil, fmt.Errorf("Market: Sell(shares): invalid value: %f, value must be >0", *shares)
		}

		body["shares"] = fmt.Sprintf("%f", *shares)
	}

	if answerID != nil {
		body["answerId"] = *answerID
	}

	result, err := s.client.POST(
		fmt.Sprintf("/market/%s/sell", url.PathEscape(id)), body,
	)
	if err != nil {
		return nil, fmt.Errorf("Market: Sell: %w: %w", ErrorPOSTFailed, err)
	}

	bet := new(Bet)
	err = json.Unmarshal(result, bet)
	if err != nil {
		return nil, fmt.Errorf("Market: Sell: %w: %w", ErrorFailedToParseResponse, err)
	}

	return bet, nil
}
