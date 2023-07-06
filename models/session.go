package models

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"fmt"

	"github.com/jimmyhogerty/lenslocked/rand"
)

type Session struct {
  ID int
  UserID int
  // Token is only set when creating a new session. When looking up a 
  // session, this will be left empty as we only store the hash of a token in 
  // our DB and we cannot reverse it into a raw token, similar to passwords.
  Token string
  TokenHash string
}

// INFO This is the minimum number of bytes to be used for each session token. Can be increased if needed.
const (
  MinBytesPerToken = 32
)

type SessionService struct {
  DB *sql.DB
  // BytesPerToken is used to determine how many bytes each session token should be generated with.
  BytesPerToken int
}


// Writes a new session to the database.
func (ss *SessionService) Create(userID int) (*Session, error) {
	bytesPerToken := ss.BytesPerToken
	if bytesPerToken < MinBytesPerToken {
		bytesPerToken = MinBytesPerToken
	}
	token, err := rand.String(bytesPerToken)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	session := Session{
		UserID:    userID,
		Token:     token,
		TokenHash: ss.hash(token),
	}
	row := ss.DB.QueryRow(`
		INSERT INTO sessions (user_id, token_hash)
    VALUES ($1, $2) ON CONFLICT (user_id) DO
    UPDATE
    SET token_hash = $2
    RETURNING id;`, session.UserID, session.TokenHash)
	err = row.Scan(&session.ID)
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}
	return &session, nil
}

// This takes a session token and returns either a User (if found) or an error.
func (ss *SessionService) User(token string) (*User, error) {
  tokenHash := ss.hash(token)
  var user User
  row := ss.DB.QueryRow(`
    SELECT users.id,
    users.email,
    users.password_hash
    FROM sessions
      JOIN users ON users.id = sessions.user_id
    WHERE sessions.token_hash = $1`, tokenHash)
  err := row.Scan(&user.ID, &user.Email, &user.PasswordHash)
  if err != nil {
    return nil, fmt.Errorf("user: %w", err)
  }
  return &user, nil;
}

// Takes a token hash and deletes corresponding session.
func (ss *SessionService) Delete(token string) error {
  tokenHash := ss.hash(token)
  _, err := ss.DB.Exec(`
    DELETE FROM sessions
    WHERE token_hash = $1`, tokenHash)
    if err!= nil {
      return fmt.Errorf("delete: %w", err)
    }
    return nil
}

// Takes a session token and returns a hash of the token.
func (ss *SessionService) hash(token string) string {
  tokenHash := sha256.Sum256([]byte(token))
  // INFO [:] turns an array into a slice, where you want all the elements of the array.
  return base64.URLEncoding.EncodeToString(tokenHash[:])
}
