package client

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/PabloGamiz/SafeEvents-Backend/model/session"
	"github.com/joho/godotenv"
	"google.golang.org/api/oauth2/v2"
)

func newTestTokenInfo() *oauth2.Tokeninfo {
	return &oauth2.Tokeninfo{
		Email:     "testing@gmail.com",
		ExpiresIn: 1,
		UserId:    "1234",
	}
}

func TestPostcondition(t *testing.T) {
	if err := godotenv.Load("../../testing.env"); err != nil {
		t.Fatalf("Got error %s; while loading dotenv", err.Error())
	}

	subject := &txSignin{
		info: newTestTokenInfo(),
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	if _, err := subject.Postcondition(ctx); err != nil {
		t.Fatalf("Got error %s; while executing Postcondition", err.Error())
	}

	session.AllInstancesByEmail.Range(func(key interface{}, value interface{}) bool {
		log.Printf("Got session for email %s", key)
		return true
	})

	if _, err := session.GetSessionByEmail(subject.info.Email); err != nil {
		t.Fatalf("Got error %s, while getting session by email %s", err.Error(), subject.info.Email)
	}
}
