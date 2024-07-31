# api_sample

## Run

```bash
cd cmd/api
air
```

## Tips

```bash
```

## 注意
- airのrootをルートディレクトリにすることであらゆるモジュールの変更検知を行なっている
- モジュールごとにdiするのは難しかったので, 一つのdiにした.
  - e.g. todoモジュールのdiでbaseモジュールのdi済みのものを使うのは難しかった. wireのコード生成の限界かもしれない？

## Reference
- https://github.com/golang-standards/project-layout

## usecase
- base
  - authenticate
- todo
  - list
  - show
  - create
  - update
  - delete
