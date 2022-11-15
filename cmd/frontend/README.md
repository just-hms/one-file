# One-file frontend

This is an executable that watches a file and upload it's changes to the server

## Upload

- a watcher is linked to a file in the system
- every time that the file is modified the watcher raise an event
- the program handle the event making a PUT request to the server

## Update

- a watcher is linked to a file in the system
- if the file is open
	- every 2 second a GET request is maid to the server to check if any update was maid
 