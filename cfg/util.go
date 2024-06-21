package cfg

func ConditionalString(b bool, s1 string, s2 string) string {
	res := ""
	if b {
		res = s1
	} else {
		res = s2
	}
	return res
}
