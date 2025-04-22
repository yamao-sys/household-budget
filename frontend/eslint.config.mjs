import globals from "globals";
import reactHooks from "eslint-plugin-react-hooks";
import reactRefresh from "eslint-plugin-react-refresh";
import reactPlugin from "eslint-plugin-react";
import importPlugin from "eslint-plugin-import";
import { dirname } from "path";
import { fileURLToPath } from "url";
import { FlatCompat } from "@eslint/eslintrc";
import eslintConfigPrettier from "eslint-config-prettier";
import * as fs from "fs";
import typescriptParser from "@typescript-eslint/parser";
import typescriptPlugin from "@typescript-eslint/eslint-plugin";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);

const compat = new FlatCompat({
  baseDirectory: __dirname,
});

// const domains = fs.readdirSync('./features');
// const zones = domains.map((domain) => ({
//   from: `./features/${domain}/!(public)/**/*`,
//   target: `./features/!(${domain})/**/*`,
// }));

const eslintConfig = [
  ...compat.extends("plugin:prettier/recommended"),
  reactPlugin.configs.flat.recommended,
  reactPlugin.configs.flat["jsx-runtime"],
  {
    settings: {
      react: {
        version: "detect",
      },
    },
  },
  {
    rules: {
      // NOTE: Prettierとの競合を防ぐ設定を追加
      //     : prettierConfigからルールを追加
      ...eslintConfigPrettier.rules,

      // NOTE: features配下はそれぞれが専属で使用するため、他のドメインのディレクトリからimport不可にする
      // 'import/no-restricted-paths': [
      //   'error',
      //   {
      //     zones,
      //   },
      // ],
      "react/jsx-uses-react": "error",
      "react/jsx-uses-vars": "error",
      "@typescript-eslint/no-unused-vars": "warn",
    },
  },
  { files: ["**/*.{js,mjs,cjs,ts,jsx,tsx}"] },
  {
    languageOptions: {
      parser: typescriptParser,
      parserOptions: {
        ecmaVersion: 2022,
        sourceType: "module",
        project: "./tsconfig.json",
        tsconfigRootDir: __dirname,
      },
      globals: globals.browser,
    },
  },
  {
    plugins: {
      "react-hooks": reactHooks,
      "react-refresh": reactRefresh,
      import: importPlugin,
      "@typescript-eslint": typescriptPlugin,
    },
  },
  { ignores: ["features/**/__generated__/**/*.ts", ".react-router/**/*.ts"] },
];

export default eslintConfig;
