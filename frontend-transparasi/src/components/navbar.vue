
  <template>

      <v-navigation-drawer
        app
        permanent="true"
        class="custom-drawer"
      >
        <v-list
        >
        <v-list-item-title class="logo-wrap">
          <a href="#"
          target="_blank"
          >
            <img src="../assets/Logo-dprd.png" alt="" class="sidebar-logo mb-3">
          </a>
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
          active-class="active-item"
          ></v-list-item>
      </v-list>
</v-navigation-drawer>
<v-app-bar
  color=""
  prominent
>
  <v-spacer></v-spacer>

</v-app-bar>
</template>
<script>
  export default {
    data: () => ({
      drawer: false,
      group: null,
      items: [
        {
          text: 'Dashboard',
          icon: 'foo',
          route:'/dashboard'
        },
        {
          text: 'Data Transparasi',
          icon: 'foo',
          route:'/datatransparasi'
        },
        {
          text: 'Pengguna',
          icon: 'foo',
          route:'/pengguna'
        },
        {
          text: 'Dana',
          icon: 'foo',
          route:'/dana'
        },
        {
          text: 'Laporan',
          icon: 'foo',
          route:'/laporan'
        },

      ],
    }),

    watch: {
      group () {
        this.drawer = false
      },
    },
  }
</script>
<style scoped>
</style>
