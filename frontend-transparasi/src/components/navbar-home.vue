<template>
  <v-app-bar
    app
    :color="navbarColor"
    :elevation="navbarElevation"
    class="px-10"
    height="80"
  >
    <v-img
    v-if="!isScrolled"
      src="../assets/logo-dprd.png"
      max-width="150"
      contain
    ></v-img>
    <v-img
    v-else-if="isScrolled"
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
        to="/home"
        :class="{ 'text-white': !isScrolled, 'text-black': isScrolled }"
      >
        Beranda
      </v-btn>
      <v-btn
        text
        style="text-transform: none;"
        class=""
        to="/program"
        :class="{ 'text-white': !isScrolled, 'text-black': isScrolled }"
      >
        Program
      </v-btn>
      <v-btn
      v-if="role == 'admin'"
        text
        style="text-transform: none;"
        class=""
        href="/admin/dashboard"
        :class="{ 'text-white': !isScrolled, 'text-black': isScrolled }"
      >
        Dashboard
      </v-btn>
      <v-btn
      v-else-if="role == 'dprd'"
        text
        style="text-transform: none;"
        class=""
        href="/dprd/dashboard"
        :class="{ 'text-white': !isScrolled, 'text-black': isScrolled }"
      >
        Dashboard
      </v-btn>
    </div>
  <div class="">
    <div class="d-flex" style="">
    <div class="" v-if="login ===false">
      <v-btn
        variant="outlined"
        style="text-transform: none;"
        :style="{'border-color': isScrolled ? '#BF3232' : 'white'}"
        class="ml-6 mr-2"
        href="/register"
        :class="{ 'text-white': !isScrolled, 'text-black': isScrolled }"
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
      <div class="d-flex" v-else>
        <v-btn
          text
          :class="{ 'text-white': !isScrolled, 'text-black': isScrolled }"
          style="text-transform: none;"
          class=""
          @click="logout"
        >
        <v-icon>mdi-logout</v-icon>
        </v-btn>
        <div class="profile-container " style="margin-top:1.8px;"><a href="/profile">
          <img v-if="userPhoto == ''" src="../assets/profile.png"  alt="" class="profile-pictures" >
          <img v-else :src="`${getImageUrl(userPhoto)}`"  alt="" class="profile-pictures" >
        </a>
      </div>
      </div>
    </div>
  </div>
  </v-app-bar>
</template>

<script>
import Swal from 'sweetalert2';
import axios from 'axios';
import { getImageUrl } from '@/config/foto';

export default {
  name: 'NavbarHome',
  data() {
    return {
      role: localStorage.getItem('Role'),
      getImageUrl,
      isScrolled: false,
      login: false,
      userPhoto: ''
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
    this.getUserPhoto()
  },

  beforeDestroy() {
    window.removeEventListener('scroll', this.handleScroll)
  },

  methods: {
    getUserPhoto() {
      axios.get('/api/user').then((res)=>{
        this.userPhoto = res.data.data.foto_profil
      })
    },

    handleScroll() {
      this.isScrolled = window.scrollY > 50
    },

    async logout() {
      try {
        // Konfirmasi logout
        const result = await Swal.fire({
          title: 'Apakah anda yakin?',
          text: "Anda akan keluar dari aplikasi",
          icon: 'warning',
          showCancelButton: true,
          confirmButtonColor: '#3085d6',
          cancelButtonColor: '#d33',
          confirmButtonText: 'Ya, Keluar!',
          cancelButtonText: 'Batal'
        });

        // Jika user mengkonfirmasi
        if (result.isConfirmed) {
          // Tampilkan loading
          Swal.fire({
            title: 'Sedang memproses...',
            allowOutsideClick: false,
            didOpen: () => {
              Swal.showLoading();
            }
          });

          // Proses logout
          await axios.post('/api/logout');
          localStorage.removeItem('Role');
          
          // Tampilkan sukses
          await Swal.fire({
            icon: 'success',
            title: 'Berhasil Keluar',
            text: 'Anda telah berhasil logout',
            showConfirmButton: false,
            timer: 1500
          });

          // Refresh halaman
          this.$router.push('/home');
          location.reload();
        }
      } catch (error) {
        console.error('Error during logout:', error);
        
        // Tampilkan error
        Swal.fire({
          icon: 'error',
          title: 'Oops...',
          text: 'Terjadi kesalahan saat logout',
          confirmButtonText: 'Tutup'
        });
      }
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
