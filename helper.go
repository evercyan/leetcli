package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/iikira/BaiduPCS-Go/pcsliner/args"
	"github.com/mozillazg/request"
	"github.com/peterh/liner"
	"github.com/urfave/cli"
)

var (
	lcDifficulty = []string{"", "简单", "中等", "困难"} // 难度类型
)

var (
	lcProbAllURL  = "https://leetcode-cn.com/api/problems/all/" // 问题列表地址
	lcGraphqlURL  = "https://leetcode-cn.com/graphql"           // 问题数据地址
	lcProbLinkURL = "https://leetcode-cn.com/problems/%s/"      // 问题页面地址
	lcTagLinkURL  = "https://leetcode-cn.com/tag/%s/"           // 标签页面地址
	lcProbSetURL  = "https://leetcode-cn.com/problemset/all/"   // 题库首页
)

var (
	tplReadme        = "./tpl/readme.tpl"          // readme 模板
	tplProblemReadme = "./tpl/problem_readme.tpl"  // 问题 readme 模板
	tplProblemGoTest = "./tpl/problem_go_test.tpl" // 问题 go 测试模板
)

// 只预设 golang 的答题模板
var langConf = map[string]map[string]string{
	"golang": {
		"file":        "solution.go",
		"fileTpl":     "package solution\n\n%s",
		"testfile":    "solution_test.go",
		"testfileTpl": read(tplProblemGoTest),
	},
}

// 默认答题配置
var (
	langFile    = "solution.%s"
	langFileTpl = "%s"
	langSuffix  = map[string]string{
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

func write(path, content string) error {
	err := ioutil.WriteFile(path, []byte(content), 0755)
	return err
}

func read(path string) string {
	file, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return ""
	}
	return string(data)
}

func string2int(v string) int {
	if r, err := strconv.Atoi(v); err == nil {
		return r
	}
	return 0
}

func jsonEncode(req interface{}) string {
	bytes, _ := json.Marshal(req)
	return string(bytes)
}

func pathExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func format(s string) string {
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
	s = re.ReplaceAllString(s, "")
	return fmt.Sprintf("```\n%s\n```", s)
}

func contains(val interface{}, target interface{}) bool {
	targetV := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetV.Len(); i++ {
			if targetV.Index(i).Interface() == val {
				return true
			}
		}
	case reflect.Map:
		if targetV.MapIndex(reflect.ValueOf(val)).IsValid() {
			return true
		}
	}
	return false
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

func color(text string, color int) string {
	return fmt.Sprintf("\033[1;%dm%s\033[0m", color, text)
}

/**
 * ******************************** leetcode
 */

type leetcode struct {
	ProblemInfo      map[int]lcProblem
	ProblemList      []lcProblem
	ProblemTitleInfo map[int]string
	TagList          []lcTagInfo
}

type lcResponse struct {
	Problems         []lcProblem    `json:"stat_status_pairs"`
	ProblemTitleInfo map[int]string `json:"problem_titles"`
}

type lcProblem struct {
	Status     string         `json:"status"`
	Stat       lcProblemStat  `json:"stat"`
	Difficulty lcProblemLevel `json:"difficulty"`
	Progress   int            `json:"progress"`
}
type lcProblemStat struct {
	ProbID         int    `json:"frontend_question_id"`
	ProbTitle      string `json:"question__title"`
	ProbTitleSlug  string `json:"question__title_slug"`
	IsNew          bool   `json:"is_new_question"`
	TotalAcs       int    `json:"total_acs"`
	TotalSubmitted int    `json:"total_submitted"`
}

type lcProblemLevel struct {
	Level int `json:"level"`
}

type lcProbInfo struct {
	ID         int                          `json:"id"`
	Title      string                       `json:"title"`
	TitleEn    string                       `json:"title_en"`
	Link       string                       `json:"link"`
	Content    string                       `json:"content"`
	ContentEn  string                       `json:"content_en"`
	Difficulty string                       `json:"difficulty"`
	Langs      []string                     `json:"langs"`
	Tags       []map[string]string          `json:"tags"`
	Codes      map[string]map[string]string `json:"codes"`
}

type lcTitleInfo struct {
	Title    string              `json:"title"`
	Question lcTitleQuestionInfo `json:"question"`
}

type lcTitleQuestionInfo struct {
	QuestionID string `json:"questionId"`
}

type lcTagInfo struct {
	Title   string
	TitleEn string
	Link    string
	Count   int
}

// 问题链接
func (lc *leetcode) getProblemURL(slugtitle string) string {
	return fmt.Sprintf(lcProbLinkURL, slugtitle)
}

// 问题标题
func (lc *leetcode) getProblemTitle(ProbID int, title string) string {
	if _, ok := lc.ProblemTitleInfo[ProbID]; !ok {
		return title
	}
	return lc.ProblemTitleInfo[ProbID]
}

// Github 链接
func (lc *leetcode) getGithubURL(ProbID int, slugtitle string) string {
	link := fmt.Sprintf("./src/%04d-%s/", ProbID, slugtitle)
	if !pathExist(link) {
		return ""
	}
	return link
}

// 标签链接
func (lc *leetcode) getTagURL(slugtitle string) string {
	return fmt.Sprintf(lcTagLinkURL, slugtitle)
}

// 获取问题列表
func (lc *leetcode) getProblemList() error {
	notice("拼命加载问题列表...")

	req := request.NewRequest(new(http.Client))
	// 获取问题列表
	req.Headers = map[string]string{
		"Accept-Encoding": "",
		"Referer":         "",
	}
	resp, err := req.Get(lcProbAllURL)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	result := &lcResponse{}
	if err := json.Unmarshal(body, result); err != nil {
		return err
	}

	// 获取问题中文标题
	req.Headers = map[string]string{
		"Content-Type": "application/json",
	}
	req.Json = map[string]interface{}{
		"operationName": "getQuestionTranslation",
		"query":         "query getQuestionTranslation($lang: String) {translations: allAppliedQuestionTranslations(lang: $lang) {title question {questionId __typename}\n    __typename}}",
		"variables":     map[string]string{},
	}
	pResp, err := req.Post(lcGraphqlURL)
	defer pResp.Body.Close()
	if err != nil {
		return err
	}
	titleList := []lcTitleInfo{}
	pJSON, _ := pResp.Json()
	translations, _ := pJSON.Get("data").Get("translations").Array()
	if err := json.Unmarshal([]byte(jsonEncode(translations)), &titleList); err != nil {
		return err
	}
	result.ProblemTitleInfo = map[int]string{}
	for _, item := range titleList {
		result.ProblemTitleInfo[string2int(item.Question.QuestionID)] = item.Title
	}

	// 问题列表
	lc.ProblemList = result.Problems
	// 标题信息
	lc.ProblemTitleInfo = result.ProblemTitleInfo
	// 保存问题状态信息
	lc.ProblemInfo = map[int]lcProblem{}
	for _, p := range lc.ProblemList {
		lc.ProblemInfo[p.Stat.ProbID] = p
	}
	// 保存标签列表
	lc.TagList = lc.getTagList()
	return nil
}

/**
 * 获取标签列表
 */
func (lc *leetcode) getTagList() []lcTagInfo {
	ret := []lcTagInfo{}
	resp, err := http.Get(lcProbSetURL)
	if err != nil {
		return ret
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return ret
	}
	doc.Find("#all-topic-tags a.tags-btn").Each(func(i int, s *goquery.Selection) {
		item := lcTagInfo{}
		item.Title = strings.TrimSpace(s.Find("span.text-gray").Text())
		item.TitleEn, _ = s.Attr("title")
		item.Link, _ = s.Attr("href")
		item.Count = string2int(strings.TrimSpace(s.Find("span.badge").Text()))
		ret = append(ret, item)
	})
	return ret
}

/**
 * 获取单个问题详情(包括描述, 标签, 语言支持, 代码片段等)
 */
func (lc *leetcode) getProblemInfo(ProbID int) (*lcProbInfo, error) {
	probInfo := &lcProbInfo{}
	p := lc.ProblemInfo[ProbID]
	req := request.NewRequest(new(http.Client))
	req.Headers = map[string]string{
		"Content-Type": "application/json",
	}
	req.Json = map[string]interface{}{
		"operationName": "questionData",
		"query":         "query questionData($titleSlug: String!) {question(titleSlug: $titleSlug) { questionId questionFrontendId boundTopicId title titleSlug content translatedTitle translatedContent isPaidOnly difficulty likes dislikes isLiked similarQuestions contributors {username profileUrl avatarUrl  __typename} langToValidPlayground topicTags { name slug translatedName __typename} companyTagStats codeSnippets { lang langSlug code __typename} stats hints solution {id canSeeDetail __typename } status sampleTestCase metaData judgerAvailable judgeType mysqlSchemas enableRunCode enableTestMode envInfo __typename}}",
		"variables": map[string]string{
			"titleSlug": p.Stat.ProbTitleSlug,
		},
	}
	pResp, err := req.Post(lcGraphqlURL)
	defer pResp.Body.Close()
	if err != nil {
		return nil, err
	}
	// 解析返回 json
	pJSON, _ := pResp.Json()
	probInfo.ID = ProbID
	probInfo.Link = lc.getProblemURL(p.Stat.ProbTitleSlug)
	probInfo.Title, _ = pJSON.Get("data").Get("question").Get("translatedTitle").String()
	probInfo.TitleEn, _ = pJSON.Get("data").Get("question").Get("title").String()
	probInfo.Content, _ = pJSON.Get("data").Get("question").Get("translatedContent").String()
	probInfo.ContentEn, _ = pJSON.Get("data").Get("question").Get("content").String()
	probInfo.Difficulty, _ = pJSON.Get("data").Get("question").Get("difficulty").String()
	// 分类
	tags := []map[string]string{}
	tagsArray, _ := pJSON.Get("data").Get("question").Get("topicTags").Array()
	if err := json.Unmarshal([]byte(jsonEncode(tagsArray)), &tags); err == nil {
		probInfo.Tags = tags
	}
	// 代码片段
	codesList := []map[string]string{}
	codesArray, _ := pJSON.Get("data").Get("question").Get("codeSnippets").Array()
	if err := json.Unmarshal([]byte(jsonEncode(codesArray)), &codesList); err != nil {
		return nil, err
	}
	// 语言代码
	codes := map[string]map[string]string{}
	// 支持语言类型
	langs := []string{}
	for _, item := range codesList {
		langs = append(langs, item["langSlug"])
		codes[item["langSlug"]] = item
	}
	probInfo.Langs = langs
	probInfo.Codes = codes
	return probInfo, nil
}

/**
 * ******************************** 文件相关
 */

/**
 * 生成 readme.md
 */
func (lc *leetcode) GenerateReadme() error {
	var b bytes.Buffer
	tmpl := template.Must(template.New("readme").Parse(read(tplReadme)))
	err := tmpl.Execute(&b, lc)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("./README.md", b.Bytes(), 0755)
	return err
}

/**
 * readme - 渲染题目表格
 */
func (lc *leetcode) DrawProblemList() string {
	if len(lc.ProblemList) <= 0 {
		return ""
	}
	res := fmt.Sprintln("|#|标题|难度|通过率|总提交次数|")
	res += fmt.Sprintln("|:-:|:-|:-: | :-: | :-: |")
	for _, p := range lc.ProblemList {
		res += fmt.Sprintf("|[%04d](%s)|", p.Stat.ProbID, lc.getProblemURL(p.Stat.ProbTitleSlug))

		probGithubURL := lc.getGithubURL(p.Stat.ProbID, p.Stat.ProbTitleSlug)
		if probGithubURL == "" {
			res += fmt.Sprintf("%s|", lc.getProblemTitle(p.Stat.ProbID, p.Stat.ProbTitle))
		} else {
			res += fmt.Sprintf("[%s](%s)|", lc.getProblemTitle(p.Stat.ProbID, p.Stat.ProbTitle), probGithubURL)
		}
		res += fmt.Sprintf("%s|", lcDifficulty[p.Difficulty.Level])
		totalPer := "0%"
		if p.Stat.TotalSubmitted > 0 {
			totalPer = fmt.Sprintf("%d%%", p.Stat.TotalAcs*100/p.Stat.TotalSubmitted)
		}
		res += fmt.Sprintf("%s|", totalPer)
		res += fmt.Sprintf("%d|\n", p.Stat.TotalSubmitted)
	}
	return res
}

/**
 * readme - 渲染类似题型
 */

func (lc *leetcode) DrawSimilarProblem() string {
	if len(lc.ProblemList) <= 0 {
		return ""
	}
	similarMap := make(map[string][]lcProblem)
	reg := regexp.MustCompile(`^([\S]+) II$`)
	for _, p := range lc.ProblemList {
		regRet := reg.FindStringSubmatch(lc.getProblemTitle(p.Stat.ProbID, p.Stat.ProbTitle))
		if len(regRet) > 1 {
			similarMap[regRet[1]] = []lcProblem{}
		}
	}
	for _, p := range lc.ProblemList {
		for key := range similarMap {
			if ok, _ := regexp.MatchString(key, lc.getProblemTitle(p.Stat.ProbID, p.Stat.ProbTitle)); ok {
				similarMap[key] = append(similarMap[key], p)
			}
		}
	}
	res := fmt.Sprintln("|#|标题|")
	res += fmt.Sprintln("|:-:|:-|")
	for _, problems := range similarMap {
		for _, p := range problems {
			res += fmt.Sprintf("|[%04d](%s)|", p.Stat.ProbID, lc.getProblemURL(p.Stat.ProbTitleSlug))
			probGithubURL := lc.getGithubURL(p.Stat.ProbID, p.Stat.ProbTitleSlug)
			if probGithubURL == "" {
				res += fmt.Sprintf("%s|\n", lc.getProblemTitle(p.Stat.ProbID, p.Stat.ProbTitle))
			} else {
				res += fmt.Sprintf("[%s](%s)|\n", lc.getProblemTitle(p.Stat.ProbID, p.Stat.ProbTitle), probGithubURL)
			}
		}
		res += fmt.Sprintln("||||\n||||")
	}
	return res
}

/**
 * readme - 渲染标签列表
 *
 * 徽章图标: https://shields.io/
 */
func (lc *leetcode) DrawTagList() string {
	res := ""
	if len(lc.TagList) <= 0 {
		return res
	}
	// [![数组](https://img.shields.io/badge/数组-99-red.svg)](https://shields.io/)
	for i := 0; i < len(lc.TagList); i++ {
		tag := lc.TagList[i]
		// 标签链接
		tagLinks := strings.Split(tag.Link, "/")
		if len(tagLinks) < 4 {
			continue
		}
		link := fmt.Sprintf(lcTagLinkURL, tagLinks[2])
		// 标题处理
		title := strings.Replace(tag.Title, " ", "", -1)
		// 图标颜色
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
		res += fmt.Sprintf("[![%s](https://img.shields.io/badge/%s-%d-%s.svg?style=flat)](%s)\n", title, title, tag.Count, color, link)
	}
	return res
}

/**
 * 终端 - 渲染单个问题
 */
func (lc *leetcode) DrawProblemLine(ProbID int) string {
	p := lc.ProblemInfo[ProbID]
	res := fmt.Sprintf("题号 : %d\n", p.Stat.ProbID)
	res += fmt.Sprintf("标题 : %s\n", lc.getProblemTitle(p.Stat.ProbID, p.Stat.ProbTitle))
	res += fmt.Sprintf("链接 : %s\n", lc.getProblemURL(p.Stat.ProbTitleSlug))
	res += fmt.Sprintf("难度 : %s\n", lcDifficulty[p.Difficulty.Level])
	res += fmt.Sprintf("比率 : %s", fmt.Sprintf("%d%%", p.Stat.TotalAcs*100/p.Stat.TotalSubmitted))
	res = color(res, 34)
	return res
}

/**
 * 生成问题信息
 */
func (lc *leetcode) GenerateProblem(ProbID int, lang string) (err error) {
	probInfo, err := lc.getProblemInfo(ProbID)
	if err != nil {
		return err
	}
	// 列表中的问题详情(包含状态等)
	p := lc.ProblemInfo[ProbID]
	probDir := fmt.Sprintf("./src/%04d-%s/", p.Stat.ProbID, p.Stat.ProbTitleSlug)
	if !pathExist(probDir) {
		dirErr := os.MkdirAll(probDir, 0755)
		if dirErr != nil {
			return errors.New("创建答题文件目录失败")
		}
	}
	// 生成问题标签
	tags := ""
	for _, tag := range probInfo.Tags {
		if tag["slug"] != "" {
			tags += fmt.Sprintf(" [%s](%s) ", tag["translatedName"], lc.getTagURL(tag["slug"]))
		}
	}
	if tags != "" {
		tags = fmt.Sprintf("> 分类: %s", tags)
	}
	difficulty := fmt.Sprintf("> 难度: %s", lcDifficulty[p.Difficulty.Level])
	// 内容需要格式化处理
	content := format(probInfo.Content)
	probTplInfo := struct {
		ID         int
		Title      interface{}
		Link       string
		Content    interface{}
		Difficulty interface{}
		Tags       interface{}
	}{
		ID:         probInfo.ID,
		Title:      template.HTML(probInfo.Title),
		Link:       probInfo.Link,
		Content:    template.HTML(content),
		Difficulty: template.HTML(difficulty),
		Tags:       template.HTML(tags),
	}
	// 生成文件 readme
	probPageFile := fmt.Sprintf("%s/README.md", probDir)
	var b bytes.Buffer
	tmpl := template.Must(template.New("question").Parse(read(tplProblemReadme)))
	err = tmpl.Execute(&b, probTplInfo)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(probPageFile, b.Bytes(), 0755)
	if err != nil {
		return err
	}
	file, fileTpl, testfile, testfileTpl := "", "", "", ""
	if _, ok := langConf[lang]; ok {
		// 如果有编辑语言配置
		file, fileTpl = langConf[lang]["file"], langConf[lang]["fileTpl"]
		testfile, testfileTpl = langConf[lang]["testfile"], langConf[lang]["testfileTpl"]
	} else if _, ok := langSuffix[lang]; ok {
		// 如果有编辑语言后缀配置
		file = fmt.Sprintf(langFile, langSuffix[lang])
		fileTpl = langFileTpl
	} else {
		return nil
	}

	// 创建答题文件
	probCodeFile := fmt.Sprintf("%s/%s", probDir, file)
	if pathExist(probCodeFile) {
		notice("文件已存在 - " + probCodeFile)
	} else {
		write(probCodeFile, fmt.Sprintf(fileTpl, probInfo.Codes[lang]["code"]))
	}

	// 创建测试文件
	if testfile != "" {
		probCodeTestFile := fmt.Sprintf("%s/%s", probDir, testfile)
		if pathExist(probCodeTestFile) {
			notice("测试文件已存在 - " + probCodeTestFile)
		} else {
			write(probCodeTestFile, testfileTpl)
		}
	}
	return nil
}

/**
 * ******************************** leetcode
 */

var lc = new(leetcode)

func main() {
	err := lc.getProblemList()
	if err != nil {
		warning("加载问题列表失败")
		return
	}
	app := cli.NewApp()
	app.Name = "leetcli"
	app.Usage = "leetcode cli"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{}
	app.Commands = []cli.Command{
		{
			Name:    "readme",
			Aliases: []string{},
			Usage:   "生成 README.md",
			Action: func(c *cli.Context) {
				err := lc.GenerateReadme()
				if err != nil {
					warning(err.Error())
					return
				}
				success("生成 README.md 成功")
				return
			},
		},
		{
			Name:    "problem",
			Aliases: []string{},
			Usage:   "生成答题文件相关 [eg: problem 101]",
			Action: func(c *cli.Context) {
				if len(c.Args()) <= 0 {
					warning("无效的问题 ID")
					return
				}
				ProbID := string2int(c.Args()[0])
				if _, ok := lc.ProblemInfo[ProbID]; !ok {
					warning("无效的问题 ID")
					return
				}
				probInfo, err := lc.getProblemInfo(ProbID)
				if err != nil {
					warning("获取问题信息失败 - " + err.Error())
					return
				}
				notice("问题详情如下")
				fmt.Println(lc.DrawProblemLine(ProbID))
				notice("支持的编程语言类型如下")
				fmt.Println(strings.Replace(strings.Trim(fmt.Sprint(probInfo.Langs), "[]"), " ", ", ", -1))

				line := liner.NewLiner()
				defer line.Close()
				lang, _ := line.Prompt("请输出答题编程语言 > ")
				if lang == "" || !contains(lang, probInfo.Langs) {
					warning("无效的编程语言")
					return
				}
				notice("已选择的编辑语言 - " + lang)
				err = lc.GenerateProblem(ProbID, lang)
				if err != nil {
					warning("创建答题文件失败 - " + err.Error())
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
			cmdArgs := args.Parse(commandLine)
			if len(cmdArgs) == 0 {
				continue
			}
			if !contains(cmdArgs[0], commandList) {
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
