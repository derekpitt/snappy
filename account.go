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

// Accounts gets all of the accounts that you have access to
func (s *Snappy) Accounts() (a []Account, err error) {
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

// Staff returns all of the staff associated with an account
func (s *Snappy) Staff(accountID int) (staff []Employee, err error) {
  up := urlAndParams{
    url: fmt.Sprintf("/account/%d/staff", accountID),
  }
  err = s.unmarshalJSONAtURL(up, &staff)
  return
}

// Mailboxes returns all of the mailboxes associated with an account
func (s *Snappy) Mailboxes(accountID int) (mailboxes []Mailbox, err error) {
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

// ContactByID returns a Contact matching a contactID
func (s *Snappy) ContactByID(accountID, contactID int) (contact Contact, err error) {
  up := urlAndParams{
    url: fmt.Sprintf("/account/%d/contacts/%d", accountID, contactID),
  }
  err = s.unmarshalJSONAtURL(up, &contact)
  return
}

// ContactByEmail returns a Contact matching an email address
func (s *Snappy) ContactByEmail(accountID int, email string) (contact Contact, err error) {
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

// Search returns all of the tickets that match the query. You can specify a page.
// page should start at 1. SearchResults.Meta.Total contains information you can use to determine how
// many pages there are.
func (s *Snappy) Search(accountID int, query string, page int) (results SearchResults, err error) {
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

// Documents gets all documents for an account
func (s *Snappy) Documents(accountID int) (documents []Document, err error) {
  up := urlAndParams{
    url: fmt.Sprintf("/account/%d/documents", accountID),
  }
  err = s.unmarshalJSONAtURL(up, &documents)
  return
}

// DownloadDocument downloads an attachment.
// Close the read closer after you are done with it please :)
func (s *Snappy) DownloadDocument(accountID, documentID int) (rc io.ReadCloser, err error) {
  up := urlAndParams{
    url: fmt.Sprintf("/account/%d/document/%d/download", accountID, documentID),
  }

  return s.get(up)
}

// WallPost holds information about a Wall Post
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

// NewWallPost holds information for a New Wall Post
type NewWallPost struct {
  Content string   `json:"content"`
  Type    string   `json:"type"`
  Tags    []string `json:"tags"`

  TicketID int `json:"ticket,omitempty"`
  NoteID   int `json:"note,omitempty"`
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

// Wall gets the latest 25 wall posts
func (s *Snappy) Wall(accountID int) (posts []WallPost, err error) {
  up := urlAndParams{
    url: fmt.Sprintf("/account/%d/wall", accountID),
  }
  err = s.unmarshalJSONAtURL(up, &posts)
  return
}

// WallAfter gets 25 wall posts that come after afterWallPostID
func (s *Snappy) WallAfter(accountID, afterWallPostID int) (posts []WallPost, err error) {
  up := urlAndParams{
    url: fmt.Sprintf("/account/%d/wall", accountID),
    params: map[string][]string{
      "after": []string{strconv.Itoa(afterWallPostID)},
    },
  }
  err = s.unmarshalJSONAtURL(up, &posts)
  return
}

// LikeWallPost will like a wall post
func (s *Snappy) LikeWallPost(accountID, wallPostID int) (err error) {
  up := urlAndParams{
    url: fmt.Sprintf("/account/%d/wall/%d/like", accountID, wallPostID),
  }

  rc, err := s.post(up, "", nil)
  defer rc.Close()

  return
}

// UnlikeWallPost will unlike a wall post
func (s *Snappy) UnlikeWallPost(accountID, wallPostID int) (err error) {
  up := urlAndParams{
    url: fmt.Sprintf("/account/%d/wall/%d/like", accountID, wallPostID),
  }

  rc, err := s.del(up)
  defer rc.Close()

  return
}

// CommentWallPost will comment on a wall post
func (s *Snappy) CommentWallPost(accountID, wallPostID int, comment string) (err error) {
  up := urlAndParams{
    url: fmt.Sprintf("/account/%d/wall/%d/comment", accountID, wallPostID),
  }

  rc, err := s.postForm(up, map[string][]string{
    "content": []string{comment},
  })
  defer rc.Close()

  return
}

// DeleteComment will delete a comment
func (s *Snappy) DeleteComment(accountID, wallPostID, commentID int) (err error) {
  up := urlAndParams{
    url: fmt.Sprintf("/account/%d/wall/%d/comment/%d", accountID, wallPostID, commentID),
  }

  rc, err := s.del(up)
  defer rc.Close()

  return
}

// CreateWallPost creates a wall post using NewWallPost
func (s *Snappy) CreateWallPost(accountID int, newPost NewWallPost) (err error) {
  up := urlAndParams{
    url: fmt.Sprintf("/account/%d/wall", accountID),
  }

  _, err = s.postAsJSON(up, newPost)

  return
}

// DeleteWallPost deletes a wall post
func (s *Snappy) DeleteWallPost(accountID, wallPostID int) (err error) {
  up := urlAndParams{
    url: fmt.Sprintf("/account/%d/wall/%d", accountID, wallPostID),
  }

  rc, err := s.del(up)
  defer rc.Close()

  return
}
