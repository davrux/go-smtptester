package main

import (
	"fmt"
	"log"
	"net/smtp"
	"time"

	"github.com/davrux/go-smtptester"
)

func main() {
	s := smtptester.Standard()

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("smtp server response %s", err)
		}
	}()

	defer s.Close()

	// Wait a second to let the server come up.
	time.Sleep(time.Second)

	// Send email.
	from := "alice@i.com"
	to := []string{"bob@e.com", "mal@b.com"}
	msg := []byte("Test\r\n")
	if err := smtp.SendMail(s.Addr, nil, from, to, msg); err != nil {
		log.Fatalf("error sending mail %+v", err)
	}

	// Lookup email.
	m, found := smtptester.GetBackend(s).Load(from, to)
	fmt.Printf("Found %t, mail %+v\n", found, m)
}
