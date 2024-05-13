package testing

import (
	"log"
	"testing"

	"github.com/golang-jwt/jwt/v4"
	json_web_token "github.com/robbymaul/golang-jwt.git/jwt"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	claims := jwt.MapClaims{}
	claims["username"] = "robby"
	claims["password"] = "robby"

	token, err := json_web_token.GenerateToken(&claims)
	if err != nil {
		t.Fatal(err.Error())
	}

	log.Println(token)
	assert.NotNil(t, token)
}

func TestDecodeToken(t *testing.T) {
	s := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzd29yZCI6InJvYmJ5IiwidXNlcm5hbWUiOiJyb2JieSJ9.Ul5uAnIwtGJW13HnliWV2PBG4BzEblknU4tMJ9YPB6M"

	claims, err := json_web_token.DecodeToken(s)
	if err != nil {
		t.Fatal(err.Error())
	}

	username := claims["username"]
	password := claims["password"]

	assert.Equal(t, "robby", username)
	assert.Equal(t, "robby", password)

}
