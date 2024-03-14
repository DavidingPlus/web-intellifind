# 跳出率较高
class highBounceRate:
    desc = "跳出率较高"  # 问题描述
    weight = 0.15  # 权重

    stayTime: float = 0.0  # 数据字段，用户停留的时间，单位毫秒


# 重复点击
class repeatClick:
    desc = "重复点击"  # 问题描述
    weight = 0.05  # 权重

    repeatClick: bool = False  # 数据字段，是否重复点击


# 页面打开慢
class slowPageLoading:
    desc = "页面打开慢"  # 问题描述
    weight = 0.15  # 权重

    pageLoad: float = 0.0  # 数据字段，页面加载时间


# 点击后网络反馈慢
class slowNetworkFeedback:
    desc = "点击后网络反馈慢"  # 问题描述
    weight = 0.15  # 权重

    feedbackInterval: float = 0.0  # 数据字段，延迟时间


# 点击无反应
class clickNoResponse:
    desc = "点击无反应"  # 问题描述
    weight = 0.1  # 权重


# 点击报错
class clickError:
    desc = "点击报错"  # 问题描述
    weight = 0.1  # 权重

    errorCount: int = 0  # 数据字段， errorCount 个数


# 页面加载报错
class pageLoadingError:
    desc = "页面加载报错"  # 问题描述
    weight = 0.1  # 权重

    consoleErrors: int = 0  # 数据字段， consoleErrors 个数


# 页面加载白屏
class pageLoadingBlank:
    desc = "页面加载白屏"  # 问题描述
    weight = 0.05  # 权重

    isBlank: bool = False  # 数据字段，是否白屏


# 多个同时出现
class multipleProblems:
    desc = "多个同时出现"  # 问题描述
    weight = 0.15  # 权重

    num: int = 0  # 数据字段，多次出现的问题个数
