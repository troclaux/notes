
# angular

> front-end web application framework used to build SPAs

- typescript-based
- developed by google

## core concepts

- components
  - composed by
    - typescript class (logic)
    - html template (view)
    - css file (styling)
- modules: organize an app into cohesive blocks
- template: HTML with additional features, [similar to jsx/tsx in react.js](/react.md#jsx-javascript-xml)
- template variables: used to reference DOM elements, component instances or directive instances
  - e.g. `<input #myInput type="text">` => `<button (click)="logValue(myInput.value)">Log</button>`
    - `#myInput` creates a template variable referencing the `<input>` element
    - `myInput.value`: accesses the `value` property of the DOM input element
- data binding: connects component's logic to the UI
  - `{{ ... }}`: interpolation (one-way) that contains a expression (what is written inside the `{{ ... }}`)
  - `[property]`: property binding
  - `(click)`: event binding
  - `[(ngModel)]`: two-way binging, which means that if the user types in an input field, the component will update
    - `<input [(ngModel)]="name">`
    - `<p>Hello, {{ name }}!</p>`

## CLI

- install angular: `npm install -g @angular/cli`
- create new angular project: `ng new my-app`
- start angular development server: `ng serve`
- generate production build: `ng build` or `ng build --configuration=production`
  - creates static files in `dist/` folder (usually `dist/app-name`)
    - which must be served by a web server (e.g. nginx/apache) in production, often inside a container

- apply lint: `ng lint`
- apply tests: `ng test`
- update dependencies: `ng update`
- generate/modify files: `ng generate`

## template

## event

- `(click)`: on mouse click
  - e.g. `<button (click)="handleClick()">Click me</button>`
- `(input)`: when user types in input
- `(change)`: when input value changes and loses focus
- `(submit)`: on form submit (not often used directly)
- `(ngSubmit)`: Angular-specific form submit (use with ngForm)
- `(keydown)`: on key press
- `(mouseover)`: on mouse hover

example:

```html
<input (input)="onType($event)">
<button (click)="sayHello()">Say Hello</button>
```

in the component

```typescript
export class AppComponent {
  sayHello() {
    console.log('Hello!');
  }

  onType(event: Event) {
    const input = event.target as HTMLInputElement;
    console.log(input.value);
  }
}
```

## expressions

- angular expressions don't support conditionals, loops and exceptions

- `currency`
- `date`
- `filter`
- `json`
- `limitTo`
- `lowercase`, `uppercase`
- `number`
- `orderBy`

## decorator

- angular decorator != design pattern decorator
- angular decorator: special function that adds metadata to a class, method, property or parameter
  - similar to a label or tag

types of decorators:

| decorator        | used for                     |
|------------------|------------------------------|
| `@Component`     | declares a component          |
| `@NgModule`      | declares a module             |
| `@Injectable`    | declares a service that can be injected |
| `@Directive`     | declares a directive          |
| `@Pipe`          | declares a pipe               |

### component

the code below renders `<h1>hello, angular!</h1>` wherever you use `<app-hello>` in html

```ts
import { Component } from '@angular/core';

@Component({
  selector: 'app-hello',
  template: `<h1>Hello, {{ name }}!</h1>`,
  styles: [`h1 { color: blue; }`]
})
export class HelloComponent {
  name = 'Angular';
}
```

### module

> groups related parts of an angular app, like components, services and other modules

- each angular app must have at least one module, the root module (`AppModule`), defined in `app.module.ts`
- a module is a class with `@NgModule` decorator

```ts
@NgModule({
  declarations: [AppComponent, HomeComponent], // components, directives, pipes
  imports: [BrowserModule, FormsModule],       // other modules you want to use
  providers: [],                               // services
  bootstrap: [AppComponent]                    // root component to start the app
})
export class AppModule { }
```

### injectable

> marks a class as available to be provided and injected as a dependency

- used for services that can be injected into components, directives, pipes, or other services
- enables dependency injection system in Angular

```ts
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root' // makes the service available throughout the app
})
export class DataService {
  getData() {
    return ['item1', 'item2', 'item3'];
  }
}
```

### directive

> extends HTML with custom behavior

- three types: components, structural directives, and attribute directives
- components are directives with templates
- structural directives change DOM layout (e.g., `ngIf`, `ngFor`)
- attribute directives change appearance or behavior of elements

```ts
import { Directive, ElementRef, Input, OnInit } from '@angular/core';

@Directive({
  selector: '[appHighlight]' // used as an attribute on elements
})
export class HighlightDirective implements OnInit {
  @Input() appHighlight: string = 'yellow';

  constructor(private el: ElementRef) {}

  ngOnInit() {
    this.el.nativeElement.style.backgroundColor = this.appHighlight;
  }
}
```

### pipe

> transforms data for display in templates

- used in templates to transform strings, currency amounts, dates, and other data
- can be chained and accept parameters

```ts
import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'reverse'
})
export class ReversePipe implements PipeTransform {
  transform(value: string): string {
    return value.split('').reverse().join('');
  }
}
```

```html
<p>{{ 'hello' | reverse }}</p> <!-- Output: 'olleh' -->
```

## controllers
## services
## angular api
## routing
