import {nodeResolve} from "@rollup/plugin-node-resolve"
import commonjs from "@rollup/plugin-commonjs"
import typescript from "@rollup/plugin-typescript"
import {terser} from "rollup-plugin-terser"
import peerDepsExternal from 'rollup-plugin-peer-deps-external'
import postcss from 'rollup-plugin-postcss'
import dts from "rollup-plugin-dts"
import babel from 'rollup-plugin-babel'

const packageJson = require("./package.json")

export default [
  {
    input: "src/index.ts",
    output: [
      {
        file: packageJson.main,
        format: "cjs",
        sourcemap: true,
      },
      {
        file: packageJson.module,
        format: "esm",
        sourcemap: true,
      },
    ],
    plugins: [
      nodeResolve(),
      peerDepsExternal(),
      postcss({
        extract: true,
      }),
      babel({
        exclude: 'node_modules/**',
      }),
      commonjs(),
      typescript({ tsconfig: "./tsconfig.json" }),
      terser()
    ],
    external: ["react", "react-dom", "styled-components", "next-themes", "@emotion/cache", "@mui/material"]
  },
  {
    input: "dist/esm/types/index.d.ts",
    output: [{ file: "dist/index.d.ts", format: "es" }],
    plugins: [
      postcss({
        extract: 'styles.css', // this will generate a specific file not being used, but we need this part of code
        autoModules: true,
        include: '**/*.css',
        extensions: ['.css'],
        plugins: [],
      }),
      dts({
        compilerOptions: {
          baseUrl: 'src',
        },
      }),
    ],
  },
];
