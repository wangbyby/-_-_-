import json

import math
INF = 100000
class minQueue:
    queue = []
    cost = {}
    def push_back(self, element):
        #向头部添加元素
        self.queue.insert( 0,element)
        self.shiftdown(0)
    def shiftdown(self,i):
        left = 2*i + 1
        right = 2*i + 2
        mini = i

        if left < len(self.queue) and self.cost[self.queue[left]] < self.cost[self.queue[mini]]:
            mini = left
        if right <  len(self.queue) and self.cost[self.queue[right]] < self.cost[self.queue[mini]]:
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
    location_file_path = "location.json"
    location_data = None #位置

    #最初构造矩阵的方法
    #根据json文件构造二维矩阵
    # def getMatrix(self,file_path):
    #
    #     with open(file_path) as f:
    #         self.location_data = json.load(f)
    #     location_len_list = range(len(self.location_data))
    #     self.martix = [_ for _ in location_len_list]
    #     for i in location_len_list:
    #         tmp = [INF for _ in location_len_list]
    #         for j in location_len_list:
    #             if j in self.location_data[str(i)][2]:
    #                 tmp[j] = -1
    #         self.martix[i] = tmp
    #     #矩阵初始化结束
    #     url_head = 'https://restapi.amap.com/v3/distance?key=76a251e04fd7611225f906c85b2a0163'
    #     tmp_url = "&origins="
    #     #然后通过web获得时间(以时间代替路程)
    #     for i in location_len_list:
    #         ii = str(i)
    #         tmp_url += str(self.location_data[ii][0])+','+str(self.location_data[ii][1]) + '|'
    #     tmp_url = tmp_url[:-1]
    #     # print("tmp url",tmp_url)
    #     #计算距离
    #     for ii in location_len_list:
    #         i = str(ii)
    #         destination = "&destination=" + str(self.location_data[i][0] ) +','+ str(self.location_data[i][1])
    #
    #         url = url_head + tmp_url + destination
    #         data = json.loads(requests.get(url).text)['results']
    #         print(data)
    #         for j in data:
    #             a = int(j['origin_id']) - 1
    #             if self.martix[a][ii] != INF:
    #                 self.martix[a][ii] = j['duration']
    #                 self.martix[ii][a] = j['duration']
    #     #写入文件
    #     with open(self.distance_file_path,"w") as f:
    #         json.dump( json.dumps(self.martix),f)

    def read_from_file(self):
        with open(self.distance_file_path) as f:
            self.martix = json.loads(json.load(f))
    def search(self, points = []):
        print(points)
        #Ex . points = [1,2,3,4,1]
        #方法描述       #根据points列表查询两点之间的距离
        #返回值 points数组点之间的路径 以及花费的时间
        n_list = range(len(points)-1)
        search_res_path = []
        search_res_time = {}
        for i in n_list:
            tmp = self.dijkstra(points[i],points[i+1])
            search_res_path.extend(tmp[0])

            search_res_time[tmp[0][-1]] = tmp[1]
        return search_res_path,search_res_time
    def dijkstra(self, start,end):
        q = minQueue()
        len_list = range(len(self.martix))
        for i in len_list:
            q.cost[i] = INF
        q.cost[start] = 0
        q.push_back(start)
        pre = {}
        while len(q.queue) !=0 :
            u = q.pop_head()
            for i in self.adj(u):
                if q.cost[i] > q.cost[u] + float(self.martix[u][i] ):
                    q.cost[i] = q.cost[u] + float(self.martix[u][i])
                    pre[i] = u
                    q.push_back(i)
        return self.get_path(pre, end)

    def get_path(self, parent={},end = 0):
        u = end
        weight = 0
        path_a = [end]
        while parent.get(u) != None:
            weight += float(self.martix[u][parent.get(u)])
            u = parent.get(u)
            path_a.append(u)
        path_a.reverse()
        return path_a, math.ceil(weight / 60)
    def adj(self, u): #返回u的邻接
        with open(self.location_file_path) as f:
            self.location_data = json.load(f)
        return self.location_data[str(u)][2]
    # def Astar(self, start, end):
    #     # 返回一条从start到end的路径,以及路径花费的时间
    #     len_list = range(len(self.martix))
    #     closed = set()
    #     open_q = minQueue()
    #     g_func = {}
    #     h_func = {}
    #     #对其初始化
    #     for i in len_list:
    #         g_func[i] = INF
    #         open_q.cost[i] = INF
    #     g_func[start] = 0
    #     open_q.cost[start] = 0
    #     open_q.push_back(start)
    #     parent = {}
    #
    #     while len(open_q.queue) != 0:
    #         u = open_q.pop_head()
    #         if u == end:
    #             #构造路径
    #             #路径权重:
    #             return  self.get_path(parent,end)
    #         closed.add(u)
    #         # print("closed = ",closed)
    #         for v in self.adj(u):
    #             gnmi = float(self.martix[u][v]) + g_func[u] # g(v)
    #             # h_func[v] = float(self.martix[u][v])  # h(v)#没有好的启发函数
    #             h_func[v] = 0
    #             fnmi = gnmi + h_func[v] #实际上退化为 dijkstra算法
    #             if v not in closed: #v不在closed里面
    #                 if v not in open_q.queue: #v不在open里面
    #                     open_q.cost[v] = fnmi
    #                     open_q.push_back(v)
    #                     parent[v] = u
    #                 else:
    #                     if gnmi < g_func[v]:
    #                         g_func[v] = gnmi
    #                         parent[v] = u
    #             else:
    #
    #                 if gnmi < g_func[v]:
    #                     closed.remove(v)
    #                     g_func[v] = gnmi
    #                     open_q.cost[v] =fnmi
    #                     open_q.push_back(v)
    #                     parent[v] = u
    #     return [],-1

if __name__ == '__main__':
    a,b = 30,50
    g = Graph()
    # g.read_from_file()
    print(g.adj(1))
    # start = time.time()
    # res = g.Astar(a,b)
    # end = time.time()
    # print("A* time",end-start)
    # print(res)
    # start = time.time()
    # res = g.dijkstra(a,b)
    # end = time.time()
    # print("dijkstra time", end - start)
    # print(res)
