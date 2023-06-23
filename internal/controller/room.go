package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/you-together/internal/controller/dao"
	"github.com/wtkeqrf0/you-together/internal/controller/dto"
	"github.com/wtkeqrf0/you-together/pkg/log"
	"github.com/wtkeqrf0/you-together/pkg/middleware/errs"
	"net/http"
)

func (h *Handler) createRoom(c *gin.Context, r dto.Room, info *dao.Session) error {

	log.Infof("websocket connection: %v", c.IsWebsocket())

	err := h.room.UpsertRoom(r, info.ID)
	if err != nil {
		return err
	}

	/*OTP := c.Query("otp")
	if OTP == "" {
		c.Status(http.StatusUnauthorized)
		return nil
	}

	if !m.otps.verifyOTP(OTP) {
		c.Status(http.StatusUnauthorized)
		return nil
	}*/

	if err = h.ws.Connect(c.Writer, c.Request); err != nil {
		//TODO handle
		return errs.WebsocketNotFound.AddErr(err)
	}

	c.Status(http.StatusOK)
	return nil
}

func (h *Handler) joinRoom(c *gin.Context) error {

	log.Infof("websocket connection: %v", c.IsWebsocket())

	if err := h.ws.Connect(c.Writer, c.Request); err != nil {
		//TODO handle
		return errs.WebsocketNotFound.AddErr(err)
	}
	return nil
}
