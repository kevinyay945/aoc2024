package day2

func Q2(reports [][]int) int {
	output := 0
	for _, report := range reports {
		if len(report) == 0 {
			output++
			continue
		}
		if reportIsValid(report) {
			output++
			continue
		}
		for i := 0; i < len(report); i++ {
			var newReport []int
			newReport = append(newReport, report[:i]...)
			newReport = append(newReport, report[i+1:]...)
			if reportIsValid(newReport) {
				output++
				break
			}
		}
	}
	return output
}

func reportIsValid(report []int) bool {
	status := 0 // 0 for ok, 1 for increase, -1 for decrease
	isValid := true
	for i := 0; i < len(report); i++ {
		if i == 0 {
			continue
		}

		pre := report[i-1]
		curr := report[i]
		minus := curr - pre
		if minus == 0 {
			isValid = false
			break
		} else if minus < 0 {
			if status == 1 {
				isValid = false
				break
			}
			status = -1
		} else if minus > 0 {
			if status == -1 {
				isValid = false
				break
			}
			status = 1
		}

		if minus > 3 || minus < -3 {
			isValid = false
			break
		}
	}
	return isValid
}
