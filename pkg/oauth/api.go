package oauth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"golang.org/x/oauth2"
)

const (
	state = "state"
)

type Client struct {
	Oauth         *oauth2.Config
	IntrospectURI string
}

func NewClient(ClientID, ClientSecret, Add string, Scopes []string, redirectURl string) *Client {
	tokenURL := Add + "/v1/oauth/tokens"
	authURL := Add + "/web/authorize"
	introspectURI := Add + "/v1/oauth/introspect"
	return &Client{
		Oauth: &oauth2.Config{
			ClientID:     ClientID,
			ClientSecret: ClientSecret,
			Scopes:       Scopes,
			Endpoint: oauth2.Endpoint{
				TokenURL: tokenURL,
				AuthURL:  authURL,
			},
			RedirectURL: redirectURl,
		},
		IntrospectURI: introspectURI,
	}
}

func (c *Client) GetToken(ctx context.Context, code string) (*oauth2.Token, error) {
	return c.Oauth.Exchange(ctx, code)
}
func (c *Client) AuthCodeURL(urlParams map[string]string) string {
	oauthUrlParams := []oauth2.AuthCodeOption{oauth2.AccessTypeOffline}

	for k, v := range urlParams {
		oauthUrlParams = append(oauthUrlParams, oauth2.SetAuthURLParam(k, v))
	}

	return c.Oauth.AuthCodeURL(state, oauthUrlParams...)
}

func (c *Client) IntrospectToken(ctx context.Context, token string, typeToken string) (UserInfo, error) {
	return introspectionToken(ctx, c.IntrospectURI, token, typeToken, c.Oauth)
}

func introspectionToken(ctx context.Context, addr, token string, typeToken string, cfg *oauth2.Config) (UserInfo, error) {
	if cfg == nil {
		return UserInfo{}, fmt.Errorf("no config")
	}
	// Use the custom HTTP client when requesting a token.

	httpClient := &http.Client{Timeout: 2 * time.Second}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, addr,
		strings.NewReader(fmt.Sprintf("token=%s&token_type_hint=%s", token, typeToken)))
	if err != nil {
		return UserInfo{}, err
	}

	req.SetBasicAuth(cfg.ClientID, cfg.ClientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return UserInfo{}, err
	}
	defer resp.Body.Close()

	var userInfo UserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return UserInfo{}, err
	}

	return userInfo, nil

}
