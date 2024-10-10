
  <template>
      <v-app-bar
        color=""
        prominent
      >
        <v-app-bar-nav-icon  variant="text" @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
        <v-toolbar-title>Transparasi Publik</v-toolbar-title>
        <v-spacer></v-spacer>

      </v-app-bar>

      <v-navigation-drawer
        v-model="drawer"
      >
        <v-list
        >
        <v-list-item
          class="list"
          min-width="100"
          density="default"
          v-for="(links, i) in items"
          :key="i"
          :value="links"
          :to="links.route"
          :title="links.text"
          ></v-list-item>
      </v-list>
</v-navigation-drawer>
</template>
<script>
  export default {
    data: () => ({
      drawer: false,
      group: null,
      items: [
        {
          text: 'Home',
          icon: 'foo',
          route:'/dashboard'
        },
        {
          text: 'Data Transparasi',
          icon: 'foo',
          route:'/dashboard/datatransparasi'
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
