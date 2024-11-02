<template>
  <div class="">
    <v-navigation-drawer
    app
    v-model="drawer"
    class="custom-drawer"
  >
    <v-list>
      <v-list-item-title class="logo-wrap">
        <a href="/home">
          <img src="../assets/Logo-dprd.png" alt="" class="sidebar-logo mb-3">
        </a>
      </v-list-item-title>

      <!-- Loop untuk item utama lainnya -->
      <v-list-item
        class="list"
        min-width="100"
        density="default"
        v-for="(links, i) in items"
        :key="i"
        :value="links"
        :to="links.route"
        :title="links.text"
        :prepend-icon="links.icon"
        active-class="active-item"
      ></v-list-item>

      <!-- Manajemen Pengguna with 2 Child items -->
       <div class="" v-if="auth == 'admin'">
      <v-list-group
      append-icon="false"
      prepend-icon="mdi-chevron-down"
      >
        <template v-slot:activator="{props: Props}">
          <v-list-item
          v-bind="Props"
          value="true"
          title="Manajemen Pengguna"
          >
        </v-list-item>
        </template>

        <!-- Child item: DPRD -->
        <v-list-item
        to="/admin/manajemen-pengguna/dprd"
        title="Anggota DPRD"
        prepend-icon="mdi-account"
        active-class="active-item"
        >

        </v-list-item>

        <!-- Child item: User -->
        <v-list-item
        to="/admin/manajemen-pengguna/user"
        title="Masyarakat"
        prepend-icon="mdi-account"
        active-class="active-item"
        >
        </v-list-item>
      </v-list-group>
    </div>
    </v-list>
  </v-navigation-drawer>

  <v-app-bar
    color=""
    prominent
    elevation="0"
  >
    <div class="d-lg-none">
      <v-app-bar-nav-icon @click="drawer = !drawer">

      </v-app-bar-nav-icon>
    </div>
    <v-spacer></v-spacer>
    <div class="mr-5">
      <v-icon class="" @click="handleLogout()">mdi-logout</v-icon>
    </div>
   <div class="profile-container mr-5"><a href="/profile">
    <img v-if="userPhoto == ''" src="../assets/profile.png"  alt="" class="profile-pictures" >
    <img v-else :src="`${getImageUrl(userPhoto)}`"  alt="" class="profile-pictures" >
   </a>
   </div>
  </v-app-bar>
  </div>
</template>

<script>
import { getImageUrl } from '@/config/foto';
import router from '@/router';
import { authStore } from '@/store/auth';
import { navitemstore } from '@/store/navitem';
import axios from 'axios';
import Swal from 'sweetalert2';

export default {
  setup() {

  },
  data: () => ({
    getImageUrl,
    drawer: true,
    items: [],
    auth:'',
    userPhoto: '',

  }),
  mounted(){
    this.handleManajemenUser();
    this.getUserPhoto();

  },
  methods:{
    getUserPhoto(){
      axios.get('/api/user')
      .then((res)=>{
        this.userPhoto = res.data.data.foto_profil
      })
    },
    handleManajemenUser(){
      var auth = navitemstore();
      this.auth = auth.stat;
      const authen = authStore();
      const item = navitemstore();
      this.items = item.navitemlist;
      authen.check(router, this.$swal);
      item.check();

    },
    async handleLogout() {
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
          await axios.post("/api/logout");
          localStorage.removeItem('Role');
          const auth = authStore();
          auth.check(router, this.$swal);
          const item = navitemstore();
          item.reset();

          // Tampilkan sukses
          await Swal.fire({
            icon: 'success',
            title: 'Berhasil Keluar',
            text: 'Anda telah berhasil logout',
            showConfirmButton: false,
            timer: 1500
          });

          // Redirect ke login
          this.$router.push("/home");
        }
      } catch (error) {
        console.error("Error during logout:", error);
        
        // Tampilkan error
        Swal.fire({
          icon: 'error',
          title: 'Oops...',
          text: 'Terjadi kesalahan saat logout',
          confirmButtonText: 'Tutup'
        });
      }
    },
  }
}
</script>

<style lang="scss">
.profile-container{
  width: 30px;
  height: 30px;
  border-radius: 50%;
  .profile-pictures{
    width: 100%;
    height: 100%;
    border-radius: 50%;
  }
}
</style>
