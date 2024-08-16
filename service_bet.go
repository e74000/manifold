package manifold

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// BetService provides methods for interacting with bets, including retrieving bets, creating new bets, and canceling existing bets.
type BetService struct {
	client *Client
}

// Bets retrieves a list of bets based on various filtering criteria.
//
// Parameters:
//   - userID: Filter bets by the ID of the user who placed them. Optional.
//   - username: Filter bets by the username of the user who placed them. Optional.
//   - contractID: Filter bets by the ID of the contract. Optional.
//   - contractSlug: Filter bets by the slug of the contract. Optional.
//   - limit: Limits the number of results returned. Must be between 0 and 1000. Optional.
//   - before: Only return bets placed before this cursor (e.g., a timestamp or ID). Optional.
//   - after: Only return bets placed after this cursor (e.g., a timestamp or ID). Optional.
//   - beforeTime: Only return bets placed before this timestamp. Optional.
//   - afterTime: Only return bets placed after this timestamp. Optional.
//   - kinds: Filter bets by their kind (e.g., "open-limit"). Optional.
//   - order: Sort results in "asc" or "desc" order based on placement time. Optional.
//
// Returns:
//   - []Bet: A slice of bets matching the specified criteria.
//   - error: An error object if the request fails or if input validation fails.
func (s *BetService) Bets(userID *string, username *string, contractID *string, contractSlug *string, limit *int, before *string, after *string, beforeTime *time.Time, afterTime *time.Time, kinds *string, order *string) ([]Bet, error) {
	params := make(map[string]string, 11)

	if userID != nil {
		params["userId"] = *userID
	}

	if username != nil {
		params["username"] = *username
	}

	if contractID != nil {
		params["contractId"] = *contractID
	}

	if contractSlug != nil {
		params["contractSlug"] = *contractSlug
	}

	if limit != nil {
		if err := checkInRange(*limit, 0, 1000); err != nil {
			return nil, fmt.Errorf("Bet: Bets(limit): %w", err)
		}

		params["limit"] = fmt.Sprintf("%d", *limit)
	}

	if before != nil {
		params["before"] = *before
	}

	if after != nil {
		params["after"] = *after
	}

	if beforeTime != nil {
		params["beforeTime"] = fmt.Sprintf("%d", beforeTime.UnixMilli())
	}

	if afterTime != nil {
		params["afterTime"] = fmt.Sprintf("%d", afterTime.UnixMilli())
	}

	if kinds != nil {
		if err := checkOneOf(*kinds, "open-limit"); err != nil {
			return nil, fmt.Errorf("Bet: Bets(kinds): %w", err)
		}

		params["kinds"] = *kinds
	}

	if order != nil {
		if err := checkOneOf(*order, "asc", "desc"); err != nil {
			return nil, fmt.Errorf("Bet: Bets(order): %w", err)
		}

		params["order"] = *order
	}

	result, err := s.client.GET("/bets", params)
	if err != nil {
		return nil, fmt.Errorf("Bet: Bets: %w: %w", ErrorGETFailed, err)
	}

	bets := make([]Bet, 0)
	err = json.Unmarshal(result, &bets)
	if err != nil {
		return nil, fmt.Errorf("Bet: Bets: %w: %w", ErrorFailedToParseResponse, err)
	}

	return bets, nil
}

// Create places a new bet on a contract.
//
// Parameters:
//   - amount: The amount of the bet. Required.
//   - contractID: The ID of the contract on which the bet is being placed. Required.
//   - outcome: The outcome of the bet (e.g., "YES" or "NO"). Optional.
//   - limitProb: Probability threshold for a limit order. Must be between 0 and 1. Optional.
//   - expiresAt: Expiration time for a limit order. Only valid if limitProb is set. Optional.
//   - dryRun: If true, simulates the bet without placing it. Optional.
//
// Returns:
//   - *Bet: The created bet object.
//   - error: An error object if the request fails, input validation fails, or the response cannot be parsed.
func (s *BetService) Create(amount float64, contractID string, outcome *string, limitProb *float64, expiresAt *time.Time, dryRun *bool) (*Bet, error) {
	body := map[string]string{
		"amount":     fmt.Sprintf("%f", amount),
		"contractId": contractID,
	}

	if outcome != nil {
		if err := checkOneOf(*outcome, "YES", "NO"); err != nil {
			return nil, fmt.Errorf("Bet: Create(outcome): %w", err)
		}

		body["outcome"] = *outcome
	}

	if limitProb != nil {
		if err := checkInRange(*limitProb, 0, 1); err != nil {
			return nil, fmt.Errorf("Bet: Create(limitProb): %w", err)
		}

		body["limitProb"] = fmt.Sprintf("%f", *limitProb)
	}

	if expiresAt != nil {
		if limitProb == nil {
			return nil, fmt.Errorf("Bet: Create(expiresAt): only limit orders can have an expiresAt")
		}

		if time.Now().After(*expiresAt) {
			return nil, fmt.Errorf("Bet: Create(expiresAt): limit order cannot expire in the past")
		}

		body["expiresAt"] = fmt.Sprintf("%d", expiresAt.UnixMilli())
	}

	if dryRun != nil {
		if *dryRun {
			body["dryRun"] = "true"
		} else {
			body["dryRun"] = "false"
		}
	}

	result, err := s.client.POST("/bet", body)
	if err != nil {
		return nil, fmt.Errorf("Bet: Create: %w: %w", ErrorPOSTFailed, err)
	}

	bet := new(Bet)
	err = json.Unmarshal(result, bet)
	if err != nil {
		return nil, fmt.Errorf("Bet: Create: %w: %w", ErrorFailedToParseResponse, err)
	}

	return bet, nil
}

// Cancel cancels an existing bet.
//
// Parameters:
//   - id: The ID of the bet to cancel. Required.
//
// Returns:
//   - error: An error object if the request fails or if the response cannot be parsed.
func (s *BetService) Cancel(id string) error {
	_, err := s.client.POST(
		fmt.Sprintf("/bet/cancel/%s", url.PathEscape(id)), nil,
	)
	if err != nil {
		return fmt.Errorf("Bet: Cancel: %w: %w", ErrorPOSTFailed, err)
	}

	return nil
}
