package manifold

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// UserService provides methods for interacting with user data, including retrieving user details by username or ID,
// listing users, and getting the authenticated user's information.
type UserService struct {
	client *Client
}

// Users retrieves a list of users with optional pagination.
//
// Parameters:
//   - limit: Limits the number of results returned. Must be between 0 and 1000. Optional.
//   - before: Retrieves users before this cursor (e.g., timestamp or ID). Optional.
//
// Returns:
//   - []User: A slice of users matching the specified criteria.
//   - error: An error object if the request fails or if input validation fails.
func (s *UserService) Users(limit *int, before *string) ([]User, error) {
	params := make(map[string]string, 2)

	if limit != nil {
		if err := checkInRange(*limit, 0, 1000); err != nil {
			return nil, fmt.Errorf("User: Users(limit): %w", err)
		}

		params["limit"] = fmt.Sprintf("%d", *limit)
	}

	if before != nil {
		params["before"] = *before
	}

	result, err := s.client.GET(
		"/users", params,
	)
	if err != nil {
		return nil, fmt.Errorf("User: Users: %w: %w", ErrorGETFailed, err)
	}

	users := make([]User, 0)
	err = json.Unmarshal(result, &users)
	if err != nil {
		return nil, fmt.Errorf("User: Users: %w: %w", ErrorFailedToParseResponse, err)
	}

	return users, nil
}

// User retrieves detailed information about a user by their username.
//
// Parameters:
//   - username: The username of the user to retrieve. Required.
//
// Returns:
//   - *User: A pointer to the retrieved user object.
//   - error: An error object if the request fails or if the response cannot be parsed.
func (s *UserService) User(username string) (*User, error) {
	result, err := s.client.GET(
		fmt.Sprintf("/user/%s", url.PathEscape(username)), nil,
	)
	if err != nil {
		return nil, fmt.Errorf("User: User: %w: %w", ErrorGETFailed, err)
	}

	user := new(User)
	err = json.Unmarshal(result, user)
	if err != nil {
		return nil, fmt.Errorf("User: User: %w: %w", ErrorFailedToParseResponse, err)
	}

	return user, nil
}

// UserLite retrieves basic information about a user by their username.
//
// Parameters:
//   - username: The username of the user to retrieve. Required.
//
// Returns:
//   - *DisplayUser: A pointer to the retrieved display user object, containing basic information.
//   - error: An error object if the request fails or if the response cannot be parsed.
func (s *UserService) UserLite(username string) (*DisplayUser, error) {
	result, err := s.client.GET(
		fmt.Sprintf("/user/%s/lite", url.PathEscape(username)), nil,
	)
	if err != nil {
		return nil, fmt.Errorf("User: UserLite: %w: %w", ErrorGETFailed, err)
	}

	user := new(DisplayUser)
	err = json.Unmarshal(result, user)
	if err != nil {
		return nil, fmt.Errorf("User: UserLite: %w: %w", ErrorFailedToParseResponse, err)
	}

	return user, nil
}

// ID retrieves detailed information about a user by their ID.
//
// Parameters:
//   - id: The ID of the user to retrieve. Required.
//
// Returns:
//   - *User: A pointer to the retrieved user object.
//   - error: An error object if the request fails or if the response cannot be parsed.
func (s *UserService) ID(id string) (*User, error) {
	result, err := s.client.GET(
		fmt.Sprintf("/user/by-id/%s", url.PathEscape(id)), nil,
	)
	if err != nil {
		return nil, fmt.Errorf("User: ID: %w: %w", ErrorGETFailed, err)
	}

	user := new(User)
	err = json.Unmarshal(result, user)
	if err != nil {
		return nil, fmt.Errorf("User: ID: %w: %w", ErrorFailedToParseResponse, err)
	}

	return user, nil
}

// IDLite retrieves basic information about a user by their ID.
//
// Parameters:
//   - id: The ID of the user to retrieve. Required.
//
// Returns:
//   - *DisplayUser: A pointer to the retrieved display user object, containing basic information.
//   - error: An error object if the request fails or if the response cannot be parsed.
func (s *UserService) IDLite(id string) (*DisplayUser, error) {
	result, err := s.client.GET(
		fmt.Sprintf("/user/by-id/%s/lite", url.PathEscape(id)), nil,
	)
	if err != nil {
		return nil, fmt.Errorf("User: IDLite: %w: %w", ErrorGETFailed, err)
	}

	user := new(DisplayUser)
	err = json.Unmarshal(result, user)
	if err != nil {
		return nil, fmt.Errorf("User: IDLite: %w: %w", ErrorFailedToParseResponse, err)
	}

	return user, nil
}

// Me retrieves information about the authenticated user.
//
// Returns:
//   - *User: A pointer to the authenticated user's object.
//   - error: An error object if the request fails or if the response cannot be parsed.
func (s *UserService) Me() (*User, error) {
	result, err := s.client.GET(
		"/me", nil,
	)
	if err != nil {
		return nil, fmt.Errorf("User: Me: %w: %w", ErrorGETFailed, err)
	}

	user := new(User)
	err = json.Unmarshal(result, user)
	if err != nil {
		return nil, fmt.Errorf("User: Me: %w: %w", ErrorFailedToParseResponse, err)
	}

	return user, nil
}
