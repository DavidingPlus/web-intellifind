# -*- coding: UTF-8 -*-

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
            'role': 'user', 'content': '您好，请您简述一下中国上下五千年的历史。'
        }
    ]
    # 非流式调用
    # gpt_35_api(model, messages)
    # 流式调用
    print(gpt_35_api_stream(model, messages))
