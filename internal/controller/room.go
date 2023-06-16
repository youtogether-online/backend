package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wtkeqrf0/you-together/internal/controller/dao"
	"github.com/wtkeqrf0/you-together/internal/controller/dto"
	"net/http"
)

func (h *Handler) createRoom(c *gin.Context, r dto.Room, info *dao.Session) error {

	room, err := h.room.Create(r, info.ID)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, room)
	return nil
}
