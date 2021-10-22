package sqs2go

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/chaseisabelle/sqsc"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNew_Logger_Success(t *testing.T) {
	exp := "test"
	act := ""

	han := func(_ string) error {
		return nil
	}

	lgr := func(e error) {
		act = e.Error()
	}

	s2g, err := New(han, lgr)

	if err != nil {
		t.Fatal(err)
	}

	lgr = s2g.Logger()

	lgr(errors.New("test"))

	if act != exp {
		t.Errorf("expected %s, got %s", exp, act)
	}
}

func TestNew_ConfigureWorkers_Success(t *testing.T) {
	exp := 1

	han := func(_ string) error {
		return nil
	}

	s2g, err := New(han, nil)

	if err != nil {
		t.Fatal(err)
	}

	err = s2g.Configure(nil)

	if err != nil {
		t.Fatal(err)
	}

	act := s2g.Configuration().Workers

	if act != exp {
		t.Errorf("expected %d, got %d", exp, act)
	}
}

func TestNew_ConfigureSQSC_Success(t *testing.T) {
	exp := &sqsc.Config{
		ID:       "poop",
		Key:      "fart",
		Secret:   "plop",
		Region:   "peepee",
		Queue:    "poopoo",
		URL:      "weewee",
		Endpoint: "dungle",
		Retries:  3,
		Timeout:  10,
		Wait:     5,
	}

	cfg := &Configuration{
		Workers: 1,
		SQSC:    exp,
	}

	han := func(_ string) error {
		return nil
	}

	s2g, err := New(han, nil)

	if err != nil {
		t.Fatal(err)
	}

	err = s2g.Configure(cfg)

	if err != nil {
		t.Fatal(err)
	}

	act := s2g.Configuration().SQSC

	if act != exp {
		t.Errorf("expected %+v, got %+v", exp, act)
	}
}

func TestNew_Start_Success(t *testing.T) {
	msg := "im watching svu right now"
	pl, err := payload(msg)

	if err != nil {
		t.Error(err)
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(pl))

		if err != nil {
			t.Error(err)
		}
	}))

	cfg := &Configuration{
		Workers: 1,
		SQSC: &sqsc.Config{
			ID:       "poop",
			Key:      "fart",
			Secret:   "plop",
			Region:   "peepee",
			Queue:    "poopoo",
			URL:      srv.URL,
			Endpoint: srv.URL,
			Retries:  3,
			Timeout:  10,
			Wait:     5,
		},
	}

	chn := make(chan string)

	han := func(bod string) error {
		chn <- bod

		return nil
	}

	s2g, err := New(han, nil)

	if err != nil {
		t.Fatal(err)
	}

	err = s2g.Configure(cfg)

	if err != nil {
		t.Fatal(err)
	}

	go func() {
		err = s2g.Start()

		if err != nil {
			t.Error(err)
		}
	}()

	act := <-chn

	if act != msg {
		t.Errorf("expected %s, got %s", msg, act)
	}
}

func payload(bod string) (string, error) {
	bpl, err := ioutil.ReadFile("payload.xml")

	if err != nil {
		return "", err
	}

	return fmt.Sprintf(string(bpl), fmt.Sprintf("%x", md5.Sum([]byte(bod))), bod), nil
}
