# NekoACM Client 🐱🐾

NekoACM: Neural-network Engine Kit of ACM-ICPC

## 操作模式

### 终端启动模式

直接通过终端启动程序并通过命令行参数选择功能。

- `neko`: 以命令行模式启动
- `neko chat`: 直接以对话模式启动
- `neko server`: 直接以服务器模式启动
- `neko problem`: 功能同命令行模式的 `problem` 命令
- `neko testcase`: 功能同命令行模式的 `testcase` 命令
- `neko solution`: 功能同命令行模式的 `solution` 命令
- `neko translate`: 功能同命令行模式的 `translate` 命令
- `neko joke`: 功能同命令行模式的 `joke` 命令

### 对话模式

进入对话模式后，可以与算法竞赛编程助手进行对话。输入 `exit` 退出对话模式。

### 命令行模式

在终端不带参数启动，进入命令行模式后，通过程序内置的命令进行操作。

#### 命令

- `chat`: 启动对话模式
- `parse`: 解析题目为 JSON 格式
- `translate`: 翻译题目为指定语言，支持多种自然语言，可以以 JSON 格式保存到文件
- `problem`: 根据用户提供的题目信息或题解出题，可以以 JSON 格式保存到文件
- `testcase`: 根据用户提供的题目信息或题解生成测试数据，可以以 JSON 格式保存到文件
- `solution`: 根据用户提供的题目信息，生成指定编程语言的题解和解释，可以以 JSON 格式保存到文件
- `joke`: 讲个关于 ACM-ICPC 算法竞赛或 OI 信息学竞赛的冷笑话
- `server`: 启动服务器模式
- `exit`: 退出程序

### 服务器模式

进入服务器模式后，通过 HTTP 请求调用 API 接口，默认端口号为 `14515`。发送请求时，需要在 HTTP 请求头中添加 `Authorization` 字段，值为 `Bearer <token>`。其中 `<token>` 为用户的访问令牌，在 `config.yaml` 配置文件中配置。

#### API 文档

- Apifox: [https://neko-acm.apifox.cn](https://neko-acm.apifox.cn)

#### API 接口

| 功能名称   | 请求方法 | 路由路径                   | 操作者 | 功能简述                               |
|--------|------|------------------------|-----|------------------------------------|
| 服务运行状态 | GET  | /api                   | 用户  | 检查服务是否正常运行                         |
| 编程助手对话 | POST | /api/chat/assistant    | 用户  | 与算法竞赛 AI 编程助手进行对话                  |
| 解析题目   | POST | /api/problem/parse     | 用户  | 将其他格式的题目数据解析为 JSON 格式              |
| 翻译题目   | POST | /api/problem/translate | 用户  | 将题目翻译为指定语言                         |
| 生成题目   | POST | /api/problem/generate  | 用户  | 根据用户提供的题目信息或题解生成题目                 |
| 生成测试数据 | POST | /api/testcase/generate | 用户  | 根据用户提供的题目信息或题解生成测试数据               |
| 生成题解代码 | POST | /api/solution/generate | 用户  | 根据用户提供的题目信息生成指定编程语言的题解             |
| 虚拟评测   | POST | /api/judge/submit      | 用户  | 提交代码进行评测                           |
| 算法笑话   | GET  | /api/misc/joke         | 用户  | 返回一个关于 ACM-ICPC 算法竞赛或 OI 信息学竞赛的冷笑话 |

## 参与贡献

- 如果你觉得这个项目对你有所帮助，欢迎给个 Star⭐️。
- 如果你有任何问题或建议，欢迎提交 Issue。
- 如果你有兴趣参与贡献代码，欢迎提交 Pull Request。

## 许可与声明

本项目采用 LGPL-3.0 license 进行许可，详情请参阅 [LICENSE](LICENSE) 文件。

如果你的软件通过 API 调用了 NekoACM 的服务，不会对你的软件是否开源做出限制，不影响你的软件的开源协议。如果可以的话，欢迎在你的软件和使用 NekoACM 生成的内容中加入 NekoACM 的来源说明。但是，如果你要修改 NekoACM 的源码并重新发布，需要遵循 LGPL-3.0 license 的规定。
