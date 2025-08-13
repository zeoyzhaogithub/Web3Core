# 02-恒和自动做市商

![](./jpg/WeChatb098b4d817b5d92eed324bea5abd9170.jpg)

部署与测试：
1、准备两个代币
2、部署ERC20在deploy中构造代币
   name:MyToken0
   symbol:MT0
   decimals:18
   构造之后得到地址d0，
      给当前的管理员地址mint 1000
   name:MyToken1
   symbol:MT1
   decimals:18
   构造之后得到地址d1
      给当前的管理员地址mint 1000
3、部署CSAMM，将d0,d1传入构造函数
4、授权d0给CSAMM，授权d1给CSAMM
   分别在d0和d1中授权CSAMM,执行approve方法，参数为CSAMM地址，授权数量为1000

5、在CSAMM中执行addLiquidity方法，参数为d0地址，d1地址，1000，1000，当前账户地址

6、切换账户进行兑换
   在d0执行mint，地址为当前账户 1000
   在d0执行mint，地址为CSAMM地址 1000

   在CSAMM中执行swap方法，参数为:d0地址，1000，
