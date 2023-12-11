# 4. Accessibility

Date: 2023-11-08

## Status

Accepted

## Context

We need to make our UI accessible for people with disabilities.

## Decision

We will use:

### [CSS prefers-reduced-transparency](https://developer.chrome.com/blog/css-prefers-reduced-transparency/) to reduce transparency effects for people with visual impairments.

Example:

```css
.bg-white {
  @media (prefers-reduced-transparency: reduce) {
    --tw-bg-opacity: 0;
  }
}
```

## Consequences

- We will be able to make our UI accessible for people with disabilities.

### References

- [What is accessibility?](https://developer.mozilla.org/en-US/docs/Learn/Accessibility/What_is_accessibility)
