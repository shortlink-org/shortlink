<script>
  import { mapGetters, mapActions } from 'vuex'

  export default {
    data() {
      let { links } = this.$store.state;

      return {
        links,
      }
    },
    render(h) {
      return (
        <el-main>
          <h1>Links</h1>

          <el-table data={this.links} height="100%">
            <el-table-column fixed prop="Url" label="URL" width="240"></el-table-column>
            <el-table-column prop="Hash" label="Hash" width="140"></el-table-column>
            <el-table-column prop="Describe" label="Describe"></el-table-column>
            <el-table-column prop="CreatedAt" label="Create at" width="140" formatter={ this.formatterTime }></el-table-column>
            <el-table-column prop="UpdatedAt" label="Update at" width="140" formatter={ this.formatterTime }></el-table-column>
          </el-table>

          <md-speed-dial md-direction="top">
            <md-speed-dial-target>
              <md-icon>add</md-icon>
            </md-speed-dial-target>

            <md-speed-dial-content>
              <md-button class="md-icon-button">
                <md-icon>link</md-icon>
              </md-button>
            </md-speed-dial-content>
          </md-speed-dial>
        </el-main>
      )
    },
    head: {
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
