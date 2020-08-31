package protocol

type Message struct {
	Player     int
	ActionType string
	Action     string
	Data       interface{}
}

func BuildResultMessage(player int, result bool) Message {
	if result {
		return Message{Player: player, ActionType: "result", Data: "true"}
	}

	return Message{Player: player, ActionType: "result", Data: "false"}
}

func BuildActionMessage(player int, action string, data interface{}) Message {
	return Message{Player: player, ActionType: "action", Action: action, Data: data}
}

func BuildNoticeMessage(player int, action string, data interface{}) Message {
	return Message{Player: player, ActionType: "notice", Action: action, Data: data}
}
