#pgsql
create database steam
\c steam
#游戏发行商
#一般来说 发行商与开发商 不一样
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

#游戏
create table game(
    gname character varying(20) primary key,
    price money NOT NUll default 100
)WITH (
    OIDS = FALSE
)TABLESPACE pg_default;

#社区
create table community(
    cid bigserial primary key,
    content character varying(120) NOT NUll default '',
    ctype character varying(10) NOT NUll default '',
    ctime timestamp with time zone NOT NUll 
)WITH (
    OIDS = FALSE
)TABLESPACE pg_default;

#游戏发行商 发行 游戏
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

#玩家 购买 游戏
#lpt : last play time 最近运行时间
#bt : 购买时间
#btype : 购买方式 : 微信, 支付宝, visa, 比特币...
#sid : 序列号, 自动增长
create table own(
    uname character varying(15) references player(uname) on update CASCADE on delete CASCADE,
    gname character varying(20) references game(gname) on update CASCADE on delete CASCADE ,
    lrt timestamp  with time zone NOT NUll,  
    bt timestamp  with time zone NOT NUll,
    btype  character varying(10) NOT NUll,
    sid bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    primary key(uname,gname)
)WITH (
    OIDS = FALSE
)TABLESPACE pg_default;

#游戏发布社区内容
create table com_con(
    cid bigserial references community(cid) on update CASCADE on delete CASCADE,
    gname character varying(20) references game(gname) on update CASCADE on delete CASCADE ,
    primary key(cid,gname)
)WITH (
    OIDS = FALSE
)TABLESPACE pg_default;


#'' : 字符串
insert into steam.gc(gcname,address,info) values ('NewWorld', 'American','小公司');
insert into steam.gc(gcname,address,info) values ('Valve', 'American','黑心商人');
insert into steam.gc(gcname,address,info) values ('Paradox', 'Sweden','沙盒类游戏');
insert into steam.gc(gcname,address,info) values ('Crytek', 'Germany frankfurt','深受盗版困扰');
insert into steam.gc(gcname,address,info) values ('Petroglyph', 'American Las Vegas','前西木员工')
insert into steam.gc(gcname,address,info) values ('505 Games', 'Italy','没啥')
