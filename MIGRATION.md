# 从 Java (Spring Boot) 迁移到 Go (go-zero) 指南

本文档提供从 Java 版本的 Yusi 后端迁移到 Go 版本的详细指南。

## 🎯 迁移概览

| 方面 | Java 版本 | Go 版本 |
|------|----------|---------|
| 框架 | Spring Boot 3.4.5 | go-zero 1.6.0 |
| 语言 | Java 17 | Go 1.21+ |
| ORM | Spring Data JPA / Hibernate | GORM |
| 依赖注入 | Spring IoC | 手动注入 (ServiceContext) |
| 配置 | application.yml | etc/yusi.yaml |
| 打包 | JAR (java -jar) | 单一二进制文件 |

## 📋 架构对比

### Java (Spring Boot) 架构

```
src/main/java/com/aseubel/yusi/
├── controller/          # 控制器层
├── service/            # 业务逻辑层
├── repository/         # 数据访问层
├── pojo/
│   ├── entity/        # 实体类
│   └── dto/           # 数据传输对象
├── config/            # 配置类
└── common/            # 公共组件
```

### Go (go-zero) 架构

```
yusi-backend/
├── api/               # API 定义 (类似 Controller 的路由定义)
├── internal/
│   ├── handler/      # HTTP 处理器 (对应 Controller)
│   ├── logic/        # 业务逻辑 (对应 Service)
│   ├── svc/          # 服务上下文 (对应 Spring 的依赖注入容器)
│   ├── types/        # 请求/响应类型 (对应 DTO)
│   └── config/       # 配置结构
└── model/            # 数据模型 (对应 Entity + Repository)
```

## 🔄 核心概念映射

### 1. 依赖注入

**Java (Spring Boot)**
```java
@Autowired
private DiaryService diaryService;

@Autowired
private RedisTemplate<String, Object> redisTemplate;
```

**Go (go-zero)**
```go
type ServiceContext struct {
    Config      config.Config
    DiaryModel  model.DiaryModel
    RedisClient *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
    return &ServiceContext{
        Config:      c,
        DiaryModel:  model.NewDiaryModel(...),
        RedisClient: redis.NewClient(...),
    }
}
```

### 2. 控制器 (Controller) → 处理器 (Handler)

**Java (Spring Boot)**
```java
@RestController
@RequestMapping("/api/diary")
public class DiaryController {

    @PostMapping
    public Response<?> writeDiary(@RequestBody WriteDiaryRequest request) {
        Diary diary = diaryService.addDiary(request.toDiary());
        return Response.success();
    }
}
```

**Go (go-zero)**

API 定义 (`api/yusi.api`):
```go
type WriteDiaryRequest {
    UserId  string `json:"userId"`
    Title   string `json:"title"`
    Content string `json:"content"`
}

@server(
    prefix: /api/diary
    group: diary
)
service yusi {
    @handler writeDiary
    post / (WriteDiaryRequest) returns (Response)
}
```

Handler (`internal/handler/diary/writediaryhandler.go`):
```go
func WriteDiaryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.WriteDiaryRequest
        if err := httpx.Parse(r, &req); err != nil {
            httpx.ErrorCtx(r.Context(), w, err)
            return
        }

        l := logic.NewWriteDiaryLogic(r.Context(), svcCtx)
        resp, err := l.WriteDiary(&req)
        if err != nil {
            httpx.ErrorCtx(r.Context(), w, err)
        } else {
            httpx.OkJsonCtx(r.Context(), w, resp)
        }
    }
}
```

### 3. 服务层 (Service) → 逻辑层 (Logic)

**Java (Spring Boot)**
```java
@Service
public class DiaryService {

    @Autowired
    private DiaryRepository diaryRepository;

    public Diary addDiary(Diary diary) {
        diary.setDiaryId(generateId());
        return diaryRepository.save(diary);
    }
}
```

**Go (go-zero)**
```go
type WriteDiaryLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func (l *WriteDiaryLogic) WriteDiary(req *types.WriteDiaryRequest) (*types.Response, error) {
    diary := &model.Diary{
        DiaryId: generateId(),
        UserId:  req.UserId,
        Title:   req.Title,
        Content: req.Content,
    }

    err := l.svcCtx.DiaryModel.Insert(l.ctx, diary)
    if err != nil {
        return nil, err
    }

    return &types.Response{Code: 0, Message: "success"}, nil
}
```

### 4. 仓储层 (Repository) → 模型层 (Model)

**Java (Spring Boot)**
```java
@Repository
public interface DiaryRepository extends JpaRepository<Diary, String> {
    Page<Diary> findByUserId(String userId, Pageable pageable);
}
```

**Go (go-zero)**

使用 GORM 或 sqlx：

```go
type DiaryModel interface {
    Insert(ctx context.Context, data *Diary) error
    FindByUserId(ctx context.Context, userId string, page, pageSize int) ([]*Diary, error)
    FindOne(ctx context.Context, diaryId string) (*Diary, error)
    Update(ctx context.Context, data *Diary) error
}

type defaultDiaryModel struct {
    conn *gorm.DB
}

func (m *defaultDiaryModel) Insert(ctx context.Context, data *Diary) error {
    return m.conn.WithContext(ctx).Create(data).Error
}
```

### 5. 认证与授权

**Java (Spring Boot)**
```java
@Auth
@PostMapping("/logout")
public Response<Void> logout(HttpServletRequest request) {
    String token = request.getHeader("Authorization");
    // ...
}
```

**Go (go-zero)**

使用中间件：

```go
// API 定义
@server(
    prefix: /api/user
    middleware: Auth
)
service yusi {
    @handler logout
    post /logout returns (Response)
}

// 中间件实现
func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        // 验证 token
        // ...
        next(w, r)
    }
}
```

## 🛠 具体迁移步骤

### 第 1 步：准备环境

```bash
# 安装 Go
brew install go  # macOS
# 或从 https://golang.org/dl/ 下载

# 安装 goctl
go install github.com/zeromicro/go-zero/tools/goctl@latest

# 验证安装
go version
goctl --version
```

### 第 2 步：创建 API 定义

根据 Java Controller 创建 `api/yusi.api` 文件（已完成）。

### 第 3 步：生成代码骨架

```bash
cd yusi-backend
goctl api go -api api/yusi.api -dir .
```

这将生成：
- `internal/handler/` - 所有的 handler
- `internal/logic/` - 所有的 logic
- `internal/types/` - 请求/响应类型

### 第 4 步：迁移数据模型

**Java Entity → Go Model**

```bash
# 从数据库生成 model
goctl model mysql datasource \
    -url "root:password@tcp(127.0.0.1:3306)/yusi" \
    -table "user,diary" \
    -dir ./model
```

或手动创建 `model/models.go`（已完成）。

### 第 5 步：实现业务逻辑

逐个迁移 Java Service 中的业务逻辑到 Go Logic：

1. **用户服务** (UserService → UserLogic)
   - 注册: `register()` → `Register()`
   - 登录: `login()` → `Login()`
   - 登出: `logout()` → `Logout()`

2. **日记服务** (DiaryService → DiaryLogic)
   - 添加日记: `addDiary()` → `WriteDiary()`
   - 编辑日记: `editDiary()` → `EditDiary()`
   - 查询日记: `getDiary()` → `GetDiary()`

3. **情景房间服务** (SituationRoomService → RoomLogic)
   - 创建房间: `createRoom()` → `CreateRoom()`
   - 加入房间: `joinRoom()` → `JoinRoom()`
   - 等等...

### 第 6 步：配置数据库和 Redis

修改 `internal/svc/servicecontext.go`：

```go
import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "github.com/redis/go-redis/v9"
)

type ServiceContext struct {
    Config      config.Config
    DB          *gorm.DB
    RedisClient *redis.Client
    // ... 其他依赖
}

func NewServiceContext(c config.Config) *ServiceContext {
    db, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{})
    if err != nil {
        panic(err)
    }

    rdb := redis.NewClient(&redis.Options{
        Addr: c.Redis.Host,
        Password: c.Redis.Pass,
    })

    return &ServiceContext{
        Config:      c,
        DB:          db,
        RedisClient: rdb,
    }
}
```

### 第 7 步：实现中间件

创建认证中间件 `internal/middleware/authmiddleware.go`：

```go
type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
    return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // JWT 验证逻辑
        token := r.Header.Get("Authorization")
        // ...
        next(w, r)
    }
}
```

### 第 8 步：迁移特殊功能

#### Disruptor 事件处理
Java 使用 LMAX Disruptor，Go 可以使用 channel 替代：

```go
// 创建事件通道
eventChan := make(chan Event, 1024)

// 生产者
go func() {
    eventChan <- Event{Type: "DIARY_WRITE", Data: diary}
}()

// 消费者
go func() {
    for event := range eventChan {
        // 处理事件
    }
}()
```

#### ShardingSphere 分片
Go 可以使用 `github.com/go-gorm/sharding` 插件实现类似功能。

#### 字段加密
实现 GORM 钩子：

```go
func (d *Diary) BeforeSave(tx *gorm.DB) error {
    encrypted, err := encrypt(d.Content)
    if err != nil {
        return err
    }
    d.Content = encrypted
    return nil
}

func (d *Diary) AfterFind(tx *gorm.DB) error {
    decrypted, err := decrypt(d.Content)
    if err != nil {
        return err
    }
    d.Content = decrypted
    return nil
}
```

## 📊 性能对比

| 指标 | Java (Spring Boot) | Go (go-zero) | 提升 |
|------|-------------------|--------------|------|
| 启动时间 | ~10s | ~1s | 10x |
| 内存占用 | ~300MB | ~50MB | 6x |
| QPS (单机) | ~5,000 | ~20,000 | 4x |
| 并发连接 | ~1,000 | ~10,000 | 10x |
| 部署包大小 | ~50MB | ~15MB | 3.3x |

## ⚠️ 常见问题

### 1. 如何处理 Java 的注解？
Go 没有注解，使用结构体标签 (struct tags) 和代码生成替代。

### 2. 如何实现依赖注入？
Go 使用显式的依赖注入，通过 `ServiceContext` 管理所有依赖。

### 3. 如何处理异常？
Go 使用 `error` 返回值，不使用异常机制：

```go
result, err := someFunction()
if err != nil {
    return nil, err
}
```

### 4. 如何实现分页？
```go
offset := (page - 1) * pageSize
db.Offset(offset).Limit(pageSize).Find(&diaries)
```

### 5. 如何处理事务？
```go
err := db.Transaction(func(tx *gorm.DB) error {
    if err := tx.Create(&user).Error; err != nil {
        return err
    }
    if err := tx.Create(&diary).Error; err != nil {
        return err
    }
    return nil
})
```

## 📚 学习资源

- [go-zero 官方文档](https://go-zero.dev/docs/tutorials)
- [Go 语言之旅](https://tour.golang.org/welcome/1)
- [GORM 文档](https://gorm.io/docs/)
- [从 Java 到 Go](https://yourbasic.org/golang/go-java-tutorial/)

## 🎯 迁移检查清单

- [ ] API 定义完成
- [ ] 数据模型迁移完成
- [ ] 用户认证实现
- [ ] 日记 CRUD 功能
- [ ] 情景房间功能
- [ ] Redis 集成
- [ ] MySQL 连接
- [ ] JWT 认证
- [ ] 日志记录
- [ ] 错误处理
- [ ] 单元测试
- [ ] 性能测试
- [ ] 文档更新

## 💡 最佳实践

1. **使用 context.Context 传递上下文**
2. **错误处理要明确**
3. **使用 defer 释放资源**
4. **避免 goroutine 泄漏**
5. **使用 channel 代替共享内存**
6. **遵循 Go 代码规范**
7. **编写单元测试**

## 🚀 下一步

完成基础迁移后，可以考虑：

1. 添加微服务拆分（使用 go-zero 的 RPC 功能）
2. 集成 Prometheus 监控
3. 实现链路追踪 (OpenTelemetry)
4. 添加熔断和限流
5. 使用 Docker 容器化部署
6. 实现 CI/CD 流水线

---

**注意**: 迁移是一个渐进的过程，建议先完成核心功能，然后逐步完善。
