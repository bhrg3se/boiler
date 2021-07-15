package chat

import (
	"boiler/models"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

func Join(ws *websocket.Conn, user *models.User) {
	var msg struct {
		Data string `json:"data"`
	}
	for ws.ReadJSON(&msg) == nil {
		msgReply := struct {
			Data string `json:"data"`
		}{"hello " + msg.Data}
		err := ws.WriteJSON(msgReply)
		if err != nil {
			logrus.Debug("probably disconnected: ", err)
			return
		}
	}
}
