package templates

import (
	"embed"
	"github.com/roessland/withoutings/internal/repos/db"
	"github.com/roessland/withoutings/internal/service/sleep"
	"github.com/roessland/withoutings/internal/withoutings/adapters/withingsapi"
	"html/template"
	"io"
)

//go:embed templates
var fs embed.FS

type Templates struct {
	template *template.Template
}

func LoadTemplates() Templates {
	templates := Templates{}
	t, err := template.ParseFS(fs, "*/**")
	if err != nil {
		panic(err)
	}
	templates.template = t
	return templates
}

type HomePageVars struct {
	Error   string
	Account *db.Account
}

func (t Templates) RenderHomePage(w io.Writer, account *db.Account) error {
	return t.template.ExecuteTemplate(w, "homepage.gohtml", HomePageVars{
		Account: account,
	})
}

type SleepSummariesVars struct {
	Error     string
	Token     *withingsapiadapter.Token
	SleepData interface{}
}

func (t Templates) RenderSleepSummaries(w io.Writer, sleepData *sleep.GetSleepSummaryOutput, err string) error {
	return t.template.ExecuteTemplate(w, "sleepsummaries.gohtml", SleepSummariesVars{
		SleepData: sleepData,
		Error:     err,
	})
}

type RefreshAccessTokenVars struct {
	Error            string
	Token, PrevToken *withingsapiadapter.Token
}

func (t Templates) RenderRefreshAccessToken(w io.Writer, token, prevToken *withingsapiadapter.Token) error {
	return t.template.ExecuteTemplate(w, "refreshaccesstoken.gohtml", RefreshAccessTokenVars{
		Token:     token,
		PrevToken: prevToken,
	})
}
