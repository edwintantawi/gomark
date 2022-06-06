CREATE TABLE IF NOT EXISTS bookmarks(
    id varchar(50) unique primary key ,
    title varchar(60) not null ,
    description text not null ,
    url text not null ,
    created_at timestamp default now(),
    updated_at timestamp default now()
);