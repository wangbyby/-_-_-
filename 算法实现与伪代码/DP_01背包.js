// /*伪代码
// M //最大质量
// m //质量数组 下标 1 to n
// v //价值数组 下标 1 to n 
// 时间复杂度 : O(nM)
// 空间复杂度: O(M)
// */
// zero_one_KP(M,m,v) //优化后 sudo code
//     //init
//     let n = m.length
//     let dp =[1,M] //向量[1,M]
//     let path = Map //哈希表
//     //init end
//     //动态规划
//     for i=1 to n //递增
//         for j = M to m[i] //递减
//             let tmp = dp[j-m[i]]+v[i]
//             if dp[j] < tmp
//                 dp[j] = tmp
//                 path[(i,j)] = 1
//     //回溯解
//     let j = M
//     let i = n
//     let x  //解向量 下标 1 to n, 全部初始化为0
//     while j>0&&i>0
//         if path[(i,j)] ==1
//             x[i] = 1
//             j = j - m[i]
//         i = i-1
//     return dp[n+1][M], x


//0-1背包问题
function zero_one_KP(M, m, v) {
    var n = m.length
    /*
        M:背包最大承重量
        n:物品种类个数
        m:物品质量
        时间复杂度: O(nM)
        空间复杂度: O(M)
    */
    var i, j;
    let path = new Map()
    let f = new Array(M + 1).fill(0)
    for (i = 1; i <= n; i++) {
        for (j = M; j >= m[i - 1]; j--) { //从大到小
            var tmp = f[j - m[i - 1]] + v[i - 1]
            if (f[j] < tmp) {
                f[j] = tmp
                path[[i, j]] = 1
            }
        }
    }
    j = M
    i = n
    var x = new Array(n).fill(0)
    while (i > 0 && j > 0) {
        if (path[[i, j]] == 1) {
            x[i - 1] = 1
            j = j - m[i - 1]
        }
        i = i - 1 //注意这里与广义背包不一样
    }

    return [f[M], x]
}