package auth

type InternalAuthChecker struct {
	password string
}

func NewInternalAuthChecker(password string) *InternalAuthChecker {
	return &InternalAuthChecker{
		password: password,
	}
}

func (c *InternalAuthChecker) Valid(creds AuthCredentials) (bool, error) {
	return creds.Password == c.password, nil
}
