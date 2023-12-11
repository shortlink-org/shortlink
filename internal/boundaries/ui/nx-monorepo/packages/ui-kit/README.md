## @shortlink/ui-kit

The UI Kit is a collection of React components that are used in the Shortlink app.

### Installation

```bash
npm install @shortlink/ui-kit # for GitLab registry
npm i shortlink-ui-kit        # for NPM registry
```

### Usage

```jsx
import { ToggleDarkMode } from '@shortlink/ui-kit'

const App = () => {
  return <ToggleDarkMode />
}
```

### Storybook

The UI Kit is documented with Storybook. You can run it locally with:

```bash
npm run storybook
```

### Stack

- ReactJS
- Pretty code base
  - Typescript
  - ESLint/Prettier
- Pretty UI
  - TailwindCSS
- Storybook (for UI)
- Jest
