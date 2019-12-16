insert into public.stu (id,name) values (1,'wang');
insert into public.stu (name) values ('zhang');

delete from public.stu where id =1;

select user; #查看用户
#自增ID
create table test_c 
(
  id integer PRIMARY KEY,
  name character varying(128)
);  
CREATE SEQUENCE test_c_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
    
alter table test_c alter column id set default nextval('test_c_id_seq');