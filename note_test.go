package snappy

import (
  "encoding/json"
  "io/ioutil"
  "net/http"
  "reflect"
  "testing"
)

func TestCreateNote(t *testing.T) {
  expected := NewNote{
    Subject:     "test from api",
    MailboxID:   1,
    TicketNonce: "123",
    StaffID:     1,
    To: []NoteAddress{
      NoteAddress{
        Name:    "Test 1",
        Address: "test@test.com",
      },
    },
    Message: "test message",
  }

  setup()
  defer teardown()

  mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
      t.Error("Expected POST method in CreateNote()")
    }

    b, _ := ioutil.ReadAll(r.Body)
    defer r.Body.Close()

    got := NewNote{}
    err := json.Unmarshal(b, &got)

    if err != nil {
      t.Error("Expected no error in unmarshaling note")
      return
    }

    if reflect.DeepEqual(expected, got) == false {
      t.Error("expected != got")
      return
    }

    w.WriteHeader(http.StatusOK)
  })

  err := client.CreateNote(expected)

  if err != nil {
    t.Error("Expected no error in CreateNote()")
  }

}
