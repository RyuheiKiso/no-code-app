import { useState, useEffect } from 'react';

// RestAPIClientProps インターフェース
// url: 接続先のURL
// requestData: リクエストデータ
// method: HTTPメソッド ('GET', 'POST', 'PUT', 'DELETE')
interface RestAPIClientProps {
  url: string;
  requestData: any;
  method: 'GET' | 'POST' | 'PUT' | 'DELETE';
}

// useRestAPIClient フック
// REST APIを使用してデータを送受信するカスタムフック
const useRestAPIClient = ({ url, requestData, method }: RestAPIClientProps) => {
  // データ取得用のステート
  const [data, setData] = useState<any | null>(null);
  // エラー内容を保持するステート
  const [error, setError] = useState<string | null>(null);

  // コンポーネントのマウント時および依存関係の変更時に実行される
  useEffect(() => {
    // データを取得する非同期関数
    const fetchData = async () => {
      try {
        // fetch APIを使用してHTTPリクエストを送信
        const response = await fetch(url, {
          // HTTPメソッドを設定
          method,
          // リクエストヘッダーを設定
          headers: {
            'Content-Type': 'application/json',
          },
          // GETメソッド以外の場合、リクエストボディを設定
          body: method !== 'GET' ? JSON.stringify(requestData) : null,
        });

        // レスポンスが正常か確認
        if (!response.ok) {
          throw new Error(`Error: ${response.statusText}`);
        }

        // レスポンスデータをJSON形式で取得
        const responseData = await response.json();
        // データをステートにセット
        setData(responseData);
      } catch (err) {
        // エラーメッセージをステートにセット
        if (err instanceof Error) {
          setError(err.message);
        } else {
          setError('不明なエラーが発生しました');
        }
      }
    };

    // データ取得関数を呼び出す
    fetchData();
  }, [url, requestData, method]);

  // データとエラーを返す
  return { data, error };
};

export { useRestAPIClient };