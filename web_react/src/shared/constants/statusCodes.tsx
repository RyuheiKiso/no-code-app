// HTTPステータスコードを定義するオブジェクト
const statusCodes = {
    // 成功
    OK: 200,
    // 作成成功
    CREATED: 201,
    // 受理
    ACCEPTED: 202,
    // コンテンツなし
    NO_CONTENT: 204,
    // 不正なリクエスト
    BAD_REQUEST: 400,
    // 認証エラー
    UNAUTHORIZED: 401,
    // 禁止
    FORBIDDEN: 403,
    // 見つからない
    NOT_FOUND: 404,
    // サーバー内部エラー
    INTERNAL_SERVER_ERROR: 500,
    // 未実装
    NOT_IMPLEMENTED: 501,
    // 不正なゲートウェイ
    BAD_GATEWAY: 502,
    // サービス利用不可
    SERVICE_UNAVAILABLE: 503,
    // ゲートウェイタイムアウト
    GATEWAY_TIMEOUT: 504,
};

export default statusCodes;

/*
使用例:

import statusCodes from './statusCodes';

if (response.status === statusCodes.OK) {
    console.log('リクエストが成功しました');
} else if (response.status === statusCodes.NOT_FOUND) {
    console.log('リソースが見つかりません');
} else if (response.status === statusCodes.INTERNAL_SERVER_ERROR) {
    console.log('サーバー内部エラーが発生しました');
}
*/