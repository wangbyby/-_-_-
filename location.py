class Location:
    id = -1 #unique id
    lng = 0
    lat = 0
    adj_locations = []
    def __init__(self, args={}):
        self.id = args["id"]
        self.lng = args["lng"]
        self.lat = args["lat"]
    
class Graph:
    martix = []

    #根据json文件构造二维矩阵
    def getGraph(self,data):
        pass
    #弗洛伊德算法
    def floyd(self): #a为起点,b为终点 a,b都为id
        n = len(self.martix) # 矩阵的行数
        d = self.martix.copy() # 复制 self

        map = {} #储存path
        range_n = range(0,n)
        for i in range_n:
            map[i] = {}
        for k in range_n:
            for i in range_n:
                for j in range_n:
                    if d[i][j] > d[i][k]+ d[k][j]:
                        d[i][j] = d[i][k] + d[k][j]
                        map[i][j] = k
        return {"m":d, "path":map}
