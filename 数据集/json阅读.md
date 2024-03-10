# JSON字段分析

# 赛题问题

**①点击无响应 ②跳出率较高 ③重复点击 ④页面打开慢**

**⑤点击后网络反馈慢 ⑥点击报错 ⑦页面加载报错**

**⑧页面加载白屏 ⑨多个同时出现**

# json数据总结

1. `desc`：描述字段，描述干了什么事情
2. `deviceGroupId`：设备组`ID`
3. `envAttr`：包含用户的环境相关信息
4. `eventId`：事件`ID`
5. `eventName`：事件名称
6. `eventTypeId`：事件类型`ID`
7. `interactionAttr`：包含有关用户交互的信息
8. `pageAttr`：包含有关页面的信息
9. `performanceAttr`：包含有关性能的信息
10. `projectId`：项目`ID`
11. `screenDirect`：屏幕方向
12. `sessionId`：会话`ID`
13. `userId`：用户`ID`

其中和问题相关的就是`interactionAttr`和`performanceAttr`

1. `interactionAttr`
   - `errorCount`：错误计数
   - `isBlank`：是否为白屏
   - `isFirstTimeVisit`：是否是第一次访问
   - `isRefresh`： 是否是通过刷新访问
   - `pageActiveDuration`：页面活跃时长，单位毫秒
   - `pageDuration`：页面持续时间，单位毫秒；用户从进入页面到最终离开页面持续的时间间隔
   - `stayTime`：视口停留时间，单位毫秒；`pageDuration`中用户实际看到页面内容的时间，比如有一些页面加载时间不算在里面
   - `timeSincePageLoad`：进入页面前的持续时间
   - `timeSinceSessionStart`：从会话开始的时间，单位毫秒
   - `textInput`：输入文本内容
   - `textInputTime`：输入的文本时间
   - `viewportHeightChangeRate`：窗口高度变化率
   - `viewporChinaidthChangeRate`：窗口宽度变化率（他这个单词是不是打错了）
   - `viewportPosition`：视口距离顶部的距离
   - `viewportPositionPercentage`：视口相对于页面的比例
2. `performanceAttr`
   - `consoleErrors`：控制台错误个数
   - `consoleInfos`：控制台日志个数
   - `consoleWarnings`：控制台警告个数
   - `largestContentfulPaint`：最大内容绘制时间
   - `loadingPerformance`：加载性能
   - `pageLoad`：总页面加载时间
   - `firstInputDelay`：用户首次与页面交互（例如点击按钮）到实际响应之间的延迟时间
   - `firstInteractionAction`：表示用户首次交互的动作类型
   - `firstInteractionPerformance`：表示首次交互的性能评估
   - `errorText`：错误文本
   - `isError`：是否有错误
   - `feedbackInterval`：用户点击到页面反馈的延迟时间
   - `slowNeChinaork`：表示是否存在延迟的网络

# 问题总结

- 每个问题满分 `100` 分，设置不同的权重（加起来等于 `1` ），最终计算出一个最终的结果

其中**点击无反应、点击报错、页面加载报错**问题有待更深入的分析（这几个我想不出来了）

- 点击无反应
  - 页面打开慢可能导致用户感觉点击无反应
  - 权重：
- 跳出率较高
  - `interactionAttr`的`pageDuration`和`stayTime`字段，如果页面停留时间较短，可能用户找不到他们需要的信息，选择离开
    - 这一条可以结合`interactionAttr`的`isFirstTimeVisit`，用户如果是第一次访问，可能不感兴趣就跳出
  - 页面打开慢（见下），也可能导致跳出率高
  - 权重：
- 重复点击
  - `interactionAttr`的`repeatClick`字段
  - 权重：
- 页面打开慢
  - `performanceAttr`的`pageLoad`（记录加载时间）和`loadingPerformance`（记录加载性能）
  - 权重：
- 点击后网络反馈慢
  - `performanceAttr`的`feedbackInterval`记录延迟时间字段和`slowNeChinaork`表示是否存在延迟的网络
  - 权重：
- 点击报错
  - `interactionAttr`的`errorCount`字段，可以表示用户交互过程中的信息交互
  - 权重：
  - `100`分：`0`
  - 每有一个错，扣xxx
- 页面加载报错
  - `performanceAttr`的`consoleErrors`字段，表示页面加载过程中控制台的报错信息
  - 权重：
  - 100分：0
  - 每有一个错，扣xxx
  - `consoleWarnings`算不算在其中？
- 页面加载白屏
  - `interactionAttr`的`isBlank`字段
  - 权重：
  - `100`分：`true`
  - `0`分：`false`
- 多个同时出现
  - 上面的问题同时出现
  - 权重：

