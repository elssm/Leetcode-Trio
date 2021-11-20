package goLeetCode

func groupAnagrams(strs []string) [][]string {
	var res [][]string
	hashMap := map[[26]byte][]string{}
	for _, str := range strs {
		m := [26]byte{}
		for _, v := range str {
			m[v-'a']++
		}
		hashMap[m] = append(hashMap[m], str)
	}
	for _, v := range hashMap {
		res = append(res, v)
	}
	return res
}
