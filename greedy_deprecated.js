function Graph() {
    this.a = []
    this.TSP = function(uu) { //贪心法
        var u = uu
        var n = this.a.length
        var p = []
        var vis = {}
        for (var i = 0; i < n; i++) {
            vis[i] = false
        }
        vis[u] = true
        while (p.length < n - 1) {
            var e = new Edge()
            e.a = u
            //e["a"] = u
            e.cost = 10000
            //e["w"] = 10000 //max
            for (var i = 0; i < n; i++) {
                if (this.a[u][i] < e.cost && vis[i] == false) {
                    e.cost = this.a[u][i]
                    e.b = i
                }
            }
            vis[u] = true
            p.push(e)
            u = e.b
        }
        var tmp = new Edge()
        tmp["a"] = p[p.length - 1]["b"]
        tmp["b"] = uu
        tmp["cost"] = this.a[uu][tmp["a"]]
        p.push(tmp)
        
        return p
    }
}

function Edge() {}
Edge.prototype.a = ""
Edge.prototype.b = ""
Edge.prototype.cost = 0