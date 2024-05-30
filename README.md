# gsls

aliyun sls tools


## 使用方式

```

Usage of ./gsls:
  -ak string
        Access Key for authentication
  -end string
        end time, unix timestamp (default "time.now()")
  -endpoint string
        sls endpoint (default "cn-hangzhou.log.aliyuncs.com")
  -logstore string
        sls logstore name (default "test-logstore")
  -project string
        sls project (default "test-project")
  -query string
        sls query expr (default "*")
  -size string
        query size (default "10")
  -sk string
        Secret Key for authentication
  -start string
        start time, unix timestamp (default "1717051210")
  -timeout string
        query timeout (default "300")

```


## 示例

```

// cli command
./gsls -endpoint="cn-xxx.log.aliyuncs.com" -ak="xxx" -sk="xxx" -project="xxx" -logstore="xxx" 

// response
--start sls query, req: *, start time: 2024-05-30 14:40:10 +0800 CST, end time: 2024-05-30 14:51:22.083391 +0800 CST m=+0.001372728 --
--query success! log count: 10 
, use time: 110.22594ms --

 [1] doc, content: xxx 
 [2] doc, content: xxx 
 ...
 
```
