# CSF: Credits Score Format

CSF 是一个用于表示如同[Credits / Frums 的 BGA]()一样使用纯文本构成的 BGA 的格式。若阅读之后仍有不解之处，还请参阅[用 CSF 来表现 Credits / Frums BGA 的例子](../static/csf-root)。

一般的音乐视频的目录结构如下：

```
csf_root_dir
├─ meta.yaml          # 定义歌曲的元数据
├─ data/              # 存放将在画面上显示的文本数据
└─ scores/            # 用于设定文本数据表示的时间和位置等等
```

## meta.yaml

使用 yaml 格式定义的歌曲元数据。

```
BPM: 179                      # 每分钟拍数，即 BPM
AudioFilePath: credits.mp3    # 音乐文件相对于 meta.yaml 的路径
AudioOffsetSec: 1.341         # 音乐和画面的偏移。对于本例，当音乐播放到 1.341 秒时，才显示第一帧
```

## data/

用于存放记录了将要显示到画面上的文本数据的文件。
您可以配置多个文件，也可使用类似于 `data/hoge/fuga_data` 的目录结构。
文本将会根据`scores/`文件夹内的`.score`文件来决定显示的时机。

## scores/

用于存放定义文本数据的显示位置和显示时机等，以`.score`为后缀名的文件。
您可以配置多个以`.score`为后缀的文件，也可使用类似于 `scores/hoge/fuga.score` 的目录结构。

### `.score`文件

`.score`文件的格式如下：

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

`.score`文件由多个小节构成的；`---`作为小节线，用于分割不同的小节。

小节中的一行可以为 DisplayCommand，Command，Comment 之一。

#### DisplayCommand

当小节内有 N 个 DisplayCommand 时，在画面上显示的文字将只持续 N 分音符的时长。

DisplayCommand 分为 DataDisplayCommand 和 InlineDisplayCommand 两种。

DataDisplayCommand 允许你在画面上显示`data/`文件夹内配置的文本文件的内容。例如需要显示`data/hoge/fuga_data`的时候，输入`hoge/fuga_data`就可以了。但是，如果文件不存在，则会直接显示这一行。

InlineDisplayCommand 允许你显示使用 「`"`」 包括的文字。例如`"hoge"`会在画面上显示`hoge`。

##### 例子

在 Credits / Frums 歌词的开头有一个 8 分间隔的「Fun-ding-for-this-pro-gram-was-made」，以下是表示它的方法。

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

注：虽然各行都没有「`"`」，但这其实是利用了当文件不存在时会将这行直接输出的特性，所以是没有问题的。

#### Command

一个 Command 以 `# ` 开始，用于立刻变更文本的显示位置和显示优先级，且不影响 DisplayCommand 的时机。
以下是所有的有效命令：

- `# MOVETO [x int] [y int]`
  - 将文本的显示位置移到`y`行`x`列。
  - 例子：`# MOVETO 1 2`（从这以后会从画面的2行1列的位置显示文本。）
- `# ZINDEX [i int]`
  - 变更文本的 zIndex（显示优先级）到`i`。当画面上有多个 DisplayCommand 的显示重叠时，优先显示优先级更高的那一个。用于处理重叠关系。
  - 例子：`# ZINDEX 5`（从这以后显示的文本的 zIndex 将会为 5。）
- `# FLIP vertical (on | off)`
  - 打开或关闭文本的左右反转模式。左右反转模式默认是关闭的。
  - 例子：`# FLIP vertical on`（从这以后显示的文字将会左右反转。）

#### Comment，BlankLine

由`/`开始的一行或者空行都将被忽略，不会影响画面的显示。
