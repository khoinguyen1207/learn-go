insert into users (name, email) values ('Nguyen', 'nguyen@gmail.com');
insert into users (name, email) values ('Hehe', 'nguyen1@gmail.com');
insert into users (name, email) values ('Ookok', 'nguyen2@gmail.com');

insert into user_profiles (user_id, phone, address) values (4, 1293213992, 'abc');
insert into categories (name) values ('Dien Thoai'), ('Laptop');

update users set name = 'Nguyen Khoi', email = 'nguyenkhoi@gmail.com' where id  = 4;

delete from users where id=4;