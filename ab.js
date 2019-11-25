/*
    α-β剪枝法伪代码
    node : 该节点
    depth : 深度
    a : α的值
    b : β的值
*/
const palyer0 = 0;
const player1 = 1;
const nothing = 2;
const cb_len = 15; //棋盘大小为 15x15
//棋盘用Map
var cb = new Map()
for(var i=0;i<cb_len;i++){
    cb[i] = new Map()
    for(var j=0;j<cb_len;j++){
        cb[i][j] = nothing
    }
}
Map.prototype.get = function (key) {
    if (key in this) {
        return this[key]
    } else {
        return nothing;
    }
}

function AlphaBeta(node, depth, a, b, player)
    if depth == 0 || node is the end status
        return eval_value(node, player)
    //end of  if depth ==0 || node is the end 
    if player == MaxPlayer
        for each child of node
            a = max(a, AlphaBeta(child, depth - 1, a, b, !palyer))
            if b <= a // beta剪枝
                break
        return a
    
    else
        for each child of node
            b = min(b, AlphaBeta(child, depth - 1, a, b, !palyer))
            if b <= a
                break
        return b

function max(a,b) {
    if (a<b) {
        return b
    }
    return a
}
function min(a,b) {
    if (a>b) {
        return b
    }
    return a
}
//分数数组
const score = [10000,9050,9020,9020,9020,9020,9020,
    6000,6000,6000,6000,6000,6000,
    6000,6000,6000,
    4000,4000,40004000
]
//匹配模式
const pattern = [
    [/0{5}/,
    /0{4}2/,
    /10{4}2/,
    /0{3}20/, 
    /00200/,
    /20002/,
    /200202/,
    /100022/,
    /100202/,
    /102002/,
    /200220/,
    /02020/,
    /1200021/,
    /220022/,
    /20202/,
    /202202/,
    /100222/,
    /102022/,
    /102202/,
    /02220/]
    ,
    [/1{5}/,
    /21{4}2/,
    /01{4}2/,
    /1{3}21/,
    /11211/,
    /21112/,
    /211212/,
    /011122/,
    /011212/,
    /012112/,
    /211221/,
    /12121/,
    /0211120/,
    /221122/,
    /21212/,
    /212212/,
    /011222/,
    /012122/,
    /012212/,
    /12221/
    ]
]
//注意这里的是字符
// pattern[0] = /(0{5})|1{5}/ //连五
// pattern[1] = /2(1{4}|0{4})2/ //活四
// pattern[2] = /(01{4})|(10{4})2/ //冲四
// pattern[3] = /(1{3}|0{3})2(1|0)/ //冲四
// pattern[4] = /(11211)|(00200)/ //冲四
// pattern[5] = /(21112)|(20002)/ //活三
// pattern[6] = /(211212)|(200202)/ //活三
// //眠三
// pattern[7] = /(011122)|(100022)/
// pattern[8] = /(011212)|(100202)/
// pattern[9] = /(012112)|(102002)/
// pattern[10] = /(211221)|(200220)/
// pattern[11] = /(12121)|(02020)/
// pattern[12] = /(1200021)|(0211120)/
// //眠三结束
// //活二
// pattern[13] = /22((1{2})|(0{2}))22/
// pattern[14] = /2((121)|(020))2/
// pattern[15] = /2((0220)|(1221))2/
// //活二结束
// //眠二
// pattern[16] = /((100)|(011))222/
// pattern[17] = / ((1020)|(0121))22/
// pattern[18] = /(012212)|(102202)/
// pattern[19] = /(02220|12221)/
// //眠二结束

// 返回node的评估值
function eval_value(node, player)
    var value= 0;
    x = node.x
    y = node.y
    cb[x][y] = player
    var lines = getAllWays(x,y)
    for i in lines
        const line = lines[i]
        for p in pattern
            for pa in pattern[p]
                const patt = pattern[p][pa]
                var num = line.match(patt)
                if num != null
                    if pa == player
                        value += score[pa]*num
                    else
                        value-= score[pa]*num
    //end for i in getAllWays(x,y)
    cb[x][y] = nothing
    return value
// end of function eval_value

function getAllWays(x, y)
    var res = new Array(4).fill([])
    var i = 0;
    var x1, y1;
    for (i = -4; i <= 4; i++) {
        x1 = x + i
        y1 = y + i
        res[0].push(cb[x1][y1])
    }
    for (i = -4; i <= 4; i++) {
        x1 = x - i
        y1 = y + i
        res[1].push(cb[x1][y1])
    }
    x1 = x
    for (i = -4; i <= 4; i++) {
        y1 = y + i
        res[2].push(cb[x1][y1])
    }
    y1 = y
    for (i = -4; i <= 4; i++) {
        x1 = x + i
        res[3].push(cb[x1][y1])
    }
    return res.toString().split(',') //返回字符串数组
//end of getAllWays
//参考资料
/*
https://blog.csdn.net/marble_xu/article/details/90450436
*/


// α-β减枝法 弃用的
/*
    伪代码
    none // 空
    black // 黑子
    white //白子
    cb //棋盘nxn
    alpha_beta(alpha, beta, h,player) 
        
        alpha = -∞
        beta = +∞
        
        ansx,ansy // variable
        if h==Limit_Deep || end of game
                return eval(player) - eval(player^1)
        if player //自己
            for i=1 to n
                for j =1 to n
                    if cb[i][j] == none
                        cb[i][j] = black_white(player)
                        ans = alpha_beta(alpha,beta,h+1,player^1)
                        cb[i][j] = none
                        if ans > alpha
                            alpha = ans
                            ansx = i
                            ansy= j
                        if alpha ≥ beta
                            return alpha
            return alpha
        else //对手
            for i=1 to n
                for j =1 to n
                    if cb[i][j] == none
                        cb[i][j] = black_white(player)
                        ans = alpha_beta(alpha,beta,h+1,player^1)
                        cb[i][j] = none
                        if ans < beta
                            beta = ans
                        if alpha ≥ beta
                            return beta
            return beta
    //alpha_beta end
    black_white(player)
        //根据palyer 返回black或者white
        return ...//待定
    //black_white end

    //eval begin
    eval(player)
        // f(p) = （所有空格放上我方棋子后，n子连线的总个数）-（所有空格放上对方棋子后，n子连线的总个数）
        x = black_beta(player)
        ans = 0
        //横 + 竖
        for i =1 to n
            w = 0
            for j =1 to n
                if cb[i][j] == x || cb[i][j] == none || cb[j][i] == x
                    w++
            if w==m
                ans++
        

        //正对角线 + 反对角线
        w = 0
        for i =1 to n
            if cb[i][i] ==x || cb[i][i] == none || cb[i][n-i+1] == x || cb[i][n-i+1] == none
                w++
        if w==m
            ans++
        return ans

    //eval end
    */

/*
function Point() {
    this.a = -1
    this.b = -1
}

function CheckBoard() {
    this.checkboard = [] //nxn矩阵
    this.player = -1
    this.LimitDeep = 6
    this.alpha_beta = function (alpha, beta, h, player) {
        
        var n = this.checkboard.length
        if (h == this.LimitDeep || ) {
            return this.evalFunction(player) - this.evalFunction(player ^ 1)
        }
        var ansx,ansy;
        if (player) {
            for (let i = 0; i < n; i++) {
                for (let j = 0; j < n; j++) {
                    if(this.checkboard[i][j] == null) {
                        
                        this.checkboard[i][j] = this.black_white(player)
                        var ans = this.alpha_beta(alpha, beta, h+1, player^1)
                        this.checkboard[i][j] = null
                        if (ans < beta) {
                            beta = ans
                            ansx = i
                            ansy = j
                        }
                        if(alpha >= beta) {
                            return [alpha,ansx,ansy]
                        } 
                    }
                }
            }
            return [alpha,ansx,ansy]
        } else {
            for (let i = 0; i < n; i++) {
                for (let j = 0; j < n; j++) {
                    if(this.checkboard[i][j] == null) {
                        
                        this.checkboard[i][j] = this.black_white(player)
                        var ans = this.alpha_beta(alpha, beta, h+1, player^1)
                        this.checkboard[i][j] = null
                        if(ans < beta) {
                            beta = ans
                            ansx = i
                            ansy = j
                        }
                        if(alpha >= beta) {
                            return [beta,ansx,ansy]
                        }
                    }
                }
            }
            return [beta,ansx,ansy]
        }
    }
    this.evalFunction = function () {
        return 0
    }
    this.black_white = function (player) {
        //判断palyer用的什么棋子
        return player
    }
}
*/