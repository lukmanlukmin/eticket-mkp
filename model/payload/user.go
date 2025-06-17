package payload

// CreateUser ...
type CreateUser struct {
	ID       uint64
	Username string `json:"username"`
	Email    string `json:"email"`
	Msisdn   string `json:"msisdn"`
	Password string `json:"password"`
}
