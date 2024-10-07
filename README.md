# Record_Toilet_Backend
トイレの記録を行うRESTAPIを実装したバックエンド

### GOOS=linux GOARCH=amd64 go build -o bootstrap main.go
buildファイルを作成　bootstrapという名前は固定

### sam package --output-template-file packaged.yaml --s3-bucket mytoiletrecord
samを用いてs３にbuildファイルをアップ

### sam deploy --template-file packaged.yaml --stack-name gin-lambda-api --capabilities CAPABILITY_IAM --region ap-northeast-1 
デプロイ