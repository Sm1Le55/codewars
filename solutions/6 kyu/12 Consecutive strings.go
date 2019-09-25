package kata

func LongestConsec(strarr []string, k int) (result string) {
	if len(strarr)==0 || len(strarr) < k || k <=0 {
		return
	}
  
	maxWordsLen := 0
	
  	for i := 0; i <= len(strarr)-k; i++ {
		wordsLen := 0
		buf := ""
		for j := 0; j < k; j++ {
			wordsLen += len(strarr[i+j])
			buf += strarr[i+j]
		}
		if wordsLen > maxWordsLen {
			maxWordsLen = wordsLen
			result = buf
		}
	}

  return result
}
