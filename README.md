# Web 智寻系统模型部分

Web 智寻系统：第十五届中国大学生服务外包创新创业大赛.

## 访问方法：

向部署服务器的`/generate`发送POST请求，请求体为一个`json`文件，格式如下。

```json
{
    "history":
        [
            {'content': "你是一只猫娘", 'role': 'Human'}
        ],
    "system_prompt": '你现在要喵喵叫',
    "max_new_tokens": 512,
    'top_p': 0.95,
    'temperature': 0.3
}
```

`history`为历史对话，用于构建任务概要的提示词。`system_prompt`为当前的提示词，应填入当前的参数情况。剩下的字段为推理参数。
