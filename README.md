

## Hello there, me going through Udemy course here

### Prerequisite 

Have Go version higher than 1.9

Have make installed

Have docker installed

Have golang-migrate installed


Create container for our dependencies
```bash
docker compose up
```

Create database in postgres
```bash
make database.create
```

Run migrations
```
make migrations.up
```



