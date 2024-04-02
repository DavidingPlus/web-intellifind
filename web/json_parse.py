# -*- coding: UTF-8 -*-

import json
from web.chat import gpt_35_api_stream
from web.reply import replyT


def jsonParse(jsonPath: str, weights: list) -> replyT:
    data = dict()
    score = 0.0
    allScores = [0.0] * 9
    proWeights = weights

    # with open as ... 不用考虑文件关闭， python 会自动再合适的时候关闭
    with open(jsonPath, mode="r", encoding="utf-8") as file:
        data = json.load(file)

    # 一. 遍历整个 json ，拿到数据
    # 跳出率较高
    stayTime: float = 0.0
    # 重复点击
    repeatClick: list[int] = [0, 0]
    # 页面打开慢
    pageLoad: float = 0.0
    pageLoadCnt: int = 0
    # 点击后网络反馈慢
    feedbackInterval: float = 0.0
    feedbackIntervalCnt: int = 0
    # 点击无反应
    clickNoReaction: list[int] = [0, 0]
    # 点击报错
    errorCount: int = 0
    # 页面加载报错
    consoleErrors: int = 0
    # 页面加载白屏
    isBlank: list[int] = [0, 0]

    for e in data['data']:
        if 'stayTime' in e['interactionAttr']:
            stayTime += e['interactionAttr']['stayTime']['value']
        if 'repeatClick' in e['interactionAttr']:
            if 'True' == e['interactionAttr']['repeatClick']['value']:
                repeatClick[0] += 1
            else:
                repeatClick[1] += 1
        if 'pageLoad' in e['performanceAttr']:
            pageLoad += e['performanceAttr']['pageLoad']['value']
            pageLoadCnt += 1
        if 'feedbackInterval' in e['performanceAttr']:
            feedbackInterval += e['performanceAttr']['feedbackInterval']['value']
            feedbackIntervalCnt += 1
        if 'clickNoReaction' in e['interactionAttr']:
            if 'True' == e['interactionAttr']['clickNoReaction']['value']:
                clickNoReaction[0] += 1
            else:
                clickNoReaction[1] += 1
        if 'errorCount' in e['interactionAttr']:
            errorCount += e['interactionAttr']['errorCount']['value']
        if 'consoleErrors' in e['performanceAttr']:
            consoleErrors += e['performanceAttr']['consoleErrors']['value']
        if 'isBlank' in e['interactionAttr']:
            if 'True' == e['interactionAttr']['isBlank']['value']:
                isBlank[0] += 1
            else:
                isBlank[1] += 1

    # 二. 计算得分
    # 1. 跳出率较高
    stayTime = round(stayTime, 2)
    if stayTime >= 30 * 1000:
        allScores[0] = 100
    elif stayTime >= 10 * 1000 and stayTime < 30 * 1000:
        allScores[0] = 80 + (100 - 80) / (30 * 1000 - 10 * 1000) * \
            (stayTime - 10 * 1000)
    elif stayTime >= 5 * 1000 and stayTime < 10 * 1000:
        allScores[0] = 70 + (80 - 70) / (10 * 1000 - 5 * 1000) * \
            (stayTime - 5 * 1000)
    else:
        allScores[0] = 60 + (70 - 60) / (5 * 1000 - 0 * 1000) * \
            (stayTime - 0 * 1000)

    allScores[0] = round(allScores[0], 2)  # 得分保留两位小数
    score += allScores[0] * proWeights[0]

    # 2. 重复点击
    allScores[1] = 60 + repeatClick[1] / \
        (repeatClick[0] + repeatClick[1]) * (100 - 60)

    allScores[1] = round(allScores[1], 2)
    score += allScores[1] * proWeights[1]

    # # 3. 页面打开慢
    # TODO 如果 json 里面没有 pageLoad 字段的处理（测试样例和其他的绝大多数情况不会）
    pageLoad = round(pageLoad / pageLoadCnt, 2)
    if pageLoad >= 0 * 1000 and pageLoad < 2 * 1000:
        allScores[2] = 100
    elif pageLoad >= 2 * 1000 and pageLoad < 5 * 1000:
        allScores[2] = 85 + (100 - 85) / (5 * 1000 - 2 *
                                          1000) * (pageLoad - 2 * 1000)
    else:
        allScores[2] = 85 - 3 * (pageLoad - 5 * 1000)
        if (allScores[2] < 60):
            allScores[2] = 60

    allScores[2] = round(allScores[2], 2)
    score += allScores[2] * proWeights[2]

    # 4. 点击后网络反馈慢
    feedbackInterval = round(feedbackInterval / feedbackIntervalCnt, 2)
    if feedbackInterval >= 0 and feedbackInterval < 60:
        allScores[3] = 100
    elif feedbackInterval >= 60 and feedbackInterval < 100:
        allScores[3] = 90 + \
            (100 - 90) / (100 - 60) * (feedbackInterval - 60)
    elif feedbackInterval >= 100 and feedbackInterval < 200:
        allScores[3] = 80 + \
            (90 - 80) / (200 - 100) * (feedbackInterval - 100)
    elif feedbackInterval >= 200 and feedbackInterval < 500:
        allScores[3] = 60 + \
            (80 - 60) / (500 - 200) * (feedbackInterval - 200)
    else:
        allScores[3] = 60

    allScores[3] = round(allScores[3], 2)
    score += allScores[3] * proWeights[3]

    # 5. 点击无反应
    allScores[4] = 60 + clickNoReaction[1] / \
        (clickNoReaction[0] + clickNoReaction[1]) * (100 - 60)

    allScores[4] = round(allScores[4], 2)
    score += allScores[4] * proWeights[4]

    # 6. 点击报错
    allScores[5] = 100 - 2 * errorCount
    if allScores[5] < 60:
        allScores[5] = 60

    allScores[5] = round(allScores[5], 2)
    score += allScores[5] * proWeights[5]

    # 7. 页面加载报错
    allScores[6] = 100 - 2 * consoleErrors
    if allScores[6] < 60:
        allScores[6] = 60

    allScores[6] = round(allScores[6], 2)
    score += allScores[6] * proWeights[6]

    # 8. 页面加载白屏
    allScores[7] = 60 + isBlank[1] / \
        (isBlank[0] + isBlank[1]) * (100 - 60)

    allScores[7] = round(allScores[7], 2)
    score += allScores[7] * proWeights[7]

    # 9. 多个同时出现
    allScores[8] = 100
    for i in range(8):
        if 100 != allScores[i]:
            allScores[8] -= 5

    allScores[8] = round(allScores[8], 2)
    score += allScores[8] * proWeights[8]

    score = round(score, 2)  # 总分同样保留两位小数

    # 三. 问题简述
    briefDesc = "经过检测，您本次体验存在的问题如下，以下是问题简述：\n"
    briefDesc += f"1. 关于跳出率较高，总共停留时长为 {stayTime} 毫秒，得分 {allScores[0]} 分；\n"
    briefDesc += f"2. 关于重复点击，重复点击次数为 {repeatClick[0]} 次，未重复点击次数为 {repeatClick[1]} 次，得分为 {allScores[1]} 分；\n"
    briefDesc += f"3. 关于页面打开慢，页面平均加载时长为 {pageLoad} 毫秒，得分 {allScores[2]} 分；\n"
    briefDesc += f"4. 关于点击后网络反馈慢，页面平均延迟时间为 {feedbackInterval} 毫秒，得分 {allScores[3]} 分；\n"
    briefDesc += f"5. 关于点击无反应，点击无反应次数为 {clickNoReaction[0]} 次，点击响应正常次数为 {clickNoReaction[1]} 次，得分 {allScores[4]} 分；\n"
    briefDesc += f"6. 关于点击报错，页面错误信息总数为 {errorCount}，得分 {allScores[5]} 分；\n"
    briefDesc += f"7. 关于页面加载报错，控制台的报错信息总数为 {consoleErrors}，得分 {allScores[6]} 分；\n"
    briefDesc += f"8. 关于页面加载白屏，白屏次数为 {isBlank[0]} 次，未白屏次数为 {isBlank[1]} 次，得分 {allScores[7]} 分；\n"
    briefDesc += f"9. 关于多个同时出现，上述问题总共出现 {(100 - allScores[8]) // 5} 个，得分 {allScores[8]} 分。\n"

    # 四. 询问 GPT ，问题详述
    model = "gpt-3.5-turbo"
    messages = [
        {
            'role': 'system', 'content': '你是一个网站体验分析师。现在我给你一些网站体验过程中的问题简述，并且附有一些数据和得分。请结合数据和得分详细分析一下这些问题，并且给出合适的优化建议，每一条请严格按照我的格式提行进行补充。请使用第三人称进行回答。'
        },
        {
            'role': 'user', 'content': briefDesc
        },]
    detailDesc = gpt_35_api_stream(model=model, messages=messages)

    # 四. 封装为结构体返回
    return replyT(score, allScores, briefDesc, detailDesc)
