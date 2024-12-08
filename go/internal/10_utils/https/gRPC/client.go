package grpc

import (
	"context"
	"log"
	"time"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// NewClientConn は新しいgRPCクライアント接続を作成します
func NewClientConn(address string, timeout time.Duration, maxRetries int, retryBackoff time.Duration, authToken string) (*grpc.ClientConn, error) {
	// タイムアウト付きのコンテキストを作成
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	// 関数終了時にキャンセル
	defer cancel()

	// リトライオプションの設定
	retryOpts := []grpc_retry.CallOption{
		// 最大リトライ回数の設定
		grpc_retry.WithMax(uint(maxRetries)),
		// リトライのバックオフ時間の設定
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(retryBackoff)),
	}

	// KeepAliveパラメータの設定
	keepAliveParams := keepalive.ClientParameters{
		// Pingを送信する間隔
		Time: 10 * time.Second,
		// Pingの応答を待つ時間
		Timeout: 20 * time.Second,
		// ストリームがなくてもPingを許可
		PermitWithoutStream: true,
	}

	// メタデータの設定
	md := metadata.Pairs(
		// 現在のタイムスタンプをメタデータに追加
		"timestamp", time.Now().Format(time.StampNano),
		// 認証トークンをメタデータに追加
		"authorization", "Bearer "+authToken,
	)
	// メタデータをコンテキストに追加
	ctx = metadata.NewOutgoingContext(ctx, md)

	// カスタムロギングインターセプター
	loggingInterceptor := func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		// リクエスト開始時間を記録
		start := time.Now()
		// 実際のRPC呼び出し
		err := invoker(ctx, method, req, reply, cc, opts...)
		// リクエスト終了時間を記録
		end := time.Now()
		// エラーからステータスを取得
		st, _ := status.FromError(err)
		// ログ出力
		log.Printf("method: %s, req: %+v, reply: %+v, duration: %s, error: %v, status: %v", method, req, reply, end.Sub(start), err, st)
		return err
	}

	// トレーシングインターセプター
	tracingInterceptor := func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		// トレーサーを取得
		tracer := otel.Tracer("grpc-client")
		// スパンを開始
		ctx, span := tracer.Start(ctx, method)
		// スパンを終了
		defer span.End()

		// 実際のRPC呼び出し
		err := invoker(ctx, method, req, reply, cc, opts...)
		if err != nil {
			// エラーをスパンに記録
			span.RecordError(err)
		}
		return err
	}

	// エラーハンドリングインターセプター
	errorHandlingInterceptor := func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		// 実際のRPC呼び出し
		err := invoker(ctx, method, req, reply, cc, opts...)
		if err != nil {
			// エラーからステータスを取得
			st, _ := status.FromError(err)
			// エラーログを出力
			log.Printf("RPC failed with status: %v", st)
		}
		return err
	}

	// gRPC接続の作成
	conn, err := grpc.DialContext(ctx, address,
		// セキュリティ設定
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// リトライインターセプター
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(retryOpts...)),
		// ロギングインターセプター
		grpc.WithUnaryInterceptor(loggingInterceptor),
		// トレーシングインターセプター
		grpc.WithUnaryInterceptor(tracingInterceptor),
		// エラーハンドリングインターセプター
		grpc.WithUnaryInterceptor(errorHandlingInterceptor),
		// ストリームリトライインターセプター
		grpc.WithStreamInterceptor(grpc_retry.StreamClientInterceptor(retryOpts...)),
		// KeepAliveパラメータ
		grpc.WithKeepaliveParams(keepAliveParams),
	)
	if err != nil {
		// 接続失敗時のログ出力
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}
	return conn, nil
}

// CloseClientConn はgRPCクライアント接続を閉じます
func CloseClientConn(conn *grpc.ClientConn) {
	// 接続クローズ失敗時のログ出力
	if err := conn.Close(); err != nil {
		log.Fatalf("failed to close connection: %v", err)
	}
}
