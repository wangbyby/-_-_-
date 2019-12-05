const maxDepth = 5
const palyer0 = 0;
const player1 = 1;
const nothing = 2;
const cb_len = 15; //棋盘大小为 15x15
var cb_empty = cb_len * cb_len;
//棋盘用Map
Map.prototype.get = function (key) {
    if (key in this) {
        return this[key]
    } else {
        return nothing;
    }
}
var cb = new Map()
for (var i = 0; i < cb_len; i++) {
    cb[i] = new Map()
    for (var j = 0; j < cb_len; j++) {
        cb[i][j] = nothing
    }
}
//分数数组
const score = [10000, 9050, 9020, 9020, 9020, 9020, 9020,
    6000, 6000, 6000, 6000, 6000, 6000,
    6000, 6000, 6000,
    4000, 4000, 40004000
]
//匹配模式
//注意这里的是字符
const pattern = [
    [/0{5}/, /*连五*/ /0{4}2/,/*活四*/ 
    /10{4}2/, /0{3}20/,   /00200/,/*冲四*/  
    /20002/,  /200202/,/*活三*/  
    /100022/,/100202/,/102002/,/200220/, /02020/, /1200021/, /*眠三*/
    /220022/, /20202/, /202202/, /*活二*/
    /100222/,/102022/, /102202/, /02220/ /*眠二*/
    ],
    [/1{5}/, /21{4}2/,
    /01{4}2/,/1{3}21/,/11211/,
    /21112/,/211212/,
    /011122/,/011212/,/012112/,/211221/,/12121/,/0211120/,
    /221122/,/21212/,/212212/,
    /011222/,/012122/,/012212/,/12221/ 
    ]
]



//a -- α
//b -- β
/*
伪代码描述
function AlphaBeta(node, depth, a, b, player)
    if depth == 0 || node is the end status
        return eval_value(node, player)
    //end of  if depth ==0 || node is the end 
    if player == MaxPlayer
        for each child of node
            a = max(a, AlphaBeta(child, depth - 1, a, b, !palyer))
            if b <= a // β剪枝
                break
        return a
    else
        for each child of node
            b = min(b, AlphaBeta(child, depth - 1, a, b, !palyer))
            if b <= a // α剪枝
                break
        return b
*/
function AlphaBeta(node, depth, a, b, player) {
    if (depth == 0 || node is the end status)
        return eval_value(node, player)
    //end of  if depth ==0 || node is the end 
    if (player == MaxPlayer) {
        for each child of node
            a = max(a, AlphaBeta(child, depth - 1, a, b, !palyer))
            if (b <= a) // β剪枝
                break;
        return a
    } else {
        for each child of node
            b = min(b, AlphaBeta(child, depth - 1, a, b, !palyer))
            if (b <= a) // α剪枝
                break
        return b
    }
}

function max(a, b) {
    if (a < b) {
        return b
    }
    return a
}

function min(a, b) {
    if (a > b) {
        return b
    }
    return a
}
// 返回node的评估值
function eval_value(node, player) 
    var value = 0
    x = node.x
    y = node.y
    cb[x][y] = player
    var lines = getAllWays(x, y)
    for _, line := range lines
        for k1,v1 := range pattern
            for k2,v2 := range v1
                const patt = v2
                var num = line.match(patt)
                if num != null
                    if k1 == player
                        value += score[k2] * num
                    else
                        value -= score[k2] * num
    cb[x][y] = nothing
    return value

// end of function eval_value
function getAllWays(x, y) {
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
}
//end of getAllWays

var win = 0;
function Play() {
    var a = -Infinity
    var b = Infinity

    while ( win==0) {
        
    }
}
function CBNode() {
}
CBNode.prototype.x = -1;
CBNode.prototype.y = -1;
// α-β减枝法 弃用的
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


// /*深度优先伪代码*/
// DFS(node)
//     foreach child of node
//         DFS(child)

// //极大极小伪代码
// minmax(node, depth)
//     if node is terminal or depth ==0
//         return eval_value(node)
//     if 对手
//         let b = +∞
//         foreach child of node
//             b = min(a,minmax(child,depth-1))
//         return b
//     else 
//         let a = -∞
//         foreach child of node
//             a = max(b,minmax(child,depth-1))
//         return a