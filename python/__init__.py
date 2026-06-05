"""
三角洲数据帝 - 开放平台 API SDK
=================================
Base URL: https://orzice.com/workApi

用法:
    from deltaforce import DeltaForceClient

    # 像 BASE_URL 一样设置 TOKEN
    DeltaForceClient.TOKEN = "your_api_token"

    client = DeltaForceClient()
    data = client.get("/v1/sjz_api/item_info_all")
    data = client.get("/v1/sjz_api/item_list", {"lx": "list", "p": 1, "limit": 10, "types": "acc"})
    data = client.post("/ide/", params={"iChartId": "317814", "sIdeToken": "xxx"})

接口列表参考 llms.txt
"""

from .client import DeltaForceClient

__version__ = "1.0.0"
__all__ = ["DeltaForceClient"]
