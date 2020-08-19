package firebase

import (
	"time"
)

// Firebase ...
type Firebase struct {
	FirebaseID      string    `db:"FirebaseID" json:"Firebase_id"`
	FirebaseName    string    `db:"FirebaseName" json:"Firebase_name"`
	FirebaseDOB     time.Time `db:"FirebaserDOB" json:"Firebase_dob"`
	FirebaseAddress string    `db:"FirebaseAddress" json:"Firebase_address"`
	FirebaseNoHP    string    `db:"FirebaseNoHP" json:"Firebase_nohp"`
	FirebaseEmail   string    `db:"FirebaseEmail" json:"Firebase_email"`
}
