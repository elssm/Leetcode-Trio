package goLeetCode

func fullJustify(words []string, maxWidth int) []string {
	llen, index := 0, 0
	var ans []string
	for i := 0; i < len(words); i++ {
		// fmt.Println(words[index], llen)
		llen += len(words[i]) + 1
		if i == len(words)-1 && llen < maxWidth+1 {
			str := ""
			for j := index; j <= i; j++ {
				if j != i {
					str = str + words[j] + " "
				} else {
					str = str + words[j]
				}
			}
			for ll := len(str); ll < maxWidth; ll++ {
				str += " "
			}
			ans = append(ans, str)
			index = i + 1
			llen = 0
		}
		if llen == maxWidth+1 {
			str := ""
			for j := index; j <= i; j++ {
				if j != i {
					str = str + words[j] + " "
				} else {
					str = str + words[j]
				}
			}
			index = i + 1
			llen = 0
			ans = append(ans, str)
		} else if llen > maxWidth+1 {
			llen -= len(words[i]) + 1 + 1
			bn := maxWidth - llen
			wn := i - index
			abn := bn
			aabn := 0
			str := ""
			if wn > 1 {
				abn = bn / (wn - 1)
				aabn = bn % (wn - 1)
				for j := index; j < i; j++ {
					if j != i-1 {
						str = str + words[j] + " "
						for k := 0; k < abn; k++ {
							str += " "
						}
						if aabn != 0 {
							str += " "
							aabn--
						}
					} else {
						str = str + words[j]
					}
				}
			} else {
				str = str + words[index]
				for k := 0; k < abn; k++ {
					str += " "
				}
			}
			ans = append(ans, str)
			llen = 0
			index = i
			i -= 1
		}
	}
	return ans

}
