package mahjong

import (
	"time"
)

//麻将种类选项
const (
	Mahjong_ErRen       int = iota           //二人
	Mahjong_Max
)



//吃碰选项
const (
	ChowPongsOption_CAndP  int32 = iota //能吃能碰
	ChowPongsOption_OnlyP               //只能碰
	ChowPongsOption_OnlyC              //只能吃
	ChowPongsOption_None               //都不能
	ChowPongOption_Max
)


//二人麻将胡牌选项
const (
	ErRenWinOption_All   int32 = iota        //自摸抢杠
	ErRenWinOption_BySelfDrawn              //自摸胡
	ErRenWinOption_ByQiangGang              //抢杠
	ErRenWinOption_Max
)


//风牌选项
const (
	WindOption_None  int32 = iota             //不带风
	WindOption_Have //带风牌
	WindOption_Max
)

//坐庄选项
const (
	BankerOption_Winner int32 = iota //赢家坐庄
	BankerOption_Turn                //轮庄
	BankerOption_Max
)




//二人麻将参数信息
const (
	ErRenParam_TimeoutOption   int = iota           //超时选项
	ErRenParam_ChowPongsOption            //吃碰选项
	ErRenParam_WinOption                  //胡牌选项
	ErRenParam_WindOption                 //风牌选项
	ErRenParam_BaseScore                  //游戏底分（废弃，原钻石场选项
	ErRenParam_SameIPOption               //同IP不可进
	ErRenParam_HuangKongOption            //荒庄不荒杠选项
	ErRenParam_Max
)


//超时选项
var TimeoutOption = []time.Duration{time.Second * 30, time.Second * 60, time.Second * 15}

const DefaultWaitCPKHTimeout = time.Second * 24



//房卡场底分选项
var BaseRates = []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}


