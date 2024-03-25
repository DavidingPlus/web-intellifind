class replyT:

    def __init__(self, score, allScore, briefDesc, detailDesc):
        self.m_score = score
        self.m_allScore = allScore
        self.m_briefDesc = briefDesc
        self.m_detailDesc = detailDesc

    # 总分
    m_score: float = 0

    # 各个问题的分数数组

    #todo 把这个拆开，不要数组，九个问题就用九个float
    m_allScore: float = [0.0] * 9

    # 问题简述
    m_briefDesc = str()

    # 分析和优化建议详述
    m_detailDesc = str()
