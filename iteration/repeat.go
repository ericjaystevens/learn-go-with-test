package iteration

func Repeat(text string, repititions int) (concatinatedString string) {
	
	for i := 0; i < repititions; i++{
		concatinatedString += text	
	}
	return
}