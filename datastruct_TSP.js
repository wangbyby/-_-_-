const inf = 1000000 //100000代表无穷大

function Result() {}
Result.prototype.matrix = null
Result.prototype.paths = {}

function Edge() {
    this.a = ""
    this.b = ""

    this.copy_self = function () {
        var copy = new Edge()
        copy.a = this.a
        copy.b = this.b
        copy.cost = this.cost
        copy.right = this.right
        copy.left = this.left
        return copy
    }
}
Edge.prototype.cost = 0
Edge.prototype.left = null
Edge.prototype.right = null



function Graph() {
    this.info = []
    this.Vertex = function () {
        var tmp = []
        for (const iterator of this.info.keys()) {
            tmp.push(iterator)
        }
        return tmp
    }

    this.info = function (u, accessed) {
        var tmp = []
        for (var i = 0; i < this.info.length; i++) {
            if (!accessed.includes(i) || accessed.length == 0) {
                var e = new Edge()
                e.a = u
                e.b = i
                e.cost = this.info[u][i]
                tmp.push(e)
            }
        }
        return tmp
    }
    this.visted_unvis = function (accessed, unaccess) {
        var path = []
        for (var index = 0; index < accessed.length; index++) {
            for (var j = 0; j < unaccess.length; j++) {
                var tmp = this.info[accessed[index]][unaccess[j]]
                var e = new Edge()
                e.a = accessed[index]
                e.b = unaccess[j]
                e.cost = tmp
                path.push(e)
            }

        }
        var ee = new Edge()
        ee.cost = inf
        for (var index = 0; index < path.length; index++) {
            if (ee.cost > path[index].cost) {
                ee.cost = path[index].cost
                ee.a = path[index].a
                ee.b = path[index].b
            }

        }
        return ee;
    };
}

function TSP(G, u) {
    var ver = G.Vertex()
    var access = []
    var path = []
    access.push(u)
    var dd = ver.indexOf(u)
    ver.splice(dd, 1)

    var root = new Edge()
    while (ver.length != 0) {
        var e = G.visted_unvis(access, ver)
        //console.log("edge=", e)
        access.push(e.b)
        path.push(e)
        dd = ver.indexOf(e.b)
        ver.splice(dd, 1)

        var ecopy = e.copy_self()

        InsertNode(root, ecopy)
    }

    console.log("root=", root)
    var hamitonTree = []

    PreOrder2(root, hamitonTree)
    hamitonTree[0] = u
    hamitonTree.push(u)
    console.log("hamiton=", hamitonTree)
    return hamitonTree
}

//寻找前驱节点
function FindPreNode(nowNode, findNode) {
    if (nowNode == null) {
        return null
    }
    if (nowNode.b == findNode.a) {
        return nowNode
    } else {
        return FindPreNode(nowNode.left, findNode) || FindPreNode(nowNode.right, findNode)
    }
}

function InsertNode(root, node) {
    var t = FindPreNode(root, node)
    if (t == null) {
        root = node
        return
    }
    if (t.left == null) {
        t.left = node
        return
    }
    if (t.right == null) {
        t.right = node
        return
    }
}

function PreOrder2(root, l) {
    if (root != null) {
        l.push(root.b)
        PreOrder2(root.left, l)
        PreOrder2(root.right, l)
    }
}



function Dijkstra(m, start, end) { //Dijkstra
    var n = m.length
    var Q = new minQueue()
    var pre = {}
    for (var i = 0; i < n; i++) {
        Q.f[i] = inf
    }
    Q.f[start] = 0
    Q.push_back(start)
    while (Q.list.length != 0) {
        var u = Q.pop_head()
        // console.log("u=", u)
        for (var j = 0; j < n; j++) {
            if (m[u][j] == inf) {
                continue
            }
            if (Q.f[j] > Q.f[u] + m[u][j]) {
                Q.f[j] = Q.f[u] + m[u][j]
                pre[j] = u
                Q.push_back(j)
            }
        }
        // console.log("Queue=", Q.list)
        // console.log("pre=", pre)
    }
    var resPath = getPath(pre, end).reverse()
    //获得权重
    var weight = 0 //此时weight单位为 s(秒)
    for(var i=0; i<resPath.length - 1;i++){
        weight += m[resPath[i]][resPath[i+1]]
    }
    weight = Math.ceil(weight/60)
    return {'path':resPath, 'weight':weight } //此时weight单位为 min(分钟)
}


function getPath(cameFrom, current) {
    var total_path = [current]
    while (cameFrom[current] != undefined) {
        current = cameFrom[current]
        total_path.push(current)
    }
    return total_path
}

function minQueue() {
    this.list = []
    this.f = {}
    this.push_back = function (element) {
        this.list.unshift(element)
        this.shiftdown(0)
    }
    this.Less = function (a, b) {
        return a < b
    }
    this.Swap = function (a, b) {
        [this.list[a], this.list[b]] = [this.list[b], this.list[a]]
    }

    this.pop_head = function () {
        var tmp = this.list[0]
        this.list = this.list.splice(1)

        this.shiftdown(0)
        //delete this.f[tmp]
        return tmp
    }

    this.shiftdown = function (index) {
        let left = index * 2 + 1
        let right = index * 2 + 2
        let min = index
        if (left < this.list.length && this.Less(this.f[this.list[left]], this.f[this.list[index]])) {
            min = left
        }
        if (right < this.list.length && this.Less(this.f[this.list[right]], this.f[this.list[min]])) {
            min = right
        }
        if (min != index) {
            this.Swap(min, index)
            this.shiftdown(min)
        }
    }
}
// exports.FLOYD = FLOYD
// exports.TSP = TSP
// exports.Graph = Graph

//在数组中遍历树
// function PreOrder(array) {
//     for (var index = 0; index < array.length; index++) {
//         const node = array[index];
//         // array.find(x=> x.a == element.b)
//         array.forEach(element => {
//             if (element.a == node.b) {
//                 if (node.left == null) {
//                     node.left = array.indexOf(element)
//                 } else {
//                     node.right = array.indexOf(element)
//                 }
//             } else if (element.a == node.a && array.indexOf(element) > array.indexOf(node)) {
//                 if (node.left == null) {
//                     node.left = array.indexOf(element)
//                 } else {
//                     node.right = array.indexOf(element)
//                 }

//             }
//         });
//     }

//     var h = []
//     h.push(array[0].a)
//     rangeTree(array[0], array, h)
//     h.push(array[0].a)
//     return h
// }

// function rangeTree(node, array, a) {
//     if (node == null) {
//         return
//     }
//     a.push(node.b)
//     if (node.left != null) {
//         rangeTree(array[node.left], array, a)
//     }
//     if (node.right != null) {
//         rangeTree(array[node.right], array, a)
//     }
// }

// function TSP(G, u) { //符合三角不等式

//     return PreOrder(Prim(G, u))
// }

function FLOYD(G) {
    var result = new Result()
    result.matrix = G
    var lenRows = result.matrix.length
    for (var index = 0; index < lenRows; index++) {
        result.paths[index] = {}
    }
    for (var k = 0; k < lenRows; k++) {
        for (var i = 0; i < lenRows; i++) {
            for (var j = 0; j < lenRows; j++) {
                if (result.matrix[i][j] > result.matrix[i][k] + result.matrix[k][j]) {
                    result.matrix[i][j] = result.matrix[i][k] + result.matrix[k][j]
                    result.paths[i][j] = k
                }
            }
        }
    }
    return result
}

