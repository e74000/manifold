package manifold

import (
	"encoding/json"
	"fmt"
	"time"
)

type ManaService struct {
	client *Client
}

// Managrams retrieves a list of Managram transactions based on optional filtering criteria.
//
// Parameters:
//   - toID: Filter transactions by the recipient's user ID. Optional.
//   - fromID: Filter transactions by the sender's user ID. Optional.
//   - limit: Limits the number of results returned. Must be between 0 and 1000. Optional.
//   - before: Only return transactions before this timestamp. Optional.
//   - after: Only return transactions after this timestamp. Optional.
//
// Returns:
//   - []Txn: A slice of transactions matching the specified criteria.
//   - error: An error object if the request fails or if the response cannot be parsed.
func (s *ManaService) Managrams(toID *string, fromID *string, limit *int, before *time.Time, after *time.Time) ([]Txn, error) {
	params := make(map[string]string, 5)

	if toID != nil {
		params["toId"] = *toID
	}

	if fromID != nil {
		params["fromId"] = *fromID
	}

	if limit != nil {
		if err := checkInRange(*limit, 0, 1000); err != nil {
			return nil, fmt.Errorf("Misc: Managrams(limit): %w", err)
		}

		params["limit"] = fmt.Sprintf("%d", *limit)
	}

	if before != nil {
		params["before"] = fmt.Sprintf("%d", before.UnixMilli())
	}

	if after != nil {
		params["after"] = fmt.Sprintf("%d", after.UnixMilli())
	}

	result, err := s.client.GET("/managrams", params)
	if err != nil {
		return nil, fmt.Errorf("Mana: Managrams: %w: %w", ErrorGETFailed, err)
	}

	managrams := make([]Txn, 0)
	err = json.Unmarshal(result, &managrams)
	if err != nil {
		return nil, fmt.Errorf("Mana: Managrams: %w: %w", ErrorFailedToParseResponse, err)
	}

	return managrams, nil
}

// Managram sends a Managram to one or more users.
//
// Parameters:
//   - toIDs: A list of user IDs to send the Managram to. Required.
//   - amount: The amount of Mana to send. Required.
//   - message: An optional message to include with the Managram. Optional.
//
// Returns:
//   - error: An error object if the request fails or if the response cannot be parsed.
func (s *ManaService) Managram(toIDs []string, amount float64, message *string) error {
	body := map[string]interface{}{
		"toIds":  toIDs,
		"amount": amount,
	}

	if message != nil {
		body["message"] = *message
	}

	_, err := s.client.POST("/managram", body)
	if err != nil {
		return fmt.Errorf("Mana: Managram: %w: %w", ErrorPOSTFailed, err)
	}

	return nil
}
