package snappy

import (
  "bytes"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "net/url"
  "testing"
)

func TestTicket(t *testing.T) {
  setup()
  defer teardown()

  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, `
     {
        "id":1,
        "account_id":1,
        "mailbox_id":1,
        "created_via":"email",
        "last_reply_by":"customer",
        "last_reply_at":1387831051,
        "opened_by_staff_id":null,
        "opened_by_contact_id":1,
        "opened_at":1387831051,
        "status":"waiting",
        "first_staff_reply_at":null,
        "default_subject":"Default Subject",
        "summary":"Summary",
        "nonce":"123",
        "next_recipients":{
           "cc":[

           ],
           "bcc":[

           ],
           "to":[
              {
                 "name":"To Test",
                 "address":"test@test.com"
              }
           ]
        },
        "created_at":1387831051,
        "updated_at":"2013-12-23 20:37:33",
        "tags":[
           "@test1",
           "#support"
        ],
        "unread":true,
        "contacts":[
           {
              "id":2,
              "account_id":1,
              "first_name":"Test",
              "last_name":"1",
              "value":"test@test.com",
              "provider":"email",
              "created_at":"2013-12-23 20:37:31",
              "updated_at":"2013-12-23 20:37:31",
              "address":"test@test.com",
              "type":"from"
           }
        ],
        "mailbox":{
           "id":1,
           "account_id":1,
           "type":"email",
           "address":"test@test.besnappy.com",
           "display":"Display",
           "auto_responding":1,
           "auto_response":"Auto Response",
           "active":1,
           "created_at":"2013-07-04 04:16:56",
           "updated_at":"2013-09-03 01:25:41",
           "custom_address":"test@test.com",
           "theme":"snappy",
           "local_part":"hello"
        },
        "opener":{
           "id":1,
           "account_id":1,
           "first_name":"Test",
           "last_name":"1",
           "value":"test@test.com",
           "provider":"email",
           "created_at":"2013-12-23 20:37:31",
           "updated_at":"2013-12-23 20:37:31",
           "address":"test@test.com"
        }
     }
    `)
  })

  _, err := client.Ticket(1)

  if err != nil {
    t.Error("Expected no error in Ticket()")
  }

}

func TestTicketNotes(t *testing.T) {
  setup()
  defer teardown()

  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, `
    [
       {
          "id":1,
          "account_id":1,
          "ticket_id":1,
          "facebook_message":null,
          "created_by_staff_id":null,
          "created_by_contact_id":1,
          "scope":"public",
          "created_at":1387832548,
          "updated_at":"2013-12-23 21:02:28",
          "content":"Content",
          "system":0,
          "contacts":[
             {
                "id":1,
                "account_id":1,
                "first_name":"Test",
                "last_name":"1",
                "value":"test@test.com",
                "provider":"email",
                "created_at":"2013-12-23 21:02:28",
                "updated_at":"2013-12-23 21:02:28",
                "address":"test@test.com",
                "type":"from"
             }
          ],
          "staff_creator":null,
          "contact_creator":{
             "id":1,
             "account_id":1,
             "first_name":"Test",
             "last_name":"1",
             "value":"test@test.com",
             "provider":"email",
             "created_at":"2013-12-23 21:02:28",
             "updated_at":"2013-12-23 21:02:28",
             "address":"test@test.com"
          },
          "attachments":[

          ],
          "creator":{
             "id":1,
             "account_id":1,
             "first_name":"Test",
             "last_name":"1",
             "value":"test@test.com",
             "provider":"email",
             "created_at":"2013-12-23 21:02:28",
             "updated_at":"2013-12-23 21:02:28",
             "address":"test@test.com"
          }
       }
    ]
    `)
  })

  notes, err := client.TicketNotes(1)

  if err != nil {
    t.Error("Expected no error in TicketNotes()")
  }

  if len(notes) != 1 {
    t.Error("Expected len(notes) == 1")
  }

}

func TestDownloadTicketAttachment(t *testing.T) {
  setup()
  defer teardown()

  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, `hey now!`)
  })

  rc, err := client.DownloadTicketAttachment(1, 1)

  if err != nil {
    t.Error("Expected no error in DownloadTicketAttachment()")
  }

  defer rc.Close()

  readBytes, _ := ioutil.ReadAll(rc)

  if !bytes.Equal(readBytes, []byte("hey now!")) {
    t.Error("expected 'hey now!'")
  }
}

func TestUpdateTags(t *testing.T) {
  expectedTags := []string{"test1", "test2"}
  setup()
  defer teardown()

  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
      t.Error("Expected POST method in UpdateTags()")
    }

    b, _ := ioutil.ReadAll(r.Body)
    defer r.Body.Close()

    parsedTagsValue, err := url.ParseQuery(string(b))

    if err != nil {
      t.Error("Expected no error parsing post data")
      return
    }

    jsonTags := parsedTagsValue["tags"][0]
    gotTags := []string{}
    err = json.Unmarshal([]byte(jsonTags), &gotTags)

    if err != nil {
      t.Error("Expected no error in unmarshaling tags")
    }

    if gotTags[0] != "test1" {
      t.Error("Expected gotTags[0] = 'test1'")
    }

    if gotTags[1] != "test2" {
      t.Error("Expected gotTags[1] = 'test2'")
    }

    w.WriteHeader(http.StatusOK)
  })

  err := client.UpdateTags(1, expectedTags...)

  if err != nil {
    t.Error("Expected no error in UpdateTags()")
  }
}
