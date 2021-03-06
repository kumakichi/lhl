package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	//	if (is_someday()) {document.body.className = "someday"}
	fmt.Printf("\n    %s\n\n", getTodayString())

	pickTodaysLuck()
	printLuck()

	fmt.Printf("    座位朝向：面向%s写程序，BUG 最少。\n\n", directions[random(iday, 2)%len(directions)])
	fmt.Printf("    今日宜饮：%s\n\n", drinkStr(pickRandom(len(drinks), 2)))
	fmt.Printf("    女神亲近指数：%s\n\n", star(random(iday, 6)%5+1))
}

func getGoodTitle(isTitleLine bool) string {
	if isTitleLine {
		return "  ┃   宜  │"
	}

	return "  ┃       │"
}

func getBadTitle(isTitleLine bool) string {
	if isTitleLine {
		return "  ┃  不宜 │"
	}

	return "  ┃       │"
}

func printLuck() {
	max := calcDescriptionLength()

	goodTitleLine := len(goods) * 3 / 2
	fmt.Println("  ┏━━━━━━━┯━" + strings.Repeat("━", max) + "┓")
	titleLine := 0
	for _, v := range goods {
		titleLine++
		fmt.Printf("%s "+v.name+strings.Repeat(" ", max-calStr(v.name))+"┃\n", getGoodTitle(goodTitleLine == titleLine))
		titleLine++
		fmt.Printf("%s   "+v.good+strings.Repeat(" ", max-2-calStr(v.good))+"┃\n", getGoodTitle(goodTitleLine == titleLine))
		titleLine++
		fmt.Printf("%s "+strings.Repeat(" ", max)+"┃\n", getGoodTitle(goodTitleLine == titleLine))
	}
	fmt.Printf("  ┠───────┼─" + strings.Repeat("─", max) + "┨\n")

	badTitleLine := len(bads) * 3 / 2
	titleLine = 0
	for _, v := range bads {
		titleLine++
		fmt.Printf("%s "+v.name+strings.Repeat(" ", max-calStr(v.name))+"┃\n", getBadTitle(badTitleLine == titleLine))
		titleLine++
		fmt.Printf("%s   "+v.bad+strings.Repeat(" ", max-2-calStr(v.bad))+"┃\n", getBadTitle(badTitleLine == titleLine))
		titleLine++
		fmt.Printf("%s "+strings.Repeat(" ", max)+"┃\n", getBadTitle(badTitleLine == titleLine))
	}
	fmt.Printf("  ┗━━━━━━━┷━" + strings.Repeat("━", max) + "┛\n\n")
}

func drinkStr(indexes []int) string {
	var result []string
	for _, v := range indexes {
		result = append(result, drinks[v])
	}
	return strings.Join(result, "，")
}

func random(dayseed, indexseed int) int {
	var n = dayseed % 11117
	for i := 0; i < 100+indexseed; i++ {
		n = n * n
		n = n % 11117 // 11117 是个质数
	}
	return n
}

var today = time.Now()
var iday, _ = strconv.Atoi(today.Format("20060102"))

var weeks = []string{"日", "一", "二", "三", "四", "五", "六"}
var directions = []string{"北方", "东北方", "东方", "东南方", "南方", "西南方", "西方", "西北方"}

type activity struct {
	weekend         bool
	name, good, bad string
}

var activities = []activity{
	{name: "写单元测试", good: "写单元测试将减少出错", bad: "写单元测试会降低你的开发效率"},
	{name: "洗澡", good: "你几天没洗澡了？", bad: "会把设计方面的灵感洗掉", weekend: true},
	{name: "锻炼一下身体", good: "", bad: "能量没消耗多少，吃得却更多", weekend: true},
	{name: "抽烟", good: "抽烟有利于提神，增加思维敏捷", bad: "除非你活够了，死得早点没关系", weekend: true},
	{name: "白天上线", good: "今天白天上线是安全的", bad: "可能导致灾难性后果"},
	{name: "重构", good: "代码质量得到提高", bad: "你很有可能会陷入泥潭"},
	{name: "使用%t", good: "你看起来更有品位", bad: "别人会觉得你在装逼"},
	{name: "跳槽", good: "该放手时就放手", bad: "鉴于当前的经济形势，你的下一份工作未必比现在强"},
	{name: "招人", good: "你面前这位有成为牛人的潜质", bad: "这人会写程序吗？"},
	{name: "面试", good: "面试官今天心情很好", bad: "面试官不爽，会拿你出气"},
	{name: "提交辞职申请", good: "公司找到了一个比你更能干更便宜的家伙，巴不得你赶快滚蛋", bad: "鉴于当前的经济形势，你的下一份工作未必比现在强"},
	{name: "申请加薪", good: "老板今天心情很好", bad: "公司正在考虑裁员"},
	{name: "晚上加班", good: "晚上是程序员精神最好的时候", bad: "", weekend: true},
	{name: "在妹子面前吹牛", good: "改善你矮穷挫的形象", bad: "会被识破", weekend: true},
	{name: "撸管", good: "避免缓冲区溢出", bad: "强撸灰飞烟灭", weekend: true},
	{name: "浏览成人网站", good: "重拾对生活的信心", bad: "你会心神不宁", weekend: true},
	{name: "命名变量\"%v\"", good: "", bad: ""},
	{name: "写超过%l行的方法", good: "你的代码组织的很好，长一点没关系", bad: "你的代码将混乱不堪，你自己都看不懂"},
	{name: "提交代码", good: "遇到冲突的几率是最低的", bad: "你遇到的一大堆冲突会让你觉得自己是不是时间穿越了"},
	{name: "代码复审", good: "发现重要问题的几率大大增加", bad: "你什么问题都发现不了，白白浪费时间"},
	{name: "开会", good: "写代码之余放松一下打个盹，有益健康", bad: "小心被扣屎盆子背黑锅"},
	{name: "打DOTA", good: "你将有如神助", bad: "你会被虐的很惨", weekend: true},
	{name: "晚上上线", good: "晚上是程序员精神最好的时候", bad: "你白天已经筋疲力尽了"},
	{name: "修复BUG", good: "你今天对BUG的嗅觉大大提高", bad: "新产生的BUG将比修复的更多"},
	{name: "设计评审", good: "设计评审会议将变成头脑风暴", bad: "人人筋疲力尽，评审就这么过了"},
	{name: "需求评审", good: "", bad: ""},
	{name: "上微博", good: "今天发生的事不能错过", bad: "今天的微博充满负能量", weekend: true},
	{name: "上AB站", good: "还需要理由吗？", bad: "满屏兄贵亮瞎你的眼", weekend: true},
	{name: "玩FlappyBird", good: "今天破纪录的几率很高", bad: "除非你想玩到把手机砸了", weekend: true},
}

type special struct {
	date                     int
	typex, name, description string
}

var specials = []special{
	{date: 20140214, typex: "bad", name: "待在男（女）友身边", description: "脱团火葬场，入团保平安。"},
}

var tools = []string{"Eclipse写程序", "MSOffice写文档", "记事本写程序", "Windows8", "Linux", "MacOS", "IE", "Android设备", "iOS设备"}

var varNames = []string{"jieguo", "huodong", "pay", "expire", "zhangdan", "every", "free", "i1", "a", "virtual", "ad", "spider", "mima", "pass", "ui"}

var drinks = []string{"水", "茶", "红茶", "绿茶", "咖啡", "奶茶", "可乐", "鲜奶", "豆奶", "果汁", "果味汽水", "苏打水", "运动饮料", "酸奶", "酒"}

func is_someday() bool {
	return int(today.Month()) == 5 && today.Day() == 4
}

func getTodayString() string {
	return today.Format("今天是2006年1月2日 星期") + weeks[today.Weekday()]
}

func star(num int) string {
	var result string
	var i = 0
	for ; i < num; i++ {
		result += "★"
	}
	for ; i < 5; i++ {
		result += "☆"
	}
	return result
}

func pickTodaysLuck() {
	_activities := filter(activities)

	var numGood int = random(iday, 98)%3 + 2
	var numBad int = random(iday, 87)%3 + 2
	eventArr := pickRandomActivity(_activities, numGood+numBad)

	pickSpecials()

	for i := 0; i < numGood; i++ {
		addToGood(eventArr[i])
	}

	for i := 0; i < numBad; i++ {
		addToBad(eventArr[numGood+i])
	}
}

func filter(activities []activity) []activity {
	var result []activity

	if isWeekend() {
		for i := 0; i < len(activities); i++ {
			if activities[i].weekend {
				result = append(result, activities[i])
			}
		}
		return result
	}

	return activities
}

func isWeekend() bool {
	return today.Weekday() == 0 || today.Weekday() == 6
}

func pickSpecials() (good, bad int) {
	for i := 0; i < len(specials); i++ {
		s := specials[i]

		if iday == s.date {
			if s.typex == "good" {
				good++
				addToGood(activity{name: s.name, good: s.description})
			} else {
				bad++
				addToBad(activity{name: s.name, bad: s.description})
			}
		}
	}

	return
}

func pickRandomActivity(activities []activity, size int) []activity {
	var picked_events []activity
	indexes := pickRandom(len(activities), size)
	for _, v := range indexes {
		picked_events = append(picked_events, activities[v])
	}

	for i := 0; i < len(picked_events); i++ {
		picked_events[i] = parse(picked_events[i])
	}

	return picked_events
}

func pickRandom(length, size int) []int {
	var result []int

	for i := 0; i < length; i++ {
		result = append(result, i)
	}

	for j := 0; j < length-size; j++ {
		var index = random(iday, j) % len(result)
		result = splice(result, index, 1)
	}

	return result
}

func parse(event activity) activity {
	result := activity{name: event.name, good: event.good, bad: event.bad} // clone

	if strings.Index(result.name, "%v") != -1 {
		result.name = strings.Replace(result.name, "%v", varNames[random(iday, 12)%len(varNames)], 1)
	}

	if strings.Index(result.name, "%t") != -1 {
		result.name = strings.Replace(result.name, "%t", tools[random(iday, 11)%len(tools)], 1)
	}

	if strings.Index(result.name, "%l") != -1 {
		result.name = strings.Replace(result.name, "%l", fmt.Sprintf("%d", random(iday, 12)%247+30), 1)
	}

	return result
}

func addToGood(event activity) {
	goods = append(goods, event)
	//fmt.Println("=======>good:", event.name, event.good)
}

func addToBad(event activity) {
	bads = append(bads, event)
	//fmt.Println("=======>bad:", event.name, event.bad)
}

func splice(a []int, index, howmany int) []int {
	return append(a[:index], a[index+howmany:]...)
}

var (
	goods []activity
	bads  []activity
)

func calcDescriptionLength() int {
	var max int
	for _, v := range goods {
		l := calStr(v.good)
		if l > max {
			max = l
		}
	}
	for _, v := range bads {
		l := calStr(v.bad)
		if l > max {
			max = l
		}
	}

	return max + 2
}

func calStr(s string) int {
	var cnt int
	for _, v := range s {
		if v < 127 {
			cnt++
		} else {
			cnt += 2
		}
	}
	return cnt
}
