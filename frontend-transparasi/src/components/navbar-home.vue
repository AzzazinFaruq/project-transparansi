<template>
  <v-app-bar
    app
    :color="navbarColor"
    :elevation="navbarElevation"
    class="px-10"
    style="position: absolute; top: 0;"
    height="80"
  >
    <v-img
      class="ml-5"
      src="../assets/dprd-hitam.png"
      max-width="150"
      contain
    ></v-img>

    <v-spacer></v-spacer>

    <div class="mr-10" >
      <v-btn
        text
        class=""
        style="color: black !important; text-transform: none;"
      >
        Beranda
      </v-btn>
      <v-btn
        text
        class=""
        style="color: black !important; text-transform: none; "
      >
        Program
      </v-btn>
    </div>
    <div class="">
        <v-btn
        variant="outlined"
        style="border-color: #BF3232; text-transform: none;"
        >
        Daftar
        </v-btn>
        <v-btn
        class="ml-1 mr-5"
        variant="tonal"
        style="text-transform: none; background-color: #BF3232 !important;color: white;"
        >
        Masuk
        </v-btn>
      </div>
  </v-app-bar>
</template>

<script>
export default {
  name: 'NavbarHome',
  data() {
    return {
      isScrolled: false,
      menuItems: [
        { title: 'Beranda' },
        { title: 'Tentang' },
        { title: 'Layanan' },
        { title: 'Kontak' },
      ]
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
  },
  beforeDestroy() {
    window.removeEventListener('scroll', this.handleScroll)
  },
  methods: {
    handleScroll() {
      this.isScrolled = window.scrollY > 50
    }
  }
}
</script>

<style scoped>
.v-app-bar {
  position: fixed !important;
  top: 0 !important;
  left: 0 !important;
  right: 0 !important;
  transition: all 0.3s ease-in-out;
}

/* Override default transition */
.v-app-bar.v-scroll-threshold {
  transition: none;
}

/* Hapus properti app dari v-app-bar */
:deep(.v-toolbar__content) {
  padding-top: 0 !important;
}
</style>
