package tracking

import "strings"

type Collection struct {
	items []*Tracking `json:"items"`
}

func NewCollection(trackings ...*Tracking) Collection {
	return Collection{trackings}
}

func (c Collection) Add(trackings ...*Tracking) {
	for _, t := range trackings {
		c.add(t)
	}
}

func (c Collection) add(tracking *Tracking) {
	if c.Contains(tracking) {
		return
	}

	c.items = append(c.items, tracking)
}

func (c Collection) Contains(tracking *Tracking) bool {
	for _, item := range c.items {
		if item == tracking {
			return true
		}
	}

	return false
}

func (c Collection) Remove(trackings ...*Tracking) {
	for _, t := range trackings {
		c.remove(t)
	}
}

func (c Collection) remove(tracking *Tracking) {
	for i, item := range c.items {
		if item == tracking {
			c.items = append(c.items[:i], c.items[i+1:]...)
			return
		}
	}
}

func (c Collection) Filter(f func(*Tracking) bool) Collection {
	items := make([]*Tracking, 0)

	for _, item := range c.items {
		if f(item) {
			items = append(items, item)
		}
	}

	return NewCollection(items...)
}

func (c Collection) Items() []*Tracking {
	return c.items
}

func (c Collection) Len() int {
	return len(c.items)
}

func (c Collection) Swap(i, j int) {
	c.items[i], c.items[j] = c.items[j], c.items[i]
}

func (c Collection) Less(i, j int) bool {
	a, b := c.items[i], c.items[j]

	if a.Interval.Before(b.Interval) {
		return true
	}

	if a.Interval.After(b.Interval) {
		return false
	}

	return strings.Compare(a.Tag, b.Tag) <= 0
}
