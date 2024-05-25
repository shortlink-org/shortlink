import commonjs from '@rollup/plugin-commonjs'
import { nodeResolve } from '@rollup/plugin-node-resolve'
import terser from '@rollup/plugin-terser'
import ts from '@rollup/plugin-typescript'
import svgr from '@svgr/rollup'
import filesize from 'rollup-plugin-filesize'
import peerDepsExternal from 'rollup-plugin-peer-deps-external'
import postcss from 'rollup-plugin-postcss'
import typescript from 'typescript'
// import dts from 'rollup-plugin-dts'

const packageJson = require('./package.json')

export default [
  {
    input: 'src/index.ts',
    output: [
      {
        name: packageJson.name,
        file: packageJson.main,
        format: 'umd',
        sourcemap: true,
      },
      {
        file: packageJson.module,
        format: 'esm',
        sourcemap: true,
      },
    ],
    plugins: [
      peerDepsExternal(), // Preferably set as first plugin.
      svgr(),
      postcss({
        extract: 'index.css',
        autoModules: true,
        include: 'src/**/*.css',
        extensions: ['.css'],
        plugins: [],
      }),
      nodeResolve(),
      ts({
        typescript,
        tsconfig: './tsconfig.json',
        noEmitOnError: false,
        declaration: true,
      }),
      commonjs({
        include: ['node_modules/**', '../../node_modules/**'],
      }),
      filesize(),
      terser(),
    ],
    external: [
      'react',
      'react-dom',
      'prop-types',
      '@emotion/react',
      '@mui/styled-engine',
      '@mui/system',
      '@mui/material',
      '@mui/material/styles',
      '@mui/x-date-pickers',
      '@mui/icons-material',
    ],
  },
  // {
  //   name: packageJson.name,
  //   input: "dist/esm/types/index.d.ts",
  //   output: [{ file: "dist/index.d.ts", format: "es" }],
  //   plugins: [
  //     postcss({
  //       extract: 'styles.css', // this will generate a specific file not being used, but we need this part of code
  //       autoModules: true,
  //       include: '**/*.css',
  //       extensions: ['.css'],
  //       plugins: [],
  //     }),
  //     dts.default({
  //       compilerOptions: {
  //         baseUrl: 'src',
  //       },
  //     }),
  //   ],
  //   external: ['react', 'react-is', 'react-dom', 'prop-types', "@mui/material", "@emotion/react"],
  // },
]
