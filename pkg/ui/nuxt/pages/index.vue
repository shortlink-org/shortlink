<script>
  export default {
    data() {
      const item = {
        url: 'http://example.com',
        hash: 'sh35f5',
        describe: 'No. 189, Grove St, Los Angeles',
        create_at: '30/12/2020',
        update_at: '30/12/2020',
      }

      return {
        tableData: Array(20).fill(item),
      }
    },
    render(h) {
      return (
        <el-main>
          <h1>Links</h1>

          <el-table data={this.tableData} height="100%">
            <el-table-column fixed prop="url" label="URL" width="240"></el-table-column>
            <el-table-column prop="hash" label="Hash" width="140"></el-table-column>
            <el-table-column prop="describe" label="Describe"></el-table-column>
            <el-table-column prop="create_at" label="Create at" width="140"></el-table-column>
            <el-table-column prop="update_at" label="Update at" width="140"></el-table-column>
          </el-table>
        </el-main>
      )
    },
    head: {
      title: 'Links'
    }
  }
</script>
