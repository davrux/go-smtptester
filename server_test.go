package smtptester

import (
	"log"
	"os"
	"testing"
	"time"

	s "net/smtp"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	srv = Standard()
)

func TestMain(m *testing.M) {
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("smtp server response %s", err)
		}
	}()

	// Wait sec to let server come up.
	time.Sleep(time.Second)

	exitVal := m.Run()

	srv.Close()

	os.Exit(exitVal)
}

func TestBackend_AddLoad(t *testing.T) {
	b := Backend{}

	m := &Mail{
		From:       "alice@i.com",
		Recipients: []string{"bob@e.com"},
		Data:       []byte("test"),
	}

	b.Add(m)
	m1, found := b.Load(m.From, m.Recipients)
	assert.True(t, found)
	assert.Equal(t, m, m1)
}

func TestSendMail(t *testing.T) {
	from := "alice@i.com"
	recipients := []string{"bob@e.com"}
	data := []byte("Test mail\r\n")

	// Send without TLS.
	require.Nil(t, s.SendMail(srv.Addr, nil, from, recipients, data))

	m, found := GetBackend(srv).Load(from, recipients)
	assert.True(t, found)
	assert.Equal(t, from, m.From)
	assert.Equal(t, recipients, m.Recipients)
	assert.Equal(t, data, m.Data)
}
