# CSF: Credits Score Format

CSF は，[Credits / Frums の BGA]()のような，テキストベースのミュージックビデオを表現するためのフォーマットです．本ドキュメントだけでわからない部分は，[Credits / Frums BGA を CSF で表現した例](../static/csf-root)も参考にしてください．

1 つのミュージックビデオは，以下のようなディレクトリ構成で表されます．

```
csf_root_dir
├─ meta.yaml（楽曲のメタ情報を定義します）
├─ data/（画面に表示するテキストデータを配置します）
└─ scores/（テキストデータを表示する位置やタイミングなどを配置します）
```

## meta.yaml

曲のメタ情報を yaml 形式で定義します．

```
BPM: 179（曲のテンポ）
AudioFilePath: credits.mp3（音声ファイルへの，meta.yamlからの相対パス）
AudioOffsetSec: 1.341（音声と画面表示タイミングのオフセット．この例では，音声ファイルの1.341秒に，画面の1小節目が表示される）
```

## data/

画面に表示するテキストデータが書かれたファイルを`data/`内に配置します．ファイルは複数配置できます．`data/hoge/fuga_data`のようにサブディレクトリに配置しても構いません．

ここで配置されたデータは`scores/`内の`.score`ファイルによって，特定のタイミングで画面に表示されます（後述）．

## scores/

テキストデータの表示位置やタイミングなどを定義する`.score`ファイルを配置します．`.score`ファイルは複数配置できます．また，`scores/hoge/fuga.score`のようにサブディレクトリに配置しても構いません．

### `.score`ファイル

`.score`ファイルは次のような形式です．

```
[DisplayCommand] | [Command] | [Comment] | [BlankLine]
[DisplayCommand] | [Command] | [Comment] | [BlankLine]
...
---
[DisplayCommand] | [Command] | [Comment] | [BlankLine]
[DisplayCommand] | [Command] | [Comment] | [BlankLine]
...
---
...
```

`.score`ファイルは，複数の小節によって構成されます．`---`は小節線であり，小節の区切りを表します．

小節の各行には DisplayCommand，Command，Comment のいずれかを書きます．

#### DisplayCommand

小節内に DisplayCommand が N 個あるとき，画面に文字を N 分音符の長さだけ表示します．

DisplayCommand には，DataDisplayCommand と InlineDisplayCommand の 2 種類があります．

DataDisplayCommand は，`data/`内に配置したテキストファイルを画面に表示します．例えば`data/hoge/fuga_data`を表示する場合は，`hoge/fuga_data`と書きます．ただし，ファイルが存在しない場合はパスがそのまま画面に表示されます．

InlineDisplayCommand は，「`"`」で囲まれたテキストを画面に表示します．例えば，`"hoge"`と書くと`hoge`と画面に表示されます．

##### 例

Credits / Frums の歌詞の冒頭部分における，8 分間隔の「Fun-ding-for-this-pro-gram-was-made」は次のように表せます．

```
---
Fun
Funding
Funding for
Funding for this
pro
program
program was
program was made
---
```

※各行に「`"`」がありませんが，これはテキストファイルが存在しない場合に，パスがそのまま表示されるのを利用しています．

#### Command

Command は`# `で始まり，テキストの表示位置や順序を即座に変更します（DisplayCommand の表示タイミングに影響しません）．以下のコマンドが有効です．

- `# MOVETO [x int] [y int]`
  - テキストの表示位置を画面の`y`行`x`文字目に変更します
  - 例：`# MOVETO 1 2`（以降，画面の 2 行 1 文字目に移動してテキストを表示します）
- `# ZINDEX [i int]`
  - テキストの zIndex（表示順序）を`i`に変更します．画面上で複数の DisplayCommand が重なった場合，`i`が大きいものが優先して表示されます．
  - 例：`# ZINDEX 5`（以降のテキストの zIndex を 5 にします）
- `# FLIP vertical (on | off)`
  - テキストの左右反転モードを有効・無効にします．左右反転モードは，デフォルトでは無効です．
  - 例：`# FLIP vertical on`（以降のテキストを左右反転させて表示します）

#### Comment，BlankLine

`/`から始まる行や空行は無視され，画面の表示に影響しません．
