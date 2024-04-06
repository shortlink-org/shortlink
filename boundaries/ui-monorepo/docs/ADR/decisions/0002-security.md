# 2. Implementing Security Measures in UI

Date: 2023-05-16

## Status

Accepted

## Context

In order to enhance the security of our user interface and protect against potential threats, we have identified the need to implement additional security measures.

## Decision

### CAPTCHA Implementation

We have decided to use Cloudflare Turnstile, a smart CAPTCHA alternative. Turnstile can be embedded into any website without sending traffic through Cloudflare and works without showing visitors a CAPTCHA. This system is effective in verifying that user interactions are genuine and not automated bots, thus preventing spam and automated extraction of data from websites. More information about Turnstile can be found in the [documentation](https://developers.cloudflare.com/turnstile/).

![Turnstile Schema](https://developers.cloudflare.com/assets/turnstile-overview_hu857217e6cfe3055a024af7c1505ed0dc_210985_3757x2700_resize_q75_box_3-3bb896c3.png)

### UI Component

For the user interface, we will use the React Turnstile component. This component is designed to work with Cloudflare Turnstile and provides an easy-to-use interface for integrating CAPTCHA functionality into our application. More information about React Turnstile can be found in the [documentation](https://docs.page/marsidev/react-turnstile/).

## Consequences

+ The implementation of Cloudflare Turnstile will help to protect our application from spam and automated attacks.
+ The use of the React Turnstile component will allow us to easily integrate CAPTCHA functionality into our user interface.
+ Users will need to complete the CAPTCHA verification before they can interact with certain parts of our application. This may add an extra step in the user journey, but it is a necessary measure to ensure the security of our application.
