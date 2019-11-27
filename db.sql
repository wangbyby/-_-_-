#pgsql
create database steam;
\c steam;
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
    gname character varying(30) primary key,
    price money NOT NUll default 20
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


#游戏发行商
#'' : 字符串
insert into steam.gc(gcname,address,info) values ('NewWorld', 'American','小公司');
insert into steam.gc(gcname,address,info) values ('Valve', 'American','黑心公司');
insert into steam.gc(gcname,address,info) values ('Paradox', 'Sweden','沙盒类游戏');
insert into steam.gc(gcname,address,info) values ('Crytek', 'Germany frankfurt','深受盗版困扰');
insert into steam.gc(gcname,address,info) values ('Petroglyph', 'American Las Vegas','生于西木');
insert into steam.gc(gcname,address,info) values ('505 Games', 'Italy','知名游戏厂商');
insert into steam.gc(gcname,address,info) values ('2K','Boston/Ausralis','知名游戏厂商');
insert into steam.gc(gcname,address,info) values  ('Ubisoft','France','大型跨国公司');
insert into steam.gc(gcname,address,info) values    ('Activision', 'American', '第三方视频游戏');
insert into steam.gc(gcname,address,info) values    ('Running With Scissors', 'American', '没有资料');

#游戏玩家
insert into steam.player(uname, upwd, nickname,balance, level) values ('clearlove9', '000000', 'yc',0,0);
insert into steam.player(uname, upwd, nickname,balance, level) values ('cy', '123456', '深海哥',100,3);
insert into steam.player(uname, upwd, nickname,balance, level) values ('bywww', 'password', 'by',0,2);
insert into steam.player(uname, upwd, nickname,balance, level) values ('tyj', 'root', '欢乐多',50,3);
insert into steam.player(uname, upwd, nickname,balance, level) values ('DioBrando', 'theworld', 'dio',0,10);
insert into steam.player(uname, upwd, nickname,balance, level) values ('KujoJotaro', '111111', 'jojo',1000,10);
insert into steam.player(uname, upwd, nickname,balance, level) values ('taige', 'qwerty', 'tg',200,7);
insert into steam.player(uname, upwd, nickname,balance, level) values ('xwj', 'root', '深蓝的海',0,4);
insert into steam.player(uname, upwd, nickname,balance, level) values ('Kars', 'sunshine', '究极生物',10000,100);
insert into steam.player(uname, upwd, nickname,balance, level) values ('Wamuu', '1234567', '瓦乌姆',0,90);

#游戏
insert into steam.game(gname,price) values ('Insurency', 20);
insert into steam.game(gname,price) values ('CSGO',50);
insert into steam.game(gname,price) values ('CallofDuty', 40);
insert into steam.game(gname,price) values ('Dota2',0);
insert into steam.game(gname,price) values ('8BitInvadier',5);
insert into steam.game(gname,price) values ('Ass',10);
insert into steam.game(gname,price) values ();
insert into steam.game(gname,price) values ();
insert into steam.game(gname,price) values ();
insert into steam.game(gname,price) values ();