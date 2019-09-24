package kata

import "strings"

func DNAStrand(dna string) string {
  result := ""
  var strand map[string]string = map[string]string{
   "A": "T",
   "T": "A", 
   "C": "G", 
   "G": "C", 
   }
   for _, c := range dna {
     symb := string(c)
     result += strings.Replace(symb,symb,strand[symb],1)
   }
   return result   
}