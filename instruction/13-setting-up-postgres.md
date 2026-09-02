# To Install

```sh
brew install postgresql
```

# Check Version

```sh
psql --version
```

# Start Postgres

```sh
brew services start postgresql
```

# Check Postgres User

```sh
cat /etc/passwd | grep 'postgres'
```

# Login in PSQL with postgres

```sh
psql postgres
```

# Check Current User

```sh
postgres=# SELECT current_user;
 current_user
--------------
 fulim
(1 row)
```

# Creating Database, User and Extension

```sh
postgres=# CREATE DATABASE greenlight;
CREATE DATABASE
postgres=# \c greenlight
You are now connected to database "greenlight" as user "postgres".
greenlight=# CREATE ROLE greenlight WITH LOGIN PASSWORD 'pa55word';
CREATE ROLE
greenlight=# CREATE EXTENSION IF NOT EXISTS citext;
CREATE EXTENSION
```

# Login with greenlight user to greenlight database

```sh
psql --host=localhost --dbname=greenlight --username=greenlight
greenlight=> SELECT current_user;
 current_user
--------------
 greenlight
(1 row)
```
