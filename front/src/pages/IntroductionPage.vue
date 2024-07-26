<template>
  <el-carousel arrow="never" class="home-carousel-container" height="100vh">
    <el-carousel-item v-for="srcImage in srcImages" :key="srcImage">
      <img class="home-carousel-item" :src="srcImage" alt="Wallpaper of people drinking coffee">
    </el-carousel-item>
    <el-container class="home-content">
      <el-text class="home-content-title">SEU REFÚGIO DIÁRIO DE SABOR E AROMA</el-text>
      <EnterButton class="home-content-login" @click="goToLoginPage()">Entrar</EnterButton>
    </el-container>
  </el-carousel>
</template>

<script>
import EnterButton from '@/components/EnterButton.vue';

  export default {
    name: 'IntroductionPage',

    components: {
      EnterButton
    },

    data() {
      return {
        srcImages: ["/assets/coffee_1.png", "/assets/coffee_2.png", "/assets/coffee_3.png"],
      }
    },

    methods: {
      goToLoginPage() {
        this.$router.push('/login')
      }
    }
  }

</script>

<style>
@import url(./../styles/colors.css);

  .home-carousel-container {
    text-align: center;
  }

  .home-carousel-item {
    width: 100vw;
    height: 100vh;
  }

  .home-content {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    z-index: 2;
    flex-wrap: wrap;    
    gap: 10px;
  }

  .home-content-title {
    color: var(--gray-40);
    font-family: var(--font-lato);
    font-weight: bold;
    font-size: 28px;
    line-height: 32px;
    max-width: 300px;
  }

  .home-content-login {
    background-color: var(--secondary-brown);
    color: var(--gray-20);
    border: none;
    border-radius: 34px;
    width: 10em;
    height: 2.5em;
    font-size: 20px;
    line-height: 24px;
    font-weight: bold;
    opacity: 0.9;
  }

</style>

