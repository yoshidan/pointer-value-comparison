package main

import (
	"errors"
	"time"
)

// Constant Master Data Repository

type ValueSampleRepository struct {
}

func NewValueSampleRepository() *ValueSampleRepository {
	return &ValueSampleRepository{}
}

//go:noinline
func (r *ValueSampleRepository) FindByPK(groupID int64, sampleID int64) (Sample, error) {
	switch [2]int64{groupID, sampleID} {
	case [2]int64{1, 1}:
		return Sample{1, 1, 1, 3, "abcdefghijk", time.Now()}, nil
	case [2]int64{1, 2}:
		return Sample{1, 1, 1, 3, "abcdefghijk", time.Now()}, nil
	case [2]int64{1, 3}:
		return Sample{1, 1, 1, 3, "abcdefghijk", time.Now()}, nil
	case [2]int64{1, 4}:
		return Sample{1, 1, 1, 3, "abcdefghijk", time.Now()}, nil
	case [2]int64{1, 5}:
		return Sample{1, 1, 1, 3, "abcdefghijk", time.Now()}, nil
	case [2]int64{1, 6}:
		return Sample{1, 1, 1, 3, "abcdefghijk", time.Now()}, nil
	case [2]int64{1, 7}:
		return Sample{1, 1, 1, 3, "abcdefghijk", time.Now()}, nil
	case [2]int64{1, 8}:
		return Sample{1, 1, 1, 3, "abcdefghijk", time.Now()}, nil
	case [2]int64{1, 9}:
		return Sample{1, 1, 1, 3, "abcdefghijk", time.Now()}, nil
	case [2]int64{1, 10}:
		return Sample{1, 1, 1, 3, "abcdefghijk", time.Now()}, nil
	case [2]int64{1, 11}:
		return Sample{1, 1, 1, 3, "abcdefghijk", time.Now()}, nil
	case [2]int64{1, 12}:
		return Sample{1, 1, 1, 3, "abcdefghijk", time.Now()}, nil
	case [2]int64{1, 13}:
		return Sample{1, 1, 1, 3, "abcdefghijk", time.Now()}, nil
	case [2]int64{1, 14}:
		return Sample{1, 1, 1, 3, "abcdefghijk", time.Now()}, nil
	case [2]int64{1, 15}:
		return Sample{1, 1, 1, 3, "abcdefghijk", time.Now()}, nil
	case [2]int64{1, 16}:
		return Sample{1, 1, 1, 3, "abcdefghijk", time.Now()}, nil
	case [2]int64{1, 17}:
		return Sample{1, 1, 1, 3, "abcdefghijk", time.Now()}, nil
	case [2]int64{1, 18}:
		return Sample{1, 1, 1, 3, "abcdefghijk", time.Now()}, nil
	}
	return Sample{}, errors.New("invalid Sample ")
}

//go:noinline
func (r *ValueSampleRepository) FindBySampleId(groupID int64) ([]Sample, error) {
	switch [1]int64{groupID} {
	case [1]int64{1}:
		return []Sample{
			{1, 1, 1, 3, "abcdegshijk", time.Now()},
			{1, 2, 2, 3, "abcdegshijk", time.Now()},
			{1, 3, 2, 3, "abcdegshijk", time.Now()},
			{1, 4, 2, 3, "abcdegshijk", time.Now()},
			{1, 5, 2, 3, "abcdegshijk", time.Now()},
			{1, 6, 2, 3, "abcdegshijk", time.Now()},
			{1, 7, 2, 3, "abcdegshijk", time.Now()},
			{1, 8, 2, 3, "abcdegshijk", time.Now()},
			{1, 9, 2, 3, "abcdegshijk", time.Now()},
			{1, 10, 2, 3, "abcdegshijk", time.Now()},
			{1, 11, 2, 3, "abcdegshijk", time.Now()},
			{1, 12, 2, 3, "abcdegshijk", time.Now()},
			{1, 13, 2, 3, "abcdegshijk", time.Now()},
			{1, 14, 2, 3, "abcdegshijk", time.Now()},
			{1, 15, 2, 3, "abcdegshijk", time.Now()},
			{1, 16, 2, 3, "abcdegshijk", time.Now()},
			{1, 17, 2, 3, "abcdegshijk", time.Now()},
			{1, 18, 2, 3, "abcdegshijk", time.Now()},
			{1, 19, 2, 3, "abcdegshijk", time.Now()},
			{1, 20, 2, 3, "abcdegshijk", time.Now()},
			{1, 21, 2, 3, "abcdegshijk", time.Now()},
			{1, 22, 2, 3, "abcdegshijk", time.Now()},
			{1, 23, 2, 3, "abcdegshijk", time.Now()},
			{1, 24, 2, 3, "abcdegshijk", time.Now()},
			{1, 25, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
		}, nil
	}
	return nil, errors.New("error")
}

type PtrSampleRepository struct {
}

func NewPtrSampleRepository() *PtrSampleRepository {
	return &PtrSampleRepository{}
}

//go:noinline
func (r *PtrSampleRepository) FindByPK(groupID int64, sampleID int64) (*Sample, error) {
	switch [2]int64{groupID, sampleID} {
	case [2]int64{1, 1}:
		return &Sample{1, 1, 1, 3, "abcdegshijk", time.Now()}, nil
	case [2]int64{1, 2}:
		return &Sample{1, 2, 2, 3, "abcdegshijk", time.Now()}, nil
	case [2]int64{1, 3}:
		return &Sample{1, 3, 3, 3, "abcdegshijk", time.Now()}, nil
	case [2]int64{1, 4}:
		return &Sample{1, 4, 4, 3, "abcdegshijk", time.Now()}, nil
	case [2]int64{1, 5}:
		return &Sample{1, 5, 5, 3, "abcdegshijk", time.Now()}, nil
	case [2]int64{1, 6}:
		return &Sample{1, 6, 6, 3, "abcdegshijk", time.Now()}, nil
	case [2]int64{1, 7}:
		return &Sample{1, 7, 7, 3, "abcdegshijk", time.Now()}, nil
	case [2]int64{1, 8}:
		return &Sample{1, 8, 8, 3, "abcdegshijk", time.Now()}, nil
	case [2]int64{1, 9}:
		return &Sample{1, 9, 9, 3, "abcdegshijk", time.Now()}, nil
	case [2]int64{1, 10}:
		return &Sample{1, 10, 10, 3, "abcdegshijk", time.Now()}, nil
	case [2]int64{1, 11}:
		return &Sample{1, 10, 11, 3, "abcdegshijk", time.Now()}, nil
	case [2]int64{1, 12}:
		return &Sample{1, 10, 11, 3, "abcdegshijk", time.Now()}, nil
	case [2]int64{1, 13}:
		return &Sample{1, 10, 11, 3, "abcdegshijk", time.Now()}, nil
	case [2]int64{1, 14}:
		return &Sample{1, 10, 11, 3, "abcdegshijk", time.Now()}, nil
	case [2]int64{1, 15}:
		return &Sample{1, 10, 11, 3, "abcdegshijk", time.Now()}, nil
	case [2]int64{1, 16}:
		return &Sample{1, 10, 11, 3, "abcdegshijk", time.Now()}, nil
	case [2]int64{1, 17}:
		return &Sample{1, 10, 11, 3, "abcdegshijk", time.Now()}, nil
	case [2]int64{1, 18}:
		return &Sample{1, 10, 11, 3, "abcdegshijk", time.Now()}, nil
	}
	return &Sample{}, errors.New("invalid Sample ")
}

//go:noinline
func (r *PtrSampleRepository) FindBySampleId(groupID int64) ([]*Sample, error) {
	switch [1]int64{groupID} {
	case [1]int64{1}:
		return []*Sample{
			{1, 1, 1, 3, "abcdegshijk", time.Now()},
			{1, 2, 2, 3, "abcdegshijk", time.Now()},
			{1, 3, 2, 3, "abcdegshijk", time.Now()},
			{1, 4, 2, 3, "abcdegshijk", time.Now()},
			{1, 5, 2, 3, "abcdegshijk", time.Now()},
			{1, 6, 2, 3, "abcdegshijk", time.Now()},
			{1, 7, 2, 3, "abcdegshijk", time.Now()},
			{1, 8, 2, 3, "abcdegshijk", time.Now()},
			{1, 9, 2, 3, "abcdegshijk", time.Now()},
			{1, 10, 2, 3, "abcdegshijk", time.Now()},
			{1, 11, 2, 3, "abcdegshijk", time.Now()},
			{1, 12, 2, 3, "abcdegshijk", time.Now()},
			{1, 13, 2, 3, "abcdegshijk", time.Now()},
			{1, 14, 2, 3, "abcdegshijk", time.Now()},
			{1, 15, 2, 3, "abcdegshijk", time.Now()},
			{1, 16, 2, 3, "abcdegshijk", time.Now()},
			{1, 17, 2, 3, "abcdegshijk", time.Now()},
			{1, 18, 2, 3, "abcdegshijk", time.Now()},
			{1, 19, 2, 3, "abcdegshijk", time.Now()},
			{1, 20, 2, 3, "abcdegshijk", time.Now()},
			{1, 21, 2, 3, "abcdegshijk", time.Now()},
			{1, 22, 2, 3, "abcdegshijk", time.Now()},
			{1, 23, 2, 3, "abcdegshijk", time.Now()},
			{1, 24, 2, 3, "abcdegshijk", time.Now()},
			{1, 25, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
			{1, 26, 2, 3, "abcdegshijk", time.Now()},
		}, nil
	}
	return nil, errors.New("error")
}
