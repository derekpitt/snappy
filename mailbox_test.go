package snappy

import (
  "fmt"
  "net/http"
  "testing"
)

func TestWaitingAtMailbox(t *testing.T) {
  setup()
  defer teardown()

  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, `
    [
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
    ]
    `)
  })

  tickets, err := client.WaitingAtMailbox(1)

  if err != nil {
    t.Error("Expected no error in WaitingAtMailbox()")
  }

  if len(tickets) != 1 {
    t.Error("len(tickets) != 1")
  }

  // TODO: compare deeply

}

func TestInboxAtMailbox(t *testing.T) {
  setup()
  defer teardown()

  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, `
    [
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
    ]
    `)
  })

  tickets, err := client.InboxAtMailbox(1)

  if err != nil {
    t.Error("Expected no error in InboxAtMailbox()")
  }

  if len(tickets) != 1 {
    t.Error("len(tickets) != 1")
  }

}

func TestYoursAtMailbox(t *testing.T) {
  setup()
  defer teardown()

  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, `
    [
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
    ]
    `)
  })

  tickets, err := client.YoursAtMailbox(1)

  if err != nil {
    t.Error("Expected no error in YoursAtMailbox()")
  }

  if len(tickets) != 1 {
    t.Error("len(tickets) != 1")
  }

}
