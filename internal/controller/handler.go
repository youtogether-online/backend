package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/you-together/ent"
	"github.com/wtkeqrf0/you-together/internal/controller/dao"
	"github.com/wtkeqrf0/you-together/internal/controller/dto"
	"github.com/wtkeqrf0/you-together/pkg/conf"
	"mime/multipart"
)

type UserService interface {
	FindUserByUsername(username string) (*dao.User, error)
	FindUserByID(id int) (*ent.User, error)
	FindMe(id int) (*dao.Me, error)
	GetRoomByOwner(userId int) (*ent.Room, error)

	UpdateUser(customer dto.UpdateUser, id int) error
	UpdatePassword(newPassword []byte, id int) error
	UpdateEmail(email string, id int) error
	UpdateImage(imageName string, id int) error
	UpdateUsername(username string, id int) error
	UsernameExist(username string) (bool, error)
}

type RoomService interface {
	UpsertRoom(rm dto.Room, creatorId int) (int, error)
	GetRoomById(roomId int) (*ent.Room, error)
}

type AuthService interface {
	SetCodes(key string, value ...any) error
	EqualsPopCode(email string, code string) (bool, error)
	DelKeys(keys ...string)
	CompareHashAndPassword(old, new []byte) error

	CreateUserByEmail(email string, language string) (*ent.User, error)
	AuthUserByEmail(email string) (*ent.User, error)
	SetEmailVerified(email string) error
	FormatLanguage(header string) string
}

type MailSender interface {
	DialAndSend(subj, body string, to ...string) error
}

type Session interface {
	SetNewCookie(id int, c *gin.Context) error
	ValidateSession(sessionId string) (info *dao.Session, ok bool, err error)
	GenerateSecretCode() string
	GenerateFileName(c *gin.Context, file *multipart.FileHeader) (string, error)
}

type WebSocket interface {
	Connect(c *gin.Context, roomId int) error
}

type Handler struct {
	user UserService
	room RoomService
	auth AuthService
	mail MailSender
	sess Session
	ws   WebSocket
	cfg  *conf.Config
}

func NewHandler(user UserService, room RoomService, auth AuthService, mail MailSender, sess Session, ws WebSocket, cfg *conf.Config) *Handler {
	return &Handler{user: user, room: room, auth: auth, mail: mail, sess: sess, ws: ws, cfg: cfg}
}
