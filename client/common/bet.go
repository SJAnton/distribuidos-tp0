package common

import (
	"fmt"
	"strconv"
	"strings"
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

// Bet constructor that uses an array made of its values
func NewBetFromList(agencyStr string, betList []string) *Bet {
	n := len(betList)
	bet := NewBet(
		agencyStr,
		strings.Join(betList[0:n-4], " "),
		betList[n-4],
		betList[n-3],
		betList[n-2],
		betList[n-1],
	)
	return bet
}

// Constructs a message with the Bet's info as follows:
// "agency|name|surname|id|dob|number"
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

// Contructs a collection of batches from bets read
// from a CSV file, passed as a parameter
func GetBatches(records [][]string, batchSize int, config ClientConfig) []string {
	var batches []string
	var batch string
	amount := 0
	for _, record := range records {
		bet := NewBetFromList(config.ID, record)
		betMsg := bet.MakeMessage() + "\n"

		if len(batch)+len(betMsg) > batchSize || amount == config.BatchMaxAmount {
			batches = append(batches, strings.TrimSuffix(batch, "\n"))
			batch = ""
			amount = 0
		}
		batch += betMsg
		amount++
	}
	batches = append(batches, strings.TrimSuffix(batch, "\n"))
	return batches
}
