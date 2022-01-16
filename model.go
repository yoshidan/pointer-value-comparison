package main

import "time"

type Sample struct {
	GroupID  int64
	SampleID int64
	ItemID   int64
	Quantity int64
	Name     string
	Time     time.Time
}
