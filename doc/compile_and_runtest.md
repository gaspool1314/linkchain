享云链支持隐私保护交易，隐私保护用到了多种加密算法及协议，这里使用了开源项目monero实现的加密协议。

编译享云链项目首先需要使用monero编译出依赖的静态库。

golang使用1.12版本

# 准备

`$ git clone https://github.com/gaspool1314/linkchain.git ; cd linkchain ; git checkout origin/br_testnet -b br_testnet `

# 编译享云链
## 获取并运行镜像

`$ mkdir -p /home/blockdata`

`$ docker run -dit -v /home/blockdata:/blockdata  --name linkchain --net=host garrixwong/go1.12-boost-centos7:0.1.0`

```bash
Unable to find image 'garrixwong/go1.12-boost-centos7:0.1.0' locally
0.1.0: Pulling from garrixwong/go1.12-boost-centos7
d8d02d457314: Already exists 
ec488c4822b0: Pull complete 
Digest: sha256:eb468b5c0615ead9329585c28f2f4beca0fc3a61a7650285e46c2c3ec3674f07
Status: Downloaded newer image for garrixwong/go1.12-boost-centos7:0.1.0
1461aaf1b2d6e9954909105d8faca90c2193b028cec30da0340cc21ad4fe73f2
```

## 编译可执行程序
`$ docker exec -it linkchain /bin/bash`

`$ rm -rf linkchain ; git clone https://github.com/gaspool1314/linkchain.git ; cd linkchain`

`$ git checkout origin/br_testnet -b br_testnet`

`$ git pull`

`$ export PATH=$PATH:/usr/local/go/bin && scl enable devtoolset-8 bash `

`$ ./build.sh `

编译成功后在/linkchain/bin目录能看到编译后的文件：

`$ ll /linkchain/bin`

```bash
total 89668
-rwxr-xr-x 1 root root 49804544 Aug 27 02:49 lkchain
-rwxr-xr-x 1 root root 42012056 Aug 27 02:49 wallet
```
`# /linkchain/bin/lkchain version`

```bash
linkchain version: 0.1.0, gitCommit:7f5d2a3e
```
# 运行享云链
## 测试模式运行单节点本地测试网络

`$ docker exec -it linkchain /bin/bash`

```bash
[root@5f400a3ad5cf /]# 
```
初始化：

`$ cd /linkchain/pack/lkchain/sbin/ ; ./start-pre-testnet.sh init`

```bash
committee contract code nil!!!
validators white list contract code nil!!!
genesisBlock stateHash 0x0d8827403cb36d8d176cbf6257915f1b5274ba11ff2891b06a0263946ebf0b57
genesisBlock trieRoot 0x0000000000000000000000000000000000000000000000000000000000000000
genesisBlock ChainID:chainID block.Hash:0x26cb0291c88674df8614a93eb0e1b5e23b82e3117f18dade10acb0cf7c597b2d
```

启动节点：

`$ sh ./start-pre-testnet.sh start`

```bash
start lkchain ...
pid: 390
```

测试RPC:

`# curl -H 'Content-Type: application/json' -d '{"jsonrpc":"2.0","id":"0","method":"eth_blockNumber","params":[]}' http://127.0.0.1:16000`

```bash
{"jsonrpc":"2.0","id":"0","result":"0x0"}
```
查看Log:
`# tail /linkchain/pack/lkchain/data/logs/lkchain.log`

```bash
DEBUG 2019-08-27 03:04:44.797 status report                            module=mempool specGoodTxs=0 goodTxs=0 futureTxs=0
DEBUG 2019-08-27 03:04:44.819 Broadcasting proposal heartbeat message  module=consensus height=3 round=0 sequence=1
DEBUG 2019-08-27 03:04:46.820 Broadcasting proposal heartbeat message  module=consensus height=3 round=0 sequence=2
DEBUG 2019-08-27 03:04:48.821 Broadcasting proposal heartbeat message  module=consensus height=3 round=0 sequence=3
DEBUG 2019-08-27 03:04:49.797 status report                            module=mempool specGoodTxs=0 goodTxs=0 futureTxs=0
DEBUG 2019-08-27 03:04:49.865 dialOutLoop                              module=conManager maxDialOutNums=3 needDynDials=3
DEBUG 2019-08-27 03:04:49.865 ReadRandomNodes                          module=httpTable tab.seeds=[]
DEBUG 2019-08-27 03:04:49.865 after dialRandNodesFromCache             module=conManager needDynDials=3
DEBUG 2019-08-27 03:04:49.865 dialNodesFromNetLoop                     module=conManager needDynDials=3
DEBUG 2019-08-27 03:04:50.822 Broadcasting proposal heartbeat message  module=consensus height=3 round=0 sequence=4
```

关闭节点:

`# sh ./start-pre-testnet.sh stop`

```bash
kill 390
```
