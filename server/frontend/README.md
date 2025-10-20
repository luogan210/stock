# 前端文件夹说明

这个文件夹包含了 Go Demo 应用的前端部分，提供了完整的 Web 界面。

## 文件夹结构

```
frontend/
├── static/           # 静态资源文件
│   ├── css/         # CSS 样式文件
│   │   └── style.css
│   └── js/          # JavaScript 文件
│       └── app.js
├── templates/        # HTML 模板文件
│   ├── index.html   # 主页面
│   ├── login.html   # 登录页面
│   └── register.html # 注册页面
└── README.md        # 本文件
```

## 功能特性

### 1. 响应式设计
- 使用现代 CSS 框架
- 支持移动端和桌面端
- 清晰的视觉层次

### 2. 页面功能
- **首页** (`/`): 应用概览和导航
- **登录页面** (`/login`): 用户登录界面
- **注册页面** (`/register`): 用户注册界面
- **用户管理**: 在首页中的用户管理功能
- **微信功能**: 微信小程序 urlLink 生成工具

### 3. JavaScript 功能
- 表单提交处理
- API 调用封装
- 用户信息管理
- 动态内容更新
- 错误处理和提示

### 4. 样式特性
- 现代化 UI 设计
- 卡片式布局
- 按钮和表单样式
- 响应式网格系统
- 状态提示样式

## 路由配置

前端路由通过 `handler/frontend.go` 和 `handler/routes.go` 进行配置：

- `/` - 首页
- `/login` - 登录页面
- `/register` - 注册页面
- `/static/*` - 静态文件服务

## 开发说明

### 添加新页面
1. 在 `frontend/templates/` 中创建 HTML 文件
2. 在 `handler/frontend.go` 中添加对应的服务方法
3. 在 `handler/routes.go` 中注册路由

### 添加新样式
1. 在 `frontend/static/css/style.css` 中添加样式
2. 确保样式具有响应式特性

### 添加新功能
1. 在 `frontend/static/js/app.js` 中添加 JavaScript 代码
2. 更新对应的 HTML 模板

## 技术栈

- **HTML5**: 语义化标签和现代特性
- **CSS3**: Flexbox 布局、响应式设计
- **JavaScript ES6+**: 类、异步函数、模块化
- **Gin**: Go Web 框架，用于服务端渲染

## 浏览器兼容性

- Chrome 60+
- Firefox 55+
- Safari 12+
- Edge 79+

## 部署说明

前端文件通过 Go 的静态文件服务功能提供服务，无需额外的 Web 服务器。所有路由都通过 Gin 框架处理。 