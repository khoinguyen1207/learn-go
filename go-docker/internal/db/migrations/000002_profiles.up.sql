create table if not exists profiles (
	profile_id serial primary key,
	user_id int unique not null,
	phone varchar(10),
	address varchar(100),
	foreign key (user_id) references users(id) on delete cascade
);