alter table profiles rename column phone to phone_number;

alter table profiles alter column phone_number type text;