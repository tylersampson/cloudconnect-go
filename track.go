package cloudconnect

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/binary"
	"fmt"
	"sort"
	"time"
)

// Field is a wrapper around a base64 value on a Track
type Field struct {
	Base64Value string `json:"b64_value"`
}

// Fields maps a name (string) to a Field
type Fields map[string]Field

// Int converts the base64 value into a 32-bit integer
func (f Field) Int() (val int32) {
	b, _ := b64.StdEncoding.DecodeString(f.Base64Value)
	buf := bytes.NewBuffer(b)
	binary.Read(buf, binary.BigEndian, &val)
	return
}

// String directly decodes the base64 value into a string
func (f Field) String() (val string) {
	b, _ := b64.StdEncoding.DecodeString(f.Base64Value)
	val = string(b)
	return
}

// Bool converts the base64 value into a boolean
func (f Field) Bool() (val bool) {
	b, _ := b64.StdEncoding.DecodeString(f.Base64Value)
	val = b[0] == 1
	return
}

// Track represents tracking data received from a device.
type Track struct {
	ID    uint64 `json:"id"`
	Asset string `json:"asset"`
	IDStr string `json:"id_str"`

	ConnectionID uint64 `json:"connection_id"`
	Index        int    `json:"index"`
	NextIndex    int    `json:"next_index"`

	RecordedAt time.Time `json:"recorded_at"`
	ReceivedAt time.Time `json:"received_at"`

	Loc []float64 `json:"loc"`

	Fields Fields `json:"fields"`
}

//SortKey used to sort tracks by Asset, then ID (better than Recorded at)
func (t Track) SortKey() string {
	return fmt.Sprintf("%s_%v_%v", t.Asset, t.ConnectionID, t.Index)
}

// Tracks is a collection of Track
type Tracks []Track

func (ts Tracks) Len() int {
	return len(ts)
}

func (ts Tracks) Less(i, j int) bool {
	return ts[i].SortKey() < ts[j].SortKey()
}

func (ts Tracks) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}

func (ts Tracks) GroupByAsset() (tracksPerAsset map[string]Tracks, assets []string) {
	if len(ts) > 0 {
		sort.Sort(ts)

		tracksPerAsset = make(map[string]Tracks, 0)
		assets = make([]string, 0)
		lastAsset := ts[0].Asset
		lastIndex := 0

		for i, t := range ts {
			if t.Asset != lastAsset {
				tracksPerAsset[lastAsset] = ts[lastIndex:i]
				assets = append(assets, lastAsset)
				lastAsset = t.Asset
				lastIndex = i
			}
		}

		//Add remaining tracks
		tracksPerAsset[lastAsset] = ts[lastIndex:]
		assets = append(assets, lastAsset)
	}
	return
}
