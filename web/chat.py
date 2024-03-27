from openai import OpenAI

api_key = str()

client = OpenAI(
    # note： 没有使用 OpenAI 官方的 API，那个要给钱
    # 参考 github 开源项目：https://github.com/chatanywhere/GPT_API_free
    # 这是我的免费的 API KEY
    api_key="sk-ilHiwmIb9kBoQdmcH0uOMWyGdsI1K33OaxbypjfuhvMgG6l7",
    base_url="https://api.chatanywhere.tech/v1"
)


# 非流式响应
def gpt_35_api(model: str, messages: list):
    """为提供的对话消息创建新的回答

    Args:
        messages (list): 完整的对话消息
    """
    completion = client.chat.completions.create(
        model=model,
        messages=messages
    )
    return completion.choices[0].message.content


# 流式响应
# note： 流式响应的含义是， AI 有可能不会一次想出所有的结果，可能需要几次，流式可以等待 AI 思考完毕，返回最终的结果
def gpt_35_api_stream(model: str, messages: list):
    """为提供的对话消息创建新的回答 (流式传输)

    Args:
        messages (list): 完整的对话消息
    """
    res = str()
    stream = client.chat.completions.create(
        model=model,
        messages=messages,
        stream=True,
    )
    for chunk in stream:
        if chunk.choices[0].delta.content is not None:
            res += chunk.choices[0].delta.content
    return res


if __name__ == '__main__':
    model = "gpt-3.5-turbo"
    messages = [
        {
            'role': 'system', 'content': '你是一个网站体验分析师。现在我给你一些网站体验过程中的问题简述，并且附有一些数据和得分，请您结合数据和得分详细分析一下这些问题，并且给出合适的优化建议，每一条请严格按照我的格式提行进行补充。'
        },
        {
            'role': 'user', 'content':
            "经过检测，您本次体验存在的问题如下，以下是问题简述：\n1. 关于跳出率较高，总共停留时长为 519518.0 毫秒，得分 100 分；\n2. 关于重复点击，重复点击次数为 5 次，未重复点击次数为 7 次，得分为 83.33 分；\n3. 关于页面打开慢，页面平均加载时长为 4791.0 毫秒，得分 98.95 分；\n4. 关于点击后网络反馈慢，页面平均延迟时间为 166.73 毫秒，得分 86.67 分；\n5. 关于点击无反应，得分 92.81 分；\n6. 关于点击报错，页面错误信息总数为 9，得分 82 分；\n7. 关于页面加载报错，控制台的报错信息总数为 17，得分 66 分；\n8. 关于页面加载白屏，白屏次数为 0 次，未白屏次数为 1 次，得分 100.0 分；\n9. 关于多个同时出现，上述问题总共出现 6 个，得分 70 分；"
        },]
    # 非流式调用
    # gpt_35_api(model, messages)
    # 流式调用
    print(gpt_35_api_stream(model, messages))
