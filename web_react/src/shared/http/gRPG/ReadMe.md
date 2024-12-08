# gRPC クライアントとサービスの使用方法

このディレクトリには、gRPC クライアントとサービスを使用してデータを取得するためのコードが含まれている。

## ファイル構成

- `gRPCClient.tsx`: gRPC クライアントを作成する関数を定義している。
- `gRPCService.tsx`: gRPC サービスからデータを取得する関数を定義している。

## protoファイルの作成方法

gRPCサービスを定義するために、まずprotoファイルを作成します。以下は、`sample.proto`ファイルの例です。

```proto
syntax = "proto3";

package yourpackage;

// サービスの定義
service YourService {
  // RPCメソッドの定義
  rpc GetYourData (YourRequest) returns (YourResponse);
}

// リクエストメッセージの定義
message YourRequest {
  string your_field = 1;
}

// レスポンスメッセージの定義
message YourResponse {
  string your_response_field = 1;
}
```

## 使用方法

### gRPC クライアントの作成

`gRPCClient.tsx` ファイルには、gRPC クライアントを作成するための `createClient` 関数が定義されている。この関数を使用して、指定された proto ファイル、サービス名、およびサーバーアドレスに基づいて gRPC クライアントを作成できる。

```typescript
import createClient from './gRPCClient';

const protoPath = 'path/to/your.proto';
const serviceName = 'YourServiceName';
const address = 'localhost:50051';

const client = createClient(protoPath, serviceName, address);
```

### gRPC サービスからデータを取得

`gRPCService.tsx` ファイルには、gRPC サービスからデータを取得するための `getYourData` 関数が定義されている。この関数を使用して、指定されたリクエストデータを gRPC サービスに送信し、レスポンスデータを取得できる。

```typescript
import { getYourData } from './gRPCService';

const protoPath = 'path/to/your.proto';
const serviceName = 'YourServiceName';
const address = 'localhost:50051';
const requestData = { /* your request data */ };

getYourData(protoPath, serviceName, address, requestData)
  .then(response => {
    console.log('Response:', response);
  })
  .catch(error => {
    console.error('Error:', error);
  });
```

## 注意事項

- `protoPath` には、使用する proto ファイルのパスを指定する。
- `serviceName` には、使用する gRPC サービスの名前を指定する。
- `address` には、gRPC サーバーのアドレスを指定する。
- `requestData` には、gRPC サービスに送信するリクエストデータを指定する。
