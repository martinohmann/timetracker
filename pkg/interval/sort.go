package interval

// SortByDate defines an interval slice that can be sorted by the interval
// dates
type SortByDate []Interval

// Len implements the sort.Interface
func (s SortByDate) Len() int {
	return len(s)
}

// Swap implements the sort.Interface
func (s SortByDate) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less implements the sort.Interface
func (s SortByDate) Less(i, j int) bool {
	a, b := s[i], s[j]

	if a.Before(&b) {
		return true
	}

	if a.After(&b) {
		return false
	}

	return a.Tag < b.Tag
}
