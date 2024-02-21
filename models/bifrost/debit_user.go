package bifrost

type DebitUserRequest struct {
	Currency        *string      `json:"currency,omitempty"`
	Amount          *Amount      `json:"amount,omitempty"`
	Description     *string      `json:"description,omitempty"`
	TransactionId   *string      `json:"transactionId,omitempty"`
	TicketId        *string      `json:"ticketId,omitempty"`
	BonusAmount     *BonusAmount `json:"bonusAmount,omitempty"`
	Ticket          *Ticket      `json:"ticket,omitempty"`
	TransactionType *int         `json:"transactionType,omitempty"`
	Token           *string      `json:"token,omitempty"`
	UserId          *string      `json:"userId,omitempty"`
}

type DebitUserResponse struct {
	ErrorCode             *int64  `json:"errorCode,omitempty"`
	ErrorMessage          *string `json:"errorMessage,omitempty"`
	OperatorTransactionId *string `json:"operatorTransactionId,omitempty"`
}
