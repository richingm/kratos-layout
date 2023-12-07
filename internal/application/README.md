# application
驱动程序流程，处理具体user case。这里的操作与interface不关联（即同样的一个Service应可被不同的interface（Web、远程调用、异步消息）复用）。这层也适合处理事务、高层次日志(oplog)、安全（权限）。注意：这层不应该有领域逻辑，它只是协调domain层的对象完成真正的工作。
Domain event的监听注册在这层。