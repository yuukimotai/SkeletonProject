package responses

type RegisterResponse struct {
	StatusCode int    `json:"statuscode"`
	StatusText string `json:"statustext"`
	Email      string `json:"email"`
}

type LoginResponse struct {
	StatusCode int    `json:"statuscode"`
	StatusText string `json:"statustext"`
	Jwt        string `json:"jwt"`
}
