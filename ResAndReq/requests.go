package ResAndReq

type RegistRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// `json:"email,omitempty" valid:"email"`
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
