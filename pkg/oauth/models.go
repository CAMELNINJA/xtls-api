package oauth

import "fmt"

type UserInfo struct {
	Active    bool   `json:"active"`
	Scope     string `json:"scope"`
	ClientID  string `json:"client_id"`
	Username  string `json:"username"`
	TokenType string `json:"token_type"`
	Exp       int    `json:"exp"`
}

func (U UserInfo) String() string {
	return fmt.Sprintf("Active: %t, Scope: %s, ClientID: %s, Username: %s, TokenType: %s, Exp: %d", U.Active, U.Scope, U.ClientID, U.Username, U.TokenType, U.Exp)
}
