#### 状态模式-行为型设计模式
通常有 3 个组成部分：状态（State）、事件（Event）、动作（Action），事件触发状态的转移及动作的执行
常用在工作流引擎等涉及到很多状态流转的场景

用一个实际的例子超级玛丽：

<img src=mario.jpeg width=400 height=200>

游戏状态：

小马里奥（Small Mario）
超级马里奥（Super Mario）
超级火焰马里奥（Super Fire Mario）
游戏结束时的马里奥（Dead Mario）

游戏事件：
E1：吃蘑菇  E2：吃火焰花  E3：撞到怪物

状态转移图：

<img src=state_change.png width=400 height=200>

不同状态对于不同事件积分处理规则：
#### Small Mario：E1+10，E2+20
#### Super Mario：E1+20，E2+40
#### Super Fire Mario：E1+40，E2+80
