# pj-nyegs

GolangでWEBサービスを作る独学用

基本的にはDockerで開発する環境を構築
やりたいことは

* TESTの作り方
* CI/CDの仕組み周り
* フロント/バックの繋ぎこみ（別リポジトリを想定）

## Hello Worldを動かす
1. docker環境を構築

    Win or Mac で導入方法異なるので各自でお願いします。

1. 作業ディレクトリを作成し移動

    ex) workディレクトリを作成して移動
    ```
    # mkdir work
    # cd work
    ```

1. Git clone 

    ```
    # git clone https://github.com/masa-hashi/pj-nyegs.git
    ```

1. プロジェクトディレクトリに移動
 
    ```
    # cd pj-nyegs
    ```

     (以下指定がない場合はこのpj-nyegsディレクトリでコマンド実行)

1. docker build & バックグラウンド起動

    ```
    # docker-compose up -d
    ```

1. Hello World 動作確認

    ```
    # docker-compose exec app go run main.go
    ```

    「Hello, World!」が表示されることを確認。

## 止め方

    ```
    # docker-compose stop
    ```

ひとまず、以上。先は長いぞ。