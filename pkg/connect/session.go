package mbconnect

import (
	"net/http"
	"net/url"
)

// UserSession is a struct that represents a user session
type UserSession struct {
	UserID        string `json:"user_id"`
	UserName      string `json:"user_name"`
	UserShortname string `json:"user_shortname"`
	AvatarURL     string `json:"avatar_url"`
	PublicToken   string `json:"public_token"`
	KfSession     string `json:"kf_session"`
	Enctoken      string `json:"enctoken"`
	LoginTime     string `json:"login_time"`
}

// POST /session/token - Generate a user session
func (c *Client) GenerateUserSession(password, totpSecret string) (*UserSession, error) {
	totpValue, err := c.GenerateTotpValue(totpSecret)
	if err != nil {
		return nil, err
	}
	params := url.Values{
		"user_id":    {c.userId},
		"password":   {password},
		"totp_value": {totpValue},
	}
	var userSession UserSession
	if err := c.doEnvelope(http.MethodPost, URISessionLogin, params, nil, &userSession); err != nil {
		return nil, err
	}
	c.SetEnctoken(userSession.Enctoken)
	return &userSession, nil
}

// POST /session/totp - Generate a totp value
func (c *Client) GenerateTotpValue(totpSecret string) (string, error) {
	params := url.Values{
		"user_id":     {c.userId},
		"totp_secret": {totpSecret},
	}
	var totpValue string
	if err := c.doEnvelope(http.MethodPost, URISessionTotp, params, nil, &totpValue); err != nil {
		return "", err
	}
	return totpValue, nil
}

// DELETE /session/token - Delete a user session
func (c *Client) DeleteUserSession(userID, enctoken string) (bool, error) {
	enctoken = url.QueryEscape(enctoken)
	params := url.Values{
		"user_id":  {userID},
		"enctoken": {enctoken},
	}
	var deleteResponse bool
	if err := c.doEnvelope(http.MethodDelete, URISessionLogout, params, nil, &deleteResponse); err != nil {
		return false, err
	}
	c.enctoken = ""
	return deleteResponse, nil
}

// POST /session/valid - Check if the `enctoken` is valid
func (c *Client) CheckEnctokenValid(enctoken string) (bool, error) {
	params := url.Values{
		"user_id":  {c.userId},
		"enctoken": {c.enctoken},
	}
	var validResponse bool
	if err := c.doEnvelope(http.MethodPost, URISessionValid, params, nil, &validResponse); err != nil {
		return false, err
	}
	return validResponse, nil
}
