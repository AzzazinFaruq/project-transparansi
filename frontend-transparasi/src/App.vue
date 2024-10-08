<template>
  <v-app>
    <navbar v-if="!isLoginPage"  :key="$route.fullPath"></navbar>
    <v-main>
      <v-container fluid>
        <router-view/>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import axios from 'axios';
import navbar from './components/navbar.vue';
import { useRoute } from 'vue-router';
import { computed } from 'vue';
export default{
    setup(){
    const route = useRoute();
    const isLoginPage = computed(() => route.path === '/login');
    return { isLoginPage};
    }

}

</script>

<style>
.contain{
  bottom: 0;
  position: absolute;
}

</style>
