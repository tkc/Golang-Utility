package HTML

func ReplaceHtmlBr(str string) string {
		reStr := strings.Replace(str, "\n", "<br/>", -1)
		return reStr
}