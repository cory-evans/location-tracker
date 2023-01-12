package auth

type AuthCredentials struct {
	ClientId string `json:"clientid"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthChecker interface {
	Valid(creds AuthCredentials) (bool, error)
}
