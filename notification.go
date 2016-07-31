package cloudconnect

//Meta is embedded in an Event. It contains the type and account associated to the Event
type Meta struct {
	Event   string `json:"event"`
	Account string `json:"account"`
}

//Event can be of multiple types. Track / Message / Presence determined by
//the Meta.Type attribute
type Event struct {
	Meta    Meta                   `json:"meta"`
	Payload map[string]interface{} `json:"payload"`
}

//A Notification is a collection of Event
type Notification []Event
