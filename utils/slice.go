package utils

func ContainsInt(s []int, el int) bool {
	if len(s) == 0 {
		return false
	}

	for _, sEl := range s {
		if el == sEl {
			return true
		}
	}
	return false
}
