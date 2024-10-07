package main

import (
	"log"
)

// -----------------------------------------------------
// POST /session/token
// -----------------------------------------------------
func (t *APITest) GenerateUserSession() {
	password := t.cfg.KitePassword
	totpSecret := t.cfg.KiteTotpSecret
	// password = ""
	// totpSecret = ""
	userSession, err := t.mbClient.GenerateUserSession(password, totpSecret)
	if err != nil {
		log.Fatalf("Error creating user session: %v", err)
	}
	t.userSession = userSession
	PrettyPrint("UserSession", userSession, 1, userSession)
}

// -----------------------------------------------------
// POST /session/totp
// -----------------------------------------------------
func (t *APITest) GenerateTotpValue() {
	totpSecret := t.cfg.KiteTotpSecret
	// totpSecret = ""
	totpValue, err := t.mbClient.GenerateTotpValue(totpSecret)
	if err != nil {
		log.Fatalf("Error generating totp value: %v", err)
	}
	PrettyPrint("TotpValue", totpValue, 1, totpValue)
}

// -----------------------------------------------------
// DELETE /session/token
// -----------------------------------------------------
func (t *APITest) DeleteUserSession() {
	userID := t.userSession.UserID
	enctoken := t.userSession.Enctoken
	// userID = ""
	// enctoken = ""
	deleteResp, err := t.mbClient.DeleteUserSession(userID, enctoken)
	if err != nil {
		log.Fatalf("Error deleting user session: %v", err)
	}
	PrettyPrint("DeleteUserSession", deleteResp, 1, deleteResp)
}

// -----------------------------------------------------
// POST /session/valid
// -----------------------------------------------------
func (t *APITest) CheckEnctokenValid() {
	enctoken := t.userSession.Enctoken
	// enctoken = ""
	valid, err := t.mbClient.CheckEnctokenValid(enctoken)
	if err != nil {
		log.Fatalf("Error checking enctoken validity: %v", err)
	}
	PrettyPrint("CheckEnctokenValid", enctoken, 1, valid)
}
