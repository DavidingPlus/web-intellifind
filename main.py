import chat
from data import Data


# 最终后端调用的主函数
def jsonParse(data: Data):
    score = 0
    allScores = [0] * 9

    # 1. 跳出率较高
    stayTime = data.m_highBounceRate.stayTime
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

    score += allScores[0] * data.m_highBounceRate.weight

    # 2. 重复点击
    if False == data.m_repeatClick.repeatClick:
        allScores[1] = 100
    else:
        allScores[1] = 60

    score += allScores[1] * data.m_repeatClick.weight

    # 3. 页面打开慢
    pageLoad = data.m_slowPageLoading.pageLoad
    if pageLoad >= 0 and pageLoad < 2:
        allScores[2] = 100
    elif pageLoad >= 2 and pageLoad < 5:
        allScores[2] = 85 + (100 - 85) / (5 - 2) * (pageLoad - 2)
    else:
        allScores[2] = 85 - 3 * (pageLoad - 5)
        if (allScores[2] < 60):
            allScores[2] = 60

    score += allScores[2] * data.m_slowPageLoading.weight

    # 4. 点击后网络反馈慢
    feedbackInterval = data.m_slowNetworkFeedback.feedbackInterval
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

    score += allScores[3] * data.m_slowNetworkFeedback.weight

    # 5. 点击无反应
    allScores[4] = allScores[2] * \
        0.5 + allScores[3] * 0.5

    score += allScores[4] * data.m_clickNoResponse.weight

    # 6. 点击报错
    allScores[5] = 100 - 3 * data.m_clickError.errorCount
    if allScores[5] < 60:
        allScores[5] = 60

    score += allScores[5] * data.m_clickError.weight

    # 7. 页面加载报错
    allScores[6] = 100 - 3 * data.m_pageLoadingError.consoleErrors
    if allScores[6] < 60:
        allScores[6] = 60

    score += allScores[6] * data.m_pageLoadingError.weight

    # 8. 页面加载白屏
    if False == data.m_pageLoadingBlank.isBlank:
        allScores[7] = 100
    else:
        allScores[7] = 60

    score += allScores[7] * data.m_pageLoadingBlank.weight

    # 9. 多个同时出现
    allScores[8] = 100
    for i in range(8):
        if 100 != allScores[i]:
            allScores[8] -= 5

    score += allScores[8] * data.m_multipleProblems.weight

    score = round(score, 2)  # 保留两位小数
    return [score, allScores]


if __name__ == "__main__":
    data = Data()
    data.m_highBounceRate.stayTime = 26.556 * 1000  # 跳出率较高
    data.m_repeatClick.repeatClick = False  # 重复点击
    data.m_slowPageLoading.pageLoad = 3.561  # 页面打开慢
    data.m_slowNetworkFeedback.feedbackInterval = 71.567  # 点击后网络反馈慢
    data.m_clickError.errorCount = 6  # 点击报错
    data.m_pageLoadingError.consoleErrors = 3  # 页面加载报错
    data.m_pageLoadingBlank.isBlank = True  # 页面加载白屏

    print(jsonParse(data))
