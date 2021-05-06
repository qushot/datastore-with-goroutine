# datastore-with-goroutine

Cloud Datastoreから全Userを取得し、Userの持つTaskIDからそれぞれのTaskを取得するサンプル。
Cloud Traceで非同期処理になっているか確認している。

deploy
--
```shell
$ gcloud app deploy \
--project=your-project-id \
--version=datastoregoroutine \
--no-promote --quiet
```

エンドポイント
--
- /index: GUI
- /setup?num=30: テストデータを作成する。numにて作成する件数を指定する。デフォルト5件。
- /teardown: テストデータを全て削除する。
- /sync: 同期的に処理する。
- /async: 非同期的に処理する。
