package cloudconnect

import "encoding/json"

// DecodeTrack takes the "hash" payload and reassembles as a Track
// NOTE: Maybe some optimization needed instead of JSON Marshal/Unmarshal
func DecodeTrack(payload map[string]interface{}) (track Track, err error) {
	enc, err := json.Marshal(payload)
	if err == nil {
		json.Unmarshal(enc, &track)
	}
	return
}

// Decode takes a json string (as []byte) and decodes into Tracks.
// TODO: Return Messages and Presences
func Decode(c []byte) (tracks Tracks, err error) {
	n := make(Notification, 0)
	if err = json.Unmarshal(c, &n); err == nil {
		for _, e := range n {
			switch e.Meta.Event {
			case "track":
				if t, err := DecodeTrack(e.Payload); err == nil {
					t.NextIndex = t.Index + len(t.Fields) + 1
					tracks = append(tracks, t)
				}
			}
		}
	}

	return
}
