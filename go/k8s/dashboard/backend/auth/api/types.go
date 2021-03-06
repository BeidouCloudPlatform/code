package api

import (
	"k8s.io/client-go/tools/clientcmd/api"
	"time"
)

// AuthenticationMode represents auth mode supported by dashboard, i.e. basic.
type AuthenticationMode string

const (
	// Resource information that are used as encryption key storage. Can be accessible by multiple dashboard replicas.
	EncryptionKeyHolderName = "kubernetes-dashboard-key-holder"

	// Resource information that are used as certificate storage for custom certificates used by the user.
	CertificateHolderSecretName = "kubernetes-dashboard-certs"

	// Expiration time (in seconds) of tokens generated by dashboard. Default: 15 min.
	DefaultTokenTTL = 900
)

// Authentication modes supported by dashboard should be defined below.
const (
	Token AuthenticationMode = "token"
	Basic AuthenticationMode = "basic"
)

func (self AuthenticationMode) String() string {
	return string(self)
}

// TokenManager is responsible for generating and decrypting tokens used for authorization. Authorization is handled
// by K8S apiserver. Token contains AuthInfo structure used to create K8S api client.
type TokenManager interface {
	// Generate secure token based on AuthInfo structure and save it tokens' payload.
	Generate(api.AuthInfo) (string, error)
	// Decrypt generated token and return AuthInfo structure that will be used for K8S api client creation.
	Decrypt(string) (*api.AuthInfo, error)
	// Refresh returns refreshed token based on provided token. In case provided token has expired, token expiration
	// error is returned.
	Refresh(string) (string, error)
	// SetTokenTTL sets expiration time (in seconds) of generated tokens.
	SetTokenTTL(time.Duration)
}

// AuthManager is used for user authentication management.
type AuthManager interface {
	// Login authenticates user based on provided LoginSpec and returns AuthResponse. AuthResponse contains
	// generated token and list of non-critical errors such as 'Failed authentication'.
	Login(*LoginSpec) (*AuthResponse, error)
	// Refresh takes valid token that hasn't expired yet and returns a new one with expiration time set to TokenTTL. In
	// case provided token has expired, token expiration error is returned.
	Refresh(string) (string, error)
	// AuthenticationModes returns array of auth modes supported by dashboard.
	AuthenticationModes() []AuthenticationMode
	// AuthenticationSkippable tells if the Skip button should be enabled or not
	AuthenticationSkippable() bool
}

// LoginSpec is extracted from request coming from Dashboard frontend during login request. It contains all the
// information required to authenticate user.
type LoginSpec struct {
	// Username is the username for basic authentication to the kubernetes cluster.
	Username string `json:"username,omitempty"`
	// Password is the password for basic authentication to the kubernetes cluster.
	Password string `json:"password,omitempty"`
	// Token is the bearer token for authentication to the kubernetes cluster.
	Token string `json:"token,omitempty"`
	// KubeConfig is the content of users' kubeconfig file. It will be parsed and auth data will be extracted.
	// Kubeconfig can not contain any paths. All data has to be provided within the file.
	KubeConfig string `json:"kubeconfig,omitempty"`
}

// AuthResponse is returned from our backend as a response for login/refresh requests. It contains generated JWEToken
// and a list of non-critical errors such as 'Failed authentication'.
type AuthResponse struct {
	// JWEToken is a token generated during login request that contains AuthInfo data in the payload.
	JWEToken string `json:"jweToken"`
	// Errors are a list of non-critical errors that happened during login request.
	Errors []error `json:"errors"`
}
