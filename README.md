# Intersect DevOps Project

A slack bot written in Golang to query the stock market. The bot and the database will be deployed in Docker in swarm mode

## Manual Installation

1. Create a new Slackbot and get the auth key

2. go get github.com/piquette/finance-go
   go get -u github.com/nlopes/slack
   go get github.com/lib/pq

3. Setup the environment variables

4. Start up the postgreSQL database 
    docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -d stock

5. Manually setup the database
    psql -h localhost -p 5432 -U postgres -W

    CREATE USER stockuser WITH PASSWORD '';
    GRANT CONNECT ON DATABASE stock TO stockuser;
    GRANT USAGE ON SCHEMA public TO stockuser;
    CREATE TABLE stockhistory (
    id SERIAL PRIMARY KEY,
    stocktick VARCHAR(255),
    price numeric,
    username VARCHAR(255)
    );

    GRANT SELECT, INSERT, UPDATE ON stockhistory TO stockuser;
    grant all on sequence stockhistory_id_seq to stockuser;
        
6. Build 
    cd cmd/stock-bot
    go build -o stockbot .

7. Run the bot
    chmod +x stockbot
    ./stockbot

8. Test if there is a response
    @botname 'stocktick'
    ex. @stockmarket XAW.TO

## Running in Docker Swarm

1. Create the following docker secrets 

echo "mysupersecurepassword" | docker secret create secret-db-pass -
echo "mysupersecurepassword" | docker secret create secret-db-user -
echo "mysupersecurepassword" | docker secret create secret-db -
echo "mysupersecurepassword" | docker secret create slack-token -
echo "mysupersecurepassword" | docker secret create secret-db-host-name -


2. Build the Dockerfiles for the Database and Bot
    docker build -t stock-db:latest docker-database/
    docker build -t stock-bot:latest docker-bot/

3. docker stack deploy --compose-file docker-compose.yml stock


## Future Improvements
* Add second command to find the price history by querying the PostgreSQL database and when the same Stock ticker is ran again it will show the price difference between
the last run.
* Build the stock bot in the Dockerfile
* Add error checking to the whole code base
* Add linting to the Dockerfile
* Implement CI/CD
* Add better logging to the Stock App


## Diagram
![Diagram of Project](https://github.com/frankielearns/intersectproject/blob/feature/intersect-devops-project-v1/images/Intersect%20Devops%20Project.jpeg)


