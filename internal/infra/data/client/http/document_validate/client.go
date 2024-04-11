package documentValidate

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/richhh7g/mm-api-nfe/internal/infra/environment"
)

var (
	ErrInvalidBody = errors.New("invalid body")
)

var _ Client = &ClientImpl{}

type ClientRequestParams struct {
	Url    string
	Method string
	Data   interface{}
}

type Client interface {
	Request(params *ClientRequestParams) (*http.Response, error)
}

type ClientImpl struct {
	ctx        context.Context
	httpClient *http.Client
}

func NewClient(ctx context.Context, httpClient *http.Client) *ClientImpl {
	return &ClientImpl{
		ctx:        ctx,
		httpClient: httpClient,
	}
}

func (c *ClientImpl) Request(params *ClientRequestParams) (*http.Response, error) {
	ctx, cancel := context.WithTimeout(c.ctx, 1*time.Second)
	defer cancel()

	baseUrl, err := environment.Get("MM_DOCUMENT_VALIDATE_URL")
	if err != nil {
		return nil, err
	}

	var body []byte

	if params.Data != nil {
		body, err = json.Marshal(params.Data)

		if err != nil {
			return nil, fmt.Errorf("%w: %v", ErrInvalidBody, params.Data)
		}
	}

	req, err := http.NewRequestWithContext(ctx, params.Method, fmt.Sprintf("%s%s", baseUrl, params.Url), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	apiKey, err := environment.Get("MM_DOCUMENT_VALIDATE_API_KEY")
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey.(string))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
