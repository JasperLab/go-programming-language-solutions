package main

func dedup(s []string) []string {
	if len(s) == 0 {
		return s
	}
	i := 1 
	last := s[0]
	for j := 1; j < len(s); j++ {
		if s[j] != last {
			s[i] = s[j]
			last = s[i]
			i++
		}
	}
			
	return s[:i]
}
