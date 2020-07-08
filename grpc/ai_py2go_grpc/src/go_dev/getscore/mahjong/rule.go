package mahjong

import (
	"fmt"
	"strconv"
)

var DEBUG = false
var ErRenRuleCheckerSington = &BaseRuleChecker{
	HasWindTile:  false,
	HasZFBTile:   false,
	HasSevenPair: true,
}

const (
	MAHJONG_SEVENPAIRS int = 1000
	MAHJONG_FOURHUN        = 1001
)

const (
	ResultFlag_Sequence int64 = iota //123
	ResultFlag_Tile                  //111
	ResultFlag_WindTile              //风扑
	ResultFlag_ZFBTile               //中发白
	ResultFlag_19Tile                //119
	ResultFlag_28Tile                //228
)

const (
	XYWORDS_DASIXI          = iota // 大四喜
	XYWORDS_DASANYUAN              // 大三元
	XYWORDS_JIULIANBAODENG         // 九莲宝灯
	XYWORDS_FORKONG                // 四杠
	XYWORDS_LIANQIDUI              // 连七对
	XYWORDS_XIAOSIXI               // 小四喜
	XYWORDS_XIAOSANYUAN            // 小三元
	XYWORDS_WORDONLYACOLOR         // 字一色
	XYWORDS_SIANKE                 // 四暗刻
	XYWORDS_YISESHUANGLONG         // 一色双龙会
	XYWORDS_YISESITONGSHUN         // 一色四同顺
	XYWORDS_YISESIJIEGAO           // 一色四节高
	XYWORDS_TIANHU                 // 天湖
	XYWORDS_YISESIBUGAO            // 一色四步高
	XYWORDS_SANGANG                // 三杠
	XYWORDS_HUNYAOJIU              // 混幺九
	XYWORDS_QINGYISE               // 清一色
	XYWORDS_YISESANTONGSHUN        // 一色三同顺
	XYWORDS_YISESANJIEGAO          // 一色三节高
	XYWORDS_QIDUI                  // 七对
	XYWORDS_DIHU                   // 地胡
	XYWORDS_QINGLONG               // 青龙
	XYWORDS_YISESANBUGAO           // 一色三步高
	XYWORDS_SANANKE                // 三暗刻
	XYWORDS_RENHU                  // 人胡
	XYWORDS_DAYUWU                 // 大于五
	XYWORDS_XIAOYUWU               // 小于五
	XYWORDS_SANFENGKE              // 三风刻
	XYWORDS_MIAOSHOUHUICHUN        // 妙手回春
	XYWORDS_HAISILAOYUE            // 海底捞月
	XYWORDS_GANGSHANGKAIHUA        // 杠上开花
	XYWORDS_QIANGGANGHU            // 抢杠胡
	XYWORDS_PENGPENGHU             // 碰碰胡
	XYWORDS_HUNYISE                // 混一色
	XYWORDS_SHUANGJIANKE           // 双箭刻
	XYWORDS_SHUANGANGANG           // 双暗杠
	XYWORDS_QUANDAIYAO             // 全带幺
	XYWORDS_BUQIUREN               // 不求人
	XYWORDS_SHUANGMINGGANG         // 双明杠
	XYWORDS_HUJUEZHANG             // 胡绝张
	XYWORDS_JIANKE                 // 箭刻
	XYWORDS_QUANFENGKE             // 圈风刻
	XYWORDS_MENFENGKE              // 门风刻
	XYWORDS_MENQIANQING            // 门前清
	XYWORDS_PINGHU                 // 平胡
	XYWORDS_ZIMO                   // 自摸
	XYWORDS_SIGUIYI                // 四归一
	XYWORDS_SHUANGANKE             // 双暗刻
	XYWORDS_DUANYAO                // 断幺
	XYWORDS_ANGANG                 // 暗杠
	XYWORDS_ERWUBAJIANG            // 258将
	XYWORDS_YIBANGAO               // 一般高
	XYWORDS_LIANLIU                // 连六
	XYWORDS_LAOSHAOFU              // 老少副
	XYWORDS_YAOJIUKE               // 幺九刻
	XYWORDS_MINGANG                // 明杠
	XYWORDS_WUZI                   // 无字
	XYWORDS_LIANZHANG              // 边张
	XYWORDS_KANZHANG               // 坎张
	XYWORDS_DANDIAOJIANG           // 单调将
	XYWORDS_YAOJIUTOU              // 幺九头
)

func getSocreByHuType(huType int) int {
	switch huType {
	case XYWORDS_DASIXI, XYWORDS_DASANYUAN, XYWORDS_JIULIANBAODENG, XYWORDS_FORKONG, XYWORDS_LIANQIDUI:
		return 88
	case XYWORDS_XIAOSIXI, XYWORDS_XIAOSANYUAN, XYWORDS_WORDONLYACOLOR, XYWORDS_SIANKE, XYWORDS_YISESHUANGLONG:
		return 64
	case XYWORDS_YISESITONGSHUN, XYWORDS_YISESIJIEGAO, XYWORDS_TIANHU:
		return 48
	case XYWORDS_YISESIBUGAO, XYWORDS_SANGANG, XYWORDS_HUNYAOJIU:
		return 32
	case XYWORDS_QINGYISE, XYWORDS_YISESANTONGSHUN, XYWORDS_YISESANJIEGAO, XYWORDS_QIDUI, XYWORDS_DIHU:
		return 24
	case XYWORDS_QINGLONG, XYWORDS_YISESANBUGAO, XYWORDS_SANANKE, XYWORDS_RENHU:
		return 16
	case XYWORDS_DAYUWU, XYWORDS_XIAOYUWU, XYWORDS_SANFENGKE:
		return 12
	case XYWORDS_MIAOSHOUHUICHUN, XYWORDS_HAISILAOYUE, XYWORDS_GANGSHANGKAIHUA, XYWORDS_QIANGGANGHU:
		return 8
	case XYWORDS_PENGPENGHU, XYWORDS_HUNYISE, XYWORDS_SHUANGJIANKE, XYWORDS_SHUANGANGANG:
		return 6
	case XYWORDS_QUANDAIYAO, XYWORDS_BUQIUREN, XYWORDS_SHUANGMINGGANG, XYWORDS_HUJUEZHANG:
		return 4
	case XYWORDS_JIANKE, XYWORDS_QUANFENGKE, XYWORDS_MENFENGKE, XYWORDS_MENQIANQING, XYWORDS_PINGHU,
		XYWORDS_SIGUIYI, XYWORDS_SHUANGANKE, XYWORDS_DUANYAO, XYWORDS_ANGANG, XYWORDS_ERWUBAJIANG, XYWORDS_ZIMO:
		return 2
	case XYWORDS_YIBANGAO, XYWORDS_LIANLIU, XYWORDS_LAOSHAOFU, XYWORDS_YAOJIUKE, XYWORDS_MINGANG, XYWORDS_WUZI,
		XYWORDS_LIANZHANG, XYWORDS_KANZHANG, XYWORDS_DANDIAOJIANG, XYWORDS_YAOJIUTOU:
		return 1
	default:
		return -1
	}
}

//规则检测
type RuleChecker interface {
	//能不能碰
	CanPongs([]int64, int64) bool
	//能不能杠
	CanKong([]int64, int64) bool
	//当前手牌中是否有杠
	HasKong([]int64) bool
	//能不能吃
	CanChow([]int64, int64, int64) bool
	//是否有将
	HasEyes([]int64, int64) bool
	//能不能胡,胡的牌型
	CanHu([]int64, int) (bool, []HuHandResult)
	//
	GetTingCards(pool []int64, cnt int, maxCnts int, excludeCard int64) (cards []int64)
	CheckHuType(typeMap map[int64]int, tilesNum int, discards, handCards, surplusCards, otherHands, chowCards, pongCards, exposedKong, concealedKong []int64, isBanker, isZimo, robKong, lastOptIsKong, lastCard_cpk bool, lastCard, tableFeng, menFeng int64) map[int64]int
}

type HuHandResult struct {
	Sequence    [4][3]int64
	Eyes        [2]int64 //将
	EyesKind    int64    //将
	Flags       [4]int64 //每扑牌的类型
	SpecailKind int      //特殊牌型
	SeqCnt      int
}

func (hhr *HuHandResult) Pop() {
	if hhr.SeqCnt > 0 {
		hhr.SeqCnt--
		for i := int64(0); i < 3; i++ {
			hhr.Sequence[hhr.SeqCnt][i] = 0
		}
		hhr.Flags[hhr.SeqCnt] = 0

	}
}

func (hhr *HuHandResult) AddSequence(c int64) {
	if hhr.SeqCnt < len(hhr.Sequence) {
		for i := int64(0); i < 3; i++ {
			hhr.Sequence[hhr.SeqCnt][i] = c + i
		}
		hhr.Flags[hhr.SeqCnt] = ResultFlag_Sequence
		hhr.SeqCnt++
	}
}

func (hhr *HuHandResult) AddCustom(v1, v2, v3 int64, flag int64) {
	hhr.Sequence[hhr.SeqCnt][0] = v1
	hhr.Sequence[hhr.SeqCnt][1] = v2
	hhr.Sequence[hhr.SeqCnt][2] = v3
	hhr.Flags[hhr.SeqCnt] = flag
	hhr.SeqCnt++
}

func (hhr *HuHandResult) AddTiles(from int64) {
	if hhr.SeqCnt < len(hhr.Sequence) {
		for i := 0; i < 3; i++ {
			hhr.Sequence[hhr.SeqCnt][i] = from
		}
		hhr.Flags[hhr.SeqCnt] = ResultFlag_Tile
		hhr.SeqCnt++
	}
}

func (hhr *HuHandResult) Clear() {
	hhr.SeqCnt = 0
	hhr.EyesKind = -1
}

type BaseRuleChecker struct {
	HasWindTile  bool //风扑
	HasZFBTile   bool //将扑
	HasSevenPair bool //七对
	logic        MjAlgorithm
}

//能不能碰
func (brc *BaseRuleChecker) CanPongs(pool []int64, c int64) bool {
	if c < 0 || c >= MJ_MAX {
		return false
	}

	return pool[c] >= 2
}

//能不能杠
func (brc *BaseRuleChecker) CanKong(pool []int64, c int64) bool {
	if c < 0 || c >= MJ_MAX {
		return false
	}

	return pool[c] == 3
}

//当前手牌中是否有杠
func (brc *BaseRuleChecker) HasKong(pool []int64) bool {
	n := len(pool)
	for i := 0; i < n; i++ {
		if pool[i] == 4 {
			return true
		}
	}
	return false
}

//是否有将
func (brc *BaseRuleChecker) HasEyes(pool []int64, except int64) bool {
	for i := int64(0); i < MJ_MAX; i++ {
		if i != except && pool[i] >= 2 {
			return true
		}
	}
	return false
}

//能否吃
func (brc *BaseRuleChecker) CanChow(pool []int64, c, exclude int64) bool {
	if c < 0 || c > MJ_MAX {
		return false
	}
	if c >= MJ_CHARACTERS_1 && c <= MJ_CHARACTERS_9 { //万
		return brc.canSuitThreeInSequence(pool, c, exclude, MJ_CHARACTERS_1, MJ_CHARACTERS_9)
	}
	return false
}

func (brc *BaseRuleChecker) canSuitThreeInSequence(pool []int64, c, exclude, beg, end int64) bool {
	if c == beg { //右
		if (c+1) != exclude && (c+2) != exclude && pool[c+1] > 0 && pool[c+2] > 0 {
			return true
		}
	} else if c == end { //左
		if (c-1) != exclude && (c-2) != exclude && pool[c-1] > 0 && pool[c-2] > 0 {
			return true
		}
	} else if c-1 == beg { //中右
		if (c-1) != exclude && (c+1) != exclude && pool[c-1] > 0 && pool[c+1] > 0 {
			return true
		} else if (c+1) != exclude && (c+2) != exclude && pool[c+1] > 0 && pool[c+2] > 0 {
			return true
		}
	} else if c+1 == end { //中左
		if (c-1) != exclude && (c+1) != exclude && pool[c-1] > 0 && pool[c+1] > 0 {
			return true
		} else if (c-1) != exclude && (c-2) != exclude && pool[c-1] > 0 && pool[c-2] > 0 {
			return true
		}
	} else { //左中右
		if (c+1) != exclude && (c+2) != exclude && pool[c+1] > 0 && pool[c+2] > 0 {
			return true
		} else if (c-1) != exclude && (c-2) != exclude && pool[c-1] > 0 && pool[c-2] > 0 {
			return true
		} else if (c-1) != exclude && (c+1) != exclude && pool[c-1] > 0 && pool[c+1] > 0 {
			return true
		}
	}
	return false
}

func (brc *BaseRuleChecker) canSuitThreeInWind(pool []int64, c, exclude int64) bool {
	if brc.HasWindTile && c >= MJ_WIND_EAST && c <= MJ_WIND_NORTH { //风扑
		cnt := 0
		for i := MJ_WIND_EAST; i <= MJ_WIND_NORTH; i++ {
			if i != exclude && i != c {
				if pool[i] > 0 {
					cnt++
				}
			}
		}
		return cnt >= 2
	}
	return false
}

func (brc *BaseRuleChecker) canSuitThreeInZFB(pool []int64, c, exclude int64) bool {
	if brc.HasZFBTile && c >= MJ_DRAGON_RED && c <= MJ_DRAGON_WHITE { //将扑
		cnt := 0
		for i := MJ_DRAGON_RED; i <= MJ_DRAGON_WHITE; i++ {
			if i != exclude && i != c {
				if pool[i] > 0 {
					cnt++
				}
			}
		}
		return cnt >= 2
	}
	return false
}

//能不能胡,胡的牌型
func (brc *BaseRuleChecker) CanHu(pool []int64, cnt int) (hu bool, results []HuHandResult) {
	//先检查牌数量是否是3n+2
	if (cnt-2)%3 != 0 {
		return
	}

	if brc.HasSevenPair && cnt == 14 && brc.IsSevenPairs(pool) {
		results = append(results, HuHandResult{
			SpecailKind: MAHJONG_SEVENPAIRS,
		})
		return true, results
	}

	//把检测的牌也加进手牌
	var result HuHandResult
	var eyesIdx int64
	var temp [MJ_MAX]int64

	for {
		result.Clear()
		copy(temp[:], pool)
		for i := eyesIdx; i < MJ_MAX; i++ {
			if temp[i] >= 2 {
				result.EyesKind = i
				result.Eyes[0], result.Eyes[1] = i, i
				temp[i] -= 2
				eyesIdx = i + 1
				break
			}
		}
		if result.EyesKind == -1 {
			return len(results) != 0, results
		}
		if brc.tryPruneOnce(temp[:], &result) {
			results = append(results, result)
			return len(results) != 0, results
		}
	}

	return len(results) != 0, results
}

func (brc *BaseRuleChecker) GetTingCards(pool []int64, cnt int, maxCnts int, excludeCard int64) (cards []int64) {
	var temp [MJ_MAX]int64
	for i := MJ_DOTS_1; i < MJ_MAX; i++ {
		if i != excludeCard {
			copy(temp[:], pool)
			temp[i]++
			if hu, _ := brc.CanHu(temp[:], cnt+1); hu {
				cards = append(cards, i)
				if maxCnts > 0 && len(cards) == maxCnts {
					return
				}
			}
		}
	}
	return
}

func (brc *BaseRuleChecker) tryPruneOnce(pool []int64, result *HuHandResult) bool {
	for i := int64(0); i < MJ_MAX; i++ {
		switch pool[i] {
		case 1:
			if !brc.trySuitOneSquence(pool, result, i) {
				return false
			}
		case 2:
			for j := 0; j < 2; j++ {
				if !brc.trySuitOneSquence(pool, result, i) {
					return false
				}
			}
		case 3:
			pool[i] -= 3
			result.AddTiles(i)
		case 4:
			pool[i] -= 3
			result.AddTiles(i)
			if !brc.trySuitOneSquence(pool, result, i) {
				return false
			}
		}
	}
	for i := int64(0); i < MJ_MAX; i++ {
		if pool[i] != 0 {
			return false
		}
	}
	return true
}

func (brc *BaseRuleChecker) trySuitOneSquence(pool []int64, result *HuHandResult, i int64) bool {
	if i >= MJ_DOTS_1 && i <= MJ_DOTS_9 {
		if !brc.trySuitSequence(pool, result, i, MJ_DOTS_1, MJ_DOTS_9) {
			return false
		}
	} else if i >= MJ_BAMBOO_1 && i <= MJ_BAMBOO_9 {
		if !brc.trySuitSequence(pool, result, i, MJ_BAMBOO_1, MJ_BAMBOO_9) {
			return false
		}
	} else if i >= MJ_CHARACTERS_1 && i <= MJ_CHARACTERS_9 {
		if !brc.trySuitSequence(pool, result, i, MJ_CHARACTERS_1, MJ_CHARACTERS_9) {
			return false
		}
	} else if i >= MJ_WIND_EAST && i <= MJ_WIND_NORTH { //风扑
		if !brc.HasWindTile || !brc.trySuitWindSequence(pool, result, i, MJ_WIND_EAST, MJ_WIND_NORTH, ResultFlag_WindTile) {
			return false
		}
	} else if i >= MJ_DRAGON_RED && i <= MJ_DRAGON_WHITE { //将扑
		if !brc.HasZFBTile || !brc.trySuitWindSequence(pool, result, i, MJ_DRAGON_RED, MJ_DRAGON_WHITE, ResultFlag_ZFBTile) {
			return false
		}
	}
	return true
}

func (brc *BaseRuleChecker) trySuitSequence(pool []int64, result *HuHandResult, c, min, max int64) bool {
	if c+2 > max {
		return false
	}
	if pool[c+1] > 0 && pool[c+2] > 0 {
		for i := int64(0); i < 3; i++ {
			pool[c+i] -= 1
		}
		result.AddSequence(c)
		return true
	}
	return false
}

func (brc *BaseRuleChecker) trySuitWindSequence(pool []int64, result *HuHandResult, c, min, max, flag int64) bool {
	if c+2 > max {
		return false
	}

	var cards []int64
	for i := c; i <= max; i++ {
		if pool[i] > 0 {
			cards = append(cards, i)
		}
	}
	cnt := len(cards)
	if cnt < 3 {
		return false
	}

	for i := 0; i < cnt && i < 3; i++ {
		pool[cards[i]] -= 1
	}
	result.AddCustom(cards[0], cards[1], cards[2], flag)
	return true
}

func (brc *BaseRuleChecker) IsSevenPairs(pool []int64) bool {
	n := len(pool)
	for i := 0; i < n; i++ {
		if pool[i] == 0 {
			continue
		}
		if pool[i] != 2 && pool[i] != 4 {
			return false
		}
	}
	return true
}

//豪华七对
func (brc *BaseRuleChecker) IsHHSevenPairs(pool []int64) bool {
	if brc.IsSevenPairs(pool) {
		n := len(pool)
		for i := 0; i < n; i++ {
			if pool[i] == 0 {
				continue
			}
			if pool[i] == 4 {
				return true
			}
		}
	}
	return false
}

func (brc *BaseRuleChecker) GetWindCnt(pool []int64) int {
	cnt := 0
	for i := MJ_WIND_EAST; i <= MJ_WIND_NORTH; i++ {
		cnt += int(pool[i])
	}
	return cnt

}

func (brc *BaseRuleChecker) GetZFBCnt(pool []int64) int {
	cnt := 0
	for i := MJ_DRAGON_RED; i <= MJ_DRAGON_WHITE; i++ {
		cnt += int(pool[i])
	}
	return cnt
}

func in(sl []int64, v int64) bool {
	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

func log(a ...interface{}) {
	if DEBUG {
		fmt.Println(a...)
	}
}

/**
typeMap 手里的牌，不包含吃碰杠
tilesNum 吃碰杠的个数
discards 已经出过的牌，俩人的都算
handCards 手里的牌 排过续的
surplusCards 牌桌上剩余的牌
otherHands:对家手里的牌
chowCards: 自己手里的吃的牌 只记录第一位
isBanker：是否是庄家
isZimo：是否是自摸
lastOptIsKong：上一次操作是否是杠
anGangNum：暗杠数量
mingGangNum：明杠数量
otherHandDiscardNum：对方手里的某张牌的数量，胡绝张用

lastCard：最后一张牌
tableFeng：本局是圈风
menFeng：我的门风
*/
func (brc *BaseRuleChecker) CheckHuType(typeMapAll map[int64]int, tilesNum int, discards, handCards, surplusCards, otherHands, chowCards, pongCards, mingKong, anKong []int64, isBanker, isZimo, robKong, lastOptIsKong, lastCard_cpk bool, lastCard, tableFeng, menFeng int64) map[int64]int {
	//logger.Logger.Tracef("typeMapAll:%v,tilesNum:%v,discards:%v,handCards:%v,surplusCards:%v,otherHands:%v,chowCards:%v,pongCards:%v,mingKong:%v,anKong:%v,isBanker:%v,isZimo:%v,lastOptIsKong:%v,lastCard_cpk:%v,lastCard:%v,tableFeng:%v,menFeng:%v", printMap(typeMapAll), tilesNum, discards, handCards, surplusCards, otherHands, chowCards, pongCards, mingKong, anKong, isBanker, isZimo, lastOptIsKong, lastCard_cpk, lastCard, tableFeng, menFeng)
	pongkongCards := make([]int64, 0)
	pongkongCards = append(pongkongCards, mingKong...)
	pongkongCards = append(pongkongCards, anKong...)
	pongkongCards = append(pongkongCards, pongCards...)
	pongkongCards = sortCards(pongkongCards)
	sortCards(handCards)

	var config MJSpecailConfig
	config.HuQiDui = 1
	config.FengPu = false
	config.JiangPu = false
	results := brc.logic.Hu(handCards, nil, -1, []int64{-1}, &config)
	checkTypeScore := make(map[int64]int, 0)

	resultLength := len(results.Results)
	if resultLength > 0 {

		for i := 0; i < resultLength; i++ {
			result := results.Results[i]
			typeAndScore := make(map[int64]int)
			if results.qidui && len(handCards) == 14 && i == 0 {
				isLianQiDui := checkQiDui(handCards)
				cardsMap := splicChangeMap(handCards)
				if isLianQiDui {
					checkTypeScore[XYWORDS_LIANQIDUI] = getSocreByHuType(XYWORDS_LIANQIDUI)
				} else {
					checkTypeScore[XYWORDS_QIDUI] = getSocreByHuType(XYWORDS_QIDUI)
				}
				count, _ := getCountByPrope(cardsMap, MJ_CHARACTERS_2, MJ_CHARACTERS_8, 1)
				if count == 0 {
					checkTypeScore[XYWORDS_QUANDAIYAO] = getSocreByHuType(XYWORDS_QUANDAIYAO)
				}
				if cardsMap[MJ_CHARACTERS_2] >= 2 || cardsMap[MJ_CHARACTERS_5] >= 2 || cardsMap[MJ_CHARACTERS_8] >= 2 {
					checkTypeScore[XYWORDS_ERWUBAJIANG] = getSocreByHuType(XYWORDS_ERWUBAJIANG)
				}
				if cardsMap[MJ_CHARACTERS_1] >= 2 || cardsMap[MJ_CHARACTERS_9] >= 2 {
					checkTypeScore[XYWORDS_YAOJIUTOU] = getSocreByHuType(XYWORDS_YAOJIUTOU)
				}
				brc.chekOtherHuType(typeMapAll, tilesNum, discards, handCards, surplusCards, otherHands, mingKong, anKong, chowCards, pongCards, isBanker, isZimo, robKong, lastOptIsKong, lastCard_cpk, lastCard, results.Results[0].EyesKind, checkTypeScore)
				//delete(checkTypeScore, XYWORDS_QUANDAIYAO)

				brc.deleteRepeat(checkTypeScore)

			} else {
				if result.EyesKind == 0 {
					continue
				}
				brc.ordinaryHu(typeMapAll, tilesNum, discards, handCards, surplusCards, otherHands, chowCards, pongCards, mingKong, anKong, pongkongCards, isBanker, isZimo, robKong, lastOptIsKong, lastCard_cpk, lastCard, tableFeng, menFeng, result, typeAndScore)
				brc.deleteRepeat(typeAndScore)
				if len(checkTypeScore) == 0 {
					checkTypeScore = typeAndScore
				} else {
					checkScore := getTotalScoreByMap(checkTypeScore)
					typeScore := getTotalScoreByMap(typeAndScore)
					if typeScore > checkScore {
						checkTypeScore = typeAndScore
					}
				}
			}
		}
	}
	//logger.Logger.Tracef("twomahjong_result:", appendString(checkTypeScore))
	return checkTypeScore
}

// 番型中不牵扯牌型的算法
func (brc *BaseRuleChecker) chekOtherHuType(typeMapAll map[int64]int, tilesNum int, discards, handCards, surplusCards, otherHands, mingKong, anKong, chowCards, pongCards []int64, isBanker, isZimo, robKong, lastOptIsKong, lastCard_cpk bool, lastCard, eyes int64, huTypeAndScore map[int64]int) {
	handCardsNum := len(handCards)
	surplusCardsNum := len(surplusCards)
	checkGang(len(anKong), len(mingKong), huTypeAndScore)
	checkJiuLianBaoDeng(handCards, lastCard, huTypeAndScore)
	checkZiYiSe(typeMapAll, huTypeAndScore)
	checkTianHu(handCardsNum, surplusCardsNum, isZimo, isBanker, huTypeAndScore)
	checkQingYiSe(typeMapAll, huTypeAndScore)
	checkDiHu(handCardsNum, surplusCardsNum, isZimo, isBanker, huTypeAndScore)
	checkGreaterLessFive(typeMapAll, huTypeAndScore)
	checkMiaoShow(surplusCardsNum, isZimo, huTypeAndScore)
	brc.checkeLianLiu(typeMapAll, huTypeAndScore)
	checkHaiDiLaoYue(surplusCardsNum, isZimo, lastCard_cpk, huTypeAndScore)
	checkGangKaiQiangGang(robKong, lastOptIsKong, isZimo, huTypeAndScore)
	checkDanDiao(eyes, lastCard, huTypeAndScore)
	checkeReWuBa(eyes, huTypeAndScore)
	checkZimo(isZimo, huTypeAndScore)
	checkHunYiSe(typeMapAll, huTypeAndScore)
	checkSiGuiYi(typeMapAll, len(mingKong), len(anKong), huTypeAndScore)
	checkBuQiuRen(chowCards, pongCards, mingKong, isZimo, huTypeAndScore)
	checkMenQing(chowCards, pongCards, mingKong, isZimo, huTypeAndScore)
	checkHuJueZhang(isZimo, lastCard, discards, handCards, huTypeAndScore)
	checkRenHu(handCardsNum, surplusCardsNum, isZimo, isBanker, huTypeAndScore)
	checkDuanYao(typeMapAll, huTypeAndScore)
}

//牌型算法
func (brc *BaseRuleChecker) ordinaryHu(typeMapAll map[int64]int, tilesNum int, discards, handCards, surplusCards, otherHands, chowCards, pongCards, mingKong, anKong, pongkongCards []int64, isBanker, isZimo, robKong, lastOptIsKong, lastCard_cpk bool, lastCard, tableFeng, menFeng int64, huResult HuHandResult, huTypeAndScore map[int64]int) {

	brc.checkYiSeShuangLong(huResult, chowCards, huTypeAndScore)
	brc.checkYiSeSanSiTong(huResult, chowCards, huTypeAndScore)
	brc.checkYiSeSanSiBu(huResult, chowCards, huTypeAndScore)
	checkPinghu(typeMapAll, huResult, chowCards, huTypeAndScore)
	brc.checkeYiBanGao(huResult, chowCards, huTypeAndScore)
	brc.checkeLaoShaoFu(huResult, chowCards, huTypeAndScore)
	brc.checkQingLong(huResult, chowCards, huTypeAndScore)
	brc.checkeBianZhang(huResult, lastCard, huTypeAndScore)
	brc.checkeKanZhang(huResult, lastCard, huTypeAndScore)
	checkDaXiaoSiXi(huResult, pongkongCards, huTypeAndScore)
	checkDaXiaoSanYuan(huResult, pongkongCards, huTypeAndScore)
	brc.checkAnKe(isZimo, lastCard, huResult, len(anKong), huTypeAndScore)
	brc.checkYiSeSanSiJie(huResult, pongkongCards, huTypeAndScore)
	checkHunYaoJiu(typeMapAll, huResult, pongkongCards, huTypeAndScore)
	checkSanFengKe(huResult, pongkongCards, huTypeAndScore)
	checkPengPeng(huResult, pongkongCards, huTypeAndScore)
	checkJianKe(huResult, pongkongCards, huTypeAndScore)
	checkQuanFengke(huResult, pongkongCards, tableFeng, huTypeAndScore)
	checkMenFengke(huResult, pongkongCards, menFeng, huTypeAndScore)
	brc.checkeYaoJiuKe(huResult, pongkongCards, huTypeAndScore)
	brc.checkQuanDaiYao(huResult, chowCards, pongkongCards, huTypeAndScore)
	brc.chekOtherHuType(typeMapAll, tilesNum, discards, handCards, surplusCards, otherHands, mingKong, anKong, chowCards, pongCards, isBanker, isZimo, robKong, lastOptIsKong, lastCard_cpk, lastCard, huResult.EyesKind, huTypeAndScore)
	brc.deleteRepeat(huTypeAndScore)
}

func (brc *BaseRuleChecker) deleteRepeat(huTypeAndScore map[int64]int) {

	if huTypeAndScore[XYWORDS_DASIXI] > 0 {
		delete(huTypeAndScore, XYWORDS_MENFENGKE)
		delete(huTypeAndScore, XYWORDS_QUANFENGKE)
		delete(huTypeAndScore, XYWORDS_XIAOSIXI)
		delete(huTypeAndScore, XYWORDS_SANFENGKE)
		delete(huTypeAndScore, XYWORDS_PENGPENGHU)
		delete(huTypeAndScore, XYWORDS_YAOJIUKE)
	}
	if huTypeAndScore[XYWORDS_DASANYUAN] > 0 {
		delete(huTypeAndScore, XYWORDS_XIAOSANYUAN)
		delete(huTypeAndScore, XYWORDS_JIANKE)
		delete(huTypeAndScore, XYWORDS_SHUANGJIANKE)
		delete(huTypeAndScore, XYWORDS_YAOJIUKE)
	}
	if huTypeAndScore[XYWORDS_JIULIANBAODENG] > 0 {
		delete(huTypeAndScore, XYWORDS_QINGYISE)
		delete(huTypeAndScore, XYWORDS_BUQIUREN)
		delete(huTypeAndScore, XYWORDS_MENQIANQING)
		delete(huTypeAndScore, XYWORDS_YAOJIUKE)
	}
	if huTypeAndScore[XYWORDS_FORKONG] > 0 {
		delete(huTypeAndScore, XYWORDS_DANDIAOJIANG)
	}
	if huTypeAndScore[XYWORDS_LIANQIDUI] > 0 {
		delete(huTypeAndScore, XYWORDS_QINGYISE)
		delete(huTypeAndScore, XYWORDS_BUQIUREN)
		delete(huTypeAndScore, XYWORDS_DANDIAOJIANG)
		delete(huTypeAndScore, XYWORDS_MENQIANQING)
		delete(huTypeAndScore, XYWORDS_QIDUI)
		delete(huTypeAndScore, XYWORDS_LIANLIU)
		delete(huTypeAndScore, XYWORDS_YIBANGAO)
	}
	if huTypeAndScore[XYWORDS_XIAOSIXI] > 0 {
		delete(huTypeAndScore, XYWORDS_SANFENGKE)
		delete(huTypeAndScore, XYWORDS_YAOJIUKE)
	}
	if huTypeAndScore[XYWORDS_XIAOSANYUAN] > 0 {
		delete(huTypeAndScore, XYWORDS_JIANKE)
		delete(huTypeAndScore, XYWORDS_SHUANGJIANKE)
		delete(huTypeAndScore, XYWORDS_YAOJIUKE)
	}
	if huTypeAndScore[XYWORDS_WORDONLYACOLOR] > 0 {
		delete(huTypeAndScore, XYWORDS_PENGPENGHU)
		delete(huTypeAndScore, XYWORDS_HUNYAOJIU)
		delete(huTypeAndScore, XYWORDS_YAOJIUKE)
		delete(huTypeAndScore, XYWORDS_QUANDAIYAO)
	}
	if huTypeAndScore[XYWORDS_SIANKE] > 0 {
		delete(huTypeAndScore, XYWORDS_MENQIANQING)
		delete(huTypeAndScore, XYWORDS_PENGPENGHU)
		delete(huTypeAndScore, XYWORDS_BUQIUREN)
	}
	if huTypeAndScore[XYWORDS_YISESHUANGLONG] > 0 {
		delete(huTypeAndScore, XYWORDS_PINGHU)
		delete(huTypeAndScore, XYWORDS_QIDUI)
		delete(huTypeAndScore, XYWORDS_QINGYISE)
		delete(huTypeAndScore, XYWORDS_YIBANGAO)
		delete(huTypeAndScore, XYWORDS_LAOSHAOFU)
	}
	if huTypeAndScore[XYWORDS_YISESITONGSHUN] > 0 {
		delete(huTypeAndScore, XYWORDS_YISESANJIEGAO)
		delete(huTypeAndScore, XYWORDS_YIBANGAO)
		delete(huTypeAndScore, XYWORDS_SIGUIYI)
	}
	if huTypeAndScore[XYWORDS_YISESIJIEGAO] > 0 {
		delete(huTypeAndScore, XYWORDS_YISESANTONGSHUN)
		delete(huTypeAndScore, XYWORDS_PENGPENGHU)
	}
	if huTypeAndScore[XYWORDS_TIANHU] > 0 {
		delete(huTypeAndScore, XYWORDS_DANDIAOJIANG)
		delete(huTypeAndScore, XYWORDS_LIANZHANG)
		delete(huTypeAndScore, XYWORDS_KANZHANG)
	}
	if huTypeAndScore[XYWORDS_SANGANG] > 0 {
		delete(huTypeAndScore, XYWORDS_SHUANGMINGGANG)
		delete(huTypeAndScore, XYWORDS_SHUANGANGANG)
		delete(huTypeAndScore, XYWORDS_MINGANG)
		delete(huTypeAndScore, XYWORDS_ANGANG)
	}
	if huTypeAndScore[XYWORDS_YISESIBUGAO] > 0 {
		delete(huTypeAndScore, XYWORDS_YISESANBUGAO)
		delete(huTypeAndScore, XYWORDS_LIANLIU)
		delete(huTypeAndScore, XYWORDS_LAOSHAOFU)
	}
	if huTypeAndScore[XYWORDS_HUNYAOJIU] > 0 {
		delete(huTypeAndScore, XYWORDS_PENGPENGHU)
		delete(huTypeAndScore, XYWORDS_YAOJIUKE)
		delete(huTypeAndScore, XYWORDS_QUANDAIYAO)
	}
	if huTypeAndScore[XYWORDS_YISESANTONGSHUN] > 0 {
		delete(huTypeAndScore, XYWORDS_YISESANJIEGAO)
		delete(huTypeAndScore, XYWORDS_YIBANGAO)
	}
	if huTypeAndScore[XYWORDS_YISESANJIEGAO] > 0 {
		delete(huTypeAndScore, XYWORDS_YISESANTONGSHUN)
	}

	if huTypeAndScore[XYWORDS_QIDUI] > 0 {
		delete(huTypeAndScore, XYWORDS_MENQIANQING)
		delete(huTypeAndScore, XYWORDS_BUQIUREN)
		delete(huTypeAndScore, XYWORDS_DANDIAOJIANG)
	}

	if huTypeAndScore[XYWORDS_QINGLONG] > 0 {
		delete(huTypeAndScore, XYWORDS_LIANLIU)
		delete(huTypeAndScore, XYWORDS_LAOSHAOFU)
	}
	if huTypeAndScore[XYWORDS_SANFENGKE] > 0 {
		delete(huTypeAndScore, XYWORDS_YAOJIUKE)
	}
	if huTypeAndScore[XYWORDS_MIAOSHOUHUICHUN] > 0 {
		delete(huTypeAndScore, XYWORDS_ZIMO)
	}
	if huTypeAndScore[XYWORDS_GANGSHANGKAIHUA] > 0 {
		delete(huTypeAndScore, XYWORDS_ZIMO)
	}
	if huTypeAndScore[XYWORDS_QIANGGANGHU] > 0 {
		delete(huTypeAndScore, XYWORDS_HUJUEZHANG)

	}
	if huTypeAndScore[XYWORDS_BUQIUREN] > 0 {
		delete(huTypeAndScore, XYWORDS_MENQIANQING)
		delete(huTypeAndScore, XYWORDS_ZIMO)
	}
	if huTypeAndScore[XYWORDS_HUJUEZHANG] > 0 {
		delete(huTypeAndScore, XYWORDS_QIANGGANGHU)
	}
	if huTypeAndScore[XYWORDS_JIANKE] > 0 || huTypeAndScore[XYWORDS_SHUANGJIANKE] > 0 {
		delete(huTypeAndScore, XYWORDS_YAOJIUKE)
	}
	if huTypeAndScore[XYWORDS_QUANFENGKE] > 0 {
		delete(huTypeAndScore, XYWORDS_YAOJIUKE)
	}
	if huTypeAndScore[XYWORDS_MENFENGKE] > 0 {
		delete(huTypeAndScore, XYWORDS_YAOJIUKE)
	}

}

//==================================麻将翻型=========================================

func checkGang(anGangNum, mingGangNum int, huTypeAndScore map[int64]int) {
	if anGangNum+mingGangNum == 4 {
		huTypeAndScore[XYWORDS_FORKONG] = getSocreByHuType(XYWORDS_FORKONG)
	} else if anGangNum+mingGangNum == 3 {
		huTypeAndScore[XYWORDS_SANGANG] = getSocreByHuType(XYWORDS_SANGANG)
	} else {
		switch anGangNum {
		case 2:
			huTypeAndScore[XYWORDS_SHUANGANGANG] = getSocreByHuType(XYWORDS_SHUANGANGANG)
		case 1:
			huTypeAndScore[XYWORDS_ANGANG] = getSocreByHuType(XYWORDS_ANGANG)
		}

		switch mingGangNum {
		case 2:
			huTypeAndScore[XYWORDS_SHUANGMINGGANG] = getSocreByHuType(XYWORDS_SHUANGMINGGANG)
		case 1:
			if anGangNum == 0 {
				//存在暗杠 不计明杠
				huTypeAndScore[XYWORDS_MINGANG] = getSocreByHuType(XYWORDS_MINGANG)
			}
		}
	}
}

/**
大四喜 小四喜判定
*/
func checkDaXiaoSiXi(huResult HuHandResult, pongkongCards []int64, huTypeAndScore map[int64]int) {
	keziCards := getCardsByFlag(huResult, 1)
	keziCards = copyAndAddSplic(keziCards, pongkongCards)
	count := getCountBySplic(keziCards, MJ_WIND_EAST, MJ_WIND_NORTH)
	switch count {
	case 4:
		huTypeAndScore[XYWORDS_DASIXI] = getSocreByHuType(XYWORDS_DASIXI)
	case 3:
		eyes := huResult.EyesKind
		if eyes >= MJ_WIND_EAST && eyes <= MJ_WIND_NORTH {
			huTypeAndScore[XYWORDS_XIAOSIXI] = getSocreByHuType(XYWORDS_XIAOSIXI)
		}
	}
}

/**
大小三元
*/
func checkDaXiaoSanYuan(huResult HuHandResult, pongkongCards []int64, huTypeAndScore map[int64]int) {
	keziCards := getCardsByFlag(huResult, 1)
	keziCards = copyAndAddSplic(keziCards, pongkongCards)
	count := getCountBySplic(keziCards, MJ_DRAGON_RED, MJ_DRAGON_WHITE)
	switch count {
	case 3:
		huTypeAndScore[XYWORDS_DASANYUAN] = getSocreByHuType(XYWORDS_DASANYUAN)
	case 2:
		eyes := huResult.EyesKind
		if eyes >= MJ_DRAGON_RED && eyes <= MJ_DRAGON_WHITE {
			huTypeAndScore[XYWORDS_XIAOSANYUAN] = getSocreByHuType(XYWORDS_XIAOSANYUAN)
		}
	}

}

/**
九莲宝灯
*/
func checkJiuLianBaoDeng(handCards []int64, card int64, huTypeAndScore map[int64]int) {
	handCardsMap := splicChangeMap(handCards)
	if len(handCardsMap) != 9 {
		return
	}
	count, _ := getCountByPrope(handCardsMap, MJ_WIND_EAST, MJ_DRAGON_WHITE, 1)
	if count > 0 {
		return
	}
	handCardsMap[card]--
	count1, _ := getCountByPrope(handCardsMap, MJ_CHARACTERS_2, MJ_CHARACTERS_8, 1)

	if count1 == 7 && handCardsMap[MJ_CHARACTERS_1] == 3 && handCardsMap[MJ_CHARACTERS_9] == 3 {
		huTypeAndScore[XYWORDS_JIULIANBAODENG] = getSocreByHuType(XYWORDS_JIULIANBAODENG)
	}
}

/**
七连队 TODO
*/
func checkQiLianDui(cards []int64, huTypeAndScore map[int64]int) {

	if len(cards) == 7 {
		for i := 0; i < 6; i++ {
			if cards[i]-cards[i+1] != -1 {
				return
			}
		}

		huTypeAndScore[XYWORDS_LIANQIDUI] = getSocreByHuType(XYWORDS_LIANQIDUI)
	}
}

/**
字一色
*/
func checkZiYiSe(typeMapALL map[int64]int, huTypeAndScore map[int64]int) {
	count, _ := getCountByPrope(typeMapALL, MJ_CHARACTERS_1, MJ_CHARACTERS_9, 1)
	if count > 0 {
		return
	}
	huTypeAndScore[XYWORDS_WORDONLYACOLOR] = getSocreByHuType(XYWORDS_WORDONLYACOLOR)
}

/**
暗刻
*/
func (brc *BaseRuleChecker) checkAnKe(isZimo bool, lastCard int64, huResult HuHandResult, angangNum int, huTypeAndScore map[int64]int) {
	keziCards := getCardsByFlag(huResult, 1)
	keNum := len(keziCards)
	if !isZimo && checkSplicInclodCard(keziCards, lastCard) {
		keNum--
	}
	switch keNum + angangNum {
	case 4:
		huTypeAndScore[XYWORDS_SIANKE] = getSocreByHuType(XYWORDS_SIANKE)
	case 3:
		huTypeAndScore[XYWORDS_SANANKE] = getSocreByHuType(XYWORDS_SANANKE)
	case 2:
		huTypeAndScore[XYWORDS_SHUANGANKE] = getSocreByHuType(XYWORDS_SHUANGANKE)
	}
}

/**
一色双龙会
*/
func (brc *BaseRuleChecker) checkYiSeShuangLong(huResult HuHandResult, chowCards []int64, huTypeAndScore map[int64]int) {
	if huResult.EyesKind != MJ_CHARACTERS_5 {
		return
	}
	shunCards := getCardsByFlag(huResult, 0)
	shunCards = copyAndAddSplic(shunCards, chowCards)
	characters1 := 0
	characters7 := 0
	for _, card := range shunCards {
		if card == MJ_CHARACTERS_1 {
			characters1++
		}
		if card == MJ_CHARACTERS_7 {
			characters7++
		}
	}

	if characters1 == 2 && characters7 == 2 {
		huTypeAndScore[XYWORDS_YISESHUANGLONG] = getSocreByHuType(XYWORDS_YISESHUANGLONG)
	}

}

/**
一色三四同顺
*/
func (brc *BaseRuleChecker) checkYiSeSanSiTong(huResult HuHandResult, chowCards []int64, huTypeAndScore map[int64]int) {
	shunCards := getCardsByFlag(huResult, 0)
	shunCards = copyAndAddSplic(shunCards, chowCards)
	if len(shunCards) < 3 {
		return
	}
	shunCount := 1
	for i := 0; i < len(shunCards)-1; i++ {
		if shunCards[i] <= MJ_CHARACTERS_9 && shunCards[i+1] <= MJ_CHARACTERS_9 && shunCards[i] == shunCards[i+1] {
			shunCount++
		} else {
			if shunCount < 3 {
				shunCount = 1
			}

		}
	}
	switch shunCount {
	case 4:
		huTypeAndScore[XYWORDS_YISESITONGSHUN] = getSocreByHuType(XYWORDS_YISESITONGSHUN)
	case 3:
		huTypeAndScore[XYWORDS_YISESANTONGSHUN] = getSocreByHuType(XYWORDS_YISESANTONGSHUN)
	}

}

/**
一色三/四节高
*/
func (brc *BaseRuleChecker) checkYiSeSanSiJie(huResult HuHandResult, pongGangCards []int64, huTypeAndScore map[int64]int) {
	keCards := getCardsByFlag(huResult, 1)
	keCards = copyAndAddSplic(keCards, pongGangCards)
	if len(keCards) < 3 {
		return
	}
	count := 1
	for i := 0; i < len(keCards)-1; i++ {
		if keCards[i] <= MJ_CHARACTERS_9 && keCards[i+1] <= MJ_CHARACTERS_9 && keCards[i]-keCards[i+1] == -1 {
			count++
		} else {
			if count < 3 {
				count = 1
			}
		}
	}
	switch count {
	case 4:
		huTypeAndScore[XYWORDS_YISESIJIEGAO] = getSocreByHuType(XYWORDS_YISESIJIEGAO)
	case 3:
		huTypeAndScore[XYWORDS_YISESANJIEGAO] = getSocreByHuType(XYWORDS_YISESANJIEGAO)
	}
}

/**
天胡
*/
func checkTianHu(handCardsNum, surplusCardsNum int, isZiMo, isBanker bool, huTypeAndScore map[int64]int) {
	if isBanker && isZiMo && handCardsNum == 14 && surplusCardsNum == 64-27 {
		huTypeAndScore[XYWORDS_TIANHU] = getSocreByHuType(XYWORDS_TIANHU)
	}
}

/**
一色三步四步高
*/
func (brc *BaseRuleChecker) checkYiSeSanSiBu(huResult HuHandResult, chowCards []int64, huTypeAndScore map[int64]int) {

	shunCards := getCardsByFlag(huResult, 0)
	shunCards = copyAndAddSplic(shunCards, chowCards)
	shunCards = removeRepByLoop(shunCards)
	if len(shunCards) < 3 {
		return
	}
	count := 1
	for i := 0; i < len(shunCards)-1; i++ {
		if shunCards[i] <= MJ_CHARACTERS_9 && shunCards[i+1] <= MJ_CHARACTERS_9 && (shunCards[i]-shunCards[i+1] == -1 || shunCards[i]-shunCards[i+1] == -2) {
			count++
		} else {
			if count < 3 {
				count = 1
			}
		}
	}
	switch count {
	case 4:
		huTypeAndScore[XYWORDS_YISESIBUGAO] = getSocreByHuType(XYWORDS_YISESIBUGAO)
	case 3:
		huTypeAndScore[XYWORDS_YISESANBUGAO] = getSocreByHuType(XYWORDS_YISESANBUGAO)
	}
}

/**
混幺九
*/
func checkHunYaoJiu(typeMapAll map[int64]int, huResult HuHandResult, pongGangCards []int64, huTypeAndScore map[int64]int) {
	count, _ := getCountByPrope(typeMapAll, MJ_CHARACTERS_2, MJ_CHARACTERS_8, 1)
	if count > 0 {
		return
	}
	keCards := getCardsByFlag(huResult, 1)
	keCards = copyAndAddSplic(keCards, pongGangCards)
	//碰杠牌带1、9
	card1 := checkSplicInclodCard(keCards, MJ_CHARACTERS_1)
	card9 := checkSplicInclodCard(keCards, MJ_CHARACTERS_9)
	//或者 将牌带1、9
	pair := checkPairInCard(typeMapAll, MJ_CHARACTERS_1, MJ_CHARACTERS_9)
	if card1 || card9 || pair {
		huTypeAndScore[XYWORDS_HUNYAOJIU] = getSocreByHuType(XYWORDS_HUNYAOJIU)
	}
}

/**
清一色
*/
func checkQingYiSe(typeMapAll map[int64]int, huTypeAndScore map[int64]int) {
	count, _ := getCountByPrope(typeMapAll, MJ_WIND_EAST, MJ_DRAGON_WHITE, 1)
	if count > 0 {
		return
	}
	huTypeAndScore[XYWORDS_QINGYISE] = getSocreByHuType(XYWORDS_QINGYISE)
}

/**
七对 连七对
*/
func checkQiDui(handCards []int64) bool {

	cardsMap := splicChangeMap(handCards)
	count, cards := getCountByPrope1(cardsMap, MJ_CHARACTERS_1, MJ_CHARACTERS_9, 2)
	if count < 7 {
		return false
	}
	state := true
	for i := 0; i < 6; i++ {
		if cards[i]-cards[i+1] != -1 {
			state = false
			break
		}
	}
	if state {
		return true
	} else {
		return false
	}

}

/**
地胡
*/
func checkDiHu(handCardsNum, surplusCardsNum int, iszimo, isBanker bool, huTypeAndScore map[int64]int) {

	if iszimo && !isBanker && handCardsNum >= 13 && surplusCardsNum == 64-27-1 {
		huTypeAndScore[XYWORDS_DIHU] = getSocreByHuType(XYWORDS_DIHU)
	}

}

/**
青龙
*/
func (brc *BaseRuleChecker) checkQingLong(huResult HuHandResult, chowCards []int64, huTypeAndScore map[int64]int) {

	shunCards := getCardsByFlag(huResult, 0)
	shunCards = copyAndAddSplic(shunCards, chowCards)
	if len(shunCards) < 3 {
		return
	}
	card1 := checkSplicInclodCard(shunCards, MJ_CHARACTERS_1)
	card4 := checkSplicInclodCard(shunCards, MJ_CHARACTERS_4)
	card7 := checkSplicInclodCard(shunCards, MJ_CHARACTERS_7)
	if card1 && card4 && card7 {
		huTypeAndScore[XYWORDS_QINGLONG] = getSocreByHuType(XYWORDS_QINGLONG)
	}
}

/**
人胡
*/
func checkRenHu(handCardsNum, surplusCardsNum int, isziMo, isBanker bool, huTypeAndScore map[int64]int) {

	if !isziMo && !isBanker && handCardsNum >= 13 && surplusCardsNum == 64-27 {
		huTypeAndScore[XYWORDS_RENHU] = getSocreByHuType(XYWORDS_RENHU)
	}
}

/**
大于五小于五
*/
func checkGreaterLessFive(typeMapAll map[int64]int, huTypeAndScore map[int64]int) {
	count1, _ := getCountByPrope(typeMapAll, MJ_WIND_EAST, MJ_DRAGON_WHITE, 1)
	if count1 > 0 {
		return
	}
	count, _ := getCountByPrope(typeMapAll, MJ_CHARACTERS_1, MJ_CHARACTERS_5, 1)

	if count == 0 {
		huTypeAndScore[XYWORDS_DAYUWU] = getSocreByHuType(XYWORDS_DAYUWU)
	}
	count, _ = getCountByPrope(typeMapAll, MJ_CHARACTERS_5, MJ_CHARACTERS_9, 1)
	if count == 0 {
		huTypeAndScore[XYWORDS_XIAOYUWU] = getSocreByHuType(XYWORDS_XIAOYUWU)
	}

}

/**
三风刻
*/
func checkSanFengKe(huResult HuHandResult, pengkongCards []int64, huTypeAndScore map[int64]int) {

	cardsCopy := getCardsByFlag(huResult, 1)
	cardsCopy = copyAndAddSplic(cardsCopy, pengkongCards)
	if len(cardsCopy) < 3 {
		return
	}
	count := getCountBySplic(cardsCopy, MJ_WIND_EAST, MJ_WIND_NORTH)
	if count >= 3 {
		huTypeAndScore[XYWORDS_SANFENGKE] = getSocreByHuType(XYWORDS_SANFENGKE)
	}
}

/**
妙手回春
*/
func checkMiaoShow(surplusCardsNum int, isZimo bool, huTypeAndScore map[int64]int) {

	if surplusCardsNum == 0 && isZimo {
		huTypeAndScore[XYWORDS_MIAOSHOUHUICHUN] = getSocreByHuType(XYWORDS_MIAOSHOUHUICHUN)
	}

}

/**
海底捞月
*/
func checkHaiDiLaoYue(surplusCardsNum int, isZimo, lastCard_cpk bool, huTypeAndScore map[int64]int) {

	if surplusCardsNum == 0 && !isZimo && lastCard_cpk {
		huTypeAndScore[XYWORDS_HAISILAOYUE] = getSocreByHuType(XYWORDS_HAISILAOYUE)
	}

}

/**
杠上开花抢杠胡
*/
func checkGangKaiQiangGang(robKong, lastOptIsKong, isZimo bool, huTypeAndScore map[int64]int) {

	if robKong {
		huTypeAndScore[XYWORDS_QIANGGANGHU] = getSocreByHuType(XYWORDS_QIANGGANGHU)
		return
	}
	if lastOptIsKong && isZimo {
		huTypeAndScore[XYWORDS_GANGSHANGKAIHUA] = getSocreByHuType(XYWORDS_GANGSHANGKAIHUA)
	}
}

/**
是否是自摸
*/
func checkZimo(isZimo bool, huTypeAndScore map[int64]int) {

	if isZimo {
		huTypeAndScore[XYWORDS_ZIMO] = getSocreByHuType(XYWORDS_ZIMO)
	}

}

/**
碰碰胡
*/
func checkPengPeng(huResult HuHandResult, pongkongCards []int64, huTypeAndScore map[int64]int) {
	cardsCopy := getCardsByFlag(huResult, 1)
	cardsCopy = copyAndAddSplic(cardsCopy, pongkongCards)
	if len(cardsCopy) > 3 {
		huTypeAndScore[XYWORDS_PENGPENGHU] = getSocreByHuType(XYWORDS_PENGPENGHU)
	}

}

/**
混一色
*/
func checkHunYiSe(typeMapAll map[int64]int, huTypeAndScore map[int64]int) {
	count, _ := getCountByPrope(typeMapAll, MJ_WIND_EAST, MJ_DRAGON_WHITE, 1)
	count1, _ := getCountByPrope(typeMapAll, MJ_CHARACTERS_1, MJ_CHARACTERS_9, 1)
	if count == 0 || count1 == 0 {
		return
	}
	huTypeAndScore[XYWORDS_HUNYISE] = getSocreByHuType(XYWORDS_HUNYISE)
}

/**
检查箭刻/双箭刻 todo 小三元
*/
func checkJianKe(huResult HuHandResult, pongkongCards []int64, huTypeAndScore map[int64]int) {

	cardsCopy := getCardsByFlag(huResult, 1)
	cardsCopy = copyAndAddSplic(cardsCopy, pongkongCards)
	count := getCountBySplic(cardsCopy, MJ_DRAGON_RED, MJ_DRAGON_WHITE)
	switch count {
	case 2:
		huTypeAndScore[XYWORDS_SHUANGJIANKE] = getSocreByHuType(XYWORDS_SHUANGJIANKE)
	case 1:
		huTypeAndScore[XYWORDS_JIANKE] = getSocreByHuType(XYWORDS_JIANKE)
	}
}

/**
全带幺
*/
func (brc *BaseRuleChecker) checkQuanDaiYao(huResult HuHandResult, chowCards, pongkongCards []int64, huTypeAndScore map[int64]int) {
	for i := 0; i < len(huResult.Sequence); i++ {
		card := huResult.Sequence[i][0]
		if card > 0 && huResult.Flags[i] == 0 {
			if card != MJ_CHARACTERS_1 && card != MJ_CHARACTERS_7 {
				return
			}
		}
		if card > 0 && card <= MJ_CHARACTERS_9 && huResult.Flags[i] == 1 {
			if card != MJ_CHARACTERS_1 && card != MJ_CHARACTERS_9 {
				return
			}
		}
	}
	for _, card := range chowCards {
		if card != MJ_CHARACTERS_1 && card != MJ_CHARACTERS_7 {
			return
		}
	}

	for _, card := range pongkongCards {
		if card > 0 && card <= MJ_CHARACTERS_9 && card != MJ_CHARACTERS_1 && card != MJ_CHARACTERS_9 {
			return
		}
	}

	eyes := huResult.EyesKind
	if eyes > 0 && eyes <= MJ_CHARACTERS_9 && eyes != MJ_CHARACTERS_1 && eyes != MJ_CHARACTERS_9 {
		return
	}

	huTypeAndScore[XYWORDS_QUANDAIYAO] = getSocreByHuType(XYWORDS_QUANDAIYAO)
}

/**
不求人
*/
func checkBuQiuRen(chowCards, pongCards, mingKongCards []int64, isZiMo bool, huTypeAndScore map[int64]int) {
	if len(chowCards) == 0 && len(pongCards) == 0 && len(mingKongCards) == 0 && isZiMo {
		huTypeAndScore[XYWORDS_BUQIUREN] = getSocreByHuType(XYWORDS_BUQIUREN)
	}
}

/**
门清
*/
func checkMenQing(chowCards, pongCards, mingKongCards []int64, isZiMo bool, huTypeAndScore map[int64]int) {
	if len(chowCards) == 0 && len(pongCards) == 0 && len(mingKongCards) == 0 && !isZiMo {
		huTypeAndScore[XYWORDS_MENQIANQING] = getSocreByHuType(XYWORDS_MENQIANQING)
	}
}

/**
胡绝张
*/
func checkHuJueZhang(isZiMo bool, card int64, discards, hands []int64, huTypeAndScore map[int64]int) {

	disCardsMap := splicChangeMap(discards)
	if !isZiMo {
		disCardsMap[card]--
	}
	handsMap := splicChangeMap(hands)
	if disCardsMap[card] >= 3 && handsMap[card] == 1 {
		huTypeAndScore[XYWORDS_HUJUEZHANG] = getSocreByHuType(XYWORDS_HUJUEZHANG)
	}
}

/**
圈风刻
*/
func checkQuanFengke(huResult HuHandResult, pongkongCards []int64, tableFeng int64, huTypeAndScore map[int64]int) {
	cardsCopy := getCardsByFlag(huResult, 1)
	cardsCopy = copyAndAddSplic(cardsCopy, pongkongCards)
	if checkSplicInclodCard(cardsCopy, tableFeng) {
		huTypeAndScore[XYWORDS_QUANFENGKE] = getSocreByHuType(XYWORDS_QUANFENGKE)
	}
}

/**
门风刻
*/
func checkMenFengke(huResult HuHandResult, pongkongCards []int64, menFeng int64, huTypeAndScore map[int64]int) {
	cardsCopy := getCardsByFlag(huResult, 1)
	cardsCopy = copyAndAddSplic(cardsCopy, pongkongCards)
	if menFeng == 1 {
		menFeng++
	}
	if checkSplicInclodCard(cardsCopy, menFeng+MJ_WIND_EAST) {
		huTypeAndScore[XYWORDS_MENFENGKE] = getSocreByHuType(XYWORDS_MENFENGKE)
	}

}

/**
平湖
*/
func checkPinghu(typeMapAll map[int64]int, huResult HuHandResult, chowCards []int64, huTypeAndScore map[int64]int) {
	count, _ := getCountByPrope(typeMapAll, MJ_WIND_EAST, MJ_DRAGON_WHITE, 1)
	if count > 0 {
		return
	}
	cardsCopy := getCardsByFlag(huResult, 0)
	if len(cardsCopy)+len(chowCards) == 4 {
		huTypeAndScore[XYWORDS_PINGHU] = getSocreByHuType(XYWORDS_PINGHU)
	}
}

/**
断幺
*/
func checkDuanYao(typeMapAll map[int64]int, huTypeAndScore map[int64]int) {
	count, _ := getCountByPrope(typeMapAll, MJ_WIND_EAST, MJ_DRAGON_WHITE, 1)
	if count == 0 && typeMapAll[MJ_CHARACTERS_1] == 0 && typeMapAll[MJ_CHARACTERS_9] == 0 {
		huTypeAndScore[XYWORDS_DUANYAO] = getSocreByHuType(XYWORDS_DUANYAO)
	}
}

/**
二五八将/幺九头
*/
func checkeReWuBa(eyes int64, huTypeAndScore map[int64]int) {
	if eyes == MJ_CHARACTERS_2 || eyes == MJ_CHARACTERS_5 || eyes == MJ_CHARACTERS_8 {
		huTypeAndScore[XYWORDS_ERWUBAJIANG] = getSocreByHuType(XYWORDS_ERWUBAJIANG)
	}

	if eyes == MJ_CHARACTERS_1 || eyes == MJ_CHARACTERS_9 {
		huTypeAndScore[XYWORDS_YAOJIUTOU] = getSocreByHuType(XYWORDS_YAOJIUTOU)
	}
}

/**
一般高
*/
func (brc *BaseRuleChecker) checkeYiBanGao(huResult HuHandResult, chowCards []int64, huTypeAndScore map[int64]int) {
	cardsCopy := getCardsByFlag(huResult, 0)
	cardsCopy = copyAndAddSplic(cardsCopy, chowCards)
	if len(cardsCopy) < 2 {
		return
	}
	count := 1
	for i := 0; i < len(cardsCopy)-1; i++ {
		if cardsCopy[i] == cardsCopy[i+1] {
			count++
			if count >= 2 {
				break
			}
		} else {
			count = 1
		}
	}
	if count >= 2 {
		huTypeAndScore[XYWORDS_YIBANGAO] = getSocreByHuType(XYWORDS_YIBANGAO)
	}
}

/**
连六
*/
func (brc *BaseRuleChecker) checkeLianLiu(typeMapAll map[int64]int, huTypeAndScore map[int64]int) {
	cards := mapChangeSplic(typeMapAll)
	count := 1
	for i := 0; i < len(cards)-1; i++ {
		if cards[i+1] > MJ_CHARACTERS_9 {
			break
		}
		if cards[i]-cards[i+1] == -1 {
			count++
			if count >= 6 {
				break
			}
		} else {
			count = 1
		}
	}
	if count >= 6 {
		huTypeAndScore[XYWORDS_LIANLIU] = getSocreByHuType(XYWORDS_LIANLIU)
	}
}

/**
老少副
*/
func (brc *BaseRuleChecker) checkeLaoShaoFu(huResult HuHandResult, chowCards []int64, huTypeAndScore map[int64]int) {
	cardsCopy := getCardsByFlag(huResult, 0)
	cardsCopy = copyAndAddSplic(cardsCopy, chowCards)
	if checkSplicInclodCard(cardsCopy, MJ_CHARACTERS_1) && checkSplicInclodCard(cardsCopy, MJ_CHARACTERS_7) {
		huTypeAndScore[XYWORDS_LAOSHAOFU] = getSocreByHuType(XYWORDS_LAOSHAOFU)
	}
}

/**
幺九刻
*/
func (brc *BaseRuleChecker) checkeYaoJiuKe(huResult HuHandResult, pongkongCards []int64, huTypeAndScore map[int64]int) {
	cardsCopy := getCardsByFlag(huResult, 1)
	cardsCopy = copyAndAddSplic(cardsCopy, pongkongCards)
	for _, card := range cardsCopy {
		if card == MJ_CHARACTERS_1 || card == MJ_CHARACTERS_9 || (card >= MJ_WIND_EAST && card <= MJ_WIND_NORTH) {
			huTypeAndScore[XYWORDS_YAOJIUKE] = getSocreByHuType(XYWORDS_YAOJIUKE)
			return
		}
	}

}

/**
边张
*/
func (brc *BaseRuleChecker) checkeBianZhang(huResult HuHandResult, card int64, huTypeAndScore map[int64]int) {
	if card > MJ_CHARACTERS_9 {
		return
	}
	cardsCopy := getCardsByFlag(huResult, 0)
	if card == MJ_CHARACTERS_3 {
		char1 := checkSplicInclodCard(cardsCopy, MJ_CHARACTERS_1)
		char2 := checkSplicInclodCard(cardsCopy, MJ_CHARACTERS_2)
		char3 := checkSplicInclodCard(cardsCopy, MJ_CHARACTERS_3)
		if char1 && !char2 && !char3 {
			huTypeAndScore[XYWORDS_LIANZHANG] = getSocreByHuType(XYWORDS_LIANZHANG)
			delete(huTypeAndScore, XYWORDS_DANDIAOJIANG)
		}
	}
	if card == MJ_CHARACTERS_7 {
		char5 := checkSplicInclodCard(cardsCopy, MJ_CHARACTERS_5)
		char6 := checkSplicInclodCard(cardsCopy, MJ_CHARACTERS_6)
		char7 := checkSplicInclodCard(cardsCopy, MJ_CHARACTERS_7)
		if !char5 && !char6 && char7 {
			huTypeAndScore[XYWORDS_LIANZHANG] = getSocreByHuType(XYWORDS_LIANZHANG)
		}
	}

}

/**
坎张
*/
func (brc *BaseRuleChecker) checkeKanZhang(huResult HuHandResult, card int64, huTypeAndScore map[int64]int) {
	if card > MJ_CHARACTERS_9 {
		return
	}
	cardsCopy := getCardsByFlag(huResult, 0)
	state2 := checkSplicInclodCard(cardsCopy, card-2)
	state1 := checkSplicInclodCard(cardsCopy, card-1)
	state := checkSplicInclodCard(cardsCopy, card)
	if state1 && !state && !state2 {
		huTypeAndScore[XYWORDS_KANZHANG] = getSocreByHuType(XYWORDS_KANZHANG)
	}
}

/**
单调
*/
func checkDanDiao(eyes, card int64, huTypeAndScore map[int64]int) {
	if eyes == card {
		huTypeAndScore[XYWORDS_DANDIAOJIANG] = getSocreByHuType(XYWORDS_DANDIAOJIANG)
		delete(huTypeAndScore, XYWORDS_KANZHANG)
		delete(huTypeAndScore, XYWORDS_LIANZHANG)
	}

}

/**
四归一
*/
func checkSiGuiYi(typeMapAll map[int64]int, mingkong, ankong int, huTypeAndScore map[int64]int) {

	count, _ := getCountByPrope(typeMapAll, MJ_CHARACTERS_1, MJ_DRAGON_WHITE, 4)
	if count-mingkong-ankong > 0 {
		huTypeAndScore[XYWORDS_SIGUIYI] = getSocreByHuType(XYWORDS_SIGUIYI)
	}
}

/**
查询大于等于规定数量的牌有几种
type：牌型对应的数量，包括吃碰杠
begin：开始范围
end：结束范围
return:
	返回数量值
	具体的卡牌切片

*/
func getCountByPrope(typeMap map[int64]int, begin, end int64, num int) (int, []int64) {
	count := 0
	cards := make([]int64, 0, 5)
	for key, value := range typeMap {
		if key >= begin && key <= end && value >= num {
			count++
			cards = append(cards, key)
		}
	}
	return count, sortCards(cards)
}

/**
查询规定数量的牌有几种
type：牌型对应的数量，包括吃碰杠
begin：开始范围
end：结束范围
return:
	返回数量值
	具体的卡牌切片

*/
func getCountByPrope1(typeMap map[int64]int, begin, end int64, num int) (int, []int64) {
	count := 0
	cards := make([]int64, 0, 5)
	for key, value := range typeMap {
		if key >= begin && key <= end && value == num {
			count++
			cards = append(cards, key)
		}
	}
	return count, sortCards(cards)
}

func getCountBySplic(cards []int64, begin, end int64) int {
	count := 0

	for _, card := range cards {
		if card >= begin && card <= end {
			count++
		}
	}
	return count
}

func getTotalScoreByMap(typeScore map[int64]int) int {
	totalScore := 0
	for _, score := range typeScore {
		totalScore += score
	}
	return totalScore
}

/**
 	根据需求获取相应的cards
	flag 0 顺子 1 刻字
*/
func getCardsByFlag(huResult HuHandResult, fla int) []int64 {
	cards := make([]int64, 0)
	for i := 0; i < 4; i++ {
		flag := huResult.Flags[i]
		if flag == int64(fla) {
			card := huResult.Sequence[i][0]
			if card > 0 {
				cards = append(cards, huResult.Sequence[i][0])
			}

		}
	}
	return sortCards(cards)
}

/**
 	根据需求获取相应的cards
	flag 0 顺子 1 刻字
*/
func getAllCardsByFlag(huResult HuHandResult, fla int) []int64 {
	cards := make([]int64, 0)
	for i := 0; i < 4; i++ {
		flag := huResult.Flags[i]
		if flag == int64(fla) {
			card := huResult.Sequence[i][0]
			if card > 0 {
				cards = append(cards, huResult.Sequence[i][0])
				cards = append(cards, huResult.Sequence[i][1])
				cards = append(cards, huResult.Sequence[i][2])
			}

		}
	}
	return sortCards(cards)
}

func checkSplicInclodCard(cards []int64, card int64) bool {
	for _, key := range cards {
		if key == card {
			return true
		}
	}
	return false
}
func checkPairInCard(cards map[int64]int, card1, card2 int64) bool {
	for k, v := range cards {
		if v == 2 {
			if k == card1 || k == card2 {
				return true
			}
		}
	}
	return false
}

/**
 	复制并返回新数组
	flag 0 顺子 1 刻字
*/
func copyAndAddSplic(cards, cards1 []int64) []int64 {
	cardsCopy := make([]int64, len(cards))
	copy(cardsCopy, cards)
	cardsCopy = append(cardsCopy, cards1...)
	return sortCards(cardsCopy)
}

func removeRepByLoop(slc []int64) []int64 {
	result := []int64{}
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, slc[i])
		}
	}
	return result
}

func sortCards(cards []int64) []int64 {
	for i := 0; i < len(cards)-1; i++ {
		for j := 0; j < (len(cards) - 1 - i); j++ {
			if (cards)[j] > (cards)[j+1] {
				temp := (cards)[j]
				(cards)[j] = (cards)[j+1]
				(cards)[j+1] = temp
			}
		}
	}
	return cards
}
func splicChangeMap(sp []int64) map[int64]int {
	typeMap := make(map[int64]int)
	for _, value := range sp {
		if value > 0 {
			typeMap[value]++
		}
	}
	return typeMap
}

func mapChangeSplic(cardMap map[int64]int) []int64 {
	cards := make([]int64, 0)
	for card, _ := range cardMap {
		cards = append(cards, card)
	}
	return sortCards(cards)
}

func printMap(typeMap map[int64]int) string {
	str := ""
	for key, value := range typeMap {
		str += strconv.FormatInt(key, 10) + ":" + strconv.Itoa(value) + ","
	}
	return string(str)
}

func appendString(huTypeAndScore map[int64]int) string {
	str := ""
	for key, _ := range huTypeAndScore {
		switch key {
		case XYWORDS_DASIXI:
			str += "大四喜,"
		case XYWORDS_DASANYUAN:
			str += "大三元,"
		case XYWORDS_JIULIANBAODENG:
			str += "九连宝灯,"
		case XYWORDS_FORKONG:
			str += "四杠,"
		case XYWORDS_LIANQIDUI:
			str += "连七对,"
		case XYWORDS_XIAOSIXI:
			str += "小四喜,"
		case XYWORDS_XIAOSANYUAN:
			str += "小三元,"
		case XYWORDS_WORDONLYACOLOR:
			str += "字一色,"
		case XYWORDS_SIANKE:
			str += "四暗刻,"
		case XYWORDS_YISESHUANGLONG:
			str += "一色双龙会,"
		case XYWORDS_YISESITONGSHUN:
			str += "一色四同顺,"
		case XYWORDS_YISESIJIEGAO:
			str += "一色四节高,"
		case XYWORDS_TIANHU:
			str += "天湖,"
		case XYWORDS_YISESIBUGAO:
			str += "一色四步高,"
		case XYWORDS_SANGANG:
			str += "三杠,"
		case XYWORDS_HUNYAOJIU:
			str += "混幺九,"
		case XYWORDS_QINGYISE:
			str += "清一色,"
		case XYWORDS_YISESANTONGSHUN:
			str += "一色三同顺,"
		case XYWORDS_YISESANJIEGAO:
			str += "一色三节高,"
		case XYWORDS_QIDUI:
			str += "七对,"
		case XYWORDS_DIHU:
			str += "地胡,"
		case XYWORDS_QINGLONG:
			str += "青龙,"
		case XYWORDS_YISESANBUGAO:
			str += "一色三步高,"
		case XYWORDS_SANANKE:
			str += "三暗刻,"
		case XYWORDS_RENHU:
			str += "人胡,"
		case XYWORDS_DAYUWU:
			str += "大于五,"
		case XYWORDS_XIAOYUWU:
			str += "小于五,"
		case XYWORDS_SANFENGKE:
			str += "三风刻,"
		case XYWORDS_MIAOSHOUHUICHUN:
			str += "妙手回春,"
		case XYWORDS_HAISILAOYUE:
			str += "海底捞月,"
		case XYWORDS_GANGSHANGKAIHUA:
			str += "杠上开花,"
		case XYWORDS_QIANGGANGHU:
			str += "抢杠胡,"
		case XYWORDS_ZIMO:
			str += "自摸,"
		case XYWORDS_PENGPENGHU:
			str += "碰碰胡,"
		case XYWORDS_HUNYISE:
			str += "混一色,"
		case XYWORDS_SHUANGJIANKE:
			str += "双箭刻,"
		case XYWORDS_SHUANGANGANG:
			str += "双暗杠,"
		case XYWORDS_QUANDAIYAO:
			str += "全带幺,"
		case XYWORDS_BUQIUREN:
			str += "不求人,"
		case XYWORDS_SHUANGMINGGANG:
			str += "双明杠,"
		case XYWORDS_HUJUEZHANG:
			str += "胡绝张,"
		case XYWORDS_JIANKE:
			str += "箭刻,"
		case XYWORDS_QUANFENGKE:
			str += "圈风刻,"
		case XYWORDS_MENFENGKE:
			str += "门风刻,"
		case XYWORDS_MENQIANQING:
			str += "门前清,"
		case XYWORDS_PINGHU:
			str += "平胡,"
		case XYWORDS_SIGUIYI:
			str += "四归一,"
		case XYWORDS_SHUANGANKE:
			str += "双暗刻,"
		case XYWORDS_DUANYAO:
			str += "断幺,"
		case XYWORDS_ANGANG:
			str += "暗杠,"
		case XYWORDS_ERWUBAJIANG:
			str += "258将,"
		case XYWORDS_YIBANGAO:
			str += "一般高,"
		case XYWORDS_LIANLIU:
			str += "连六,"
		case XYWORDS_LAOSHAOFU:
			str += "老少副,"
		case XYWORDS_YAOJIUKE:
			str += "幺九刻,"
		case XYWORDS_MINGANG:
			str += "明杠,"
		case XYWORDS_WUZI:
			str += "无字,"
		case XYWORDS_LIANZHANG:
			str += "边张,"
		case XYWORDS_KANZHANG:
			str += "坎张,"
		case XYWORDS_DANDIAOJIANG:
			str += "单调将,"
		case XYWORDS_YAOJIUTOU:
			str += "幺九头"
		}
	}
	return string(str)
}
