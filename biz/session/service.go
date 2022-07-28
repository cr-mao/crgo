package session

import (
	"context"
	"crgo/infra/bizerror"
	"crgo/infra/db"
	"crgo/infra/util"
	"crgo/models"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"

	"crgo/infra/redis"
)

const (
	VERSION1 int64 = iota + 1
)

func keySession(sessionID string) string {
	return "session:" + sessionID
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

// 删除缓存
func delUserSession(ctx context.Context, sid string) {
	// todo session 变 配置
	err := redis.Client("session").Del(ctx, keySession(sid)).Err()
	if err != nil && err != redis.Nil {
		panic(wrapError(err, "session redis.Del failed. SessionID: %s", sid))
	}
}

func (s *Service) IsInvalid(ctx context.Context, sessionId string) bool {
	return !s.IsValid(ctx, sessionId)
}

func (s *Service) IsValid(ctx context.Context, sessionID string) bool {
	if sessionID == "" {
		return false
	}
	session := s.Get(ctx, sessionID)
	if session != nil {
		return true
	}

	return false
}

func (s *Service) IsAnonymous(ctx context.Context, sessionID string) bool {
	session := s.Get(ctx, sessionID)
	if session != nil {
		return session.Anonymous
	}
	// 后端找不到记录，客户端应该重新走 bootstrap 流程
	return false
}

//匿名 sessionId 生成
func (s *Service) GenerateAnonymous(ctx context.Context) string {
	newSessionID := uuid.New().String()
	b := util.MustMarshal(util.MarshalWrapper(proto.Marshal), &Session{
		Version:   VERSION1,
		Anonymous: true,
	})
	err := redis.Client("session").Set(ctx, keySession(newSessionID), b, time.Hour).Err()
	if err != nil {
		panic(err)
	}
	return newSessionID
}

func (s *Service) Bind(ctx context.Context, sessionID, guid string, userID int64) {
	err := db.GetDb("default").Create(&models.UserSession{
		SessionId: sessionID,
		Guid:      guid,
		UserId:    userID,
	}).Error
	if err != nil {
		panic(wrapError(err, "session db.Insert failed. sessionID: %s, guid: %s, user_id: %d", sessionID, guid, userID))
	}

	s.Set(ctx, sessionID, guid, userID)
}
func (s *Service) Set(ctx context.Context, sessionID, guid string, userID int64) {
	b := util.MustMarshal(util.MarshalWrapper(proto.Marshal), &Session{
		Version:   VERSION1,
		Anonymous: false,
		UserID:    int64(userID),
		Guid:      guid,
	})

	err := redis.Client("session").Set(ctx, keySession(sessionID), b, time.Hour*24*7).Err()
	if err != nil {
		panic(wrapError(err, "session redis.Set failed. sessionID: %s, guid: %s, user_id: %d", sessionID, guid, userID))
	}
}

func (s *Service) Get(ctx context.Context, sessionID string) *Session {
	var session *Session
	b, err := redis.Client("session").Get(ctx, keySession(sessionID)).Bytes()
	if err != nil {
		return nil
	}
	session = &Session{}
	util.MustUnmarshal(util.UnmarshalWrapper(proto.Unmarshal), b, session)
	// reset expire time
	redis.Client("session").Expire(ctx, keySession(sessionID), time.Hour*24*7).Err()
	return session
}

func (s *Service) Remove(ctx context.Context, sessionID string) {
	user_session := &models.UserSession{
		SessionId: sessionID,
	}
	err := db.GetDb("default").Where("session_id = ? ", sessionID).Delete(user_session).Error
	if err != nil {
		panic(wrapError(err, "session db.Delete failed. SessionID: %s", sessionID))
	}
	delUserSession(ctx, sessionID)
}

func Empty() *Session {
	return &Session{
		Version:   VERSION1,
		Anonymous: true,
	}
}

func wrapError(err error, format string, args ...interface{}) error {
	return bizerror.Wrapf(1010, "服务异常", err, format, args...)
}
