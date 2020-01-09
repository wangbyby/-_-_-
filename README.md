#大学写的代码
- go语言实现的一些算法
- 大学中写的程序作业(大部分收录在此)

#rust总结
- while let Some 和 if let Some 很好使
- 首选递归实现某些算法
- trait的使用
  - trait的面向对象特征
    ```rust
        trait Car{
        fn approve(self: Box<Self>) -> Box<dyn Car>;
    }
    #[derive(Debug)]
    struct Tsl{
    }

    impl Car for Tsl {
        //注意 dyn关键字
        fn approve(self: Box<Self>) -> Box<dyn Car>{
            self
        }
    }
    fn main() {
        let a = Box::new(Tsl{});
        let res = a.approve();
    }
    ```
  - trait的默认行为
    ```rust
    trait Car{
        //默认行为定义
        fn run(&self) -> String {
            String::from("i am running")
        }
    }

    #[derive(Debug)]
    struct Tsl{}

    impl Car for Tsl {
        fn run(&self) -> String {
            String::from("特斯拉在跑")
        }
    }

    struct Wl{}
    //使用默认行为
    impl Car for Wl{}

    fn main() {
        let tsl = Box::new(Tsl{});
        let s = tsl.run();
        println!("{:?}", s);

        let wl = Box::new(Wl{});
        let wres = wl.run();
        println!("{:?}", wres);
    }
    ```
- match用法
  - @与 a..=b 的用法
    ```rust
    fn main(){
        enum Message {
            Hello { id: i32 },
        }
        
        let msg = Message::Hello { id: 5 };
        
        match msg {
            //在解析的同时 使用@, 来检测值的范围
            Message::Hello { id: id_variable @ 3..=7 } => {
                println!("Found an id in range: {}", id_variable)
            },
            Message::Hello { id: 10..=12 } => {
                println!("Found an id in another range")
            },
            Message::Hello { id } => {
                println!("Found some other id: {}", id)
            },
        }
    }
    ```
- map的使用
    ```rust
        // let v = (1..20).map(|x| x+10).collect::<Vec<_>>();
        // for i in &v{
        //     println!("{}",i);
        // }
        let mut rag = rand::thread_rng();
        let mut a  = vec![0;20]; //暂时不指定类型
        for _j in 0..20 {
            let tmp:i8 = rag.gen(); //产生随机数
            a[_j] = tmp;
            
        }
        for i in &a{
            println!("{}",i);
        }
        let bb = a.iter().map(|&x| if x<0 {-x}else{x}).collect::<Vec<_>>(); //不是过滤掉...
        for i in &bb{
            println!("{}",i);
        }
    ```
- Cow 写时复制
    ```rust
    //copy on write
    //多学
    use std::borrow::Cow;
    fn cow_test<'a> (input :&'a str) -> Cow<'a,str>{
        if input.contains(' ') {
            let mut buf  = String::with_capacity(input.len());
            for i in input.chars() { //这里String是没有 buf[0](按下标取值)操作的
                if i != ' ' {
                    buf.push(i);
                }
            }
            return Cow::Owned(buf);
        }
        Cow::Borrowed(input)
    }   
    ```
- c-like遍历二叉树
    ```rust
    //现在写的还是挺烂的
    pub fn insert_nonrecursive(&mut self,k:K,v:V) {   
        let mut node = self;
        loop {
            match node.key<=k {
                true => {
                    if let Some(ref mut right) = node.right {
                        node = right;
                    }else{
                        node.right = Some(Box::new(BinNode::new(k, v)));
                        break;
                    }
                },
                false =>{
                    if let Some(ref mut l) = node.left {
                        node = l;
                    }else{
                        node.left = Some(Box::new(BinNode::new(k, v)));
                        break;
                    }
                },
            }
            }
        }
    ```
- 引用, 生命周期
  - 生命周期检查器
    ```rust
        
    fn main() {

        let s1 = String::from("hello");
        let s2 = "world";
        let c = longest(&s1, s2);
    }
    fn longest<'c>(a: &'c str, b: &'c str) -> &'c str{
        if a.len() > b.len() {
            return  a;
        }
        b
    }
    ```

  - 'static 静态变量
- 多线程

  - move关键字
    ```rust
    use std::thread;
    fn main(){
        let v = vec![1,2,3,4];
        let handle = thread::spawn(move ||{
            for i in &v{
                println!("{}", i);
            }
        });
        for i in 12..=17{
            println!("{}", i);
        }
        handle.join().unwrap();
    }
    ```
   - channel
        ```rust
        use std::thread;
        use std::sync::mpsc;
        fn main(){
            let (sender, receiver) = mpsc::channel();
            let sender2 = sender.clone(); // 只有发送端有clone, 接收端没有:)
            let v = vec![1,2,3,4];
            let handle = thread::spawn(move ||{
                for i in v{
                    println!("{}", i);

                    //想象一个船从上游到下游
                    //消耗所有权
                    sender.send(i).unwrap(); 
                }
                
            });
            thread::spawn(move || {
                for i in 2..=5{
                    sender2.send(i).unwrap();
                }
                
            });
            for i in receiver {
                println!("{}", i);
            }
            
            handle.join().unwrap();
        }
        
        ```
  - 共享资源
    ```rust
    use std::thread;
    use std::sync::{Mutex, Arc,};
    fn main(){
        
        let counter = Arc::new(Mutex::new(0));
        
        let mut handlers = vec![];
        
        for _ in 0..10{
            let counter = Arc::clone(&counter);
            let handler = thread::spawn(move || {
                let mut num = counter.lock().unwrap();
                *num +=1;
                
            });
            handlers.push(handler);
        }
        for i in handlers{
            i.join().unwrap();
        }
        println!("{}",*counter.lock().unwrap());
    }
    ```
- 测试
      - cargo test -- --nocapture
      - 单元测试
        ```
        #[cfg(test)] //只在cargo test中生效, 不编译在cargo build中
        ```
        
