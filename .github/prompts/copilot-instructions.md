---
agent: agent
---

# How to Go - Coding Standards and Guidelines

This project is a web application built with **Go (Golang)** using **Gin** for routing, **Templ** for HTML generation, and **Bulma** for CSS styling.

## Tech Stack

- **Language:** Go 1.22+
- **Router/Framework:** Gin (github.com/gin-gonic/gin)
- **Templating:** Templ (github.com/a-h/templ)
- **CSS Framework:** Bulma (via CDN or local file)
- **Styling:** Bulma classes first, Custom CSS (`static/css/style.css`) for overrides.

## Coding Standards

### 1. Go & Gin Guidelines

- **Project Layout:** Follow standard Go project layout ( `tools/`, `templates/`, `static/`).
- **Handlers:** - Handlers should exist in `handlers`.
  - Do not use `c.HTML` or `html/template`.
  - Always render responses using Templ components.
  - Create a helper function or middleware to bridge Gin and Templ (e.g., `render(c *gin.Context, component templ.Component)`).
- **Error Handling:** Use idiomatic `if err != nil` checks. Return JSON errors for API routes, but render Error UI components for page routes.
 - **Standard Template Layout**: For example pages, use .github/example-page.templ as a reference for content structure and layout.

### 2. Templ Guidelines

- **File Structure:** Place all `.templ` files in the `templates/` directory.
- **Syntax:** Use typed arguments for components. Avoid passing `interface{}` unless absolutely necessary.
- **Composition:** Break UI into small, reusable components (e.g., `templates/components/navbar.templ`, `templates/layouts/base.templ`).
- **HTMX:** If interactivity is needed, prefer HTMX attributes over raw JavaScript.

### 3. CSS & Bulma Guidelines

- **Prioritize Bulma:** Use Bulma utility classes (e.g., `is-primary`, `columns`, `section`, `container`) before writing custom CSS.
- **Responsive Design:** Ensure layouts use Bulma's responsive modifiers (e.g., `is-mobile`, `is-desktop`).
- **Custom CSS:** If a specific style is not available in Bulma, add a class in `style.css` and reference it. Avoid inline styles in Templ files.

## Preferred Patterns

### Render Helper

When writing handlers, assume a helper function exists to render Templ components into the Gin context:

```go
// internal/render/render.go
func Render(c *gin.Context, status int, component templ.Component) {
	c.Status(status)
	component.Render(c.Request.Context(), c.Writer)
}

Handler Pattern
Go

func (h *Handler) ShowHome(c *gin.Context) {
    data := getData()
    component := home.Index(data)
    render.Render(c, http.StatusOK, component)
}

Templ Layout Pattern

Always use a Base layout for pages:
Go

// views/layouts/base.templ
package layouts

templ Base(title string) {
    <!DOCTYPE html>
    <html>
        <head>
            <title>{ title }</title>
            <link rel="stylesheet" href="[https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css](https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css)">
            <link rel="stylesheet" href="/static/css/style.css">
        </head>
        <body>
            { children... }
        </body>
    </html>
}

---

// Page example using Base layout

## Follow the style of all pages in views/pages/*.templ

Each page should use the Base layout and contain a consistent structure:

- @layout layouts.Base("Page Title")
- @layout.Header()
- Main content area with Bulma classes (e.g., `section`, `container`)
- 2 Code block examples demonstrating the concepts uniquely to that page. Ensure
the most important information is included in as little lines of code as possible to demonstrate the concept.
- Bulma tables detaililng declarations, types, patterns, concepts, internals, etc. 
- Key points summarized in bullet lists.
- @layout.Footer()

---

## What I Like

    Keep a consistent layout, look and feel across all pages using Bulma and custom CSS.

    This project is inspired by www.gobyexample.com, so keep that style in mind when creating components and pages.

    You can use modern HTML5 and Bulma features to enhance the UI/UX.

    Always consider SEO optimization or implications when creating components. Ensure proper use of headings, meta tags, and semantic HTML. Google search ranking is important for this project. Ensure that the generated HTML is SEO-friendly and Google will index it well.

    Pages about Go standard library packages should always include a full table containing every function, type, and constant in that package, similar to how the documentation lists them, but with better explanations and examples.

    Keep logic simple and clean in both Go and Templ files. Focus on readability and maintainability.

    You may use other Go packages if needed, but avoid overcomplicating the project with unnecessary dependencies.

    You may use HTMX or similar libraries for interactivity, but avoid heavy JavaScript frameworks.

    Use font-awesome icons via CDN or utilize /icons directory SVGs for icons.


What to Avoid

    Do not suggest html/template standard library solutions.

    Do not suggest React, Vue, or other JS frameworks.

    Do not make any github changes like reverting commits or modifying workflows without explicit permission.

    Do not mix logic inside Templ files; keep data processing in the Go handler.
```
