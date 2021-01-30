<script>
  export default {
    render(h) {
      return (
        <div>
          <h1>About</h1>

          <p>
            GitHub URL: <a href="https://github.com/batazor/shortlink">https://github.com/batazor/shortlink</a>
          </p>
        </div>
      )
    },
    head: {
      htmlAttrs: {
        lang: 'en'
      },
      title: 'About of project'
    }
  }
</script>
