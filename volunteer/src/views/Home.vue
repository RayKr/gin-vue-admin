<template>
  <Header/>
  <!-- <div class="home">
    
    <img alt="Vue logo" src="../assets/logo.png">
    <HelloWorld msg="Welcome to Your Vue.js App"/>
  </div> -->
</template>

<script>
// @ is an alias to /src
import Header from "@/components/Header";

export default {
  name: 'Home',
  components: {
    Header
  }
}
</script>

<style scoped>
.home {
  width: 100%;
}
</style>