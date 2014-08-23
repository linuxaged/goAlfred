(* Notebook name to create these journal entries in *)
property nb : "Daily Diary"

(* BEGIN HANDLER CALL

NOTE: I have added a new portion to the script so that users who
aren't familiar with how to use a handler can run the Script
directly and check it out. If you'd like to use this as a handler
in your own AppleScript, Just delete or comment out the portion
of code between "BEGIN HANDLER CALL" and "END HANDLER CALL"

*)

(*set notetext to text returned of (display dialog "Diary Entry" default answer "")
my handle_string(notetext)

on handle_string(notetext)
if notetext is not "" then
CreateDailyEvernote(notetext)
end if
end handle_string
*)
on alfred_script(q)
if q is not "" then
CreateDailyEvernote(q)
end if
end alfred_script

(* END HANDLER CALL *)

on CreateDailyEvernote(txt)
set t to "Daily Journal " & (do shell script "date +'%Y-%m-%d'")
set crlf to (ASCII character 13) & (ASCII character 10)
set timeStr to time string of (current date)

(* keep track of the app that was running when we hit the hotkey *)
set frontApp to (path to frontmost application as Unicode text)

tell application "Evernote"
set foundNotes to find notes "notebook:\"" & nb & "\"" & " intitle:\"" & t & "\""
set found to ((length of foundNotes) is not 0)
if not found then
set curnote to create note with text crlf & timeStr & ": " & txt & crlf title t notebook nb
else
repeat with curnote in foundNotes
tell curnote to append text crlf & timeStr & ": " & txt & crlf
end repeat
end if
activate
end tell

(* put the old foreground app back on top *)
tell application frontApp to activate

(* just to be sure there are no odd effects, explicitly hide Evernote *)
tell application "System Events"
set visible of process "Evernote" to false
end tell

end CreateDailyEvernote