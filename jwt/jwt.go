package jwt

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"twitter"
	"twitter/config"

	"github.com/lestrrat-go/jwx/jwa"
	jwtGo "github.com/lestrrat-go/jwx/jwt"
)

var signatureType = jwa.HS256

type TokenService struct {
	Conf *config.Config
}

func NewTokenService(conf *config.Config) *TokenService {
	return &TokenService{
		Conf: conf,
	}
}

func (ts *TokenService) ParseTokenFromRequest(c context.Context, r *http.Request) (twitter.AuthToken, error) {

	token, err := jwtGo.ParseRequest(
		r,
		jwtGo.WithValidate(true),
		jwtGo.WithIssuer(ts.Conf.JWT.Issuer),
		jwtGo.WithVerify(signatureType, []byte(ts.Conf.JWT.Secret)),
	)

	if err != nil {
		return twitter.AuthToken{}, twitter.ErrInvalidAccessToken
	}

	return buildToken(token), nil
}

func buildToken(token jwtGo.Token) twitter.AuthToken {
	return twitter.AuthToken{
		ID:  token.JwtID(),
		Sub: token.Subject(),
	}
}

func (ts *TokenService) ParseToken(c context.Context, payload string) (twitter.AuthToken, error) {
	token, err := jwtGo.Parse(
		[]byte(payload),
		jwtGo.WithValidate(true),
		jwtGo.WithIssuer(ts.Conf.JWT.Issuer),
		jwtGo.WithVerify(signatureType, []byte(ts.Conf.JWT.Secret)),
	)

	if err != nil {
		return twitter.AuthToken{}, twitter.ErrInvalidAccessToken
	}

	return buildToken(token), nil
}

func (ts *TokenService) CreateRefreshToken(c context.Context, user twitter.User, tokenId string) (string, error) {

	t := jwtGo.New()

	if err := setDefualtToken(t, user, twitter.RefreshTokenLifetime, ts.Conf); err != nil {
		return "", err
	}

	if err := t.Set(jwtGo.JwtIDKey, tokenId); err != nil {
		return "", fmt.Errorf("error set jwt id: %v", err)
	}

	token, err := jwtGo.Sign(t, signatureType, []byte(ts.Conf.JWT.Secret))
	if err != nil {
		return "", fmt.Errorf("error sign jwt : %v", err)
	}

	return string(token), nil

}

func setDefualtToken(t jwtGo.Token, user twitter.User, lifeTime time.Duration, config *config.Config) error {

	if err := t.Set(jwtGo.SubjectKey, user.ID); err != nil {
		return fmt.Errorf("error while set sub:%v", err)
	}
	if err := t.Set(jwtGo.IssuerKey, config.JWT.Issuer); err != nil {
		return fmt.Errorf("error while set issuer key: %v", err)
	}

	if err := t.Set(jwtGo.IssuedAtKey, time.Now().Unix()); err != nil {
		return fmt.Errorf("error while set issued at key: %v", err)
	}

	if err := t.Set(jwtGo.ExpirationKey, time.Now().Add(lifeTime).Unix()); err != nil {
		return fmt.Errorf("error while set jwt expired at: %v", err)
	}

	return nil
}

func (ts *TokenService) CreateAccessToken(c context.Context, user twitter.User) (string, error) {
	t := jwtGo.New()

	if err := setDefualtToken(t, user, twitter.AccessTokenLifetime, ts.Conf); err != nil {
		return "", err
	}

	token, err := jwtGo.Sign(t, signatureType, []byte(ts.Conf.JWT.Secret))
	if err != nil {
		return "", fmt.Errorf("error sign jwt : %v", err)
	}

	return string(token), nil
}
