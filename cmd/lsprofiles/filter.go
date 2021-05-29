package main

import "strings"

type Filter interface {
	match(profile provisioningProfile) bool
}

type CompoundFilter struct {
	filters []Filter
}

func (receiver CompoundFilter) match(profile provisioningProfile) bool {
	for _, f := range receiver.filters {
		if !f.match(profile) {
			return false
		}
	}

	return true
}

type StringContainsFilter struct {
	value string
	extractFunc func(profile provisioningProfile) string
}

func (receiver StringContainsFilter) match(profile provisioningProfile) bool {
	return strings.Contains(receiver.extractFunc(profile), receiver.value)
}
