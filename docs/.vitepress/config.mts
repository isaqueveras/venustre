import { defineConfig } from 'vitepress'

export default defineConfig({
  lang: 'en-US',
  title: "🦚 venüstre.io",
  description: "A simple and zero-dependencies library to create and manage routines, schedule execution time, and control competition between other processes.",
  themeConfig: {
    nav: [
      { text: 'Home', link: '/' },
      { text: 'Docs', link: '/guide/what-is-venustre' },
      { text: 'Source', link: 'https://github.com/isaqueveras/venustre' }
    ],

    sidebar: [
      {
        text: 'Introduction',
        collapsed: false,
        items: [
          { text: 'What is Venüste?', link: '/guide/what-is-venustre' },
          { text: 'Getting Started', link: '/guide/getting-started' },
          { text: "How to Use", link: "/guide/how-to-use" },
          { text: "SRP", link: "/guide/server-routine-pattern" },
        ]
      },
      {
        text: "Implementation",
        collapsed: false,
        items: [
          { text: "What Is?", link: "/implementation/what-is" },
          { text: "How It Works", link: "/implementation/how-it-works" },
          { text: "Curated List", link: "/implementation/curated-list" },
          { text: 'Contribution', link: '/implementation/contribution' }
        ]
      },
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/isaqueveras/venustre' }
    ],

    editLink: {
      pattern: 'https://github.com/isaqueveras/venustre/edit/main/docs/:path'
    },

    footer: {
      message: "🌵🇧🇷 Developed by Isaque Veras in his free time",
      copyright: "Copyright © 2024-present · Released Under the MIT License"
    }
  }
})
