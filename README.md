# go_interpreter

[Go 言語で作るインタプリタ](https://www.amazon.co.jp/Go%E8%A8%80%E8%AA%9E%E3%81%A7%E3%81%A4%E3%81%8F%E3%82%8B%E3%82%A4%E3%83%B3%E3%82%BF%E3%83%97%E3%83%AA%E3%82%BF-Thorsten-Ball/dp/4873118220/ref=sr_1_1?adgrpid=52270124614&dchild=1&gclid=Cj0KCQjw4cOEBhDMARIsAA3XDRhYel8yWP6oLipgfaac_w5B24eVMHQbuZ20E6IXibxt5m_j2npdajsaArhjEALw_wcB&hvadid=338518266894&hvdev=c&hvlocphy=1009314&hvnetw=g&hvqmt=e&hvrand=12370438417398361186&hvtargid=kwd-456677309977&hydadcr=27267_11561158&jp-ad-ap=0&keywords=go%E8%A8%80%E8%AA%9E%E3%81%A7%E3%81%A4%E3%81%8F%E3%82%8B%E3%82%A4%E3%83%B3%E3%82%BF%E3%83%97%E3%83%AA%E3%82%BF&qid=1620139864&sr=8-1)

# Grammer

以下のような C 言語ライクな構文を持つなんちゃってプログラミング言語を自作する

```javascript

let x = 10;
let y = 15;

let add = fn(a, b) {
    return a + b;
};

```

[要加筆]

### 文

- 変数束縛: let
- return 文: return

### 式

let 文と return 文を覗くと全てが式として評価される設計になっている。
識別子も式

```javascript
add(2, 3) - 5;
!ture;
5 * 5;
4 * 3;
```

# Limitation

インタプリタが対応していないこと

- マルチバイト文字（monkey は ASCII 文字のみに対応）
  - l.ch を byte から rune に変更しなければならない。１文字を読み込むことがやや複雑になってしまう
- 整数以外の数字
  - 簡単なものをまずは作る。今後の課題（ブログのネタにでもしよう）

# Thanks

I appreciate you sharing knowledge, Thorsten Ball and Yoji Shidara. Before reading [this book](https://www.amazon.co.jp/Go%E8%A8%80%E8%AA%9E%E3%81%A7%E3%81%A4%E3%81%8F%E3%82%8B%E3%82%A4%E3%83%B3%E3%82%BF%E3%83%97%E3%83%AA%E3%82%BF-Thorsten-Ball/dp/4873118220/ref=sr_1_1?adgrpid=52270124614&dchild=1&gclid=Cj0KCQjw4cOEBhDMARIsAA3XDRhYel8yWP6oLipgfaac_w5B24eVMHQbuZ20E6IXibxt5m_j2npdajsaArhjEALw_wcB&hvadid=338518266894&hvdev=c&hvlocphy=1009314&hvnetw=g&hvqmt=e&hvrand=12370438417398361186&hvtargid=kwd-456677309977&hydadcr=27267_11561158&jp-ad-ap=0&keywords=go%E8%A8%80%E8%AA%9E%E3%81%A7%E3%81%A4%E3%81%8F%E3%82%8B%E3%82%A4%E3%83%B3%E3%82%BF%E3%83%97%E3%83%AA%E3%82%BF&qid=1620139864&sr=8-1), I couldn't imagine I can write interpreter by myself. Thanks a lot!
