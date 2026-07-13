# 砾石 (Pebble)

<div align="center">

![Wails](https://img.shields.io/badge/Wails-v2.x-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Vue.js](https://img.shields.io/badge/Vue.js-v3.x-4FC08D?style=for-the-badge&logo=vue.js&logoColor=white)
![TypeScript](https://img.shields.io/badge/TypeScript-v5.x-3178C6?style=for-the-badge&logo=typescript&logoColor=white)
![Vite](https://img.shields.io/badge/Vite-v5.x-646CFF?style=for-the-badge&logo=vite&logoColor=white)
![Pinia](https://img.shields.io/badge/Pinia-v3.x-FFD859?style=for-the-badge&logo=pinia&logoColor=black)
![TailwindCSS](https://img.shields.io/badge/TailwindCSS-v3.x-06B6D4?style=for-the-badge&logo=tailwindcss&logoColor=white)
![SQLite](https://img.shields.io/badge/SQLite-3.x-003B57?style=for-the-badge&logo=sqlite&logoColor=white)

一个优雅的项目文档管理系统，基于 Wails + Vue3 + TypeScript + Pinia + TailwindCSS + DaisyUI 构建

**保护您的文档安全，高效管理项目**

</div>

## ✨ 功能特性

### 📁 项目管理
- 创建、编辑、删除项目
- 每个项目独立密码保护
- 按时间线显示项目
- 支持自定义存储路径

### 📋 问题跟踪
- 记录项目问题
- 支持问题状态管理（待解决、进行中、已解决）
- 支持对接人信息
- 问题与文档关联

### 📄 文档管理
- 上传文档（支持拖拽、批量上传）
- AES-256加密存储
- 文件名乱码存储，防止外部访问
- 支持脱密导出原始文件

### 🏷️ 标签系统
- 创建、编辑、删除标签
- 自定义标签颜色
- 文档标签管理
- 标签搜索

### 🔍 搜索功能
- 元数据搜索（标题、标签、描述）
- 支持项目内搜索
- 搜索历史记录

### 🔐 安全特性
- 每个项目独立密码
- 启动时输入密码验证
- 软件运行期间会话保持
- 退出后会话失效

### 🎨 界面设计
- 柔和渐变式设计
- 小清新优雅风格
- 明暗主题切换
- 平滑动画效果
- 响应式布局

### 📊 统计报表
- 项目数量统计
- 问题数量统计
- 文档数量统计
- 问题状态分布

### ⚡ 快速输入
- 浮窗快速输入问题
- 快捷键支持（Ctrl+N）
- 快速创建问题记录

## 🛠️ 技术栈

| 技术 | 版本 | 描述 |
|------|------|------|
| [Wails](https://wails.io) | v2.x | 使用 Go 和 Web 技术构建桌面应用 |
| [Vue 3](https://vuejs.org) | v3.x | 渐进式 JavaScript 框架 |
| [TypeScript](https://www.typescriptlang.org) | v5.x | JavaScript 的超集，添加了静态类型 |
| [Vite](https://vitejs.dev) | v5.x | 下一代前端构建工具 |
| [Pinia](https://pinia.vuejs.org) | v3.x | Vue 的状态管理库 |
| [TailwindCSS](https://tailwindcss.com) | v3.x | 实用优先的 CSS 框架 |
| [DaisyUI](https://daisyui.com) | v5.x | TailwindCSS 的组件库 |
| [SQLite](https://www.sqlite.org) | - | 轻量级数据库 |
| [modernc.org/sqlite](https://pkg.go.dev/modernc.org/sqlite) | - | 纯Go SQLite驱动 |

## 📦 安装

### 前置要求

- [Go](https://golang.org/dl/) 1.21+
- [Node.js](https://nodejs.org/) 18+
- [Wails CLI](https://wails.io/docs/gettingstarted/installation)

### 克隆项目

```bash
git clone https://github.com/your-username/Pebble.git
cd Pebble

# 安装前端依赖
cd frontend
npm install
cd ..
```

## 🚀 开发

### 开发模式

```bash
# 启动开发服务器
wails dev
```

### 构建应用

```bash
# 构建生产版本
wails build
```

### 仅前端开发

```bash
cd frontend
npm run dev
```

## 📁 项目结构

```
Pebble/
├── app.go                 # Wails 应用入口
├── main.go                # Go 主程序
├── wails.json             # Wails 配置文件
├── build/                 # 构建输出目录
├── internal/              # 内部包
│   ├── database/          # 数据库管理
│   │   ├── database.go    # 数据库初始化
│   │   └── models.go      # 数据模型
│   ├── handlers/          # 业务处理器
│   │   ├── project.go     # 项目管理
│   │   ├── issue.go       # 问题跟踪
│   │   ├── document.go    # 文档管理
│   │   ├── tag.go         # 标签管理
│   │   ├── search.go      # 搜索功能
│   │   ├── settings.go    # 设置管理
│   │   └── activity.go    # 活动日志
│   └── crypto/            # 加密工具
│       └── crypto.go      # AES-256加密
├── frontend/              # 前端代码
│   ├── src/
│   │   ├── components/    # Vue 组件
│   │   │   ├── TitleBar.vue       # 自定义标题栏
│   │   │   ├── Sidebar.vue        # 可展开侧边栏
│   │   │   ├── ProjectList.vue    # 项目列表
│   │   │   ├── ProjectForm.vue    # 项目表单
│   │   │   ├── IssueList.vue      # 问题列表
│   │   │   ├── IssueForm.vue      # 问题表单
│   │   │   ├── DocumentList.vue   # 文档列表
│   │   │   ├── DocumentUpload.vue # 文档上传
│   │   │   ├── TagList.vue        # 标签列表
│   │   │   ├── TagForm.vue        # 标签表单
│   │   │   ├── SearchBar.vue      # 搜索栏
│   │   │   ├── ProjectLogin.vue   # 项目登录
│   │   │   └── QuickInput.vue     # 快速输入
│   │   ├── stores/        # Pinia 状态管理
│   │   │   ├── project.ts         # 项目状态
│   │   │   ├── issue.ts           # 问题状态
│   │   │   ├── document.ts        # 文档状态
│   │   │   ├── tag.ts             # 标签状态
│   │   │   ├── search.ts          # 搜索状态
│   │   │   ├── settings.ts        # 设置状态
│   │   │   ├── auth.ts            # 认证状态
│   │   │   └── quickInput.ts      # 快速输入状态
│   │   ├── views/         # 页面组件
│   │   │   ├── Home.vue           # 首页
│   │   │   ├── Projects.vue       # 项目管理
│   │   │   ├── Settings.vue       # 设置页面
│   │   │   └── Statistics.vue     # 统计报表
│   │   ├── utils/         # 工具函数
│   │   │   └── shortcuts.ts       # 快捷键
│   │   ├── App.vue        # 主应用组件
│   │   ├── main.ts        # 前端入口
│   │   ├── router.ts      # 路由配置
│   │   └── style.css      # 全局样式
│   ├── package.json       # 前端依赖
│   └── tailwind.config.cjs # TailwindCSS 配置
└── README.md
```

## 🎨 界面设计

### 配色方案
- **主色渐变**: #667eea → #764ba2（蓝紫色渐变）
- **辅助色渐变**: #f093fb → #f5576c（粉紫色渐变）
- **成功色**: #10b981（绿色）
- **警告色**: #f59e0b（橙色）
- **错误色**: #ef4444（红色）

### 界面特点
- 柔和渐变背景
- 毛玻璃效果
- 平滑动画过渡
- 圆角设计
- 响应式布局

## 🔧 数据存储

### 存储位置
- 默认路径: `用户目录/砾石/`
- 自定义路径: 可在设置中配置

### 数据库
- 使用SQLite存储元数据
- 纯Go实现，无需额外依赖

### 文档存储
- AES-256加密存储
- 文件名乱码存储
- 支持脱密导出

## 📝 使用说明

### 创建项目
1. 点击"新建项目"按钮
2. 输入项目名称、描述和密码
3. 保存项目

### 添加问题
1. 选择项目
2. 点击"新建问题"按钮
3. 输入问题标题、描述和对接人
4. 保存问题

### 上传文档
1. 选择问题
2. 点击"上传文档"按钮
3. 输入文档标题、选择文件和加密密码
4. 上传文档

### 搜索功能
1. 在搜索栏输入关键词
2. 系统自动搜索匹配的项目、问题、文档和标签
3. 点击搜索结果跳转到对应页面

### 快速输入
1. 按 Ctrl+N 打开快速输入窗口
2. 选择项目
3. 输入问题标题和描述
4. 快速创建问题

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📄 许可证

[MIT License](LICENSE)

## 🙏 致谢

- [Wails](https://wails.io) - 优秀的桌面应用框架
- [Vue.js](https://vuejs.org) - 渐进式前端框架
- [TailwindCSS](https://tailwindcss.com) - 实用优先的CSS框架
- [DaisyUI](https://daisyui.com) - 美观的组件库
- [Iconify](https://iconify.design) - 丰富的图标库
