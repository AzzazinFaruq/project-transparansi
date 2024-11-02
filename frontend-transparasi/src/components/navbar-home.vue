<template>
  <div>
    <!-- Desktop Navbar -->
    <v-app-bar
      app
      :color="navbarColor"
      :elevation="navbarElevation"
      class="px-sm-10 px-2"
      height="80"
    >
      <!-- Logo -->
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

      <!-- Desktop Menu -->
      <div class="d-none d-sm-flex">
        <v-btn
          text
          style="text-transform: none;"
          to="/home"
          :class="{ 'text-white': !isScrolled, 'text-black': isScrolled }"
        >
          Beranda
        </v-btn>
        <v-btn
          text
          style="text-transform: none;"
          to="/program"
          :class="{ 'text-white': !isScrolled, 'text-black': isScrolled }"
        >
          Program
        </v-btn>
        <v-btn
          v-if="role == 'admin'"
          text
          style="text-transform: none;"
          href="/admin/dashboard"
          :class="{ 'text-white': !isScrolled, 'text-black': isScrolled }"
        >
          Dashboard
        </v-btn>
        <v-btn
          v-else-if="role == 'dprd'"
          text
          style="text-transform: none;"
          href="/dprd/dashboard"
          :class="{ 'text-white': !isScrolled, 'text-black': isScrolled }"
        >
          Dashboard
        </v-btn>
      </div>

      <!-- Desktop Auth -->
      <div class="d-none d-sm-flex">
        <div class="d-flex" v-if="login === false">
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
            class="text-white"
          >
            Masuk
          </v-btn>
        </div>
        <div class="d-flex" v-else>
          <v-btn
            text
            :class="{ 'text-white': !isScrolled, 'text-black': isScrolled }"
            style="text-transform: none;"
            @click="logout"
          >
            <v-icon>mdi-logout</v-icon>
          </v-btn>
          <div class="profile-container" style="margin-top:1.8px;">
            <a href="/profile">
              <img v-if="userPhoto == ''" src="../assets/profile.png" alt="" class="profile-pictures">
              <img v-else :src="`${getImageUrl(userPhoto)}`" alt="" class="profile-pictures">
            </a>
          </div>
        </div>
      </div>

      <!-- Menu Icon for Mobile -->
      <v-app-bar-nav-icon
        @click="drawer = !drawer"
        class="d-sm-none"
        :class="{ 'text-white': !isScrolled, 'text-black': isScrolled }"
      ></v-app-bar-nav-icon>
    </v-app-bar>

    <!-- Mobile Navigation Drawer -->
    <v-navigation-drawer
      v-model="drawer"
      location="right"
      temporary
      width="280"
    >
      <v-list>
        <v-list-item class="px-4 py-4 mb-2">
          <div v-if="login">
            <div class="d-flex align-center mb-2">
              <div class="profile-container mr-4">
                <img v-if="userPhoto == ''" src="../assets/profile.png" alt="" class="profile-pictures">
                <img v-else :src="`${getImageUrl(userPhoto)}`" alt="" class="profile-pictures">
              </div>
              <a href="/profile" class="text-decoration-none">Profil Saya</a>
            </div>
          </div>
          <div v-else>
            <v-btn
              block
              color="#BF3232"
              style="text-transform: none;"
              class="text-white mb-2"
              href="/login"
            >
              Masuk
            </v-btn>
            <v-btn
              block
              variant="outlined"
              style="text-transform: none;"
              color="#BF3232"
              href="/register"
            >
              Daftar
            </v-btn>
          </div>
        </v-list-item>

        <v-divider></v-divider>

        <v-list-item
          to="/home"
          class="mb-2"
        >
          <template v-slot:prepend>
            <v-icon>mdi-home</v-icon>
          </template>
          <v-list-item-title>Beranda</v-list-item-title>
        </v-list-item>

        <v-list-item
          to="/program"
          class="mb-2"
        >
          <template v-slot:prepend>
            <v-icon>mdi-view-dashboard</v-icon>
          </template>
          <v-list-item-title>Program</v-list-item-title>
        </v-list-item>

        <v-list-item
          v-if="role == 'admin'"
          href="/admin/dashboard"
          class="mb-2"
        >
          <template v-slot:prepend>
            <v-icon>mdi-view-dashboard</v-icon>
          </template>
          <v-list-item-title>Dashboard</v-list-item-title>
        </v-list-item>

        <v-list-item
          v-else-if="role == 'dprd'"
          href="/dprd/dashboard"
          class="mb-2"
        >
          <template v-slot:prepend>
            <v-icon>mdi-view-dashboard</v-icon>
          </template>
          <v-list-item-title>Dashboard</v-list-item-title>
        </v-list-item>

        <v-list-item
          v-if="login"
          @click="logout"
          class="mb-2"
        >
          <template v-slot:prepend>
            <v-icon>mdi-logout</v-icon>
          </template>
          <v-list-item-title>Keluar</v-list-item-title>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
  </div>
</template>

<style scoped>
.v-app-bar {
  transition: all 0.3s ease-in-out;
}

.v-app-bar.v-scroll-threshold {
  transition: none;
}

.profile-pictures {
  width: 35px;
  height: 35px;
  border-radius: 50%;
  object-fit: cover;
}

/* Mobile styles */
@media (max-width: 960px) {
  .v-navigation-drawer {
    margin-top: 80px !important;
  }

  .profile-pictures {
    width: 40px;
    height: 40px;
  }

  .v-list-item {
    min-height: 44px;
  }

  /* Tambahan untuk memastikan drawer muncul di atas konten */
  .v-navigation-drawer {
    z-index: 1000;
  }
}

/* Animasi untuk drawer */
.v-navigation-drawer {
  transition: transform 0.3s ease-in-out;
}

.v-navigation-drawer--right {
  transform: translateX(100%);
}

.v-navigation-drawer--right.v-navigation-drawer--active {
  transform: translateX(0);
}
</style>

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
      userPhoto: '',
      drawer: false
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
