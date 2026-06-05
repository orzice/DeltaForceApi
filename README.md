# DeltaForceApi

三角洲数据帝开放平台 API SDK，提供 Python / Go / PHP 三种语言的请求封装。

<p align="center">
  <b>开放平台</b> &nbsp;·&nbsp;
  <a href="https://orzice.com/work/">orzice.com/work</a>
  &nbsp;&nbsp;|&nbsp;&nbsp;
  <b>API 文档</b> &nbsp;·&nbsp;
  <a href="https://work-api.apifox.cn">work-api.apifox.cn</a>
  &nbsp;&nbsp;|&nbsp;&nbsp;
  <b>QQ 群</b> &nbsp;·&nbsp;
  <code>791785919</code>
</p>

---

## 安装

<!-- tabs:start -->

#### **Python**

```bash
pip install requests
```

将 `python/` 目录加入 `PYTHONPATH`，或直接复制到项目中。

#### **Go**

```bash
go get github.com/orzice/DeltaForceApi/go/deltaforce
```

#### **PHP**

将 `php/Client.php` 复制到项目中 `require` 即可，依赖 cURL 扩展。

<!-- tabs:end -->

## 快速开始

三种语言 API 风格一致：配置类变量 `TOKEN`，调用 `get()` / `post()`。

<table>
<tr>
<td width="33%">

**Python**
```python
from deltaforce import DeltaForceClient

DeltaForceClient.TOKEN = "your_token"

client = DeltaForceClient()
data = client.get("/v1/sjz_api/item_info_all")
data = client.get("/v1/sjz_api/item_list",
    {"types": "acc", "limit": 10})
```

</td>
<td width="33%">

**Go**
```go
import "github.com/orzice/DeltaForceApi/go/deltaforce"

deltaforce.Token = "your_token"
client := deltaforce.NewClient(30 * time.Second)
resp, _ := client.Get(
    "/v1/sjz_api/item_info_all", nil)
resp, _ = client.Get(
    "/v1/sjz_api/item_list",
    map[string]string{
        "types": "acc", "limit": "10"})
```

</td>
<td width="33%">

**PHP**
```php
require 'Client.php';

DeltaForceClient::$TOKEN = 'your_token';
$client = new DeltaForceClient();
$data = $client->get(
    '/v1/sjz_api/item_info_all');
$data = $client->get(
    '/v1/sjz_api/item_list',
    ['types' => 'acc', 'limit' => 10]);
```

</td>
</tr>
</table>

## API 接口

| 模块 | 说明 |
|:---|:---|
| 基础接口 | 全物品基础信息、最新价格、今日密码、改枪码 |
| 鼠鼠卡战备 | 实时卡战备 V3 / V4、假账排行榜、DIY 自定义凑战备 |
| 实时交易行 | 全物品列表 / Top 排行、价格曲线（分钟 / 小时 / 天） |
| 军需处兑换 + 制造 | 兑换成本、子弹收益、制造利润 |
| 数据分析 | 钥匙卡预测、子弹分析 |

> 完整接口路径与参数参考 [`llms.txt`](llms.txt) 或 [线上 API 文档](https://work-api.apifox.cn)

## 通用设计

| 特性 | 说明 |
|:---|:---|
| 配置 | `TOKEN` / `BASE_URL` 为静态变量，配置一次全局生效 |
| 鉴权 | `get()` 自动注入 `TOKEN` |
| 错误 | 404 抛出"无权限访问" |
| UA | `DeltaForceApi/1.0` |
| 依赖 | Python 仅 `requests`，Go / PHP 零外部依赖 |
