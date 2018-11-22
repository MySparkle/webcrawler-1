package pkgtool

// appendIfAbsent
func appendIfAbsent(s []string, t ...string) []string {
	emap := map[string]struct{}{}
	for _, se := range s {
		emap[se] = struct{}{}
	}
	for _, te := range t {
		if _, ok := emap[te]; ok {
			continue
		}
		s = append(s, te)
		emap[te] = struct{}{}
	}
	return s
}
