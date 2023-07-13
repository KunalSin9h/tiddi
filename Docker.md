# Run Tiddi in Docker

#### Prerequisites

- [Docker](https://docs.docker.com/get-docker/) (20.10.22 or higher recommended)
- [Sqlite3](https://www.sqlite.org/download.html) (3.37.2 or higher recommended)

#### Steps

1. Pull the image from Docker Hub

```bash
docker pull ghcr.io/kunalsin9h/tiddi-container:latest
```

> :heavy_check_mark: This image is very small in size (only 10.5 MB)

2. Create a database folder in the root directory

```bash
mkdir database
```

> Database is stored in host machine (using [Bind Mounts](https://docs.docker.com/storage/bind-mounts/)), to make it persistent between container runs and to avoid data loss.

3. Create a sqlite3 database file in the database folder

```bash
touch database/dev.db
```

5. Run the server

```bash
docker run \
    -d -p 5656:5656 --name tiddi --rm \
    --mount type=bind,source="$(pwd)"/database/dev.db,target=/tiddi/database/dev.db \
    ghcr.io/kunalsin9h/tiddi-container:latest
```

6. Open the browser and go to [http://localhost:5656](http://localhost:5656)

### Environment Variables

- `PORT` - Port number on which the server will run (default: `5656`)
- `DB` - Path to the database file (default: `./database/dev.db`)
- `HOST` - Host name (default: `http://localhost`)

> You can see more info about environment variables in Tiddi: [Here](https://github.com/KunalSin9h/tiddi#environment-variables)

#### Full Command using Environment Variables

```bash
docker run \
    -d --name tiddi --rm \
    -p 5656:5656 \
    -e PORT=5656 \
    -e DB=./database/dev.db \
    -e HOST=https://tiddi.kunalsin9h.com \
    --mount type=bind,source="$(pwd)"/database/dev.db,target=/tiddi/database/dev.db \
    ghcr.io/kunalsin9h/tiddi-container:latest
```

Now open the browser and go to [http://localhost:5656](http://localhost:5000)

> :warning: Make sure the HOST domain name is same as the one you are using to access the server

### Notes

- The server will be running in the background, to stop the server, run the following command

```bash
docker stop tiddi
```

- To run the server in the foreground, remove the `-d` flag from the above command
