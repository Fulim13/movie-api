package data

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"time"

	"github.com/Fulim13/movie-api/internal/validator"
)

const (
	ScopeActivation = "activation"
)

type Token struct {
	PlainText string
	Hash      []byte
	UserId    int64
	Expiry    time.Time
	Scope     string
}

func generateToken(userId int64, ttl time.Duration, scope string) *Token {
	token := &Token{
		PlainText: rand.Text(),
		UserId:    userId,
		Expiry:    time.Now().Add(ttl),
		Scope:     scope,
	}

	hash := sha256.Sum256([]byte(token.PlainText))
	token.Hash = hash[:]
	return token
}

func ValidateTokenPlainText(v *validator.Validator, tokenPlaintext string) {
	v.Check(tokenPlaintext != "", "token", "must be provided")
	v.Check(len(tokenPlaintext) == 26, "token", "must be 26 bytes long")
}

type TokenModel struct {
	DB *sql.DB
}

func (t *TokenModel) New(userId int64, ttl time.Duration, scope string) (*Token, error) {
	token := generateToken(userId, ttl, scope)
	err := t.Insert(token)
	return token, err
}

func (t *TokenModel) Insert(token *Token) error {
	query := `
		INSERT INTO tokens (hash, user_id, expiry, scope)
		VALUES ($1, $2, $3, $4)
	`

	args := []any{
		token.Hash,
		token.UserId,
		token.Expiry,
		token.Scope,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := t.DB.ExecContext(ctx, query, args...)
	return err
}

func (t *TokenModel) DeleteAllForUser(scope string, userId int64) error {
	query := `
		DELETE FROM tokens
		WHERE scope = $1 AND user_id = $2
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err := t.DB.ExecContext(ctx, query, scope, userId)
	return err
}
