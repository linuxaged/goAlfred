Workflows
--------
<table>
<tr><td>Wordsbook</td><td>添加新单词及其中文翻译到 Evernote</td></tr>
</table>

创建步骤

1. go get github.com/linuxaged/goAlfred

2. Worflows -> New Templates -> Essentials -> Script Filter to AppleScript

	· right click "Script Filter" add `./wb {query}`
	
	· right click "Run NSAppleScript" add the evernote.scpt

3. go build wb.go

4. put wb executable file in workflow folder




