# NodeJs-GoLang性能对比测试

## 单核测试
#### GoLang 测试结果
``` bash
wrk -t50 -c1000 -d30s 'http://127.0.0.1:8088/'
Running 30s test @ http://127.0.0.1:8088/
  50 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    22.41ms   36.24ms   1.98s    99.72%
    Req/Sec   394.28    364.47     3.36k    85.34%
  505993 requests in 30.10s, 61.77MB read
  Socket errors: connect 0, read 1922, write 2, timeout 352
Requests/sec:  16809.00
Transfer/sec:      2.05MB
```

#### NodeJs 测试结果(v10.16.3)
``` bash
wrk -t50 -c1000 -d30s 'http://127.0.0.1:8087/'
Running 30s test @ http://127.0.0.1:8087/
  50 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    37.47ms   13.78ms 142.50ms   72.00%
    Req/Sec   432.03    216.81     2.24k    82.00%
  623107 requests in 30.10s, 65.96MB read
  Socket errors: connect 0, read 1232, write 0, timeout 0
Requests/sec:  20702.55
Transfer/sec:      2.19MB
```

### NodeJs 测试结果(v8.11.1)
``` bash
Running 30s test @ http://127.0.0.1:8087/
  50 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    53.20ms   14.27ms 164.18ms   72.24%
    Req/Sec   330.75    125.55     1.20k    78.58%
  488841 requests in 30.10s, 51.75MB read
  Socket errors: connect 0, read 1084, write 1, timeout 0
Requests/sec:  16241.47
Transfer/sec:      1.72MB
```

## 多核测试
#### GoLang 测试结果
``` bash
Running 30s test @ http://127.0.0.1:8088/
  50 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    28.25ms   92.08ms   1.99s    97.81%
    Req/Sec     1.05k   496.75    10.30k    79.96%
  1533082 requests in 30.10s, 187.14MB read
  Socket errors: connect 0, read 1227, write 1, timeout 180
Requests/sec:  50927.99
Transfer/sec:      6.22MB
```

#### NodeJs 测试结果(v10.16.3)
``` bash
wrk -t50 -c1000 -d30s 'http://127.0.0.1:8087/'
Running 30s test @ http://127.0.0.1:8087/
  50 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    37.36ms  104.41ms   1.99s    98.20%
    Req/Sec   818.90    219.91     3.78k    71.60%
  1191082 requests in 30.10s, 147.67MB read
  Socket errors: connect 0, read 339, write 5, timeout 147
Requests/sec:  39570.82
Transfer/sec:      4.91MB
```

#### NodeJs 测试结果(v8.11.1)
``` bash
wrk -t50 -c1000 -d30s 'http://127.0.0.1:8087/'
Running 30s test @ http://127.0.0.1:8087/
  50 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    41.82ms  118.47ms   2.00s    98.01%
    Req/Sec   748.88    158.83     3.87k    76.79%
  1078363 requests in 30.10s, 133.69MB read
  Socket errors: connect 0, read 359, write 12, timeout 65
Requests/sec:  35826.11
Transfer/sec:      4.44MB
```


## 密集场景多核测试
#### GoLang 测试结果
``` bash
wrk -t50 -c1000 -d30s 'http://127.0.0.1:8088/'
Running 30s test @ http://127.0.0.1:8088/
  50 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    53.22ms   68.90ms   1.99s    94.82%
    Req/Sec   217.72    158.78     1.88k    73.78%
  287529 requests in 30.10s, 35.10MB read
  Socket errors: connect 0, read 1474, write 0, timeout 462
Requests/sec:   9553.53
Transfer/sec:      1.17MB
```

#### NodeJs 测试结果(v10.16.3)
``` bash
wrk -t50 -c1000 -d30s 'http://127.0.0.1:8087/'
Running 30s test @ http://127.0.0.1:8087/
  50 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   214.79ms   88.70ms   1.99s    86.17%
    Req/Sec    93.95     61.56     1.33k    71.73%
  115586 requests in 30.09s, 14.33MB read
  Socket errors: connect 0, read 25, write 0, timeout 676
Requests/sec:   3840.90
Transfer/sec:    487.61KB
```

#### NodeJs 测试结果(v8.11.1)
``` bash
wrk -t50 -c1000 -d30s 'http://127.0.0.1:8087/'
Running 30s test @ http://127.0.0.1:8087/
  50 threads and 1000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   272.90ms  110.20ms   1.95s    78.95%
    Req/Sec    73.09     59.41     1.05k    74.93%
  83488 requests in 30.10s, 10.35MB read
  Socket errors: connect 0, read 54, write 2, timeout 712
Requests/sec:   2773.89
Transfer/sec:    352.15KB
```