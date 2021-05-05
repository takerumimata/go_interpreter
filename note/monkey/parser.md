# 構文解析 - パーサーを理解する
じっくり、コードを読むことで動きを理解する

```javascript
// 擬似コード

function parserProgram(){
    program = newProgramASTNode()

    advanceTokens()

    for ( currentToken() != EOF_TOKEN) {
        // statementはまずnullで初期化する
        statement = null

        // 現在のtokenとこの言語で定義されているtokenとをマッチングさせて、statement変数に束縛する
        if (currentToken() == LET_TOKEN) {
            statement = parseLetStatement()
        } else if (currentToken() == RETURN_TOKEN) {
            statement = ParseReturnStatement()
        } else if (currentToken() == IF_TOKEN) {
            statement = parseIfStatement()
        }

        advanceTokens()
    }

    return program
}

function parseLetStatement() {
    // 次のtokenを読む
    advanceTokens()

    // 変数名をidentとして読み込む
    identifier = parseIdentifier()

    // 次のtokenを読み込む
    advaneTokens()

    // ここにはequalがきているはず。
    if currentToken() != EQUAL_TOKEN {
        parseExpression("no equal sign!")
    }

    // 次のtokenを読み込む
    advanceTokens()

    // 整数, + などをパース
    value = parseExpression()

    variableStatement = newVariableStatementASTNode()
    variableStatement.identifier = identifier
    variableStatement.value = value
    return variableStatement
}


// identifier（変数名）をパースする
function parseIdentifier() {
    identifier = newIdentifierASTNode()
    identifier.token = currentToken()
    return identifier
}

// 整数, +、;、(などをパースする
function parseExpression() {
    if (currentToken() == INTEGER_TOKEN) {
        if(nextToken() == PLUS_TOKEN) {
            return parseOperatorExpression()
        } else if(nextToken() == SEMICOLON_TOKEN) {
            return parseIntegerLiteral()
        } 
    } else if ( currentToken() == LEFT_PAREN) {
            return parseGroupedExpression()
    }
    // [...]
}

function parseOperatorExpression() {
    operatorExpression = newOperatorExpression()

    operatorExpression.left = parseIntegerLiteral()
    operatorExpression.operator = currentToken()
    operatorExpression.right = parseExpression()

    return operatorExpression()
}


// [...]
```