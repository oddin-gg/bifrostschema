package bifrost

type CreditUserRequest struct {
	Amount                     *Amount      `json:"amount,omitempty"`
	BonusAmount                *BonusAmount `json:"bonusAmount,omitempty"`
	Currency                   *string      `json:"currency,omitempty"`
	DebitOperatorTransactionId *string      `json:"debitOperatorTransactionId,omitempty"`
	DebitTransactionId         *string      `json:"debitTransactionId,omitempty"`
	Description                *string      `json:"description,omitempty"`
	Ticket                     *Ticket      `json:"ticket,omitempty"`
	Token                      *string      `json:"token,omitempty"`
	TransactionId              *string      `json:"transactionId,omitempty"`
	UserId                     *string      `json:"userId,omitempty"`
	TicketId                   *string      `json:"ticketId,omitempty"`
}

type CreditUserResponse struct {
	ErrorCode             *int64  `json:"errorCode,omitempty"`
	ErrorMessage          *string `json:"errorMessage,omitempty"`
	OperatorTransactionId *string `json:"operatorTransactionId,omitempty"`
}
