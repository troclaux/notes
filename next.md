# Next.js

> front-end react framework that enables server-side rendering and static site generation

[SSR](/web_development.md#server-side-rendered-application-ssr)

- main features:

  - routing: how users navigate between different parts of your application
    - route: URL path that maps to a page route or to an api route
    - file-based routing
    - every file inside `app/` is a route (`app/index.tsx` => `/`)
    - dynamic routes: `app/workout/[id].tsx` => `/workout/:id`
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
- would you like to use eslint? Yes
- would you like to use Tailwind CSS? Yes
- would you like to use src/ directory? Yes
- would you like to use App Router? Recommended Yes
- would you like to use Turbopack for `next dev`? › Yes
- would you like to customize the default import alias (@/\*)? No

4. wait for the installation to finish
5. `npm run dev` to start the development server
6. start editing `src/pages/index.tsx` to see the changes in the browser
7. install basic dependencies

```bash
npm install <dependency>
```

- `clsx`: utility to conditionally apply css classes
- `jsonwebtoken`: create and verify tokens
- `bcrypt`: password hashing
- `pg`: postgresql client
- `cookies-next`: cookie management
- `tailwind-merge`: merge tailwind classes without style conflicts

- if auth provider is `better-auth`:
  - `lucide-react`: customizable icons
  - `tailwindcss-animate`: animations for tailwind
  - `sonner`: sound effects
  - `npx shadcn@latest init -d` [see the documentation before adding](https://ui.shadcn.com): beautiful components
- if auth provider is `next-auth@beta`:
  - `npx auth secret`: create a secret key
- if auth provider is `@clerk/nextjs`:
  - `@clerk/nextjs`

## directory structure

> next.js uses a file-system based routing system

```
my-app/
├── src/
│   ├── app/
│   │   ├── api/
│   │   │   └── items/
│   │   │       ├── route.ts       # Handles GET, POST for /api/items
│   │   │       └── [id]/
│   │   │           └── route.ts   # Handles GET, PUT, DELETE for /api/items/:id
│   ├── lib/
│   │   └── db.ts                  # DB mock or DB logic
├── next.config.js
├── package.json
├── tsconfig.json
```

- `src/`: optional application source folder
  - `src/app/` (new and recommended): app router, contains `api` (back end) and each page directory
    - endpoint vs route: you can have different endpoints for the same route
      - endpoint: where your API receives and sends data
        - e.g. `GET localhost:8080/employees/42`, `DELETE localhost:8080/employees/42`
      - route: URL that maps to a page in your website (like `/about` or `/products`)
        - e.g. `localhost:8080/employees/42`
    - uses nested folders to define routes
    - `src/app/page.tsx`: default page for root route (`/`)
      - equivalent to `index.tsx` in the pages router (`src/pages/`)
    - `src/app/[route]/page.tsx`: page component, `[route]` defines the route and contains the front end
      - e.g. `src/app/products/page.tsx`
    - `src/app/[route]/layout.tsx`: layout component
      - layout: ui that is shared between multiple pages
    - `/app/[route]/loading.tsx`: loading indicator or a fallback ui while a specific route segment is loading
      - route segment: portions of the url path that corresponds to a directory or file within the `app` directory
      - `<Suspense>` component: wraps the loading component
        - use `fallback={<p>Loading...</p>}` prop to handle cases where a page's content isn't available at build time
    - `src/app/lib/`: contains functions used in the application
      - e.g. reusable utility functions and data fetching functions
    - `src/app/api/`: stores all api routes
      - e.g. `src/app/api/users/route.ts`: contains all user api routes
    - `src/app/ui/`: contains all ui components for the application
  - `src/pages/` (old, not recommended): pages router, server-side
    - files directly represent routes
- `public/`: static assets such as images, fonts, etc
  - can be referenced by your code starting from base URL (`/`)
- top-level files:
  - `next.config.js`: config file for next.js application
    - customize build and runtime behaviour (e.g. enable experimental features, modify webpack configurations, etc)
  - `package.json`:
  - `middleware.ts`:
  - `.env`: file that stores environment variables at the root of the project
    - examples of env vars:
      - `DATABASE_HOST`: database credentials for connection pool
      - `DATABASE_NAME`: database credentials for connection pool
      - `DATABASE_USER`: database credentials for connection pool
      - `DATABASE_PASSWORD`: database credentials for connection pool
      - `GOOSE_DBSTRING`: allows database migration with goose
      - `GOOSE_DRIVER`: allows database migration with goose
      - `AUTH_SECRET`: used to encrypt session tokens in Auth.js
      - `NEXTAUTH_URL`: sets the canonical base URL of your application (e.g. `pesodevops.com`)
        - used for OAuth callbacks, redirect URLs, email links and CSRF protection
      - `AUTH_TRUST_HOST`: when set to `true`, allows Auth.js to trust the `NEXTAUTH_URL` in production if it's not `localhost`
        - used for deployment
      - `AUTH_GOOGLE_ID`: enables OAuth/OIDC authentication
      - `AUTH_GOOGLE_SECRET`: enables OAuth/OIDC authentication
      - `AWS_ACCESS_KEY_ID`: aws credentials for terraform
      - `AWS_SECRET_ACCESS_KEY`: aws credentials for terraform
      - `ECR_REPO_URL`
      - `EC2_PUBLIC_IP`
    - if the same variable is defined in multiple `.env` files (`.env.local`, `.env.production`, `.env.development`), the final value will be the one in the file with highest priority
    - `.env.local`: highest priority
    - `.env.production`and `.env.development`: medium priority
    - `.env`: lowest priority
  - `.env.developent` (recommended at the beggining): developent environment variables
  - `.env.production` (use for production): production environment variables
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
> routes are not accessible until a `page.(js/ts/tsx)` or `route.(js/ts/tsx)` is added to the route segment\*

> [!TIP]
> try to use `app` folder purely for routing purposes

- separate `client` and `server` folders vs single next.js application handling both frontend and api routes (recommended)
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
  { name: "Home", href: "/dashboard", icon: HomeIcon },
  {
    name: "Invoices",
    href: "/dashboard/invoices",
    icon: DocumentDuplicateIcon,
  },
  { name: "Customers", href: "/dashboard/customers", icon: UserGroupIcon },
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
          "flex h-[48px] grow items-center justify-center gap-2 rounded-md bg-gray-50 p-3 text-sm font-medium hover:bg-sky-100 hover:text-blue-600 md:flex-none md:justify-start md:p-2 md:px-3",
          {
            "bg-sky-100 text-blue-600": pathname === link.href,
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
clsx("button", isActive && "active");
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

| page                 | route              | result      |
| -------------------- | ------------------ | ----------- |
| `app/page.ts`        | `app/route.ts`     | ❌ conflict |
| `app/page.ts`        | `app/api/route.ts` | ✅ valid    |
| `app/[user]/page.ts` | `app/api/route.ts` | ✅ valid    |

- each `route.ts` or `page.ts` file must handle all http verbs for that route
  - i.e. you can't have some http methods handled in one file and others in another file for the same route

## http verbs

> a route file allows you to create custom request handlers for a given route

```ts
import type { NextRequest } from "next/server";

export async function GET(request: NextRequest) {
  const url = request.nextUrl;
  return new Response(`URL: ${url}`, { status: 200 });
}

export async function POST(request: NextRequest) {
  const body = await request.json();
  return new Response(`Received: ${JSON.stringify(body)}`, { status: 200 });
}

export async function PUT(request: NextRequest) {
  const body = await request.json();
  return new Response(`Updated: ${JSON.stringify(body)}`, { status: 200 });
}

export async function DELETE(request: NextRequest) {
  const body = await request.json();
  return new Response(`Deleted: ${JSON.stringify(body)}`, { status: 200 });
}
```

- `NextRequest`: object that gives access to request information like headers, cookies, URL parameters, and body data
  - `request.nextUrl`: URL's request
- `request.json()`: This method parses the request body as JSON
  - must be used inside an `async` function
- `NextResponse`: object that allows you to create custom responses (e.g. status code, headers, etc)
  - `return NextResponse.json({ error: 'Internal server error' }, { status: 500 });`

example:

```typescript
return new Response('Hello', {
  status: 200,
  headers: { 'Content-Type': 'text/plain' }
})
```

## client-side hooks

> next.js has come client-side hooks that allow you to interact with the browser

- `useSearchParams`: convert url parameters of the current url to json
- `usePathName`: get current URL's pathname
- `useRouter`: hook that enables client-side navigation and access to router information (current route, query parameters, etc)
- `useSession`: hook that provides access to the current user's session data (e.g. user id, email, etc)
  - available only client-side

## Client-Side Rendering vs Server-Side Rendering

- components in the `app/` folder are server components by default

- `"use client"`: add it to the top of a file before all `import` statements to convert component into a client component
  - required to use client-side react features, such as `useState`, `useEffect`, event handlers, etc
  - marks a file and all its dependencies as part of the client bundle
  - all child components and imports are also considered client components
- `"use server"`: marks all exported functions as server actions
  - enables "server actions"
    - server actions: server functions that can be called directly from client components

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
````

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
- partial prerendering: combination of the benefits of static and dynamic rendering in the same route
  - use case: pages that have some content that can be pre-rendered at build time and some dynamically generated content

---

- route segment: portions of the url path that corresponds to a directory or file within the `app` directory
- segments: each part of the url path separated by a forward slash (`/`)
- dynamic segments (like `[id]`): create routes that handle variable parts of the url
- hidration: process of making server-rendered html interactive on the client-side

---

## implementation

### add a new route

### add auth

### add layout

### get users data client-side

### add route that blocks unauthenticated users

### test api routes with authentication

1. login to get a session token
2. inspect => application => cookies => copy authjs.session-token
3. use the token in the request to test the api route

```bash
curl -X POST http://localhost:3000/api/exercises \
  -H "Content-Type: application/json" \
  -H "Cookie: authjs.session-token=<your-session-token>" \
  -d '{
    "name": "Push-up",
    "description": "A basic exercise for upper body strength"
  }'
```

---

## questions

- [x] how do i make http requests?
  - just create the routes in the `app` folder
  - [example](https://github.com/troclaux/peso/tree/main/src/app/api/users)
- [x] Should i use express.js or does next.js have a built-in solution for this?
  - no need to install express.js
  - Next.js has a built-in API routing system that can replace express.js
- [x] what is the best method to fetch data from the database?
  - in my opinion, the best way to fetch data is to use `pg`, without any ORM
- [x] if i have a route for a resource, how do i use the routing system to create endpoints for a specific resource (using the id)?
  - should all endpoints be in one file?
    - no, each resource (`/workouts`) has a folder, and the id-based routes (`/workouts/:id`) go inside `[id].ts` file
    - `route.ts`: handle GET all and POST new
    - `[id].ts`: handle GET 1, PUT 1 and DELETE 1

```
app/
  api/
    workouts/
      route.ts      → Handles `/api/workouts` (GET all, POST new)
      [id].ts → Handles `/api/workouts/:id` (GET, PUT, DELETE)
```

- [x] ensure database connection is always released
  - problem: if an error occurs before client.release(), the connection is not released
  - fix: use a try-finally block
- [x] Use Secure HTTP-Only Cookies for Tokens
  - Problem: Tokens are sent in JSON response, making them vulnerable to XSS
  - Fix: Use `NextResponse.cookies.set()` for storing tokens securely
- [x] Return Only the Access Token in Response
  - Why? Refresh tokens should not be sent to the frontend. They should be stored in HTTP-only cookies
