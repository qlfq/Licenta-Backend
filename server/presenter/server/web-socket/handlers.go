package web_socket

func (m *Manager) SetupEventHandlers() {
	m.Handlers[EventSendMessage] = SendMessageHandler
	m.Handlers[EventSendDmMessage] = SendDmMessageHandler
	m.Handlers[EventSendNewRoom] = AddRoomHandler
	m.Handlers[EventDeleteRoom] = DeleteRoomHandler
}
