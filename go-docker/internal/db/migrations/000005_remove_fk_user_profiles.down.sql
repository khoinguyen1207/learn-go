alter table profiles 
add column user_id int unique not null;

alter table profiles
add foreign key (user_id) references users(id) on delete cascade;