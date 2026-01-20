# 在浏览器上展示

## 在浏览器中展示count

## 调用count方法，增加1

## .env

重新部署ts代码：
npx hardhat run ./scripts/deploy-counnter.ts --network localhost

拿到最新的合约地址
修改全局变量为counter的to地址

编写页面内容：

启动web应用：
npm run dev

启动本地网络：
npx hardhat run ./scripts/deploy-counnter.ts --network localhost
也可以在配置文件中配置后，用下面的命令启动
npm run dev
