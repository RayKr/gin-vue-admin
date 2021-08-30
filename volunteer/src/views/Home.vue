<template>
  <!-- 顶部菜单 -->
  <Header/>
<!--  轮播图 -->
  <Swiper v-if="banners.length > 0">
    <Slide v-for="(item,index) in banners" :key="index">
      <img :src="item.src" alt="" class="swiper-img"/>
    </Slide>
  </Swiper>
</template>

<script>
// @ is an alias to /src
import Header from "@/components/Header";
import { Swiper, Slide } from 'vue-swiper-component';

export default {
  name: 'Home',
  components: {
    Header,
    Swiper,
    Slide
  },
  data() {
    return {
      banners: [{
        src: '/images/demo/banner1.jpeg'
      }, {
        src: '/images/demo/banner2.jpeg'
      }, {
        src: '/images/demo/banner3.jpeg'
      }, {
        src: '/images/demo/banner4.jpeg'
      }]
    }
  },
}
</script>

<style scoped>
.home {
  width: 100%;
}
.swiper-img {
  height: 400px;
}
</style>