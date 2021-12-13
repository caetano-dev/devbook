insert into users (name, nick, email, password)
values
("user1","user1","user1@gmail.com","$2a$10$7qH/aaE/KLKepl3F7xDvCugSYp.jpUIF7wA9SUJZxzEiRtz8CD.3O"),
("user2","user2","user2@gmail.com","$2a$10$7qH/aaE/KLKepl3F7xDvCugSYp.jpUIF7wA9SUJZxzEiRtz8CD.3O"),
("user3","user3","user3@gmail.com","$2a$10$7qH/aaE/KLKepl3F7xDvCugSYp.jpUIF7wA9SUJZxzEiRtz8CD.3O");

insert into followers(user_id, follower_id)
values
(1, 2),
(3, 1),
(1, 3);

insert into posts(title, content, author_id)
values
("post1", "post1", 1),
("post2", "post2", 2),
("post3", "post3", 3);
