import problems


# 整体数据结构体
class Data:
    # 点击无反应
    m_clickNoResponse = problems.clickNoResponse()

    # 跳出率较高
    m_highBounceRate = problems.highBounceRate()

    # 重复点击
    m_repeatClick = problems.repeatClick()

    # 页面打开慢
    m_slowPageLoading = problems.slowPageLoading()

    # 点击后网络反馈慢
    m_slowNetworkFeedback = problems.slowNetworkFeedback()

    # 点击报错
    m_clickError = problems.clickError()

    # 页面加载报错
    m_pageLoadingError = problems.pageLoadingError()

    # 页面加载白屏
    m_pageLoadingBlank = problems.pageLoadingBlank()

    # 多个同时出现
    m_multipleProblems = problems.multipleProblems()


if __name__ == "__main__":
    pass
