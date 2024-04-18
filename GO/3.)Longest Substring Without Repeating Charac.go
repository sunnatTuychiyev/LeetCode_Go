func lengthOfLongestSubstring(s string) int {
	hashmap := map[byte]int{}
	max := 0
	for i := range s {
		_, ok := hashmap[s[i]]
		if !ok {
			hashmap[s[i]] = i
			if len(hashmap) > max {
				max = len(hashmap)
			}
		} else {
			oldI := hashmap[s[i]]
			hashmap[s[i]] = i

			for key, value := range hashmap {
				if value < oldI {
					delete(hashmap, key)
				}
			}
		}
	}
	return max
}