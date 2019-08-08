package tool

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"testing"
)

var firstName = []string{
	"伟", "刚", "勇", "毅", "俊", "峰", "强", "军", "平", "保", "东", "文", "辉", "力", "明", "永", "健", "世", "广", "志", "义",
	"兴", "良", "海", "山", "仁", "波", "宁", "贵", "福", "生", "龙", "元", "全", "国", "胜", "学", "祥", "才", "发", "武", "新",
	"利", "清", "飞", "彬", "富", "顺", "信", "子", "杰", "涛", "昌", "成", "康", "星", "光", "天", "达", "安", "岩", "中", "茂",
	"进", "林", "有", "坚", "和", "彪", "博", "诚", "先", "敬", "震", "振", "壮", "会", "思", "群", "豪", "心", "邦", "承", "乐",
	"绍", "功", "松", "善", "厚", "庆", "磊", "民", "友", "裕", "河", "哲", "江", "超", "浩", "亮", "政", "谦", "亨", "奇", "固",
	"之", "轮", "翰", "朗", "伯", "宏", "言", "若", "鸣", "朋", "斌", "梁", "栋", "维", "启", "克", "伦", "翔", "旭", "鹏", "泽",
	"晨", "辰", "士", "以", "建", "家", "致", "树", "炎", "德", "行", "时", "泰", "盛", "雄", "琛", "钧", "冠", "策", "腾", "楠",
	"榕", "风", "航", "弘", "秀", "娟", "英", "华", "慧", "巧", "美", "娜", "静", "淑", "惠", "珠", "翠", "雅", "芝", "玉", "萍",
	"红", "娥", "玲", "芬", "芳", "燕", "彩", "春", "菊", "兰", "凤", "洁", "梅", "琳", "素", "云", "莲", "真", "环", "雪", "荣",
	"爱", "妹", "霞", "香", "月", "莺", "媛", "艳", "瑞", "凡", "佳", "嘉", "琼", "勤", "珍", "贞", "莉", "桂", "娣", "叶", "璧",
	"璐", "娅", "琦", "晶", "妍", "茜", "秋", "珊", "莎", "锦", "黛", "青", "倩", "婷", "姣", "婉", "娴", "瑾", "颖", "露", "瑶",
	"怡", "婵", "雁", "蓓", "纨", "仪", "荷", "丹", "蓉", "眉", "君", "琴", "蕊", "薇", "菁", "梦", "岚", "苑", "婕", "馨", "瑗",
	"琰", "韵", "融", "园", "艺", "咏", "卿", "聪", "澜", "纯", "毓", "悦", "昭", "冰", "爽", "琬", "茗", "羽", "希", "欣", "飘",
	"育", "滢", "馥", "筠", "柔", "竹", "霭", "凝", "晓", "欢", "霄", "枫", "芸", "菲", "寒", "伊", "亚", "宜", "可", "姬", "舒",
	"影", "荔", "枝", "丽", "阳", "妮", "宝", "贝", "初", "程", "梵", "罡", "恒", "鸿", "桦", "骅", "剑", "娇", "纪", "宽", "苛",
	"灵", "玛", "媚", "琪", "晴", "容", "睿", "烁", "堂", "唯", "威", "韦", "雯", "苇", "萱", "阅", "彦", "宇", "雨", "洋", "忠",
	"宗", "曼", "紫", "逸", "贤", "蝶", "菡", "绿", "蓝", "儿", "翠", "烟", "小", "轩"}















var firstNameLen = len(firstName)
func Test_GetFullName(t *testing.T)  {
	defer func(){
		if r:=recover();r!=nil{
			log.Println(r)
		}
	}()
	var times int =10
	var once int =10
	getChar :=func()string{
		index:=rand.Intn(firstNameLen)
		return firstName[index-1]
	}
	buf:=bytes.NewBufferString("")
	for t:=0;t<times;t++{
		for i:=0;i<once;i++{
			buf.WriteString(fmt.Sprintf("%6s",getChar()+getChar()))
		}
		//返回姓名
		log.Println(buf.String())
		buf.Reset()
	}
}

func Test_GetFullName2(t *testing.T)  {
	defer func(){
		if r:=recover();r!=nil{
			log.Println(r)
		}
	}()
	var times int =10
	var once int =10
	getChar :=func()string{
		index:=rand.Intn(10000)
		return fmt.Sprintf("%c",int16(19968+index))
	}
	buf:=bytes.NewBufferString("")
	for t:=0;t<times;t++{
		for i:=0;i<once;i++{
			buf.WriteString(fmt.Sprintf("%6s",getChar()+getChar()))
		}
		//返回姓名
		log.Println(buf.String())
		buf.Reset()
	}
}
var Column9 string=`项 城 政 挺 巷 荣 胡 南 标 枯 栋 相 查 柏 柳 柱
树 威 临 览 竖 省 盼 星 昨
贵 界 虹 思 品
哪 炭 峡 罚 贱 贴 骨 钞 钟 钢 钥 钩 卸 缸 拜 看 矩 怎 牲 选
适 秒 香 种 秋 科 重 复 竿 段 便 俩 贷 顺 修 保 促 侮 俭 俗
俘 信 皇 泉 鬼 侵 追 俊 盾 待 律 很 须 叙 剑 逃 食 盆 胆 胜
胞 胖 脉 勉 狭 狮 独 狡 狱 狠 贸 怨 急 饶 蚀 饺 饼 弯 将 奖
哀 亭 亮 度 迹 庭 疮 疯 疫 疤 姿 亲 音 帝 施 闻 阀 阁 差 养
美 姜 叛 送 类 迷 前 首 逆 总 炼 炸 炮 烂 剃 洁 洪 洒 浇 浊
洞 测 洗 活 派 洽 染 济 洋 洲 浑 浓 津 恒 恢 恰 恼 恨 举 觉
宣 室 宫 宪 突 穿 窃 客 冠 语 扁 袄 祖 神 祝 误 诱 说 诵 垦
退 既 屋 昼 费 陡 眉 孩 除 险 院 娃 姥 姨 姻 娇 怒 架 贺 盈
勇 怠 柔 垒 绑 绒 结 绕 骄 绘 给 络 骆 绝 绞 统`

var Column11 string =`球 理 捧 堵 描 域 掩 捷 排 掉 堆 推 掀 授 教 掏 掠 培 接 控
探 据 掘 职 基 著 勒 黄 萌 萝 菌 菜 萄 菊 萍 菠 营 械 梦 梢
梅 检 梳 梯 桶 救 副 票 戚 爽 聋 袭 盛 雪 辅 辆 虚 雀 堂 常
匙 晨 睁 眯 眼 悬 野 啦 晚 啄 距 跃 略 蛇 累 唱 患 唯 崖 崭
崇 圈 铜 铲 银 甜 梨 犁 移 笨 笼 笛 符 第 敏 做 袋 悠 偿 偶
偷 您 售 停 偏 假 得 衔 盘 船 斜 盒 鸽 悉 欲 彩 领 脚 脖 脸
脱 象 够 猜 猪 猎 猫 猛 馅 馆 凑 减 毫 麻 痒 痕 廊 康 庸 鹿
盗 章 竟 商 族 旋 望 率 着 盖 粘 粗 粒 断 剪 兽 清 添 淋 淹
渠 渐 混 渔 淘 液 淡 深 婆 梁 渗 情 惜 惭 悼 惧 惕 惊 惨 惯
寇 寄 宿 窑 密 谋 谎 祸 谜 逮 敢 屠 弹 随 蛋 隆 隐 婚 婶 颈
绩 绪 续 骑 绳 维 绵 绸 绿`

func Test_GetFullName_BIHUA(t *testing.T)  {
	defer func(){
		if r:=recover();r!=nil{
			log.Println(r)
		}
	}()
	var chars9 []string
	var chars11 []string
	 splitChar :=func(in string)[]string{
	 	array :=strings.Split(in," ")
	 	var rspArray []string
	 	if len(array)==0{
			return rspArray
		}
	 	for i:=range array{
	 		if len([]rune(array[i]))!=1{
	 			continue
			}
			rspArray=append(rspArray,array[i])
		}
	 	return rspArray
	 }
	chars9=splitChar(Column9)
	chars11=splitChar(Column11)
	var times int =10
	var once int =10
	getChar :=func(input []string)string{
		index:=rand.Intn(len(input)-1)
		return fmt.Sprintf("%s",input[index])
	}
	buf:=bytes.NewBufferString("")
	for t:=0;t<times;t++{
		for i:=0;i<once;i++{
			buf.WriteString(fmt.Sprintf("%6s",getChar(chars9)+getChar(chars9)))
		}
		//返回姓名
		log.Println(buf.String())
		buf.Reset()
	}
	fmt.Println(len(chars11))
}

func Test_Unicode(t *testing.T){
	//var b1V int16
	//hexb1,_ :=hex.DecodeString("4E00")
	//b1V =binary.BigEndian.Int16(hexb1)
	//log.Println(b1V)
	//for i:=0;i<100;i++{
	//	log.Println(fmt.Sprintf("%c",b1V+int16(i)))
	//}
}


//4E00-9FA5
