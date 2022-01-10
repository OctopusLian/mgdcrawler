# Go实现分布式版本爬虫  

## 技术栈  

- 使用ElasticSearch作为数据存储  
- 使用Go语言标准模板库实现http数据展示部分  

## Docker安装  

```
$ docker run -d -p 80:80 nginx

$ docker run -d -p 9200:9200 -p 9300:9300 elasticsearch
```

## 进度  

elasticsearch依赖问题暂时没有解决。