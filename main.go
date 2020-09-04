package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/evercyan/letitgo/crypto"
	"github.com/evercyan/letitgo/util"
	"github.com/mozillazg/request"
	"github.com/peterh/liner"
	"github.com/urfave/cli"
)

/**
 * ********************************
 * 模板配置
 * ********************************
 */

// Repo README 模板
var tplReadme = `# leetcli

[![996.icu](https://img.shields.io/badge/link-996.icu-red.svg)](https://996.icu)
[![Build Status](https://www.travis-ci.org/evercyan/leetcli.svg?branch=master)](https://www.travis-ci.org/evercyan/leetcli)
[![codecov](https://codecov.io/gh/evercyan/leetcli/branch/master/graph/badge.svg?token=RbJTUtAlvl)](https://codecov.io/gh/evercyan/leetcli)

> leetcode 刷题小助手, 帮助生成 readme, 答题文件, 答题测试文件等

---

## Tag

{{.DrawQuestionTagList}}

---

## Question

{{.DrawQuestionList}}

#### Snapshot

![leetcli](https://raw.githubusercontent.com/evercyan/cantor/master/resource/80/803bac1363e065a5e0fa7f8ac9d6db6a.png)

---
`

// 题目描述 README 模板
var tplQuestionReadme = `## [{{.FQID}}. {{.Title}}]({{.Link}})

---

{{.Difficulty}}

{{.Tags}}

---

#### 题目描述

{{.Content}}

---
`

// 题目 go 测试模板
var tplQuestionGoTestFile = `package solution

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		name    string
		inputs  []interface{}
		expects []interface{}
	}{
		{
			"Test",
			[]interface{}{
				"input",
			},
			[]interface{}{
				"output",
			},
		},
	}
	index := 1
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := FuncToReplace(c.inputs[0].(string))
			if !reflect.DeepEqual(ret, c.expects[0].(string)) {
				t.Fatalf("FAIL ----> name: %v-%v, inputs: %v, expects: %v, ret: %v", c.name, index, c.inputs, c.expects[0], ret)
			}
		})
		index++
	}
}`

/**
 * ********************************
 * 公共函数
 * ********************************
 */

func formatHtml(s string) string {
	// 先处理 img 标签, 生成 markdown 格式, ![](图片链接地址)
	rImg, _ := regexp.Compile(`<img[^<>]*src="([^"]+)"[^<>]*>`)
	s = rImg.ReplaceAllString(s, "```\n\n![]($1)\n\n```")
	replaceMap := map[string]string{
		"&amp;quot;": "\"",
		"&amp;lt;":   "<",
		"&amp;gt;":   ">",
		"&amp;ge;":   ">=",
		"&amp;le;":   "<=",
		"&amp;nbsp;": " ",
		"&amp;amp;":  "&",
		"&amp;#39;":  "'",
		"&amp;#43;":  "+",
		"&quot;":     "\"",
		"&lt;":       "<",
		"&gt;":       ">",
		"&ge;":       ">=",
		"&le;":       "<=",
		"&nbsp;":     " ",
		"&amp;":      "&",
		"&#39;":      "'",
		"&#43;":      "+",
		"\n\n\n\n\n": "\n\n",
		"\n\n\n\n":   "\n\n",
		"\n\n\n":     "\n\n",
	}
	for k, v := range replaceMap {
		s = strings.Replace(s, k, v, -1)
	}
	re, _ := regexp.Compile("\\<[^<>]+\\>")
	return fmt.Sprintf("```\n%s\n```", re.ReplaceAllString(s, ""))
}

func getQustionPath(fid string, id int, slug string) string {
	qid := int(util.ToInt64(fid))
	if qid == 0 {
		qid = id
	}
	tpl := "./src/%d-%s/"
	if id < 10000 {
		tpl = "./src/%04d-%s/"
	}
	return fmt.Sprintf(tpl, qid, slug)
}

func success(text string) {
	fmt.Println(fmt.Sprintf("\033[1;32mSuccess: %s\033[0m", text))
}

func notice(text string) {
	fmt.Println(fmt.Sprintf("\033[1;33mNotice: %s\033[0m", text))
}

func warning(text string) {
	fmt.Println(fmt.Sprintf("\033[1;31mError: %s\033[0m", text))
}

/**
 * ********************************
 * leetcode 数据
 * ********************************
 */

var (
	LCProbAllURL      = "https://leetcode-cn.com/api/problems/all/" // 问题列表地址
	LCGraphqlURL      = "https://leetcode-cn.com/graphql"           // 问题数据地址
	LCQuestionLinkURL = "https://leetcode-cn.com/problems/%s/"      // 问题页面地址
	LCTagLinkURL      = "https://leetcode-cn.com/tag/%s/"           // 标签页面地址
	LCQuestionSetURL  = "https://leetcode-cn.com/problemset/all/"   // 题库首页
	LCDifficulty      = []string{"", "简单", "中等", "困难"}              // 难度类型
)

type LeetCode struct {
	QuestionList    []LCQuestionInfo
	QuestionMap     map[string]LCQuestionInfo
	QuestionTagList []LCQuestionTagInfo
}

// 问题
type LCQuestionInfo struct {
	FQID           string `json:"fqid"`            // 前端 id
	QID            int    `json:"qid"`             // id
	Title          string `json:"title"`           // 名称
	Slug           string `json:"slug"`            // 标识
	Link           string `json:"link"`            //链接
	TotalAcs       int    `json:"total_acs"`       // 总通过次数
	TotalSubmitted int    `json:"total_submitted"` // 总提交次数
	Difficulty     string `json:"difficulty"`      // 困难度
}

// 问题标签
type LCQuestionTagInfo struct {
	Title string
	Link  string
	Count int
}

// 原始问题
type LCOriginQuestionInfo struct {
	Paid_only bool        `json:"paid_only"`
	Status    interface{} `json:"status"`
	Is_favor  bool        `json:"is_favor"`
	Progress  int64       `json:"progress"`
	Frequency int64       `json:"frequency"`
	Stat      struct {
		Question_id           int    `json:"question_id"`
		Question__hide        bool   `json:"question__hide"`
		Question__title       string `json:"question__title"`
		Question__title_slug  string `json:"question__title_slug"`
		Frontend_question_id  string `json:"frontend_question_id"`
		Total_acs             int    `json:"total_acs"`
		Total_submitted       int    `json:"total_submitted"`
		Total_column_articles int    `json:"total_column_articles"`
		Is_new_question       bool   `json:"is_new_question"`
	} `json:"stat"`
	Difficulty struct {
		Level int `json:"level"`
	} `json:"difficulty"`
}

// 问题详情
type LCQuestionDetail struct {
	Slug     string                       `json:"slug"`
	Title    string                       `json:"title"`
	Content  string                       `json:"content"`
	LangList []string                     `json:"lang_list"`
	TagList  []map[string]string          `json:"tag_list"`
	CodeMap  map[string]map[string]string `json:"code_list"`
}

// 获取问题列表
func (this *LeetCode) getQuestionList() error {
	req := request.NewRequest(new(http.Client))
	resp, err := req.Get(LCProbAllURL)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	respJson, err := resp.Json()
	if err != nil {
		return err
	}

	qList1, _ := respJson.Get("stat_status_pairs").Array()
	qList2 := []LCOriginQuestionInfo{}
	if err := json.Unmarshal([]byte(crypto.JsonEncode(qList1)), &qList2); err != nil {
		return err
	}

	questionMap := make(map[string]LCQuestionInfo)
	for i := len(qList2) - 1; i >= 0; i-- {
		questionInfo := LCQuestionInfo{
			FQID:           qList2[i].Stat.Frontend_question_id,
			QID:            qList2[i].Stat.Question_id,
			Title:          qList2[i].Stat.Question__title,
			Slug:           qList2[i].Stat.Question__title_slug,
			Link:           fmt.Sprintf(LCQuestionLinkURL, qList2[i].Stat.Question__title_slug),
			TotalAcs:       qList2[i].Stat.Total_acs,
			TotalSubmitted: qList2[i].Stat.Total_submitted,
			Difficulty:     LCDifficulty[qList2[i].Difficulty.Level],
		}
		this.QuestionList = append(this.QuestionList, questionInfo)
		questionMap[questionInfo.Slug] = questionInfo
	}
	this.QuestionMap = questionMap

	return nil
}

// 获取问题标签列表
func (this *LeetCode) getQuestionTagList() error {
	resp, err := http.Get(LCQuestionSetURL)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}

	doc.Find("#all-topic-tags a.tags-btn").Each(func(i int, s *goquery.Selection) {
		item := LCQuestionTagInfo{}
		item.Title = strings.TrimSpace(s.Find("span.text-gray").Text())
		item.Link, _ = s.Attr("href")
		item.Count = int(util.ToInt64(strings.TrimSpace(s.Find("span.badge").Text())))
		this.QuestionTagList = append(this.QuestionTagList, item)
	})

	return nil
}

// 获取问题详情
func (this *LeetCode) getQuestionDetail(slug string) (*LCQuestionDetail, error) {
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
	resp, err := req.Post(LCGraphqlURL)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	respJson, _ := resp.Json()

	questionDetail := &LCQuestionDetail{}

	// 基础信息
	questionDetail.Slug = slug
	questionDetail.Title, _ = respJson.Get("data").Get("question").Get("translatedTitle").String()
	questionDetail.Content, _ = respJson.Get("data").Get("question").Get("translatedContent").String()
	// 处理 html 标签
	questionDetail.Content = formatHtml(questionDetail.Content)

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

// leetcode 数据初始化
func (this *LeetCode) Init() error {
	err := this.getQuestionList()
	if err != nil {
		return err
	}
	err = this.getQuestionTagList()
	if err != nil {
		return err
	}
	return nil
}

/**
 * ********************************
 * 答题文件处理
 * ********************************
 */

// 默认答题配置
var (
	// 预设 golang 的答题模板
	langConf = map[string]map[string]string{
		"golang": {
			"file":        "solution.go",
			"fileTpl":     "package solution\n\n%s",
			"testfile":    "solution_test.go",
			"testfileTpl": tplQuestionGoTestFile,
		},
	}
	// 答题文件名称
	langFile = "solution.%s"
	// 答题文件模板
	langFileTpl = "%s"
	// 答题文件类型后缀
	langSuffix = map[string]string{
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
)

type LeetCodeFile struct {
}

// 生成答题相关文件
func (this *LeetCodeFile) GenerateQuestion(slug string, lang string, questionDetail *LCQuestionDetail) (err error) {
	questionInfo, _ := LC.QuestionMap[slug]
	questionPath := getQustionPath(questionInfo.FQID, questionInfo.QID, questionInfo.Slug)
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
		tagStr += fmt.Sprintf(" [%s](%s) ", tag["translatedName"], fmt.Sprintf(LCTagLinkURL, tag["slug"]))
	}
	if tagStr != "" {
		tagStr = fmt.Sprintf("> 分类: %s", tagStr)
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
		Difficulty: template.HTML(fmt.Sprintf("> 难度: %s", questionInfo.Difficulty)),
		Tags:       template.HTML(tagStr),
	}

	// 生成答题文件 README.md
	questionReadmeFile := fmt.Sprintf("%s/README.md", questionPath)
	var b bytes.Buffer
	tpl := template.Must(template.New("question").Parse(tplQuestionReadme))
	err = tpl.Execute(&b, questionReadmeInfo)
	if err != nil {
		return errors.New("答题 README 文件模板解析失败")
	}
	if err = util.WriteFile(questionReadmeFile, string(b.Bytes())); err != nil {
		return errors.New("创建答题 README 文件失败")
	}

	// 生成答题文件, 答题测试文件
	file, fileTpl, testfile, testfileTpl := "", "", "", ""
	if _, ok := langConf[lang]; ok {
		file, fileTpl = langConf[lang]["file"], langConf[lang]["fileTpl"]
		testfile, testfileTpl = langConf[lang]["testfile"], langConf[lang]["testfileTpl"]
	} else if _, ok := langSuffix[lang]; ok {
		file = fmt.Sprintf(langFile, langSuffix[lang])
		fileTpl = langFileTpl
	} else {
		// 未配置编程语言的, 生成完 README.md 直接结束
		return nil
	}

	// 创建答题文件
	questionFile := fmt.Sprintf("%s/%s", questionPath, file)
	if util.IsExist(questionFile) {
		notice("答题文件已存在: " + questionFile)
	} else {
		util.WriteFile(questionFile, fmt.Sprintf(fileTpl, questionDetail.CodeMap[lang]["code"]))
	}

	// 创建答题测试文件
	if testfile != "" {
		questionTestFile := fmt.Sprintf("%s/%s", questionPath, testfile)
		if util.IsExist(questionTestFile) {
			notice("答题测试文件已存在: " + questionTestFile)
		} else {
			util.WriteFile(questionTestFile, testfileTpl)
		}
	}

	return nil
}

// 生成 Resp README.md
func (this *LeetCodeFile) GenerateReadme() error {
	var b bytes.Buffer
	tmpl := template.Must(template.New("readme").Parse(tplReadme))
	err := tmpl.Execute(&b, this)
	if err != nil {
		return err
	}
	return util.WriteFile("./README.md", string(b.Bytes()))
}

// readme - 渲染标签列表
// [![数组](https://img.shields.io/badge/数组-99-red.svg)](https://shields.io/)
func (this *LeetCodeFile) DrawQuestionTagList() string {
	tagList := LC.QuestionTagList
	if len(tagList) <= 0 {
		return ""
	}

	resp := ""
	for _, tag := range tagList {
		tagLinks := strings.Split(tag.Link, "/")
		if len(tagLinks) < 4 {
			continue
		}
		link := fmt.Sprintf(LCTagLinkURL, tagLinks[2])
		title := strings.Replace(tag.Title, " ", "", -1)
		color := "lightgray"
		if tag.Count > 100 {
			color = "brightgreen"
		} else if tag.Count > 60 {
			color = "yellowgreen"
		} else if tag.Count > 40 {
			color = "orange"
		} else if tag.Count > 20 {
			color = "blue"
		} else if tag.Count > 10 {
			color = "red"
		}
		resp += fmt.Sprintf("[![%s](https://img.shields.io/badge/%s-%d-%s.svg?style=flat)](%s)\n", title, title, tag.Count, color, link)
	}
	return resp
}

// readme - 渲染题目表格, 只显示已解决的
func (this *LeetCodeFile) DrawQuestionList() string {
	questionList := LC.QuestionList
	if len(questionList) <= 0 {
		return ""
	}

	resp := fmt.Sprintln("|#|标题|难度|")
	resp += fmt.Sprintln("|:-:|:-|:-:|")
	for _, question := range questionList {
		questionPath := getQustionPath(question.FQID, question.QID, question.Slug)
		if !util.IsExist(questionPath) {
			continue
		}
		resp += fmt.Sprintf("|[%s](%s)|", question.FQID, question.Link)
		resp += fmt.Sprintf("[%s](%s)|", question.Title, questionPath)
		resp += fmt.Sprintf("%s|\n", question.Difficulty)
	}
	return resp
}

/**
 * ********************************
 * Main
 * ********************************
 */

var LC = new(LeetCode)
var LCFile = new(LeetCodeFile)

func main() {
	notice("拼命加载数据中...")
	err := LC.Init()
	if err != nil {
		warning("加载失败: " + err.Error())
		return
	}
	app := cli.NewApp()
	app.Name = "leetcli"
	app.Usage = "A CLI tool for leetcode"
	app.Version = "1.0.1"
	app.Flags = []cli.Flag{}
	app.Commands = []cli.Command{
		{
			Name:    "readme",
			Aliases: []string{},
			Usage:   "生成 README.md",
			Action: func(c *cli.Context) {
				err := LCFile.GenerateReadme()
				if err != nil {
					warning(err.Error())
					return
				}
				success("生成 README.md 成功")
				return
			},
		},
		{
			Name:    "question",
			Aliases: []string{},
			Usage:   "生成答题文件相关 [eg: question two-sum]",
			Action: func(c *cli.Context) {
				if len(c.Args()) <= 0 {
					warning("请输入问题标识")
					return
				}
				slug := c.Args()[0]
				if _, ok := LC.QuestionMap[slug]; !ok {
					warning("无效的问题标识")
					return
				}
				questionDetail, err := LC.getQuestionDetail(slug)
				if err != nil {
					warning("获取问题信息失败: " + err.Error())
					return
				}

				notice("支持的编程语言类型如下:")
				success(strings.Replace(strings.Trim(fmt.Sprint(questionDetail.LangList), "[]"), " ", ", ", -1))

				line := liner.NewLiner()
				defer line.Close()
				lang, _ := line.Prompt("请输入答题编程语言 > ")
				if lang == "" || !util.InArray(lang, questionDetail.LangList) {
					warning("无效的编程语言")
					return
				}
				notice("已选择的答题编辑语言: " + lang)
				err = LCFile.GenerateQuestion(slug, lang, questionDetail)
				if err != nil {
					warning("创建答题文件失败: " + err.Error())
					return
				}
				success("创建答题文件成功")

				return
			},
		},
		{
			Name:    "quit",
			Aliases: []string{},
			Usage:   "退出程序",
			Action: func(c *cli.Context) error {
				return cli.NewExitError("", 0)
			},
		},
	}
	commandList := []string{"help"}
	for _, command := range app.Commands {
		commandList = append(commandList, command.Name)
	}
	app.Action = func(c *cli.Context) error {
		line := liner.NewLiner()
		defer line.Close()
		line.SetCtrlCAborts(true)
		line.SetCompleter(func(line string) (c []string) {
			for _, n := range commandList {
				if strings.HasPrefix(n, strings.ToLower(line)) {
					c = append(c, n)
				}
			}
			return
		})
		notice("输入命令 `help` 试一试~")
		for {
			commandLine, err := line.Prompt(fmt.Sprintf("%s > ", app.Name))
			if err != nil {
				fmt.Println(err)
				return err
			}
			line.AppendHistory(commandLine)
			cmdArgs := strings.Split(commandLine, " ")
			if len(cmdArgs) == 0 {
				continue
			}
			if !util.InArray(cmdArgs[0], commandList) {
				warning("无效的命令, 输入命令 `help` 试一试~")
				continue
			}
			s := []string{os.Args[0]}
			s = append(s, cmdArgs...)
			c.App.Run(s)
		}
	}
	sort.Sort(cli.FlagsByName(app.Flags))
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
