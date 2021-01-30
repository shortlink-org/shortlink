<script>
export default {
  data() {

    return {
      title: this.$route.name,
      drawer: null,
      items: [
        { title: 'Links', icon: 'mdi-account', path: "/" },
        { title: 'About', icon: 'mdi-account', path: "/about" },
      ],
    }
  },
  render(h) {
    return (
      <div>
        <v-app-bar
          id="core-toolbar"
          app
          dense
          fixed
          flat
          dense
          tile
          collapse-on-scroll
          color="grey lighten-3"
        >
          <v-app-bar-nav-icon
            color="grey"
            onClick={() => this.changeDrawer(this)}
          />

          <v-toolbar-title light>
            { this.title }
          </v-toolbar-title>

          <v-spacer />

          <v-toolbar-items>
            <v-flex
              align-center
              layout
              py-2
            >
              <v-text-field
                label="Search..."
                hide-details
                color="grey"
              />

              <v-icon color="tertiary">mdi-account</v-icon>
            </v-flex>
          </v-toolbar-items>
        </v-app-bar>

        <v-navigation-drawer
          v-model={this.drawer}
          absolute
          temporary
          light
        >
          <v-list-item>
            <v-list-item-avatar>
              <v-img src="https://randomuser.me/api/portraits/men/78.jpg"></v-img>
            </v-list-item-avatar>

            <v-list-item-content>
              <v-list-item-title>John Leider</v-list-item-title>
            </v-list-item-content>
          </v-list-item>

          <v-divider></v-divider>

          {
            this.items.map(item => (
              <nuxt-link to={item.path}>
                <v-list-item link key={ item.icon }>
                  <v-list-item-icon>
                    <v-icon>{ item.icon }</v-icon>
                  </v-list-item-icon>

                  <v-list-item-content>
                    <v-list-item-title>{ item.title }</v-list-item-title>
                  </v-list-item-content>
                </v-list-item>
              </nuxt-link>
            ))
          }
        </v-navigation-drawer>
      </div>
    )
  },
  watch: {
    '$route' (val) {
      this.title = val.name
    }
  },
  methods: {
    changeDrawer: async (state) => {
      state.$data.drawer = !state.$data.drawer
    }
  },
}
</script>
