package bifrost

type UserDetailsRequest struct {
	Token *string `json:"token,omitempty"`
}

type UserDetailsResponse struct {
	Currency *string `json:"currency,omitempty"`
	Language *string `json:"language,omitempty"`
	Nickname *string `json:"nickname,omitempty"`
	UserId   *string `json:"userId,omitempty"`

	ErrorCode    *int64  `json:"errorCode,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty"`
}
