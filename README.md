# go_interpreter
Go言語で作るインタプリタ

# Limitation
インタプリタが対応していないこと
- マルチバイト文字（monkeyはASCII文字のみに対応）
    - l.chをbyteからruneに変更しなければならない。１文字を読み込むことがやや複雑になってしまう