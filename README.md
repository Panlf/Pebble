# Wails Template Vue3

<div align="center">

![Wails](https://img.shields.io/badge/Wails-v2.x-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Vue.js](https://img.shields.io/badge/Vue.js-v3.4.x-4FC08D?style=for-the-badge&logo=vue.js&logoColor=white)
![TypeScript](https://img.shields.io/badge/TypeScript-v5.x-3178C6?style=for-the-badge&logo=typescript&logoColor=white)
![Vite](https://img.shields.io/badge/Vite-v5.x-646CFF?style=for-the-badge&logo=vite&logoColor=white)
![Pinia](https://img.shields.io/badge/Pinia-v2.x-FFD859?style=for-the-badge&logo=pinia&logoColor=black)
![TailwindCSS](https://img.shields.io/badge/TailwindCSS-v3.x-06B6D4?style=for-the-badge&logo=tailwindcss&logoColor=white)
![DaisyUI](https://img.shields.io/badge/DaisyUI-v4.x-5A67D8?style=for-the-badge&logo=daisyui&logoColor=white)

一个现代化的桌面应用模板，基于 Wails + Vue3 + TypeScript + Pinia + TailwindCSS + DaisyUI 构建

**🎯 可直接作为 Wails 官方模板使用！**

[English](./README_EN.md) | 简体中文

</div>

## 🚀 快速开始

使用此模板创建新的 Wails 项目：

```bash
wails init -n MyApp -t https://github.com/onewinner/wails-template-vue3
```

这将创建一个包含所有现代化功能的桌面应用项目！

### 模板特性

✅ **开箱即用** - 无需额外配置，直接开始开发
✅ **现代化UI** - 预配置的DaisyUI组件和29个主题
✅ **类型安全** - 完整的TypeScript支持
✅ **状态管理** - 集成Pinia状态管理
✅ **响应式设计** - 移动端友好的布局
✅ **毛玻璃效果** - Windows Acrylic原生体验

## ✨ 特性

- 🚀 **现代技术栈**: Wails v2 + Vue 3 + TypeScript + Vite
- 🎨 **美观UI**: TailwindCSS + DaisyUI 组件库
- 🌈 **多主题支持**: 29个预设主题，支持深色模式
- 📱 **响应式设计**: 移动端友好的响应式布局
- 🔧 **状态管理**: Pinia 响应式状态管理
- 🎯 **类型安全**: 完整的 TypeScript 支持
- 🪟 **原生体验**: Windows Acrylic 毛玻璃效果
- 📦 **组件化**: 模块化的组件架构
- 🔄 **可展开侧边栏**: 支持展开/收缩的侧边栏导航
- ⚙️ **完整设置页**: 主题切换、应用配置等

## 🛠️ 技术栈

| 技术 | 版本 | 描述 |
|------|------|------|
| [Wails](https://wails.io) | v2.x | 使用 Go 和 Web 技术构建桌面应用 |
| [Vue 3](https://vuejs.org) | v3.4.x | 渐进式 JavaScript 框架 |
| [TypeScript](https://www.typescriptlang.org) | v5.x | JavaScript 的超集，添加了静态类型 |
| [Vite](https://vitejs.dev) | v5.x | 下一代前端构建工具 |
| [Pinia](https://pinia.vuejs.org) | v2.x | Vue 的状态管理库 |
| [TailwindCSS](https://tailwindcss.com) | v3.x | 实用优先的 CSS 框架 |
| [DaisyUI](https://daisyui.com) | v4.x | TailwindCSS 的组件库 |

## 📦 安装

### 前置要求

- [Go](https://golang.org/dl/) 1.18+
- [Node.js](https://nodejs.org/) 16+
- [Wails CLI](https://wails.io/docs/gettingstarted/installation)

### 方式一：使用 Wails 模板（推荐）

```bash
# 使用此模板创建新项目
wails init -n <Your Appname> -t https://github.com/onewinner/wails-template-vue3

# 进入项目目录
cd <Your Appname>

# 安装前端依赖
cd frontend
npm install
cd ..
```

### 方式二：克隆项目

```bash
git clone https://github.com/onewinner/wails-template-vue3.git
cd wails-template-vue3

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
wails-template-vue3/
├── app.go                 # Wails 应用入口
├── main.go               # Go 主程序
├── wails.json            # Wails 配置文件
├── build/                # 构建输出目录
├── frontend/             # 前端代码
│   ├── src/
│   │   ├── components/   # Vue 组件
│   │   │   ├── Sidebar.vue      # 可展开侧边栏
│   │   │   ├── TitleBar.vue     # 自定义标题栏
│   │   │   ├── WindowControls.vue # 窗口控制按钮
│   │   │   └── DarkMode.vue     # 主题切换
│   │   ├── stores/       # Pinia 状态管理
│   │   │   ├── theme.ts         # 主题管理
│   │   │   └── app.ts           # 应用状态
│   │   ├── views/        # 页面组件
│   │   │   ├── Home.vue         # 首页展示
│   │   │   └── Setup.vue        # 设置页面
│   │   ├── App.vue       # 主应用组件
│   │   ├── main.ts       # 前端入口
│   │   └── router.ts     # 路由配置
│   ├── package.json      # 前端依赖
│   └── tailwind.config.cjs # TailwindCSS 配置
└── README.md
```

## 🎨 主要功能

### 🌈 多主题支持

应用支持 29 个预设主题，包括：

- 浅色/深色模式
- 纸杯蛋糕、大黄蜂、翡翠等特色主题
- 合成波、赛博朋克等科技主题
- 万圣节、情人节等节日主题

### 📱 可展开侧边栏

- 点击应用图标可展开/收缩侧边栏
- 展开时显示应用名称和菜单文本
- 收缩时仅显示图标，节省空间
- 支持工具提示显示菜单名称

### 🪟 自定义标题栏

- Windows Acrylic 毛玻璃效果
- 集成主题切换按钮
- 现代化窗口控制按钮
- 可拖拽窗口区域

### ⚙️ 完整设置页面

- 主题选择和预览
- 应用配置选项
- 应用信息展示
- 技术栈信息
- 数据管理操作

## 🔧 自定义配置

### 添加新菜单项

在 `frontend/src/App.vue` 中修改 `menuItems`：

```typescript
const menuItems = [
  { text: "首页", href: "/home", icon: "material-symbols:home-app-logo" },
  { text: "新页面", href: "/new-page", icon: "material-symbols:new-window" },
];
```

### 添加新主题

在 `frontend/src/stores/theme.ts` 中添加主题：

```typescript
const themes = [
  // 现有主题...
  { name: 'custom', label: '自定义', icon: 'icon-park:palette' },
];
```

### 修改应用信息

在 `frontend/src/stores/app.ts` 中修改 `appInfo`：

```typescript
const appInfo = ref<AppInfo>({
  name: '你的应用名称',
  version: '1.0.0',
  description: '应用描述',
  author: '你的名字',
  repository: 'https://github.com/your-username/your-repo',
  license: 'MIT'
});
```

## 📝 开发指南

### 添加新页面

1. 在 `frontend/src/views/` 创建新的 Vue 组件
2. 在 `frontend/src/router.ts` 添加路由
3. 在侧边栏菜单中添加导航项

### 使用 Pinia Store

```typescript
// 在组件中使用
import { useThemeStore } from '../stores/theme';

const themeStore = useThemeStore();
themeStore.setTheme('dark');
```

### 使用 DaisyUI 组件

```vue
<template>
  <button class="btn btn-primary">按钮</button>
  <div class="card bg-base-200 shadow-xl">
    <div class="card-body">
      <h2 class="card-title">卡片标题</h2>
      <p>卡片内容</p>
    </div>
  </div>
</template>
```

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
