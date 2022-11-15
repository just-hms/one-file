# One-file frontend

This is an executable that watches a file and upload it's changes to the server

## Upload

- a watcher is linked to a file in the system
- every time that the file is modified the watcher raise an event
- the program handle the event making a PUT request to the server

## Update

every 2 second a GET request is made to the server to check if any update was made
 