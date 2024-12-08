# QUIC クライアントの使用方法

このディレクトリには、QUIC クライアントを使用してデータを取得するためのコードが含まれています。

## ファイル構成

- `QuicClient.tsx`: QUIC クライアントを作成し、データを取得するためのフックを定義しています。

## 使用方法

### QUIC クライアントの作成

`QuicClient.tsx` ファイルには、QUIC クライアントを作成するための `useQuicProtoClient` フックが定義されています。このフックを使用して、指定された URL とリクエストデータに基づいて QUIC クライアントを作成し、データを取得できます。

```typescript
import { useQuicProtoClient } from './QuicClient';

const url = 'https://your-quic-server.com';
const requestData = { /* your request data */ };
const YourResponse = /* your response class */;

const { data, error } = useQuicProtoClient({ url, requestData, YourResponse });

if (error) {
  console.error('Error:', error);
} else {
  console.log('Response:', data);
}
```

## 注意事項

- `url` には、QUIC サーバーの URL を指定します。
- `requestData` には、QUIC サービスに送信するリクエストデータを指定します。
- `YourResponse` には、レスポンスデータをデシリアライズするためのクラスを指定します。
