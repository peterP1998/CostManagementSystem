CREATE table User(
   id int NOT NULL primary key unique auto_increment,
   username varchar(255) not null unique,
   email varchar(255) not null unique,
   password varchar(255) not null,
   admin boolean not null
);
CREATE table Expense(
   id int NOT NULL primary key unique auto_increment,
   description varchar(255) not null,
   value double not null,
   userid int,
   foreign key (userid) references User(id)
);
CREATE table Income(
   id int NOT NULL primary key unique auto_increment,
   description varchar(255) not null,
   value double not null,
   userid int,
   foreign key (userid) references User(id)
);
create table User_Group(
   id int NOT NULL primary key unique auto_increment,
   groupname varchar(255) not null unique,
   moneybynow double not null unique,
   targetmoney double not null
);
CREATE TABLE group_user(
    user_id int NOT NULL,
    group_id int NOT NULL,
    FOREIGN KEY (user_id) REFERENCES User(id), 
    FOREIGN KEY (group_id) REFERENCES User_Group(id),
    UNIQUE (user_id, group_id)
);