package manifold

import (
	"encoding/json"
	"fmt"
)

// CommentService provides methods for interacting with comments on contracts,
// including retrieving, posting text, HTML, and Markdown comments.
type CommentService struct {
	client *Client
}

// Comments retrieves a list of comments for a specific contract.
//
// Parameters:
//   - contractID: Filter comments by the ID of the contract. Optional.
//   - contractSlug: Filter comments by the slug of the contract. Optional.
//   - limit: Limits the number of results returned. Must be between 0 and 1000. Optional.
//   - offset: Skips the specified number of comments before returning results. Must be 0 or greater. Optional.
//   - userID: Filter comments by the ID of the user who posted them. Optional.
//
// Returns:
//   - []Comment: A slice of comments matching the specified criteria.
//   - error: An error object if the request fails or if input validation fails.
func (s *CommentService) Comments(contractID *string, contractSlug *string, limit *int, offset *int, userID *string) ([]Comment, error) {
	params := make(map[string]string, 5)

	if contractID != nil {
		params["contractId"] = *contractID
	}

	if contractSlug != nil {
		params["contractSlug"] = *contractSlug
	}

	if limit != nil {
		if err := checkInRange(*limit, 0, 1000); err != nil {
			return nil, fmt.Errorf("Comment: Comments(limit): %w", err)
		}

		params["limit"] = fmt.Sprintf("%d", *limit)
	}

	if offset != nil {
		if *offset < 0 {
			return nil, fmt.Errorf("Comment: Comments(offset): invalid value: %v, must be greater than 0", *offset)
		}

		params["offset"] = fmt.Sprintf("%d", *offset)
	}

	if userID != nil {
		params["userId"] = *userID
	}

	result, err := s.client.GET("/comments", params)
	if err != nil {
		return nil, fmt.Errorf("Comment: Comments: %w: %w", ErrorGETFailed, err)
	}

	comments := make([]Comment, 0)
	err = json.Unmarshal(result, &comments)
	if err != nil {
		return nil, fmt.Errorf("Comment: Comments: %w: %w", ErrorFailedToParseResponse, err)
	}

	return comments, nil
}

// Comment posts a json TipTap comment on a contract.
//
// Parameters:
//   - id: The ID of the contract to comment on. Required.
//   - content: The text content of the comment. Required.
//
// Returns:
//   - error: An error object if the request fails or if the response cannot be parsed.
func (s *CommentService) Comment(id string, content string) error {
	body := map[string]string{
		"contractId": id,
		"content":    content,
	}

	_, err := s.client.POST("/comment", body)
	if err != nil {
		return fmt.Errorf("Comment: Comment: %w: %w", ErrorPOSTFailed, err)
	}

	return nil
}

// CommentHTML posts an HTML comment on a contract.
//
// Parameters:
//   - id: The ID of the contract to comment on. Required.
//   - content: The HTML content of the comment. Required.
//
// Returns:
//   - error: An error object if the request fails or if the response cannot be parsed.
func (s *CommentService) CommentHTML(id string, content string) error {
	body := map[string]string{
		"contractId": id,
		"html":       content,
	}

	_, err := s.client.POST("/comment", body)
	if err != nil {
		return fmt.Errorf("Comment: CommentHTML: %w: %w", ErrorPOSTFailed, err)
	}

	return nil
}

// CommentMarkdown posts a Markdown-formatted comment on a contract.
//
// Parameters:
//   - id: The ID of the contract to comment on. Required.
//   - content: The Markdown content of the comment. Required.
//
// Returns:
//   - error: An error object if the request fails or if the response cannot be parsed.
func (s *CommentService) CommentMarkdown(id string, content string) error {
	body := map[string]string{
		"contractId": id,
		"markdown":   content,
	}

	_, err := s.client.POST("/comment", body)
	if err != nil {
		return fmt.Errorf("Comment: CommentMarkdown: %w: %w", ErrorPOSTFailed, err)
	}

	return nil
}
