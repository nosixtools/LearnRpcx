# LearnRpcx

通过 client 调用 hello_server,hello_server 再次调用 word_server 作为实例

通过 docker 启动  zipkin
`docker run -d -p 9411:9411 openzipkin/zipkin`

## start 

```
make hello
make world
make client
```

## image
![1.png](./image/1.png)

![2.png](./image/2.png)

