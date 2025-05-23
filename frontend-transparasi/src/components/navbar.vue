<template>
  <div class="">
    <v-navigation-drawer
    app
    v-model="drawer"
    class="custom-drawer"
  >
    <v-list >
      <v-list-item-title class="logo-wrap">
        <a href="/home">
          <img src="./logo-dprd.png" alt="" class="sidebar-logo mb-3">
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

      <div class="" v-if="auth == 'admin'">
          <!-- Manajemen Anggaran -->
          <div>
            <v-list-item
              @click="setActiveGroup('anggaran')"
              title="Manajemen Anggaran"
              class="group-header"
            >
              <template v-slot:prepend>
                <v-icon
                  :class="{ 'rotate-icon': activeGroup === 'anggaran' }"
                >
                  mdi-chevron-right
                </v-icon>
              </template>
            </v-list-item>

            <!-- Konten group dengan transition -->
            <transition name="expand">
              <div
                v-show="activeGroup === 'anggaran'"
                class="group-content"
              >
                <v-list-item
                  to="/admin/kategori-penggunaan"
                  title="Kategori Penggunaan"
                  prepend-icon="mdi-currency-usd"
                  active-class="active-item"
                  class="ml-4"
                >
                </v-list-item>

                <v-list-item
                  to="/admin/jenis-anggaran"
                  title="Jenis Anggaran"
                  prepend-icon="mdi-currency-usd"
                  active-class="active-item"
                  class="ml-4"
                >
                </v-list-item>
              </div>
            </transition>
          </div>

          <!-- Manajemen Pengguna -->
          <div>
            <v-list-item
              @click="setActiveGroup('pengguna')"
              title="Manajemen Pengguna"
              class="group-header"
            >
              <template v-slot:prepend>
                <v-icon
                  :class="{ 'rotate-icon': activeGroup === 'pengguna' }"
                >
                  mdi-chevron-right
                </v-icon>
              </template>
            </v-list-item>

            <!-- Konten group dengan transition -->
            <transition name="expand">
              <div
                v-show="activeGroup === 'pengguna'"
                class="group-content"
              >
                <v-list-item
                  to="/admin/manajemen-pengguna/dprd"
                  title="Anggota DPRD"
                  prepend-icon="mdi-account"
                  active-class="active-item"
                  class="ml-4"
                >
                </v-list-item>

                <v-list-item
                  to="/admin/manajemen-pengguna/user"
                  title="Masyarakat"
                  prepend-icon="mdi-account"
                  active-class="active-item"
                  class="ml-4"
                >
                </v-list-item>
              </div>
            </transition>
          </div>
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
    <div class="mr-6">
      <v-icon class="" @click="handleLogout()">mdi-logout</v-icon>
    </div>
   <div class="profile-container mr-8"><a href="/profile">
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
    activeGroup: null
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
    setActiveGroup(groupName) {
      if (this.activeGroup === groupName) {
        this.activeGroup = null;
      } else {
        this.activeGroup = groupName;
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

// Tambahkan style untuk animasi expand/collapse
.expand-enter-active,
.expand-leave-active {
  transition: all 0.3s ease;
  max-height: 300px; // Sesuaikan dengan kebutuhan
  opacity: 1;
}

.expand-enter-from,
.expand-leave-to {
  max-height: 0;
  opacity: 0;
  overflow: hidden;
}

.group-content {
  overflow: hidden;
}

.rotate-icon {
  transform: rotate(90deg);
  transition: transform 0.3s ease;
}

// Style untuk header group
.group-header {
  .v-icon {
    transition: transform 0.3s ease;
  }

  &:hover {
    background-color: rgba(255, 255, 255, 0.1);
  }
}

// Style untuk list items
.v-list-item {
  transition: all 0.3s ease;

  &.ml-4 {
    margin-left: 16px;
    position: relative;

    &::before {
      content: '';
      position: absolute;
      left: -8px;
      top: 0;
      bottom: 0;
      width: 2px;
      background-color: rgba(255, 255, 255, 0.1);
    }
  }
}
</style>
