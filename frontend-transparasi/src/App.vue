<template>
  <v-app>
    <navbar v-if="isAdminOrDprdRoute"></navbar>
    <v-main>
      <v-container fluid>
        <router-view/>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import { useRoute } from 'vue-router';
import { computed } from 'vue';
import Navbar from './components/navbar.vue';
import { navitemstore } from './store/navitem';

export default {
  components: {
    Navbar
  },
  setup() {
    const route = useRoute();
    const auth = navitemstore();

    auth.check();

    const isAdminOrDprdRoute = computed(() => {
      return route.path.startsWith('/admin') || route.path.startsWith('/dprd');
    });

    return {
      isAdminOrDprdRoute
    };
  }
}
</script>
