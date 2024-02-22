# Worklog

## 2024-02-22 Sorting issue

Sorting behaves in unexpected way in Postgres. It looks like it uses ASCII code instead of alphabetical sort for `text`. I tried to sort links by `name` field, and `Zard` was before `andre` in ascending direction. Steps below was done to solve the issue.

### Migrated to Postgres 16 from 14

Because of issues with sorting I decided to migrate to the newest version but it didn't solve an issue.

### Installed PgAdmin in Docker compose

I was thinking about that dozens of times. It feels much better than `adminer` tool. And it was much more pleasent to debug using it.

### Played with COLLATE and CTYPE

#### Finding what's wrong

I found out there is something called COLLATE and CTYPE and it determines how sort is being done. A value of COLLATE for collumns was equal `pg_catalog."default"` which in the end solves to the COLLATE differentiating big and small letters in sorting alphabetically. This was a root of the issue.

I asked database for list of available collations using:

```sql
SELECT * FROM pg_collation;
```

I found there `en-US-x-icu` and specifying it in the SELECT query showed me desired collate.

```sql
SELECT * FROM public.links ORDER BY name COLLATE "en-US-x-icu";
```

With that, sort works as I expected - `Z` is after `a` in ascending direction.

#### Finding solution

I understood, I should create database using `COLLATE "en-US-x-icu"`. As my database was still pretty tiny, I updated configuration of postgres container by adding the following environment variable:

```yaml
- POSTGRES_INITDB_ARGS="-E 'UTF-8' --lc-collate='en-US-x-icu' --lc-ctype='en-US-x-icu'"
```

And recreated my database from scratch, filled it with some data and finally I have sort working as expected.

### Buggy COLLATE on Mac OS

During googling session, I found many topics related to buggy behavior of COLLATE on MacOS. Postgres is using COLLATE from operating system, and something is wrong with `en_US.utf8` collate on MacOS. I wasted some time in this rabbit hole. The truth is I forgot about obvious thing, I am inside Docker container, on some Linux based image. So it looks like collate inside this image causes unexpected behavior.

As I didn't want to modify Docker image and found available collate solving my issue, I went for using `en-US-x-icu`.
