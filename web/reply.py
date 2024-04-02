# -*- coding: UTF-8 -*-

class replyT:

    def __init__(self, totalScore, allScore, briefDesc, detailDesc):
        # 总分
        self.m_totalScore = totalScore
        # 各个部分得分
        self.m_stayTimeScore, self.m_repeatClickScore, self.m_pageLoadScore, self.m_feedbackIntervalScore, self.m_noReactionScore, self.m_errorCountScore, self.m_consoleErrorsScore, self.m_isBlankScore, self.m_occurManyScore = allScore
        # 简要描述
        self.m_briefDesc = briefDesc
        # 详细描述
        self.m_detailDesc = detailDesc

    # 总分
    m_totalScore: float = 0.0

    # 各个问题的分数数组
    # 经商讨，使用 9 个 float
    m_stayTimeScore: float = 0.0
    m_repeatClickScore: float = 0.0
    m_pageLoadScore: float = 0.0
    m_feedbackIntervalScore: float = 0.0
    m_noReactionScore: float = 0.0
    m_errorCountScore: float = 0.0
    m_consoleErrorsScore: float = 0.0
    m_isBlankScore: float = 0.0
    m_occurManyScore: float = 0.0

    # 问题简述
    m_briefDesc = str()

    # 分析和优化建议详述
    m_detailDesc = str()
