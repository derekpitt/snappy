package snappy

import (
  "bytes"
  "fmt"
  "io/ioutil"
  "net/http"
  "testing"
)

func TestAccounts(t *testing.T) {
  setup()
  defer teardown()

  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, `
    [
        {
            "id": 3,
            "organization": "Snappy Help",
            "domain": "help.besnappy.com",
            "plan_id": 1,
            "active": 1,
            "created_at": "2012-12-05 15:24:20",
            "updated_at": "2013-05-07 19:48:06",
            "custom_domain": ""
        }
    ]
    `)
  })

  accounts, err := client.Accounts()

  if err != nil {
    t.Error("Expected no error in Accounts()")
  }

  if len(accounts) != 1 {
    t.Error("Expected len(accounts) = 1")
  }

  expected := Account{
    ID:           3,
    Organization: "Snappy Help",
    Domain:       "help.besnappy.com",
    PlanID:       1,
    Active:       1,
    CreatedAt:    "2012-12-05 15:24:20",
    UpdatedAt:    "2013-05-07 19:48:06",
    CustomDomain: "",
  }

  if expected != accounts[0] {
    t.Error("Expected account does not match")
  }
}

func TestStaff(t *testing.T) {
  setup()
  defer teardown()

  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, `
    [
       {
          "id":123,
          "email":"test1@test.com",
          "sms_number":"",
          "first_name":"Test",
          "last_name":"1",
          "photo":null,
          "culture":"en",
          "notify":1,
          "created_at":"2013-07-04 04:16:56",
          "updated_at":"2013-12-23 04:20:05",
          "signature":"",
          "tour_played":1,
          "timezone":"America/Chicago",
          "notify_new":1,
          "news_read_at":"2013-12-23 04:20:05",
          "address":"test1@test.com",
          "username":"test1"
       },
       {
          "id":124,
          "email":"test2@test.com",
          "sms_number":"",
          "first_name":"Test",
          "last_name":"2",
          "photo":null,
          "culture":"en",
          "notify":1,
          "created_at":"2013-07-05 16:12:24",
          "updated_at":"2013-12-21 16:11:15",
          "signature":"",
          "tour_played":1,
          "timezone":"America\/Chicago",
          "notify_new":0,
          "news_read_at":"2013-12-21 16:11:15",
          "address":"test2@test.com",
          "username":"test2"
       }
    ]
    `)
  })

  staff, err := client.Staff(1)

  if err != nil {
    t.Error("Expected no error in Staff()")
  }

  if len(staff) != 2 {
    t.Error("Expected len(staff) = 2")
  }

  expected1 := Employee{
    ID:         123,
    Email:      "test1@test.com",
    SMSNumber:  "",
    FirstName:  "Test",
    LastName:   "1",
    Photo:      "",
    Culture:    "en",
    Notify:     1,
    Signature:  "",
    TourPlayed: 1,
    TimeZone:   "America/Chicago",
    NotifyNew:  1,
    NewsReadAt: "2013-12-23 04:20:05",
    UserName:   "test1",
    CreatedAt:  "2013-07-04 04:16:56",
    UpdatedAt:  "2013-12-23 04:20:05",
    Address:    "test1@test.com",
  }

  expected2 := Employee{
    ID:         124,
    Email:      "test2@test.com",
    SMSNumber:  "",
    FirstName:  "Test",
    LastName:   "2",
    Photo:      "",
    Culture:    "en",
    Notify:     1,
    Signature:  "",
    TourPlayed: 1,
    TimeZone:   "America/Chicago",
    NotifyNew:  0,
    NewsReadAt: "2013-12-21 16:11:15",
    UserName:   "test2",
    CreatedAt:  "2013-07-05 16:12:24",
    UpdatedAt:  "2013-12-21 16:11:15",
    Address:    "test2@test.com",
  }

  if expected1 != staff[0] {
    t.Error("expected1 != staff[0]")
  }

  if expected2 != staff[1] {
    t.Error("expected2 != staff[1]")
  }
}

func TestMailboxes(t *testing.T) {
  setup()
  defer teardown()

  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, `
    [
       {
          "id":1,
          "account_id":1,
          "type":"email",
          "address":"test@test.besnappy.com",
          "display":"Test1",
          "auto_responding":0,
          "auto_response":"Test 1",
          "active":1,
          "created_at":"2013-09-06 21:05:07",
          "updated_at":"2013-09-06 21:05:07",
          "custom_address":"test1@test.com",
          "theme":"snappy",
          "local_part":"notifications"
       },
       {
          "id":2,
          "account_id":1,
          "type":"email",
          "address":"test2@test.besnappy.com",
          "display":"Test2",
          "auto_responding":0,
          "auto_response":"Test 2",
          "active":1,
          "created_at":"2013-09-06 21:05:07",
          "updated_at":"2013-09-06 21:05:07",
          "custom_address":"test2@test.com",
          "theme":"snappy",
          "local_part":"notifications"
       }
    ]
    `)
  })

  mbs, err := client.Mailboxes(1)

  if err != nil {
    t.Error("Expected no error in Mailboxes()")
  }

  if len(mbs) != 2 {
    t.Error("Expected len(mbs) = 2")
  }

  expected1 := Mailbox{
    ID:             1,
    AccountID:      1,
    Type:           "email",
    Address:        "test@test.besnappy.com",
    Display:        "Test1",
    AutoResponding: 0,
    AutoResponse:   "Test 1",
    Active:         1,
    CreatedAt:      "2013-09-06 21:05:07",
    UpdatedAt:      "2013-09-06 21:05:07",
    CustomAddress:  "test1@test.com",
    Theme:          "snappy",
    LocalPart:      "notifications",
  }

  expected2 := Mailbox{
    ID:             2,
    AccountID:      1,
    Type:           "email",
    Address:        "test2@test.besnappy.com",
    Display:        "Test2",
    AutoResponding: 0,
    AutoResponse:   "Test 2",
    Active:         1,
    CreatedAt:      "2013-09-06 21:05:07",
    UpdatedAt:      "2013-09-06 21:05:07",
    CustomAddress:  "test2@test.com",
    Theme:          "snappy",
    LocalPart:      "notifications",
  }

  if expected1 != mbs[0] {
    t.Error("expected1 != mbs[0]")
  }

  if expected2 != mbs[1] {
    t.Error("expected2 != mbs[1]")
  }
}

func TestContactByID(t *testing.T) {
  setup()
  defer teardown()

  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, `
    {
       "id":1,
       "account_id":1,
       "first_name":"Test",
       "last_name":"1",
       "value":"test@test.com",
       "provider":"email",
       "created_at":"2013-12-21 15:50:19",
       "updated_at":"2013-12-21 15:50:19",
       "address":"test@test.com"
    }
    `)
  })

  contact, err := client.ContactByID(1, 1)

  if err != nil {
    t.Error("Expected no error in ContactByID()")
  }

  expected := Contact{
    ID:        1,
    AccountID: 1,
    FirstName: "Test",
    LastName:  "1",
    Value:     "test@test.com",
    Provider:  "email",
    CreatedAt: "2013-12-21 15:50:19",
    UpdatedAt: "2013-12-21 15:50:19",
    Address:   "test@test.com",
  }

  if contact != expected {
    t.Error("contact != expected")
  }
}

func TestContactByEmail(t *testing.T) {
  setup()
  defer teardown()

  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, `
    {
       "id":1,
       "account_id":1,
       "first_name":"Test",
       "last_name":"1",
       "value":"test@test.com",
       "provider":"email",
       "created_at":"2013-12-21 15:50:19",
       "updated_at":"2013-12-21 15:50:19",
       "address":"test@test.com"
    }
    `)
  })

  contact, err := client.ContactByEmail(1, "test@test.com")

  if err != nil {
    t.Error("Expected no error in ContactByEmail()")
  }

  expected := Contact{
    ID:        1,
    AccountID: 1,
    FirstName: "Test",
    LastName:  "1",
    Value:     "test@test.com",
    Provider:  "email",
    CreatedAt: "2013-12-21 15:50:19",
    UpdatedAt: "2013-12-21 15:50:19",
    Address:   "test@test.com",
  }

  if contact != expected {
    t.Error("contact != expected")
  }
}

func TestSearch(t *testing.T) {
  setup()
  defer teardown()

  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query()

    if q["page"][0] != "1" {
      t.Error("Expected Search to send 'page' = 1")
    }

    if q["query"][0] != "test" {
      t.Error("Expected Search to send 'query' = 'test'")
    }

    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, `
    {
       "meta":{
          "total":1,
          "page":"1"
       },
       "data":[
          {
             "id":1,
             "account_id":1,
             "mailbox_id":1,
             "created_via":"email",
             "last_reply_by":"customer",
             "last_reply_at":1387643050,
             "opened_by_staff_id":null,
             "opened_by_contact_id":1,
             "opened_at":1387643050,
             "status":"replied",
             "first_staff_reply_at":"2013-12-21 16:27:38",
             "default_subject":"Default Subject",
             "summary":"Summary",
             "next_recipients":{
                "cc":[

                ],
                "bcc":[

                ],
                "to":[
                   {
                      "name":"Test",
                      "address":"test@test.com"
                   }
                ]
             },
             "created_at":1387643050,
             "updated_at":"2013-12-21 18:38:38",
             "tags":[
                "@test",
                "#test"
             ],
             "mailbox":{
                "id":1,
                "account_id":1,
                "type":"email",
                "address":"hello@test.com",
                "display":"Tedt",
                "auto_responding":1,
                "auto_response":"Test Auto Response",
                "active":1,
                "created_at":"2013-07-04 04:16:56",
                "updated_at":"2013-09-03 01:25:41",
                "custom_address":"test@test.com",
                "theme":"snappy",
                "local_part":"hello"
             },
             "contacts":[
                {
                   "id":1,
                   "account_id":1,
                   "first_name":"",
                   "last_name":"",
                   "value":"test@test.com",
                   "provider":"email",
                   "created_at":"2013-07-18 16:10:05",
                   "updated_at":"2013-07-18 16:10:05",
                   "address":"test@test.com",
                   "type":"from"
                }
             ],
             "opener":{
                "id":1,
                "account_id":1,
                "first_name":"",
                "last_name":"",
                "value":"test@test.com",
                "provider":"email",
                "created_at":"2013-07-18 16:10:05",
                "updated_at":"2013-07-18 16:10:05",
                "address":"test@test.com"
             }
          }
       ]
    }
    `)
  })

  results, err := client.Search(1, "test", 1)

  if err != nil {
    t.Error("Expected no error in Search()")
  }

  if results.Meta.Page != "1" {
    t.Error("Meta.Page != '1'")
  }

  if results.Meta.Total != 1 {
    t.Error("Meta.Total != 1")
  }

  if len(results.Tickets) != 1 {
    t.Error("len(results.Tickets) != 1")
  }

  // TODO: deep equal of returned ticket result
}

func TestDocuments(t *testing.T) {
  setup()
  defer teardown()

  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, `
    [
       {
          "id":1,
          "account_id":1,
          "filename":"Filename",
          "type":"image\/png",
          "size":59788,
          "storage_key":"fake",
          "created_at":"2013-07-10 15:41:34",
          "updated_at":"2013-07-10 15:41:34",
          "store":0
       }
    ]
    `)
  })

  docs, err := client.Documents(1)

  if err != nil {
    t.Error("Expected no error in Documents()")
  }

  if len(docs) != 1 {
    t.Error("len(docs) != 1")
  }

  expected := Document{
    ID:         1,
    AccountID:  1,
    Filename:   "Filename",
    Type:       "image/png",
    Size:       59788,
    StorageKey: "fake",
    CreatedAt:  "2013-07-10 15:41:34",
    UpdatedAt:  "2013-07-10 15:41:34",
  }

  if expected != docs[0] {
    t.Error("expected != docs[0]")
  }
}

func TestDownloadDocument(t *testing.T) {
  setup()
  defer teardown()

  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, `hey now!`)
  })

  rc, err := client.DownloadDocument(1, 1)

  if err != nil {
    t.Error("Expected no error in DownloadDocument()")
  }

  defer rc.Close()

  readBytes, _ := ioutil.ReadAll(rc)

  if !bytes.Equal(readBytes, []byte("hey now!")) {
    t.Error("expected 'hey now!'")
  }
}

func TestWall(t *testing.T) {
  setup()
  defer teardown()

  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, `
    [
       {
          "id":1,
          "account_id":1,
          "staff_id":1,
          "ticket_id":null,
          "note_id":null,
          "type":"post",
          "content":"test",
          "created_at":"2013-07-12 18:08:32",
          "updated_at":"2013-07-12 21:14:30",
          "tags":[
             "tag1"
          ],
          "likes":[
             "Like 1"
          ],
          "like_count":1,
          "comments":[

          ],
          "content_markdown":"Markdown"
       }
    ]
    `)
  })

  posts, err := client.Wall(1)

  if err != nil {
    t.Error("Expected no error in Wall()")
  }

  if len(posts) != 1 {
    t.Error("len(posts) != 1")
  }

  //TODO: deep test wall post
}

func TestWallAfter(t *testing.T) {
  setup()
  defer teardown()

  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    q := r.URL.Query()

    if q["after"][0] != "1" {
      t.Error("Expected WallAfter to send 'after' = 1")
    }

    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, `
    [
       {
          "id":1,
          "account_id":1,
          "staff_id":1,
          "ticket_id":null,
          "note_id":null,
          "type":"post",
          "content":"test",
          "created_at":"2013-07-12 18:08:32",
          "updated_at":"2013-07-12 21:14:30",
          "tags":[
             "tag1"
          ],
          "likes":[
             "Like 1"
          ],
          "like_count":1,
          "comments":[

          ],
          "content_markdown":"Markdown"
       }
    ]
    `)
  })

  posts, err := client.WallAfter(1, 1)

  if err != nil {
    t.Error("Expected no error in WallAfter()")
  }

  if len(posts) != 1 {
    t.Error("len(posts) != 1")
  }
}
