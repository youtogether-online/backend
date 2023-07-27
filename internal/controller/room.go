package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/you-together/internal/controller/dao"
	"github.com/wtkeqrf0/you-together/internal/controller/dto"
	"github.com/wtkeqrf0/you-together/pkg/middleware/errs"
	"net/http"
)

func (h *Handler) createRoom(c *gin.Context, r dto.Room, info *dao.Session) error {
	roomId, err := h.room.UpsertRoom(r, info.ID)
	if err != nil {
		return err
	}

	if err = h.ws.Connect(c, roomId); err != nil {
		return errs.WebsocketNotFound.AddErr(err)
	}

	c.Status(http.StatusOK)
	return nil
}

func (h *Handler) joinRoom(c *gin.Context, r dto.RoomId) error {

	room, err := h.room.GetRoomById(r.Id)
	if err != nil {
		return err
	}

	if err = h.ws.Connect(c, r.Id); err != nil {
		return errs.WebsocketNotFound.AddErr(err)
	}

	c.JSON(http.StatusOK, room)
	return nil
}
