import { defineConfig } from "vite";
import Vue from "@vitejs/plugin-vue";
import Components from "unplugin-vue-components/vite";
import { ArcoResolver } from "unplugin-vue-components/resolvers";
import externalGlobals from "rollup-plugin-external-globals";
import Markdown from "vite-plugin-vue-markdown";

// https://vitejs.dev/config/
export default defineConfig(({ mode }) => {
  return {
    plugins: [
      Vue({
        include: [/\.vue$/, /\.md$/], // <--
      }),
      Markdown({
        // default options passed to markdown-it
        // see: https://markdown-it.github.io/markdown-it/
        markdownItOptions: { html: true },
        // A function providing the Markdown It instances gets the ability to apply custom settings/plugins
        markdownItSetup(md) {
          // add code syntax highlighting with Prism
          md.use(require("markdown-it-attrs"));
          // custom renderer for fences
          md.renderer.rules.fence = function (tokens, idx, options, env, slf) {
            const token = tokens[idx];
            return (
              "<pre" +
              slf.renderAttrs(token) +
              ">" +
              "<code>" +
              token.content +
              "</code>" +
              "</pre>"
            );
          };
          md.use(require("markdown-it-prism"), { highlightInlineCode: true });
          md.use(require("markdown-it-texmath"), {
            engine: require("katex"),
            delimiters: ["dollars", "beg_end"],
            katexOptions: { macros: { "\\RR": "\\mathbb{R}" } },
          });
        },
      }),
      Components({
        resolvers: [ArcoResolver()],
      }),
    ],
    build: {
      rollupOptions: {
        external: [
          "vue",
          "vue-router",
          //     "@arco-design/web-vue",
          "axios",
          //     "vue-demi",
        ],
        plugins: [
          externalGlobals({
            vue: "Vue",
            "vue-router": "VueRouter",
            //         "@arco-design/web-vue": "ArcoVue",
            axios: "axios",
            //         "vue-demi": "VueDemi", // pinia 源码中引入了 vue-demi 这个包
          }),
        ],
      },
    },
  };
});
