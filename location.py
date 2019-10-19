import json
import math
import requests
INF = math.inf
class minQueue:
    queue = []
    f = {}
    def push_back(self, element):
        #向头部添加元素
        self.queue.insert(element, 0)
        self.shiftdown(0)
    def shiftdown(self,i):
        left = i<<1 + 1
        right = i<<1 + 2
        mini = i
        if left < len(self.queue) and self.f[self.queue[left]] < self.f[self.queue[mini]]:
            mini = left
        if right <  len(self.queue) and self.f[self.queue[right]] < self.f[self.queue[mini]]:
            mini = right
        if mini != i:
            self.queue[mini], self.queue[i] = self.queue[i], self.queue[mini]
            self.shiftdown(mini)
    def pop_head(self):
        tmp = self.queue[0]
        self.queue = self.queue[1:]
        self.shiftdown(0)
        return tmp
class Graph:
    martix = [] #二维数组
    distance_file_path = "distance.json" #距离文件路径
    location_data = None #位置
    distance = [] #距离矩阵
    #根据json文件构造二维矩阵
    def getMatrix(self,file_path):

        with open(file_path) as f:
            self.location_data = json.load(f)
        location_len_list = range(len(self.location_data))
        self.martix = [_ for _ in location_len_list]
        for i in location_len_list:
            tmp = [INF for _ in location_len_list]
            for j in location_len_list:
                if j in self.location_data[str(i)][2]:
                    tmp[j] = -1
            self.martix[i] = tmp
        #矩阵初始化结束

        url_head = 'https://restapi.amap.com/v3/distance?key=76a251e04fd7611225f906c85b2a0163'
        tmp_url = "&origins="
        #然后通过web获得时间(以时间代替路程)
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
                if self.martix[a][ii] != INF:
                    self.martix[a][ii] = j['duration']
                    self.martix[ii][a] = j['duration']


        #写入文件
        with open(self.distance_file_path,"w") as f:
            json.dump( json.dumps(self.martix),f)

    def read_from_file(self):

        with open(self.distance_file_path) as f:
            self.distance = json.loads(json.load(f))

    #dijkstra
    def dijkstra(self,start,end):
        pass
    def search(self, points = []):
        #方法描述       #根据points列表查询两点之间的距离
        #返回值 points数组点之间的路径
        pass



if __name__ == '__main__':


    g = Graph()
    # g.getMatrix('location.json')
    #g.read_from_file()
    minq = minQueue()
    mi