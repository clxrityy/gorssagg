# go-rss-agg

Go RSS feed aggregator

---

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

```.env
PORT=8080
DB_URL=postgres://<user>:<optional_password>@localhost:PORT
```
