# Hishab

Hishab (A Bengali word) means keeping track of your assets/income/expense

## License

This project is licensed under the GNU General Public License v3.0. See the [LICENSE](LICENSE) file for details.

## Getting Started

### Setup Database

1. Create a PostgreSQL database named **hishab** using PgAdmin4 or psql

2. Set `DATABASE_URL` as environment variable.

    Windows
    ```ps
    $env:DATABASE_URL="postgres://user:password@localhost:port/hishab"
    ```
    Linux
    ```bash
    export DATABASE_URL="postgres://user:password@localhost:port/hishab"
    ```
    Don't forget to set the `user`, `password` and `port` as per your configuration.

### Run Server

Make sure, [Go](https://go.dev/doc/install) is already installed.

```ps
go run .\cmd\apiserver\ 
```