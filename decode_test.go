package cloudconnect

import "testing"

func TestDecode(t *testing.T) {
	content := ` [{
     "meta": {
         "account": "municio",
         "event": "track"
     },
     "payload": {
         "id": 580761614244905081,
         "id_str": "580761614244905081",
         "asset": "359858024375626",
         "recorded_at": "2014-05-22T14:18:57Z",
         "recorded_at_ms": "2014-05-22T14:18:57.000Z",
         "received_at": "2014-05-22T14:19:18Z",
         "loc": [2.34793, 48.83487],
         "fields": {
             "GPS_DIR": {
                 "b64_value": "AAA+Hg=="
             },
             "GPS_SPEED": {
                 "b64_value": "AAASCw=="
             },
						 "BATT": {
                 "b64_value": "Ab=="
             }
         }
     }
 }]`

	tracks, err := Decode([]byte(content))

	if err != nil {
		t.Error(len(tracks), err)
		return
	}

	if len(tracks) != 1 {
		t.Errorf("Got %v tracks. Expected %v", len(tracks), 1)
		return
	}

	if tracks[0].Fields["GPS_SPEED"].Int() != 4619 {
		t.Errorf("Expected GPS_SPEED to equal %s. Got %s", 4619, tracks[0].Fields["GPS_SPEED"].Int())
		return
	}

	if tracks[0].Fields["BATT"].Bool() != true {
		t.Errorf("Expected BATT to equal %s. Got %s", true, tracks[0].Fields["BATT"].Bool())
		return
	}

}
