package api

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"slices"
	"time"

	"github.com/FADAMIS/dashboard/db"
	"github.com/FADAMIS/dashboard/entities"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Login(ctx *gin.Context) {
	var credentials entities.Admin
	ctx.Bind(&credentials)

	sum := sha256.Sum256([]byte(credentials.Password))
	credentials.Password = hex.EncodeToString(sum[:])

	admins := db.GetAdmins()
	adminIndex := slices.IndexFunc(admins, func(a entities.Admin) bool {
		return credentials.Username == a.Username && credentials.Password == a.Password
	})

	if adminIndex == -1 {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "login failed",
		})

		return
	}

	session := newSession(credentials.Username)
	jsonSession, _ := json.Marshal(session)
	// CHANGE DOMAIN IN THE COOKIE
	ctx.SetCookie("session", string(jsonSession), 6*60*60, "/", "fadamis.live", false, true)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "login successful",
	})
}

func newSession(username string) entities.Session {
	var session entities.Session

	session.UUID = uuid.New().String()
	session.Username = username
	session.Expires = time.Now().Unix() + 6*60*60

	db.AddSession(session)

	return session
}

// Checks if a session is valid and deletes the session if it is expired
func IsSessionValid(session entities.Session) bool {
	sessions := db.GetSessions()
	sessionIndex := slices.IndexFunc(sessions, func(s entities.Session) bool { return s == session })

	if sessionIndex == -1 || session.Expires < time.Now().Unix() {
		return false
	}

	return true
}
