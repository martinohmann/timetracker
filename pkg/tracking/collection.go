package tracking

import "strings"

type Collection []Tracking

func (c Collection) Add(trackings ...Tracking) {
	for _, t := range trackings {
		c.add(t)
	}
}

func (c Collection) add(tracking Tracking) {
	if c.Contains(tracking) {
		return
	}

	c = append(c, tracking)
}

func (c Collection) Contains(tracking Tracking) bool {
	for _, item := range c {
		if item == tracking {
			return true
		}
	}

	return false
}

func (c Collection) Remove(trackings ...Tracking) {
	for _, t := range trackings {
		c.remove(t)
	}
}

func (c Collection) remove(tracking Tracking) {
	for i, item := range c {
		if item == tracking {
			c = append(c[:i], c[i+1:]...)
			return
		}
	}
}

func (c Collection) Filter(f func(Tracking) bool) Collection {
	items := make([]Tracking, 0)

	for _, item := range c {
		if f(item) {
			items = append(items, item)
		}
	}

	return Collection(items)
}

func (c Collection) Len() int {
	return len(c)
}

func (c Collection) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c Collection) Less(i, j int) bool {
	a, b := c[i], c[j]

	if a.Interval.Before(b.Interval) {
		return true
	}

	if a.Interval.After(b.Interval) {
		return false
	}

	return strings.Compare(a.Tag, b.Tag) <= 0
}
