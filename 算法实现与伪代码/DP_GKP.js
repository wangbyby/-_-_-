/*
    *1.
    函数输入说明
        M:背包最大承重量
        v:价值
        m:物品质量
    //向量 m与v 下标 1 to n
    GKP(M,m,v)
        let n = m.length //物品种类个数
        for i=1 to n 
            for j = 0 to M
                let count = j/m[i] 
                for k=0 to count
                    dp[i][j] = dp[i-1][j]
                    let tmp = dp[i - 1][j - k * m[i]] + k * v[i]
                    if dp[i][j] < tmp
                        dp[i][j] = tmp
                        path[(i,j)] = 1
        let j = M
        let i = n
        let x  be a new Vector// 下标范围 [1,n]
        while i>0 && j>0
            if path[(i,j)] ==1
                x[i]++
                j = j- m[i]
            else 
                i = i-1
        return dp[n][M] , x    
    *时间复杂度: O(n*M*M/(min(m)))
    *空间复杂度 : O(nM)
    *2.
    GKP(M,m,v) 
        
        for i = 1 to n
            for j = m[i] to M
                let tmp = f[j-m[i]] + v[i]
                if f[j] < tmp
                    f[j] = tmp
                    path[(i,j)]  = 1
        let j = M
        let i = n
        let x  be a new Vector// 下标范围 [1,n]
        while i>0 && j>0
            if path[(i,j)] ==1
                x[i]++
                j = j- m[i]
            else 
                i = i-1
        return f[M] , x 
    *时间复杂度: O(n*M)
    *空间复杂度 : O(M)
    */



function GKP0(M, m, v) {
    var n = m.length
    var i, j, k;
    var dp = new Array(n + 1)
    for (i = 0; i <= n; i++) {
        dp[i] = new Array(M + 1).fill(0)
    }
    var x = new Array(n + 1).fill(0)
    var path = {}
    for (i = 1; i <= n; i++) {
        for (j = 0; j <= M; j++) {
            var count = j / m[i - 1]
            for (k = 0; k <= count; k++) {
                dp[i][j] = dp[i - 1][j]
                var tmp = dp[i - 1][j - k * m[i - 1]] + k * v[i - 1]
                if (dp[i][j] < tmp) {
                    dp[i][j] = tmp
                    path[[i, j]] = 1
                }
            }
        }
    }
    //解路径
    j = M
    i = n
    while (i > 0 && j > 0) {
        if (path[[i, j]] == 1) {
            x[i - 1]++
            j = j - m[i - 1]
        } else {
            i = i - 1
        }
    }
    x.pop()
    console.log("解=", x)
    return [dp[n][M], x]
}


function GKP(M, m, v) {
    var n = m.length
    
    var i, j;
    var x = {} //解向量
    for (i = 0; i <= n; i++) {
        x[i] = 0
    }
    var path = new Map()

    let f = new Array(M + 1).fill(0)
    for (i = 1; i <= n; i++) {
        /*
        时间复杂度 : O(nM)
        空间复杂度: O(M)
        */
        for (j = m[i - 1]; j <= M; j++) {
            var tmp = f[j - m[i - 1]] + v[i - 1]
            if (f[j] < tmp) {
                f[j] = tmp
                path[[i, j]] = 1
            }
        }
    }
    //解路径
    j = M
    i = n
    while (i > 0 && j > 0) {
        if (path[[i, j]] == 1) {
            x[i - 1]++
            j = j - m[i - 1]
        } else {
            i = i - 1
        }
    }
    console.log("最少花费 = ", f[M])
    console.log("解向量 = ", x)

    return [f[M], x]
}


//广义背包测试数据
var M = 10
var m = [1, 6, 4, 3]
var v = [1, 3, 2, 6]
console.log(GKP(M, m, v))

M = 230
m = [20, 25, 40, 12, 31]
v = [5, 2, 3, 1, 5]
console.log(GKP(M, m, v))