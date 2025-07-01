// vite.config.js
import { defineConfig } from "file:///Users/duzhipeng/Documents/Magic/duzhipeng.com/www/node_modules/vite/dist/node/index.js";
import Vue from "file:///Users/duzhipeng/Documents/Magic/duzhipeng.com/www/node_modules/@vitejs/plugin-vue/dist/index.mjs";
import Components from "file:///Users/duzhipeng/Documents/Magic/duzhipeng.com/www/node_modules/unplugin-vue-components/dist/vite.js";
import { ArcoResolver } from "file:///Users/duzhipeng/Documents/Magic/duzhipeng.com/www/node_modules/unplugin-vue-components/dist/resolvers.js";
import externalGlobals from "file:///Users/duzhipeng/Documents/Magic/duzhipeng.com/www/node_modules/rollup-plugin-external-globals/index.js";
import Markdown from "file:///Users/duzhipeng/Documents/Magic/duzhipeng.com/www/node_modules/unplugin-vue-markdown/dist/vite.js";
import { vitePluginForArco } from "file:///Users/duzhipeng/Documents/Magic/duzhipeng.com/www/node_modules/@arco-plugins/vite-vue/lib/index.js";
import markdownItAttrs from "file:///Users/duzhipeng/Documents/Magic/duzhipeng.com/www/node_modules/markdown-it-attrs/index.js";
import markdownItPrism from "file:///Users/duzhipeng/Documents/Magic/duzhipeng.com/www/node_modules/markdown-it-prism/build/index.js";
import markdownItTexmath from "file:///Users/duzhipeng/Documents/Magic/duzhipeng.com/www/node_modules/markdown-it-texmath/texmath.js";
import katex from "file:///Users/duzhipeng/Documents/Magic/duzhipeng.com/www/node_modules/katex/dist/katex.mjs";
var vite_config_default = defineConfig(({ mode }) => {
  return {
    plugins: [
      Vue({
        include: [/\.vue$/, /\.md$/]
        // <--
      }),
      vitePluginForArco({
        theme: "@arco-themes/vue-duzhipeng",
        style: true
      }),
      Markdown({
        // default options passed to markdown-it
        // see: https://markdown-it.github.io/markdown-it/
        markdownItOptions: { html: true },
        // A function providing the Markdown It instances gets the ability to apply custom settings/plugins
        markdownItSetup(md) {
          md.use(markdownItAttrs);
          md.renderer.rules.fence = function(tokens, idx, options, env, slf) {
            const token = tokens[idx];
            return "<pre" + slf.renderAttrs(token) + "><code>" + token.content + "</code></pre>";
          };
          md.use(markdownItPrism, { highlightInlineCode: true });
          md.use(markdownItTexmath, {
            engine: katex,
            delimiters: ["dollars", "beg_end"],
            katexOptions: { macros: { "\\RR": "\\mathbb{R}" } }
          });
        }
      }),
      Components({
        resolvers: [ArcoResolver()]
      })
    ],
    build: {
      rollupOptions: {
        external: [
          // "vue",
          // "vue-router",
          //     "@arco-design/web-vue",
          // "axios",
          //     "vue-demi",
        ],
        plugins: [
          externalGlobals({
            // vue: "Vue",
            // "vue-router": "VueRouter",
            //         "@arco-design/web-vue": "ArcoVue",
            // axios: "axios",
            //         "vue-demi": "VueDemi", // pinia 源码中引入了 vue-demi 这个包
          })
        ]
      }
    }
  };
});
export {
  vite_config_default as default
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsidml0ZS5jb25maWcuanMiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImNvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9kaXJuYW1lID0gXCIvVXNlcnMvZHV6aGlwZW5nL0RvY3VtZW50cy9NYWdpYy9kdXpoaXBlbmcuY29tL3d3d1wiO2NvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9maWxlbmFtZSA9IFwiL1VzZXJzL2R1emhpcGVuZy9Eb2N1bWVudHMvTWFnaWMvZHV6aGlwZW5nLmNvbS93d3cvdml0ZS5jb25maWcuanNcIjtjb25zdCBfX3ZpdGVfaW5qZWN0ZWRfb3JpZ2luYWxfaW1wb3J0X21ldGFfdXJsID0gXCJmaWxlOi8vL1VzZXJzL2R1emhpcGVuZy9Eb2N1bWVudHMvTWFnaWMvZHV6aGlwZW5nLmNvbS93d3cvdml0ZS5jb25maWcuanNcIjtpbXBvcnQgeyBkZWZpbmVDb25maWcgfSBmcm9tIFwidml0ZVwiO1xuaW1wb3J0IFZ1ZSBmcm9tIFwiQHZpdGVqcy9wbHVnaW4tdnVlXCI7XG5pbXBvcnQgQ29tcG9uZW50cyBmcm9tIFwidW5wbHVnaW4tdnVlLWNvbXBvbmVudHMvdml0ZVwiO1xuaW1wb3J0IHsgQXJjb1Jlc29sdmVyIH0gZnJvbSBcInVucGx1Z2luLXZ1ZS1jb21wb25lbnRzL3Jlc29sdmVyc1wiO1xuaW1wb3J0IGV4dGVybmFsR2xvYmFscyBmcm9tIFwicm9sbHVwLXBsdWdpbi1leHRlcm5hbC1nbG9iYWxzXCI7XG5pbXBvcnQgTWFya2Rvd24gZnJvbSBcInVucGx1Z2luLXZ1ZS1tYXJrZG93bi92aXRlXCI7XG5pbXBvcnQgeyB2aXRlUGx1Z2luRm9yQXJjbyB9IGZyb20gXCJAYXJjby1wbHVnaW5zL3ZpdGUtdnVlXCI7XG5pbXBvcnQgbWFya2Rvd25JdEF0dHJzIGZyb20gXCJtYXJrZG93bi1pdC1hdHRyc1wiO1xuaW1wb3J0IG1hcmtkb3duSXRQcmlzbSBmcm9tIFwibWFya2Rvd24taXQtcHJpc21cIjtcbmltcG9ydCBtYXJrZG93bkl0VGV4bWF0aCBmcm9tIFwibWFya2Rvd24taXQtdGV4bWF0aFwiO1xuaW1wb3J0IGthdGV4IGZyb20gXCJrYXRleFwiO1xuXG4vLyBodHRwczovL3ZpdGVqcy5kZXYvY29uZmlnL1xuZXhwb3J0IGRlZmF1bHQgZGVmaW5lQ29uZmlnKCh7IG1vZGUgfSkgPT4ge1xuICByZXR1cm4ge1xuICAgIHBsdWdpbnM6IFtcbiAgICAgIFZ1ZSh7XG4gICAgICAgIGluY2x1ZGU6IFsvXFwudnVlJC8sIC9cXC5tZCQvXSwgLy8gPC0tXG4gICAgICB9KSxcbiAgICAgIHZpdGVQbHVnaW5Gb3JBcmNvKHtcbiAgICAgICAgdGhlbWU6IFwiQGFyY28tdGhlbWVzL3Z1ZS1kdXpoaXBlbmdcIixcbiAgICAgICAgc3R5bGU6IHRydWUsXG4gICAgICB9KSxcbiAgICAgIE1hcmtkb3duKHtcbiAgICAgICAgLy8gZGVmYXVsdCBvcHRpb25zIHBhc3NlZCB0byBtYXJrZG93bi1pdFxuICAgICAgICAvLyBzZWU6IGh0dHBzOi8vbWFya2Rvd24taXQuZ2l0aHViLmlvL21hcmtkb3duLWl0L1xuICAgICAgICBtYXJrZG93bkl0T3B0aW9uczogeyBodG1sOiB0cnVlIH0sXG4gICAgICAgIC8vIEEgZnVuY3Rpb24gcHJvdmlkaW5nIHRoZSBNYXJrZG93biBJdCBpbnN0YW5jZXMgZ2V0cyB0aGUgYWJpbGl0eSB0byBhcHBseSBjdXN0b20gc2V0dGluZ3MvcGx1Z2luc1xuICAgICAgICBtYXJrZG93bkl0U2V0dXAobWQpIHtcbiAgICAgICAgICAvLyBhZGQgY29kZSBzeW50YXggaGlnaGxpZ2h0aW5nIHdpdGggUHJpc21cbiAgICAgICAgICBtZC51c2UobWFya2Rvd25JdEF0dHJzKTtcbiAgICAgICAgICAvLyBjdXN0b20gcmVuZGVyZXIgZm9yIGZlbmNlc1xuICAgICAgICAgIG1kLnJlbmRlcmVyLnJ1bGVzLmZlbmNlID0gZnVuY3Rpb24gKHRva2VucywgaWR4LCBvcHRpb25zLCBlbnYsIHNsZikge1xuICAgICAgICAgICAgY29uc3QgdG9rZW4gPSB0b2tlbnNbaWR4XTtcbiAgICAgICAgICAgIHJldHVybiAoXG4gICAgICAgICAgICAgIFwiPHByZVwiICtcbiAgICAgICAgICAgICAgc2xmLnJlbmRlckF0dHJzKHRva2VuKSArXG4gICAgICAgICAgICAgIFwiPlwiICtcbiAgICAgICAgICAgICAgXCI8Y29kZT5cIiArXG4gICAgICAgICAgICAgIHRva2VuLmNvbnRlbnQgK1xuICAgICAgICAgICAgICBcIjwvY29kZT5cIiArXG4gICAgICAgICAgICAgIFwiPC9wcmU+XCJcbiAgICAgICAgICAgICk7XG4gICAgICAgICAgfTtcbiAgICAgICAgICBtZC51c2UobWFya2Rvd25JdFByaXNtLCB7IGhpZ2hsaWdodElubGluZUNvZGU6IHRydWUgfSk7XG4gICAgICAgICAgbWQudXNlKG1hcmtkb3duSXRUZXhtYXRoLCB7XG4gICAgICAgICAgICBlbmdpbmU6IGthdGV4LFxuICAgICAgICAgICAgZGVsaW1pdGVyczogW1wiZG9sbGFyc1wiLCBcImJlZ19lbmRcIl0sXG4gICAgICAgICAgICBrYXRleE9wdGlvbnM6IHsgbWFjcm9zOiB7IFwiXFxcXFJSXCI6IFwiXFxcXG1hdGhiYntSfVwiIH0gfSxcbiAgICAgICAgICB9KTtcbiAgICAgICAgfSxcbiAgICAgIH0pLFxuICAgICAgQ29tcG9uZW50cyh7XG4gICAgICAgIHJlc29sdmVyczogW0FyY29SZXNvbHZlcigpXSxcbiAgICAgIH0pLFxuICAgIF0sXG4gICAgYnVpbGQ6IHtcbiAgICAgIHJvbGx1cE9wdGlvbnM6IHtcbiAgICAgICAgZXh0ZXJuYWw6IFtcbiAgICAgICAgICAvLyBcInZ1ZVwiLFxuICAgICAgICAgIC8vIFwidnVlLXJvdXRlclwiLFxuICAgICAgICAgIC8vICAgICBcIkBhcmNvLWRlc2lnbi93ZWItdnVlXCIsXG4gICAgICAgICAgLy8gXCJheGlvc1wiLFxuICAgICAgICAgIC8vICAgICBcInZ1ZS1kZW1pXCIsXG4gICAgICAgIF0sXG4gICAgICAgIHBsdWdpbnM6IFtcbiAgICAgICAgICBleHRlcm5hbEdsb2JhbHMoe1xuICAgICAgICAgICAgLy8gdnVlOiBcIlZ1ZVwiLFxuICAgICAgICAgICAgLy8gXCJ2dWUtcm91dGVyXCI6IFwiVnVlUm91dGVyXCIsXG4gICAgICAgICAgICAvLyAgICAgICAgIFwiQGFyY28tZGVzaWduL3dlYi12dWVcIjogXCJBcmNvVnVlXCIsXG4gICAgICAgICAgICAvLyBheGlvczogXCJheGlvc1wiLFxuICAgICAgICAgICAgLy8gICAgICAgICBcInZ1ZS1kZW1pXCI6IFwiVnVlRGVtaVwiLCAvLyBwaW5pYSBcdTZFOTBcdTc4MDFcdTRFMkRcdTVGMTVcdTUxNjVcdTRFODYgdnVlLWRlbWkgXHU4RkQ5XHU0RTJBXHU1MzA1XG4gICAgICAgICAgfSksXG4gICAgICAgIF0sXG4gICAgICB9LFxuICAgIH0sXG4gIH07XG59KTtcbiJdLAogICJtYXBwaW5ncyI6ICI7QUFBd1UsU0FBUyxvQkFBb0I7QUFDclcsT0FBTyxTQUFTO0FBQ2hCLE9BQU8sZ0JBQWdCO0FBQ3ZCLFNBQVMsb0JBQW9CO0FBQzdCLE9BQU8scUJBQXFCO0FBQzVCLE9BQU8sY0FBYztBQUNyQixTQUFTLHlCQUF5QjtBQUNsQyxPQUFPLHFCQUFxQjtBQUM1QixPQUFPLHFCQUFxQjtBQUM1QixPQUFPLHVCQUF1QjtBQUM5QixPQUFPLFdBQVc7QUFHbEIsSUFBTyxzQkFBUSxhQUFhLENBQUMsRUFBRSxLQUFLLE1BQU07QUFDeEMsU0FBTztBQUFBLElBQ0wsU0FBUztBQUFBLE1BQ1AsSUFBSTtBQUFBLFFBQ0YsU0FBUyxDQUFDLFVBQVUsT0FBTztBQUFBO0FBQUEsTUFDN0IsQ0FBQztBQUFBLE1BQ0Qsa0JBQWtCO0FBQUEsUUFDaEIsT0FBTztBQUFBLFFBQ1AsT0FBTztBQUFBLE1BQ1QsQ0FBQztBQUFBLE1BQ0QsU0FBUztBQUFBO0FBQUE7QUFBQSxRQUdQLG1CQUFtQixFQUFFLE1BQU0sS0FBSztBQUFBO0FBQUEsUUFFaEMsZ0JBQWdCLElBQUk7QUFFbEIsYUFBRyxJQUFJLGVBQWU7QUFFdEIsYUFBRyxTQUFTLE1BQU0sUUFBUSxTQUFVLFFBQVEsS0FBSyxTQUFTLEtBQUssS0FBSztBQUNsRSxrQkFBTSxRQUFRLE9BQU8sR0FBRztBQUN4QixtQkFDRSxTQUNBLElBQUksWUFBWSxLQUFLLElBQ3JCLFlBRUEsTUFBTSxVQUNOO0FBQUEsVUFHSjtBQUNBLGFBQUcsSUFBSSxpQkFBaUIsRUFBRSxxQkFBcUIsS0FBSyxDQUFDO0FBQ3JELGFBQUcsSUFBSSxtQkFBbUI7QUFBQSxZQUN4QixRQUFRO0FBQUEsWUFDUixZQUFZLENBQUMsV0FBVyxTQUFTO0FBQUEsWUFDakMsY0FBYyxFQUFFLFFBQVEsRUFBRSxRQUFRLGNBQWMsRUFBRTtBQUFBLFVBQ3BELENBQUM7QUFBQSxRQUNIO0FBQUEsTUFDRixDQUFDO0FBQUEsTUFDRCxXQUFXO0FBQUEsUUFDVCxXQUFXLENBQUMsYUFBYSxDQUFDO0FBQUEsTUFDNUIsQ0FBQztBQUFBLElBQ0g7QUFBQSxJQUNBLE9BQU87QUFBQSxNQUNMLGVBQWU7QUFBQSxRQUNiLFVBQVU7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBO0FBQUEsUUFNVjtBQUFBLFFBQ0EsU0FBUztBQUFBLFVBQ1AsZ0JBQWdCO0FBQUE7QUFBQTtBQUFBO0FBQUE7QUFBQTtBQUFBLFVBTWhCLENBQUM7QUFBQSxRQUNIO0FBQUEsTUFDRjtBQUFBLElBQ0Y7QUFBQSxFQUNGO0FBQ0YsQ0FBQzsiLAogICJuYW1lcyI6IFtdCn0K
