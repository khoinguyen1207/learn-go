alter table profiles rename column phone_number to phone;

alter table profiles alter column phone type varchar(10);