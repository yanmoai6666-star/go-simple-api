# 所有 Issues 汇总

## Issue #1
- 标题：为用户和产品创建接口添加更严格的输入验证
- 状态：Open
- 创建时间：2025-12-05T03:09:46Z
- 标签：enhancement
- 内容：当前 `CreateUser` 和 `CreateProduct` 接口仅通过 `ShouldBindJSON` 进行基础绑定验证，但未充分利用 `pkg/utils/validation.go` 中的工具函数（如 `ValidateEmail`、`ValidatePrice` 等）。例如：  
  - 用户邮箱格式未验证，可能导致无效邮箱存入数据库  
  - 产品价格可能出现负数，与业务逻辑冲突  
  - 用户名长度未限制（现有 `ValidateName` 函数未在 handler 中调用）  

  建议在 handler 层添加完整验证逻辑，调用现有工具函数对关键字段（邮箱、价格、库存、用户名等）进行校验，并返回更明确的错误提示。

## Issue #2
- 标题：为列表接口添加分页功能
- 状态：Open
- 创建时间：2025-12-05T03:10:29Z
- 标签：enhancement, help wanted
- 内容：目前 `GET /api/v1/products` 和 `GET /api/v1/users` 接口会返回所有数据，当数据量增大时可能导致响应缓慢、资源消耗过高。  
  建议实现分页功能，通过 URL 查询参数（如 `?page=1&pageSize=20`）控制返回数据量，在 service 层的 `GetAllProducts` 和 `GetAllUsers` 方法中添加分页逻辑，同时在响应中返回总条数、总页数等元信息，提升接口性能。

## Issue #3
- 标题：完善 API 文档，补充请求/响应示例
- 状态：Open
- 创建时间：2025-12-05T03:10:56Z
- 标签：documentation, good first issue
- 内容：README 中已列出 API 端点，但缺少具体的请求体格式、响应示例和状态码说明。例如：  
  - `POST /api/v1/products` 需要哪些字段？`price` 字段的格式要求是什么？  
  - 接口成功/失败时的响应结构（如错误信息格式）不明确  

  建议补充每个接口的详细说明，包括：  
  - 请求头、请求体示例（JSON 格式）  
  - 成功响应示例（含状态码 200/201）  
  - 错误响应示例（如 400 无效输入、500 服务器错误）  
  - 字段说明（如 `UserID` 与用户的关联关系）  

## Issue #4
- 标题：添加用户认证功能（JWT 或 Basic Auth）
- 状态：Open
- 创建时间：2025-12-05T03:11:22Z
- 标签：enhancement, help wanted
- 内容：当前 API 接口未做权限控制，任何客户端均可直接调用创建/修改接口，存在安全风险。例如：  
  - 恶意用户可能创建大量无效用户或产品  
  - 无法验证操作发起者的身份，难以追踪责任  

  建议实现基于 JWT 的认证机制：  
  1. 添加 `POST /api/v1/auth/login` 接口，用户登录后返回令牌  
  2. 在需要权限的接口（如创建产品、修改用户）添加中间件验证令牌  
  3. 为用户模型添加 `Password` 字段（需加密存储，如使用 bcrypt）  

## Issue #5
- 标题：修复产品更新接口可能导致的外键关联问题
- 状态：Open
- 创建时间：2025-12-05T03:11:44Z
- 标签：bug
- 内容：在 `product_service.go` 的 `UpdateProduct` 方法中，允许直接修改 `UserID` 字段，但未验证该 `UserID` 是否存在于数据库中。若传入不存在的 `UserID`，会导致产品与无效用户关联，违反外键约束（虽然当前模型定义了 `UserID` 索引，但未显式设置外键约束）。  

  建议：  
  1. 在 `UpdateProduct` 中添加校验，检查 `UserID` 对应的用户是否存在  
  2. 在 GORM 模型中为 `Product.UserID` 添加外键约束（`gorm:"foreignKey:UserID;references:ID"`）  
  3. 若用户不存在，返回 400 错误提示  

## Issue #6
- 标题：扩展集成测试覆盖范围
- 状态：Open
- 创建时间：2025-12-05T03:12:07Z
- 标签：enhancement
- 内容：现有测试中，`product_service_test.go` 和 `user_service_test.go` 仅覆盖了基础的 CRUD 逻辑，但缺少以下场景：  
  - 并发请求下的接口稳定性测试  
  - 边界条件测试（如价格为 0、库存为负数的创建请求）  
  - 关联操作测试（如删除用户后，其关联的产品是否按预期处理）  


  建议补充 `tests/integration` 目录下的测试用例，使用真实 HTTP 请求测试完整接口链路（从 handler 到 service 再到数据库），并在 CI 流程中确保集成测试通过。
