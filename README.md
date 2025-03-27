# Hishab

Hishab (হিসাব), in Bengali, means to keep track of one's income/expense/asset.

## License

This project is licensed under the GNU General Public License v3.0. See the [LICENSE](LICENSE) file for details.

## Getting Started

### Run Server

1. Setup `.env` file

    ```shell
    # Server
    HOST=0.0.0.0                    # for the docker container to listen on all interfaces
    PORT={your_port}
    # Database
    POSTGRES_USER={your_user}
    POSTGRES_PASSWORD={your_password}
    POSTGRES_DB={your_database}
    POSTGRES_HOST=postgres          # according to service name in docker-compose.yml
    POSTGRES_PORT=5432              # inter-container communication in docker, ignores port mapping
    POSTGRES_SSLMODE={your_value}
    # Cache
    REDIS_HOST=redis                # according to service name in docker-compose.yml
    REDIS_PORT=6379                 # inter-container communication in docker, ignores port mapping
    REDIS_PASSWORD={your_password}
    REDIS_DB=0                      # redis logical database number (typically 0-15)
    REDIS_PROTOCOL=3                # redis protocol version (2 or 3)
    # Others
    PRODUCTION={true/false}
    ```

2. Install [docker](https://www.docker.com/) on the system

3. The root directory contains `docker-compose.yaml` which specifies the containers' configurations.
    - postgres
    - redis
    - apiserver

4. Run the docker containers
    -   **Commmands**
        View available commands provided by Makefile:
        ```shell
        make help
        ```
    -   **Production**
        ```shell
        make prod-build
        ```
        Which basically runs `docker compose up` with `--build` under the hood. Check [Makefile](./Makefile) for more information. For running without build run this:
        ```shell
        make prod
        ```


    -   **Development** (with hot reloading)
        With `--build`
        ```shell
        make dev-build
        ```

        After first time build you can run without `--build`
        ```shell
        make dev
        ```


### Reset Server

Use this command:
```shell
make reset
```

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

1. Run Migrations

    ```ps
    migrate -path "internal/database/migrations/" -database "postgres://{your_user}:{your_password}@{your_host}:{your_port}/{your_database}?sslmode=disable" up
    ```

