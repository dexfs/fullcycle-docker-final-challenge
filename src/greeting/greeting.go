package main

func greeting(text string) string {
	if text == "" {
		return "<b>Code.education Rocks!<b>"
	} else {
		return "<b>" + text + "</b>"
	}
}
