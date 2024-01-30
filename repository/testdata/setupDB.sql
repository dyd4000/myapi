create table if not exists article (
article_id integer unsigned auto_increment primary key, title varchar(100) not null,
contents text not null,
username varchar(100) not null,
nice integer not null,
created_at datetime
);
create table if not exists comment (
comment_id integer unsigned auto_increment primary key, article_id integer unsigned not null,
message text not null,
created_at datetime,
foreign key (article_id) references article(article_id)
);
insert into article (title, contents, username, nice, created_at) values
('firstPost', 'This is my first blog', 'daniel', 2, now()); insert into article (title, contents, username, nice) values
('2ndPost', 'Second blog post', 'daniel', 4);
insert into comment (article_id, message, created_at) values
    (1, '1st comment yeah', now());
insert into comment (article_id, message) values (1, 'welcome');