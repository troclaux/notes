
# tailwind

> css framework that speeds up the development process

## basics

- compose designs with pre-built utility classes
- no need to write custom css for most designs
- responsive design with prefixes:
  - `sm:`: applies styles at small screen sizes (640px and above)
  - `md:`: applies styles at medium screen sizes (768px and above)
  - `lg:`: applies styles at large screen sizes (1024px and above)
  - `xl:`: applies styles at extra large screen sizes (1280px and above)
  - `2xl:`: applies styles at 2x extra large screen sizes (1536px and above)
- dark mode with `dark:` prefix
- hover states with `hover:` prefix

style elements by adding class names:

```html
<h1 className="text-blue-500 hover:text-blue-700 md:text-xl">I'm blue!</h1>
```

## essential utility classes

### typography

- `text-{size}`: font size (`text-xs`, `text-sm`, `text-base`, `text-lg`, `text-xl`, `text-2xl`, etc.)
- `font-{weight}`: font weight (`font-normal`, `font-medium`, `font-bold`)
- `text-{color}`: text color (`text-blue-500`, `text-red-600`, `text-gray-800`)
- `text-{alignment}`: text alignment (`text-left`, `text-center`, `text-right`)
- `tracking-{spacing}`: letter spacing (`tracking-tight`, `tracking-normal`, `tracking-wide`)
- `leading-{height}`: line height (`leading-none`, `leading-tight`, `leading-normal`, `leading-loose`)

### spacing

- `m-{size}`: margin on all sides (`m-0`, `m-1`, `m-2`, `m-4`, `m-8`)
- `mx-{size}`: horizontal margin (left and right)
- `my-{size}`: vertical margin (top and bottom)
- `mt-{size}`, `mr-{size}`, `mb-{size}`, `ml-{size}`: margin top, right, bottom, left
- `p-{size}`: padding on all sides (`p-0`, `p-1`, `p-2`, `p-4`, `p-8`)
- `px-{size}`: horizontal padding (left and right)
- `py-{size}`: vertical padding (top and bottom)
- `pt-{size}`, `pr-{size}`, `pb-{size}`, `pl-{size}`: padding top, right, bottom, left

### layout

- `flex`: display flex
- `grid`: display grid
- `block`: display block
- `hidden`: display none
- `inline`, `inline-block`, `inline-flex`: display inline variants
- `container`: container with responsive max-width
- `w-{size}`: width (`w-full`, `w-1/2`, `w-1/3`, `w-1/4`, `w-auto`)
- `h-{size}`: height (`h-full`, `h-screen`, `h-64`, `h-auto`)
- `max-w-{size}`: max width (`max-w-sm`, `max-w-md`, `max-w-lg`, `max-w-full`)

### flexbox

- `flex-row`, `flex-col`: flex direction
- `justify-{position}`: justify content (`justify-start`, `justify-center`, `justify-end`, `justify-between`)
- `items-{position}`: align items (`items-start`, `items-center`, `items-end`, `items-stretch`)
- `flex-wrap`, `flex-nowrap`: flex wrapping
- `flex-1`, `flex-auto`, `flex-initial`, `flex-none`: flex grow/shrink

### backgrounds

- `bg-{color}`: background color (`bg-blue-500`, `bg-gray-200`, `bg-white`)
- `bg-opacity-{value}`: background opacity (`bg-opacity-25`, `bg-opacity-50`, `bg-opacity-75`)
- `bg-{position}`: background position (`bg-center`, `bg-top`, `bg-bottom`)
- `bg-{size}`: background size (`bg-cover`, `bg-contain`)

### borders

- `border`: add border with default width
- `border-{width}`: border width (`border-0`, `border-2`, `border-4`)
- `border-{color}`: border color (`border-gray-300`, `border-blue-500`)
- `rounded-{size}`: border radius (`rounded`, `rounded-sm`, `rounded-md`, `rounded-lg`, `rounded-full`)
- `rounded-t-{size}`, `rounded-r-{size}`, etc.: border radius on specific sides

### effects

- `shadow-{size}`: box shadow (`shadow-sm`, `shadow`, `shadow-md`, `shadow-lg`, `shadow-xl`)
- `opacity-{value}`: opacity (`opacity-25`, `opacity-50`, `opacity-75`, `opacity-100`)

### example of combining classes

```html
<!-- A card component using Tailwind classes -->
<div class="bg-white rounded-lg shadow-md p-6 m-4 max-w-sm mx-auto">
  <h2 class="text-xl font-bold text-gray-800 mb-2">Card Title</h2>
  <p class="text-gray-600 mb-4">This is a simple card built with Tailwind CSS utility classes.</p>
  <button class="bg-blue-500 hover:bg-blue-600 text-white font-medium py-2 px-4 rounded">
    Click Me
  </button>
</div>
```
