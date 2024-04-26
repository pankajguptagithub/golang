func checkPalindrome(s string)bool{
    var flag bool 
    flag = true
    for i:=0;i<len(s)/2;i++{
            if s[i] != s[len(s) -i-1 ]{
                flag = false
                break
            }
    }
    return flag
}
func generateSubstring(s string)[]string{
    var subStringList []string
    for i:=0;i<len(s);i++{
        for j:=i;j<len(s);j++{
            str := s[i:j+1]
            if checkPalindrome(str){
                subStringList = append(subStringList,s[i:j+1])
            }
        }
    }
    return subStringList
}
func longestPalindrome(s string) string {
    var max_length,l int
    var res string
    max_length,l = 0,0
    if len(s) == 1{
        return s
    }else{
        subStringList := generateSubstring(s)
        for i:=0;i<len(subStringList);i++{
            str := subStringList[i]           
            l = len(str)
            if l >= max_length{                
                res = str
                max_length = l
            }         
        }
    }    
    fmt.Println("Result is ",res)
    return res
}
