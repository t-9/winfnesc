/*
Package winfnesc is windows file name escaper.
*/
package winfnesc

import "strings"

// Escape returns that escape a string
// that cannot be used for Windows filenames.
func Escape(s string) string {
	return strings.NewReplacer(
		`\`, "￥",
		`/`, "／",
		`:`, "：",
		`*`, "＊",
		`?`, "？",
		`"`, "”",
		`<`, "＜",
		`>`, "＞",
		`|`, "｜",
	).Replace(s)
}
