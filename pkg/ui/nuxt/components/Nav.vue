<script>
export default {
  data() {
    let activeIndex = 'links'

    switch (this.$route.path) {
      case '/about':
        activeIndex = 'about'
        break
      default:
        activeIndex = 'links'
    }

    return {
      activeIndex,
    }
  },
  render(h) {
    return (
      <md-toolbar class="md-medium md-primary">
        <div class="md-toolbar-row">
          <div class="md-toolbar-section-start">
            <md-button class="md-icon-button">
              <md-icon>menu</md-icon>
            </md-button>
            <h3 class="md-title">Shortlink</h3>
          </div>

          <div class="md-toolbar-section-end">
            <div class="md-toolbar-row md-toolbar-offset">
              <md-tabs class="md-primary" md-alignment="fixed" md-active-tab={this.activeIndex}>
                <md-tab id="links" md-label="Links" href="/"></md-tab>
                <md-tab id="about" md-label="About" href="/about"></md-tab>
              </md-tabs>
            </div>

            <md-button class="md-icon-button">
              <md-icon>refresh</md-icon>
            </md-button>

            <md-button class="md-icon-button">
              <md-icon>more_vert</md-icon>
            </md-button>
          </div>
        </div>
      </md-toolbar>
    )
  },
}
</script>
