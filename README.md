# S3+Lambda発火のサンプル

## 手順
```
# localstack起動
TMPDIR=/private$TMPDIR docker-compose up -d

# ビルド
GOOS=linux GOARCH=amd64 go build -o main main.go
zip main.zip main
rm main

# Lambda作成
aws lambda create-function \
  --endpoint-url=http://localhost:4574 \
  --function-name s3-trigger-test \
  --runtime go1.x \
  --handler main \
  --zip-file fileb://main.zip \
  --role r1

# バケット作成
aws --endpoint-url=http://localhost:4572 s3 mb s3://test-bucket

# Lambdaに権限付与
aws lambda add-permission \
  --endpoint-url=http://localhost:4574 \
  --region us-east-1 \
  --function-name s3-trigger-test \
  --statement-id s3-put-event \
  --principal s3.amazonaws.com \
  --action "lambda:InvokeFunction" \
  --source-arn arn:aws:s3:::test-bucket

# バケットのイベント通知設定
aws s3api put-bucket-notification-configuration \
  --endpoint-url=http://localhost:4572 \
  --bucket test-bucket \
  --notification-configuration file://s3test-event.json

# S3アップロード
aws s3 cp main.go s3://test-bucket \
  --endpoint-url=http://localhost:4572
```

## ログ確認
```
aws logs describe-log-groups \
  --endpoint-url=http://localhost:4586 \
  --log-group-name-prefix "/aws/lambda" \
  --region us-east-1

aws logs tail --follow /aws/lambda/s3-trigger-test --endpoint-url http://localhost:4586 --region=us-east-1
```

## バケット確認
```
aws s3 ls s3://test-bucket \
  --endpoint-url=http://localhost:4572

aws lambda list-functions \
  --endpoint-url=http://localhost:4574

## Lambda発火
aws lambda invoke \
  --endpoint-url=http://localhost:4574 \
  --function-name s3-trigger-test \
  out --log-type Tail; rm out
```
