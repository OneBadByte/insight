# INSIGHT
Insight is my personal tool for logging daily activities, todos, and goal tracker.

## DATABASE
* Create a database called orbitmessengerdb
run:
```
cd src
psql -U postgres -d insight -f sql/createTables.sql
```

## BUILD SERVER
its possible to run the runServer binary on any linux machine, but to build type
```
cd src
./buildServer.py -r
```

## HTTP actions
to get messenges from the server you will need to connect to the server by port 3000
to test type
```
curl -X GET localhost:3000/getAllMessages -u username:password
```
or to add a message type
```
curl -X POST localhost:3000/addMessage -u username:password -d '{"message": "Testing"}'

```


