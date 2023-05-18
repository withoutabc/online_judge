package util

import (
	"sort"
	"time"
)

func CheckTime(from, to string) error {
	_, err := time.Parse("2006-01-02 15:04:05", from)
	if from != "" && err != nil {
		return err
	}
	_, err = time.Parse("2006-01-02 15:04:05", to)
	if to != "" && err != nil {
		return err
	}
	return nil
}

func InArray(val interface{}, array interface{}) (exists bool, index int) {
	index = sort.Search(len(array.([]interface{})), func(i int) bool {
		return array.([]interface{})[i] == val
	})
	exists = index < len(array.([]interface{})) && array.([]interface{})[index] == val
	return exists, index
}
