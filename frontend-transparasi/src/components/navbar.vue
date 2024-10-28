<template>
  <v-navigation-drawer
    app
    v-model="drawer"
    class="custom-drawer"
  >
    <v-list>
      <v-list-item-title class="logo-wrap">
        <img src="../assets/Logo-dprd.png" alt="" class="sidebar-logo mb-3">
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
    <div class="mx-4">
      <v-icon class="mx-2" @click="handleLogout()">mdi-logout</v-icon>
      <v-icon class="mx-2">mdi-bell-outline</v-icon>
      <v-btn append-icon="mdi-chevron-down" class="mx-2" rounded="lg" color="black" variant="outlined" style="border-color: #BF3232;  text-transform: none; font-weight: 600;letter-spacing: 0.5px;">Admin</v-btn>
    </div>
  </v-app-bar>
</template>

<script>
import router from '@/router';
import { authStore } from '@/store/auth';
import { navitemstore } from '@/store/navitem';
import axios from 'axios';

export default {
  setup() {

  },
  data: () => ({
    drawer: true,
    items: [],
    auth:'',


  }),
  mounted(){
    const auth = authStore();
    const item = navitemstore();
    this.items = item.navitemlist;
    auth.check(router, this.$swal);
    item.check();
    this.handleManajemenUser();

  },
  methods:{
    handleManajemenUser(){
      var auth = navitemstore();
      this.auth = auth.stat;
    },
    handleLogout(){
      try{
        axios.post("/api/logout")
        .then((res)=>{
          const auth = authStore();
          auth.check(router, this.$swal);
          const item = navitemstore();
          item.reset();
          localStorage.removeItem('Role')
        })
      }catch(error){
        console.error("Error during logout:", error);
      }
    },
  }
}
</script>

<style scoped>
</style>
