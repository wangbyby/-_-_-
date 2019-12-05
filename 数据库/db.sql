--pgsql
create database steam;
\c steam;
--游戏发行商
--一般来说 发行商与开发商 不一样
create table gc(
    gcname character varying(40) primary key,
    address character varying(40) NOT NUll default '',
    info character varying(40) NOT NULL default '游戏公司'
)WITH (
    OIDS = FALSE
)TABLESPACE pg_default;


--游戏玩家
create table player( 
    uname character varying(20) primary key,
    upwd character varying(15) NOT NUll default '',
    nickname char(15) NOT NUll default '',
    balance money NOT NUll default 0,
    ulevel smallint NOT NUll default 0
)WITH (
    OIDS = FALSE
)TABLESPACE pg_default;

--游戏
create table game(
    gname character varying(40) primary key,
    price money NOT NUll default 20
)WITH (
    OIDS = FALSE
)TABLESPACE pg_default;

--社区
create table community(
    cid bigserial primary key, --自动递增的id
    content character varying(156) NOT NUll default '',
    ctype character varying(10) NOT NUll default '',
    ctime timestamp with time zone NOT NUll 
)WITH (
    OIDS = FALSE
)TABLESPACE pg_default;

--游戏发行商 发行 游戏
create table issue(
    gcname character varying(40) references gc(gcname) on update CASCADE on delete CASCADE   ,
    gname  character varying(40) REFERENCES game(gname) on update CASCADE on delete CASCADE ,
    itime timestamp  with time zone NOT NUll , --发行时间
    iinfo character varying(256) NOT NUll default '', --发行信息
    itype character varying(20) NOT NUll default '', --游戏所属类型
    primary key(gcname,gname)
)WITH (
    OIDS = FALSE
)TABLESPACE pg_default;

--玩家 购买 游戏
--lpt : last play time 最近运行时间
--bt : 购买时间
--btype : 购买方式 : 微信, 支付宝, visa, 比特币...
--sid : 序列号, 自动增长
create table own(
    uname character varying(20) references player(uname) on update CASCADE on delete CASCADE,
    gname character varying(40) references game(gname) on update CASCADE on delete CASCADE ,
    lrt timestamp  with time zone NOT NUll,  
    bt timestamp  with time zone NOT NUll,
    btype  character varying(10) NOT NUll,
    bprice money NOT null default 0,
    sid bigserial not null,
    primary key(uname,gname)
)WITH (
    OIDS = FALSE
)TABLESPACE pg_default;

--社区新闻属于游戏
create table com_con(
    cid bigserial references community(cid) on update CASCADE on delete CASCADE,
    gname character varying(40) references game(gname) on update CASCADE on delete CASCADE ,
    primary key(cid,gname)
)WITH (
    OIDS = FALSE
)TABLESPACE pg_default;


--游戏发行商
--'' : 字符串
insert into gc(gcname,address,info) values ('NewWorld', 'American','小公司');
insert into gc(gcname,address,info) values ('Valve', 'American','黑心公司');
insert into gc(gcname,address,info) values ('Paradox', 'Sweden','沙盒类游戏');
insert into gc(gcname,address,info) values ('Crytek', 'Germany frankfurt','深受盗版困扰');
insert into gc(gcname,address,info) values ('Petroglyph', 'American Las Vegas','生于西木');
insert into gc(gcname,address,info) values ('505 Games', 'Italy','知名游戏厂商');
insert into gc(gcname,address,info) values ('2K','Boston/Ausralis','知名游戏厂商');
insert into gc(gcname,address,info) values  ('Ubisoft','France','大型跨国公司');
insert into gc(gcname,address,info) values    ('Activision', 'American', '第三方视频游戏');
insert into gc(gcname,address,info) values    ('Running With Scissors', 'American', '没有资料');
--游戏
insert into game(gname,price) values ('Insurgency', 20);
insert into game(gname,price) values ('CSGO',50);
insert into game(gname,price) values ('Destiny2', 40);
insert into game(gname,price) values ('Dota2',0);
insert into game(gname,price) values ('8Bit Invaders',5);
insert into game(gname,price) values ('Assass',10);
insert into game(gname,price) values ('BioShock',10);
insert into game(gname,price) values ('Crysis',9);
insert into game(gname,price) values ('Europa',15);
insert into game(gname,price) values ('ABZU',8);
insert into game(gname,price) values ('POSTAL2',1);


--游戏玩家
insert into player(uname, upwd, nickname,balance, ulevel) values ('dhg', '000000', '黑子',0,0);
insert into player(uname, upwd, nickname,balance, ulevel) values ('cy', '123456', '深海哥',100,3);
insert into player(uname, upwd, nickname,balance, ulevel) values ('wby', 'password', 'by',0,2);
insert into player(uname, upwd, nickname,balance, ulevel) values ('jty', 'root', '欢乐多',50,3);
insert into player(uname, upwd, nickname,balance, ulevel) values ('yc', 'theworld', 'clearlove9',0,10);
insert into player(uname, upwd, nickname,balance, ulevel) values ('shx', '111111', '猫',1000,10);
insert into player(uname, upwd, nickname,balance, ulevel) values ('yl', 'qwerty', '操作怪',200,7);
insert into player(uname, upwd, nickname,balance, ulevel) values ('jxw', 'root', '深蓝的海',0,4);
insert into player(uname, upwd, nickname,balance, ulevel) values ('gf', 'sunshine', '操作dota',10000,100);
insert into player(uname, upwd, nickname,balance, ulevel) values ('lyl', '1234567', '交响乐',0,90);

--社区
--current_timestamp 当前时间
insert into community(content,ctype,ctime) values ('CSGO更新九头蛇大行动','更新', '2019-11-27 05:47');
insert into community(content,ctype,ctime) values ('叛乱:沙尘暴Steam大奖提名','新闻', '2019-11-27 02:08');
insert into community(content,ctype,ctime) values ('CSGO重大更新 裂网大行动','重大更新', '2019-11-20 02:12');
insert into community(content,ctype,ctime) values ('CSGO全新 "裂网大行动"任务已解锁','更新', '2019-11-27 05:47');
insert into community(content,ctype,ctime) values ('CSGO更新 创意工坊提交流程更新','新闻', '2019-4-3 08:46');
insert into community(content,ctype,ctime) values ('叛乱:沙尘暴 游戏开发近况','新闻', '2019-10-30 03:42');
insert into community(content,ctype,ctime) values ('叛乱:沙尘暴 新游戏模式Frontline','新闻', '2019-9-27 22:31');
insert into community(content,ctype,ctime) values ('Bungie每周快报万圣节糖果','新闻', '2019-11-15 07:04');
insert into community(content,ctype,ctime) values ('CS20 Submission Deadline Extended','新闻', '2019-9-24 05:14');
insert into community(content,ctype,ctime) values ('CSGO Watch the Berlin Major Championship','现场直播', '2019-9-9 03:30');

--游戏发行商 发行 游戏 
-- FPS : 第一人称射击
-- RTS : 即使战略
insert into issue(gcname,gname,itime,iinfo,itype) values ('NewWorld','Insurgency', '2014-1-23','Take to the streets for intense close quarters combat','FPS');
insert into issue(gcname,gname,itime,iinfo,itype) values ('Valve','CSGO','2012-8-22','CSGO延续了 1999 年原作在团队竞技类游戏上取得的成就','FPS');
insert into issue(gcname,gname,itime,iinfo,itype) values ('Valve','Dota2','2013-7-9','Dota 2已真正地焕发了生命','Free');
insert into issue(gcname,gname,itime,iinfo,itype) values ('Paradox','Europa','2013-8-13','The empire building game','沙盒');
insert into issue(gcname,gname,itime,iinfo,itype) values ('505 Games','ABZU','2016-8-2','ABZÛ 是一款能唤起潜水梦想的唯美海底冒险游戏','音乐');
insert into issue(gcname,gname,itime,iinfo,itype) values ('2K','BioShock','2013-3-26','Booker must rescue Elizabeth, a mysterious girl imprisoned since childhood and locked up in the flying city of Columbia.','FPS,steampunk');
insert into issue(gcname,gname,itime,iinfo,itype) values ('Crytek','Crysis','2007-11-13','纳米服,自定义武器部件','FPS');
insert into issue(gcname,gname,itime,iinfo,itype) values ('Petroglyph','8Bit Invaders','2016-12-17','PC 平台的快节奏、复古式即时战略游戏','RTS');
insert into issue(gcname,gname,itime,iinfo,itype) values ('Ubisoft','Assass','2015-3-10','游戏《刺客信条：叛变》将带来《刺客信条》系列至今最黑暗的篇章','开放世界');
insert into issue(gcname,gname,itime,iinfo,itype) values ('Activision','Destiny2','2019-10-1','进入命运2的免费游戏世界来体验第一人称射击战斗','FPS');
insert into issue(gcname,gname,itime,iinfo,itype) values ('Running With Scissors','POSTAL2','2003-4-13','到底哪里出错了呢？','FPS');

--玩家购买游戏
-- bprice 代表免费游戏
insert into own(uname,gname,lrt,bt,btype,bprice) values ('wby','CSGO','2019-11-30','2017-6-10','微信',50);
insert into own(uname,gname,lrt,bt,btype,bprice) values ('wby','Dota2','2019-3-10','2018-12-10','steam',0);
insert into own(uname,gname,lrt,bt,btype,bprice) values ('wby','BioShock','2019-8-1','2019-1-2','微信',50);
insert into own(uname,gname,lrt,bt,btype,bprice) values ('wby','Assass','2018-1-29','2017-12-2','微信',30);
insert into own(uname,gname,lrt,bt,btype,bprice) values ('wby','Crysis','2018-3-12','2017-10-4','微信',99);
insert into own(uname,gname,lrt,bt,btype,bprice) values ('cy','CSGO','2019-11-2','2018-9-1','steam',0);
insert into own(uname,gname,lrt,bt,btype,bprice) values ('cy','BioShock','2019-1-1','2018-12-7','微信',50);
insert into own(uname,gname,lrt,bt,btype,bprice) values ('cy','Dota2','2019-4-5','2019-4-5','steam',0);
insert into own(uname,gname,lrt,bt,btype,bprice) values ('cy','Destiny2','2019-8-29','2019-8-1','steam',0);
insert into own(uname,gname,lrt,bt,btype,bprice) values ('jty','CSGO','2019-11-2','2019-3-5','steam',0);

--游戏发布社区新闻
insert into com_con(gname,cid) values('CSGO',1);
insert into com_con(gname,cid) values('CSGO',3);
insert into com_con(gname,cid) values('CSGO',4);
insert into com_con(gname,cid) values('CSGO',5);
insert into com_con(gname,cid) values('CSGO',9);
insert into com_con(gname,cid) values('CSGO',10);
insert into com_con(gname,cid) values('Insurgency',2);
insert into com_con(gname,cid) values('Insurgency',6);
insert into com_con(gname,cid) values('Insurgency',7);
insert into com_con(gname,cid) values('Destiny2',8);


--简单查询
select (gcname,gname) from issue ;
select * from own ;
--选择表中若干数组
select (gcname,gname) from issue where gcname='Valve';
select * from own where uname='wby';

--对查询结果排序
select (uname,ulevel) from player order by ulevel;
select (uname,gname,lrt,bt) from own where uname='cy' order by lrt-bt;

--计算函数汇总
select SUM(bprice) from own where uname='cy';

--分组计算
select (uname, Max(bprice)) from own group by uname;

--多表连接查询
select * from own  inner join player on own.uname=player.uname; 
select (own.uname,own.gname,own.bprice,own.btype) from own cross join game  where own.bprice=game.price;

--9
insert into own(uname,gname,lrt,bt,btype,bprice) values ('wby','POSTAL2');
insert into game(gname,price) values ('Insurgency:SandStorm', 50);

--
update player SET upwd='gog$name' where uname='yc';
update player SET nickname='赌怪',upwd='thereisnopoint' where uname='cy';


delete from own ;
delete from issue where itype='FPS';