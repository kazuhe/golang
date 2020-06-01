# golang
はじめてのGo!

## Git commit rule
コミットメッセージには必ずステータスを記載する。また、commitは可能な範囲で細かく行う。
``` bash
# [wip] 作業進行中
$ git commit -m 'wip:作業中の内容'

# [fix] バグ修正
$ git commit -m 'fix:バグの内容'

# [mod] バグ以外の修正
$ git commit -m 'mod:修正内容'

# [add] 機能 / ファイル追加
$ git commit -m 'add:追加する機能'

# [cln] リファクタリング / ファイル整理
$ git commit -m 'cln:ざっくり内容'

# [rvt] 変更の取り消し
$ git commit -m 'rvt:取り消す内容'
```