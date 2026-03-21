package common

import (
	"fmt"
	"strconv"
)

// Bet entity containing its information
type Bet struct {
	Agency  int
	Name    string
	Surname string
	Id      int
	Dob     string
	Number  int
}

// Returns a Bet object made from the parameters passed
func NewBet(agencyStr string, name string, surname string, idStr string, dob string, numberStr string) *Bet {
	agency, agency_err := strconv.Atoi(agencyStr)
	if agency_err != nil {
		log.Criticalf(
			"action: agencyStr_to_int | result: fail | error: %v",
			agency_err,
		)
	}
	id, id_err := strconv.Atoi(idStr)
	if id_err != nil {
		log.Criticalf(
			"action: idStr_to_int | result: fail | error: %v",
			id_err,
		)
	}
	number, number_err := strconv.Atoi(numberStr)
	if number_err != nil {
		log.Criticalf(
			"action: numberStr_to_int | result: fail | error: %v",
			number_err,
		)
	}

	bet := &Bet{
		Agency:  agency,
		Name:    name,
		Surname: surname,
		Id:      id,
		Dob:     dob,
		Number:  number,
	}
	return bet
}

// Constructs a message with the Bet's info as follows:
// "msg_size|agency|name|surname|id|dob|number"
func (bet *Bet) MakeMessage() string {
	return fmt.Sprintf(
		"%d|%s|%s|%d|%s|%d",
		bet.Agency,
		bet.Name,
		bet.Surname,
		bet.Id,
		bet.Dob,
		bet.Number,
	)
}
