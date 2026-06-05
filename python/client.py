"""
DeltaForce API 客户端 - 纯请求封装
"""

import requests
from typing import Optional


class DeltaForceClient:
    """三角洲数据帝 API 客户端

    使用方式:
        from deltaforce import DeltaForceClient

        # 类变量配置
        DeltaForceClient.TOKEN = "your_token"
        # 可选：修改 BASE_URL
        # DeltaForceClient.BASE_URL = "https://custom.url"

        client = DeltaForceClient()
        resp = client.get("/v1/sjz_api/item_info_all")
        resp = client.get("/v1/sjz_api/item_list", {"lx": "list", "p": 1, "limit": 10, "types": "acc"})
        resp = client.post("/ide/", params={"iChartId": "317814", "sIdeToken": "xxx"})
    """

    BASE_URL = "https://orzice.com/workApi"
    """
    你的API密钥 TOKEN
    """
    TOKEN = ""

    def __init__(self, timeout: int = 30):
        """
        Args:
            timeout: 请求超时时间（秒）
        """
        self._timeout = timeout
        self._session = requests.Session()
        self._session.headers["User-Agent"] = "DeltaForceApi/1.0"

    def get(self, path: str, params: Optional[dict] = None) -> dict:
        """发送GET请求，自动附加 TOKEN

        Args:
            path: 接口路径，如 /v1/sjz_api/item_info_all
            params: 查询参数
        """
        if params is None:
            params = {}
        params.setdefault("token", self.TOKEN)
        url = f"{self.BASE_URL}{path}"
        resp = self._session.get(url, params=params, timeout=self._timeout)
        if resp.status_code == 404:
            raise PermissionError("无权限访问，请检查 token 是否正确")
        resp.raise_for_status()
        return resp.json()

    def post(self, path: str, params: Optional[dict] = None, data: Optional[dict] = None) -> dict:
        """发送POST请求

        Args:
            path: 接口路径
            params: 查询参数
            data: 请求体
        """
        if params is None:
            params = {}
        url = f"{self.BASE_URL}{path}"
        resp = self._session.post(url, params=params, data=data, timeout=self._timeout)
        if resp.status_code == 404:
            raise PermissionError("无权限访问，请检查 token 是否正确")
        resp.raise_for_status()
        return resp.json()

    def close(self):
        """关闭HTTP会话"""
        self._session.close()

    def __enter__(self):
        return self

    def __exit__(self, *args):
        self.close()
