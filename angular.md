
# angular

> front-end web application framework used to build SPAs

- typescript-based
- developed by Google

## core concepts

- components
  - composed by
    - typescript class (logic)
    - html template (view)
    - css file (styling)
- modules: organize an app into cohesive blocks
- templates: HTML with additional features such as `*ngIf`, `*ngFor`, `{{ expression }}`
  - e.g. `<h1 *ngIf="isLoggedIn">Welcome back!</h1>`
- data binding: connects component's logic to the UI
  - `{{ }}`: interpolation (one-way)
  - `[(ngModel)]`: two-way binding
  - `[property]`: property binding
  - `(event)`: event binding

## CLI

- install angular: `npm install -g @angular/cli`
- create new angular project: `ng new my-app`
- start angular development server: `ng serve`
- generate production build: `ng build` or `ng build --configuration=production`
  - creates static files in `dist/` folder (usually `dist/app-name`)
    - which must be served by a web server (e.g. nginx/apache) in production, often inside a container
