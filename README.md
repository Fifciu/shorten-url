# Shorten URL

Simple project created to play more with backend development including database's usage. 

## Functional requirements

v0.0.1 (go):

- [x] user can sign up,
- [x] user can sign in,
- [x] user can sign out,
- [x] user can shorten an URL,
- [x] user can shorten an URL with custom alias,
- [x] anyone can access shortened URL
- [x] redirects to 404 if shortened URL doesn't exist

v0.0.2:

- [x] user can remove his links
- [ ] user can update his links
- [ ] user can read his data (for now it's just name), somethine like GET /me endpoint
- [IN PROGRESS] rewrite to Go
- [x] user can show his links
- [ ] user can change password (my account)
- [ ] user can change name (my account)
- [ ] user can reset password (sending email with code)
- [ ] link has field updatedAt
- [ ] user has field updatedAt

v0.0.3:

- [ ] create client project
- [ ] sign up
- [ ] sign in
- [ ] sign out
- [ ] user can shorten URL
- [ ] user can shorten URL with custom alias
- [ ] user can shorten an URL with custom expiry date, lower or equal 5 years,
- [ ] user can list his links
- [ ] user can remove his links
- [ ] user can update his links

Backlog:

- [ ] URLs expire after 5 years (I still need to specify how will it look like, I mean - expiration process),
- [ ] Counting views
- [ ] add linter
