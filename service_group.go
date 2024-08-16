package manifold

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// GroupService provides methods for interacting with groups,
// including retrieving a list of groups, getting details of a specific group by slug or ID.
type GroupService struct {
	client *Client
}

// Groups retrieves a list of groups based on optional filtering criteria.
//
// Parameters:
//   - beforeTime: Only return groups created before this timestamp. Optional.
//   - availableToUserID: Filter groups that are available to the specified user ID. Optional.
//
// Returns:
//   - []Group: A slice of groups matching the specified criteria.
//   - error: An error object if the request fails or if the response cannot be parsed.
func (s *GroupService) Groups(beforeTime *time.Time, availableToUserID *string) ([]Group, error) {
	params := make(map[string]string, 2)

	if beforeTime != nil {
		params["beforeTime"] = fmt.Sprintf("%d", beforeTime.UnixMilli())
	}

	if availableToUserID != nil {
		params["availableToUserID"] = *availableToUserID
	}

	result, err := s.client.GET(
		"/groups", params,
	)
	if err != nil {
		return nil, fmt.Errorf("Group: Groups: %w: %w", ErrorGETFailed, err)
	}

	groups := make([]Group, 0)
	err = json.Unmarshal(result, &groups)
	if err != nil {
		return nil, fmt.Errorf("Group: Groups: %w: %w", ErrorFailedToParseResponse, err)
	}

	return groups, nil
}

// Group retrieves the details of a specific group using its slug.
//
// Parameters:
//   - slug: The slug of the group to retrieve. Required.
//
// Returns:
//   - *Group: A pointer to the retrieved group object.
//   - error: An error object if the request fails or if the response cannot be parsed.
func (s *GroupService) Group(slug string) (*Group, error) {
	result, err := s.client.GET(
		fmt.Sprintf("/group/%s", url.PathEscape(slug)), nil,
	)
	if err != nil {
		return nil, fmt.Errorf("Group: Group: %w: %w", ErrorGETFailed, err)
	}

	group := new(Group)
	err = json.Unmarshal(result, group)
	if err != nil {
		return nil, fmt.Errorf("Group: Group: %w: %w", ErrorFailedToParseResponse, err)
	}

	return group, nil
}

// ID retrieves the details of a specific group using its ID.
//
// Parameters:
//   - id: The ID of the group to retrieve. Required.
//
// Returns:
//   - *Group: A pointer to the retrieved group object.
//   - error: An error object if the request fails or if the response cannot be parsed.
func (s *GroupService) ID(id string) (*Group, error) {
	result, err := s.client.GET(
		fmt.Sprintf("/group/by-id/%s", url.PathEscape(id)), nil,
	)
	if err != nil {
		return nil, fmt.Errorf("Group: ID: %w: %w", ErrorGETFailed, err)
	}

	group := new(Group)
	err = json.Unmarshal(result, group)
	if err != nil {
		return nil, fmt.Errorf("Group: ID: %w: %w", ErrorFailedToParseResponse, err)
	}

	return group, nil
}
