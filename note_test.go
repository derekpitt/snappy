package snappy

import (
  "encoding/json"
  "io/ioutil"
  "net/http"
  "reflect"
  "testing"
)

func TestCreateNote(t *testing.T) {
  expectedNote := NewNote{
    Subject:   "test from api",
    MailboxID: 1,
    TicketID:  1,
    StaffID:   1,
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

    gotNote := NewNote{}
    err := json.Unmarshal(b, &gotNote)

    if err != nil {
      t.Error("Expected no error in unmarshaling note")
      return
    }

    if reflect.DeepEqual(expectedNote, gotNote) == false {
      t.Error("expectedNote != gotNote")
      return
    }

    w.WriteHeader(http.StatusOK)
  })

  err := client.CreateNote(expectedNote)

  if err != nil {
    t.Error("Expected no error in CreateNote()")
  }

}
