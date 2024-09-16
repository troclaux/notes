
> react framework that enables server-side rendering and static site generation

# Client-Side Rendering vs Server-Side Rendering

## Client-Side Rendering (CSR)

- client-side rendering is the most common way to render a web page
- the browser downloads a minimal HTML page, then downloads and executes JavaScript that renders the page
- the JavaScript can make requests to the server to fetch data and update the page

## Server-Side Rendering (SSR)

- server-side rendering is when the server generates the HTML for the page and sends it to the browser
- the browser receives a fully rendered page, which can be faster to load and better for SEO
- server-side rendering can be more complex to set up and maintain

# Static Site Generation (SSG)

> when the server generates the HTML for the page at build time

- build time: the time when your application is being built and prepared for deployment
  - Next.js pre-renders your pages to HTML during this phase
  - each page is generated once at build time and reused on each request
  - this allows for very fast page loads and SEO benefits
  - server can serve static files instead of executing server-side code for each request

to use SSG, you export an async function called getStaticProps in a page:

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
