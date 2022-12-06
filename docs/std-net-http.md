### net/http包

#### 参考资料
* [net/http库知道吗？能说说优缺点吗？](https://mp.weixin.qq.com/s/IelVDnMzGtT5y7hGSb_OxA)
* [Go HttpClient 超时机制](https://mp.weixin.qq.com/s/HPzoclfCB3UxLScXm4J83w)
* [golang http client 连接池](https://two.github.io/2018/06/01/golang-http-client-connection-pool/)

#### 服务端
```go
    addr := ":803"
    server := http.Server{
        Addr: addr,
    }
    http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
        _, _ = writer.Write([]byte("hello"))
    })
    // net.Listen最终会执行系统调用socket、bind、listen
    listener, _ := net.Listen("tcp", addr)
    // 每accept一个连接就会起一个协程
    _ = server.Serve(listener)
```

#### 客户端
```go
    client := http.Client{
        Timeout: 10 * time.Second,
    }
    // 最终执行net.http.Transport.roundTrip(req)
    resp, _ := client.Get("http://httpbin.org/get?a=1")
    body, _ := io.ReadAll(resp.Body)
    println(string(body))
```

net.http.Transport.roundTrip(req)的代码：
```go
func (t *Transport) roundTrip(req *Request) (*Response, error) {
    for {
        select {
        case <-req.Context().Done():
            req.closeBody()
            return nil, ctx.Err()
        default:
        }
        treq := &transportRequest{Request: req, cancelKey: cancelKey}
        pconn, err := t.getConn(treq, cm)
        resp, err = pconn.roundTrip(treq)
    }
}
```

t.getConn(treq, cm)的代码：
```go
func (t *Transport) getConn(treq *transportRequest, cm connectMethod) (pc *persistConn, err error) {
    w := &wantConn{
        cm:         cm,
        key:        cm.key(),
    }
    // 有空闲连接，直接用
    if delivered := t.queueForIdleConn(w); delivered {
        pc := w.pc
        return pc, nil
    }
    // 没有则排号
    t.queueForDial(w)
    select {
    case <-w.ready:                         // 连接建立好了
        return w.pc, w.err
    case <-req.Cancel:                      // Deprecated，其它业务可以通过此channel终止请求
        return nil, errRequestCanceledConn
    case <-req.Context().Done():            // 通过context来终止请求 
        return nil, req.Context().Err()
    case err := <-cancelc:                  // Deprecated，终止请求
        if err == errRequestCanceled {
            err = errRequestCanceledConn
        }
        return nil, err
    }
}
```

t.queueForDial(w)会执行`go t.dialConnFor(w)`，里面执行`t.dialConn(w.ctx, w.cm)`，其中关键的三步是：
```go
    conn, err := t.dial(ctx, "tcp", cm.addr())  // 关键一：会执行t.DialContext(ctx, network, addr)
                                                // 接着执行c, err = sd.dialSerial(ctx, primaries)
                                                // 接着执行系统调用socket、connnect
                                                
    go pconn.readLoop()                         // 关键二，新协程代码如下：
    alive := true
	for alive {
		rc := <-pc.reqch
		resp, err = pc.readResponse(rc, trace)
		if err != nil {
			select {
			case rc.ch <- responseAndError{err: err}:
			case <-rc.callerGone:
				return
			}
			return
		}
		select {
		case rc.ch <- responseAndError{res: resp}:
		case <-rc.callerGone:
			return
		}
		select {
		case bodyEOF := <-waitForBodyRead:          // caller协程已读完  
			alive = alive &&
				bodyEOF &&
				!pc.sawEOF &&
				pc.wroteRequest() &&
				replaced && tryPutIdleConn(trace)
			if bodyEOF {
				eofc <- struct{}{}
			}
		case <-rc.req.Cancel:                       // Deprecated，其它业务可以通过此channel终止请求
			alive = false
			pc.t.CancelRequest(rc.req)
		case <-rc.req.Context().Done():             // 通过context来终止请求 
			alive = false
			pc.t.cancelRequest(rc.cancelKey, err)
		case <-pc.closech:                          // 连接关闭
			alive = false
		}
	}
    
    go pconn.writeLoop()                            // 关键三，新协程代码如下
    for {
		select {
		case wr := <-pc.writech:
			startBytesWritten := pc.nwrite
			err := wr.req.Request.write(pc.bw, pc.isProxy, wr.req.extra, pc.waitForContinue(wr.continueCh))
			if bre, ok := err.(requestBodyReadError); ok {
				wr.req.setError(bre.error)
			}
			if err == nil {
				err = pc.bw.Flush()
			}
			pc.writeErrCh <- err
			wr.ch <- err
			if err != nil {
				pc.close(err)
				return
			}
		case <-pc.closech:
			return
		}
	}
```

pconn.roundTrip(treq)的代码：
```go
func (pc *persistConn) roundTrip(req *transportRequest) (resp *Response, err error) {
    pc.writech <- writeRequest{req, writeErrCh, continueCh} // writeLoop协程来处理
    pc.reqch <- requestAndChan{                             // readLoop协程来处理
        req:        req.Request,
        cancelKey:  req.cancelKey,
        ch:         resc,
        addedGzip:  requestedGzip,
        continueCh: continueCh,
        callerGone: gone,
    }
    for {
        select {
        case err := <-writeErrCh:        // 写出错 
            return
        case <-pc.closech:               // 连接关闭
            return
        case <-respHeaderTimer:          // 写response超时
            return nil, errTimeout
        case re := <-resc:               // 读到response
            return re.res, nil
        case <-req.Request.Cancel:       // Deprecated，其它业务可以通过此channel终止请求
            return
        case <-req.Context().Done():     // 通过context来终止请求     
            return
        }
    }
}
```