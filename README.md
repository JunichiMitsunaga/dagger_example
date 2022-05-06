# Dagger go example

daggerの初期化
```sh
❯ dagger project init
Project initialized! To install dagger packages, run `dagger project update`

❯ dagger project update
Project updated
```

テスト実行例
```sh
❯ dagger do test
[✔] actions.test
[✔] client.filesystem."./hello".read
```

ビルド実行例
```sh
[✔] actions.build.
[✔] client.filesystem."./hello".read
[✔] actions.build.container.export
```
