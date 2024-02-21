package bifrost

import "time"

type Amount struct {
	Amount      *float64 `json:"amount,omitempty"`
	Tax         *float64 `json:"tax,omitempty"`
	TotalAmount *float64 `json:"totalAmount,omitempty"`
}

type BonusAmount struct {
	Free       *float64 `json:"free,omitempty"`
	NonWinning *float64 `json:"nonWinning,omitempty"`
	Tax        *float64 `json:"tax,omitempty"`
	Total      *float64 `json:"total,omitempty"`
	Virtual    *float64 `json:"virtual,omitempty"`
	Winning    *float64 `json:"winning,omitempty"`
}

type Bonus struct {
	Id *string `json:"id,omitempty"`
}

type TicketBet struct {
	BonusStake     *BonusAmount      `json:"bonusStake,omitempty"`
	BonusWonAmount *BonusAmount      `json:"bonusWonAmount,omitempty"`
	Bonuses        []Bonus           `json:"bonuses,omitempty"`
	Id             *string           `json:"id,omitempty"`
	Selections     []TicketSelection `json:"selections,omitempty"`
	Stake          *Amount           `json:"stake,omitempty"`
	Systems        []int             `json:"systems,omitempty"`
	TotalOdds      *float64          `json:"totalOdds,omitempty"`
	Voided         *bool             `json:"voided,omitempty"`
	Won            *bool             `json:"won,omitempty"`
	WonAmount      *Amount           `json:"wonAmount,omitempty"`
}

type TicketEvent struct {
	DateStart      *time.Time        `json:"dateStart,omitempty"`
	Id             *string           `json:"id,omitempty"`
	IsLive         *bool             `json:"isLive,omitempty"`
	Name           *string           `json:"name,omitempty"`
	SportId        *string           `json:"sportId,omitempty"`
	SportName      *string           `json:"sportName,omitempty"`
	Teams          []TicketEventTeam `json:"teams,omitempty"`
	TournamentId   *string           `json:"tournamentId,omitempty"`
	TournamentName *string           `json:"tournamentName,omitempty"`
}

type TicketEventTeam struct {
	Id     *string `json:"id,omitempty"`
	IsHome *bool   `json:"isHome,omitempty"`
	Name   *string `json:"name,omitempty"`
}

type Ticket struct {
	Bets              []TicketBet  `json:"bets,omitempty"`
	BonusStake        *BonusAmount `json:"bonusStake,omitempty"`
	BonusWonAmount    *BonusAmount `json:"bonusWonAmount,omitempty"`
	Date              *time.Time   `json:"date,omitempty"`
	PossibleWinAmount *Amount      `json:"possibleWinAmount,omitempty"`
	Stake             *Amount      `json:"stake,omitempty"`
	Status            *string      `json:"status,omitempty"`
	WonAmount         *Amount      `json:"wonAmount,omitempty"`
}

type TicketOutcome struct {
	// Deprecated: use OutcomeId
	Id         *int64  `json:"id,omitempty"`
	MarketId   *string `json:"marketId,omitempty"`
	MarketName *string `json:"marketName,omitempty"`
	Name       *string `json:"name,omitempty"`
	OutcomeId  *string `json:"outcomeId,omitempty"`
}

type TicketSelection struct {
	Id      *string        `json:"id"`
	Event   *TicketEvent   `json:"event,omitempty"`
	Odds    *float64       `json:"odds,omitempty"`
	Outcome *TicketOutcome `json:"outcome,omitempty"`
	Voided  *bool          `json:"voided,omitempty"`
	Won     *bool          `json:"won,omitempty"`
}
