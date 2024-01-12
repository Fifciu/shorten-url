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
- [x] user can't add two links with the same name
- [x] user can update his links
- [x] rewrite to Go
- [x] user can show his links
- [x] user can change password (my account)
- [x] user can change name (my account)
- [x] link has field updatedAt
- [x] user has field updatedAt

v0.0.3:

- [ ] refactor (tidy auth/user controllers)
- [ ] fix controller tests
- [x] create TS Client
- [ ] write e2e tests in ts client
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
- [ ] add linter (go)
- [ ] add db migrations on launch
- [ ] Wishlist
- [ ] SSO
- [ ] Bugfix: Just added link is wrongly sorted by date