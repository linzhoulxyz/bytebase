package mysql

import (
	"context"

	"github.com/antlr4-go/antlr/v4"
	"github.com/pkg/errors"

	mysql "github.com/bytebase/mysql-parser"

	storepb "github.com/bytebase/bytebase/backend/generated-go/store"
	"github.com/bytebase/bytebase/backend/plugin/advisor"
	mysqlparser "github.com/bytebase/bytebase/backend/plugin/parser/mysql"
)

var (
	_ advisor.Advisor = (*StatementMaximumStatementsInTransactionAdvisor)(nil)
)

func init() {
	advisor.Register(storepb.Engine_MYSQL, advisor.MySQLStatementMaximumStatementsInTransaction, &StatementMaximumStatementsInTransactionAdvisor{})
}

type StatementMaximumStatementsInTransactionAdvisor struct {
}

func (*StatementMaximumStatementsInTransactionAdvisor) Check(_ context.Context, checkCtx advisor.Context) ([]*storepb.Advice, error) {
	stmtList, ok := checkCtx.AST.([]*mysqlparser.ParseResult)
	if !ok {
		return nil, errors.Errorf("failed to convert to mysql parse result")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(checkCtx.Rule.Level)
	if err != nil {
		return nil, err
	}
	payload, err := advisor.UnmarshalNumberTypeRulePayload(checkCtx.Rule.Payload)
	if err != nil {
		return nil, err
	}
	checker := &statementMaximumStatementsInTransactionChecker{
		level:         level,
		title:         string(checkCtx.Rule.Type),
		limitMaxValue: payload.Number,
	}

	for _, stmt := range stmtList {
		checker.baseLine = stmt.BaseLine
		antlr.ParseTreeWalkerDefault.Walk(checker, stmt.Tree)
	}

	return checker.adviceList, nil
}

type statementMaximumStatementsInTransactionChecker struct {
	*mysql.BaseMySQLParserListener

	baseLine      int
	adviceList    []*storepb.Advice
	level         storepb.Advice_Status
	title         string
	text          string
	limitMaxValue int
}

func (checker *statementMaximumStatementsInTransactionChecker) EnterQuery(ctx *mysql.QueryContext) {
	checker.text = ctx.GetParser().GetTokenStream().GetTextFromRuleContext(ctx)
}
