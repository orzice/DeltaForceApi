<?php
/**
 * 三角洲数据帝 - 开放平台 API SDK (PHP)
 * Base URL: https://orzice.com/workApi
 */

class DeltaForceClient
{
    /** @var string 基础URL */
    public static string $BASE_URL = 'https://orzice.com/workApi';

    /** @var string API密钥 */
    public static string $TOKEN = '';

    /** @var int 请求超时（秒） */
    private int $timeout;

    /**
     * @param int $timeout 请求超时时间（秒）
     */
    public function __construct(int $timeout = 30)
    {
        $this->timeout = $timeout;
    }

    /**
     * 发送GET请求，自动附加 TOKEN
     *
     * @param string $path   接口路径，如 /v1/sjz_api/item_info_all
     * @param array  $params 查询参数
     * @return array
     * @throws Exception
     */
    public function get(string $path, array $params = []): array
    {
        $params['token'] = self::$TOKEN;
        $url = self::$BASE_URL . $path . '?' . http_build_query($params);

        $ch = curl_init();
        curl_setopt_array($ch, [
            CURLOPT_URL            => $url,
            CURLOPT_RETURNTRANSFER => true,
            CURLOPT_TIMEOUT        => $this->timeout,
            CURLOPT_HTTPHEADER     => [
                'Accept: application/json',
                'User-Agent: DeltaForceApi/1.0',
            ],
        ]);

        $body = curl_exec($ch);
        $httpCode = curl_getinfo($ch, CURLINFO_HTTP_CODE);
        $error  = curl_error($ch);
        curl_close($ch);

        if ($error) {
            throw new Exception("请求失败: {$error}");
        }
        if ($httpCode == 404) {
            throw new Exception('无权限访问，请检查 token 是否正确');
        }
        if ($httpCode != 200) {
            throw new Exception("请求失败，状态码: {$httpCode}, 响应: {$body}");
        }

        return json_decode($body, true) ?: [];
    }

    /**
     * 发送POST请求
     *
     * @param string $path   接口路径
     * @param array  $params 查询参数
     * @param array  $data   请求体
     * @return array
     * @throws Exception
     */
    public function post(string $path, array $params = [], array $data = []): array
    {
        $url = self::$BASE_URL . $path;
        if ($params) {
            $url .= '?' . http_build_query($params);
        }

        $ch = curl_init();
        curl_setopt_array($ch, [
            CURLOPT_URL            => $url,
            CURLOPT_RETURNTRANSFER => true,
            CURLOPT_TIMEOUT        => $this->timeout,
            CURLOPT_POST           => true,
            CURLOPT_POSTFIELDS     => http_build_query($data),
            CURLOPT_HTTPHEADER     => [
                'Accept: application/json',
                'User-Agent: DeltaForceApi/1.0',
            ],
        ]);

        $body = curl_exec($ch);
        $httpCode = curl_getinfo($ch, CURLINFO_HTTP_CODE);
        $error  = curl_error($ch);
        curl_close($ch);

        if ($error) {
            throw new Exception("请求失败: {$error}");
        }
        if ($httpCode == 404) {
            throw new Exception('无权限访问，请检查 token 是否正确');
        }
        if ($httpCode != 200) {
            throw new Exception("请求失败，状态码: {$httpCode}, 响应: {$body}");
        }

        return json_decode($body, true) ?: [];
    }
}
