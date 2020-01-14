<script>
  export default {
    data() {
      return {}
    },
    render(h) {
      return (
        <div>
          <h1>Navigation</h1>
          <nuxt-link to="/">Links</nuxt-link>
          <nuxt-link to="/about">About</nuxt-link>
        </div>
      )
    },
    head: {
      title: 'About of project'
    }
  }
</script>
