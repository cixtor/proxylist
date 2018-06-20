package main

type byFilter []Settings

func (s byFilter) Len() int {
	return len(s)
}

func (s byFilter) Swap(i int, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byFilter) Less(i int, j int) bool {
	return s[i].Filter < s[j].Filter
}
