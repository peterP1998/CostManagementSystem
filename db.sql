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