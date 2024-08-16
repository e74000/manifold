package manifold

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Client represents the Manifold API client, used to interact with various services such as users, groups, markets, and more.
// It manages API requests and provides access to all the available services.
type Client struct {
	BaseURL    string       // The base URL for the Manifold API.
	APIKey     string       // The API key used for authentication with the Manifold API.
	HTTPClient *http.Client // The HTTP client used to perform requests.

	User    *UserService    // Service for user-related API calls.
	Group   *GroupService   // Service for group-related API calls.
	Market  *MarketService  // Service for market-related API calls.
	Bet     *BetService     // Service for bet-related API calls.
	Comment *CommentService // Service for comment-related API calls.
	Mana    *ManaService    // Service for mana-related API calls.
}

// NewClient creates a new instance of the Manifold API client.
//
// Parameters:
//   - apiKey: The API key used for authenticating with the Manifold API.
//
// Returns:
//   - *Client: A pointer to the newly created Client instance, pre-configured with services.
func NewClient(apiKey string) *Client {
	c := &Client{
		BaseURL:    "https://api.manifold.markets/v0",
		APIKey:     apiKey,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}

	// Initialize all services associated with the client.
	c.User = &UserService{client: c}
	c.Group = &GroupService{client: c}
	c.Market = &MarketService{client: c}
	c.Bet = &BetService{client: c}
	c.Comment = &CommentService{client: c}
	c.Mana = &ManaService{client: c}

	return c
}

// GET performs a GET request to the Manifold API.
//
// Parameters:
//   - endpoint: The API endpoint to send the GET request to (relative to BaseURL).
//   - params: A map of query parameters to include in the request. Optional.
//
// Returns:
//   - []byte: The response body as a byte slice.
//   - error: An error object if the request fails or if the response cannot be read.
func (c *Client) GET(endpoint string, params map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", c.BaseURL, endpoint), nil)
	if err != nil {
		return nil, err
	}

	// Add query parameters to the request.
	q := req.URL.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	// Add the Authorization header if an API key is provided.
	if c.APIKey != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Key %s", c.APIKey))
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// POST performs a POST request to the Manifold API.
//
// Parameters:
//   - endpoint: The API endpoint to send the POST request to (relative to BaseURL).
//   - body: The body to include in the POST request. Must be serializable to JSON. Optional.
//
// Returns:
//   - []byte: The response body as a byte slice.
//   - error: An error object if the request fails or if the response cannot be read.
func (c *Client) POST(endpoint string, body interface{}) ([]byte, error) {
	var (
		jsonBody []byte
		err      error
	)
	if body != nil {
		// Serialize the body to JSON.
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	} else {
		jsonBody = []byte{}
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.BaseURL, endpoint), bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	// Set request headers.
	req.Header.Add("Content-Type", "application/json")
	if c.APIKey != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Key %s", c.APIKey))
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
