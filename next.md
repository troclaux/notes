[SSR](/web_development.md#server-side-rendered-application-ssr)

# Next.js

> front-end react framework that enables server-side rendering and static site generation

- main features:
  - routing: how users navigate between different parts of your application
    - route: a path that maps to a page component and determines what content is displayed
    - file-based routing
    - every file inside `pages/` is a route (`pages/index.tsx` => `/`)
    - dynamic routes: `pages/workout/[id].tsx` => `/workout/:id`
    - route vs endpoint:
      - route: a page in your website (like `/about` or `/products`)
        - what users see and interact with in the browser
        - example: `/dashboard` shows your dashboard page
      - endpoint: where your API receives and sends data
        - handles data operations in the background
        - example: `/api/users` might get or update user data
  - Server-side Rendering (SSR) vs Static Site Generation (SSG)
    - use `getServerSideProps` for dynamic, always-updated data
    - use `getStaticProps` for pre-generated pages (faster)
  - typescript support
  - allows for optimization of images and assets

- basic `npm` commands:
  - update dependencies: `npm update`
  - **only** update next.js, react.js and react-dom: `npm install next@latest react@latest react-dom@latest`
  - run next.js application in development mode: `npm run dev`
  - start production server: `npm run start`
  - run tests: `npm test`

in `package.json`, you can set scripts:

```json
{
  "scripts": {
    "dev": "next dev",
    "build": "next build",
    "start": "next start",
    "lint": "next lint"
  }
}
```

- with this `package.json`:
  - view application by running dev script: `npm run dev`
    - by default, next.js uses `http://localhost:3000`
  - to run build script: `npm run build`
  - to run start script: `npm run start`
  - to run lint script: `npm run lint`

## starting a new project

1. create and clone github repository
2. initialize a new Next.js application in the current directory: `npx create-next-app@latest .`
3. answer the questions
  - would you like to use TypeScript? Yes
  - would you like to use eslint? No
  - would you like to use Tailwind CSS? Yes
  - would you like to use src/ directory? Yes
  - would you like to use App Router? Recommended Yes
  - would you like to use Turbopack for `next dev`? › Yes
  - would you like to customize the default import alias (@/*)? No
4. wait for the installation to finish
5. `npm run dev` to start the development server
6. start editing `src/pages/index.tsx` to see the changes in the browser

## directory structure

next.js uses a file-system based routing system

- `src/`: optional application source folder
  - `src/app/` (new and recommended): app router, client-side
    - uses nested folders to define routes
    - `src/app/page.tsx`: default page for root route (`/`)
      - equivalent to `index.tsx` in the pages router (`src/pages/`)
    - `src/app/[route]/page.tsx`: page component, `[route]` defines a route segment?
      - e.g. `src/app/products/page.tsx`
    - `src/app/[route]/layout.tsx`: layout component
      - layout: ui that is shared between multiple pages
    - `src/app/lib`: contains functions used in the application
      - e.g. reusable utility functions and data fetching functions
    - `src/app/ui`: contains all ui components for the application
  - `src/pages/` (traditional): pages router, server-side
    - files directly represent routes
  - `src/pages/api/`: create backend logic inside a frontend project
- `public/`: static assets such as images, fonts, etc
  - can be referenced by your code starting from base URL (`/`)
- top-level files:
  - `next.config.js`: config file for next.js application
    - customize build and runtime behaviour (e.g. enable experimental features, modify webpack configurations, etc)
  - `package.json`: 
  - `middleware.ts`: 
  - `.env`: file that stores environment variables at the root of the project
  - `.env.local`: local environment variables
  - `.env.production`: production environment variables
  - `.env.developent`: developent environment variables
  - `eslintrc.json`: config file for eslint
  - `tsconfig.json`: config file for typescript
  - `jsconfig.json`: config file for javascript
- routing files:
  - `layout.(js/ts/tsx)`: shared UI that wraps multiple pages (like headers/navigation)
  - `page.(js/ts/tsx)`: the actual page content and UI for a route
  - `loading.(js/ts/tsx)`: loading state shown while page content loads
  - `not-found.(js/ts/tsx)`: custom 404 page when a route doesn't exist
  - `error.(js/ts/tsx)`: error UI shown when a page has an error
  - `global-error.(js/ts/tsx)`: error UI for errors outside pages
  - `route.(js/ts/tsx)`: API endpoints and request handlers
  - `template.(js/ts/tsx)`: similar to layout but creates new instance on navigation
  - `default.(js/ts/tsx)`: fallback page for parallel routes

> [!WARNING]
> routes are not accessible until a `page.(js/ts/tsx)` or `route.(js/ts/tsx)` is added to the route segment*

- separate `client` and `server` folders vs single next.js application handling bothe frontend and api routes (recommended)
  - separate `client` and `server` folders:
    - advantages
      - separation of concerts
      - independent development
  - single next.js application handling both frontend and api routes (recommended)
    - advantages
      - simplicity: no need to manage separate express server
      - integrated routing: next.js handles both page and api routing

### component hierarchy

- root layout (`layout.tsx`): wraps all pages
  - contains persistent components (e.g. navbar, footer)
- page components (`page.tsx`): represents individual pages (e.g., `/workouts`, `/workouts/[id]`)
  - fetches data and passes it down to child components
- ui components (`components/`): reusable components

### data flow in hierarchy

- top-down (props): parent components pass data down to child components through props
- bottom-up (callbacks/state management): child components communicate with parents through callback functions or state management libraries like Redux

## app router

- client component: 
  - server-sede redering
- server component: 
  - server-sede redering
  - hidration: process of making pre-rendered html interactive on the client-side

---

## Client-Side Rendering vs Server-Side Rendering

### Client-Side Rendering (CSR)

- client-side rendering is the most common way to render a web page
- the browser downloads a minimal HTML page, then downloads and executes JavaScript that renders the page
- the JavaScript can make requests to the server to fetch data and update the page

### Server-Side Rendering (SSR)

- server-side rendering is when the server generates the HTML for the page and sends it to the browser
- the browser receives a fully rendered page, which can be faster to load and better for SEO
- server-side rendering can be more complex to set up and maintain

## Static Site Generation (SSG)

> pre-render pages at build time for performance

- build time: the time when your application is being built and prepared for deployment
  - Next.js pre-renders your pages to HTML during this phase
  - each page is generated once at build time and reused on each request
  - this allows for very fast page loads and SEO benefits
  - server can serve static files instead of executing server-side code for each request

to use SSG, you export an async function called `getStaticProps` in pages:

```javascript
// pages/index.js
export async function getStaticProps() {
  return {
    props: {
      message: 'Hello from static generation!'
    }
  };
}

export default function Home({ message }) {
  return <h1>{message}</h1>;
}
```

---

- route segment: portions of the url path that corresponds to a directory or file within the `app` directory
- segments: each part of the url path separated by a forward slash (`/`)
- dynamic segments (like `[id]`): create routes that handle variable parts of the url
