import * as grpc from '@grpc/grpc-js';
import * as protoLoader from '@grpc/proto-loader';

/**
 * gRPCクライアントを作成する関数
 * @param protoPath - protoファイルのパス
 * @param serviceName - サービス名
 * @param address - サーバーのアドレス
 * @returns gRPCクライアント
 */
const createClient = (protoPath: string, serviceName: string, address: string) => {
  // protoファイルをロードしてパッケージ定義を作成
  const packageDefinition = protoLoader.loadSync(protoPath, {
    // フィールド名の大文字小文字を保持
    keepCase: true,
    // long型を文字列として扱う
    longs: String,
    // enum型を文字列として扱う
    enums: String,
    // デフォルト値を含める
    defaults: true,
    // oneofフィールドを含める
    oneofs: true,
  });

  // パッケージ定義からgRPCオブジェクトを作成
  const grpcObject = grpc.loadPackageDefinition(packageDefinition);
  // 指定されたサービス名のクライアントを取得
  const ServiceClient = grpcObject[serviceName] as any;

  // サービスが見つからない場合はエラーをスロー
  if (!ServiceClient) {
    throw new Error(`Service ${serviceName} not found in proto file`);
  }

  // gRPCクライアントを作成して返す
  const client = new ServiceClient(address, grpc.credentials.createInsecure());
  return client;
};

// createClient関数をエクスポート
export default createClient;