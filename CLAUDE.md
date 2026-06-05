# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 项目概述

三角洲数据帝开放平台 Python SDK，对 `https://orzice.com/workApi` 的 HTTP 请求封装。仅依赖 `requests`。

## 代码结构

```
python/
  __init__.py    # 导出 DeltaForceClient
  client.py      # DeltaForceClient 类：get() / post()
llms.txt         # 所有 API 接口文档链接（Apifox），接口路径和参数参考此文件
```

`python/` 本身即 Python 包目录。

## 核心设计

- `DeltaForceClient.TOKEN` 和 `DeltaForceClient.BASE_URL` 是类变量，在实例化前全局配置一次即可
- `get()` 自动将 `TOKEN` 注入 `params["token"]`；`post()` 不自动注入
- 404 响应统一抛出 `PermissionError("无权限访问，请检查 token 是否正确")`
- 支持上下文管理器：`with DeltaForceClient() as client: ...`
- 新加 API 无需修改 SDK，直接调用 `client.get("/path", params)` 即可，接口路径和参数参考 `llms.txt`

## 依赖

```
requests
```
