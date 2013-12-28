package snappy

import (
  "bytes"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
  "net/url"
  "reflect"
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

  got, err := client.Accounts()

  if err != nil {
    t.Error("Expected no error in Accounts()")
  }

  expected := []Account{
    Account{
      ID:           3,
      Organization: "Snappy Help",
      Domain:       "help.besnappy.com",
      PlanID:       1,
      Active:       1,
      CreatedAt:    "2012-12-05 15:24:20",
      UpdatedAt:    "2013-05-07 19:48:06",
      CustomDomain: "",
    },
  }

  if reflect.DeepEqual(expected, got) == false {
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

  got, err := client.Staff(1)

  if err != nil {
    t.Error("Expected no error in Staff()")
  }

  expected := []Employee{
    Employee{
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
    },
    Employee{
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
    },
  }

  if reflect.DeepEqual(expected, got) == false {
    t.Error("expected != got")
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

  got, err := client.Mailboxes(1)

  if err != nil {
    t.Error("Expected no error in Mailboxes()")
  }

  expected := []Mailbox{
    Mailbox{
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
    },

    Mailbox{
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
    },
  }

  if reflect.DeepEqual(expected, got) == false {
    t.Error("expected != got")
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

  got, err := client.ContactByID(1, 1)

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

  if expected != got {
    t.Error("expected != got")
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

  got, err := client.ContactByEmail(1, "test@test.com")

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

  if expected != got {
    t.Error("expected != got")
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

  got, err := client.Documents(1)

  if err != nil {
    t.Error("Expected no error in Documents()")
  }

  expected := []Document{
    Document{
      ID:         1,
      AccountID:  1,
      Filename:   "Filename",
      Type:       "image/png",
      Size:       59788,
      StorageKey: "fake",
      CreatedAt:  "2013-07-10 15:41:34",
      UpdatedAt:  "2013-07-10 15:41:34",
    },
  }

  if reflect.DeepEqual(expected, got) == false {
    t.Error("expected != got")
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

  got, err := client.Wall(1)

  if err != nil {
    t.Error("Expected no error in Wall()")
  }

  expected := []WallPost{
    WallPost{
      ID:              1,
      AccountID:       1,
      StaffID:         1,
      Type:            "post",
      Content:         "test",
      CreatedAt:       "2013-07-12 18:08:32",
      UpdatedAt:       "2013-07-12 21:14:30",
      Tags:            []string{"tag1"},
      Likes:           []string{"Like 1"},
      LikeCount:       1,
      ContentMarkdown: "Markdown",
      Comments:        []WallComment{},
    },
  }

  if reflect.DeepEqual(expected, got) == false {
    t.Error("expected != got")
  }
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

  got, err := client.WallAfter(1, 1)

  if err != nil {
    t.Error("Expected no error in WallAfter()")
  }

  if len(got) != 1 {
    t.Error("len(got) != 1")
  }
}

func TestLikeWallPost(t *testing.T) {
  setup()
  defer teardown()

  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
      t.Error("Expected POST method in LikeWallPost()")
    }

    w.WriteHeader(http.StatusOK)
  })

  err := client.LikeWallPost(1, 1)

  if err != nil {
    t.Error("Expected no error in LikeWallPost()")
  }
}

func TestUnlikeWallPost(t *testing.T) {
  setup()
  defer teardown()

  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    if r.Method != "DELETE" {
      t.Error("Expected DELETE method in LikeWallPost()")
    }

    w.WriteHeader(http.StatusOK)
  })

  err := client.UnlikeWallPost(1, 1)

  if err != nil {
    t.Error("Expected no error in UnlikeWallPost()")
  }
}

func TestCommentWallPost(t *testing.T) {
  setup()
  defer teardown()

  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
      t.Error("Expected POST method in CommentWallPost()")
    }

    b, _ := ioutil.ReadAll(r.Body)
    defer r.Body.Close()

    parsedCommentFormValue, err := url.ParseQuery(string(b))

    if err != nil {
      t.Error("Expected no error parsing post data")
      return
    }

    if parsedCommentFormValue["content"][0] != "test comment" {
      t.Error("Expected POST key value content='test comment'")
    }

    w.WriteHeader(http.StatusOK)
  })

  err := client.CommentWallPost(1, 1, "test comment")

  if err != nil {
    t.Error("Expected no error in CommentWallPost()")
  }
}

func TestDeleteComment(t *testing.T) {
  setup()
  defer teardown()

  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    if r.Method != "DELETE" {
      t.Error("Expected DELETE method in DeleteComment()")
    }

    w.WriteHeader(http.StatusOK)
  })

  err := client.DeleteComment(1, 1, 1)

  if err != nil {
    t.Error("Expected no error in DeleteComment()")
  }
}

func TestCreateWallPost(t *testing.T) {
  expectedNewWallPost := NewWallPost{
    Content: "this is a test",
    Type:    "post",
    Tags:    []string{"test1", "test2"},
  }

  setup()
  defer teardown()

  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
      t.Error("Expected POST method in CreateWallPost()")
    }

    b, _ := ioutil.ReadAll(r.Body)
    defer r.Body.Close()

    gotWallPost := NewWallPost{}
    err := json.Unmarshal(b, &gotWallPost)

    if err != nil {
      t.Error("Expected no error unmarshaling json")
      return
    }

    if reflect.DeepEqual(expectedNewWallPost, gotWallPost) == false {
      t.Error("expectedNewWallPost != gotWallPost")
    }

    w.WriteHeader(http.StatusOK)
  })

  err := client.CreateWallPost(1, expectedNewWallPost)

  if err != nil {
    t.Error("Expected no error in CreateWallPost()")
  }
}

func TestDeleteWallPost(t *testing.T) {
  setup()
  defer teardown()

  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    if r.Method != "DELETE" {
      t.Error("Expected DELETE method in DeleteWallPost()")
    }
    w.WriteHeader(http.StatusOK)
  })

  err := client.DeleteWallPost(1, 1)

  if err != nil {
    t.Error("Expected no error in DeleteWallPost()")
  }
}
