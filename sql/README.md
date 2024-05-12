## PostgreSQL server & client

> Steps for **MacOS**

- install PostgreSQL

```zsh
brew install postgresql
```

### Server

- start a PostgreSQL server

```zsh
brew services start postgresql
```

**OR**

- use [DBngin](https://dbngin.com/) to manage your local database servers with a simple UI:
  > ![dbngin example](https://tableplus.com/assets/images/dbngin/dbngin-local.png)

### Client

I am using [pgAdmin](https://www.pgadmin.org/) as my PostgreSQL client

- Connect to the server with the name & port you specified (the default/standard port is `5432`)
  - Easier to specify with **DBngin**.

---

- Add your database URL to `.env`

```env
PORT=8080
DB_URL=postgres://<user>:<optional_password>@localhost:<PORT>/<database_name>
```

---

### `sqlc.yaml`

```yaml
version: "2"
sql:
  - schema: "sql/schema"
    queries: "sql/queries"
    engine: "postgresql"
    gen:
      go:
        out: "internal/database"
```

### Commands

- Install [sqlc](https://sqlc.dev/) & [goose](https://github.com/c9s/goose) (globally with **brew** or locally with **go**)

---

- run `up` migration

```zsh
cd sql/schema
```

```zsh
goose postgres <database_url> <database_name> up
```

- run `down` migration

```zsh
goose postgres <database_url> <database_name> down
```
> generates/drops the schemas in the database

- generate query

```zsh
cd ../../
```

```zsh
sqlc generate
```

> generates `internal/database/` files to interact with your sql schemas