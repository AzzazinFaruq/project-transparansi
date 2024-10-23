
  <template>

      <v-navigation-drawer
        app
        v-model="drawer"
        class="custom-drawer"
      >
        <v-list
        >
        <v-list-item-title class="logo-wrap">
            <img src="../assets/Logo-dprd.png" alt="" class="sidebar-logo mb-3">
        </v-list-item-title>
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
      </v-list>
</v-navigation-drawer>
<v-app-bar
  color=""
  prominent
  elevation="0"
>
<div class="d-lg-none">
  <v-app-bar-nav-icon @click="drawer = !drawer">
    <img src="../assets/favicon.png" alt="" style="width: 40px;margin-left: 25px;margin-top: 5px;">
  </v-app-bar-nav-icon>
</div>
  <v-spacer></v-spacer>
  <div class="mx-4">
    <v-icon class="mx-2" @click="handleLogout()">mdi-logout</v-icon>
    <v-icon class="mx-2">mdi-bell</v-icon>
    <v-btn class="mx-2 " rounded="lg" color="black" variant="outlined"><b>admin</b></v-btn>
  </div>
</v-app-bar>
</template>
<script>
import router from '@/router';
import { authStore } from '@/store/auth';
import axios from 'axios';


  export default {
    data: () => ({
      drawer: true,
      items: [
        {
          text: 'Dashboard',
          icon: 'mdi-home',
          route:'/dashboard'
        },
        {
          text: 'Manajemen Pengguna',
          icon: 'mdi-account',
          route:'/manajemen-pengguna'
        },
        {
          text: 'Manajemen Program',
          icon: 'mdi-list-box',
          route:'/manajemen-program'
        },
        {
          text: 'Keluhan',
          icon: 'mdi-list-box',
          route:'/keluhan'
        },

      ],
    }),
    mounted(){
      const auth = authStore();
      auth.check(router,this.$swal)
    },
    methods:{
      handleLogout(){
        try{
          axios.post("/api/logout")
          .then((res)=>{
            const auth = authStore();
            auth.check(router,this.$swal)

          })
        }catch(error){

        }
      }
    }
  }
</script>
<style scoped>
</style>
