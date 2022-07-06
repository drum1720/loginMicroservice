package models

import (
	"math"
	"time"
)

type TimeLogEntry struct {
	Id        int64
	UserId    int64
	StartTime time.Time
	EndTime   time.Time
	IsEnded   bool
	Minutes   int
}

func (t *TimeLogEntry) CalculateMinutes() {
	t.Minutes = int(math.Round(t.EndTime.UTC().Sub(t.StartTime.UTC()).Minutes()))
}
