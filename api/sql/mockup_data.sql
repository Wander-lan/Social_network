/* Use this commands to insert this mockup data on the database */
insert into users (name, nick, email, password)
values
("User 1", "user_1", "user1@gmail.com", "$2a$10$1VxCD6l5h5d.K4vv7zvLoesZ6z3MBP4mFI2CHVy5ZgGIA.VRDPWiO"),
("User 2", "user_2", "user2@gmail.com", "$2a$10$1VxCD6l5h5d.K4vv7zvLoesZ6z3MBP4mFI2CHVy5ZgGIA.VRDPWiO"),
("User 3", "user_3", "user3@gmail.com", "$2a$10$1VxCD6l5h5d.K4vv7zvLoesZ6z3MBP4mFI2CHVy5ZgGIA.VRDPWiO");

insert into followers(user_id, follower_id)
values
(1, 2),
(3, 1),
(1, 3);

insert into posts(title, content, author_id)
values
("Post from user 1", "This is a post from user 1! Yey!", 1),
("Post from user 2", "This is a post from user 2! Yey!", 2),
("Post from user 3", "This is a post from user 3! Yey!", 3);