package constant

type ErrorCode string

const (
	// 400系エラーコード
	GP4000 ErrorCode = "GP4000" // 指定されたコンテンツが存在しない
	GP4100 ErrorCode = "GP4100" // Headerエラー(必須Headerが不足している)

	// 500系エラーコード
	GP5000 ErrorCode = "GP5000" // サーバー内部エラー
)
