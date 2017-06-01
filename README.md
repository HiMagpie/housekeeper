# housekeeper
HiMagpie 管家，维护并验证 Token、Session和接受推送消息和分发。

### 客户端验证
验证接入 magpie 的客户端（ClientId）是否合法。

### 推送消息接收
提供 HTTP API 给第三方服务 Post 消息进行推送给指定客户端（ClientId）。

### 推送服务器（magpie）和客户端关系维护
1. 维护客户端、客户端连接和推送服务器（magpie）的关系。  
2. 将待推送消息进行归类、整理以及放到推送服务器的待消费队列。
