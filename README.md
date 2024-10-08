![awanfiles banner](https://github.com/user-attachments/assets/1bc1c9e2-acbf-4842-93ad-884ab814cef6)

# Awan Files

Awan Files is primarily focused towards storing photos, pictures and audios in the cloud that is secure, and easily sharable. Original intention was to create a file server running on a VPS with a web dashboard, and an API, as a fun project to play with.

"Cloud" in Malay, is called "Awan", hence why this project is named Awan Files - or Cloud Files.

## Solutions

### Cloud storage

Upload and store your files in our cloud server. Manage and access your files anywhere via our web dashboard, or the API - easily integrate Awan Files into your existing projects.

## Stack

- Backend: Go, with Echo
- Frontend: HTML, served by backend
- Database: SQLite

## To-do

- [x] Upload file via API
  - [x] Add file to database
  - [x] Create file
  - [x] Copy file, then write
- [x] Download file via API
  - [ ] Verify file ownership
  - [x] Download file as attachment
- [ ] Web dashboard
  - [ ] Landing page
- [ ] Authentication
  - [ ] Create user
  - [ ] Log in
  - [ ] Verify user on each request

----

Please note the Awan Files source code cannot be used for commercial purposes. Awan Files is made open-source to provide transparency, build trust and credibility to our users.

© Awan Files 2024 | Designed by [theluqmn](theluqmn.github.io)
