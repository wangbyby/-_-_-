//D 为二维数组(nxn矩阵)
function TSP(D) { //起点为0
    const INF = 10000 //定义的最大值
    var n = D.length // n的个数
    var i, j, k, min, tmp;
    var b = 1 << (n - 1); // 点集状态总数
    var dp = {} // 记录状态
    var bridge = {} //记录中间节点
    for (i = 0; i < n; i++) { //初始化dp与bridge
        dp[i] = {}
        bridge[i] = {}
        for (j = 0; j < b; j++) {
            dp[i][j] = 0
            bridge[i][j] = INF
        }
    }
    for (i = 0; i < n; i++) { //初始化 dp的第0列
        dp[i][0] = D[i][0]
    }
    //init end

    //遍历二维数组即遍历dp
    for (i = 1; i < b - 1; i++) {
        for (j = 1; j < n; j++) {
            if ((1 << (j - 1) & i) == 0) {
                //点j未访问
                min = INF
                for (k = 1; k < n; k++) { //遍历点集
                    if (1 << (k - 1) & i) {
                        //点k在集合中

                        // 松弛操作
                        tmp = D[j][k] + dp[k][i - (1 << (k - 1))]
                        if (tmp < min) {
                            min = tmp
                            dp[j][i] = min
                            bridge[j][i] = k
                        }
                    }
                }
            }
        }
    }
    //处理最后一列
    min = INF
    for (k = 1; k < n; k++) {
        // 松弛操作
        tmp = D[0][k] + dp[k][b - 1 - (1 << (k - 1))] //b-1-(1<<(k-1)) :  去掉k节点
        if (tmp < min) {
            min = tmp
            dp[0][b - 1] = min
            bridge[0][b - 1] = k
        }
    }
    var mincost = dp[0][b - 1]
    var path = [0]
    for (i = b - 1, j = 0; i > 0;) {
        j = bridge[j][i] // 下一个节点
        i = i - (1 << (j - 1))
        path.push(j)
    }
    path.push(0)
    //返回值说明 path为路径, mincost为最短花费
    return [path, mincost]
}


var m = [
    [0, 7, 6, 10, 13],
    [7, 0, 7, 10, 10],
    [6, 7, 0, 5, 9],
    [10, 10, 5, 0, 6],
    [13, 10, 9, 6, 0]
]


var res = TSP(m)
console.log("最短路径 : ", res[0])
console.log("最少花费 : ", res[1])

var n = 23
var m = new Array(n)
for (var i = 0; i < n; i++) {
    m[i] = new Array(n).fill(0)
}
var s1 = new Date().getTime()
res = TSP(m)
var s2 = new Date().getTime()
console.log("花费时间(单位:ms):",s2 - s1)


// for i = 1 to n
//     for j = 0 to M
//         for x = 0;x * w[i] <= j;x++
//             m[i][j] = max(
//                 m[i - 1][j],
//                 m[i - 1][j - x * w[i]] + x * v[i]
//             )
function GKP(M, m, v) {
    var n = m.length

    /*
        M:背包最大承重量
        n:物品种类个数
        m:物品质量
    */
    var i, j, x;
    var dp = new Array(n + 1)
    for (i = 0; i <= n; i++) {
        dp[i] = new Array(M + 1).fill(0)
    }
    var numGoods = new Array(n + 1).fill(0)
    for (i = 1; i <= n; i++) {
        for (j = 0; j <= M; j++) {
            var count = j / m[i - 1]
            for (x = 0; x <= count; x++) {

                var tmp = dp[i - 1][j - x * m[i - 1]] + x * v[i - 1]
                if (dp[i][j] < tmp) {
                    dp[i][j] = tmp

                }
            }
        }
    }
    console.log("动态规划表", dp)
    console.log("最大值=", dp[n][M])
    j = M;
    i = n;
    //构造解路径
    while (j >= 0 && i > 0) {
        // console.log(" (i,j) = ",[i,j])
        if (dp[i][j] == dp[i - 1][j]) {
            i--
        } else {
            j -= m[i - 1]
            numGoods[i]++

        }
    }
    numGoods = numGoods.splice(1)
    console.log("解=", numGoods)
    return [dp[n][M], numGoods]
}
var M = 10

var m = [1, 6, 4, 3]
var v = [1, 3, 2, 6]
GKP(M, m, v)

M = 230
m = [20,25,40,12,31]
v = [1,2,3,1,5]
GKP(M, m, v)