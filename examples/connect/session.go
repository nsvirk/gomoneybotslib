package main

import (
	"log"
)

func (t *APITest) GenerateUserSession() {
	userSession, err := t.mbClient.GenerateUserSession(t.cfg.KitePassword, t.cfg.KiteTotpSecret)
	if err != nil {
		log.Fatalf("Error creating user session: %v", err)
	}
	t.mbClient.SetEnctoken(userSession.Enctoken)
	PrettyPrint("UserSession", userSession)
}

func (t *APITest) GenerateTotpValue() {
	totpValue, err := t.mbClient.GenerateTotpValue(t.cfg.KiteTotpSecret)
	if err != nil {
		log.Fatalf("Error generating totp value: %v", err)
	}
	PrettyPrint("TotpValue", totpValue)
}

func (t *APITest) DeleteUserSession() {
	deleteResp, err := t.mbClient.DeleteUserSession()
	if err != nil {
		log.Fatalf("Error deleting user session: %v", err)
	}
	PrettyPrint("DeleteUserSession", deleteResp)
}
