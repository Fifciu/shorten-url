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

- [x] create TS Client
- [x] Bug: Fix sorting by ASCII code to alphabetical sort (collation)
- [x] create client project
- [x] sign up
- [x] sign in
- [x] sign out
- [x] user can shorten URL
- [x] user can shorten URL with custom alias
- [x] user can list his links
- [x] user can remove his links
- [x] user can update his links

v0.0.4:

- [ ] user can update email
- [ ] user can update password
- [x] adjust date format in links' updated_at field
- [ ] refactor basebutton to tailwind
- [ ] make it possible to use inline for link in base button

Backlog:

- [ ] Counting views
- [ ] add linter (go)
- [ ] add db migrations on launch
- [ ] Wishlist
- [ ] SSO
- [x] Bugfix: Just added link is wrongly sorted by date
- [ ] Bugfix: Width of no links yet is different than table of links
- [ ] Bugfix: Use scoped styles in layouts' headers
- [ ] Frontend : Map font size to rem
- [ ] Frontend: new breakpoints
- [ ] Frontend: prepare tailwind config
- [ ] Frontend: migrate to tailwind
- [ ] Frontend a11y: Focus trap for modals
- [ ] write E2E tests
- [ ] Backend: Refactor (tidy auth/user controllers)
- [ ] Backend: Improve logging
- [ ] Backend: Improve error responses
- [ ] Backend: Add swagger
- [ ] Frontend: Add loaders
- [ ] Frontend: Icons in link buttons should be at least 18px wide
