package interval

type SortByDate []Interval

func (s SortByDate) Len() int {
	return len(s)
}

func (s SortByDate) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortByDate) Less(i, j int) bool {
	a, b := s[i], s[j]

	if a.Before(b) {
		return true
	}

	if a.After(b) {
		return false
	}

	return a.Tag < b.Tag
}
