package config

import (
	"fmt"
	"text/template"

	"github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	tb "gopkg.in/telebot.v3"
)

type RunType string

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"

	ProjectName          string = "flowerss"
	BotToken             string
	Socks5               string
	TelegraphToken       []string
	TelegraphAccountName string
	TelegraphAuthorName  string = "flowerss-bot"
	TelegraphAuthorURL   string

	// EnableTelegraph 是否启用telegraph
	EnableTelegraph       bool = false
	PreviewText           int  = 0
	DisableWebPagePreview bool = false
	mysqlConfig           *mysql.Config
	SQLitePath            string
	EnableMysql           bool = false

	// UpdateInterval rss抓取间隔
	UpdateInterval int = 10

	// ErrorThreshold rss源抓取错误阈值
	ErrorThreshold uint = 100

	// MessageTpl rss更新推送模版
	MessageTpl *template.Template

	// MessageMode telegram消息渲染模式
	MessageMode tb.ParseMode

	// TelegramEndpoint telegram bot 服务器地址，默认为空
	TelegramEndpoint string = tb.DefaultApiURL

	// UserAgent User-Agent
	UserAgent string

	// RunMode 运行模式 Release / Debug
	RunMode RunType = ReleaseMode

	// AllowUsers 允许使用bot的用户
	AllowUsers []int64

	// DBLogMode 是否打印数据库日志
	DBLogMode bool = false
)

const (
	defaultMessageTplMode = tb.ModeHTML
	defaultMessageTpl     = `
{{ if .PreviewText }}<b>{{.ContentTitle}}</b>

{{.PreviewText}}...
{{if .EnableTelegraph}}
<a href="{{.RawLink}}">查看原文</a> | <a href="{{.TelegraphURL}}">Telegraph</a>
{{- else }}
<a href="{{.RawLink}}">查看原文</a>
{{- end }}
{{- else }}{{if .EnableTelegraph}}<a href="{{.RawLink}}"><b>{{.ContentTitle}}</b></a> | <a href="{{.TelegraphURL}}">Telegraph</a>
{{- else }}<a href="{{.RawLink}}"><b>{{.ContentTitle}}</b></a>{{- end }}{{- end }}
{{.Tags}}
`
	defaultMessageMarkdownTpl = `
{{ if .PreviewText }}** {{.ContentTitle}} **

{{.PreviewText}}...
{{if .EnableTelegraph}}
[查看原文]({{.RawLink}}) |  [Telegraph]({{.TelegraphURL}})
{{- else }}
[查看原文]({{.RawLink}})
{{- end }}
{{- else }}{{if .EnableTelegraph}}[** {{.ContentTitle}} **]({{.RawLink}}) |  [Telegraph]({{.TelegraphURL}})
{{- else }}[** {{.ContentTitle}} **]({{.RawLink}}){{- end}}{{- end }}
{{.Tags}}
`
	TestMode    RunType = "Test"
	ReleaseMode RunType = "Release"
)

type TplData struct {
	SourceTitle     string
	ContentTitle    string
	RawLink         string
	PreviewText     string
	TelegraphURL    string
	Tags            string
	EnableTelegraph bool
}

func AppVersionInfo() (s string) {
	s = fmt.Sprintf("version %v, commit %v, built at %v", version, commit, date)
	return
}

// GetString get string config value by key
func GetString(key string) string {
	var value string
	if viper.IsSet(key) {
		value = viper.GetString(key)
	}

	return value
}

func GetMysqlDSN() string {
	return mysqlConfig.FormatDSN()
}
