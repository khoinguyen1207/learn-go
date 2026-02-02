create table if not exists users (
	id serial primary key,
	name varchar(50) not null,
	email varchar(100) unique not null
);

create table if not exists user_profiles (
	profile_id serial primary key,
	user_id int unique not null,
	phone varchar(10),
	address varchar(100),
	constraint fk_user foreign key (user_id) references users(id) on delete cascade
);

drop table users;
drop table user_profiles;

create table if not exists categories (
	category_id serial primary key,
	name varchar(100) not null
);

create table if not exists products (
	product_id serial primary key,
	category_id int not null,
	name varchar(100) not null,
	price int not null check(price >= 0),
	image varchar(200),
	status int not null check(status in (1, 2)),
	foreign key (category_id) references categories(category_id) on delete restrict
);
drop table products;
drop table categories;

create table if not exists students (
	student_id serial primary key,
	name varchar(100) not null
);

create table if not exists courses (
	course_id serial primary key,
	name varchar(100) not null
);

create table if not exists students_courses (
	student_id int,
	course_id int,
	primary key (student_id, course_id),
	foreign key (student_id) references students(student_id) on delete cascade,
	foreign key (course_id) references courses(course_id) on delete cascade
);

drop table students_courses;
drop table students;
drop table courses;



	


