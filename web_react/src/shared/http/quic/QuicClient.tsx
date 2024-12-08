import { useState, useEffect } from 'react';

interface QuicProtoClientProps {
  url: string;
  requestData: any;
  YourResponse: any;
}

const useQuicProtoClient = ({ url, requestData, YourResponse }: QuicProtoClientProps) => {
  // データ取得用のステート
  const [data, setData] = useState<any | null>(null);
  // エラー内容を保持するステート
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
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

    fetchData();
  }, [url, requestData, YourResponse]);

  // データとエラーを返す
  return { data, error };
};

export { useQuicProtoClient };