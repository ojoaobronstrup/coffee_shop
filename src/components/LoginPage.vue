<template>
  <v-carousel show-arrows="false" hide-delimiter-background cycle height="100vh">
    <v-carousel-item v-for="slider in sliders"
      :src="slider"
      :key="slider"
      cover
    >
      <div>
        <h1>AOS QUE POUCO DORMEM E MUITO SONHAM</h1>
        <h3>Seu refúgio diário de sabor e aroma</h3>
        <v-btn>LOGIN</v-btn>
      </div>
    </v-carousel-item>
  </v-carousel>
</template>

<script>
export default {
  name: 'LoginPage',
  data() {
    return {
      sliders: ["/assets/coffee_1.jpeg", "/assets/coffee_2.jpeg", "/assets/coffee_3.jpeg"],
    }
  }
}
</script>
