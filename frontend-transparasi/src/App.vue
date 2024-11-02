<template>
  <v-app>
    <navbar v-if="isAdminOrDprdRoute"></navbar>
    <navbar-program v-else-if="isProfile"></navbar-program>
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
import NavbarProgram from './components/navbar-program.vue';
export default {
  components: {
    Navbar,
    NavbarProgram
  },
  setup() {
    const route = useRoute();
    const auth = navitemstore();

    auth.check();

    const isAdminOrDprdRoute = computed(() => {
      return route.path.startsWith('/admin') || route.path.startsWith('/dprd');
    });
    const isProfile = computed(() => {
      return route.path.startsWith('/profile');
    });

    return {
      isAdminOrDprdRoute,
      isProfile
    };
  }
}
</script>
