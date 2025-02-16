# Hishab

Hishab (A Bengali word) means keeping track of your assets/income/expense

## License

This project is licensed under the GNU General Public License v3.0. See the [LICENSE](LICENSE) file for details.

## Getting Started

### Setup Database (PostgreSQL)

1. Setup `.env` file

    ```shell
    POSTGRES_USER={your_user}
    POSTGRES_PASSWORD={your_password}
    POSTGRES_DB={your_database}
    POSTGRES_HOST={your_host}
    POSTGRES_PORT={your_port}
    POSTGRES_SSLMODE=disable
    ```

2. Make sure you got [docker](https://www.docker.com/) installed.

3. There's a `docker-compose.yml` file in the root directory, specifying the containers configurations. 

4. Running PostgreSQL Server

    ```
    docker-compose up -d
    ```

### Run Server

Make sure, [Go](https://go.dev/doc/install) is already installed.

```ps
go run .\cmd\apiserver\ 
```

Note that, the database migrations (up) will be run everytime you spin up the server.

## Development

### sqlc

1. Install **sqlc** following this [guide](https://docs.sqlc.dev/en/latest/overview/install.html)

2. For Windows
    - Download the pre-built binary from the [guide](https://docs.sqlc.dev/en/latest/overview/install.html).
    - Extract the `.zip` file (`sqlc.exe` will be inside) and put the `.exe` into a location of your choice. 
    - Add the location to system environment variables (for ease of use). 
    - Then, `sqlc.exe` can be used from any location to perform necessary tasks.

### golang-migrate

1. Follow this [guide](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md)

2. For Windows
    - From [here](https://github.com/golang-migrate/migrate/releases), download pre-built binary as per system requirement.
    - Extract `.zip`.
    - Add the path to system environment variables.

### Powershell Commands

1. Reset Database

    ```ps
    docker-compose down | docker volume rm {your_database}_postgres_data | docker-compose up -d
    ```

2. Run Migrations

    ```ps
    migrate -path "internal/database/migrations/" -database "postgres://{your_user}:{your_password}@{your_host}:{your_port}/{your_database}?sslmode=disable" up
    ```

