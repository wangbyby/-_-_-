// α-β减枝法
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