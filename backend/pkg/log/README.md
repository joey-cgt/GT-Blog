1. **github.com/natefinch/lumberjack** (v2.0.0+incompatible)
   - 这是一个用于日志文件轮转(log rotation)的库
   - 在你的代码中用于管理日志文件的大小、备份数量和保留时间
2. **go.uber.org/zap** (v1.27.0)
   - Uber开发的高性能、结构化日志库
   - 提供了比标准库更快、更灵活的日志记录功能
3. **go.uber.org/zap/zapcore**
   - 这是zap包的一部分，不需要单独安装
   - 提供了日志核心组件，用于自定义日志格式、级别等

该日志系统支持：

- 不同的日志级别(debug/info/warn/error)
- 输出到控制台或文件
- 日志文件轮转
- JSON格式的结构化日志
- 调用栈跟踪
- 各种便捷的日志记录方法