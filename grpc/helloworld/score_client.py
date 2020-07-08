# Copyright 2015 gRPC authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
"""The Python implementation of the GRPC helloworld.Greeter client."""

from __future__ import print_function
import logging

import grpc

from rpc_package import getscore_pb2
from rpc_package import getscore_pb2_grpc

"""
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
"""
def run():
    # NOTE(gRPC Python Team): .close() is possible on a channel and should be
    # used in circumstances in which the with statement does not fit the needs
    # of the code.
    with grpc.insecure_channel('localhost:50052') as channel:
        stub = getscore_pb2_grpc.HuScoreStub(channel)
        response = stub.CheckHuType(getscore_pb2.CardInfo( typeMapAll = {18:1,19:2,20:2,21:1,25:3,26:3,33:2},
                                                           tilesNum = 0,
                                                           discards=[16,17,18,19,20,21,21,22,22,22,22,23,23,23,23,24,24,26,27,27,27,27,28,28,28,28,29,29,29,29,30,30,31,32,33,33],
                                                           handCards=[18,19,19,20,20,21,25,25,25,26,26,26,33,33],
                                                           surplusCards =[21,24,24,25,30,30,31,31,31,32,32,32],
                                                           otherHands=[16,16,16,17,17,17,18,18,18,19,20,20,21],
                                                           chowCards=[],
                                                           pongCards=[],
                                                           mingKong=[],
                                                           anKong=[],
                                                           isBanker=True,
                                                           isZimo=True,
                                                           robKong=False,
                                                           lastOptIsKong=False,
                                                           lastCard_cpk=False,
                                                           lastCard=17,
                                                           tableFeng=0,
                                                           menFeng=1
                                                           ))
    print(response.message)

if __name__ == '__main__':
    logging.basicConfig()
    run()
