[SSR](/web_development.md#server-side-rendered-application-ssr)

# Next.js

> front-end react framework that enables server-side rendering and static site generation

- main features:
  - routing: how users navigate between different parts of your application
    - route: URL path that maps to a page route or to an api route
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
  - streaming: data transfer technique that breaks down a route into smaller chunks and progressively transfers them to the client
    - prevents slow data request from blocking the entire page

- basic `npm` commands:
  - update dependencies: `npm update`
  - **only** update next.js, react.js and react-dom: `npm install next@latest react@latest react-dom@latest`
  - run next.js application in development mode: `npm run dev`
  - start production server: `npm run start`
  - run tests: `npm test`

setting scripts in `package.json`:

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
    - `/app/[route]/loading.tsx`: loading indicator or a fallback ui while a specific route segment is loading
      - route segment: portions of the url path that corresponds to a directory or file within the `app` directory
      - `<Suspense>` component: wraps the loading component
        - use `fallback={<p>Loading...</p>}` prop to handle cases where a page's content isn't available at build time
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

> [!TIP]
> try to use `app` folder purely for routing purposes


- separate `client` and `server` folders vs single next.js application handling bothe frontend and api routes (recommended)
  - separate `client` and `server` folders:
    - advantages
      - separation of concerts
      - independent development
  - single next.js application handling both frontend and api routes (recommended)
    - advantages
      - simplicity: no need to manage separate express server
      - integrated routing: next.js handles both page and api routing

### route groups

> organize routes into groups e.g. by site section, intent or team

- created by wrapping a folder in parenthesis: `(foldername)`
  - examples:
    - `src/app/(admin)/dashboard/page.ts` file path => `/dashboard` URL path
    - `src/app/(marketing)/about/page.ts` file path => `/about` URL path
    - `src/app/(marketing)/blog/page.ts` file path => `/blog` URL path

### component hierarchy

- root layout (`layout.tsx`): wraps all pages
  - contains persistent components (e.g. navbar, footer)
- page components (`page.tsx`): represents individual pages (e.g., `/workouts`, `/workouts/[id]`)
  - fetches data and passes it down to child components
- ui components (`components/`): reusable components

### data flow in hierarchy

- top-down (props): parent components pass data down to child components through props
- bottom-up (callbacks/state management): child components communicate with parents through callback functions or state management libraries like Redux

## navigation

- use `<Link href="/users"></Link>` tag instead of `<a href="/users"></a>` to optimize page navigation
- `import Link from "next/link";`
- `href` (Hypertext REFerence): destination URL/path that the link will navigate to

### add navigation buttons

1. create an array of navigation links

```typescript
const links = [
  { name: 'Home', href: '/dashboard', icon: HomeIcon },
  { name: 'Invoices', href: '/dashboard/invoices', icon: DocumentDuplicateIcon },
  { name: 'Customers', href: '/dashboard/customers', icon: UserGroupIcon },
];
```

2. use `map` to create a link for each

```tsx
<>
  {links.map((link) => {
    const LinkIcon = link.icon;
    return (
      <Link
        key={link.name}
        href={link.href}
        className={clsx(
          'flex h-[48px] grow items-center justify-center gap-2 rounded-md bg-gray-50 p-3 text-sm font-medium hover:bg-sky-100 hover:text-blue-600 md:flex-none md:justify-start md:p-2 md:px-3',
          {
            'bg-sky-100 text-blue-600': pathname === link.href,
          },
        )}
      >
        <LinkIcon className="w-6" />
        <p className="hidden md:block">{link.name}</p>
      </Link>
    );
  })}
</>
```

- [Tailwind](./tailwind.md) CSS classes explanation (`hidden md:block`):
  - `hidden`: element is not visible by default (display: none)
  - `md:block`: on medium screens (768px and up), element becomes visible (display: block)
    - on medium screens and up: text is visible

- `clsx`: utility to conditionally apply css classes
  - `clsx('<css classes 1>', {'<css classes 2>': pathname === link.href,})`
    - `'<css classes 1>'`: css classes 1 that are always applied
    - `'<css classes 2>'`: if the condition is true (`pathname === link.href`), the css classes 2 are applied

more examples:

```tsx
// Conditional classes
clsx('button', isActive && 'active')
// If isActive is true: "button active"
// If isActive is false: "button"
```

## route handling

- route handlers are defined in a `route.js|ts` file inside the `app` directory
  - there cannot be a `route.js|ts` file at the same route segment level as `page.js|ts`
  - not cached by default
    - each request triggers a fresh database query, rather than serving a stored (cached) response

```
app/
  blog/
    page.js      ❌
    route.js     ❌

app/
  blog/
    page.js      ✅ (handles the UI)
    api/
      route.js   ✅ (handles API requests)
```

| page | route | result |
| --------------- | --------------- | --------------- |
| `app/page.ts` | `app/route.ts` | ❌ conflict |
| `app/page.ts` | `app/api/route.ts` | ✅ valid |
| `app/[user]/page.ts` | `app/api/route.ts` | ✅ valid |

- each `route.ts` or `page.ts` file must handle all http verbs for that route
  - i.e. you can't have some http methods handled in one file and others in another file for the same route

## http verbs

> a route file allows you to create custom request handlers for a given route

```ts
import type { NextRequest } from 'next/server'

export async function GET(request: NextRequest) {
  const url = request.nextUrl
  return new Response(`URL: ${url}`, { status: 200 })
}

export async function POST(request: NextRequest) {
  const body = await request.json()
  return new Response(`Received: ${JSON.stringify(body)}`, { status: 200 })
}

export async function PUT(request: NextRequest) {
  const body = await request.json()
  return new Response(`Updated: ${JSON.stringify(body)}`, { status: 200 })
}

export async function DELETE(request: NextRequest) {
  const body = await request.json()
  return new Response(`Deleted: ${JSON.stringify(body)}`, { status: 200 })
}
```

- `NextRequest`: object that gives access to request information like headers, cookies, URL parameters, and body data
  - `request.nextUrl`: URL's request
- `request.json()`: This method parses the request body as JSON
  - must be used inside an `async` function
- `Response(body, options)`: Web API object representing an HTTP response, allowing you to create custom responses with:
  - response body
  - status code
  - headers
  - other response properties

Example:

```typescript
return new Response('Hello', {
  status: 200,
  headers: { 'Content-Type': 'text/plain' }
})

---

## Client-Side Rendering vs Server-Side Rendering

> [!IMPORTANT]
> components in the `app/` folder are server components by default
> to convert a server component into a Client component, add this to the beggining of the file: `'use client';`


> [!TIP]
> a project usually uses a mix of both types
> try to maximize the use of server components and minimize client components

### Client-Side Rendering (CSR)

- client-side rendering is the most common way to render a web page
- the browser downloads a minimal HTML page, then downloads and executes JavaScript that renders the page
- the JavaScript can make requests to the server to fetch data and update the page

- client component: components that are rendered in the user's browser
  - bigger payload size

### Server-Side Rendering (SSR)

- server-side rendering is when the server generates the HTML for the page and sends it to the browser
- the browser receives a fully rendered page, which can be faster to load and better for SEO
- server-side rendering can be more complex to set up and maintain

- server component (default): components that are rendered in the server
  - smaller payload size
  - improves SEO (Search Engine Optimization)
  - cannot
    - listen to browser events (e.g. click, change, submit, etc)
    - access browser apis (e.g. localstorage)
    - maintain state
    - use effects

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

## static vs dynamic rendering

- static rendering: render at build time
  - use case: UI with no data OR with data shared across users
  - faster
  - easier caching
  - reduce server load
  - better SEO
- dynamic rendering: render page in response to a user's request
  - use case: all cases that don't fit static rendering
    - user-specific content
    - real-time data

---

- route segment: portions of the url path that corresponds to a directory or file within the `app` directory
- segments: each part of the url path separated by a forward slash (`/`)
- dynamic segments (like `[id]`): create routes that handle variable parts of the url
- hidration: process of making server-rendered html interactive on the client-side

---

## questions

- how do i make http requests?
  - do i use express.js or does next.js have a built-in solution for this?

- best way to fetch data?
- how do i make http requests with next.js for postgresql?
