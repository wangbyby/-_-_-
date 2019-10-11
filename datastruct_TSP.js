function Edge() {
    this.a = ""
    this.b = ""

    this.copy_self = function () {
        let copy = new Edge()
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
    this.Vertex = function () {
        var tmp = []
        for (const iterator of this.info.keys()) {
            tmp.push(iterator)
        }
        return tmp
    }
    this.info = []
    this.adj = function (u, accessed) {
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
        for (let index = 0; index < accessed.length; index++) {
            for (let j = 0; j < unaccess.length; j++) {
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
        for (let index = 0; index < path.length; index++) {
            if (ee.cost > path[index].cost) {
                ee.cost = path[index].cost
                ee.a = path[index].a
                ee.b = path[index].b
            }

        }
        return ee
    }

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
        console.log("edge=", e)
        access.push(e.b)
        path.push(e)
        dd = ver.indexOf(e.b)
        ver.splice(dd, 1)

        var ecopy = e.copy_self()

        InsertNode(root, ecopy)
    }

    console.log("root=", root)
    let hamitonTree = []

    PreOrder2(root, hamitonTree)
    hamitonTree[0] = 0
    hamitonTree.push(0)
    console.log("hamiton=", hamitonTree)
    return hamitonTree
}


const inf = 100000 //100000代表无穷大

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



// function PreOrder(array) {
//     for (let index = 0; index < array.length; index++) {
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

//     let h = []
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
