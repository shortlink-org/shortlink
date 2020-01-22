<template>
  <v-app>
    <Nav />

    <!-- Sizes your content based upon application components -->
    <v-content>
      <!-- Provides the application the proper gutter -->
      <v-container fluid>
        <nuxt/>
      </v-container>
    </v-content>
  </v-app>
</template>

<script>
import Nav from '../components/Nav.vue'

export default {
  components: {
    Nav,
  }
}
</script>

<style lang="scss">
@import '@/styles/index.scss';

/* Remove in 1.2 */
.v-datatable thead th.column.sortable i {
  vertical-align: unset;
}
</style>
