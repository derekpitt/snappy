package snappy

// NoteAddress holds information about a Name and an Email address
type NoteAddress struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

// NewNote holds information about a new note
type NewNote struct {
	Subject string `json:"subject"`
	Message string `json:"message"`

	MailboxID   int           `json:"mailbox_id,omitempty"`
	To          []NoteAddress `json:"to,omitempty"`
	From        []NoteAddress `json:"from,omitempty"`
	StaffID     int           `json:"staff_id,omitempty"`
	TicketNonce string        `json:"id,omitempty"`
}

// CreateNote will create a note using NewNote
func (s *Snappy) CreateNote(newNote NewNote) (err error) {
	up := urlAndParams{
		url: "/note",
	}

	_, err = s.postAsJSON(up, newNote)
	return
}
