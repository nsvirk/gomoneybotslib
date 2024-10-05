package mbconnect

import (
	"net/http"
	"net/url"
	"testing"
)

// MockHTTPClient is a mock implementation of the HTTPClient interface
type MockHTTPClient struct {
	DoEnvelopeFunc func(method, url string, params url.Values, headers http.Header, obj interface{}) error
}

func (m *MockHTTPClient) Do(method, rURL string, params url.Values, headers http.Header) (HTTPResponse, error) {
	return HTTPResponse{}, nil
}

func (m *MockHTTPClient) DoRaw(method, rURL string, reqBody []byte, headers http.Header) (HTTPResponse, error) {
	return HTTPResponse{}, nil
}

func (m *MockHTTPClient) DoEnvelope(method, url string, params url.Values, headers http.Header, obj interface{}) error {
	return m.DoEnvelopeFunc(method, url, params, headers, obj)
}

func (m *MockHTTPClient) DoJSON(method, url string, params url.Values, headers http.Header, obj interface{}) (HTTPResponse, error) {
	return HTTPResponse{}, nil
}

func (m *MockHTTPClient) GetClient() *httpClient {
	return nil
}

func TestGenerateUserSession(t *testing.T) {
	mockClient := &MockHTTPClient{
		DoEnvelopeFunc: func(method, url string, params url.Values, headers http.Header, obj interface{}) error {
			switch url {
			case URISessionTotp:
				// Handle TOTP generation
				*(obj.(*string)) = "123456"
			case URISessionLogin:
				// Handle user session generation
				session, ok := obj.(*UserSession)
				if !ok {
					t.Fatalf("Expected obj to be *UserSession, got %T", obj)
				}
				*session = UserSession{
					UserID:        "testuser",
					UserName:      "Test User",
					UserShortname: "TU",
					AvatarURL:     "http://example.com/avatar.jpg",
					PublicToken:   "publictoken123",
					KfSession:     "kfsession123",
					Enctoken:      "enctoken123",
					LoginTime:     "2024-10-02 10:00:00",
				}
			default:
				t.Fatalf("Unexpected URL: %s", url)
			}
			return nil
		},
	}

	client := &Client{
		userId:     "testuser",
		httpClient: mockClient,
	}

	session, err := client.GenerateUserSession("password123", "totpsecret123")
	if err != nil {
		t.Fatalf("GenerateUserSession failed: %v", err)
	}

	if session.UserID != "testuser" {
		t.Errorf("Expected UserID 'testuser', got '%s'", session.UserID)
	}
	if session.Enctoken != "enctoken123" {
		t.Errorf("Expected Enctoken 'enctoken123', got '%s'", session.Enctoken)
	}
}

func TestGenerateTotpValue(t *testing.T) {
	mockClient := &MockHTTPClient{
		DoEnvelopeFunc: func(method, url string, params url.Values, headers http.Header, obj interface{}) error {
			*(obj.(*string)) = "123456"
			return nil
		},
	}

	client := &Client{
		userId:     "testuser",
		httpClient: mockClient,
	}

	totpValue, err := client.GenerateTotpValue("totpsecret123")
	if err != nil {
		t.Fatalf("GenerateTotpValue failed: %v", err)
	}

	if totpValue != "123456" {
		t.Errorf("Expected TOTP value '123456', got '%s'", totpValue)
	}
}

func TestDeleteUserSession(t *testing.T) {
	mockClient := &MockHTTPClient{
		DoEnvelopeFunc: func(method, url string, params url.Values, headers http.Header, obj interface{}) error {
			*(obj.(*bool)) = true
			return nil
		},
	}

	client := &Client{
		userId:     "testuser",
		enctoken:   "enctoken123",
		httpClient: mockClient,
	}

	success, err := client.DeleteUserSession()
	if err != nil {
		t.Fatalf("DeleteUserSession failed: %v", err)
	}

	if !success {
		t.Errorf("Expected DeleteUserSession to return true, got false")
	}
}
