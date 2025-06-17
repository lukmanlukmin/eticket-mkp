package constant

import "errors"

var (
	INVALID_ID                       = errors.New("invalid ID")
	REGISTRATION_NOTFOUND_OR_EXPIRED = errors.New("registration not found or expired")
	FAIL_CREATE_USER                 = errors.New("fail to create user")
	USER_ALREADY_REGISTERED          = errors.New("user already registered")
	FAIL_TO_PROCCESS_REGISTRATION    = errors.New("failed to proccess data registration")

	INVALID_USER_CREDENTIAL   = errors.New("invalid user credentials")
	FAIL_TO_GENERATE_TOKEN    = errors.New("failed to generate token")
	FAIL_TO_VALIDATE_TOKEN_ID = errors.New("failed to validate token id")
)
