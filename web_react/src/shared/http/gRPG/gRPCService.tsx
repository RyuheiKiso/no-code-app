import createClient from './gRPCClient';

/**
 * gRPCサービスからデータを取得する関数
 * @param protoPath - protoファイルのパス
 * @param serviceName - サービス名
 * @param address - サーバーのアドレス
 * @param request - リクエストデータ
 * @returns レスポンスデータのPromise
 */
export const getYourData = async (protoPath: string, serviceName: string, address: string, request: any): Promise<any> => {
  // gRPCクライアントを作成
  const ServiceClient = createClient(protoPath, serviceName, address);

  // Promiseを返す
  return new Promise((resolve, reject) => {
    // gRPCサービスを呼び出してデータを取得
    ServiceClient.getYourData(request, (error: any, response: any) => {
      if (error) {
        // エラーが発生した場合はreject
        reject(error);
      } else {
        // 正常にデータを取得できた場合はresolve
        resolve(response);
      }
    });
  });
};