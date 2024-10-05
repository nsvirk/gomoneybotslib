package mbconnect

import (
	"fmt"
	"net/http"
	"net/url"
)

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

func (c *Client) GenerateUserSession(password, totpSecret string) (*UserSession, error) {
	totpValue, err := c.GenerateTotpValue(totpSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to generate TOTP value: %w", err)
	}

	params := url.Values{
		"user_id":    {c.userId},
		"password":   {password},
		"totp_value": {totpValue},
	}

	var userSession UserSession
	if err := c.doEnvelope(http.MethodPost, URISessionLogin, params, nil, &userSession); err != nil {
		return nil, fmt.Errorf("failed to generate user session: %w", err)
	}

	c.SetEnctoken(userSession.Enctoken)
	return &userSession, nil
}

func (c *Client) GenerateTotpValue(totpSecret string) (string, error) {
	params := url.Values{
		"user_id":     {c.userId},
		"totp_secret": {totpSecret},
	}

	var totpValue string
	if err := c.doEnvelope(http.MethodPost, URISessionTotp, params, nil, &totpValue); err != nil {
		return "", fmt.Errorf("failed to generate TOTP value: %w", err)
	}

	return totpValue, nil
}

func (c *Client) DeleteUserSession() (bool, error) {
	if c.enctoken == "" {
		return false, fmt.Errorf("no enctoken set, please login first")
	}

	params := url.Values{
		"user_id":  {c.userId},
		"enctoken": {url.QueryEscape(c.enctoken)},
	}

	var deleteResponse bool
	if err := c.doEnvelope(http.MethodDelete, URISessionLogout, params, nil, &deleteResponse); err != nil {
		return false, fmt.Errorf("failed to delete user session: %w", err)
	}

	c.enctoken = ""
	return deleteResponse, nil
}
