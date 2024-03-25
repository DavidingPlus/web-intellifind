
from chat import gpt_35_api_stream
from web.reply import replyT
import json
from weight import defaultWeight

def jsonParse(jsonPath: str, weights: list) -> replyT:
    data = dict()
    score = 0.0
    allScores = [0.0] * 9
    proWeights = []
    if [] == weights:
        proWeights = defaultWeight
    else:
        proWeights = weights

    # with open as ... 不用考虑文件关闭， python 会自动再合适的时候关闭
    with open(jsonPath, mode="r", encoding="utf-8") as file:
        data = json.load(file)

    # 一. 计算得分

    # 1. 跳出率较高
    stayTime = data['interactionAttr']['stayTime']['value']
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

    score += allScores[0] * proWeights[0]

    # 2. 重复点击
    if str("False") == data['interactionAttr']['repeatClick']['value']:
        allScores[1] = 100
    else:
        allScores[1] = 60

    score += allScores[1] * proWeights[1]

    # 3. 页面打开慢
    pageLoad = data['performanceAttr']['pageLoad']['value']
    if pageLoad >= 0 * 1000 and pageLoad < 2 * 1000:
        allScores[2] = 100
    elif pageLoad >= 2 * 1000 and pageLoad < 5 * 1000:
        allScores[2] = 85 + (100 - 85) / (5 * 1000 - 2 *
                                          1000) * (pageLoad - 2 * 1000)
    else:
        allScores[2] = 85 - 3 * (pageLoad - 5 * 1000)
        if (allScores[2] < 60):
            allScores[2] = 60

    score += allScores[2] * proWeights[2]

    # 4. 点击后网络反馈慢
    feedbackInterval = data['performanceAttr']['feedbackInterval']['value']
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

    score += allScores[3] * proWeights[3]

    # 5. 点击无反应
    allScores[4] = allScores[2] * \
        0.5 + allScores[3] * 0.5

    score += allScores[4] * proWeights[4]

    # 6. 点击报错
    allScores[5] = 100 - 3 * data['interactionAttr']['errorCount']['value']
    if allScores[5] < 60:
        allScores[5] = 60

    score += allScores[5] * proWeights[5]

    # 7. 页面加载报错
    allScores[6] = 100 - 3 * data['performanceAttr']['consoleErrors']['value']
    if allScores[6] < 60:
        allScores[6] = 60

    score += allScores[6] * proWeights[6]

    # 8. 页面加载白屏
    if str("False") == data['interactionAttr']['isBlank']['value']:
        allScores[7] = 100
    else:
        allScores[7] = 60

    score += allScores[7] * proWeights[7]

    # 9. 多个同时出现
    allScores[8] = 100
    for i in range(8):
        if 100 != allScores[i]:
            allScores[8] -= 5

    score += allScores[8] * proWeights[8]

    score = round(score, 2)  # 保留两位小数

    # 二. 问题简述
    briefDesc = "经过检测，您本次体验存在的问题如下，以下是问题简述：\n"
    briefDesc += f"1. 关于跳出率较高，停留时长为 {stayTime} 毫秒，得分 {round(allScores[0] , 2)} 分；\n"
    briefDesc += f"2. 关于重复点击，是否重复点击为 {data['interactionAttr']['repeatClick']['value']} ，得分为 {round(allScores[1] , 2)} 分；\n"
    briefDesc += f"3. 关于页面打开慢，页面加载时长为 {pageLoad} 毫秒，得分 {round(allScores[2] , 2)} 分；\n"
    briefDesc += f"4. 关于点击后网络反馈慢，延迟时间为 {feedbackInterval} 毫秒，得分 {round(allScores[3] , 2)} 分；\n"
    briefDesc += f"5. 关于点击无反应，得分 {round(allScores[4] , 2)} 分；\n"
    briefDesc += f"6. 关于点击报错，页面错误信息个数为 {data['interactionAttr']['errorCount']['value']}，得分 {round(allScores[5] , 2)} 分；\n"
    briefDesc += f"7. 关于页面加载报错，控制台的报错信息个数为 {data['performanceAttr']['consoleErrors']['value']}，得分 {round(allScores[6] , 2)} 分；\n"
    briefDesc += f"8. 关于页面加载白屏，是否白屏为 {data['interactionAttr']['isBlank']['value']}，得分 {round(allScores[7] , 2)} 分；\n"
    briefDesc += f"9. 关于多个同时出现，上述问题总共出现 {(100 - allScores[8]) // 5} 个，得分 {round(allScores[8] , 2)} 分；\n"

    # 三. 询问 GPT ，问题详述
    askGPTStr = f"现在我给你一些网站体验过程中的问题简述，并且附有一些数据和得分，请您结合数据和得分详细分析一下这些问题，并且给出合适的优化建议，每一条请严格按照我的格式提行进行补充。\n{briefDesc}"
    detailDesc = gpt_35_api_stream(
        "gpt-3.5-turbo", [{'role': 'user', 'content': askGPTStr}])

    # 四. 封装为结构体返回
    return replyT(score, allScores, briefDesc, detailDesc)
