# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## 项目概述

三角洲数据帝开放平台 API SDK，提供 Python / Go / PHP 三种语言的 HTTP 请求封装。

Base URL: `https://orzice.com/workApi`  
GitHub: `github.com/orzice/DeltaForceApi`

## 代码结构

```
python/           # Python SDK（依赖 requests）
  __init__.py     # 导出 DeltaForceClient
  client.py       # get() / post()
go/deltaforce/    # Go SDK（零外部依赖，仅 net/http）
  client.go       # Client.Get() / Client.Post()
  go.mod
php/              # PHP SDK（零外部依赖，仅 cURL）
  Client.php      # DeltaForceClient::get() / post()
llms.txt          # API 接口文档链接（Apifox），接口路径和参数参考此文件
```

## 三语言共通的 API 设计

- `TOKEN` 和 `BASE_URL` 为静态/包级变量，实例化前配置一次
- `get()` 自动将 `TOKEN` 注入请求参数，`post()` 不自动注入
- 所有请求 UA 为 `DeltaForceApi/1.0`
- 404 响应统一抛出"无权限访问"错误
- 新加 API 无需修改 SDK，直接传 path 和 params 即可

## 重要约定

- 三语言 SDK 保持 API 风格完全一致，修改一个语言时需同步其他两个
- 接口路径和参数以 `llms.txt` 中的 Apifox 文档为准
- Python: UTF-8 无 BOM；README.md 编码为 UTF-8
