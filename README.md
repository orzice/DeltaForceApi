# DeltaForceApi

三角洲数据帝开放平台 API SDK，提供 Python / Go / PHP 三种语言的请求封装。

Base URL: `https://orzice.com/workApi`

## 安装

**Python**

```bash
pip install requests
```

将 `python/` 目录加入 `PYTHONPATH`，或直接复制到项目中。

**Go**

```bash
go get github.com/orzice/deltaforce-api/go/deltaforce
```

**PHP**

将 `php/Client.php` 复制到项目中 `require` 即可，依赖 cURL 扩展。

## 快速开始

所有语言 API 风格一致：先配置 `TOKEN`，再调用 `get()` / `post()`。

**Python**

```python
from deltaforce import DeltaForceClient

DeltaForceClient.TOKEN = "your_api_token"

client = DeltaForceClient()
data = client.get("/v1/sjz_api/item_info_all")
data = client.get("/v1/sjz_api/item_list", {"lx": "list", "p": 1, "limit": 10, "types": "acc"})
data = client.post("/ide/", params={"iChartId": "317814", "sIdeToken": "xxx"})
```

**Go**

```go
import "github.com/orzice/deltaforce-api/go/deltaforce"

deltaforce.Token = "your_api_token"
client := deltaforce.NewClient(30 * time.Second)
resp, _ := client.Get("/v1/sjz_api/item_info_all", nil)
resp, _ = client.Get("/v1/sjz_api/item_list", map[string]string{"types": "acc", "limit": "10"})
```

**PHP**

```php
require 'Client.php';

DeltaForceClient::$TOKEN = 'your_api_token';
$client = new DeltaForceClient();
$data = $client->get('/v1/sjz_api/item_info_all');
$data = $client->get('/v1/sjz_api/item_list', ['types' => 'acc', 'limit' => 10]);
```

## API 接口

所有接口路径和参数参考 `llms.txt`，主要包括：

| 模块 | 说明 |
|------|------|
| 基础接口 | 全物品基础信息、最新价格、今日密码、改枪码 |
| 鼠鼠卡战备 | 实时卡战备V3/V4、假账排行榜、DIY自定义凑战备 |
| 实时交易行 | 全物品列表/Top排行、价格曲线（分钟/小时/天） |
| 军需处兑换+制造 | 兑换成本、子弹收益、制造利润 |
| 数据分析 | 钥匙卡预测、子弹分析 |
| 玩家接口 | 授权登录、角色信息、战绩查询、对局详情 |

## 通用设计

- `TOKEN` 和 `BASE_URL` 为静态/包级变量，配置一次全局生效
- `get()` 自动将 `TOKEN` 注入请求参数
- 404 响应统一抛出"无权限访问"错误
- 请求 UA 为 `DeltaForceApi/1.0`
