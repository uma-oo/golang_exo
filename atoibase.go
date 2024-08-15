package piscine 
func checkBase(base string) bool {
	if len(base) >= 2 && len(base) <= 20 {
		result := true
		for i := 0; i < len(base); i++ {
			for j := i + 1; j < len(base); j++ {
				if base[i] == base[j] || base[i] == '+' || base[i] == '-' {
					result = false
					break
				} else {
					continue
				}
			}
		}
		return result
	} else {
		return false
	}
}
func Index(s string, toFind string) int {
	result := -1
	for i := 0; i+len(toFind) <= len(s); i++ {
		if s[i:len(toFind)+i] == toFind {
			result = i
			break
		}
	}
	return result
}
func Power(number, x int) int {
  nbr:=1
  if x==0{
    nbr=1
  }
  for i:=0;i<x;i++{
    nbr=number*nbr
  }
  return nbr
}
func AtoiBase(s string, base string) int {
if !checkBase(base){
  return 0
} else {
  nbr:=0
 for i,j:=len(s)-1,0;i>=0 && j<len(s);i,j=i-1,j+1{
   nbr+=Power(len(base),i)*Index(base,string(s[j]))
 }
  return nbr
  
}
 
}
