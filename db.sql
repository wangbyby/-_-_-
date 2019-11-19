#pgsql
create database steam
\c steam
create table gc(
    gcname character varying(20) primary key,
    address character varying(40) NOT NUll default '',
    info character varying(40) NOT NULL default '游戏公司'
)WITH (
    OIDS = FALSE
)TABLESPACE pg_default;

#游戏玩家
create table player( 
    uname character varying(15) primary key,
    upwd character varying(15) NOT NUll default '',
    nickname char(15) NOT NUll default '',
    balance money NOT NUll default 0,
    level smallint NOT NUll default 0
)WITH (
    OIDS = FALSE
)TABLESPACE pg_default;


create table game(
    gname character varying(20) primary key,
    price money NOT NUll default 100
)WITH (
    OIDS = FALSE
)TABLESPACE pg_default;


create table community(
    cid bigserial primary key,
    content character varying(120) NOT NUll default '',
    ctype character varying(10) NOT NUll default '',
    ctime timestamp with time zone NOT NUll 
)WITH (
    OIDS = FALSE
)TABLESPACE pg_default;

create table issue(
    gcname character varying(20) references gc(gcname) on update CASCADE on delete CASCADE   ,
    gname  character varying(20) REFERENCES game(gname) on update CASCADE on delete CASCADE ,
    itime timestamp  with time zone NOT NUll ,
    iinfo character varying(40) NOT NUll default '',
    itype character varying(20) NOT NUll default '',
    primary key(gcname,gname)
)WITH (
    OIDS = FALSE
)TABLESPACE pg_default;

#lpt : last play time 最近运行时间
#bt : 购买时间
create table own(
    uname character varying(15) references player(uname) on update CASCADE on delete CASCADE   ,
    gname character varying(20) references game(gname) on update CASCADE on delete CASCADE ,
    lrt timestamp  with time zone NOT NUll,  
    bt timestamp  with time zone NOT NUll,
    btype smallint NOT NUll default 0,
    sid bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    primary key(uname,gname)
)WITH (
    OIDS = FALSE
)TABLESPACE pg_default;

create table com_con(
    cid bigserial references community(cid) on update CASCADE on delete CASCADE,
    gname character varying(20) references game(gname) on update CASCADE on delete CASCADE ,
    primary key(cid,gname)
)WITH (
    OIDS = FALSE
)TABLESPACE pg_default;


