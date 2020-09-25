package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"html/template"
	"net/http"
	"os"
	"os/user"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/evercyan/gocli/color"
	"github.com/evercyan/gocli/table"
	"github.com/evercyan/letitgo/crypto"
	"github.com/evercyan/letitgo/util"
	"github.com/mozillazg/request"
	"github.com/peterh/liner"
	cli "github.com/urfave/cli/v2"
)

var (
	lcAllURL         = "https://leetcode-cn.com/api/problems/all/" // 问题列表地址
	lcGraphqlURL     = "https://leetcode-cn.com/graphql"           // 问题数据地址
	lcQuestionURL    = "https://leetcode-cn.com/problems/%s/"      // 问题页面地址
	lcTagURL         = "https://leetcode-cn.com/tag/%s/"           // 标签页面地址
	lcQuestionSetURL = "https://leetcode-cn.com/problemset/all/"   // 题库首页
	lcDifficulty     = []string{"", "简单", "中等", "困难"}              // 难度类型
	commandList      = []string{"help", "h"}                       // 命令列表
)

/**
 * ******************************** 模板配置
 */

// Repo README 模板
var tplReadme = `<div align="center">

![leetcli](https://raw.githubusercontent.com/evercyan/cantor/master/resource/69/69f055fa7ccfe73114bf6608a2789d8f.png)

[![travis-ci](https://www.travis-ci.org/evercyan/leetcli.svg?branch=master)](https://www.travis-ci.org/evercyan/leetcli)
[![codecov](https://codecov.io/gh/evercyan/leetcli/branch/master/graph/badge.svg?token=RbJTUtAlvl)](https://codecov.io/gh/evercyan/leetcli)
[![goreportcard](https://goreportcard.com/badge/github.com/evercyan/leetcli)](https://goreportcard.com/report/github.com/evercyan/leetcli)

leetcode 刷题小工具, 生成 README, 答题文件, 测试文件等
</div>

---

## Usage

` + "```sh" + `
go get -v github.com/evercyan/leetcli
leetcli --help
` + "```" + `

---

## Tag List

{{.DrawQuestionTagList}}

---

## Question List

{{.DrawQuestionList}}

`

// 题目 README 模板
var tplQuestionReadme = `# [{{.FQID}}. {{.Title}}]({{.Link}})

` + "`" + `[{{.Difficulty}}]` + "`" + ` {{.Tags}}

---

{{.Content}}

`

// 题目 go 测试模板
var tplQuestionGoTestFile = `package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		inputs  []interface{}
		expects []interface{}
	}{
		{
			[]interface{}{
				"input",
			},
			[]interface{}{
				"output",
			},
		},
	}
	for _, c := range cases {
		t.Run("FuncToReplace", func(t *testing.T) {
			ret := FuncToReplace(c.inputs[0].(string))
			assert.Equal(t, c.expects[0].(string), ret)
		})
	}
}`

/**
 * ******************************** 公共函数
 */

func formatContent(str string) string {
	// 处理 img 标签, 生成 markdown 格式, ![](图片链接地址)
	rImg, _ := regexp.Compile(`<img[^<>]*src="([^"]+)"[^<>]*>`)
	str = rImg.ReplaceAllString(str, "```\n\n![]($1)\n\n```")
	// 处理 &nbsp; &lt; 等 html escape
	str = html.UnescapeString(str)
	// 处理多行换行
	str = regexp.MustCompile(`(\n){3,}`).ReplaceAllString(str, "\n\n")
	// 处理标签
	re, _ := regexp.Compile("\\<[^<>]+\\>")
	return fmt.Sprintf("```json\n%s\n```", re.ReplaceAllString(str, ""))
}

func getQustionPath(fid string, id int, slug string) string {
	qid := int(util.ToInt64(fid))
	if qid == 0 {
		qid = id
	}
	tpl := "/src/%d-%s/"
	if id < 10000 {
		tpl = "/src/%04d-%s/"
	}
	return fmt.Sprintf(tpl, qid, slug)
}

func getConfigFile() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}
	appPath := user.HomeDir + "/.leetcli"
	if !util.IsExist(appPath) {
		os.Mkdir(appPath, os.ModePerm)
	}
	return appPath + "/config.json", nil
}

func success(texts ...string) {
	color.New(texts...).Color(color.Green).Render()
}

func notice(texts ...string) {
	color.New(texts...).Color(color.Yellow).Render()
}

func fail(texts ...string) {
	color.New(texts...).Color(color.Red).Render()
}

func show(prefixStr, successStr, failStr string) {
	content := color.New(prefixStr).Color(color.Yellow).Content() + " "
	if successStr != "" {
		content += color.New(successStr).Color(color.Green).Content()
	} else {
		content += color.New(failStr).Color(color.Red).Content()
	}
	fmt.Println(content)
}

/**
 * ******************************** leetcode
 */

type leetCode struct {
	QuestionList    []lcQuestionInfo
	QuestionMap     map[string]lcQuestionInfo
	QuestionTagList []lcQuestionTagInfo
}

type lcQuestionInfo struct {
	FQID       string `json:"fqid"`       // 前端 id
	QID        int    `json:"qid"`        // id
	Title      string `json:"title"`      // 名称
	Slug       string `json:"slug"`       // 标识
	Link       string `json:"link"`       //链接
	Difficulty string `json:"difficulty"` // 困难度
}

type lcQuestionTagInfo struct {
	Title string
	Link  string
	Count int
}

type lcOriginQuestionInfo struct {
	Status    interface{} `json:"status"`
	Progress  int64       `json:"progress"`
	Frequency int64       `json:"frequency"`
	Stat      struct {
		QuestionId          int    `json:"question_id"`
		QuestionHide        bool   `json:"question__hide"`
		QuestionTitle       string `json:"question__title"`
		QuestionTitleSlug   string `json:"question__title_slug"`
		FrontendQuestionId  string `json:"frontend_question_id"`
		TotalColumnArticles int    `json:"total_column_articles"`
		IsNewQuestion       bool   `json:"is_new_question"`
	} `json:"stat"`
	Difficulty struct {
		Level int `json:"level"`
	} `json:"difficulty"`
}

type lcQuestionDetail struct {
	Slug     string                       `json:"slug"`
	Title    string                       `json:"title"`
	Content  string                       `json:"content"`
	LangList []string                     `json:"lang_list"`
	TagList  []map[string]string          `json:"tag_list"`
	CodeMap  map[string]map[string]string `json:"code_list"`
}

func (l *leetCode) getQuestionList() error {
	req := request.NewRequest(new(http.Client))
	resp, err := req.Get(lcAllURL)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	respJson, err := resp.Json()
	if err != nil {
		return err
	}
	qList1, _ := respJson.Get("stat_status_pairs").Array()
	qList2 := []lcOriginQuestionInfo{}
	if err := json.Unmarshal([]byte(crypto.JsonEncode(qList1)), &qList2); err != nil {
		return err
	}
	questionMap := make(map[string]lcQuestionInfo)
	for i := len(qList2) - 1; i >= 0; i-- {
		questionInfo := lcQuestionInfo{
			FQID:       qList2[i].Stat.FrontendQuestionId,
			QID:        qList2[i].Stat.QuestionId,
			Title:      qList2[i].Stat.QuestionTitle,
			Slug:       qList2[i].Stat.QuestionTitleSlug,
			Link:       fmt.Sprintf(lcQuestionURL, qList2[i].Stat.QuestionTitleSlug),
			Difficulty: lcDifficulty[qList2[i].Difficulty.Level],
		}
		l.QuestionList = append(l.QuestionList, questionInfo)
		questionMap[questionInfo.Slug] = questionInfo
	}
	l.QuestionMap = questionMap
	return nil
}

func (l *leetCode) getQuestionTagList() error {
	resp, err := http.Get(lcQuestionSetURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}
	doc.Find("#all-topic-tags a.tags-btn").Each(func(i int, s *goquery.Selection) {
		item := lcQuestionTagInfo{}
		item.Title = strings.TrimSpace(s.Find("span.text-gray").Text())
		item.Link, _ = s.Attr("href")
		item.Count = int(util.ToInt64(strings.TrimSpace(s.Find("span.badge").Text())))
		l.QuestionTagList = append(l.QuestionTagList, item)
	})
	return nil
}

func (l *leetCode) getQuestionDetail(slug string) (*lcQuestionDetail, error) {
	req := request.NewRequest(new(http.Client))
	req.Headers = map[string]string{
		"Content-Type": "application/json",
	}
	req.Json = map[string]interface{}{
		"operationName": "questionData",
		"query":         "query questionData($titleSlug: String!) {question(titleSlug: $titleSlug) {questionId questionFrontendId title titleSlug content translatedTitle translatedContent topicTags {name slug translatedName} codeSnippets {lang langSlug code}}}",
		"variables": map[string]string{
			"titleSlug": slug,
		},
	}
	resp, err := req.Post(lcGraphqlURL)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	respJson, _ := resp.Json()

	questionDetail := &lcQuestionDetail{}
	// 基础信息
	questionDetail.Slug = slug
	questionDetail.Title, _ = respJson.Get("data").Get("question").Get("translatedTitle").String()
	questionDetail.Content, _ = respJson.Get("data").Get("question").Get("translatedContent").String()
	// 处理 html 标签
	questionDetail.Content = formatContent(questionDetail.Content)
	// 标签
	tagListTmp, _ := respJson.Get("data").Get("question").Get("topicTags").Array()
	tagList := []map[string]string{}
	if err := json.Unmarshal([]byte(crypto.JsonEncode(tagListTmp)), &tagList); err != nil {
		return nil, err
	}
	questionDetail.TagList = tagList
	// 代码片段
	codeListTmp, _ := respJson.Get("data").Get("question").Get("codeSnippets").Array()
	codeList := []map[string]string{}
	if err := json.Unmarshal([]byte(crypto.JsonEncode(codeListTmp)), &codeList); err != nil {
		return nil, err
	}
	// 语言代码
	codeMap := map[string]map[string]string{}
	// 支持语言类型
	langList := []string{}
	for _, item := range codeList {
		langList = append(langList, item["langSlug"])
		codeMap[item["langSlug"]] = item
	}
	questionDetail.CodeMap = codeMap
	questionDetail.LangList = langList
	return questionDetail, nil
}

func (l *leetCode) Init() error {
	err := l.getQuestionList()
	if err != nil {
		return err
	}
	err = l.getQuestionTagList()
	if err != nil {
		return err
	}
	return nil
}

/**
 * ******************************** 答题文件处理
 */

// 默认答题配置
var (
	LangConf = map[string]map[string]string{
		"golang": {
			"file":        "solution.go",
			"fileTpl":     "package solution\n\n%s",
			"testfile":    "solution_test.go",
			"testfileTpl": tplQuestionGoTestFile,
		},
	}
	LangFile    = "solution.%s"
	LangFileTpl = "%s"
	LangSuffix  = map[string]string{
		"golang":     "go",
		"php":        "php",
		"c":          "c",
		"cpp":        "c",
		"java":       "java",
		"python":     "py",
		"python3":    "py",
		"javascript": "js",
		"mysql":      "sql",
		"bash":       "sh",
	}
	LangList = []string{
		"golang", "php", "c", "cpp", "java", "python", "python3", "javascript", "mysql", "bash",
	}
)

type leetCodeFile struct {
	Path string `json:"path" desc:"答题文件目录"`
	Lang string `json:"lang" desc:"默认编程语言"`
}

func (lf *leetCodeFile) Init() error {
	configFile, err := getConfigFile()
	if err != nil {
		return err
	}
	content := util.ReadFile(configFile)
	if content == "" {
		return nil
	}
	return json.Unmarshal([]byte(content), lf)
}

func (lf *leetCodeFile) Save() error {
	configFile, err := getConfigFile()
	if err != nil {
		return err
	}
	return util.WriteFile(configFile, crypto.JsonEncode(lf))
}

// 生成答题相关文件
func (lf *leetCodeFile) GenerateQuestion(slug string, lang string, questionDetail *lcQuestionDetail) (err error) {
	questionInfo, _ := lc.QuestionMap[slug]
	questionPath := getQustionPath(questionInfo.FQID, questionInfo.QID, questionInfo.Slug)
	questionPath = lf.Path + questionPath
	if !util.IsExist(questionPath) {
		err = os.MkdirAll(questionPath, 0755)
		if err != nil {
			return errors.New("创建答题文件目录失败")
		}
	}

	// 问题标签
	tagStr := ""
	for _, tag := range questionDetail.TagList {
		if tag["slug"] == "" {
			continue
		}
		tagStr += fmt.Sprintf(" [%s](%s) ", tag["translatedName"], fmt.Sprintf(lcTagURL, tag["slug"]))
		tagStr = strings.Trim(tagStr, " ")
	}

	// 问题模板数据
	questionReadmeInfo := struct {
		FQID       string
		Title      interface{}
		Link       string
		Content    interface{}
		Difficulty interface{}
		Tags       interface{}
	}{
		FQID:       questionInfo.FQID,
		Title:      template.HTML(questionInfo.Title),
		Link:       questionInfo.Link,
		Content:    template.HTML(questionDetail.Content),
		Difficulty: template.HTML(questionInfo.Difficulty),
		Tags:       template.HTML(tagStr),
	}

	// 生成答题文件 README.md
	questionReadmeFile := fmt.Sprintf("%s/README.md", questionPath)
	var b bytes.Buffer
	tpl := template.Must(template.New("question").Parse(tplQuestionReadme))
	err = tpl.Execute(&b, questionReadmeInfo)
	if err != nil {
		return errors.New("解析答题 README 模板失败")
	}
	if err = util.WriteFile(questionReadmeFile, string(b.Bytes())); err != nil {
		return errors.New("创建答题 README 文件失败")
	}

	// 生成答题文件, 答题测试文件
	file, fileTpl, testfile, testfileTpl := "", "", "", ""
	if _, ok := LangConf[lang]; ok {
		file, fileTpl = LangConf[lang]["file"], LangConf[lang]["fileTpl"]
		testfile, testfileTpl = LangConf[lang]["testfile"], LangConf[lang]["testfileTpl"]
	} else if _, ok := LangSuffix[lang]; ok {
		file = fmt.Sprintf(LangFile, LangSuffix[lang])
		fileTpl = LangFileTpl
	} else {
		// 未配置编程语言的, 生成完 README.md 直接结束
		return nil
	}

	// 创建答题文件
	questionFile := fmt.Sprintf("%s/%s", questionPath, file)
	if util.IsExist(questionFile) {
		show("答题文件已存在:", questionFile, "")
	} else {
		util.WriteFile(questionFile, fmt.Sprintf(fileTpl, questionDetail.CodeMap[lang]["code"]))
	}

	// 创建答题测试文件
	if testfile != "" {
		questionTestFile := fmt.Sprintf("%s/%s", questionPath, testfile)
		if util.IsExist(questionTestFile) {
			show("答题测试文件已存在:", questionTestFile, "")
		} else {
			// 替换测试文件中函数名称
			matchs := regexp.MustCompile(`func ([^\(]+)\(`).FindStringSubmatch(questionDetail.CodeMap[lang]["code"])
			if len(matchs) >= 2 {
				testfileTpl = strings.Replace(testfileTpl, "FuncToReplace", matchs[1], -1)
			}
			util.WriteFile(questionTestFile, testfileTpl)
		}
	}

	return nil
}

func (lf *leetCodeFile) GenerateReadme() error {
	var b bytes.Buffer
	tpl := template.Must(template.New("readme").Parse(tplReadme))
	err := tpl.Execute(&b, lf)
	if err != nil {
		return err
	}
	return util.WriteFile(lf.Path+"/README.md", string(b.Bytes()))
}

// [![数组](https://img.shields.io/badge/数组-99-red.svg)](https://shields.io/)
func (lf *leetCodeFile) DrawQuestionTagList() string {
	if len(lc.QuestionTagList) <= 0 {
		return ""
	}

	resp, preColor := "", ""
	for _, tag := range lc.QuestionTagList {
		tagLinks := strings.Split(tag.Link, "/")
		if len(tagLinks) < 4 {
			continue
		}
		url := fmt.Sprintf(lcTagURL, tagLinks[2])
		title := strings.Replace(tag.Title, " ", "", -1)
		color := "ff9985"
		if tag.Count > 200 {
			color = "8a0808"
		} else if tag.Count > 100 {
			color = "b80909"
		} else if tag.Count > 50 {
			color = "e64546"
		} else if tag.Count > 10 {
			color = "f57567"
		}
		if preColor != "" && preColor != color {
			resp += "\n"
		}
		preColor = color
		resp += fmt.Sprintf("[![%s](https://img.shields.io/badge/%s-%d-%s.svg?style=flat)](%s)\n", title, title, tag.Count, color, url)
	}
	return strings.Trim(resp, "\n")
}

// readme - 渲染题目表格, 只显示已解决的
func (lf *leetCodeFile) DrawQuestionList() string {
	questionList := lc.QuestionList
	if len(questionList) <= 0 {
		return ""
	}
	resp := fmt.Sprintln("|#|标题|难度|")
	resp += fmt.Sprintln("|:-:|:-|:-:|")
	for _, question := range questionList {
		questionPath := getQustionPath(question.FQID, question.QID, question.Slug)
		questionPath = "." + questionPath
		if !util.IsExist(questionPath) {
			continue
		}
		resp += fmt.Sprintf("|[%s](%s)|", question.FQID, question.Link)
		resp += fmt.Sprintf("[%s](%s)|", question.Title, questionPath)
		resp += fmt.Sprintf("%s|\n", question.Difficulty)
	}
	return strings.Trim(resp, "\n")
}

/**
 * ******************************** main
 */

var lc = new(leetCode)
var lcFile = new(leetCodeFile)

func main() {
	if err := lc.Init(); err != nil {
		fail("加载数据失败:", err.Error())
		return
	}
	if err := lcFile.Init(); err != nil {
		fail("加载配置失败:", err.Error())
		return
	}
	app := cli.NewApp()
	app.Name = "leetcli"
	app.Usage = "leetcode 刷题小工具, 生成 README.md, 答题文件, 测试文件等"
	app.Version = "v0.0.4"
	app.Commands = []*cli.Command{
		{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "配置",
			Subcommands: []*cli.Command{
				{
					Name:  "path",
					Usage: "设置答题文件目录 [eg: config path /tmp]",
					Action: func(c *cli.Context) error {
						path := c.Args().Get(0)
						if path == "" {
							show("答题文件目录:", lcFile.Path, "未设置")
							return nil
						}
						if !util.IsExist(path) {
							fail("答题文件目录不存在:", path)
							return nil
						}
						lcFile.Path = path
						if err := lcFile.Save(); err != nil {
							fail("更新答题文件目录失败:", err.Error())
							return nil
						}
						success("更新答题文件目录成功")
						return nil
					},
				},
				{
					Name:  "lang",
					Usage: "设置默认编程语言 [eg: config lang golang]",
					Action: func(c *cli.Context) error {
						lang := c.Args().Get(0)
						if lang == "" {
							show("默认编程语言:", lcFile.Lang, "未设置")
							return nil
						}
						if !util.InArray(lang, LangList) {
							fail("无效的编程语言:", lang)
							show("支持的编程语言:", crypto.JsonEncode(LangList), "")
							return nil
						}
						lcFile.Lang = lang
						if err := lcFile.Save(); err != nil {
							fail("更新默认编程语言失败:", err.Error())
							return nil
						}
						success("更新默认编程语言成功")
						return nil
					},
				},
			},
			Before: func(c *cli.Context) error {
				subCommand := c.Args().Get(0)
				if subCommand != "" && !util.InArray("config "+subCommand, commandList) {
					fail("无效命令, 输入 `help` 试一试")
					return errors.New("")
				}
				return nil
			},
		},
		{
			Name:    "readme",
			Aliases: []string{"r"},
			Usage:   "生成 README.md",
			Action: func(c *cli.Context) error {
				err := lcFile.GenerateReadme()
				if err != nil {
					fail(err.Error())
					return nil
				}
				success("生成 README.md 成功")
				return nil
			},
		},

		{
			Name:    "question",
			Aliases: []string{"q"},
			Usage:   "生成答题文件 [eg: question two-sum]",
			Action: func(c *cli.Context) error {
				slug := c.Args().Get(0)
				if slug == "" {
					show("请输入问题标识:", "[eg: question two-sum]", "")
					return nil
				}
				if _, ok := lc.QuestionMap[slug]; !ok {
					fail("无效的问题标识")
					return nil
				}
				questionDetail, err := lc.getQuestionDetail(slug)
				if err != nil {
					fail("读取问题信息失败:", err.Error())
					return nil
				}

				lang := ""
				if lcFile.Lang != "" && util.InArray(lcFile.Lang, questionDetail.LangList) {
					lang = lcFile.Lang
					show("已使用默认编程语言:", lang, "")
				} else {
					show("支持的编程语言:", crypto.JsonEncode(questionDetail.LangList), "")
					// 监听输入, 并对编程语言自动补全
					line := liner.NewLiner()
					defer line.Close()
					line.SetCtrlCAborts(true)
					line.SetCompleter(func(line string) (c []string) {
						for _, command := range questionDetail.LangList {
							if strings.HasPrefix(command, strings.ToLower(line)) {
								c = append(c, command)
							}
						}
						return
					})
					for {
						lang, err := line.Prompt("请输入编程语言 > ")
						if err == liner.ErrPromptAborted {
							os.Exit(0)
						}
						if lang != "" {
							break
						}
					}
					if !util.InArray(lang, questionDetail.LangList) {
						fail("无效的编程语言")
						return nil
					}
				}

				err = lcFile.GenerateQuestion(slug, lang, questionDetail)
				if err != nil {
					fail("生成答题文件失败:", err.Error())
					return nil
				}
				success("生成答题文件成功")
				return nil
			},
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "问题列表 [eg: list two-sum; 最多显示 20 条]",
			Action: func(c *cli.Context) error {
				keyword := c.Args().Get(0)
				if keyword == "" {
					show("请输入关键字:", "[eg: list two-sum]", "")
					return nil
				}
				list := [][]interface{}{}
				count := 0
				for _, item := range lc.QuestionList {
					if !strings.Contains(item.FQID, keyword) && !strings.Contains(item.Title, keyword) && !strings.Contains(item.Slug, keyword) {
						continue
					}
					if count >= 20 {
						break
					}
					count++
					list = append(list, []interface{}{
						item.FQID,
						item.Title,
						item.Slug,
						item.Difficulty,
					})
				}
				success(table.New(list).Header([]string{"id", "标题", "标识", "难度"}).Content())
				return nil
			},
		},
		{
			Name:    "exit",
			Aliases: []string{"e"},
			Usage:   "退出",
			Action: func(c *cli.Context) error {
				return cli.NewExitError("", 0)
			},
		},
	}

	// 命令列表
	for _, command := range app.Commands {
		commandList = append(commandList, command.Name)
		commandList = append(commandList, command.Aliases...)
		if len(command.Subcommands) > 0 {
			for _, subCommand := range command.Subcommands {
				commandList = append(commandList, command.Name+" "+subCommand.Name)
			}
		}
	}

	app.Action = func(c *cli.Context) error {
		line := liner.NewLiner()
		defer line.Close()

		// 处理命令补全
		line.SetCtrlCAborts(true)
		line.SetCompleter(func(line string) (c []string) {
			for _, command := range commandList {
				if strings.HasPrefix(command, strings.ToLower(line)) {
					c = append(c, command)
				}
			}
			return
		})

		for {
			commandLine, err := line.Prompt(fmt.Sprintf("%s > ", app.Name))
			// ctrl + c
			if err == liner.ErrPromptAborted {
				os.Exit(0)
			}
			commandLine = strings.Trim(commandLine, " ")
			if commandLine == "" {
				continue
			}
			cmdArgs := strings.Split(commandLine, " ")
			if len(cmdArgs) == 0 {
				continue
			}
			if !util.InArray(cmdArgs[0], commandList) {
				fail("无效命令, 输入 `help` 试一试")
				continue
			}
			if !strings.HasPrefix(cmdArgs[0], "config") && lcFile.Path == "" {
				fail("请先设置答题文件目录, 输入 `config` 试一试")
				continue
			}
			line.AppendHistory(commandLine)
			s := []string{os.Args[0]}
			s = append(s, cmdArgs...)
			c.App.Run(s)
		}
	}

	app.Run(os.Args)
}
