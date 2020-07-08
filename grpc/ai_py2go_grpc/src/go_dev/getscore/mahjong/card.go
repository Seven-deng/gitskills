package mahjong

import (
	"math"
	"math/rand"
	"sort"
	"time"
)

const (
	/*0~8(1到9筒)*/
	MJ_DOTS_1 int64 = iota
	MJ_DOTS_2
	MJ_DOTS_3
	MJ_DOTS_4
	MJ_DOTS_5
	MJ_DOTS_6
	MJ_DOTS_7
	MJ_DOTS_8
	MJ_DOTS_9
	/*9~17(1到9索)*/
	MJ_BAMBOO_1
	MJ_BAMBOO_2
	MJ_BAMBOO_3
	MJ_BAMBOO_4
	MJ_BAMBOO_5
	MJ_BAMBOO_6
	MJ_BAMBOO_7
	MJ_BAMBOO_8
	MJ_BAMBOO_9
	/*18~26(1到9万)*/
	MJ_CHARACTERS_1
	MJ_CHARACTERS_2
	MJ_CHARACTERS_3
	MJ_CHARACTERS_4
	MJ_CHARACTERS_5
	MJ_CHARACTERS_6
	MJ_CHARACTERS_7
	MJ_CHARACTERS_8
	MJ_CHARACTERS_9
	/*27~30(东南西北风)*/
	MJ_WIND_EAST
	MJ_WIND_SOUTH
	MJ_WIND_WEST
	MJ_WIND_NORTH
	/*31~33(中发白)*/
	MJ_DRAGON_RED
	MJ_DRAGON_GREEN
	MJ_DRAGON_WHITE
	/*最大34*/
	MJ_MAX
)

const (
	MJ_Four int = iota
	MJ_Three
	MJ_Double
	MJ_Link
	MJ_Random2
	MJ_Random1
	MJ_Style_Max
)

const (
	MJ_DOTS int = iota
	MJ_BAMBOO
	MJ_CHARACTERS
	MJ_Word
)

func IsSameType(c1, c2 int64) bool {
	if c1 >= MJ_CHARACTERS_1 && c1 <= MJ_CHARACTERS_9 && c2 >= MJ_CHARACTERS_1 && c2 <= MJ_CHARACTERS_9 {
		return true
	}
	return false
}

func IsIsolated(pool []int64, card int64) bool {

	if pool[card-1] != 0 || pool[card+1] != 0 || pool[card-2] != 0 || pool[card+2] != 0 {
		return false
	} else {
		return true
	}

	return false
}

func NoNeighbor(pool []int64, card int64) bool {

	if pool[card-1] == 0 && pool[card+1] == 0 {
		return true
	}

	return false
}

//对这张牌进行评分，分数越高牌越好
func CalcuCardScore(pool []int64, card int64) int {
	switch pool[card] { //刚好缺张
	case 0:
		if card >= MJ_WIND_EAST && card <= MJ_DRAGON_WHITE {
			return 0
		} else {
			var color int64 = card / 9
			var num int64 = card % 9
			switch num {
			case 0, 8: //首尾
				var n, nn int64
				if num == 0 {
					n = color*9 + num + 1
					nn = color*9 + num + 2
				} else {
					n = color*9 + num - 1
					nn = color*9 + num - 2
				}
				switch pool[n] {
				case 0: //00x
					return 0
				case 1:
					switch pool[nn] {
					case 1: //011
						return 90
					case 2: //012
						return 80
					case 3: //013
						return 85
					case 4: //014
						return 90
					case 0: //010
						return 0
					}
				case 2:
					switch pool[nn] {
					case 1: //021
						return 80
					case 2: //022
						return 70
					case 3: //023
						return 60
					case 4: //024
						return 0
					case 0: //020
						return 0
					}
				case 3:
					switch pool[nn] {
					case 1: //031
						return 80
					case 2: //032
						return 30
					case 3: //033
						return 0
					case 4: //034
						return 0
					case 0: //030
						return 0
					}
				case 4:
					switch pool[nn] {
					case 1: //041
						return 5
					case 2: //042
						return 0
					case 3: //043
						return 0
					case 4: //044
						return 0
					case 0: //040
						return 0
					}
				}
			default: //中间
				min := color * 9
				max := color*9 + 8
				pre := color*9 + num - 1
				ppre := pre - 1
				next := color*9 + num + 1
				nnext := next + 1
				if (pool[pre] == 1 && pool[next] == 1) ||
					(nnext <= max && pool[next] == 1 && pool[nnext] == 1) ||
					(ppre >= min && pool[ppre] == 1 && pool[pre] == 1) {
					return 95
				}
				if (pool[pre] == 2 && pool[next] == 2) ||
					(nnext <= max && pool[next] == 2 && pool[nnext] == 2) ||
					(ppre >= min && pool[ppre] == 2 && pool[pre] == 2) {
					return 40
				}
				if pool[pre] >= 1 || pool[next] >= 1 {
					return 50
				}
				return 0
			}
		}
	case 1:
		if card >= MJ_WIND_EAST && card <= MJ_DRAGON_WHITE {
			return 80
		} else {
			var color int64 = card / 9
			var num int64 = card % 9
			switch num {
			case 0, 8: //首尾
				var n, nn int64
				if num == 0 {
					n = color*9 + num + 1
					nn = color*9 + num + 2
				} else {
					n = color*9 + num - 1
					nn = color*9 + num - 2
				}
				switch pool[n] {
				case 0: //10x
					return 75
				case 1:
					switch pool[nn] {
					case 1: //111
						return 0
					case 2: //112
						return 10
					case 3: //113
						return 15
					case 4: //114
						return 10
					case 0: //110
						return 20
					}
				case 2:
					switch pool[nn] {
					case 1: //121
						return 40
					case 2: //122
						return 93
					case 3: //123
						return 85
					case 4: //124
						return 70
					case 0: //120
						return 50
					}
				case 3:
					switch pool[nn] {
					case 1: //131
						return 20
					case 2: //132
						return 85
					case 3: //133
						return 80
					case 4: //134
						return 70
					case 0: //130
						return 60
					}
				case 4:
					switch pool[nn] {
					case 0: //140
						return 60
					case 1: //141
						return 65
					case 2: //142
						return 60
					case 3: //143
						return 65
					case 4: //144
						return 70
					}
				}
			default: //中间
				min := color * 9
				max := color*9 + 8
				pre := color*9 + num - 1
				ppre := pre - 1
				next := color*9 + num + 1
				nnext := next + 1
				if (pool[pre] == 1 && pool[next] == 1) ||
					(nnext <= max && pool[next] == 1 && pool[nnext] == 1) ||
					(ppre >= min && pool[ppre] == 1 && pool[pre] == 1) {
					return 95
				}
				if (pool[pre] == 2 && pool[next] == 2) ||
					(nnext <= max && pool[next] == 2 && pool[nnext] == 2) ||
					(ppre >= min && pool[ppre] == 2 && pool[pre] == 2) {
					return 40
				}
				if pool[pre] >= 1 || pool[next] >= 1 {
					return 50
				}
				return 0
			}
		}
	case 2:
		return 90
	case 3:
		return 95
	}

	return 0
}

type CardPool struct {
	buf      []int64
	pos      int
	cnt      int
	HasWinds bool
	HasWord  bool
	HasRed   bool
}

func NewCardPool(hasWinds bool) *CardPool {
	cp := &CardPool{}
	cp.cnt = 36
	if hasWinds {
		cp.cnt += 28
	}
	cp.buf = make([]int64, 0, cp.cnt)
	cp.HasWinds = hasWinds
	cp.HasWord = hasWinds
	cp.HasRed = true
	cp.Reset()
	return cp
}

func PDSDHNewCardPool(hasWinds bool) *CardPool {
	cp := &CardPool{}
	cp.cnt = 36
	if hasWinds {
		cp.cnt += 28
	}
	cp.buf = make([]int64, 0, cp.cnt)
	cp.HasWinds = hasWinds
	cp.HasWord = hasWinds
	cp.HasRed = true
	cp.PDSDHReset()
	return cp
}

func NewCardPoolHasRed(hasRed bool) *CardPool {
	cp := &CardPool{}
	cp.cnt = 36
	cp.HasWinds = true
	cp.HasWord = true
	cp.HasRed = hasRed
	if hasRed {
		cp.cnt += 28
	}
	cp.buf = make([]int64, 0, cp.cnt)
	cp.Reset()
	return cp
}

func NewCardPoolEx(hasWinds, hasWord bool) *CardPool {
	cp := &CardPool{}
	cp.cnt = 36
	if hasWinds {
		cp.cnt += 16
	}
	if hasWord {
		cp.cnt += 12
	}
	cp.buf = make([]int64, 0, cp.cnt)
	cp.HasWinds = hasWinds
	cp.HasWord = hasWord
	cp.HasRed = true
	cp.Reset()
	return cp
}
func NewBloodCardPoolEx() *CardPool {
	cp := &CardPool{}
	cp.buf = make([]int64, 0, 0)
	for i := MJ_DOTS_1; i <= MJ_DOTS_9; i++ {
		for j := 0; j < 4; j++ {
			cp.buf = append(cp.buf, i)
		}
	}
	for i := MJ_BAMBOO_1; i <= MJ_BAMBOO_9; i++ {
		for j := 0; j < 4; j++ {
			cp.buf = append(cp.buf, i)
		}
	}
	for i := MJ_CHARACTERS_1; i <= MJ_CHARACTERS_9; i++ {
		for j := 0; j < 4; j++ {
			cp.buf = append(cp.buf, i)
		}
	}
	cp.Shuffle()
	cp.cnt = len(cp.buf)
	return cp
}
func (cp *CardPool) CheatCard(oldc, newc int64) bool {
	for k, v := range cp.buf {
		if k >= cp.pos && v == newc {
			cp.buf[k] = oldc
			return true
		}
	}
	return false
}
func (cp *CardPool) Reset() {
	cp.buf = cp.buf[0:0]
	for i := MJ_CHARACTERS_1; i <= MJ_CHARACTERS_9; i++ {
		for j := 0; j < 4; j++ {
			cp.buf = append(cp.buf, i)
		}
	}
	if cp.HasWinds {
		for i := MJ_WIND_EAST; i <= MJ_WIND_NORTH; i++ {
			for j := 0; j < 4; j++ {
				cp.buf = append(cp.buf, i)
			}
		}
	}
	if cp.HasWord {
		for i := MJ_DRAGON_RED; i < MJ_MAX; i++ {
			for j := 0; j < 4; j++ {
				cp.buf = append(cp.buf, i)
			}
		}
	} else if cp.HasRed {
		for j := 0; j < 4; j++ {
			cp.buf = append(cp.buf, MJ_DRAGON_RED)
		}
	}

	cp.Shuffle()
}

func (cp *CardPool) PDSDHReset() {
	cp.buf = cp.buf[0:0]
	for i := MJ_CHARACTERS_1; i <= MJ_CHARACTERS_9; i++ {
		for j := 0; j < 4; j++ {
			cp.buf = append(cp.buf, i)
		}
	}
	if cp.HasWinds {
		for i := MJ_WIND_EAST; i <= MJ_WIND_NORTH; i++ {
			for j := 0; j < 4; j++ {
				cp.buf = append(cp.buf, i)
			}
		}
	}
	if cp.HasWord {
		for i := MJ_DRAGON_RED; i < MJ_MAX; i++ {
			for j := 0; j < 4; j++ {
				cp.buf = append(cp.buf, i)
			}
		}
	}

	cp.Shuffle()
}

func (cp *CardPool) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	cnt := len(cp.buf)
	for i := 0; i < cnt; i++ {
		j := rand.Intn(i + 1)
		cp.buf[i], cp.buf[j] = cp.buf[j], cp.buf[i]
	}
	cp.pos = 0
}

func (cp *CardPool) ShuffleLastCard() {
	cnt := len(cp.buf)
	for i := cp.pos; i < cnt; i++ {
		j := RandInt(cp.pos, i+1)
		cp.buf[i], cp.buf[j] = cp.buf[j], cp.buf[i]
	}
}

//[l..u)
func RandInt(args ...int) int {
	switch len(args) {
	case 0:
		return rand.Int()
	case 1:
		if args[0] > 0 {
			return rand.Intn(args[0])
		} else {
			return 0
		}
	default:
		l := args[0]
		u := args[1]
		switch {
		case l == u:
			{
				return l
			}
		case l > u:
			{
				return u + rand.Intn(l-u)
			}
		default:
			{
				return l + rand.Intn(u-l)
			}
		}
	}
}
func (cp *CardPool) Next() int64 {
	if cp.pos >= len(cp.buf) {
		return -1
	}
	c := cp.buf[cp.pos]
	cp.pos++
	return c
}
func (cp *CardPool) GetCardWall() []int64 {
	return cp.buf[cp.pos:]
}
func (cp *CardPool) GetNoLackColor(lackColor int64) int64 {
	if cp.buf[cp.pos]/9 != lackColor {
		return cp.Next()
	}
	for k, v := range cp.buf {
		if k < cp.pos {
			continue
		}
		if v/9 != lackColor {
			cp.buf[cp.pos], cp.buf[k] = cp.buf[k], cp.buf[cp.pos]
			return cp.Next()
		}
	}
	return -1
}
func (cp *CardPool) GetLackColor(lackColor int64) int64 {
	if cp.buf[cp.pos]/9 == lackColor {
		return cp.Next()
	}
	for k, v := range cp.buf {
		if k < cp.pos {
			continue
		}
		if v/9 == lackColor {
			cp.buf[cp.pos], cp.buf[k] = cp.buf[k], cp.buf[cp.pos]
			return cp.Next()
		}
	}
	return -1
}

func (cp *CardPool) GetHuNextCard(cards []int64) int64 {
	if cp.pos >= len(cp.buf) {
		return -1
	}
	c := cp.buf[cp.pos]
	tingMap := make(map[int64]int)
	for _, nc := range cards {
		tingMap[nc]++
	}
	if _, ok := tingMap[c]; !ok {
		needcardInx := -1
		for k, v := range cp.buf {
			if k < cp.pos {
				continue
			}
			if _, ok1 := tingMap[v]; ok1 {
				needcardInx = k
				break
			}
		}
		if needcardInx != -1 {
			cp.buf[cp.pos], cp.buf[needcardInx] = cp.buf[needcardInx], cp.buf[cp.pos]
			return cp.Next()
		}
	}
	return -1
}

func (cp *CardPool) GetNextNoIn(cards []int64) int64 {
	if cp.pos >= len(cp.buf) {
		return -1
	}
	c := cp.buf[cp.pos]
	tingMap := make(map[int64]int)
	for _, nc := range cards {
		tingMap[nc]++
	}
	if _, ok := tingMap[c]; ok {
		needcardInx := -1
		for k, v := range cp.buf {
			if k < cp.pos {
				continue
			}
			if _, ok1 := tingMap[v]; !ok1 {
				needcardInx = k
				break
			}
		}
		if needcardInx != -1 {
			cp.buf[cp.pos], cp.buf[needcardInx] = cp.buf[needcardInx], cp.buf[cp.pos]
			return cp.Next()
		}
	} else {
		return cp.Next()
	}
	return -1
}
func (cp *CardPool) GetCardByStyle(style int, cardType int) []int64 {
	hands := make([]int64, 0)
	switch style {
	case MJ_Four:
		cards := cp.GetFourCard(cardType)
		hands = append(hands, cards...)
	case MJ_Three:
		cards := cp.GetThreeCard(cardType)
		hands = append(hands, cards...)
	case MJ_Double:
		cards := cp.GetDoubleCard(cardType)
		hands = append(hands, cards...)
	case MJ_Link:
		cards := cp.GetLinkCard(cardType)
		hands = append(hands, cards...)
	case MJ_Random1:
		cards := cp.GetRandom1Card(cardType)
		hands = append(hands, cards...)
	case MJ_Random2:
		cards := cp.GetRandom2Card(cardType)
		hands = append(hands, cards...)
	}

	return hands
}

func (cp *CardPool) GetFourCard(cardType int) []int64 {
	hands := make([]int64, 0)
	maxNum := int64(0)
	minNum := int64(0)
	if cardType == MJ_Word {
		maxNum = MJ_DRAGON_WHITE
		minNum = MJ_WIND_EAST
	} else {
		maxNum = int64((cardType+1)*9) - 1
		minNum = int64((cardType) * 9)
	}

	indexArr := make([]int, 0)
	first := int64(-1)
	for pos := cp.pos; pos < len(cp.buf); pos++ {
		card := cp.buf[pos]
		if card <= maxNum && card >= minNum {
			if first == -1 {
				first = card
				indexArr = make([]int, 0)
				indexArr = append(indexArr, pos)
				for findIndex := pos + 1; findIndex < len(cp.buf); findIndex++ {
					if cp.buf[findIndex] == first {
						indexArr = append(indexArr, findIndex)
					}

					if len(indexArr) == 4 {
						//需要删除元素
						for _, v := range indexArr {
							cp.buf[cp.pos], cp.buf[v] = cp.buf[v], cp.buf[cp.pos]
							cp.pos += 1
							hands = append(hands, first)
						}

						return hands
					}
				}
				first = -1
			}
		}
	}
	return nil
}

func (cp *CardPool) GetThreeCard(cardType int) []int64 {
	hands := make([]int64, 0)
	maxNum := int64(0)
	minNum := int64(0)
	if cardType == MJ_Word {
		maxNum = MJ_DRAGON_WHITE
		minNum = MJ_WIND_EAST
	} else {
		maxNum = int64((cardType+1)*9) - 1
		minNum = int64((cardType) * 9)
	}

	indexArr := make([]int, 0)
	first := int64(-1)
	for pos := cp.pos; pos < len(cp.buf); pos++ {
		card := cp.buf[pos]
		if card <= maxNum && card >= minNum {
			if first == -1 {
				first = card
				indexArr = make([]int, 0)
				indexArr = append(indexArr, pos)
				for findIndex := pos + 1; findIndex < len(cp.buf); findIndex++ {
					if cp.buf[findIndex] == first {
						indexArr = append(indexArr, findIndex)
					}

					if len(indexArr) == 3 {
						//需要删除元素
						for _, v := range indexArr {
							cp.buf[cp.pos], cp.buf[v] = cp.buf[v], cp.buf[cp.pos]
							cp.pos += 1
							hands = append(hands, first)
						}

						return hands
					}
				}
				first = -1
			}
		}
	}
	return nil
}

func (cp *CardPool) GetDoubleCard(cardType int) []int64 {
	hands := make([]int64, 0)
	maxNum := int64(0)
	minNum := int64(0)
	if cardType == MJ_Word {
		maxNum = MJ_DRAGON_WHITE
		minNum = MJ_WIND_EAST
	} else {
		maxNum = int64((cardType+1)*9) - 1
		minNum = int64((cardType) * 9)
	}

	indexArr := make([]int, 0)
	first := int64(-1)
	for pos := cp.pos; pos < len(cp.buf); pos++ {
		card := cp.buf[pos]
		if card <= maxNum && card >= minNum {
			if first == -1 {
				first = card
				indexArr = make([]int, 0)
				indexArr = append(indexArr, pos)
				for findIndex := pos + 1; findIndex < len(cp.buf); findIndex++ {
					if cp.buf[findIndex] == first {
						indexArr = append(indexArr, findIndex)
					}

					if len(indexArr) == 2 {
						//需要删除元素
						for _, v := range indexArr {
							cp.buf[cp.pos], cp.buf[v] = cp.buf[v], cp.buf[cp.pos]
							cp.pos += 1
							hands = append(hands, first)
						}

						return hands
					}
				}
				first = -1
			}
		}
	}
	return nil
}

func (cp *CardPool) GetLinkCard(cardType int) []int64 {
	hands := make([]int64, 0)
	maxNum := int64(0)
	minNum := int64(0)
	if cardType == MJ_Word {
		maxNum = MJ_DRAGON_WHITE
		minNum = MJ_WIND_EAST
	} else {
		maxNum = int64((cardType+1)*9) - 1
		minNum = int64((cardType) * 9)
	}

	first := int64(-1)
	findArr := [5]int{-1, -1, -1, -1, -1}
	findNum := 0
	for pos := cp.pos; pos < len(cp.buf); pos++ {
		card := cp.buf[pos]
		if card <= maxNum && card >= minNum {
			if first == -1 {
				first = card
				findArr = [5]int{-1, -1, -1, -1, -1}
				findArr[2] = pos
				findNum = 1
				for findIndex := pos + 1; findIndex < len(cp.buf); findIndex++ {
					findCard := cp.buf[findIndex]
					if findCard <= maxNum && findCard >= minNum {
						offset := findCard - first
						if math.Abs(float64(offset)) <= 2 {
							offset = offset + 2
							if findArr[offset] == -1 {
								findArr[offset] = findIndex
								findNum += 1
							}
						}
					}

					//找到连续的3个值
					firstIndex := -1
					continueNum := 0
					if findNum >= 3 {
						for loop := 0; loop < len(findArr); loop++ {
							if findArr[loop] != -1 {
								if firstIndex == -1 {
									firstIndex = loop
								}

								continueNum += 1
								if continueNum >= 3 {
									//需要针对位置重新排序，否则会出现处理过的的换后边的问题
									sortArry := []int{}
									for i := 0; i < 3; i++ {
										sortArry = append(sortArry, findArr[firstIndex+i])
									}
									sort.Ints(sortArry)
									for i := 0; i < 3; i++ {
										newPos := sortArry[i]
										hands = append(hands, cp.buf[newPos])
										cp.buf[cp.pos], cp.buf[newPos] = cp.buf[newPos], cp.buf[cp.pos]
										cp.pos += 1
									}

									return hands
								}
							} else {
								firstIndex = -1
								continueNum = 0
							}
						}
					}
				}
				first = -1
			}
		}
	}
	return nil
}

func (cp *CardPool) GetRandom2Card(cardType int) []int64 {
	hands := make([]int64, 0)
	maxNum := int64(0)
	minNum := int64(0)
	if cardType == MJ_Word {
		maxNum = MJ_DRAGON_WHITE
		minNum = MJ_WIND_EAST
	} else {
		maxNum = int64((cardType+1)*9) - 1
		minNum = int64((cardType) * 9)
	}

	indexArr := make([]int, 0)
	first := int64(-1)
	for pos := cp.pos; pos < len(cp.buf); pos++ {
		card := cp.buf[pos]
		if card <= maxNum && card >= minNum {
			if first == -1 {
				first = card
				indexArr = make([]int, 0)
				indexArr = append(indexArr, pos)
				for findIndex := pos + 1; findIndex < len(cp.buf); findIndex++ {
					if cp.buf[findIndex] <= maxNum && cp.buf[findIndex] >= minNum {
						indexArr = append(indexArr, findIndex)
					}

					if len(indexArr) == 2 {
						//需要删除元素
						for _, v := range indexArr {
							hands = append(hands, cp.buf[v])
							cp.buf[cp.pos], cp.buf[v] = cp.buf[v], cp.buf[cp.pos]
							cp.pos += 1
						}

						return hands
					}
				}
				first = -1
			}
		}
	}
	return nil
}

func (cp *CardPool) GetRandom1Card(cardType int) []int64 {
	hands := make([]int64, 0)
	maxNum := int64(0)
	minNum := int64(0)
	if cardType == MJ_Word {
		maxNum = MJ_DRAGON_WHITE
		minNum = MJ_WIND_EAST
	} else {
		maxNum = int64((cardType+1)*9) - 1
		minNum = int64((cardType) * 9)
	}

	indexArr := make([]int, 0)
	first := int64(-1)
	for pos := cp.pos; pos < len(cp.buf); pos++ {
		card := cp.buf[pos]
		if card <= maxNum && card >= minNum {
			if first == -1 {
				first = card
				indexArr = make([]int, 0)
				indexArr = append(indexArr, pos)
				//需要删除元素
				for _, v := range indexArr {
					cp.buf[cp.pos], cp.buf[v] = cp.buf[v], cp.buf[cp.pos]
					cp.pos += 1
					hands = append(hands, first)
				}
				return hands
			}
		}
	}
	return nil
}

func (cp *CardPool) SendHandCard() []int64 {
	hands := make([]int64, 0)
	for i := 0; i < 13; i++ {
		hands = append(hands, cp.Next())
	}

	return hands
}

func (cp *CardPool) TryNextN(n int) int64 {
	if cp.pos+n >= len(cp.buf) {
		return -1
	}
	return cp.buf[cp.pos+n]
}

func (cp *CardPool) ChangeNextN(n int, c int64) bool {
	if cp.pos+n >= len(cp.buf) {
		return false
	}
	cp.buf[cp.pos+n] = c
	return true
}

func (cp *CardPool) GetIndex(tingCards []int64) (int, int64) {
	for i := cp.pos; i < len(cp.buf); i++ {
		for _, card := range tingCards {
			if cp.buf[i] == card {
				return i - cp.pos, card
			}
		}
	}
	return -1, -1
}

func (cp *CardPool) GetNoInCludIndex(tingCards []int64) (int, int64) {
	for i := cp.pos; i < len(cp.buf); i++ {
		state := true
		for _, card := range tingCards {
			if cp.buf[i] == card {
				state = false
				continue
			}
		}
		if state {
			return i - cp.pos, cp.buf[i]
		} else {
			continue
		}
	}
	return -1, -1
}

//n从0开始,即n=0取最后一个元素
func (cp *CardPool) Back(n int) int64 {
	cnt := len(cp.buf)
	if cnt <= n {
		return -1
	}
	return cp.buf[cnt-n-1]
}

//n从0开始,即n=0取最后一个元素
func (cp *CardPool) PopBackNth(n int) int64 {
	if cp.Count() <= 0 {
		return -1
	}
	cnt := len(cp.buf)
	if n > cnt {
		return -1
	}
	ret := cp.buf[cnt-n-1]
	for i := cnt - n - 1; i < cnt-1; i++ {
		cp.buf[i] = cp.buf[i+1]
	}
	cp.buf = cp.buf[:cnt-1]
	return ret
}

func (cp *CardPool) Count() int {
	if len(cp.buf) >= cp.pos {
		cnt := len(cp.buf) - cp.pos
		if cnt < 0 {
			cnt = 0
		}
		return cnt
	}
	return 0
}

func (cp *CardPool) Pos() int {
	return cp.pos
}

func (cp *CardPool) GetTotalCount() int {
	return cp.cnt
}

func (cp *CardPool) Buf() []int64 {
	return cp.buf
}

/**
hhx 发牌器专用开始==============
*/

func (cp *CardPool) ChangeBuf(bufss []int64) {

	cp.buf = bufss
}

/**
hhx 发牌器专用结束===================
*/
func (cp *CardPool) CheckCardInPool(card int64) int {
	index := cp.pos
	for ; index < len(cp.buf); index++ {
		poolCard := cp.buf[index]
		if poolCard == card {
			return index
		}
	}
	return -1
}

func (cp *CardPool) ChangeDesignatedCard(card int64, cardIndex int) bool {
	if -1 == cardIndex || cardIndex < cp.pos || cardIndex > len(cp.buf) {
		return false
	}
	cp.buf[cardIndex] = card
	return true
}

//获取手牌的积分
func (cp *CardPool) TotalHandCardsScore(hands []int64) int {
	var typePool [MJ_MAX]int64
	for _, card := range hands {
		typePool[card]++
	}
	score := 0
	for i := MJ_CHARACTERS_1; i < MJ_MAX; i++ {
		switch typePool[i] {
		case 2:
			score += 5
		case 3:
			score += 20
		case 4:
			score += 30
		case 1:
			if i > MJ_CHARACTERS_9 {
				score++
			} else {
				if i+2 <= MJ_CHARACTERS_9 && typePool[i+1] == 1 && typePool[i+2] == 1 {
					score += 10
					i = i + 3
				} else {
					score++
				}
			}
		}
	}
	return score
}
