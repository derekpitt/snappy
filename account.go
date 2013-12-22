package snappy

import (
  "fmt"
  "io"
  "net/url"
  "strconv"
)

// Account holds information about am account
type Account struct {
  ID           int    `json:"id"`
  Organization string `json:"organization"`
  Domain       string `json:"domain"`
  PlanID       int    `json:"plan_id"`
  Active       int    `json:"active"`
  CreatedAt    string `json:"created_at"`
  UpdatedAt    string `json:"updated_at"`
  CustomDomain string `json:"custom_domain"`
}

func (s *snappy) Accounts() (a []Account, err error) {
  up := urlAndParams{
    url: "/accounts",
  }
  err = s.unmarshalJSONAtURL(up, &a)
  return
}

// Employee holds information about users that can access Snappy
type Employee struct {
  ID         int    `json:"id"`
  Email      string `json:"email"`
  SMSNumber  string `json:"sms_number"`
  FirstName  string `json:"first_name"`
  LastName   string `json:"last_name"`
  Photo      string `json:"photo"`
  Culture    string `json:"culture"`
  Notify     int    `json:"notify"`
  Signature  string `json:"signature"`
  TourPlayed int    `json:"tour_played"`
  TimeZone   string `json:"timezone"`
  NotifyNew  int    `json:"notify_new"`
  NewsReadAt string `json:"news_read_at"`
  UserName   string `json:"username"`
  CreatedAt  string `json:"created_at"`
  UpdatedAt  string `json:"updated_at"`
  Address    string `json:"address"`
}

func (s *snappy) Staff(accountID int) (staff []Employee, err error) {
  up := urlAndParams{
    url: fmt.Sprintf("/account/%d/staff", accountID),
  }
  err = s.unmarshalJSONAtURL(up, &staff)
  return
}

func (s *snappy) Mailboxes(accountID int) (mailboxes []Mailbox, err error) {
  up := urlAndParams{
    url: fmt.Sprintf("/account/%d/mailboxes", accountID),
  }
  err = s.unmarshalJSONAtURL(up, &mailboxes)
  return
}

// Contact hold information about a contact for an account
type Contact struct {
  ID        int    `json:"id"`
  AccountID int    `json:"account_id"`
  FirstName string `json:"first_name"`
  LastName  string `json:"last_name"`
  Value     string `json:"value"`
  Provider  string `json:"provider"`
  Address   string `json:"address"`
  CreatedAt string `json:"created_at"`
  UpdatedAt string `json:"updated_at"`
}

func (s *snappy) ContactByID(accountID, contactID int) (contact Contact, err error) {
  up := urlAndParams{
    url: fmt.Sprintf("/account/%d/contacts/%d", accountID, contactID),
  }
  err = s.unmarshalJSONAtURL(up, &contact)
  return
}

func (s *snappy) ContactByEmail(accountID int, email string) (contact Contact, err error) {
  up := urlAndParams{
    url: fmt.Sprintf("/account/%d/contacts/%s", accountID, url.QueryEscape(email)),
  }
  err = s.unmarshalJSONAtURL(up, &contact)
  return
}

// SearchResults holds the meta data for a search and the tickets for the current page
type SearchResults struct {
  Meta struct {
    Total int    `json:"total"`
    Page  string `json:"page"`
  } `json:"meta"`

  Tickets []Ticket `json:"data"`
}

func (s *snappy) Search(accountID int, query string, page int) (results SearchResults, err error) {
  up := urlAndParams{
    url: fmt.Sprintf("/account/%d/search", accountID),
    params: map[string][]string{
      "query": []string{query},
      "page":  []string{strconv.Itoa(page)},
    },
  }
  err = s.unmarshalJSONAtURL(up, &results)
  return
}

func (s *snappy) Documents(accountID int) (documents []Document, err error) {
  up := urlAndParams{
    url: fmt.Sprintf("/account/%d/documents", accountID),
  }
  err = s.unmarshalJSONAtURL(up, &documents)
  return
}

// DownloadDocument downloads an attachment.
// Close the read closer after you are done with it please :)
func (s *snappy) DownloadDocument(accountID, documentID int) (rc io.ReadCloser, err error) {
  up := urlAndParams{
    url: fmt.Sprintf("/account/%d/document/%d/download", accountID, documentID),
  }

  return s.get(up)
}

// WallPost hold information about a Wall Post
type WallPost struct {
  ID              int    `json:"id"`
  AccountID       int    `json:"account_id"`
  StaffID         int    `json:"staff_id"`
  TicketID        int    `json:"ticket_id"`
  NoteID          int    `json:"note_id"`
  Type            string `json:"type"`
  Content         string `json:"content"`
  ContentMarkdown string `json:"content_markdown"`

  CreatedAt string `json:"created_at"`
  UpdatedAt string `json:"updated_at"`

  Tags      []string      `json:"tags"`
  Likes     []string      `json:"likes"`
  LikeCount int           `json:"like_count"`
  Comments  []WallComment `json:"comments"`
}

// WallComment holds information about a comment on a Wall Post
type WallComment struct {
  ID              int    `json:"id"`
  WallPostID      int    `json:"post_id"`
  StaffID         int    `json:"staff_id"`
  Content         string `json:"content"`
  ContentMarkdown string `json:"content"`

  CreatedAt string `json:"created_at"`
  UpdatedAt string `json:"updated_at"`

  Staff Employee `json:"staff"`
}

func (s *snappy) Wall(accountID int) (posts []WallPost, err error) {
  up := urlAndParams{
    url: fmt.Sprintf("/account/%d/wall", accountID),
  }
  err = s.unmarshalJSONAtURL(up, &posts)
  return
}

func (s *snappy) WallAfter(accountID, afterWallPostID int) (posts []WallPost, err error) {
  up := urlAndParams{
    url: fmt.Sprintf("/account/%d/wall", accountID),
    params: map[string][]string{
      "after": []string{strconv.Itoa(afterWallPostID)},
    },
  }
  err = s.unmarshalJSONAtURL(up, &posts)
  return
}
