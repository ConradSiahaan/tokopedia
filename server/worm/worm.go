package worm

func FindMinuteCount(n, u, d int) int {
	
	minuteCount := 0
	var i int

	i = 0

	if n == 0 {
		return minuteCount
	}else{
		for i < n {
			i += u
			minuteCount++
			if n <= i {
				break
			}
			minuteCount++
			i -= d
		}
		return minuteCount
	}

}

func WormTemplateStr() string {

	templateStr := `<html><head><title>Generate Minute Count For Worm</title></head><body><form action="/worm/" name=f method="GET"><input maxLength=100 size=25 name=n value="" title="inches deep" placeholder="N Inches Deep">&nbsp;<input maxLength=100 size=25 name=u value="" title="climb size" placeholder="Climb U Inches">&nbsp;<input maxLength=100 size=25 name=d value="" title="slip down size" placeholder="Slips D Inches"><input type=submit value="Find Minute Count" name=mc></form></body></html>`

	return templateStr
}