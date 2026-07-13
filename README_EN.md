# Wails Template Vue3

<div align="center">

![Wails](https://img.shields.io/badge/Wails-v2.x-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Vue.js](https://img.shields.io/badge/Vue.js-v3.4.x-4FC08D?style=for-the-badge&logo=vue.js&logoColor=white)
![TypeScript](https://img.shields.io/badge/TypeScript-v5.x-3178C6?style=for-the-badge&logo=typescript&logoColor=white)
![Vite](https://img.shields.io/badge/Vite-v5.x-646CFF?style=for-the-badge&logo=vite&logoColor=white)
![Pinia](https://img.shields.io/badge/Pinia-v2.x-FFD859?style=for-the-badge&logo=pinia&logoColor=black)
![TailwindCSS](https://img.shields.io/badge/TailwindCSS-v3.x-06B6D4?style=for-the-badge&logo=tailwindcss&logoColor=white)
![DaisyUI](https://img.shields.io/badge/DaisyUI-v4.x-5A67D8?style=for-the-badge&logo=daisyui&logoColor=white)

A modern desktop application template built with Wails + Vue3 + TypeScript + Pinia + TailwindCSS + DaisyUI

**🎯 Can be used directly as an official Wails template!**

English | [简体中文](./README.md)

</div>

## 🚀 Quick Start

Create a new Wails project using this template:

```bash
wails init -n MyApp -t https://github.com/onewinner/wails-template-vue3
```

This will create a desktop application project with all modern features!

### Template Features

✅ **Ready to Use** - No additional configuration needed, start developing immediately
✅ **Modern UI** - Pre-configured DaisyUI components with 29 themes
✅ **Type Safe** - Complete TypeScript support
✅ **State Management** - Integrated Pinia state management
✅ **Responsive Design** - Mobile-friendly layout
✅ **Glass Effect** - Windows Acrylic native experience

## ✨ Features

- 🚀 **Modern Tech Stack**: Wails v2 + Vue 3 + TypeScript + Vite
- 🎨 **Beautiful UI**: TailwindCSS + DaisyUI component library
- 🌈 **Multi-theme Support**: 29 preset themes with dark mode support
- 📱 **Responsive Design**: Mobile-friendly responsive layout
- 🔧 **State Management**: Pinia reactive state management
- 🎯 **Type Safety**: Complete TypeScript support
- 🪟 **Native Experience**: Windows Acrylic glass effect
- 📦 **Component-based**: Modular component architecture
- 🔄 **Expandable Sidebar**: Collapsible sidebar navigation
- ⚙️ **Complete Settings**: Theme switching, app configuration, etc.

## 🛠️ Tech Stack

| Technology | Version | Description |
|------------|---------|-------------|
| [Wails](https://wails.io) | v2.x | Build desktop apps using Go and web technologies |
| [Vue 3](https://vuejs.org) | v3.4.x | Progressive JavaScript framework |
| [TypeScript](https://www.typescriptlang.org) | v5.x | Superset of JavaScript with static typing |
| [Vite](https://vitejs.dev) | v5.x | Next generation frontend build tool |
| [Pinia](https://pinia.vuejs.org) | v2.x | State management library for Vue |
| [TailwindCSS](https://tailwindcss.com) | v3.x | Utility-first CSS framework |
| [DaisyUI](https://daisyui.com) | v4.x | Component library for TailwindCSS |

## 📦 Installation

### Prerequisites

- [Go](https://golang.org/dl/) 1.18+
- [Node.js](https://nodejs.org/) 16+
- [Wails CLI](https://wails.io/docs/gettingstarted/installation)

### Method 1: Use Wails Template (Recommended)

```bash
# Create new project using this template
wails init -n <Your Appname> -t https://github.com/onewinner/wails-template-vue3

# Enter project directory
cd <Your Appname>

# Install frontend dependencies
cd frontend
npm install
cd ..
```

### Method 2: Clone the project

```bash
git clone https://github.com/onewinner/wails-template-vue3.git
cd wails-template-vue3

# Install frontend dependencies
cd frontend
npm install
cd ..
```

## 🚀 Development

### Development mode

```bash
# Start development server
wails dev
```

### Build application

```bash
# Build production version
wails build
```

### Frontend only development

```bash
cd frontend
npm run dev
```

## 📁 Project Structure

```
wails-template-vue3/
├── app.go                 # Wails app entry
├── main.go               # Go main program
├── wails.json            # Wails configuration
├── build/                # Build output directory
├── frontend/             # Frontend code
│   ├── src/
│   │   ├── components/   # Vue components
│   │   │   ├── Sidebar.vue      # Expandable sidebar
│   │   │   ├── TitleBar.vue     # Custom title bar
│   │   │   ├── WindowControls.vue # Window control buttons
│   │   │   └── DarkMode.vue     # Theme switcher
│   │   ├── stores/       # Pinia state management
│   │   │   ├── theme.ts         # Theme management
│   │   │   └── app.ts           # App state
│   │   ├── views/        # Page components
│   │   │   ├── Home.vue         # Home showcase
│   │   │   └── Setup.vue        # Settings page
│   │   ├── App.vue       # Main app component
│   │   ├── main.ts       # Frontend entry
│   │   └── router.ts     # Route configuration
│   ├── package.json      # Frontend dependencies
│   └── tailwind.config.cjs # TailwindCSS configuration
└── README.md
```

## 🎨 Key Features

### 🌈 Multi-theme Support

The app supports 29 preset themes including:

- Light/Dark modes
- Cupcake, Bumblebee, Emerald themed styles
- Synthwave, Cyberpunk tech themes
- Halloween, Valentine seasonal themes

### 📱 Expandable Sidebar

- Click app icon to expand/collapse sidebar
- Shows app name and menu text when expanded
- Icon-only mode saves space when collapsed
- Tooltip support for menu item names

### 🪟 Custom Title Bar

- Windows Acrylic glass effect
- Integrated theme switcher
- Modern window control buttons
- Draggable window area

### ⚙️ Complete Settings Page

- Theme selection and preview
- App configuration options
- Application information display
- Tech stack information
- Data management operations

## 🤝 Contributing

Issues and Pull Requests are welcome!

## 📄 License

[MIT License](LICENSE)

## 🙏 Acknowledgments

- [Wails](https://wails.io) - Excellent desktop app framework
- [Vue.js](https://vuejs.org) - Progressive frontend framework
- [TailwindCSS](https://tailwindcss.com) - Utility-first CSS framework
- [DaisyUI](https://daisyui.com) - Beautiful component library
- [Iconify](https://iconify.design) - Rich icon library
