# To Install Migration package

```sh
brew install golang-migrate
```

# Check Migrate Version

```sh
migrate -version
```

# Create Migration

```sh
migrate create -seq -ext=.sql -dir=./migrations create_movies_table
```

- The -seq flag indicates that we want to use sequential numbering like 0001, 0002, ... for the migration files (instead of a Unix timestamp, which is the default).
- The -ext flag indicates that we want to give the migration files the extension .sql.
- The -dir flag indicates that we want to store the migration files in the ./migrations directory (which will be created automatically if it doesn’t already exist).

# Create Second Migration

```sh
migrate create -seq -ext=.sql -dir=./migrations add_movies_check_constraints
```

# Migrate migrations files

```sh
migrate -path=./migrations -database=postgres://greenlight:pa55word@localhost/greenlight?sslmode=disable up
```

Important: If you are using PostgreSQL v15, you may get the error error: pq: permission denied for schema public... when running this command. If you do, this is because v15 revokes the CREATE permission from all users except a database owner. See this StackOverflow post for more information

To get around this, set the database owner to the greenlight user:

```sh
psql postgres
ALTER DATABASE greenlight OWNER TO greenlight;
```

# Check the Migration Schema

```sh
greenlight=> \dt
                List of relations
 Schema |       Name        | Type  |   Owner
--------+-------------------+-------+------------
 public | movies            | table | greenlight
 public | schema_migrations | table | greenlight

greenlight=> SELECT * FROM schema_migrations;
 version | dirty
---------+-------
       2 | f

greenlight-> \d movies
                                        Table "public.movies"
   Column   |            Type             | Collation | Nullable |              Default
------------+-----------------------------+-----------+----------+------------------------------------
 id         | bigint                      |           | not null | nextval('movies_id_seq'::regclass)
 created_at | timestamp(0) with time zone |           | not null | now()
 title      | text                        |           | not null |
 year       | integer                     |           | not null |
 runtime    | integer                     |           | not null |
 genres     | text[]                      |           | not null |
 version    | integer                     |           | not null | 1
Indexes:
    "movies_pkey" PRIMARY KEY, btree (id)
Check constraints:
    "genres_length_check" CHECK (array_length(genres, 1) >= 1 AND array_length(genres, 1) <= 5)
    "movies_runtime_check" CHECK (runtime >= 0)
    "movies_year_check" CHECK (year >= 1888 AND year::double precision <= date_part('year'::text, now()))
```

Sure — here’s a cleaner **1 title → 1 command** Markdown format for your notes:

# Check Current Migration Version

```bash
migrate -path=./migrations -database=$EXAMPLE_DSN version
```

# Migrate to a Specific Version

```bash
migrate -path=./migrations -database=$EXAMPLE_DSN goto 1
```

# Roll Back the Most Recent Migration

```bash
migrate -path=./migrations -database=$EXAMPLE_DSN down 1
```

# Roll Back All Migrations

```bash
migrate -path=./migrations -database=$EXAMPLE_DSN down
```

# Drop All Tables

Removes all tables **including `schema_migrations`**. The database itself remains.

```bash
migrate -path=./migrations -database=$EXAMPLE_DSN drop
```

# Force Migration Version

Use this when a migration fails and the database becomes **dirty**.

```bash
migrate -path=./migrations -database=$EXAMPLE_DSN force 1
```
