alter table profiles 
drop constraint profiles_user_id_fkey;

alter table profiles
drop column user_id;