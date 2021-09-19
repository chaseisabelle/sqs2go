package sqs2go

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/chaseisabelle/sqsc"
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
	sum := md5.Sum([]byte(msg))

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(fmt.Sprintf(`
			<ReceiveMessageResponse>
			  <ReceiveMessageResult>
				<Message>
				  <MessageId>5fea7756-0ea4-451a-a703-a558b933e274</MessageId>
				  <ReceiptHandle>
					MbZj6wDWli+JvwwJaBV+3dcjk2YW2vA3+STFFljTM8tJJg6HRG6PYSasuWXPJB+Cw
					Lj1FjgXUv1uSj1gUPAWV66FU/WeR4mq2OKpEGYWbnLmpRCJVAyeMjeU5ZBdtcQ+QE
					auMZc8ZRv37sIW2iJKq3M9MFx1YvV11A2x/KSbkJ0=
				  </ReceiptHandle>
				  <MD5OfBody>%s</MD5OfBody>
				  <Body>%s</Body>
				  <Attribute>
					<Name>SenderId</Name>
					<Value>195004372649</Value>
				  </Attribute>
				  <Attribute>
					<Name>SentTimestamp</Name>
					<Value>1238099229000</Value>
				  </Attribute>
				  <Attribute>
					<Name>ApproximateReceiveCount</Name>
					<Value>5</Value>
				  </Attribute>
				  <Attribute>
					<Name>ApproximateFirstReceiveTimestamp</Name>
					<Value>1250700979248</Value>
				  </Attribute>
				</Message>
			  </ReceiveMessageResult>
			  <ResponseMetadata>
				<RequestId>b6633655-283d-45b4-aee4-4e84e0ae6afa</RequestId>
			  </ResponseMetadata>
			</ReceiveMessageResponse>
		`, fmt.Sprintf("%x", sum), msg)))

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
