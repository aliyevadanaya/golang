create table if not exists users (
    id serial primary key,
    name varchar(255) not null,
    age int,
    gender varchar(255), --M male, F female, P prefer not to say haha
    city varchar(255),
    deleted_at timestamp null
);

-- insert into users(name) values ('Danaya Aliyeva')