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
    gname character varying(40) primary key,
    price money NOT NUll default 20
)WITH (
    OIDS = FALSE
)TABLESPACE pg_default;

#社区
create table community(
    cid bigserial primary key, #自动递增的id
    content character varying(156) NOT NUll default '',
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
#游戏
insert into steam.game(gname,price) values ('Insurgency', 20);
insert into steam.game(gname,price) values ('CSGO',50);
insert into steam.game(gname,price) values ('Destiny2', 40);
insert into steam.game(gname,price) values ('Dota2',0);
insert into steam.game(gname,price) values ('8Bit Invaders',5);
insert into steam.game(gname,price) values ('Assass',10);
insert into steam.game(gname,price) values ('BioShock',10);
insert into steam.game(gname,price) values ('Crysis',9);
insert into steam.game(gname,price) values ('Europa',15);
insert into steam.game(gname,price) values ('ABZU',8);
insert into steam.game(gname,price) values ('POSTAL2',1);


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

#社区
#current_timestamp 当前时间
insert into steam.community(content,ctype,ctime) values ('CSGO更新九头蛇大行动','更新', '2019-11-27 05:47');
insert into steam.community(content,ctype,ctime) values ('叛乱:沙尘暴Steam大奖提名','新闻', '2019-11-27 02:08');
insert into steam.community(content,ctype,ctime) values ('CSGO重大更新 裂网大行动','重大更新', '2019-11-20 02:12');
insert into steam.community(content,ctype,ctime) values ('CSGO全新 "裂网大行动"任务已解锁','更新', '2019-11-27 05:47');
insert into steam.community(content,ctype,ctime) values ('CSGO更新 创意工坊提交流程更新','新闻', '2019-4-3 08:46');
insert into steam.community(content,ctype,ctime) values ('叛乱:沙尘暴 游戏开发近况','新闻', '2019-10-30 03:42');
insert into steam.community(content,ctype,ctime) values ('叛乱:沙尘暴 新游戏模式Frontline','新闻', '2019-9-27 22:31');
insert into steam.community(content,ctype,ctime) values ('Bungie每周快报万圣节糖果','新闻', '2019-11-15 07:04');
insert into steam.community(content,ctype,ctime) values ('CS20 Submission Deadline Extended','新闻', '2019-9-24 05:14');
insert into steam.community(content,ctype,ctime) values ('CSGO Watch the Berlin Major Championship','现场直播', '2019-9-9 03:30');

#游戏发行商 发行 游戏
insert into steam.issue(gcname,gname,itime,iinfo,itype) values ('NewWorld','Insurgency');
insert into steam.issue(gcname,gname,itime,iinfo,itype) values ('Valve','CSGO');
insert into steam.issue(gcname,gname,itime,iinfo,itype) values ('Valve','Dota2');
insert into steam.issue(gcname,gname,itime,iinfo,itype) values ('Paradox','Europa');
insert into steam.issue(gcname,gname,itime,iinfo,itype) values ('505 Games','ABZU');
insert into steam.issue(gcname,gname,itime,iinfo,itype) values ('2k','BioShock');
insert into steam.issue(gcname,gname,itime,iinfo,itype) values ('Crytek','Crysis');
insert into steam.issue(gcname,gname,itime,iinfo,itype) values ('Petroglyph','8Bit Invaders');
insert into steam.issue(gcname,gname,itime,iinfo,itype) values ('Ubisoft','Assass');
insert into steam.issue(gcname,gname,itime,iinfo,itype) values ('Acticision','Destiny2');
insert into steam.issue(gcname,gname,itime,iinfo,itype) values ('Running With Scissors','POSTAL2');
