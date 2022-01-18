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

func FilterValue(sl []Sample) []Sample {
	v := make([]Sample, 0, len(sl))
	for _, e := range sl {
		if e.Quantity < 10000 {
			v = append(v, e)
		}
	}
	return v
}

func FilterPtr(sl []*Sample) []*Sample {
	v := make([]*Sample, 0, len(sl))
	for _, e := range sl {
		if e.Quantity < 10000 {
			v = append(v, e)
		}
	}
	return v
}
