

# Hello there, me going through Udemy course here

## Learnings so far
- Design the database for a simple bank project, (I added a users table, I'll adapt as the course goes)
- Using ```start``` and ```commit``` to ensure atomicity among multiple queries for a proper workflow
- Replicate a situation where deadlock might happened
- Learn how to mitigate deadlock 
  - Example shows a situation where we want to update the balance of 2 accounts in one atomic transaction
- Setup a Go server
  - Currently the file structure in the repo are my own designs. (This part I'm not following the course, a good time to practice scaffolding a Go backend repo)

## Prerequisite 

- Have Go version higher than 1.9
- Have make installed
- Have docker installed
- Have golang-migrate installed


### Create container for our dependencies
```bash
docker compose up
```

### Create database in postgres
```bash
make database.create
```

### Run migrations
```
make migrations.up
```



