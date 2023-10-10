package session

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tot0p/SharePhoto/model"
)

// init initializes the sessions map
func init() {
	SessionsManager.Sessions = make(map[string]*session)
}

// SessionsManager is the sessions manager
var SessionsManager sessionsManager

// sessionsManager is the sessions manager
type sessionsManager struct {
	Sessions map[string]*session
}

// session is a session
type session struct {
	uuid string
	user *model.User
}

// DeleteSession deletes the current session
func (s *sessionsManager) DeleteSession(ctx *gin.Context) {
	cookie, err := ctx.Cookie("session")
	if err != nil {
		return
	}
	s.RemoveSession(cookie)
	domain := ctx.Request.Host
	ctx.SetCookie("session", "", -1, "/", domain, false, false)
}

// IsLogged returns true if the user is logged
func (s *sessionsManager) IsLogged(ctx *gin.Context) bool {
	cookie, err := ctx.Cookie("session")
	if err != nil {
		return false
	}
	_, ok := s.Sessions[cookie]
	return ok
}

// GetUser returns the user of the current session and nil if the user is not logged
func (s *sessionsManager) GetUser(ctx *gin.Context) *model.User {
	cookie, err := ctx.Cookie("session")
	if err != nil {
		return nil
	}
	user, ok := s.Sessions[cookie]
	if !ok {
		return nil
	}
	return user.GetUser()
}

// CreateSession creates a new session for the given user and set the cookie in the context
func (s *sessionsManager) CreateSession(ctx *gin.Context, user *model.User) {
	ses := s.addSession(user)
	domain := ctx.Request.Host
	ctx.SetCookie("session", ses.GetUUID(), 36000, "/", domain, false, false)
}

// addSession adds a new session for the given user
func (s *sessionsManager) addSession(user *model.User) *session {
	u := uuid.New().String()
	var ses = &session{
		uuid: u,
		user: user,
	}
	s.Sessions[u] = ses
	return ses
}

// RemoveSession removes the session with the given uuid
func (s *sessionsManager) RemoveSession(uuid string) {
	delete(s.Sessions, uuid)
}

// GetUser returns the user of the current session
func (s *session) GetUser() *model.User {
	return s.user
}

// GetUUID returns the uuid of the current session
func (s *session) GetUUID() string {
	return s.uuid
}
