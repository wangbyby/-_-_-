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
    distance_file_path = "distance.json"
    location_data = None
    #根据json文件构造二维矩阵
    def getMatrix(self,file_path):

        with open(file_path) as f:
            self.location_data = json.load(f)
        location_len_list = range(len(self.location_data))
        self.martix = [_ for _ in location_len_list]
        for i in location_len_list:
            tmp = [INF for _ in location_len_list]
            self.martix[i] = tmp
        #矩阵初始化结束

        url_head = 'https://restapi.amap.com/v3/distance?key=76a251e04fd7611225f906c85b2a0163'
        tmp_url = "&origins="
        #然后通过web获得距离
        for i in location_len_list:
            ii = str(i)

            tmp_url += str(self.location_data[ii][0])+','+str(self.location_data[ii][1]) + '|'
        tmp_url = tmp_url[:-1]
        # print("tmp url",tmp_url)
        #计算距离
        for ii in location_len_list:
            i = str(ii)
            destination = "&destination=" + str(self.location_data[i][0] ) +','+ str(self.location_data[i][1])

            url = url_head + tmp_url + destination
            data = json.loads(requests.get(url).text)['results']
            print(data)
            for j in data:
                a = int(j['origin_id']) - 1
                self.martix[a][ii] = j['duration']
                self.martix[ii][a] = j['duration']

        #写入文件
        with open(self.distance_file_path,"w") as f:
            json.dump( json.dumps(self.martix),f)
    # 从文件读数据
    def read_m_from_file(self):
        with open(self.distance_file_path) as f:
            data = json.load(f)
            print(data)
    #dijkstra
    def dijkstra(self,start,end):
        pass

INF = math.inf



if __name__ == '__main__':


    g = Graph()
    #g.getMatrix('location.json')
    g.read_m_from_file()