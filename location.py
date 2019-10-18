import json
import math
import requests
class Location:
    id = -1 #unique id
    lng = 0
    lat = 0
    adj_locations = []
    def __init__(self, args={}):
        self.id = args["id"]
        self.lng = args["lng"]
        self.lat = args["lat"]

class minQueue:
    pass
class Graph:
    martix = []

    #根据json文件构造二维矩阵
    def getMatrix(self,file_path):
        location_data = None
        with open(file_path) as f:
            location_data = json.load(f)
        location_len_list = range(len(location_data))
        self.martix = [_ for _ in location_len_list]
        for i in location_len_list:
            tmp = [INF for _ in location_len_list]
            self.martix[i] = tmp
        #矩阵初始化结束

        url_head = 'https://restapi.amap.com/v3/distance'
        tmp_url = "&origin="
        #然后通过web获得距离
        for i in location_len_list:
            ii = str(i)

            tmp_url+=  str(location_data[ii][0])+','+str(location_data[ii][1]) +'|'
        tmp_url = tmp_url[:-1]
        # print("tmp url",tmp_url)
        #计算距离
        for ii in location_len_list:
            i = str(ii)
            destination = "&destination=" + str(location_data[i][0] ) +','+ str(location_data[i][1])
            print("url",url_head+tmp_url+destination + '&type=1')
            map_data = {
                'key': '76a251e04fd7611225f906c85b2a0163',
                'origins':tmp_url,
                'destination':destination
            }
            data = json.loads(requests.get(url_head, map_data ).text)
            print(data)
            for j in data:
                a = data[j]['origin_id'] - 1
                self.martix[a][ii] = data[j]['distance']
                self.martix[ii][a] = data[j]['distance']
        #写入文件
        with open("distance.json","w") as f:
            json.dump( json.dumps(self.martix),f)

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
    #dijkstra
    def dijkstra(self,start,end):
        pass

INF = math.inf



if __name__ == '__main__':

    """https://restapi.amap.com/v3/distance?origins=116.481028,39.989643|114.481028,39.989643|115.481028,39.989643&destination=114.465302,40.004717&output=json&key=76a251e04fd7611225f906c85b2a0163"""
    # location_data = None
    #
    # with open('location.json') as f:
    #     location_data = json.load(f)
    # location_len_list = range( len(location_data) )
    # matrix = [_ for _ in location_len_list]
    # for i in location_len_list:
    #     tmp = [INF for _ in location_len_list]
    #     matrix[i] = tmp
    # #初始化矩阵结束
    g = Graph()
    g.getMatrix('location.json')

    """"
    
    
    
    """