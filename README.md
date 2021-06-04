## server
server listen on :80

## db

### db_conf
```
User:     postgres
Password: postgres
Addr:     localhost:5432
Database: postgres
```
### tables
#### relationships 
```
create table relationships
(
from_uid bigint,
to_uid   bigint,
state    varchar(20),
id       serial not null
constraint relationships_pk
primary key
);

alter table relationships
owner to postgres;

create unique index relationships_from_uid_to_uid_uindex
on relationships (from_uid, to_uid);

create unique index relationships_id_uindex
on relationships (id);
```

#### users
```
create table users
(
uid  bigserial not null
constraint users_pk
primary key,
name text,
type text
);

alter table users
owner to postgres;

```
