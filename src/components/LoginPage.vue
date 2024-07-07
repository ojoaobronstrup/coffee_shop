<template>
  <el-carousel height="100vh">
    <el-carousel-item v-for="srcImage in srcImages" :key="srcImage">
      <img class="carousel-item" :src="srcImage" alt="Wallpaper of people drinking coffee">
    </el-carousel-item>
    <el-container class="text-overlay">
      <el-text class="text-main">AOS QUE POUCO DORMEM E MUITO SONHAM</el-text>
      <el-text class="text-main">Seu refúgio diário de sabor e aroma</el-text>
    </el-container>
  </el-carousel>
</template>

<script>
  export default {
    name: 'LoginPage',
    data() {
      return {
        srcImages: ["/assets/coffee_1.png", "/assets/coffee_2.png", "/assets/coffee_3.png"],
      }
    }
  }
</script>

<style>
  .carousel-item {
    width: 100vw;
    height: 100vh;
  }

  .text-overlay {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 2;
    pointer-events: none;
  }

  .text-main {
    color: white;
  }
</style>

