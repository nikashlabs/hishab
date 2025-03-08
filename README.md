# Hishab

Hishab (A Bengali word) means keeping track of your assets/income/expense

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
    POSTGRES_HOST={your_host}       # localhost
    POSTGRES_PORT={your_port}
    POSTGRES_SSLMODE={your_value}
    # Cache
    REDIS_HOST={your_host}          # localhost
    REDIS_PORT={your_port}
    REDIS_PASSWORD={your_password}
    REDIS_DB=0                      # Redis logical database number (typically 0-15)
    REDIS_PROTOCOL=3                # Redis protocol version (2 or 3)
    # Others
    PRODUCTION={true/false}
    ```

2. Make sure you got [docker](https://www.docker.com/) installed.

3. There's a `docker-compose.yml` file in the root directory, specifying the containers' configurations. 
    - postgres
    - redis
    - apiserver

4. Run the docker containers

    ```
    docker-compose up -d
    ```

## Development

### Run the Server Locally (Development)

Make sure, [Go](https://go.dev/doc/install) is already installed.

1. **Windows**
    
    To start the server, run the following powershell commands

    ```ps
    go build -o ./bin/hishab-api.exe ./cmd/apiserver
    .\bin\hishab-api.exe
    ```

    To stop the server, invoke keyboard interrupt (Press `CTRL+C`) or run the following powershell command

    ```ps
    Get-Process hishab-api | Stop-Process -Force
    ```

2. **Linux/MacOS**

    To start the server, run the following shell commands

    ```shell
    go build -o ./bin/hishab-api ./cmd/apiserver
    ./bin/hishab-api
    ```

    To stop the server, invoke keyboard interrupt (Press `CTRL+C`) or run the following shell command

    ```shell
    pkill -f hishab-api
    ```

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

1. Reset Docker Containers (Database, Cache)

    ```ps
    docker-compose down | docker volume rm hishab_postgres_data hishab_redis_data | docker-compose up -d
    ```

2. Run Migrations

    ```ps
    migrate -path "internal/database/migrations/" -database "postgres://{your_user}:{your_password}@{your_host}:{your_port}/{your_database}?sslmode=disable" up
    ```

