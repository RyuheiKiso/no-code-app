import { useState, useEffect } from 'react';

// QuicProtoClientProps インターフェース
// url: 接続先のURL
// requestData: リクエストデータ
// YourResponse: レスポンスデータの型
interface QuicProtoClientProps {
  url: string;
  requestData: any;
  YourResponse: any;
}

// useQuicProtoClient フック
// QUICプロトコルを使用してデータを送受信するカスタムフック
const useQuicProtoClient = ({ url, requestData, YourResponse }: QuicProtoClientProps) => {
  // データ取得用のステート
  const [data, setData] = useState<any | null>(null);
  // エラー内容を保持するステート
  const [error, setError] = useState<string | null>(null);

  // コンポーネントのマウント時および依存関係の変更時に実行される
  useEffect(() => {
    // データを取得する非同期関数
    const fetchData = async () => {
      try {
        // リクエストデータをバイナリ形式にシリアライズ
        const requestBinary = requestData.serializeBinary();

        // QUIC通信を使用してデータを送信
        const response = await fetch(url, {
          method: 'POST',
          body: requestBinary,
          headers: {
            'Content-Type': 'application/x-protobuf',
          },
        });

        // レスポンスが正常か確認
        if (!response.ok) {
          throw new Error(`Error: ${response.statusText}`);
        }

        // レスポンスデータをバイナリで取得
        const responseBinary = await response.arrayBuffer();

        // バイナリデータをデシリアライズ
        const responseData = YourResponse.deserializeBinary(new Uint8Array(responseBinary));

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
  }, [url, requestData, YourResponse]);

  // データとエラーを返す
  return { data, error };
};

export { useQuicProtoClient };