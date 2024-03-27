# 自己作算法本地测试的主程序
from web.json_parse import jsonParse

if __name__ == "__main__":
    res = jsonParse("./res/test.json",
                    [0.1, 0.1, 0.15, 0.15, 0.1, 0.1, 0.1, 0.05, 0.15])
    print(res.m_totalScore)
    print(res.m_stayTimeScore, res.m_repeatClickScore, res.m_pageLoadScore, res.m_feedbackIntervalScore,
          res.m_noReactionScore, res.m_errorCountScore, res.m_consoleErrorsScore, res.m_isBlankScore, res.m_occurManyScore)
    print(res.m_briefDesc)
    print(res.m_detailDesc)
