## requirements
1. Make sure that the db-manager project has been cloned `git clone https://github.com/go-transcoder/db-manager.git`
2. Check README.md to start the db-manager. This will create the db and create migrations

# Setup
1. Start the docker containers `docker compose up -d` 
2. Make sure to watch the logs from the container. As discussed the automatic build will be done inside the container `docker logs --follow uploader-app`

