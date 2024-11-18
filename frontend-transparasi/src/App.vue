<template>
  <v-app>
    <navbar v-if="isAdminOrDprdRoute"></navbar>
    <navbar-program v-else-if="isProfile"></navbar-program>
    <v-main>
      <router-view/>
    </v-main>
  </v-app>
</template>

<script>
import { useRoute } from 'vue-router';
import { watch, onUnmounted, computed } from 'vue';
import Navbar from './components/navbar.vue';
import { navitemstore } from './store/navitem';
import NavbarProgram from './components/navbar-program.vue';

export default {
  components: { Navbar, NavbarProgram },
  setup() {
    const route = useRoute();
    const auth = navitemstore();
    auth.check();

    // Style handler
    const handleStyles = (apply = true) => {
      const style = apply ? 'hidden' : '';
      document.documentElement.style.overflowX = style;
      document.body.style.overflowX = style;

      // Tambahkan pengecekan null
      const vMain = document.querySelector('.v-main');
      if (vMain) {
        vMain.style.padding = apply ? '0' : '';
      }
    };

    // Watch route dengan setTimeout
    watch(
      () => route.path,
      (path) => {
        // Tunggu sampai DOM ter-render
        setTimeout(() => {
          handleStyles(path === '/home' || path.startsWith('/program'));
        }, 0);
      },
      { immediate: true }
    );

    onUnmounted(() => handleStyles(false));

    return {
      isAdminOrDprdRoute: computed(() => route.path.startsWith('/admin') || route.path.startsWith('/dprd')),
      isProfile: computed(() => route.path.startsWith('/profile'))
    };
  }
}
</script>

<style>
/* Style bisa dikosongkan karena sudah ditangani oleh JavaScript */
</style>
