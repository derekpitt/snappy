package snappy

import (
  "fmt"
  "io"
)

// Ticket holds information about a ticket
type Ticket struct {
  ID                int      `json:"id"`
  AccountID         int      `json:"account_id"`
  MailboxID         int      `json:"mailbox_id"`
  CreatedVia        string   `json:"created_via"`
  LastReplyBy       string   `json:"last_reply_by"`
  LastReplyAt       int      `json:"last_reply_at"`
  OpenedByStaffID   int      `json:"opened_by_staff_id"`
  OpenedByContactID int      `json:"opened_by_contact_id"`
  OpenedAt          int      `json:"opened_at"`
  Status            string   `json:"status"`
  FirstStaffReplyAt string   `json:"first_staff_reply_at"`
  DefaultSubject    string   `json:"default_subject"`
  Summary           string   `json:"summary"`
  CreatedAt         int      `json:"created_at"`
  UpdatedAt         string   `json:"updated_at"`
  Unread            bool     `json:"unread"`
  Tags              []string `json:"tags"`

  Contacts []Contact `json:"contacts"`
  Mailbox  Mailbox   `json:"mailbox"`
  Opener   Contact   `json:"opener"`
}

// Ticket gets the details of a ticket
func (s *Snappy) Ticket(ticketID int) (ticket Ticket, err error) {
  up := urlAndParams{
    url: fmt.Sprintf("/ticket/%d", ticketID),
  }

  err = s.unmarshalJSONAtURL(up, &ticket)
  return
}

// Document holds information about a document. Can be a Document on the account or a document attached to a ticket
type Document struct {
  ID         int    `json:"id"`
  AccountID  int    `json:"account_id"`
  NoteID     int    `json:"note_id"`
  Filename   string `json:"filename"`
  Type       string `json:"type"`
  Size       int    `json:"size"`
  StorageKey string `json:"storage_key"`
  CreatedAt  string `json:"created_at"`
  UpdatedAt  string `json:"updated_at"`
}

// Note holds information about a note.
type Note struct {
  ID                 int    `json:"id"`
  AccountID          int    `json:"account_id"`
  TicketID           int    `json:"ticket_id"`
  CreatedByStaffID   int    `json:"created_by_staff_id"`
  CreatedByContactID int    `json:"created_by_contact_id"`
  Scope              string `json:"scope"`
  CreatedAt          int    `json:"created_at"`
  UpdatedAt          string `json:"updated_at"`
  Content            string `json:"content"`

  Contacts    []Contact  `json:"contacts"`
  Creator     Contact    `json:"creator"`
  Attachments []Document `json:"attachments"`
}

// TicketNotes gets the notes attached to a ticketID
func (s *Snappy) TicketNotes(ticketID int) (notes []Note, err error) {
  up := urlAndParams{
    url: fmt.Sprintf("/ticket/%d/notes", ticketID),
  }

  err = s.unmarshalJSONAtURL(up, &notes)
  return
}

// DownloadTicketAttachment downloads an attachment.
// Close the read closer after you are done with it please :)
func (s *Snappy) DownloadTicketAttachment(ticketID, attachmentID int) (rc io.ReadCloser, err error) {
  up := urlAndParams{
    url: fmt.Sprintf("/ticket/%d/attachment/%d/download", ticketID, attachmentID),
  }

  return s.get(up)
}
