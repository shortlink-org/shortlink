<template>
  <div class="content">
    <Nav />
    <nuxt/>
  </div>
</template>

<script>
import Nav from '../components/Nav.vue'

export default {
  components: {
    Nav,
  }
}
</script>

<style>
.content {
  display: grid;
  grid-template-rows: auto 1fr;
}
</style>
