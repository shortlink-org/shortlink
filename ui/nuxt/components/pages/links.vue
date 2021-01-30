<script>
  import { mapGetters, mapActions } from 'vuex'

  export default {
    data() {
      let { links } = this.$store.state;

      return {
        links: links || [],

        // table setting
        headers: [
          { text: 'URL', value: 'Url' },
          { text: 'Hash', value: 'Hash' },
          { text: 'Describe', value: 'Describe' },
          { text: 'Created At', value: 'CreatedAt' },
          { text: 'Updated At', value: 'UpdatedAt' },
        ],
      }
    },
    render(h) {
      return (
        <div>
          <h1>Links</h1>

          <v-data-table
            headers={this.headers}
            items={this.links}
            items-per-page={5}
            class="elevation-1"
            fixed-header
          />
        </div>
      )
    },
    head: {
      htmlAttrs: {
        lang: 'en'
      },
      title: 'Links'
    },
    async fetch ({ store, params }) {
      await store.dispatch('GET_LINKS');
    },
    methods: {
      formatterTime(row, prop, value) {
        return this.$dateFns.format(new Date(value))
      }
    },
  }
</script>
