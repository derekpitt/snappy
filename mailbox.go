package snappy

import (
	"fmt"
)

// Mailbox holds information about a mailbox attached to an account
type Mailbox struct {
	ID             int    `json"id"`
	AccountID      int    `json:"account_id"`
	Type           string `json:"type"`
	Address        string `json:"address"`
	Display        string `json:"display"`
	AutoResponding int    `json:"auto_responding"`
	AutoResponse   string `json:"auto_response"`
	Active         int    `json:"active"`
	CustomAddress  string `json:"custom_address"`
	Theme          string `json:"theme"`
	LocalPart      string `json:"local_part"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}

func (s *Snappy) ticketsAtMailboxEndpoint(mailboxID int, endpoint string) (tickets []Ticket, err error) {
	up := urlAndParams{
		url: fmt.Sprintf("/mailbox/%d/%s", mailboxID, endpoint),
	}

	err = s.unmarshalJSONAtURL(up, &tickets)
	return
}

// WaitingAtMailbox gets the tickets that are a waiting status
func (s *Snappy) WaitingAtMailbox(mailboxID int) (tickets []Ticket, err error) {
	return s.ticketsAtMailboxEndpoint(mailboxID, "tickets")
}

// InboxAtMailbox gets the tickets that are a new and unanssigned status
func (s *Snappy) InboxAtMailbox(mailboxID int) (tickets []Ticket, err error) {
	return s.ticketsAtMailboxEndpoint(mailboxID, "inbox")
}

// YoursAtMailbox get the tickets that are waiting and assigned to you
func (s *Snappy) YoursAtMailbox(mailboxID int) (tickets []Ticket, err error) {
	return s.ticketsAtMailboxEndpoint(mailboxID, "yours")
}
