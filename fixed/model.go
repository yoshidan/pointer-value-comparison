package fixed

type Sample struct {
	GroupID  int64
	SampleID int64
	ItemID   int64
	Quantity int64
}

//go:noinline
func FilterValue(sl []Sample) []Sample {
	v := make([]Sample, 0, len(sl))
	for _, e := range sl {
		if e.Quantity < 10000 {
			v = append(v, e)
		}
	}
	return v
}

//go:noinline
func FilterPtr(sl []*Sample) []*Sample {
	v := make([]*Sample, 0, len(sl))
	for _, e := range sl {
		if e.Quantity < 10000 {
			v = append(v, e)
		}
	}
	return v
}

//go:noinline
func JustReturnValue(v []Sample) []Sample {
	return v
}

//go:noinline
func JustReturnPtr(v []*Sample) []*Sample {
	return v
}
