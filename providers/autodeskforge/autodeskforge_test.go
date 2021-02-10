package autodeskforge_test

import (
	"os"
	"testing"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/autodeskforge"
	"github.com/stretchr/testify/assert"
)

func Test_New(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	p := provider()

	a.Equal(p.ClientKey, os.Getenv("ADSK_FORGE_CLIENT_ID"))
	a.Equal(p.Secret, os.Getenv("ADSK_FORGE_CLIENT_SECRET"))
	a.Equal(p.CallbackURL, "/foo")
}

func Test_Implements_Provider(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	a.Implements((*goth.Provider)(nil), provider())
}

func Test_BeginAuth(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	p := provider()
	session, err := p.BeginAuth("test_state")
	s := session.(*autodeskforge.Session)
	a.NoError(err)
	a.Contains(s.AuthURL, "https://developer.api.autodesk.com/authentication/v1/authorize")
}

func Test_SessionFromJSON(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	p := provider()
	session, err := p.UnmarshalSession(`{"AuthURL":"https://developer.api.autodesk.com/authentication/v1/authorize","AccessToken":"1234567890"}`)
	a.NoError(err)

	s := session.(*autodeskforge.Session)
	a.Equal(s.AuthURL, "https://developer.api.autodesk.com/authentication/v1/authorize")
	a.Equal(s.AccessToken, "1234567890")
}

func provider() *autodeskforge.Provider {
	return autodeskforge.New(os.Getenv("ADSK_FORGE_CLIENT_ID"), os.Getenv("ADSK_FORGE_CLIENT_SECRET"), "/foo")
}
