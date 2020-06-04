package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)


// JSONTime format json time field by myself
type jsonTime struct {
	time.Time
}

// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (t jsonTime) MarshalJSON() ([]byte, error) {
	if (t == jsonTime{}) {
		formatted := fmt.Sprintf("\"%s\"", "")
		return []byte(formatted), nil
	} else {
		formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
		return []byte(formatted), nil
	}
}

// Value insert timestamp into mysql need this function.
func (t jsonTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *jsonTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = jsonTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

type TimeModel struct {
	CreatedAt jsonTime `json:"created_at"`
	UpdatedAt jsonTime `json:"updated_at"`
}

type DeletedTimeModel struct {
	DeletedAt jsonTime `json:"deleted_at"`
}