<template>
  <div id="app">
    <app-navbar />
    <div class="container">
      <router-view/>
    </div>
  </div>
</template>

<script>
import Navbar from './components/Navbar.vue'

export default {
  name: 'App',
  components: {
    appNavbar: Navbar
  }
}
</script>

<style>

</style>
