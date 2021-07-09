package util

import "os"

// get missing string in str1 - (str1 & str2.name)
func FindMissing(str1 []string, entries []os.DirEntry) []string {
	s := make(map[string]bool, len(entries))
	for i := range entries {
		s[entries[i].Name()] = true
	}

	missing := make([]string, 0, len(str1))
	for i := range str1 {
		if _, ok := s[str1[i]]; !ok {
			missing = append(missing, str1[i])
		}
	}
	return missing
}
