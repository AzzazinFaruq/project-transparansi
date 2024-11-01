<template>
  <v-app-bar
    app
    :color="navbarColor"
    :elevation="navbarElevation"
    class="px-10"
    height="80"
  >
    <v-img
      src="../assets/logo-dprd-hitam.png"
      max-width="150"
      contain
    ></v-img>

    <v-spacer></v-spacer>

    <div class="d-flex">
      <v-btn
        text
        style="text-transform: none;"
        class=""
        href="/home"
        :class="{ 'text-black': !isScrolled, 'text-black': isScrolled }"
      >
        Beranda
      </v-btn>
      <v-btn
        text
        style="text-transform: none;"
        class=""
        href="/program"
        :class="{ 'text-black': !isScrolled, 'text-black': isScrolled }"
      >
        Program
      </v-btn>
      <div class="" v-if="login ===false">
        <v-btn
          variant="outlined"
        style="text-transform: none;"
        :style="{'border-color': isScrolled ? '#BF3232' : '#BF3232'}"
        class="ml-6 mr-2"
        href="/register"
        :class="{ 'text-black': !isScrolled, 'text-black': isScrolled }"
      >
        Daftar
      </v-btn>
      <v-btn
        variant="tonal"
        style="background-color: #BF3232; text-transform: none;"
        href="/login"
          class=" text-white "
          >
          Masuk
        </v-btn>
      </div>
      <div class="" v-if="login ===true">
        <v-btn
          variant="text"
          style="text-transform: none;"
          class="ml-6 mr-2"
          @click="logout"
        >
          <v-icon>mdi-logout</v-icon>
        </v-btn>
      </div>
    </div>
  </v-app-bar>
</template>

<script>
import axios from 'axios';
export default {
  name: 'NavbarHome',
  data() {
    return {
      isScrolled: false,
      login: false
    }
  },
  computed: {
    navbarColor() {
      return this.isScrolled ? 'white' : 'transparent'
    },
    navbarElevation() {
      return this.isScrolled ? 4 : 0
    }
  },
  mounted() {
    window.addEventListener('scroll', this.handleScroll)
    this.login = localStorage.getItem('Role') ? true : false
  },
  beforeDestroy() {
    window.removeEventListener('scroll', this.handleScroll)
  },
  methods: {
    handleScroll() {
      this.isScrolled = window.scrollY > 50
      },
    logout() {
      axios.post('/api/logout').then((res)=>{
        localStorage.removeItem('Role')
        this.$router.push('/home')
      })
    }
  }
}
</script>

<style scoped>
.v-app-bar {
  transition: all 0.3s ease-in-out;
}

/* Override default transition */
.v-app-bar.v-scroll-threshold {
  transition: none;
}
</style>
