package mahjong

import (
	"sort"
)

type MjAlgorithm struct {
}

type cv struct {
	value int64
	index []int
}

type card struct {
	value int64
	cards []int
}

type MJResult struct {
	qidui   bool
	laizi4  bool
	temp    HuHandResult
	Results []HuHandResult
}

type MJSpecailConfig struct {
	FengPu         bool //风扑
	JiangPu        bool //将扑
	OnlyChk        bool //查到胡即可
	HuQiDui        int  //
}

type Int64Slice []int64

func (p Int64Slice) Len() int           { return len(p) }
func (p Int64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Int64Slice) Sort()              { sort.Sort(p) }

func removeCard(zu []cv, k int) []cv {
	if k >= 0 && k < len(zu) {
		cnt := len(zu[k].index)
		if cnt > 0 {
			zu[k].index = zu[k].index[:cnt-1]
		}
		if len(zu[k].index) == 0 {
			count := len(zu)
			if k == 0 {
				zu = zu[1:]
			} else if k == count-1 {
				zu = zu[:count-1]
			} else {
				arr := zu[:k]
				arr = append(arr, zu[k+1:]...)
				zu = arr
			}
		}
	}
	return zu
}

func insertCard(arr []card, value card) []card {
	for i := 0; i < len(arr); i++ {
		if arr[i].value == value.value {
			arr[i].cards = append(arr[i].cards, value.cards...)
			return arr
		}
	}
	arr = append(arr, value)
	return arr
}

func getCard(zu []cv, k int) card {
	if k >= 0 && k < len(zu) {
		cnt := len(zu[k].index)
		if cnt > 0 {
			return card{value: zu[k].value, cards: []int{zu[k].index[cnt-1]}}
		}
	}
	return card{value: -1, cards: nil}
}

func cloneq(zu []cv) []cv {
	cnt := len(zu)
	cache := make([]cv, cnt, cnt)
	copy(cache, zu)
	return cache
}

func (this *MjAlgorithm) analysis(hands []int64) []cv {
	var zu []cv
	t := cv{value: -1}
	for k, c := range hands {
		if t.value != c {
			if len(t.index) != 0 {
				zu = append(zu, t)
				t = cv{value: -1}
			}
			t.value = c
		}
		t.index = append(t.index, k)
	}
	if len(t.index) != 0 {
		zu = append(zu, t)
	}
	return zu
}

func (this *MjAlgorithm) Hu(myCard []int64, otherCard []int64, lastCard int64, laizi []int64, config *MJSpecailConfig) *MJResult {
	cnt := len(myCard)
	cc := make([]int64, cnt, cnt)
	copy(cc, myCard)

	if lastCard >= 0 && lastCard < MJ_MAX {
		cc = append(cc, lastCard)
	}
	sort.Sort(Int64Slice(cc))

	var laiziCard []int64
	var cards []int64
	//拆分出癞子和普通牌
	for _, v := range cc {
		if in(laizi, v) {
			laiziCard = append(laiziCard, v)
		} else {
			cards = append(cards, v)
		}
	}

	var temp int64
	result := &MJResult{}
	//4癞子
	laiziCnt := len(laiziCard)
	if laiziCnt == 4 {
		if laiziCnt >= 7 {
			result.laizi4 = true
		} else {
			sort.Sort(Int64Slice(laiziCard))
			var tmp int
			val := laiziCard[0]
			for i := 1; i < len(laiziCard); i++ {
				if val == laiziCard[i] {
					tmp++
				} else {
					tmp = 0
				}
				if tmp >= 3 {
					result.laizi4 = true
				}
			}
		}
	}

	//4赖
	if result.laizi4 {

	}

	//牌整理成组
	zu := this.analysis(cards)
	//7对判断
	if len(otherCard) == 0 && config.HuQiDui > 0 {
		//7小对
		var item1 []card
		var laiziZu []int64
		zu1 := cloneq(zu)
		laiziCnt := len(laiziCard)
		//1代表混牌可以通配，2 代表不能通配，1不用处理
		if config.HuQiDui == 2 {
			cards = append(cards, laiziCard...)
			zu1 = this.analysis(cards)
			laiziCnt = 0
		}

		var duiziCount int
		for len(zu1) > 0 {
			if len(zu1[0].index) > 1 {
				zu1 = removeCard(zu1, 0)
				zu1 = removeCard(zu1, 0)
				duiziCount++
			} else {
				if laiziCnt > 0 {
					laiziZu = append(laiziZu, zu1[0].value)
					zu1 = removeCard(zu1, 0)
					laiziCnt--
					duiziCount++
				} else {
					item1 = insertCard(item1, getCard(zu1, 0))
					zu1 = removeCard(zu1, 0)
				}
			}
		}

		if len(item1) == 0 {
			//七对
			result.qidui = true
		} else {
			if duiziCount == 6 {
			}
		}
	}

	if result.qidui {
		result.Results = append(result.Results, HuHandResult{SpecailKind: MAHJONG_SEVENPAIRS})
		if config.OnlyChk {
			return result
		}
	}

	laiziCnt = len(laiziCard)
	var lzCount int
	var count int
	for k, _ := range zu {
		zu1 := cloneq(zu)
		lzCount = laiziCnt
		count = len(zu1[k].index)
		if count > 1 {
			//剪将
			v := zu1[k].value
			zu1 = removeCard(zu1, k)
			zu1 = removeCard(zu1, k)
			result.temp.Clear()
			result.temp.EyesKind = v
			result.temp.Eyes[0], result.temp.Eyes[1] = v, v
			this.analyseHu(zu1, lzCount, result, config)
			if config.OnlyChk && len(result.Results) != 0 {
				return result
			}
		}

		if laiziCnt > 0 {
			zu1 := cloneq(zu)
			lzCount = laiziCnt
			temp = zu1[k].value
			//剪将
			zu1 = removeCard(zu1, k)
			result.temp.Clear()
			result.temp.EyesKind = temp
			result.temp.Eyes[0], result.temp.Eyes[1] = temp, temp
			this.analyseHu(zu1, lzCount-1, result, config)
			if config.OnlyChk && len(result.Results) != 0 {
				return result
			}
		}
	}

	return result
}

func (this *MjAlgorithm) analyseHu(zu []cv, laiziCnt int, result *MJResult, config *MJSpecailConfig) bool {
	zuCnt := len(zu)
	if zuCnt == 0 {
		result.Results = append(result.Results, result.temp)
		return true
	}
	var ret bool
	//牌,牌,牌
	cache := cloneq(zu)
	v := cache[0].value
	if v < MJ_WIND_EAST {
		/*ret = this.analysePu1928(cache, laiziCnt, result, config)
		if !ret {*/
		if len(cache) >= 3 {
			v1 := cache[1].value - 1
			v2 := cache[2].value - 2
			if v1 == v && v2 == v && cache[2].value/9 == v/9 && cache[1].value/9 == v/9 {
				cache = removeCard(cache, 2)
				cache = removeCard(cache, 1)
				cache = removeCard(cache, 0)
				result.temp.AddSequence(v)
				ret = this.analyseHu(cache, laiziCnt, result, config)
				result.temp.Pop()
			}
		}
	}

	//牌,牌,癞子
	if laiziCnt > 0 {
		v := zu[0].value
		if v < MJ_WIND_EAST {
			if len(zu) >= 2 {
				cache := cloneq(zu)
				v1 := cache[1].value
				if v == v1-1 && v/9 == v1/9 {
					cache = removeCard(cache, 1)
					cache = removeCard(cache, 0)
					result.temp.AddCustom(v, v1, -1, ResultFlag_Tile)
					ret = this.analyseHu(cache, laiziCnt-1, result, config)
					result.temp.Pop()
				} else if v == v1-2 && v/9 == v1/9 {
					cache = removeCard(cache, 1)
					cache = removeCard(cache, 0)
					result.temp.AddCustom(v, v1-1, v1, ResultFlag_Tile)
					ret = this.analyseHu(cache, laiziCnt-1, result, config)
					result.temp.Pop()
				}
			}

			if len(zu) >= 3 {
				cache = cloneq(zu)
				v1 := cache[2].value
				if v == v1-2 && v/9 == v1/9 {
					cache = removeCard(cache, 2)
					cache = removeCard(cache, 0)
					result.temp.AddCustom(v, v+1, v1, ResultFlag_Tile)
					ret = this.analyseHu(cache, laiziCnt-1, result, config)
					result.temp.Pop()
				}
			}
		}
	}

	//牌,癞子,癞子
	if laiziCnt > 1 {
		v := zu[0].value
		if v < MJ_WIND_EAST {
			cache := cloneq(zu)
			cache = removeCard(cache, 0)
			result.temp.AddCustom(v, v, v, ResultFlag_Tile)
			ret = this.analyseHu(cache, laiziCnt-2, result, config)
			result.temp.Pop()
		}
	}

	//牌牌牌
	count := len(zu[0].index)
	if count >= 3 {
		cache := cloneq(zu)
		result.temp.AddTiles(cache[0].value)
		for i := 0; i < 3; i++ {
			cache = removeCard(cache, 0)
		}
		ret = this.analyseHu(cache, laiziCnt, result, config)
		result.temp.Pop()
	}

	//牌牌癞子(刻)
	if laiziCnt > 0 {
		count := len(zu[0].index)
		if count >= 2 {
			v := zu[0].value
			cache = removeCard(cache, 0)
			cache = removeCard(cache, 0)
			result.temp.AddTiles(v)
			ret = this.analyseHu(cache, laiziCnt-1, result, config)
			result.temp.Pop()
		}
	}

	//牌 癞子 癞子(刻)
	if laiziCnt > 1 {
		count := len(zu[0].index)
		if count >= 1 {
			v := zu[0].value
			result.temp.AddTiles(v)
			cache = removeCard(cache, 0)
			ret = this.analyseHu(cache, laiziCnt-2, result, config)
			result.temp.Pop()
		}
	}

	return ret
}

func (this *MjAlgorithm) Ting(mycard []int64, otherCard []int64, laizi []int64, config *MJSpecailConfig) map[int64][]int64 {
	var result map[int64][]int64
	for k, v := range mycard {
		if result != nil {
			if _, ok := result[v]; ok {
				continue
			}
		}
		var tmp []int64
		var cache []int64
		cache = append(cache, mycard[:k]...)
		cache = append(cache, mycard[k+1:]...)
		var card int64 = -1
		for i := int64(0); i < MJ_MAX; i++ {
			if this.Hu(cache, otherCard, i, laizi, config) != nil {
				card = v
				tmp = append(tmp, i)
			}
		}
		if card >= 0 {
			if result == nil {
				result = make(map[int64][]int64)
			}
			result[card] = tmp
		}
	}
	return result
}

func (this *MjAlgorithm) GetTingCards(pool []int64, cnt int, config *MJSpecailConfig, maxCnts int, excludeCard int32) (cards []int64) {
	var temp [MJ_MAX]int32
	for _, c := range pool {
		temp[c]++
	}
	config.OnlyChk = true
	/*for i := MJ_DOTS_1; i < MJ_MAX; i++ {
		if i != excludeCard {
			if temp[i] == 0 { //过滤掉不可能的牌
				k := i / 9
				n := i % 9
				if k < 3 {
					switch n {
					case 0:
						if !config.YaoJiuPu && (temp[i+1] == 0 || temp[i+2] == 0) {
							continue
						}
					case 1:
						if !config.ErBaPu && (temp[i-1] == 0 || temp[i+1] == 0) {
							continue
						}
					case 7:
						if !config.ErBaPu && (temp[i-1] == 0 || temp[i+1] == 0) {
							continue
						}
					case 8:
						if !config.YaoJiuPu && (temp[i-1] == 0 || temp[i-2] == 0) {
							continue
						}
					default:
						if (temp[i-1] == 0 && temp[i+1] == 0) && (temp[i-1] == 0 && temp[i-2] == 0) && (temp[i+1] == 0 && temp[i+2] == 0) {
							continue
						}
					}
				} else {
					if n < 4 {
						if !config.FengPu {
							continue
						}
					} else {
						if !config.JiangPu {
							continue
						}
					}
				}
			}
			result := this.Hu(pool, nil, i, nil, config)
			if len(result.Results) != 0 {
				cards = append(cards, i)
				if maxCnts > 0 && len(cards) == maxCnts {
					return
				}
			}
		}
	}*/
	return
}
